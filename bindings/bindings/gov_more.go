// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"morph-l2/bindings/solc"
)

const GovStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1017_storage\"},{\"astId\":1003,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1016_storage\"},{\"astId\":1005,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"batchBlockInterval\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"batchMaxBytes\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"batchTimeout\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_uint256\"},{\"astId\":1008,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"maxChunks\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"rollupEpoch\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_uint256\"},{\"astId\":1010,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"rollupEpochUpdateTime\",\"offset\":0,\"slot\":\"106\",\"type\":\"t_uint256\"},{\"astId\":1011,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"proposalInterval\",\"offset\":0,\"slot\":\"107\",\"type\":\"t_uint256\"},{\"astId\":1012,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"currentProposalID\",\"offset\":0,\"slot\":\"108\",\"type\":\"t_uint256\"},{\"astId\":1013,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"proposalData\",\"offset\":0,\"slot\":\"109\",\"type\":\"t_mapping(t_uint256,t_struct(ProposalData)1019_storage)\"},{\"astId\":1014,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"proposalInfos\",\"offset\":0,\"slot\":\"110\",\"type\":\"t_mapping(t_uint256,t_struct(ProposalInfo)1020_storage)\"},{\"astId\":1015,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"votes\",\"offset\":0,\"slot\":\"111\",\"type\":\"t_mapping(t_uint256,t_struct(AddressSet)1018_storage)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_bytes32)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"bytes32[]\",\"numberOfBytes\":\"32\"},\"t_array(t_uint256)1016_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1017_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_struct(AddressSet)1018_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e struct EnumerableSetUpgradeable.AddressSet)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_struct(AddressSet)1018_storage\"},\"t_mapping(t_uint256,t_struct(ProposalData)1019_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e struct IGov.ProposalData)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_struct(ProposalData)1019_storage\"},\"t_mapping(t_uint256,t_struct(ProposalInfo)1020_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e struct IGov.ProposalInfo)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_struct(ProposalInfo)1020_storage\"},\"t_struct(AddressSet)1018_storage\":{\"encoding\":\"inplace\",\"label\":\"struct EnumerableSetUpgradeable.AddressSet\",\"numberOfBytes\":\"64\"},\"t_struct(ProposalData)1019_storage\":{\"encoding\":\"inplace\",\"label\":\"struct IGov.ProposalData\",\"numberOfBytes\":\"160\"},\"t_struct(ProposalInfo)1020_storage\":{\"encoding\":\"inplace\",\"label\":\"struct IGov.ProposalInfo\",\"numberOfBytes\":\"64\"},\"t_struct(Set)1021_storage\":{\"encoding\":\"inplace\",\"label\":\"struct EnumerableSetUpgradeable.Set\",\"numberOfBytes\":\"64\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var GovStorageLayout = new(solc.StorageLayout)

var GovDeployedBin = "0x608060405234801561000f575f80fd5b5060043610610179575f3560e01c80638da5cb5b116100d2578063b511328d11610088578063e5aec99511610063578063e5aec9951461038e578063ecded2ae14610397578063f2fde38b146103aa575f80fd5b8063b511328d14610334578063bb881e4114610372578063de00d3fd1461037b575f80fd5b8063929a9cbe116100b8578063929a9cbe146102b457806396dea936146102bd5780639f50395214610321575f80fd5b80638da5cb5b1461026f5780638e21d5fb1461028d575f80fd5b806349c1a5811161013257806377c793801161010d57806377c7938014610211578063807de4431461021a5780638596305214610266575f80fd5b806349c1a581146101dd5780636396619014610200578063715018a614610209575f80fd5b80632d7aa82b116101625780632d7aa82b146101a55780634063a84e146101b85780634428c1a4146101d4575f80fd5b80630121b93f1461017d5780630d61b51914610192575b5f80fd5b61019061018b3660046118c7565b6103bd565b005b6101906101a03660046118c7565b6106c0565b6101906101b33660046118de565b6107cf565b6101c1606b5481565b6040519081526020015b60405180910390f35b6101c1606a5481565b6101f06101eb36600461193e565b610c91565b60405190151581526020016101cb565b6101c1606c5481565b610190610cb1565b6101c160675481565b6102417f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101cb565b6101c160655481565b60335473ffffffffffffffffffffffffffffffffffffffff16610241565b6102417f000000000000000000000000000000000000000000000000000000000000000081565b6101c160665481565b6102f96102cb3660046118c7565b606d6020525f9081526040902080546001820154600283015460038401546004909401549293919290919085565b604080519586526020860194909452928401919091526060830152608082015260a0016101cb565b61019061032f3660046118c7565b610cc4565b61035d6103423660046118c7565b606e6020525f90815260409020805460019091015460ff1682565b604080519283529015156020830152016101cb565b6101c160685481565b61019061038936600461196c565b610d88565b6101c160695481565b6101f06103a53660046118c7565b6110e6565b6101906103b8366004611982565b61112b565b5f73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016636d46e987336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401602060405180830381865afa158015610462573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061048691906119a4565b9050806104f4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f6f6e6c792073657175656e6365722063616e2070726f706f736500000000000060448201526064015b60405180910390fd5b5f828152606e6020526040902060010154829060ff1615610571576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f70726f706f73616c20616c726561647920617070726f7665640000000000000060448201526064016104eb565b5f818152606e60205260409020544211156105e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f70726f706f73616c206f7574206f66206461746500000000000000000000000060448201526064016104eb565b6105ff335f858152606f60205260409020906111e2565b1561068c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f73657175656e63657220616c726561647920766f746520666f7220746869732060448201527f70726f706f73616c00000000000000000000000000000000000000000000000060648201526084016104eb565b6106a3335f858152606f6020526040902090611210565b506106ad83611231565b156106bb576106bb83611368565b505050565b5f818152606e6020526040902060010154819060ff161561073d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f70726f706f73616c20616c726561647920617070726f7665640000000000000060448201526064016104eb565b5f818152606e60205260409020544211156107b4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f70726f706f73616c206f7574206f66206461746500000000000000000000000060448201526064016104eb565b6107bd82611231565b156107cb576107cb82611368565b5050565b5f54610100900460ff16158080156107ed57505f54600160ff909116105b806108065750303b15801561080657505f5460ff166001145b610892576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104eb565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156108ee575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b5f8711610957576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f696e76616c69642070726f706f73616c20696e74657276616c0000000000000060448201526064016104eb565b5f83116109c0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f696e76616c6964206d6178206368756e6b73000000000000000000000000000060448201526064016104eb565b5f8211610a29576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f696e76616c696420726f6c6c75702065706f636800000000000000000000000060448201526064016104eb565b85151580610a3657508415155b80610a4057508315155b610aa6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f696e76616c696420626174636820706172616d7300000000000000000000000060448201526064016104eb565b610aae611647565b606b8790556065869055606685905560678490556068839055606982905542606a55604080515f8152602081018990527f9e890086ea51933fb82fde9166ba4d58ecb0fdb81559ee03743b7ac052f43f7b910160405180910390a1604080515f8152602081018890527fa044538eba1b21d23eb13fa35811ca9d1d7ff9ef1c81ee4dc594fca08412531b910160405180910390a1604080515f8152602081018790527f11b7e0f5b30d2753fcf7151b7a907cc343034c6a7572d56c261ae00c411d16a7910160405180910390a1604080515f8152602081018690527fab2cb47d396c5d12c082ac9b6512d332af2767ca8e1fa5bcef40fa6970626569910160405180910390a1604080515f8152602081018590527fd4cf36ce0d0f667d929d7bdf98e8774da275ea7f990c012c308516650d85839a910160405180910390a1604080515f8152602081018490527f9b20ee151d057f4f3ece7fdf4ca1370cf143f181760e7712b722572f2dcba88f910160405180910390a18015610c88575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b5f828152606f60205260408120610ca890836111e2565b90505b92915050565b610cb96116e5565b610cc25f611766565b565b610ccc6116e5565b5f81118015610cdd5750606b548114155b610d43576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f696e76616c6964206e65772070726f706f73616c20696e74657276616c00000060448201526064016104eb565b606b80549082905560408051828152602081018490527f9e890086ea51933fb82fde9166ba4d58ecb0fdb81559ee03743b7ac052f43f7b910160405180910390a15050565b5f73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016636d46e987336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401602060405180830381865afa158015610e2d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e5191906119a4565b905080610eba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f6f6e6c792073657175656e6365722063616e2070726f706f736500000000000060448201526064016104eb565b81608001355f03610f27576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f696e76616c696420726f6c6c75702065706f636800000000000000000000000060448201526064016104eb565b5f826060013511610f94576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f696e76616c6964206d6178206368756e6b73000000000000000000000000000060448201526064016104eb565b8135151580610fa65750602082013515155b80610fb45750604082013515155b61101a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f696e76616c696420626174636820706172616d7300000000000000000000000060448201526064016104eb565b606c8054905f611029836119f0565b9091555050606c545f908152606d602052604090208290611075828281358155602082013560018201556040820135600282015560608201356003820155608082013560048201555050565b9050506040518060400160405280606b54426110919190611a27565b81525f6020918201819052606c548152606e82526040902082518155910151600190910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790555050565b5f818152606e602052604081206001015460ff161561110657505f919050565b5f828152606e602052604090205442111561112257505f919050565b610cab82611231565b6111336116e5565b73ffffffffffffffffffffffffffffffffffffffff81166111d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104eb565b6111df81611766565b50565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526001830160205260408120541515610ca8565b5f610ca88373ffffffffffffffffffffffffffffffffffffffff84166117dc565b5f807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166377d7dffb6040518163ffffffff1660e01b81526004015f60405180830381865afa15801561129b573d5f803e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526112e09190810190611a77565b90505f805b82518110156113455761132a83828151811061130357611303611b55565b6020026020010151606f5f8881526020019081526020015f206111e290919063ffffffff16565b1561133d5761133a826001611a27565b91505b6001016112e5565b506003825160026113569190611b82565b6113609190611b99565b109392505050565b5f818152606d6020526040902054606554146113d757606580545f838152606d60205260409081902054928390555190917fa044538eba1b21d23eb13fa35811ca9d1d7ff9ef1c81ee4dc594fca08412531b916113cd91848252602082015260400190565b60405180910390a1505b5f818152606d60205260409020600101546066541461144c57606680545f838152606d60205260409081902060010154928390555190917f11b7e0f5b30d2753fcf7151b7a907cc343034c6a7572d56c261ae00c411d16a79161144291848252602082015260400190565b60405180910390a1505b5f818152606d6020526040902060020154606754146114c157606780545f838152606d60205260409081902060020154928390555190917fab2cb47d396c5d12c082ac9b6512d332af2767ca8e1fa5bcef40fa6970626569916114b791848252602082015260400190565b60405180910390a1505b5f818152606d60205260409020600301546068541461153657606880545f838152606d60205260409081902060030154928390555190917fd4cf36ce0d0f667d929d7bdf98e8774da275ea7f990c012c308516650d85839a9161152c91848252602082015260400190565b60405180910390a1505b5f818152606d6020526040902060040154606954146115aa57606980545f838152606d6020908152604091829020600401805490945542606a55925481518381529384015290917f9b20ee151d057f4f3ece7fdf4ca1370cf143f181760e7712b722572f2dcba88f910160405180910390a1505b5f818152606e6020908152604091829020600190810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690911790556065546066546067546068546069548651948552948401929092528285015260608201526080810191909152905182917f146676d233683eb1ec2a813a7f97a7aa3241ae78af1ee6df4a4548c47178cbfa919081900360a00190a250565b5f54610100900460ff166116dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104eb565b610cc2611828565b60335473ffffffffffffffffffffffffffffffffffffffff163314610cc2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104eb565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f81815260018301602052604081205461182157508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610cab565b505f610cab565b5f54610100900460ff166118be576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104eb565b610cc233611766565b5f602082840312156118d7575f80fd5b5035919050565b5f805f805f8060c087890312156118f3575f80fd5b505084359660208601359650604086013595606081013595506080810135945060a0013592509050565b73ffffffffffffffffffffffffffffffffffffffff811681146111df575f80fd5b5f806040838503121561194f575f80fd5b8235915060208301356119618161191d565b809150509250929050565b5f60a0828403121561197c575f80fd5b50919050565b5f60208284031215611992575f80fd5b813561199d8161191d565b9392505050565b5f602082840312156119b4575f80fd5b8151801515811461199d575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611a2057611a206119c3565b5060010190565b80820180821115610cab57610cab6119c3565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b8051611a728161191d565b919050565b5f6020808385031215611a88575f80fd5b825167ffffffffffffffff80821115611a9f575f80fd5b818501915085601f830112611ab2575f80fd5b815181811115611ac457611ac4611a3a565b8060051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f83011681018181108582111715611b0757611b07611a3a565b604052918252848201925083810185019188831115611b24575f80fd5b938501935b82851015611b4957611b3a85611a67565b84529385019392850192611b29565b98975050505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b8082028115828204841417610cab57610cab6119c3565b5f82611bcc577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b50049056fea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(GovStorageLayoutJSON), GovStorageLayout); err != nil {
		panic(err)
	}

	layouts["Gov"] = GovStorageLayout
	deployedBytecodes["Gov"] = GovDeployedBin
}
