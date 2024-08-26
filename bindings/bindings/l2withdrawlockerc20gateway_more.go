// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"morph-l2/bindings/solc"
)

const L2WithdrawLockERC20GatewayStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/l2/gateways/L2WithdrawLockERC20Gateway.sol:L2WithdrawLockERC20Gateway\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/l2/gateways/L2WithdrawLockERC20Gateway.sol:L2WithdrawLockERC20Gateway\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/l2/gateways/L2WithdrawLockERC20Gateway.sol:L2WithdrawLockERC20Gateway\",\"label\":\"_status\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"contracts/l2/gateways/L2WithdrawLockERC20Gateway.sol:L2WithdrawLockERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_array(t_uint256)1015_storage\"},{\"astId\":1004,\"contract\":\"contracts/l2/gateways/L2WithdrawLockERC20Gateway.sol:L2WithdrawLockERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_array(t_uint256)1016_storage\"},{\"astId\":1005,\"contract\":\"contracts/l2/gateways/L2WithdrawLockERC20Gateway.sol:L2WithdrawLockERC20Gateway\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_address\"},{\"astId\":1006,\"contract\":\"contracts/l2/gateways/L2WithdrawLockERC20Gateway.sol:L2WithdrawLockERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)1015_storage\"},{\"astId\":1007,\"contract\":\"contracts/l2/gateways/L2WithdrawLockERC20Gateway.sol:L2WithdrawLockERC20Gateway\",\"label\":\"counterpart\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_address\"},{\"astId\":1008,\"contract\":\"contracts/l2/gateways/L2WithdrawLockERC20Gateway.sol:L2WithdrawLockERC20Gateway\",\"label\":\"router\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_address\"},{\"astId\":1009,\"contract\":\"contracts/l2/gateways/L2WithdrawLockERC20Gateway.sol:L2WithdrawLockERC20Gateway\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"153\",\"type\":\"t_address\"},{\"astId\":1010,\"contract\":\"contracts/l2/gateways/L2WithdrawLockERC20Gateway.sol:L2WithdrawLockERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"154\",\"type\":\"t_array(t_uint256)1014_storage\"},{\"astId\":1011,\"contract\":\"contracts/l2/gateways/L2WithdrawLockERC20Gateway.sol:L2WithdrawLockERC20Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"200\",\"type\":\"t_array(t_uint256)1016_storage\"},{\"astId\":1012,\"contract\":\"contracts/l2/gateways/L2WithdrawLockERC20Gateway.sol:L2WithdrawLockERC20Gateway\",\"label\":\"tokenMapping\",\"offset\":0,\"slot\":\"250\",\"type\":\"t_mapping(t_address,t_address)\"},{\"astId\":1013,\"contract\":\"contracts/l2/gateways/L2WithdrawLockERC20Gateway.sol:L2WithdrawLockERC20Gateway\",\"label\":\"withdrawLock\",\"offset\":0,\"slot\":\"251\",\"type\":\"t_mapping(t_address,t_bool)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1014_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[46]\",\"numberOfBytes\":\"1472\"},\"t_array(t_uint256)1015_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1016_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_address)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e address)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_address\"},\"t_mapping(t_address,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_bool\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2WithdrawLockERC20GatewayStorageLayout = new(solc.StorageLayout)

var L2WithdrawLockERC20GatewayDeployedBin = "0x608060405260043610610109575f3560e01c8063a93a4af9116100a1578063cdd0da7c11610071578063f2fde38b11610057578063f2fde38b1461033b578063f887ea401461035a578063fac752eb14610386575f80fd5b8063cdd0da7c146102de578063ebc137d01461031c575f80fd5b8063a93a4af91461024c578063ba27f50b1461025f578063c0c53b8b146102a0578063c676ad29146102bf575f80fd5b8063715018a6116100dc578063715018a6146101cf578063797594b0146101e35780638431f5c11461020f5780638da5cb5b14610222575f80fd5b80633cb747bf1461010d57806354bbd59c14610163578063575361b6146101a75780636c07ea43146101bc575b5f80fd5b348015610118575f80fd5b506099546101399073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561016e575f80fd5b5061013961017d366004611908565b73ffffffffffffffffffffffffffffffffffffffff9081165f90815260fa60205260409020541690565b6101ba6101b536600461196f565b6103a5565b005b6101ba6101ca3660046119e5565b6103f0565b3480156101da575f80fd5b506101ba61042e565b3480156101ee575f80fd5b506097546101399073ffffffffffffffffffffffffffffffffffffffff1681565b6101ba61021d366004611a17565b610441565b34801561022d575f80fd5b5060655473ffffffffffffffffffffffffffffffffffffffff16610139565b6101ba61025a366004611aa9565b61088f565b34801561026a575f80fd5b50610139610279366004611908565b60fa6020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b3480156102ab575f80fd5b506101ba6102ba366004611aec565b6108a1565b3480156102ca575f80fd5b506101396102d9366004611908565b610aaf565b3480156102e9575f80fd5b5061030c6102f8366004611908565b60fb6020525f908152604090205460ff1681565b604051901515815260200161015a565b348015610327575f80fd5b506101ba610336366004611b34565b610b13565b348015610346575f80fd5b506101ba610355366004611908565b610c21565b348015610365575f80fd5b506098546101399073ffffffffffffffffffffffffffffffffffffffff1681565b348015610391575f80fd5b506101ba6103a0366004611b6f565b610cd8565b6103e886868686868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250889250610e18915050565b505050505050565b6104298333845f5b6040519080825280601f01601f191660200182016040528015610422576020820181803683370190505b5085610e18565b505050565b6104366112b1565b61043f5f611332565b565b60995473ffffffffffffffffffffffffffffffffffffffff163381146104c8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064015b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610511573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105359190611bc8565b60975473ffffffffffffffffffffffffffffffffffffffff9081169116146105b9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e74657270617274000000000000000060448201526064016104bf565b6105c16113a8565b3415610629576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6e6f6e7a65726f206d73672e76616c756500000000000000000000000000000060448201526064016104bf565b73ffffffffffffffffffffffffffffffffffffffff88166106a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f7420626520300000000000000060448201526064016104bf565b73ffffffffffffffffffffffffffffffffffffffff8088165f90815260fa6020526040902054898216911614610738576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6c3120746f6b656e206d69736d6174636800000000000000000000000000000060448201526064016104bf565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8681166004830152602482018690528816906340c10f19906044015f604051808303815f87803b1580156107a5575f80fd5b505af11580156107b7573d5f803e3d5ffd5b505050506107fa8584848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061141b92505050565b8573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167f165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34888888886040516108749493929190611be3565b60405180910390a461088560018055565b5050505050505050565b61089b8484845f6103f8565b50505050565b5f54610100900460ff16158080156108bf57505f54600160ff909116105b806108d85750303b1580156108d857505f5460ff166001145b610964576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104bf565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156109c0575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff8316610a3d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f7a65726f20726f7574657220616464726573730000000000000000000000000060448201526064016104bf565b610a488484846114cb565b801561089b575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f756e696d706c656d656e7465640000000000000000000000000000000000000060448201525f906064016104bf565b610b1b6112b1565b73ffffffffffffffffffffffffffffffffffffffff8216610b98576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f7420626520300000000000000060448201526064016104bf565b73ffffffffffffffffffffffffffffffffffffffff82165f81815260fb602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001685151590811790915591519182527fd8f6792507085b7664354b4599c60b3b600bd3f7e1a758f5e37134d4816b044a910160405180910390a25050565b610c296112b1565b73ffffffffffffffffffffffffffffffffffffffff8116610ccc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104bf565b610cd581611332565b50565b610ce06112b1565b73ffffffffffffffffffffffffffffffffffffffff8116610d5d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f7420626520300000000000000060448201526064016104bf565b73ffffffffffffffffffffffffffffffffffffffff8083165f81815260fa6020908152604080832080548787167fffffffffffffffffffffffff00000000000000000000000000000000000000008216811790925560fb90935281842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790559051919094169392849290917f2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d9190a4505050565b610e206113a8565b73ffffffffffffffffffffffffffffffffffffffff85165f90815260fb602052604090205460ff1615610eaf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f7769746864726177206c6f636b0000000000000000000000000000000000000060448201526064016104bf565b73ffffffffffffffffffffffffffffffffffffffff8086165f90815260fa60205260409020541680610f3d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f6e6f20636f72726573706f6e64696e67206c3120746f6b656e0000000000000060448201526064016104bf565b5f8411610fa6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7769746864726177207a65726f20616d6f756e7400000000000000000000000060448201526064016104bf565b609854339073ffffffffffffffffffffffffffffffffffffffff16819003610fe15783806020019051810190610fdc9190611c75565b945090505b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff828116600483015260248201879052881690639dc29fac906044015f604051808303815f87803b15801561104e575f80fd5b505af1158015611060573d5f803e3d5ffd5b505050505f82888389898960405160240161108096959493929190611d9a565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84bd13b00000000000000000000000000000000000000000000000000000000017905260995482517fecc7042800000000000000000000000000000000000000000000000000000000815292519394505f9373ffffffffffffffffffffffffffffffffffffffff9091169263ecc704289260048083019391928290030181865afa158015611164573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111889190611df4565b6099546097546040517fb2267a7b00000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff9182169263b2267a7b9234926111ed929116905f9088908c90600401611e0b565b5f604051808303818588803b158015611204575f80fd5b505af1158015611216573d5f803e3d5ffd5b50505050508273ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fa9967b740f3fc69dfbf4744b4b1c7dfdb0b1b63f1fa4cf573bcdcb9f3ac687c48b8b8b876040516112959493929190611e0b565b60405180910390a4505050506112aa60018055565b5050505050565b60655473ffffffffffffffffffffffffffffffffffffffff16331461043f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104bf565b6065805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b600260015403611414576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016104bf565b6002600155565b5f815111801561144157505f8273ffffffffffffffffffffffffffffffffffffffff163b115b156114c1576040517f444b281f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83169063444b281f90611498908490600401611e50565b5f604051808303815f87803b1580156114af575f80fd5b505af11580156103e8573d5f803e3d5ffd5b5050565b60018055565b73ffffffffffffffffffffffffffffffffffffffff8316611548576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f7a65726f20636f756e746572706172742061646472657373000000000000000060448201526064016104bf565b73ffffffffffffffffffffffffffffffffffffffff81166115c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f7a65726f206d657373656e67657220616464726573730000000000000000000060448201526064016104bf565b6115cd611676565b6115d5611714565b6097805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255609980548484169216919091179055821615610429576098805473ffffffffffffffffffffffffffffffffffffffff84167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b5f54610100900460ff1661170c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104bf565b61043f6117b2565b5f54610100900460ff166117aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104bf565b61043f611848565b5f54610100900460ff166114c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104bf565b5f54610100900460ff166118de576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104bf565b61043f33611332565b73ffffffffffffffffffffffffffffffffffffffff81168114610cd5575f80fd5b5f60208284031215611918575f80fd5b8135611923816118e7565b9392505050565b5f8083601f84011261193a575f80fd5b50813567ffffffffffffffff811115611951575f80fd5b602083019150836020828501011115611968575f80fd5b9250929050565b5f805f805f8060a08789031215611984575f80fd5b863561198f816118e7565b9550602087013561199f816118e7565b945060408701359350606087013567ffffffffffffffff8111156119c1575f80fd5b6119cd89828a0161192a565b979a9699509497949695608090950135949350505050565b5f805f606084860312156119f7575f80fd5b8335611a02816118e7565b95602085013595506040909401359392505050565b5f805f805f805f60c0888a031215611a2d575f80fd5b8735611a38816118e7565b96506020880135611a48816118e7565b95506040880135611a58816118e7565b94506060880135611a68816118e7565b93506080880135925060a088013567ffffffffffffffff811115611a8a575f80fd5b611a968a828b0161192a565b989b979a50959850939692959293505050565b5f805f8060808587031215611abc575f80fd5b8435611ac7816118e7565b93506020850135611ad7816118e7565b93969395505050506040820135916060013590565b5f805f60608486031215611afe575f80fd5b8335611b09816118e7565b92506020840135611b19816118e7565b91506040840135611b29816118e7565b809150509250925092565b5f8060408385031215611b45575f80fd5b8235611b50816118e7565b915060208301358015158114611b64575f80fd5b809150509250929050565b5f8060408385031215611b80575f80fd5b8235611b8b816118e7565b91506020830135611b64816118e7565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f60208284031215611bd8575f80fd5b8151611923816118e7565b73ffffffffffffffffffffffffffffffffffffffff8516815283602082015260606040820152816060820152818360808301375f818301608090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01601019392505050565b5f5b83811015611c6d578181015183820152602001611c55565b50505f910152565b5f8060408385031215611c86575f80fd5b8251611c91816118e7565b602084015190925067ffffffffffffffff80821115611cae575f80fd5b818501915085601f830112611cc1575f80fd5b815181811115611cd357611cd3611b9b565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611d1957611d19611b9b565b81604052828152886020848701011115611d31575f80fd5b611d42836020830160208801611c53565b80955050505050509250929050565b5f8151808452611d68816020860160208601611c53565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b5f73ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525083608083015260c060a0830152611de860c0830184611d51565b98975050505050505050565b5f60208284031215611e04575f80fd5b5051919050565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152608060408201525f611e3f6080830185611d51565b905082606083015295945050505050565b602081525f6119236020830184611d5156fea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(L2WithdrawLockERC20GatewayStorageLayoutJSON), L2WithdrawLockERC20GatewayStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2WithdrawLockERC20Gateway"] = L2WithdrawLockERC20GatewayStorageLayout
	deployedBytecodes["L2WithdrawLockERC20Gateway"] = L2WithdrawLockERC20GatewayDeployedBin
}