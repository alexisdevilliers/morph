package iface

import (
	"math/big"

	"github.com/morph-l2/bindings/bindings"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
)

type IRollup interface {
	LatestL2BlockNumber(opts *bind.CallOpts) (*big.Int, error)
	LastCommittedBatchIndex(opts *bind.CallOpts) (*big.Int, error)
	CommitBatch(opts *bind.TransactOpts, batchData bindings.IRollupBatchData, version *big.Int, sequencers []common.Address, signature []byte) (*types.Transaction, error)
	LastFinalizedBatchIndex(opts *bind.CallOpts) (*big.Int, error)
	FinalizeBatch(opts *bind.TransactOpts, _batchIndex *big.Int) (*types.Transaction, error)
	FINALIZATIONPERIODSECONDS(opts *bind.CallOpts) (*big.Int, error)
	CommittedBatchStores(opts *bind.CallOpts, arg0 *big.Int) (struct {
		BatchHash              [32]byte
		OriginTimestamp        *big.Int
		FinalizeTimestamp      *big.Int
		PrevStateRoot          [32]byte
		PostStateRoot          [32]byte
		WithdrawalRoot         [32]byte
		DataHash               [32]byte
		L1MessagePopped        *big.Int
		TotalL1MessagePopped   *big.Int
		SkippedL1MessageBitmap []byte
		BlockNumber            *big.Int
		BlobVersionedHash      [32]byte
	}, error)
	BatchInsideChallengeWindow(opts *bind.CallOpts, batchIndex *big.Int) (bool, error)
}

// IL2Submitter is the interface for the submitter on L2
type IL2Submitter interface {
	GetCurrentSubmitter(opts *bind.CallOpts) (common.Address, *big.Int, *big.Int, error)
}

// IL2Sequencer is the interface for the sequencer on L2
type IL2Sequencer interface {
	InSequencersSet(opts *bind.CallOpts, previous bool, checkAddr common.Address) (bool, *big.Int, error)
}
