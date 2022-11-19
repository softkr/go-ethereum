// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package auction

import (
	"math/big"
	"strings"

	"gitee.com/aqchain/go-ethereum"
	"gitee.com/aqchain/go-ethereum/accounts/abi"
	"gitee.com/aqchain/go-ethereum/accounts/abi/bind"
	"gitee.com/aqchain/go-ethereum/common"
	"gitee.com/aqchain/go-ethereum/core/types"
	"gitee.com/aqchain/go-ethereum/event"
)

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
const AddressBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820fb3395e9478145a576a5cd6b4d5ad35f3ddad544c862098809480741bb05543f0029`

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// AuctionContractABI is the input ABI used to generate the binding from.
const AuctionContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getDeadline\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getHighestBid\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getHighestBidder\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"reservePrice\",\"type\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"finish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getReservePrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getSeller\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"itemContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"dci\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"reservePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"Create\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"dci\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"dci\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"highestBid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"highestBidder\",\"type\":\"address\"}],\"name\":\"Finish\",\"type\":\"event\"}]"

// AuctionContractBin is the compiled bytecode used for deploying new contracts.
const AuctionContractBin = `0x608060405234801561001057600080fd5b5060405160208061114d8339810180604052602081101561003057600080fd5b505160008054600160a060020a039092166101000261010060a860020a031960ff19909316600117929092169190911790556110dc806100716000396000f3fe6080604052600436106100a3576000357c010000000000000000000000000000000000000000000000000000000090048063c750cb7911610076578063c750cb791461016b578063c97fe94b146101b1578063d353a1cb146101e7578063d45c35ff14610211578063d6a9de511461023b576100a3565b8063454a2ab3146100a85780634f558e79146100c757806382862275146101055780638f28864414610141575b600080fd5b6100c5600480360360208110156100be57600080fd5b5035610265565b005b3480156100d357600080fd5b506100f1600480360360208110156100ea57600080fd5b5035610632565b604080519115158252519081900360200190f35b34801561011157600080fd5b5061012f6004803603602081101561012857600080fd5b503561064f565b60408051918252519081900360200190f35b34801561014d57600080fd5b5061012f6004803603602081101561016457600080fd5b5035610664565b34801561017757600080fd5b506101956004803603602081101561018e57600080fd5b5035610679565b60408051600160a060020a039092168252519081900360200190f35b3480156101bd57600080fd5b506100c5600480360360608110156101d457600080fd5b5080359060208101359060400135610697565b3480156101f357600080fd5b506100c56004803603602081101561020a57600080fd5b5035610a0e565b34801561021d57600080fd5b5061012f6004803603602081101561023457600080fd5b5035610e8e565b34801561024757600080fd5b506101956004803603602081101561025e57600080fd5b5035610ea4565b60005460ff1615156102c1576040805160e560020a62461bcd02815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6000805460ff191681558181526001602052604090208054600160a060020a03161515610338576040805160e560020a62461bcd02815260206004820181905260248201527f41756374696f6e3a20746f6b656e206973206e6f74206f6e2061756374696f6e604482015290519081900360640190fd5b600181015434101561037e5760405160e560020a62461bcd0281526004018080602001828103825260308152602001806110816030913960400191505060405180910390fd5b60008281526001602052604090206002015442106103d05760405160e560020a62461bcd0281526004018080602001828103825260218152602001806110316021913960400191505060405180910390fd5b6003810154600160a060020a03161561044d57600481015434116104285760405160e560020a62461bcd02815260040180806020018281038252602f815260200180611052602f913960400191505060405180910390fd5b6004810154600382015461044d91600160a060020a039091169063ffffffff610ebf16565b3460048083019190915560038201805473ffffffffffffffffffffffffffffffffffffffff19163317905560008054604080517fdb1573390000000000000000000000000000000000000000000000000000000081529384018690525185937fb773ce5048c68e3c2b907d5b3fbf2875976f4c973ae9494721266036d3d9af8b93610100909304600160a060020a03169263db15733992602480840193829003018186803b1580156104fe57600080fd5b505afa158015610512573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561053b57600080fd5b81019080805164010000000081111561055357600080fd5b8201602081018481111561056657600080fd5b815164010000000081118282018710171561058057600080fd5b50509291905050503334604051808060200184600160a060020a0316600160a060020a03168152602001838152602001828103825285818151815260200191508051906020019080838360005b838110156105e55781810151838201526020016105cd565b50505050905090810190601f1680156106125780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a250506000805460ff19166001179055565b600090815260016020526040902054600160a060020a0316151590565b60009081526001602052604090206002015490565b60009081526001602052604090206004015490565b600090815260016020526040902060030154600160a060020a031690565b60005460ff1615156106f3576040805160e560020a62461bcd02815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6000805460ff1916905561070683610632565b156107455760405160e560020a62461bcd02815260040180806020018281038252602381526020018061100e6023913960400191505060405180910390fd5b6040805160a0810182523381526020808201858152828401858152600060608501818152608086018281528a83526001958690528783209651875473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a0392831617895595519688019690965592516002870155516003860180549094169085161790925551600493840155805484517fb7760c8f00000000000000000000000000000000000000000000000000000000815293840188905230602485015293516101009094049091169263b7760c8f92604480820193929182900301818387803b15801561083157600080fd5b505af1158015610845573d6000803e3d6000fd5b505060008054604080517fdb1573390000000000000000000000000000000000000000000000000000000081526004810189905290518895507f843836e9bbe62d07b095e8e28eac630dac44c1b739bf84765a826b25dd88e33e9450610100909204600160a060020a03169263db15733992602480840193829003018186803b1580156108d157600080fd5b505afa1580156108e5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561090e57600080fd5b81019080805164010000000081111561092657600080fd5b8201602081018481111561093957600080fd5b815164010000000081118282018710171561095357600080fd5b5050929190505050338585604051808060200185600160a060020a0316600160a060020a03168152602001848152602001838152602001828103825286818151815260200191508051906020019080838360005b838110156109bf5781810151838201526020016109a7565b50505050905090810190601f1680156109ec5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a250506000805460ff1916600117905550565b60005460ff161515610a6a576040805160e560020a62461bcd02815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6000805460ff191681558181526001602052604090208054600160a060020a03161515610ae1576040805160e560020a62461bcd02815260206004820181905260248201527f41756374696f6e3a20746f6b656e206973206e6f74206f6e2061756374696f6e604482015290519081900360640190fd5b6002810154421015610b275760405160e560020a62461bcd028152600401808060200182810382526021815260200180610fed6021913960400191505060405180910390fd5b6003810154600160a060020a031615610beb57600080546003830154604080517fb7760c8f00000000000000000000000000000000000000000000000000000000815260048101879052600160a060020a03928316602482015290516101009093049091169263b7760c8f9260448084019382900301818387803b158015610bae57600080fd5b505af1158015610bc2573d6000803e3d6000fd5b50505060048201548254610be69250600160a060020a03169063ffffffff610ebf16565b610c75565b600080548254604080517fb7760c8f00000000000000000000000000000000000000000000000000000000815260048101879052600160a060020a03928316602482015290516101009093049091169263b7760c8f9260448084019382900301818387803b158015610c5c57600080fd5b505af1158015610c70573d6000803e3d6000fd5b505050505b6000828152600160208190526040808320805473ffffffffffffffffffffffffffffffffffffffff199081168255928101849055600281018490556003810180549093169092556004918201839055825481517fdb157339000000000000000000000000000000000000000000000000000000008152928301869052905185937f24d7d71f521fe04fd826928ee3b3acad562942a8348377a8f1edd96b36a3e3e093610100909304600160a060020a03169263db1573399260248083019392829003018186803b158015610d4857600080fd5b505afa158015610d5c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610d8557600080fd5b810190808051640100000000811115610d9d57600080fd5b82016020810184811115610db057600080fd5b8151640100000000811182820187101715610dca57600080fd5b505086546004880154600389015460408051600160a060020a03948516602080830182905292820185905292909416606085018190526080808652865190860152855195985091965091945092829160a08301919088019080838360005b83811015610e40578181015183820152602001610e28565b50505050905090810190601f168015610e6d5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a250506000805460ff19166001179055565b6000908152600160208190526040909120015490565b600090815260016020526040902054600160a060020a031690565b3031811115610f18576040805160e560020a62461bcd02815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e6365000000604482015290519081900360640190fd5b604051600090600160a060020a0384169083908381818185875af1925050503d8060008114610f63576040519150601f19603f3d011682016040523d82523d6000602084013e610f68565b606091505b50509050801515610fad5760405160e560020a62461bcd02815260040180806020018281038252603a815260200180610fb3603a913960400191505060405180910390fd5b50505056fe416464726573733a20756e61626c6520746f2073656e642076616c75652c20726563697069656e74206d6179206861766520726576657274656441756374696f6e3a2061756374696f6e206973206e6f7420656e6465642079657441756374696f6e436f6e74726163743a206974656d206973206f6e2061756374696f6e41756374696f6e436f6e74726163743a2061756374696f6e20697320656e64656441756374696f6e3a206269642063616e206e6f74206c6f776572207468616e2074686520686967686573742062696441756374696f6e3a206269642063616e206e6f74206c6f776572207468616e20746865206c6f77657374207072696365a165627a7a723058209be735b11d92c480ccf048fcb9c3a921fb5df702765ec67c82d8e8022368953a0029`

// DeployAuctionContract deploys a new Ethereum contract, binding an instance of AuctionContract to it.
func DeployAuctionContract(auth *bind.TransactOpts, backend bind.ContractBackend, itemContract common.Address) (common.Address, *types.Transaction, *AuctionContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuctionContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AuctionContractBin), backend, itemContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AuctionContract{AuctionContractCaller: AuctionContractCaller{contract: contract}, AuctionContractTransactor: AuctionContractTransactor{contract: contract}, AuctionContractFilterer: AuctionContractFilterer{contract: contract}}, nil
}

// AuctionContract is an auto generated Go binding around an Ethereum contract.
type AuctionContract struct {
	AuctionContractCaller     // Read-only binding to the contract
	AuctionContractTransactor // Write-only binding to the contract
	AuctionContractFilterer   // Log filterer for contract events
}

// AuctionContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionContractSession struct {
	Contract     *AuctionContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionContractCallerSession struct {
	Contract *AuctionContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AuctionContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionContractTransactorSession struct {
	Contract     *AuctionContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AuctionContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionContractRaw struct {
	Contract *AuctionContract // Generic contract binding to access the raw methods on
}

// AuctionContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionContractCallerRaw struct {
	Contract *AuctionContractCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionContractTransactorRaw struct {
	Contract *AuctionContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuctionContract creates a new instance of AuctionContract, bound to a specific deployed contract.
func NewAuctionContract(address common.Address, backend bind.ContractBackend) (*AuctionContract, error) {
	contract, err := bindAuctionContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuctionContract{AuctionContractCaller: AuctionContractCaller{contract: contract}, AuctionContractTransactor: AuctionContractTransactor{contract: contract}, AuctionContractFilterer: AuctionContractFilterer{contract: contract}}, nil
}

// NewAuctionContractCaller creates a new read-only instance of AuctionContract, bound to a specific deployed contract.
func NewAuctionContractCaller(address common.Address, caller bind.ContractCaller) (*AuctionContractCaller, error) {
	contract, err := bindAuctionContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionContractCaller{contract: contract}, nil
}

// NewAuctionContractTransactor creates a new write-only instance of AuctionContract, bound to a specific deployed contract.
func NewAuctionContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionContractTransactor, error) {
	contract, err := bindAuctionContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionContractTransactor{contract: contract}, nil
}

// NewAuctionContractFilterer creates a new log filterer instance of AuctionContract, bound to a specific deployed contract.
func NewAuctionContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionContractFilterer, error) {
	contract, err := bindAuctionContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionContractFilterer{contract: contract}, nil
}

// bindAuctionContract binds a generic wrapper to an already deployed contract.
func bindAuctionContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuctionContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionContract *AuctionContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AuctionContract.Contract.AuctionContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionContract *AuctionContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionContract.Contract.AuctionContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionContract *AuctionContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionContract.Contract.AuctionContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionContract *AuctionContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AuctionContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionContract *AuctionContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionContract *AuctionContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionContract.Contract.contract.Transact(opts, method, params...)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(tokenId uint256) constant returns(bool)
func (_AuctionContract *AuctionContractCaller) Exists(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AuctionContract.contract.Call(opts, out, "exists", tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(tokenId uint256) constant returns(bool)
func (_AuctionContract *AuctionContractSession) Exists(tokenId *big.Int) (bool, error) {
	return _AuctionContract.Contract.Exists(&_AuctionContract.CallOpts, tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(tokenId uint256) constant returns(bool)
func (_AuctionContract *AuctionContractCallerSession) Exists(tokenId *big.Int) (bool, error) {
	return _AuctionContract.Contract.Exists(&_AuctionContract.CallOpts, tokenId)
}

// GetDeadline is a free data retrieval call binding the contract method 0x82862275.
//
// Solidity: function getDeadline(tokenId uint256) constant returns(uint256)
func (_AuctionContract *AuctionContractCaller) GetDeadline(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AuctionContract.contract.Call(opts, out, "getDeadline", tokenId)
	return *ret0, err
}

// GetDeadline is a free data retrieval call binding the contract method 0x82862275.
//
// Solidity: function getDeadline(tokenId uint256) constant returns(uint256)
func (_AuctionContract *AuctionContractSession) GetDeadline(tokenId *big.Int) (*big.Int, error) {
	return _AuctionContract.Contract.GetDeadline(&_AuctionContract.CallOpts, tokenId)
}

// GetDeadline is a free data retrieval call binding the contract method 0x82862275.
//
// Solidity: function getDeadline(tokenId uint256) constant returns(uint256)
func (_AuctionContract *AuctionContractCallerSession) GetDeadline(tokenId *big.Int) (*big.Int, error) {
	return _AuctionContract.Contract.GetDeadline(&_AuctionContract.CallOpts, tokenId)
}

// GetHighestBid is a free data retrieval call binding the contract method 0x8f288644.
//
// Solidity: function getHighestBid(tokenId uint256) constant returns(uint256)
func (_AuctionContract *AuctionContractCaller) GetHighestBid(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AuctionContract.contract.Call(opts, out, "getHighestBid", tokenId)
	return *ret0, err
}

// GetHighestBid is a free data retrieval call binding the contract method 0x8f288644.
//
// Solidity: function getHighestBid(tokenId uint256) constant returns(uint256)
func (_AuctionContract *AuctionContractSession) GetHighestBid(tokenId *big.Int) (*big.Int, error) {
	return _AuctionContract.Contract.GetHighestBid(&_AuctionContract.CallOpts, tokenId)
}

// GetHighestBid is a free data retrieval call binding the contract method 0x8f288644.
//
// Solidity: function getHighestBid(tokenId uint256) constant returns(uint256)
func (_AuctionContract *AuctionContractCallerSession) GetHighestBid(tokenId *big.Int) (*big.Int, error) {
	return _AuctionContract.Contract.GetHighestBid(&_AuctionContract.CallOpts, tokenId)
}

// GetHighestBidder is a free data retrieval call binding the contract method 0xc750cb79.
//
// Solidity: function getHighestBidder(tokenId uint256) constant returns(address)
func (_AuctionContract *AuctionContractCaller) GetHighestBidder(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AuctionContract.contract.Call(opts, out, "getHighestBidder", tokenId)
	return *ret0, err
}

// GetHighestBidder is a free data retrieval call binding the contract method 0xc750cb79.
//
// Solidity: function getHighestBidder(tokenId uint256) constant returns(address)
func (_AuctionContract *AuctionContractSession) GetHighestBidder(tokenId *big.Int) (common.Address, error) {
	return _AuctionContract.Contract.GetHighestBidder(&_AuctionContract.CallOpts, tokenId)
}

// GetHighestBidder is a free data retrieval call binding the contract method 0xc750cb79.
//
// Solidity: function getHighestBidder(tokenId uint256) constant returns(address)
func (_AuctionContract *AuctionContractCallerSession) GetHighestBidder(tokenId *big.Int) (common.Address, error) {
	return _AuctionContract.Contract.GetHighestBidder(&_AuctionContract.CallOpts, tokenId)
}

// GetReservePrice is a free data retrieval call binding the contract method 0xd45c35ff.
//
// Solidity: function getReservePrice(tokenId uint256) constant returns(uint256)
func (_AuctionContract *AuctionContractCaller) GetReservePrice(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AuctionContract.contract.Call(opts, out, "getReservePrice", tokenId)
	return *ret0, err
}

// GetReservePrice is a free data retrieval call binding the contract method 0xd45c35ff.
//
// Solidity: function getReservePrice(tokenId uint256) constant returns(uint256)
func (_AuctionContract *AuctionContractSession) GetReservePrice(tokenId *big.Int) (*big.Int, error) {
	return _AuctionContract.Contract.GetReservePrice(&_AuctionContract.CallOpts, tokenId)
}

// GetReservePrice is a free data retrieval call binding the contract method 0xd45c35ff.
//
// Solidity: function getReservePrice(tokenId uint256) constant returns(uint256)
func (_AuctionContract *AuctionContractCallerSession) GetReservePrice(tokenId *big.Int) (*big.Int, error) {
	return _AuctionContract.Contract.GetReservePrice(&_AuctionContract.CallOpts, tokenId)
}

// GetSeller is a free data retrieval call binding the contract method 0xd6a9de51.
//
// Solidity: function getSeller(tokenId uint256) constant returns(address)
func (_AuctionContract *AuctionContractCaller) GetSeller(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AuctionContract.contract.Call(opts, out, "getSeller", tokenId)
	return *ret0, err
}

// GetSeller is a free data retrieval call binding the contract method 0xd6a9de51.
//
// Solidity: function getSeller(tokenId uint256) constant returns(address)
func (_AuctionContract *AuctionContractSession) GetSeller(tokenId *big.Int) (common.Address, error) {
	return _AuctionContract.Contract.GetSeller(&_AuctionContract.CallOpts, tokenId)
}

// GetSeller is a free data retrieval call binding the contract method 0xd6a9de51.
//
// Solidity: function getSeller(tokenId uint256) constant returns(address)
func (_AuctionContract *AuctionContractCallerSession) GetSeller(tokenId *big.Int) (common.Address, error) {
	return _AuctionContract.Contract.GetSeller(&_AuctionContract.CallOpts, tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(tokenId uint256) returns()
func (_AuctionContract *AuctionContractTransactor) Bid(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _AuctionContract.contract.Transact(opts, "bid", tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(tokenId uint256) returns()
func (_AuctionContract *AuctionContractSession) Bid(tokenId *big.Int) (*types.Transaction, error) {
	return _AuctionContract.Contract.Bid(&_AuctionContract.TransactOpts, tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(tokenId uint256) returns()
func (_AuctionContract *AuctionContractTransactorSession) Bid(tokenId *big.Int) (*types.Transaction, error) {
	return _AuctionContract.Contract.Bid(&_AuctionContract.TransactOpts, tokenId)
}

// Create is a paid mutator transaction binding the contract method 0xc97fe94b.
//
// Solidity: function create(tokenId uint256, reservePrice uint256, deadline uint256) returns()
func (_AuctionContract *AuctionContractTransactor) Create(opts *bind.TransactOpts, tokenId *big.Int, reservePrice *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AuctionContract.contract.Transact(opts, "create", tokenId, reservePrice, deadline)
}

// Create is a paid mutator transaction binding the contract method 0xc97fe94b.
//
// Solidity: function create(tokenId uint256, reservePrice uint256, deadline uint256) returns()
func (_AuctionContract *AuctionContractSession) Create(tokenId *big.Int, reservePrice *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AuctionContract.Contract.Create(&_AuctionContract.TransactOpts, tokenId, reservePrice, deadline)
}

// Create is a paid mutator transaction binding the contract method 0xc97fe94b.
//
// Solidity: function create(tokenId uint256, reservePrice uint256, deadline uint256) returns()
func (_AuctionContract *AuctionContractTransactorSession) Create(tokenId *big.Int, reservePrice *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _AuctionContract.Contract.Create(&_AuctionContract.TransactOpts, tokenId, reservePrice, deadline)
}

// Finish is a paid mutator transaction binding the contract method 0xd353a1cb.
//
// Solidity: function finish(tokenId uint256) returns()
func (_AuctionContract *AuctionContractTransactor) Finish(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _AuctionContract.contract.Transact(opts, "finish", tokenId)
}

// Finish is a paid mutator transaction binding the contract method 0xd353a1cb.
//
// Solidity: function finish(tokenId uint256) returns()
func (_AuctionContract *AuctionContractSession) Finish(tokenId *big.Int) (*types.Transaction, error) {
	return _AuctionContract.Contract.Finish(&_AuctionContract.TransactOpts, tokenId)
}

// Finish is a paid mutator transaction binding the contract method 0xd353a1cb.
//
// Solidity: function finish(tokenId uint256) returns()
func (_AuctionContract *AuctionContractTransactorSession) Finish(tokenId *big.Int) (*types.Transaction, error) {
	return _AuctionContract.Contract.Finish(&_AuctionContract.TransactOpts, tokenId)
}

// AuctionContractBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the AuctionContract contract.
type AuctionContractBidIterator struct {
	Event *AuctionContractBid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AuctionContractBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionContractBid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AuctionContractBid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AuctionContractBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionContractBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionContractBid represents a Bid event raised by the AuctionContract contract.
type AuctionContractBid struct {
	TokenId *big.Int
	Dci     string
	Bidder  common.Address
	Bid     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0xb773ce5048c68e3c2b907d5b3fbf2875976f4c973ae9494721266036d3d9af8b.
//
// Solidity: e Bid(tokenId indexed uint256, dci string, bidder address, bid uint256)
func (_AuctionContract *AuctionContractFilterer) FilterBid(opts *bind.FilterOpts, tokenId []*big.Int) (*AuctionContractBidIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AuctionContract.contract.FilterLogs(opts, "Bid", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AuctionContractBidIterator{contract: _AuctionContract.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0xb773ce5048c68e3c2b907d5b3fbf2875976f4c973ae9494721266036d3d9af8b.
//
// Solidity: e Bid(tokenId indexed uint256, dci string, bidder address, bid uint256)
func (_AuctionContract *AuctionContractFilterer) WatchBid(opts *bind.WatchOpts, sink chan<- *AuctionContractBid, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AuctionContract.contract.WatchLogs(opts, "Bid", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionContractBid)
				if err := _AuctionContract.contract.UnpackLog(event, "Bid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AuctionContractCreateIterator is returned from FilterCreate and is used to iterate over the raw logs and unpacked data for Create events raised by the AuctionContract contract.
type AuctionContractCreateIterator struct {
	Event *AuctionContractCreate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AuctionContractCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionContractCreate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AuctionContractCreate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AuctionContractCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionContractCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionContractCreate represents a Create event raised by the AuctionContract contract.
type AuctionContractCreate struct {
	TokenId      *big.Int
	Dci          string
	Seller       common.Address
	ReservePrice *big.Int
	Deadline     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreate is a free log retrieval operation binding the contract event 0x843836e9bbe62d07b095e8e28eac630dac44c1b739bf84765a826b25dd88e33e.
//
// Solidity: e Create(tokenId indexed uint256, dci string, seller address, reservePrice uint256, deadline uint256)
func (_AuctionContract *AuctionContractFilterer) FilterCreate(opts *bind.FilterOpts, tokenId []*big.Int) (*AuctionContractCreateIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AuctionContract.contract.FilterLogs(opts, "Create", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AuctionContractCreateIterator{contract: _AuctionContract.contract, event: "Create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0x843836e9bbe62d07b095e8e28eac630dac44c1b739bf84765a826b25dd88e33e.
//
// Solidity: e Create(tokenId indexed uint256, dci string, seller address, reservePrice uint256, deadline uint256)
func (_AuctionContract *AuctionContractFilterer) WatchCreate(opts *bind.WatchOpts, sink chan<- *AuctionContractCreate, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AuctionContract.contract.WatchLogs(opts, "Create", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionContractCreate)
				if err := _AuctionContract.contract.UnpackLog(event, "Create", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AuctionContractFinishIterator is returned from FilterFinish and is used to iterate over the raw logs and unpacked data for Finish events raised by the AuctionContract contract.
type AuctionContractFinishIterator struct {
	Event *AuctionContractFinish // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AuctionContractFinishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionContractFinish)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AuctionContractFinish)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AuctionContractFinishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionContractFinishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionContractFinish represents a Finish event raised by the AuctionContract contract.
type AuctionContractFinish struct {
	TokenId       *big.Int
	Dci           string
	Seller        common.Address
	HighestBid    *big.Int
	HighestBidder common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFinish is a free log retrieval operation binding the contract event 0x24d7d71f521fe04fd826928ee3b3acad562942a8348377a8f1edd96b36a3e3e0.
//
// Solidity: e Finish(tokenId indexed uint256, dci string, seller address, highestBid uint256, highestBidder address)
func (_AuctionContract *AuctionContractFilterer) FilterFinish(opts *bind.FilterOpts, tokenId []*big.Int) (*AuctionContractFinishIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AuctionContract.contract.FilterLogs(opts, "Finish", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AuctionContractFinishIterator{contract: _AuctionContract.contract, event: "Finish", logs: logs, sub: sub}, nil
}

// WatchFinish is a free log subscription operation binding the contract event 0x24d7d71f521fe04fd826928ee3b3acad562942a8348377a8f1edd96b36a3e3e0.
//
// Solidity: e Finish(tokenId indexed uint256, dci string, seller address, highestBid uint256, highestBidder address)
func (_AuctionContract *AuctionContractFilterer) WatchFinish(opts *bind.WatchOpts, sink chan<- *AuctionContractFinish, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AuctionContract.contract.WatchLogs(opts, "Finish", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionContractFinish)
				if err := _AuctionContract.contract.UnpackLog(event, "Finish", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// CountersABI is the input ABI used to generate the binding from.
const CountersABI = "[]"

// CountersBin is the compiled bytecode used for deploying new contracts.
const CountersBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a7230582048c828fece9be6dc90fc86a9b6055bfc3e4aeb185930b7d3e3331c6799011bd40029`

// DeployCounters deploys a new Ethereum contract, binding an instance of Counters to it.
func DeployCounters(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counters, error) {
	parsed, err := abi.JSON(strings.NewReader(CountersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CountersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counters{CountersCaller: CountersCaller{contract: contract}, CountersTransactor: CountersTransactor{contract: contract}, CountersFilterer: CountersFilterer{contract: contract}}, nil
}

// Counters is an auto generated Go binding around an Ethereum contract.
type Counters struct {
	CountersCaller     // Read-only binding to the contract
	CountersTransactor // Write-only binding to the contract
	CountersFilterer   // Log filterer for contract events
}

// CountersCaller is an auto generated read-only Go binding around an Ethereum contract.
type CountersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CountersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CountersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CountersSession struct {
	Contract     *Counters         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CountersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CountersCallerSession struct {
	Contract *CountersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CountersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CountersTransactorSession struct {
	Contract     *CountersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CountersRaw is an auto generated low-level Go binding around an Ethereum contract.
type CountersRaw struct {
	Contract *Counters // Generic contract binding to access the raw methods on
}

// CountersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CountersCallerRaw struct {
	Contract *CountersCaller // Generic read-only contract binding to access the raw methods on
}

// CountersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CountersTransactorRaw struct {
	Contract *CountersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounters creates a new instance of Counters, bound to a specific deployed contract.
func NewCounters(address common.Address, backend bind.ContractBackend) (*Counters, error) {
	contract, err := bindCounters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counters{CountersCaller: CountersCaller{contract: contract}, CountersTransactor: CountersTransactor{contract: contract}, CountersFilterer: CountersFilterer{contract: contract}}, nil
}

// NewCountersCaller creates a new read-only instance of Counters, bound to a specific deployed contract.
func NewCountersCaller(address common.Address, caller bind.ContractCaller) (*CountersCaller, error) {
	contract, err := bindCounters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CountersCaller{contract: contract}, nil
}

// NewCountersTransactor creates a new write-only instance of Counters, bound to a specific deployed contract.
func NewCountersTransactor(address common.Address, transactor bind.ContractTransactor) (*CountersTransactor, error) {
	contract, err := bindCounters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CountersTransactor{contract: contract}, nil
}

// NewCountersFilterer creates a new log filterer instance of Counters, bound to a specific deployed contract.
func NewCountersFilterer(address common.Address, filterer bind.ContractFilterer) (*CountersFilterer, error) {
	contract, err := bindCounters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CountersFilterer{contract: contract}, nil
}

// bindCounters binds a generic wrapper to an already deployed contract.
func bindCounters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CountersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counters *CountersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Counters.Contract.CountersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counters *CountersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counters.Contract.CountersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counters *CountersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counters.Contract.CountersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counters *CountersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Counters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counters *CountersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counters *CountersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counters.Contract.contract.Transact(opts, method, params...)
}

// EnumerableMapABI is the input ABI used to generate the binding from.
const EnumerableMapABI = "[]"

// EnumerableMapBin is the compiled bytecode used for deploying new contracts.
const EnumerableMapBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820ebf14ba25eb19e1870e24dbb6622a7fb943bc76099e8ec4829d671b257a9675e0029`

// DeployEnumerableMap deploys a new Ethereum contract, binding an instance of EnumerableMap to it.
func DeployEnumerableMap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableMap, error) {
	parsed, err := abi.JSON(strings.NewReader(EnumerableMapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EnumerableMapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableMap{EnumerableMapCaller: EnumerableMapCaller{contract: contract}, EnumerableMapTransactor: EnumerableMapTransactor{contract: contract}, EnumerableMapFilterer: EnumerableMapFilterer{contract: contract}}, nil
}

// EnumerableMap is an auto generated Go binding around an Ethereum contract.
type EnumerableMap struct {
	EnumerableMapCaller     // Read-only binding to the contract
	EnumerableMapTransactor // Write-only binding to the contract
	EnumerableMapFilterer   // Log filterer for contract events
}

// EnumerableMapCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnumerableMapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableMapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnumerableMapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableMapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnumerableMapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableMapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnumerableMapSession struct {
	Contract     *EnumerableMap    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumerableMapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnumerableMapCallerSession struct {
	Contract *EnumerableMapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EnumerableMapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnumerableMapTransactorSession struct {
	Contract     *EnumerableMapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnumerableMapRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnumerableMapRaw struct {
	Contract *EnumerableMap // Generic contract binding to access the raw methods on
}

// EnumerableMapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnumerableMapCallerRaw struct {
	Contract *EnumerableMapCaller // Generic read-only contract binding to access the raw methods on
}

// EnumerableMapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnumerableMapTransactorRaw struct {
	Contract *EnumerableMapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnumerableMap creates a new instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMap(address common.Address, backend bind.ContractBackend) (*EnumerableMap, error) {
	contract, err := bindEnumerableMap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableMap{EnumerableMapCaller: EnumerableMapCaller{contract: contract}, EnumerableMapTransactor: EnumerableMapTransactor{contract: contract}, EnumerableMapFilterer: EnumerableMapFilterer{contract: contract}}, nil
}

// NewEnumerableMapCaller creates a new read-only instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMapCaller(address common.Address, caller bind.ContractCaller) (*EnumerableMapCaller, error) {
	contract, err := bindEnumerableMap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableMapCaller{contract: contract}, nil
}

// NewEnumerableMapTransactor creates a new write-only instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMapTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableMapTransactor, error) {
	contract, err := bindEnumerableMap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableMapTransactor{contract: contract}, nil
}

// NewEnumerableMapFilterer creates a new log filterer instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMapFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableMapFilterer, error) {
	contract, err := bindEnumerableMap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableMapFilterer{contract: contract}, nil
}

// bindEnumerableMap binds a generic wrapper to an already deployed contract.
func bindEnumerableMap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnumerableMapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableMap *EnumerableMapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EnumerableMap.Contract.EnumerableMapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableMap *EnumerableMapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableMap.Contract.EnumerableMapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableMap *EnumerableMapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableMap.Contract.EnumerableMapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableMap *EnumerableMapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EnumerableMap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableMap *EnumerableMapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableMap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableMap *EnumerableMapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableMap.Contract.contract.Transact(opts, method, params...)
}

// EnumerableSetABI is the input ABI used to generate the binding from.
const EnumerableSetABI = "[]"

// EnumerableSetBin is the compiled bytecode used for deploying new contracts.
const EnumerableSetBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a723058207f8aad7f8b6e085399c78c41dc0ebbae4c06295012edf29867adb3552e50f67a0029`

// DeployEnumerableSet deploys a new Ethereum contract, binding an instance of EnumerableSet to it.
func DeployEnumerableSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableSet, error) {
	parsed, err := abi.JSON(strings.NewReader(EnumerableSetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EnumerableSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// EnumerableSet is an auto generated Go binding around an Ethereum contract.
type EnumerableSet struct {
	EnumerableSetCaller     // Read-only binding to the contract
	EnumerableSetTransactor // Write-only binding to the contract
	EnumerableSetFilterer   // Log filterer for contract events
}

// EnumerableSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnumerableSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnumerableSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnumerableSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnumerableSetSession struct {
	Contract     *EnumerableSet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumerableSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnumerableSetCallerSession struct {
	Contract *EnumerableSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EnumerableSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnumerableSetTransactorSession struct {
	Contract     *EnumerableSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnumerableSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnumerableSetRaw struct {
	Contract *EnumerableSet // Generic contract binding to access the raw methods on
}

// EnumerableSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnumerableSetCallerRaw struct {
	Contract *EnumerableSetCaller // Generic read-only contract binding to access the raw methods on
}

// EnumerableSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnumerableSetTransactorRaw struct {
	Contract *EnumerableSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnumerableSet creates a new instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSet(address common.Address, backend bind.ContractBackend) (*EnumerableSet, error) {
	contract, err := bindEnumerableSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// NewEnumerableSetCaller creates a new read-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetCaller(address common.Address, caller bind.ContractCaller) (*EnumerableSetCaller, error) {
	contract, err := bindEnumerableSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetCaller{contract: contract}, nil
}

// NewEnumerableSetTransactor creates a new write-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableSetTransactor, error) {
	contract, err := bindEnumerableSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetTransactor{contract: contract}, nil
}

// NewEnumerableSetFilterer creates a new log filterer instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableSetFilterer, error) {
	contract, err := bindEnumerableSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetFilterer{contract: contract}, nil
}

// bindEnumerableSet binds a generic wrapper to an already deployed contract.
func bindEnumerableSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnumerableSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.EnumerableSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transact(opts, method, params...)
}

// ItemContractABI is the input ABI used to generate the binding from.
const ItemContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"dci\",\"type\":\"string\"}],\"name\":\"getTokenId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dci\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"create\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getCreator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getDCI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"dci\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"Create\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"dci\",\"type\":\"string\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"dci\",\"type\":\"string\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ItemContractBin is the compiled bytecode used for deploying new contracts.
const ItemContractBin = `0x608060405234801561001057600080fd5b506114f1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100c6576000357c0100000000000000000000000000000000000000000000000000000000900480636b8ff5741161008e5780636b8ff5741461042a578063b7760c8f14610447578063c41a360a14610473578063d48e638a14610490578063db157339146104ad578063febe4909146104ca576100c6565b8063081812fc146100cb5780631e7663bc146101045780633bb3a24d146101bc5780634f0cd27b1461024e5780635d28560a14610274575b600080fd5b6100e8600480360360208110156100e157600080fd5b50356104f6565b60408051600160a060020a039092168252519081900360200190f35b6101aa6004803603602081101561011a57600080fd5b81019060208101813564010000000081111561013557600080fd5b82018360208201111561014757600080fd5b8035906020019184600183028401116401000000008311171561016957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061055d945050505050565b60408051918252519081900360200190f35b6101d9600480360360208110156101d257600080fd5b50356105c5565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102135781810151838201526020016101fb565b50505050905090810190601f1680156102405780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101aa6004803603602081101561026457600080fd5b5035600160a060020a03166106b3565b6104286004803603606081101561028a57600080fd5b8101906020810181356401000000008111156102a557600080fd5b8201836020820111156102b757600080fd5b803590602001918460018302840111640100000000831117156102d957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561032c57600080fd5b82018360208201111561033e57600080fd5b8035906020019184600183028401116401000000008311171561036057600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156103b357600080fd5b8201836020820111156103c557600080fd5b803590602001918460018302840111640100000000831117156103e757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610726945050505050565b005b6101d96004803603602081101561044057600080fd5b5035610a59565b6104286004803603604081101561045d57600080fd5b5080359060200135600160a060020a0316610b23565b6100e86004803603602081101561048957600080fd5b5035610d89565b6100e8600480360360208110156104a657600080fd5b5035610db8565b6101d9600480360360208110156104c357600080fd5b5035610e22565b610428600480360360408110156104e057600080fd5b5080359060200135600160a060020a0316610ed5565b600061050182610fde565b15156105415760405160e560020a62461bcd0281526004018080602001828103825260318152602001806114146031913960400191505060405180910390fd5b50600090815260066020526040902054600160a060020a031690565b60006004826040518082805190602001908083835b602083106105915780518252601f199092019160209182019101610572565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054949350505050565b60606105d082610fde565b15156106105760405160e560020a62461bcd0281526004018080602001828103825260318152602001806114146031913960400191505060405180910390fd5b6000828152600560209081526040918290206002908101805484516001821615610100026000190190911692909204601f8101849004840283018401909452838252909290918301828280156106a75780601f1061067c576101008083540402835291602001916106a7565b820191906000526020600020905b81548152906001019060200180831161068a57829003601f168201915b50505050509050919050565b6000600160a060020a03821615156106ff5760405160e560020a62461bcd02815260040180806020018281038252602e815260200180611392602e913960400191505060405180910390fd5b600160a060020a038216600090815260016020526040902061072090610ff1565b92915050565b6107376107328461055d565b610fde565b1561078c576040805160e560020a62461bcd02815260206004820152601b60248201527f4974656d436f6e74726163743a206578697374656e74206974656d0000000000604482015290519081900360640190fd5b6107966000610ffc565b60006107a26000611005565b3360009081526001602052604090209091506107c4908263ffffffff61100916565b506107d76002823363ffffffff61101c16565b50806004856040518082805190602001908083835b6020831061080b5780518252601f1990920191602091820191016107ec565b51815160209384036101000a60001901801990921691161790529201948552506040805194859003820185209590955560808401855288845283810188905283850187905233606085015260008681526005825294909420835180519495919461087b94508593509101906112f6565b50602082810151805161089492600185019201906112f6565b50604082015180516108b09160028401916020909101906112f6565b5060608201518160030160006101000a815481600160a060020a030219169083600160a060020a031602179055509050508033600160a060020a03167fefebf1e9cc6cf1906f62a4ed1545a6c54e027262b597e06b8e00abc4a3d8f50186868660405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561095557818101518382015260200161093d565b50505050905090810190601f1680156109825780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b838110156109b557818101518382015260200161099d565b50505050905090810190601f1680156109e25780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b83811015610a155781810151838201526020016109fd565b50505050905090810190601f168015610a425780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a350505050565b6060610a6482610fde565b1515610aa45760405160e560020a62461bcd0281526004018080602001828103825260318152602001806114146031913960400191505060405180910390fd5b600560008381526020019081526020016000206001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106a75780601f1061067c576101008083540402835291602001916106a7565b600160a060020a0381161515610b6d5760405160e560020a62461bcd02815260040180806020018281038252602a815260200180611445602a913960400191505060405180910390fd5b610b7682610fde565b1515610bcc576040805160e560020a62461bcd02815260206004820152601f60248201527f4974656d436f6e74726163743a206e6f6e6578697374656e7420746f6b656e00604482015290519081900360640190fd5b6000610bd783610d89565b9050610be2836104f6565b600160a060020a031633600160a060020a03161480610c09575033600160a060020a038216145b1515610c495760405160e560020a62461bcd02815260040180806020018281038252602d8152602001806113c0602d913960400191505060405180910390fd5b600160a060020a0381166000908152600160205260409020610c71908463ffffffff61103a16565b50600160a060020a0382166000908152600160205260409020610c9a908463ffffffff61100916565b50610cad6002848463ffffffff61101c16565b506000838152600660205260409020805473ffffffffffffffffffffffffffffffffffffffff1916905582600160a060020a038381169083167fcd6e659e4c2e75c3bfe47fecaccf39aeb368116a0ee52afb532e07f6cba6c0d1610d1084610e22565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610d4a578181015183820152602001610d32565b50505050905090810190601f168015610d775780820380516001836020036101000a031916815260200191505b509250505060405180910390a4505050565b600061072082606060405190810160405280602e8152602001611498602e91396002919063ffffffff61104616565b6000610dc382610fde565b1515610e035760405160e560020a62461bcd0281526004018080602001828103825260318152602001806114146031913960400191505060405180910390fd5b50600090815260056020526040902060030154600160a060020a031690565b6060610e2d82610fde565b1515610e6d5760405160e560020a62461bcd0281526004018080602001828103825260318152602001806114146031913960400191505060405180910390fd5b60008281526005602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156106a75780601f1061067c576101008083540402835291602001916106a7565b6000610ee083610d89565b9050600160a060020a038281169082161415610f305760405160e560020a62461bcd0281526004018080602001828103825260278152602001806113ed6027913960400191505060405180910390fd5b33600160a060020a03821614610f7a5760405160e560020a62461bcd02815260040180806020018281038252602981526020018061146f6029913960400191505060405180910390fd5b6000838152600660205260409020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384169081179091558390337ffa07013316165a7c424e6be19a166ba4784477c405b341b4bb535985391e7baa610d1084610e22565b600061072060028363ffffffff61105316565b600061072082611005565b80546001019055565b5490565b6000611015838361105f565b9392505050565b60006110328484600160a060020a0385166110ab565b949350505050565b60006110158383611145565b600061103284848461120f565b600061101583836112de565b600061106b83836112de565b15156110a357508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610720565b506000610720565b6000828152600184016020526040812054801515611112575050604080518082018252838152602080820184815286546001818101895560008981528481209551600290930290950191825591519082015586548684528188019092529290912055611015565b845483908690600019840190811061112657fe5b9060005260206000209060020201600101819055506000915050611015565b60008181526001830160205260408120548015611205578354600019808301919081019060009087908390811061117857fe5b9060005260206000200154905080876000018481548110151561119757fe5b60009182526020808320909101929092558281526001898101909252604090209084019055865487908015156111c957fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050610720565b6000915050610720565b6000828152600184016020526040812054828115156112af5760405160e560020a62461bcd0281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561127457818101518382015260200161125c565b50505050905090810190601f1680156112a15780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b508454859060001983019081106112c257fe5b9060005260206000209060020201600101549150509392505050565b60009081526001919091016020526040902054151590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061133757805160ff1916838001178555611364565b82800160010185558215611364579182015b82811115611364578251825591602001919060010190611349565b50611370929150611374565b5090565b61138e91905b80821115611370576000815560010161137a565b9056fe4974656d436f6e74726163743a20636f756e7420717565727920666f7220746865207a65726f20616464726573734974656d436f6e74726163743a2063616c6c6572206973206e6f74206f776e6572206f7220617070726f76616c4974656d436f6e74726163743a20617070726f76616c20746f2063757272656e74206f776e65724974656d436f6e74726163743a20617070726f76656420717565727920666f72206e6f6e6578697374656e74206974656d4974656d436f6e74726163743a207472616e7366657220746f20746865207a65726f20616464726573734974656d436f6e74726163743a20617070726f76652063616c6c6572206973206e6f74206f776e65724974656d436f6e74726163743a206f776e657220717565727920666f72206e6f6e6578697374656e74206974656da165627a7a72305820170e36cce226a7e876127b1c9aaf2e9ed1b9a95a958502007863e22650a270000029`

// DeployItemContract deploys a new Ethereum contract, binding an instance of ItemContract to it.
func DeployItemContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ItemContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ItemContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ItemContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ItemContract{ItemContractCaller: ItemContractCaller{contract: contract}, ItemContractTransactor: ItemContractTransactor{contract: contract}, ItemContractFilterer: ItemContractFilterer{contract: contract}}, nil
}

// ItemContract is an auto generated Go binding around an Ethereum contract.
type ItemContract struct {
	ItemContractCaller     // Read-only binding to the contract
	ItemContractTransactor // Write-only binding to the contract
	ItemContractFilterer   // Log filterer for contract events
}

// ItemContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ItemContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItemContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ItemContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItemContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ItemContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItemContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ItemContractSession struct {
	Contract     *ItemContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ItemContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ItemContractCallerSession struct {
	Contract *ItemContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ItemContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ItemContractTransactorSession struct {
	Contract     *ItemContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ItemContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ItemContractRaw struct {
	Contract *ItemContract // Generic contract binding to access the raw methods on
}

// ItemContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ItemContractCallerRaw struct {
	Contract *ItemContractCaller // Generic read-only contract binding to access the raw methods on
}

// ItemContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ItemContractTransactorRaw struct {
	Contract *ItemContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewItemContract creates a new instance of ItemContract, bound to a specific deployed contract.
func NewItemContract(address common.Address, backend bind.ContractBackend) (*ItemContract, error) {
	contract, err := bindItemContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ItemContract{ItemContractCaller: ItemContractCaller{contract: contract}, ItemContractTransactor: ItemContractTransactor{contract: contract}, ItemContractFilterer: ItemContractFilterer{contract: contract}}, nil
}

// NewItemContractCaller creates a new read-only instance of ItemContract, bound to a specific deployed contract.
func NewItemContractCaller(address common.Address, caller bind.ContractCaller) (*ItemContractCaller, error) {
	contract, err := bindItemContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ItemContractCaller{contract: contract}, nil
}

// NewItemContractTransactor creates a new write-only instance of ItemContract, bound to a specific deployed contract.
func NewItemContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ItemContractTransactor, error) {
	contract, err := bindItemContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ItemContractTransactor{contract: contract}, nil
}

// NewItemContractFilterer creates a new log filterer instance of ItemContract, bound to a specific deployed contract.
func NewItemContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ItemContractFilterer, error) {
	contract, err := bindItemContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ItemContractFilterer{contract: contract}, nil
}

// bindItemContract binds a generic wrapper to an already deployed contract.
func bindItemContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ItemContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ItemContract *ItemContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ItemContract.Contract.ItemContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ItemContract *ItemContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ItemContract.Contract.ItemContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ItemContract *ItemContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ItemContract.Contract.ItemContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ItemContract *ItemContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ItemContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ItemContract *ItemContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ItemContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ItemContract *ItemContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ItemContract.Contract.contract.Transact(opts, method, params...)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(tokenId uint256) constant returns(address)
func (_ItemContract *ItemContractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ItemContract.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(tokenId uint256) constant returns(address)
func (_ItemContract *ItemContractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ItemContract.Contract.GetApproved(&_ItemContract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(tokenId uint256) constant returns(address)
func (_ItemContract *ItemContractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ItemContract.Contract.GetApproved(&_ItemContract.CallOpts, tokenId)
}

// GetCount is a free data retrieval call binding the contract method 0x4f0cd27b.
//
// Solidity: function getCount(owner address) constant returns(uint256)
func (_ItemContract *ItemContractCaller) GetCount(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ItemContract.contract.Call(opts, out, "getCount", owner)
	return *ret0, err
}

// GetCount is a free data retrieval call binding the contract method 0x4f0cd27b.
//
// Solidity: function getCount(owner address) constant returns(uint256)
func (_ItemContract *ItemContractSession) GetCount(owner common.Address) (*big.Int, error) {
	return _ItemContract.Contract.GetCount(&_ItemContract.CallOpts, owner)
}

// GetCount is a free data retrieval call binding the contract method 0x4f0cd27b.
//
// Solidity: function getCount(owner address) constant returns(uint256)
func (_ItemContract *ItemContractCallerSession) GetCount(owner common.Address) (*big.Int, error) {
	return _ItemContract.Contract.GetCount(&_ItemContract.CallOpts, owner)
}

// GetCreator is a free data retrieval call binding the contract method 0xd48e638a.
//
// Solidity: function getCreator(tokenId uint256) constant returns(address)
func (_ItemContract *ItemContractCaller) GetCreator(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ItemContract.contract.Call(opts, out, "getCreator", tokenId)
	return *ret0, err
}

// GetCreator is a free data retrieval call binding the contract method 0xd48e638a.
//
// Solidity: function getCreator(tokenId uint256) constant returns(address)
func (_ItemContract *ItemContractSession) GetCreator(tokenId *big.Int) (common.Address, error) {
	return _ItemContract.Contract.GetCreator(&_ItemContract.CallOpts, tokenId)
}

// GetCreator is a free data retrieval call binding the contract method 0xd48e638a.
//
// Solidity: function getCreator(tokenId uint256) constant returns(address)
func (_ItemContract *ItemContractCallerSession) GetCreator(tokenId *big.Int) (common.Address, error) {
	return _ItemContract.Contract.GetCreator(&_ItemContract.CallOpts, tokenId)
}

// GetDCI is a free data retrieval call binding the contract method 0xdb157339.
//
// Solidity: function getDCI(tokenId uint256) constant returns(string)
func (_ItemContract *ItemContractCaller) GetDCI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ItemContract.contract.Call(opts, out, "getDCI", tokenId)
	return *ret0, err
}

// GetDCI is a free data retrieval call binding the contract method 0xdb157339.
//
// Solidity: function getDCI(tokenId uint256) constant returns(string)
func (_ItemContract *ItemContractSession) GetDCI(tokenId *big.Int) (string, error) {
	return _ItemContract.Contract.GetDCI(&_ItemContract.CallOpts, tokenId)
}

// GetDCI is a free data retrieval call binding the contract method 0xdb157339.
//
// Solidity: function getDCI(tokenId uint256) constant returns(string)
func (_ItemContract *ItemContractCallerSession) GetDCI(tokenId *big.Int) (string, error) {
	return _ItemContract.Contract.GetDCI(&_ItemContract.CallOpts, tokenId)
}

// GetName is a free data retrieval call binding the contract method 0x6b8ff574.
//
// Solidity: function getName(tokenId uint256) constant returns(string)
func (_ItemContract *ItemContractCaller) GetName(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ItemContract.contract.Call(opts, out, "getName", tokenId)
	return *ret0, err
}

// GetName is a free data retrieval call binding the contract method 0x6b8ff574.
//
// Solidity: function getName(tokenId uint256) constant returns(string)
func (_ItemContract *ItemContractSession) GetName(tokenId *big.Int) (string, error) {
	return _ItemContract.Contract.GetName(&_ItemContract.CallOpts, tokenId)
}

// GetName is a free data retrieval call binding the contract method 0x6b8ff574.
//
// Solidity: function getName(tokenId uint256) constant returns(string)
func (_ItemContract *ItemContractCallerSession) GetName(tokenId *big.Int) (string, error) {
	return _ItemContract.Contract.GetName(&_ItemContract.CallOpts, tokenId)
}

// GetOwner is a free data retrieval call binding the contract method 0xc41a360a.
//
// Solidity: function getOwner(tokenId uint256) constant returns(address)
func (_ItemContract *ItemContractCaller) GetOwner(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ItemContract.contract.Call(opts, out, "getOwner", tokenId)
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0xc41a360a.
//
// Solidity: function getOwner(tokenId uint256) constant returns(address)
func (_ItemContract *ItemContractSession) GetOwner(tokenId *big.Int) (common.Address, error) {
	return _ItemContract.Contract.GetOwner(&_ItemContract.CallOpts, tokenId)
}

// GetOwner is a free data retrieval call binding the contract method 0xc41a360a.
//
// Solidity: function getOwner(tokenId uint256) constant returns(address)
func (_ItemContract *ItemContractCallerSession) GetOwner(tokenId *big.Int) (common.Address, error) {
	return _ItemContract.Contract.GetOwner(&_ItemContract.CallOpts, tokenId)
}

// GetTokenId is a free data retrieval call binding the contract method 0x1e7663bc.
//
// Solidity: function getTokenId(dci string) constant returns(uint256)
func (_ItemContract *ItemContractCaller) GetTokenId(opts *bind.CallOpts, dci string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ItemContract.contract.Call(opts, out, "getTokenId", dci)
	return *ret0, err
}

// GetTokenId is a free data retrieval call binding the contract method 0x1e7663bc.
//
// Solidity: function getTokenId(dci string) constant returns(uint256)
func (_ItemContract *ItemContractSession) GetTokenId(dci string) (*big.Int, error) {
	return _ItemContract.Contract.GetTokenId(&_ItemContract.CallOpts, dci)
}

// GetTokenId is a free data retrieval call binding the contract method 0x1e7663bc.
//
// Solidity: function getTokenId(dci string) constant returns(uint256)
func (_ItemContract *ItemContractCallerSession) GetTokenId(dci string) (*big.Int, error) {
	return _ItemContract.Contract.GetTokenId(&_ItemContract.CallOpts, dci)
}

// GetTokenURI is a free data retrieval call binding the contract method 0x3bb3a24d.
//
// Solidity: function getTokenURI(tokenId uint256) constant returns(string)
func (_ItemContract *ItemContractCaller) GetTokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ItemContract.contract.Call(opts, out, "getTokenURI", tokenId)
	return *ret0, err
}

// GetTokenURI is a free data retrieval call binding the contract method 0x3bb3a24d.
//
// Solidity: function getTokenURI(tokenId uint256) constant returns(string)
func (_ItemContract *ItemContractSession) GetTokenURI(tokenId *big.Int) (string, error) {
	return _ItemContract.Contract.GetTokenURI(&_ItemContract.CallOpts, tokenId)
}

// GetTokenURI is a free data retrieval call binding the contract method 0x3bb3a24d.
//
// Solidity: function getTokenURI(tokenId uint256) constant returns(string)
func (_ItemContract *ItemContractCallerSession) GetTokenURI(tokenId *big.Int) (string, error) {
	return _ItemContract.Contract.GetTokenURI(&_ItemContract.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0xfebe4909.
//
// Solidity: function approve(tokenId uint256, to address) returns()
func (_ItemContract *ItemContractTransactor) Approve(opts *bind.TransactOpts, tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _ItemContract.contract.Transact(opts, "approve", tokenId, to)
}

// Approve is a paid mutator transaction binding the contract method 0xfebe4909.
//
// Solidity: function approve(tokenId uint256, to address) returns()
func (_ItemContract *ItemContractSession) Approve(tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _ItemContract.Contract.Approve(&_ItemContract.TransactOpts, tokenId, to)
}

// Approve is a paid mutator transaction binding the contract method 0xfebe4909.
//
// Solidity: function approve(tokenId uint256, to address) returns()
func (_ItemContract *ItemContractTransactorSession) Approve(tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _ItemContract.Contract.Approve(&_ItemContract.TransactOpts, tokenId, to)
}

// Create is a paid mutator transaction binding the contract method 0x5d28560a.
//
// Solidity: function create(dci string, name string, tokenURI string) returns()
func (_ItemContract *ItemContractTransactor) Create(opts *bind.TransactOpts, dci string, name string, tokenURI string) (*types.Transaction, error) {
	return _ItemContract.contract.Transact(opts, "create", dci, name, tokenURI)
}

// Create is a paid mutator transaction binding the contract method 0x5d28560a.
//
// Solidity: function create(dci string, name string, tokenURI string) returns()
func (_ItemContract *ItemContractSession) Create(dci string, name string, tokenURI string) (*types.Transaction, error) {
	return _ItemContract.Contract.Create(&_ItemContract.TransactOpts, dci, name, tokenURI)
}

// Create is a paid mutator transaction binding the contract method 0x5d28560a.
//
// Solidity: function create(dci string, name string, tokenURI string) returns()
func (_ItemContract *ItemContractTransactorSession) Create(dci string, name string, tokenURI string) (*types.Transaction, error) {
	return _ItemContract.Contract.Create(&_ItemContract.TransactOpts, dci, name, tokenURI)
}

// Transfer is a paid mutator transaction binding the contract method 0xb7760c8f.
//
// Solidity: function transfer(tokenId uint256, to address) returns()
func (_ItemContract *ItemContractTransactor) Transfer(opts *bind.TransactOpts, tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _ItemContract.contract.Transact(opts, "transfer", tokenId, to)
}

// Transfer is a paid mutator transaction binding the contract method 0xb7760c8f.
//
// Solidity: function transfer(tokenId uint256, to address) returns()
func (_ItemContract *ItemContractSession) Transfer(tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _ItemContract.Contract.Transfer(&_ItemContract.TransactOpts, tokenId, to)
}

// Transfer is a paid mutator transaction binding the contract method 0xb7760c8f.
//
// Solidity: function transfer(tokenId uint256, to address) returns()
func (_ItemContract *ItemContractTransactorSession) Transfer(tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _ItemContract.Contract.Transfer(&_ItemContract.TransactOpts, tokenId, to)
}

// ItemContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ItemContract contract.
type ItemContractApprovalIterator struct {
	Event *ItemContractApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ItemContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ItemContractApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ItemContractApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ItemContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ItemContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ItemContractApproval represents a Approval event raised by the ItemContract contract.
type ItemContractApproval struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Dci     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0xfa07013316165a7c424e6be19a166ba4784477c405b341b4bb535985391e7baa.
//
// Solidity: e Approval(from indexed address, to indexed address, tokenId indexed uint256, dci string)
func (_ItemContract *ItemContractFilterer) FilterApproval(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ItemContractApprovalIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ItemContract.contract.FilterLogs(opts, "Approval", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ItemContractApprovalIterator{contract: _ItemContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0xfa07013316165a7c424e6be19a166ba4784477c405b341b4bb535985391e7baa.
//
// Solidity: e Approval(from indexed address, to indexed address, tokenId indexed uint256, dci string)
func (_ItemContract *ItemContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ItemContractApproval, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ItemContract.contract.WatchLogs(opts, "Approval", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ItemContractApproval)
				if err := _ItemContract.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ItemContractCreateIterator is returned from FilterCreate and is used to iterate over the raw logs and unpacked data for Create events raised by the ItemContract contract.
type ItemContractCreateIterator struct {
	Event *ItemContractCreate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ItemContractCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ItemContractCreate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ItemContractCreate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ItemContractCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ItemContractCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ItemContractCreate represents a Create event raised by the ItemContract contract.
type ItemContractCreate struct {
	From     common.Address
	TokenId  *big.Int
	Dci      string
	Name     string
	TokenURI string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCreate is a free log retrieval operation binding the contract event 0xefebf1e9cc6cf1906f62a4ed1545a6c54e027262b597e06b8e00abc4a3d8f501.
//
// Solidity: e Create(from indexed address, tokenId indexed uint256, dci string, name string, tokenURI string)
func (_ItemContract *ItemContractFilterer) FilterCreate(opts *bind.FilterOpts, from []common.Address, tokenId []*big.Int) (*ItemContractCreateIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ItemContract.contract.FilterLogs(opts, "Create", fromRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ItemContractCreateIterator{contract: _ItemContract.contract, event: "Create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0xefebf1e9cc6cf1906f62a4ed1545a6c54e027262b597e06b8e00abc4a3d8f501.
//
// Solidity: e Create(from indexed address, tokenId indexed uint256, dci string, name string, tokenURI string)
func (_ItemContract *ItemContractFilterer) WatchCreate(opts *bind.WatchOpts, sink chan<- *ItemContractCreate, from []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ItemContract.contract.WatchLogs(opts, "Create", fromRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ItemContractCreate)
				if err := _ItemContract.contract.UnpackLog(event, "Create", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ItemContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ItemContract contract.
type ItemContractTransferIterator struct {
	Event *ItemContractTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ItemContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ItemContractTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ItemContractTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ItemContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ItemContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ItemContractTransfer represents a Transfer event raised by the ItemContract contract.
type ItemContractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Dci     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xcd6e659e4c2e75c3bfe47fecaccf39aeb368116a0ee52afb532e07f6cba6c0d1.
//
// Solidity: e Transfer(from indexed address, to indexed address, tokenId indexed uint256, dci string)
func (_ItemContract *ItemContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ItemContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ItemContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ItemContractTransferIterator{contract: _ItemContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xcd6e659e4c2e75c3bfe47fecaccf39aeb368116a0ee52afb532e07f6cba6c0d1.
//
// Solidity: e Transfer(from indexed address, to indexed address, tokenId indexed uint256, dci string)
func (_ItemContract *ItemContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ItemContractTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ItemContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ItemContractTransfer)
				if err := _ItemContract.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ReentrancyGuardABI is the input ABI used to generate the binding from.
const ReentrancyGuardABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ReentrancyGuardBin is the compiled bytecode used for deploying new contracts.
const ReentrancyGuardBin = `0x`

// DeployReentrancyGuard deploys a new Ethereum contract, binding an instance of ReentrancyGuard to it.
func DeployReentrancyGuard(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReentrancyGuard, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReentrancyGuardBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// ReentrancyGuard is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuard struct {
	ReentrancyGuardCaller     // Read-only binding to the contract
	ReentrancyGuardTransactor // Write-only binding to the contract
	ReentrancyGuardFilterer   // Log filterer for contract events
}

// ReentrancyGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardSession struct {
	Contract     *ReentrancyGuard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardCallerSession struct {
	Contract *ReentrancyGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReentrancyGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardTransactorSession struct {
	Contract     *ReentrancyGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReentrancyGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardRaw struct {
	Contract *ReentrancyGuard // Generic contract binding to access the raw methods on
}

// ReentrancyGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardCallerRaw struct {
	Contract *ReentrancyGuardCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactorRaw struct {
	Contract *ReentrancyGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuard creates a new instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuard(address common.Address, backend bind.ContractBackend) (*ReentrancyGuard, error) {
	contract, err := bindReentrancyGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// NewReentrancyGuardCaller creates a new read-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardCaller, error) {
	contract, err := bindReentrancyGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardCaller{contract: contract}, nil
}

// NewReentrancyGuardTransactor creates a new write-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardTransactor, error) {
	contract, err := bindReentrancyGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardTransactor{contract: contract}, nil
}

// NewReentrancyGuardFilterer creates a new log filterer instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardFilterer, error) {
	contract, err := bindReentrancyGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardFilterer{contract: contract}, nil
}

// bindReentrancyGuard binds a generic wrapper to an already deployed contract.
func bindReentrancyGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.ReentrancyGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820d8f1e0516007cc5aafd39c88fcdfd650219c492c1c1f0bed233a35b14e397e900029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
