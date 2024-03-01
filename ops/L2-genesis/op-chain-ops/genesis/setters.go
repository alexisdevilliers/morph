package genesis

import (
	"fmt"
	"math/big"

	"github.com/morph-l2/bindings/bindings"
	"github.com/morph-l2/bindings/predeploys"
	"github.com/morph-l2/morph-deployer/op-chain-ops/immutables"
	"github.com/morph-l2/morph-deployer/op-chain-ops/state"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/vm"
	"github.com/scroll-tech/go-ethereum/log"
)

// UntouchableCodeHashes contains code hashes of all the contracts
// that should not be touched by the migration process.
type ChainHashMap map[uint64]common.Hash

var (
	// UntouchablePredeploys are addresses in the predeploy namespace
	// that should not be touched by the migration process.
	UntouchablePredeploys = map[common.Address]bool{
		predeploys.MorphStandardERC20Addr: true,
	}

	// UntouchableCodeHashes represent the bytecode hashes of contracts
	// that should not be touched by the migration process.
	UntouchableCodeHashes = map[common.Address]ChainHashMap{
		//predeploys.GovernanceTokenAddr: {
		//	1: common.HexToHash("0x8551d935f4e67ad3c98609f0d9f0f234740c4c4599f82674633b55204393e07f"),
		//	5: common.HexToHash("0xc4a213cf5f06418533e5168d8d82f7ccbcc97f27ab90197c2c051af6a4941cf9"),
		//},
		//predeploys.WETH9Addr: {
		//	1: common.HexToHash("0x779bbf2a738ef09d961c945116197e2ac764c1b39304b2b4418cd4e42668b173"),
		//	5: common.HexToHash("0x779bbf2a738ef09d961c945116197e2ac764c1b39304b2b4418cd4e42668b173"),
		//},
	}

	// FrozenStoragePredeploys represents the set of predeploys that
	// will not have their storage wiped during the migration process.
	// It is very explicitly set in its own mapping to ensure that
	// changes elsewhere in the codebase do no alter the predeploys
	// that do not have their storage wiped. It is safe for all other
	// predeploys to have their storage wiped.
	FrozenStoragePredeploys = map[common.Address]bool{
		//predeploys.GovernanceTokenAddr:     true,
		//predeploys.WETH9Addr:               true,
		//predeploys.LegacyMessagePasserAddr: true,
		//predeploys.LegacyERC20ETHAddr:      true,
		//predeploys.DeployerWhitelistAddr:   true,
	}
)

// FundDevAccounts will fund each of the development accounts.
func FundDevAccounts(db vm.StateDB) {
	for _, account := range DevAccounts {
		db.CreateAccount(account)
		db.AddBalance(account, devBalance)
	}
}

// SetL2Proxies will set each of the proxies in the state. It requires
// a Proxy and ProxyAdmin deployment present so that the Proxy bytecode
// can be set in state and the ProxyAdmin can be set as the admin of the
// Proxy.
func SetL2Proxies(db vm.StateDB) error {
	return setProxies(db, predeploys.ProxyAdminAddr, bigL2PredeployNamespace_53, 128)
}

func setProxies(db vm.StateDB, proxyAdminAddr common.Address, namespace *big.Int, count uint64) error {
	depBytecode, err := bindings.GetDeployedBytecode("TransparentUpgradeableProxy")
	if err != nil {
		return err
	}

	for i := uint64(0); i <= count; i++ {
		bigAddr := new(big.Int).Or(namespace, new(big.Int).SetUint64(i))
		addr := common.BigToAddress(bigAddr)
		if UntouchablePredeploys[addr] {
			log.Info("Skipping setting proxy", "address", addr)
			continue
		}

		if !db.Exist(addr) {
			db.CreateAccount(addr)
		}

		db.SetCode(addr, depBytecode)
		db.SetState(addr, AdminSlot, proxyAdminAddr.Hash())
		log.Trace("Set proxy", "address", addr, "admin", proxyAdminAddr)
	}

	return nil
}

// SetImplementations will set the implementations of the contracts in the state
// and configure the proxies to point to the implementations. It also sets
// the appropriate storage values for each contract at the proxy address.
func SetImplementations(db vm.StateDB, storage state.StorageConfig, immutable immutables.ImmutableConfig, imuConfig *immutables.Config) error {
	deployResults, slotResults, err := immutables.BuildMorph(immutable, imuConfig)
	if err != nil {
		return err
	}

	for name, address := range predeploys.Predeploys {
		if UntouchablePredeploys[*address] {
			err = SetTouchable(db, name, *address, storage, deployResults, slotResults)
			if err != nil {
				return err
			}
		} else {
			err = SetUntouchable(db, name, *address, storage, deployResults, slotResults)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func SetUntouchable(db vm.StateDB, name string, address common.Address, storage state.StorageConfig, deployResults immutables.DeploymentResults, slotResults immutables.SlotResults) error {
	codeAddr, err := AddressToCodeNamespace(address)
	if err != nil {
		return fmt.Errorf("error converting to code namespace: %w", err)
	}

	if !db.Exist(codeAddr) {
		db.CreateAccount(codeAddr)
	}

	db.SetState(address, ImplementationSlot, codeAddr.Hash())

	if err := setupPredeploy(db, deployResults, slotResults, storage, name, address, codeAddr); err != nil {
		return err
	}

	code := db.GetCode(codeAddr)
	if len(code) == 0 {
		return fmt.Errorf("code not set for %s", name)
	}
	return nil
}

func SetTouchable(db vm.StateDB, name string, address common.Address, storage state.StorageConfig, deployResults immutables.DeploymentResults, slotResults immutables.SlotResults) error {
	codeAddr := address
	if !db.Exist(codeAddr) {
		db.CreateAccount(codeAddr)
	}
	db.SetState(address, ImplementationSlot, codeAddr.Hash())
	if err := setupPredeploy(db, deployResults, slotResults, storage, name, address, codeAddr); err != nil {
		return err
	}
	code := db.GetCode(codeAddr)
	if len(code) == 0 {
		return fmt.Errorf("Untouchable predeploys code not set for %s", name)
	}
	return nil
}

// SetPrecompileBalances will set a single wei at each precompile address.
// This is an optimization to make calling them cheaper. This should only
// be used for devnets.
func SetPrecompileBalances(db vm.StateDB) {
	for i := 0; i < 256; i++ {
		addr := common.BytesToAddress([]byte{byte(i)})
		db.CreateAccount(addr)
		db.AddBalance(addr, common.Big1)
	}
}

func SetL2CrossDomainMessengerBalances(db *state.MemoryStateDB) {
	db.AddBalance(predeploys.L2CrossDomainMessengerAddr, lockedBalance)
	log.Info("Set balance to address of L2CrossDomainMessenger =>", "address:", predeploys.L2CrossDomainMessengerAddr, "balance:", lockedBalance)
}

func setupPredeploy(db vm.StateDB, deployResults immutables.DeploymentResults, slotResults immutables.SlotResults, storage state.StorageConfig, name string, proxyAddr common.Address, implAddr common.Address) error {
	// Use the generated bytecode when there are immutables
	// otherwise use the artifact deployed bytecode
	if bytecode, ok := deployResults[name]; ok {
		log.Info("Setting deployed bytecode with immutables", "name", name, "address", implAddr)
		db.SetCode(implAddr, bytecode)
	} else {
		depBytecode, err := bindings.GetDeployedBytecode(name)
		if err != nil {
			return err
		}
		log.Info("Setting deployed bytecode from solc compiler output", "name", name, "address", implAddr)
		db.SetCode(implAddr, depBytecode)
	}

	// Set the storage values
	if storageConfig, ok := storage[name]; ok {
		log.Info("Setting storage", "name", name, "address", proxyAddr)
		if err := state.SetStorage(name, proxyAddr, storageConfig, db); err != nil {
			return err
		}
	}

	if name == "L2Sequencer" || name == "Submitter" {
		// set slots directly
		if slots, ok := slotResults[name]; ok {
			for slotK, slotV := range slots {
				db.SetState(proxyAddr, slotK, slotV)
			}
		}
	}

	return nil
}