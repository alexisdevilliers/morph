package services

import (
	"sort"
	"sync"
	"time"

	"morph-l2/bindings/bindings"
	"morph-l2/tx-submitter/utils"

	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
)

type TxInfo struct {
	sendTime uint64
	tx       types.Transaction

	queryTimes uint64
}

type PendingTxs struct {
	mu sync.Mutex

	txinfos map[common.Hash]TxInfo
	pnonce  uint64 // pending nonce

	failedIndex *uint64
	pindex      uint64 // pending batch index

	pfinalize       uint64
	commitBatchId   []byte
	finalizeBatchId []byte
}

func NewPendingTxs(cid []byte, fid []byte) *PendingTxs {
	return &PendingTxs{
		txinfos:         make(map[common.Hash]TxInfo),
		commitBatchId:   cid,
		finalizeBatchId: fid,
	}
}

func (pt *PendingTxs) Add(tx types.Transaction) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.txinfos[tx.Hash()] = TxInfo{
		sendTime: uint64(time.Now().Unix()),
		tx:       tx,
	}
}

func (pt *PendingTxs) Remove(txHash common.Hash) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	delete(pt.txinfos, txHash)
}

func (pt *PendingTxs) Recover(txs []types.Transaction, a *abi.ABI) {
	// restore state from mempool
	if len(txs) > 0 {
		var pbindex, pfindex uint64

		for _, tx := range txs {
			pt.Add(tx)

			method := utils.ParseMethod(tx, a)
			if method == "commitBatch" {

				index := utils.ParseParentBatchIndex(tx.Data())
				if index > pbindex {
					pbindex = index
				}
			} else if method == "finalizeBatch" {
				findex := utils.ParseFBatchIndex(tx.Data())
				if findex > pfindex {
					pfindex = findex
				}
			}
		}

		pt.SetPindex(pbindex)
		pt.SetPFinalize(pfindex)
		pt.SetNonce(txs[len(txs)-1].Nonce())
	}
}

func (pt *PendingTxs) GetAll() []TxInfo {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	// copy txs and return
	txs := make([]TxInfo, 0, len(pt.txinfos))
	for _, tx := range pt.txinfos {
		txs = append(txs, tx)
	}

	// sort by nonce
	sort.SliceStable(txs, func(i, j int) bool {
		return txs[i].tx.Nonce() < txs[j].tx.Nonce()
	})

	return txs
}

func (pt *PendingTxs) Get(txHash common.Hash) (TxInfo, bool) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	tx, ok := pt.txinfos[txHash]
	return tx, ok
}

func (pt *PendingTxs) IncQueryTimes(txHash common.Hash) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.txinfos[txHash] = TxInfo{tx: pt.txinfos[txHash].tx, queryTimes: pt.txinfos[txHash].queryTimes + 1, sendTime: pt.txinfos[txHash].sendTime}
}

func (pt *PendingTxs) SetFailedStatus(index uint64) {
	pt.mu.Lock()
	defer pt.mu.Unlock()

	// failed index must be less than pindex
	if pt.failedIndex != nil || index >= pt.pindex {
		return
	}

	pt.failedIndex = &index
}
func (pt *PendingTxs) SetPindex(index uint64) {
	pt.mu.Lock()
	defer pt.mu.Unlock()

	pt.pindex = index
}

func (pt *PendingTxs) SetNonce(nonce uint64) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.pnonce = nonce
}

func (pt *PendingTxs) SetPFinalize(finalize uint64) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.pfinalize = finalize
}

func (pt *PendingTxs) RemoveRollupRestriction() {
	pt.mu.Lock()
	defer pt.mu.Unlock()

	pt.failedIndex = nil
}

func (pt *PendingTxs) HaveFailed() bool {
	return pt.failedIndex != nil
}

func (pt *PendingTxs) ExistedIndex(index uint64) bool {

	txs := pt.GetAll()

	abi, _ := bindings.RollupMetaData.GetAbi()
	pt.mu.Lock()
	defer pt.mu.Unlock()
	for i := len(txs) - 1; i >= 0; i-- {
		tx := txs[i].tx
		if utils.ParseMethod(tx, abi) == "commitBatch" {
			pindex := utils.ParseParentBatchIndex(tx.Data()) + 1
			if index == pindex {
				return true
			}

		}

	}
	return false

}

func (pt *PendingTxs) ResetFailedIndex(index uint64) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.failedIndex = &index
}