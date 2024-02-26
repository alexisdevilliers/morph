// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const L1CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1025_storage\"},{\"astId\":1003,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1024_storage\"},{\"astId\":1005,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_paused\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":1006,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)1024_storage\"},{\"astId\":1007,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_status\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1008,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)1024_storage\"},{\"astId\":1009,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"xDomainMessageSender\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_address\"},{\"astId\":1010,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"counterpart\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_address\"},{\"astId\":1011,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"feeVault\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_address\"},{\"astId\":1012,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__rateLimiter\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":1013,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_array(t_uint256)1023_storage\"},{\"astId\":1014,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"provenWithdrawals\",\"offset\":0,\"slot\":\"251\",\"type\":\"t_mapping(t_bytes32,t_struct(ProvenWithdrawal)1026_storage)\"},{\"astId\":1015,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"finalizedWithdrawals\",\"offset\":0,\"slot\":\"252\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1016,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"messageSendTimestamp\",\"offset\":0,\"slot\":\"253\",\"type\":\"t_mapping(t_bytes32,t_uint256)\"},{\"astId\":1017,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"isL1MessageDropped\",\"offset\":0,\"slot\":\"254\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1018,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"rollup\",\"offset\":0,\"slot\":\"255\",\"type\":\"t_address\"},{\"astId\":1019,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"messageQueue\",\"offset\":0,\"slot\":\"256\",\"type\":\"t_address\"},{\"astId\":1020,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"maxReplayTimes\",\"offset\":0,\"slot\":\"257\",\"type\":\"t_uint256\"},{\"astId\":1021,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"replayStates\",\"offset\":0,\"slot\":\"258\",\"type\":\"t_mapping(t_bytes32,t_struct(ReplayState)1027_storage)\"},{\"astId\":1022,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"prevReplayIndex\",\"offset\":0,\"slot\":\"259\",\"type\":\"t_mapping(t_uint256,t_uint256)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1023_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[46]\",\"numberOfBytes\":\"1472\"},\"t_array(t_uint256)1024_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1025_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_mapping(t_bytes32,t_struct(ProvenWithdrawal)1026_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e struct L1CrossDomainMessenger.ProvenWithdrawal)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_struct(ProvenWithdrawal)1026_storage\"},\"t_mapping(t_bytes32,t_struct(ReplayState)1027_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e struct L1CrossDomainMessenger.ReplayState)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_struct(ReplayState)1027_storage\"},\"t_mapping(t_bytes32,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_uint256\"},\"t_struct(ProvenWithdrawal)1026_storage\":{\"encoding\":\"inplace\",\"label\":\"struct L1CrossDomainMessenger.ProvenWithdrawal\",\"numberOfBytes\":\"96\"},\"t_struct(ReplayState)1027_storage\":{\"encoding\":\"inplace\",\"label\":\"struct L1CrossDomainMessenger.ReplayState\",\"numberOfBytes\":\"32\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L1CrossDomainMessengerDeployedBin = "0x6080604052600436106101b2575f3560e01c80638da5cb5b116100e7578063c0c53b8b11610087578063e965084c11610062578063e965084c14610586578063ea7ec514146105db578063ecc7042814610607578063f2fde38b1461061b575f80fd5b8063c0c53b8b14610510578063cb23bcb51461052f578063e70fc93b1461055b575f80fd5b8063a14238e7116100c2578063a14238e714610482578063b2267a7b146104b0578063b604bf4c146104c3578063bedb86fb146104f1575f80fd5b80638da5cb5b146104155780638ef1332e1461043f578063946130d81461045e575f80fd5b80635c975abb11610152578063715018a61161012d578063715018a61461033757806371a4caab1461034b578063797594b01461036a578063846d4d7a14610396575f80fd5b80635c975abb146102e15780635f7b1577146102f85780636e296e451461030b575f80fd5b80633b70c18a1161018d5780633b70c18a14610231578063407c195514610283578063478222c2146102a257806355004105146102ce575f80fd5b806329907acd146101bd5780632a6cccb2146101de578063340735f7146101fd575f80fd5b366101b957005b5f80fd5b3480156101c8575f80fd5b506101dc6101d73660046132da565b61063a565b005b3480156101e9575f80fd5b506101dc6101f8366004613347565b610a3b565b348015610208575f80fd5b5061021c61021736600461337e565b610aca565b60405190151581526020015b60405180910390f35b34801561023c575f80fd5b506101005461025e9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610228565b34801561028e575f80fd5b506101dc61029d3660046133ba565b610b93565b3480156102ad575f80fd5b5060cb5461025e9073ffffffffffffffffffffffffffffffffffffffff1681565b6101dc6102dc3660046133d1565b610bda565b3480156102ec575f80fd5b5060655460ff1661021c565b6101dc610306366004613468565b6112d6565b348015610316575f80fd5b5060c95461025e9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610342575f80fd5b506101dc61132a565b348015610356575f80fd5b506101dc610365366004613506565b61133d565b348015610375575f80fd5b5060ca5461025e9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156103a1575f80fd5b506103ec6103b03660046133ba565b6101026020525f90815260409020546fffffffffffffffffffffffffffffffff8082169170010000000000000000000000000000000090041682565b604080516fffffffffffffffffffffffffffffffff938416815292909116602083015201610228565b348015610420575f80fd5b5060335473ffffffffffffffffffffffffffffffffffffffff1661025e565b34801561044a575f80fd5b506101dc6104593660046132da565b61184e565b348015610469575f80fd5b506104746101015481565b604051908152602001610228565b34801561048d575f80fd5b5061021c61049c3660046133ba565b60fc6020525f908152604090205460ff1681565b6101dc6104be36600461358f565b611e8a565b3480156104ce575f80fd5b5061021c6104dd3660046133ba565b60fe6020525f908152604090205460ff1681565b3480156104fc575f80fd5b506101dc61050b3660046135ea565b611ea5565b34801561051b575f80fd5b506101dc61052a366004613609565b611ec6565b34801561053a575f80fd5b5060ff5461025e9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610566575f80fd5b506104746105753660046133ba565b60fd6020525f908152604090205481565b348015610591575f80fd5b506105c06105a03660046133ba565b60fb6020525f908152604090208054600182015460029092015490919083565b60408051938452602084019290925290820152606001610228565b3480156105e6575f80fd5b506104746105f53660046133ba565b6101036020525f908152604090205481565b348015610612575f80fd5b50610474612197565b348015610626575f80fd5b506101dc610635366004613347565b61222e565b6106426122e2565b60c95473ffffffffffffffffffffffffffffffffffffffff1661dead146106ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4d65737361676520697320616c726561647920696e20657865637574696f6e0060448201526064015b60405180910390fd5b6101005473ffffffffffffffffffffffffffffffffffffffff165f6106f2878787878761234f565b90505f818051906020012090505f60fd5f8381526020019081526020015f20541161079f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f50726f7669646564206d65737361676520686173206e6f74206265656e20656e60448201527f717565756564000000000000000000000000000000000000000000000000000060648201526084016106c1565b5f81815260fe602052604090205460ff1615610817576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f4d65737361676520616c72656164792064726f7070656400000000000000000060448201526064016106c1565b5f818152610102602052604081205470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16908190036108575750845b6040517f916524610000000000000000000000000000000000000000000000000000000081526004810182905273ffffffffffffffffffffffffffffffffffffffff8516906391652461906024015f604051808303815f87803b1580156108bc575f80fd5b505af11580156108ce573d5f803e3d5ffd5b5050505f918252506101036020526040902054801561090e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610857565b5f82815260fe60205260409081902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905560c980547fffffffffffffffffffffffff000000000000000000000000000000000000000016736f297c61b5c92ef107ffd30cd56affe5a273e841179055517f14298c5100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a16906314298c519089906109d69089906004016136b4565b5f604051808303818588803b1580156109ed575f80fd5b505af11580156109ff573d5f803e3d5ffd5b505060c980547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055505050505050505050505050565b610a436123eb565b60cb805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f591015b60405180910390a15050565b5f84815b6020811015610b87578085901c600116600103610b3457858160208110610af757610af76136c6565b602002013582604051602001610b17929190918252602082015260400190565b604051602081830303815290604052805190602001209150610b7f565b81868260208110610b4757610b476136c6565b6020020135604051602001610b66929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b600101610ace565b50909114949350505050565b610b9b6123eb565b61010180549082905560408051828152602081018490527fd700562df02eb66951f6f5275df7ebd7c0ec58b3422915789b3b1877aab2e52b9101610abe565b610be26122e2565b60c95473ffffffffffffffffffffffffffffffffffffffff1661dead14610c65576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4d65737361676520697320616c726561647920696e20657865637574696f6e0060448201526064016106c1565b6101005460ca5473ffffffffffffffffffffffffffffffffffffffff91821691165f610c948a8a8a8a8a61234f565b90505f818051906020012090505f60fd5f8381526020019081526020015f205411610d41576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f50726f7669646564206d65737361676520686173206e6f74206265656e20656e60448201527f717565756564000000000000000000000000000000000000000000000000000060648201526084016106c1565b5f81815260fe602052604090205460ff1615610db9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f4d65737361676520616c72656164792064726f7070656400000000000000000060448201526064016106c1565b6040517fd7704bae00000000000000000000000000000000000000000000000000000000815263ffffffff871660048201525f9073ffffffffffffffffffffffffffffffffffffffff86169063d7704bae90602401602060405180830381865afa158015610e29573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e4d91906136f3565b905080341015610eb9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f496e73756666696369656e74206d73672e76616c756520666f7220666565000060448201526064016106c1565b8015610f885760cb546040515f9173ffffffffffffffffffffffffffffffffffffffff169083908381818185875af1925050503d805f8114610f16576040519150601f19603f3d011682016040523d82523d5f602084013e610f1b565b606091505b5050905080610f86576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4661696c656420746f206465647563742074686520666565000000000000000060448201526064016106c1565b505b5f8573ffffffffffffffffffffffffffffffffffffffff1663fd0ad31e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fd2573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ff691906136f3565b6040517f9b15978200000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff871690639b1597829061104f9088908c90899060040161370a565b5f604051808303815f87803b158015611066575f80fd5b505af1158015611078573d5f803e3d5ffd5b5050505f848152610102602090815260408083208151808301909252546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000090910416918101829052925090036110e8575f8281526101036020526040902060018c019055611119565b80602001516001016fffffffffffffffffffffffffffffffff166101035f8481526020019081526020015f20819055505b6fffffffffffffffffffffffffffffffff8083166020830152610101548251909116106111a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f457863656564206d6178696d756d207265706c61792074696d6573000000000060448201526064016106c1565b80516fffffffffffffffffffffffffffffffff600191909101811682525f8581526101026020908152604090912083519184015183167001000000000000000000000000000000000291909216179055348381039084146112c5575f8973ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f8114611253576040519150601f19603f3d011682016040523d82523d5f602084013e611258565b606091505b50509050806112c3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4661696c656420746f20726566756e642074686520666565000000000000000060448201526064016106c1565b505b505050505050505050505050505050565b6112de6122e2565b611322868686868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525088925087915061246c9050565b505050505050565b6113326123eb565b61133b5f612969565b565b6113456122e2565b60c95473ffffffffffffffffffffffffffffffffffffffff1661dead146113c8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4d65737361676520697320616c726561647920696e20657865637574696f6e0060448201526064016106c1565b6101005473ffffffffffffffffffffffffffffffffffffffff90811690871603611474576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f4d657373656e6765723a20466f7262696420746f2063616c6c206d657373616760448201527f652071756575650000000000000000000000000000000000000000000000000060648201526084016106c1565b61147d866129df565b60c95473ffffffffffffffffffffffffffffffffffffffff90811690881603611528576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f4d657373656e6765723a20496e76616c6964206d6573736167652073656e646560448201527f720000000000000000000000000000000000000000000000000000000000000060648201526084016106c1565b5f611536888888888861234f565b80516020918201205f81815260fb83526040908190208151606081018352815481526001820154948101859052600290910154918101919091529092509015611601576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4d657373656e6765723a207769746864726177616c206861736820686173206160448201527f6c7265616479206265656e2070726f76656e000000000000000000000000000060648201526084016106c1565b60ff546040517f04d772150000000000000000000000000000000000000000000000000000000081526004810185905273ffffffffffffffffffffffffffffffffffffffff909116905f9082906304d7721590602401602060405180830381865afa158015611672573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061169691906136f3565b90505f8111611727576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f4d657373656e6765723a20646f206e6f74207375626d6974207769746864726160448201527f77616c526f6f740000000000000000000000000000000000000000000000000060648201526084016106c1565b61173384878a88610aca565b6117bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4d657373656e6765723a20696e76616c6964207769746864726177616c20696e60448201527f636c7573696f6e2070726f6f660000000000000000000000000000000000000060648201526084016106c1565b604080516060810182528681524260208083019182528284018c81525f89815260fb90925284822093518455915160018401559051600290920191909155905173ffffffffffffffffffffffffffffffffffffffff808d1692908e169187917f67a6208cfcc0801d50f6cbe764733f4fddf66ac0b04442061a8a8c0cb6b63f6291a45050505050505050505050565b6118566122e2565b60c95473ffffffffffffffffffffffffffffffffffffffff1661dead146118d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4d65737361676520697320616c726561647920696e20657865637574696f6e0060448201526064016106c1565b5f6118e7868686868661234f565b80516020918201205f81815260fc90925260409091205490915060ff1615611991576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4d657373656e6765723a207769746864726177616c2068617320616c7265616460448201527f79206265656e2066696e616c697a65640000000000000000000000000000000060648201526084016106c1565b5f81815260fb602090815260408083208151606081018352815481526001820154938101849052600290910154918101919091529103611a53576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4d657373656e6765723a207769746864726177616c20686173206e6f7420626560448201527f656e2070726f76656e207965740000000000000000000000000000000000000060648201526084016106c1565b611a608160200151612a5e565b611aee57604080517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260248101919091527f4d657373656e6765723a2070726f76656e207769746864726177616c2066696e60448201527f616c697a6174696f6e20706572696f6420686173206e6f7420656c617073656460648201526084016106c1565b60ff5481516040517f04d7721500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909216915f9183916304d7721591611b4c9160040190815260200190565b602060405180830381865afa158015611b67573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b8b91906136f3565b90505f8111611c1c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f4d657373656e6765723a20646f206e6f74207375626d6974207769746864726160448201527f77616c526f6f740000000000000000000000000000000000000000000000000060648201526084016106c1565b6040517f2571098d000000000000000000000000000000000000000000000000000000008152600481018290525f9073ffffffffffffffffffffffffffffffffffffffff841690632571098d90602401602060405180830381865afa158015611c87573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611cab91906136f3565b905080611d14576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4d657373656e6765723a206261746368206e6f7420766572696669656400000060448201526064016106c1565b5050508660c95f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f8673ffffffffffffffffffffffffffffffffffffffff168685604051611d7e919061374d565b5f6040518083038185875af1925050503d805f8114611db8576040519150601f19603f3d011682016040523d82523d5f602084013e611dbd565b606091505b505060c980547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead17905590508015611e55575f83815260fc602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555184917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a2611e80565b60405183907f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f905f90a25b5050505050505050565b611e926122e2565b611e9f848484843361246c565b50505050565b611ead6123eb565b8015611ebe57611ebb612b01565b50565b611ebb612b86565b5f54610100900460ff1615808015611ee457505f54600160ff909116105b80611efd5750303b158015611efd57505f5460ff166001145b611f89576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016106c1565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015611fe5575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff8316158061201c575073ffffffffffffffffffffffffffffffffffffffff8216155b15612053576040517fecc6fdf000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61207173530000000000000000000000000000000000000185612bdd565b60ff805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925561010080549285169282169290921790915560ca805490911673530000000000000000000000000000000000000717905560036101018190556040517fd700562df02eb66951f6f5275df7ebd7c0ec58b3422915789b3b1877aab2e52b91612128915f9190918252602082015260400190565b60405180910390a18015611e9f575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b61010054604080517ffd0ad31e00000000000000000000000000000000000000000000000000000000815290515f9273ffffffffffffffffffffffffffffffffffffffff169163fd0ad31e9160048083019260209291908290030181865afa158015612205573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061222991906136f3565b905090565b6122366123eb565b73ffffffffffffffffffffffffffffffffffffffff81166122d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016106c1565b611ebb81612969565b60655460ff161561133b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016106c1565b6060858585858560405160240161236a959493929190613768565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f8ef1332e00000000000000000000000000000000000000000000000000000000179052905095945050505050565b60335473ffffffffffffffffffffffffffffffffffffffff16331461133b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016106c1565b612474612d29565b6101005460ca54604080517ffd0ad31e000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff93841693909216915f91849163fd0ad31e916004808201926020929091908290030181865afa1580156124ee573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061251291906136f3565b90505f612522338a8a858b61234f565b6040517fd7704bae000000000000000000000000000000000000000000000000000000008152600481018890529091505f9073ffffffffffffffffffffffffffffffffffffffff86169063d7704bae90602401602060405180830381865afa158015612590573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906125b491906136f3565b90506125c089826137b7565b341015612629576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f496e73756666696369656e74206d73672e76616c75650000000000000000000060448201526064016106c1565b80156126f85760cb546040515f9173ffffffffffffffffffffffffffffffffffffffff169083908381818185875af1925050503d805f8114612686576040519150601f19603f3d011682016040523d82523d5f602084013e61268b565b606091505b50509050806126f6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4661696c656420746f206465647563742074686520666565000000000000000060448201526064016106c1565b505b6040517f9b15978200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861690639b1597829061274e9087908b9087906004016137ef565b5f604051808303815f87803b158015612765575f80fd5b505af1158015612777573d5f803e3d5ffd5b505050505f8280519060200120905060fd5f8281526020019081526020015f20545f14612800576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f4475706c696361746564206d657373616765000000000000000000000000000060448201526064016106c1565b5f81815260fd6020526040902042905573ffffffffffffffffffffffffffffffffffffffff8b163373ffffffffffffffffffffffffffffffffffffffff167f104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e8c878c8e6040516128739493929190613823565b60405180910390a3348290038a8103908b14612951575f8873ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f81146128df576040519150601f19603f3d011682016040523d82523d5f602084013e6128e4565b606091505b505090508061294f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4661696c656420746f20726566756e642074686520666565000000000000000060448201526064016106c1565b505b505050505050506129626001609755565b5050505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b3073ffffffffffffffffffffffffffffffffffffffff821603611ebb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f4d657373656e6765723a20466f7262696420746f2063616c6c2073656c66000060448201526064016106c1565b60ff54604080517ff4daa29100000000000000000000000000000000000000000000000000000000815290515f9273ffffffffffffffffffffffffffffffffffffffff169163f4daa2919160048083019260209291908290030181865afa158015612acb573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612aef91906136f3565b612af990836137b7565b421192915050565b612b096122e2565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612b5c3390565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b612b8e612da3565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33612b5c565b5f54610100900460ff16612c73576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106c1565b612c7b612e0f565b612c83612ead565b612c8b612f4b565b60c980547fffffffffffffffffffffffff000000000000000000000000000000000000000090811661dead1790915560ca805473ffffffffffffffffffffffffffffffffffffffff85811691909316179055811615612d255760cb80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5050565b600260975403612d95576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016106c1565b6002609755565b6001609755565b60655460ff1661133b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016106c1565b5f54610100900460ff16612ea5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106c1565b61133b612fe9565b5f54610100900460ff16612f43576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106c1565b61133b613088565b5f54610100900460ff16612fe1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106c1565b61133b613148565b5f54610100900460ff1661307f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106c1565b61133b33612969565b5f54610100900460ff1661311e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106c1565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b5f54610100900460ff16612d9c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106c1565b803573ffffffffffffffffffffffffffffffffffffffff81168114613201575f80fd5b919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112613242575f80fd5b813567ffffffffffffffff8082111561325d5761325d613206565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019082821181831017156132a3576132a3613206565b816040528381528660208588010111156132bb575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f805f805f60a086880312156132ee575f80fd5b6132f7866131de565b9450613305602087016131de565b93506040860135925060608601359150608086013567ffffffffffffffff81111561332e575f80fd5b61333a88828901613233565b9150509295509295909350565b5f60208284031215613357575f80fd5b613360826131de565b9392505050565b806104008101831015613378575f80fd5b92915050565b5f805f806104608587031215613392575f80fd5b843593506133a38660208701613367565b939693955050505061042082013591610440013590565b5f602082840312156133ca575f80fd5b5035919050565b5f805f805f805f60e0888a0312156133e7575f80fd5b6133f0886131de565b96506133fe602089016131de565b95506040880135945060608801359350608088013567ffffffffffffffff811115613427575f80fd5b6134338a828b01613233565b93505060a088013563ffffffff8116811461344c575f80fd5b915061345a60c089016131de565b905092959891949750929550565b5f805f805f8060a0878903121561347d575f80fd5b613486876131de565b955060208701359450604087013567ffffffffffffffff808211156134a9575f80fd5b818901915089601f8301126134bc575f80fd5b8135818111156134ca575f80fd5b8a60208285010111156134db575f80fd5b602083019650809550505050606087013591506134fa608088016131de565b90509295509295509295565b5f805f805f805f6104c0888a03121561351d575f80fd5b613526886131de565b9650613534602089016131de565b95506040880135945060608801359350608088013567ffffffffffffffff81111561355d575f80fd5b6135698a828b01613233565b9350506135798960a08a01613367565b91506104a0880135905092959891949750929550565b5f805f80608085870312156135a2575f80fd5b6135ab856131de565b935060208501359250604085013567ffffffffffffffff8111156135cd575f80fd5b6135d987828801613233565b949793965093946060013593505050565b5f602082840312156135fa575f80fd5b81358015158114613360575f80fd5b5f805f6060848603121561361b575f80fd5b613624846131de565b9250613632602085016131de565b9150613640604085016131de565b90509250925092565b5f5b8381101561366357818101518382015260200161364b565b50505f910152565b5f8151808452613682816020860160208601613649565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081525f613360602083018461366b565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f60208284031215613703575f80fd5b5051919050565b73ffffffffffffffffffffffffffffffffffffffff8416815263ffffffff83166020820152606060408201525f613744606083018461366b565b95945050505050565b5f825161375e818460208701613649565b9190910192915050565b5f73ffffffffffffffffffffffffffffffffffffffff808816835280871660208401525084604083015283606083015260a060808301526137ac60a083018461366b565b979650505050505050565b80820180821115613378577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201525f613744606083018461366b565b848152836020820152826040820152608060608201525f613847608083018461366b565b969550505050505056fea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(L1CrossDomainMessengerStorageLayoutJSON), L1CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1CrossDomainMessenger"] = L1CrossDomainMessengerStorageLayout
	deployedBytecodes["L1CrossDomainMessenger"] = L1CrossDomainMessengerDeployedBin
}
