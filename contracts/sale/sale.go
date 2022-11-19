// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sale

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

// SaleContractABI is the input ABI used to generate the binding from.
const SaleContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"create\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getBuyer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"confirm\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getSeller\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"itemContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"dci\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Create\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"dci\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Cancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"dci\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Confirm\",\"type\":\"event\"}]"

// SaleContractBin is the compiled bytecode used for deploying new contracts.
const SaleContractBin = `0x608060405234801561001057600080fd5b50604051602080610fb38339810180604052602081101561003057600080fd5b505160008054600160a060020a039092166101000261010060a860020a031960ff1990931660011792909216919091179055610f42806100716000396000f3fe60806040526004361061008d576000357c0100000000000000000000000000000000000000000000000000000000900480635bf608b81161006b5780635bf608b81461013b578063ba0179b514610181578063d6a9de511461019e578063e7572230146101c85761008d565b806323cdf3651461009257806340e58ee5146100d35780634f558e79146100fd575b600080fd5b34801561009e57600080fd5b506100d1600480360360608110156100b557600080fd5b5080359060208101359060400135600160a060020a0316610204565b005b3480156100df57600080fd5b506100d1600480360360208110156100f657600080fd5b503561051d565b34801561010957600080fd5b506101276004803603602081101561012057600080fd5b503561087e565b604080519115158252519081900360200190f35b34801561014757600080fd5b506101656004803603602081101561015e57600080fd5b503561089b565b60408051600160a060020a039092168252519081900360200190f35b6100d16004803603602081101561019757600080fd5b50356108ba565b3480156101aa57600080fd5b50610165600480360360208110156101c157600080fd5b5035610d2b565b3480156101d457600080fd5b506101f2600480360360208110156101eb57600080fd5b5035610d46565b60408051918252519081900360200190f35b61020d8361087e565b15610262576040805160e560020a62461bcd02815260206004820152601d60248201527f53616c65436f6e74726163743a206974656d206973206f6e2073616c65000000604482015290519081900360640190fd5b60408051606081018252338152600160a060020a03838116602080840191825283850187815260008981526001928390528681209551865490861673ffffffffffffffffffffffffffffffffffffffff199182161787559351928601805493861693909416929092179092559051600290930192909255815483517fb7760c8f0000000000000000000000000000000000000000000000000000000081526004810188905230602482015293516101009091049091169263b7760c8f92604480830193919282900301818387803b15801561033c57600080fd5b505af1158015610350573d6000803e3d6000fd5b505060008054604080517fdb1573390000000000000000000000000000000000000000000000000000000081526004810189905290518895507e48209bd46e9c7af682cae0cfbd459137e1a1d6cfe0b90f9767f3ef1c6f6bf09450610100909204600160a060020a03169263db15733992602480840193829003018186803b1580156103db57600080fd5b505afa1580156103ef573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561041857600080fd5b81019080805164010000000081111561043057600080fd5b8201602081018481111561044357600080fd5b815164010000000081118282018710171561045d57600080fd5b5050929190505050338486604051808060200185600160a060020a0316600160a060020a0316815260200184600160a060020a0316600160a060020a03168152602001838152602001828103825286818151815260200191508051906020019080838360005b838110156104db5781810151838201526020016104c3565b50505050905090810190601f1680156105085780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a2505050565b6105268161087e565b15156105665760405160e560020a62461bcd028152600401808060200182810382526021815260200180610ef66021913960400191505060405180910390fd5b3361057082610d2b565b600160a060020a0316146105b85760405160e560020a62461bcd028152600401808060200182810382526022815260200180610e6f6022913960400191505060405180910390fd5b6105c0610e4e565b50600081815260016020818152604080842081516060810183528154600160a060020a038082168352838701805480831685890152600286018054868901528b8b529890975273ffffffffffffffffffffffffffffffffffffffff19928316909455941690915592849055835481517fb7760c8f00000000000000000000000000000000000000000000000000000000815260048101879052336024820152915193946101009091049092169263b7760c8f9260448084019391929182900301818387803b15801561069157600080fd5b505af11580156106a5573d6000803e3d6000fd5b505060008054604080517fdb1573390000000000000000000000000000000000000000000000000000000081526004810188905290518795507f44e2a04ad594fb3b55cdf6a56531388a6ead5e22127700b9e7e85498510d32f09450610100909204600160a060020a03169263db15733992602480840193829003018186803b15801561073157600080fd5b505afa158015610745573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561076e57600080fd5b81019080805164010000000081111561078657600080fd5b8201602081018481111561079957600080fd5b81516401000000008111828201871017156107b357600080fd5b5050929190505050836000015184602001518560400151604051808060200185600160a060020a0316600160a060020a0316815260200184600160a060020a0316600160a060020a03168152602001838152602001828103825286818151815260200191508051906020019080838360005b8381101561083d578181015183820152602001610825565b50505050905090810190601f16801561086a5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a25050565b600090815260016020526040902054600160a060020a0316151590565b60009081526001602081905260409091200154600160a060020a031690565b60005460ff161515610916576040805160e560020a62461bcd02815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6000805460ff191690556109298161087e565b15156109695760405160e560020a62461bcd028152600401808060200182810382526021815260200180610ef66021913960400191505060405180910390fd5b610971610e4e565b5060008181526001602081815260409283902083516060810185528154600160a060020a03908116825293820154909316918301919091526002015491810182905290341015610a0b576040805160e560020a62461bcd02815260206004820152601e60248201527f53616c65436f6e74726163743a206e6f7420656e6f7567682076616c75650000604482015290519081900360640190fd5b6020810151600160a060020a031615610a6c576020810151600160a060020a03163314610a6c5760405160e560020a62461bcd02815260040180806020018281038252602b815260200180610ecb602b913960400191505060405180910390fd5b60008054604080517fb7760c8f000000000000000000000000000000000000000000000000000000008152600481018690523360248201529051610100909204600160a060020a03169263b7760c8f9260448084019382900301818387803b158015610ad757600080fd5b505af1158015610aeb573d6000803e3d6000fd5b50505060408201518251610b0f9250600160a060020a03169063ffffffff610d5b16565b6000828152600160208190526040808320805473ffffffffffffffffffffffffffffffffffffffff19908116825592810180549093169092556002909101829055815481517fdb15733900000000000000000000000000000000000000000000000000000000815260048101869052915185937f4bf5cf8e0c546146096c5ee70738e1612ac80ba22ebbe1e06cfb156c3a7c77a993610100909304600160a060020a03169263db1573399260248083019392829003018186803b158015610bd557600080fd5b505afa158015610be9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610c1257600080fd5b810190808051640100000000811115610c2a57600080fd5b82016020810184811115610c3d57600080fd5b8151640100000000811182820187101715610c5757600080fd5b50509291905050508360000151338560400151604051808060200185600160a060020a0316600160a060020a0316815260200184600160a060020a0316600160a060020a03168152602001838152602001828103825286818151815260200191508051906020019080838360005b83811015610cdd578181015183820152602001610cc5565b50505050905090810190601f168015610d0a5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a250506000805460ff19166001179055565b600090815260016020526040902054600160a060020a031690565b60009081526001602052604090206002015490565b3031811115610db4576040805160e560020a62461bcd02815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e6365000000604482015290519081900360640190fd5b604051600090600160a060020a0384169083908381818185875af1925050503d8060008114610dff576040519150601f19603f3d011682016040523d82523d6000602084013e610e04565b606091505b50509050801515610e495760405160e560020a62461bcd02815260040180806020018281038252603a815260200180610e91603a913960400191505060405180910390fd5b505050565b60408051606081018252600080825260208201819052918101919091529056fe53616c65436f6e74726163743a2063616c6c6572206973206e6f742073656c6c6572416464726573733a20756e61626c6520746f2073656e642076616c75652c20726563697069656e74206d6179206861766520726576657274656453616c65436f6e74726163743a2063616c6c6572206973206e6f74206170706f696e74206163636f756e7453616c65436f6e74726163743a206974656d206973206e6f74206f6e2073616c65a165627a7a7230582035b211572c3cebe2a76a1d5a5ccebaf80a75614f676aac18799d1fb965d375100029`

// DeploySaleContract deploys a new Ethereum contract, binding an instance of SaleContract to it.
func DeploySaleContract(auth *bind.TransactOpts, backend bind.ContractBackend, itemContract common.Address) (common.Address, *types.Transaction, *SaleContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SaleContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SaleContractBin), backend, itemContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SaleContract{SaleContractCaller: SaleContractCaller{contract: contract}, SaleContractTransactor: SaleContractTransactor{contract: contract}, SaleContractFilterer: SaleContractFilterer{contract: contract}}, nil
}

// SaleContract is an auto generated Go binding around an Ethereum contract.
type SaleContract struct {
	SaleContractCaller     // Read-only binding to the contract
	SaleContractTransactor // Write-only binding to the contract
	SaleContractFilterer   // Log filterer for contract events
}

// SaleContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SaleContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SaleContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SaleContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SaleContractSession struct {
	Contract     *SaleContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaleContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SaleContractCallerSession struct {
	Contract *SaleContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SaleContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SaleContractTransactorSession struct {
	Contract     *SaleContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SaleContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SaleContractRaw struct {
	Contract *SaleContract // Generic contract binding to access the raw methods on
}

// SaleContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SaleContractCallerRaw struct {
	Contract *SaleContractCaller // Generic read-only contract binding to access the raw methods on
}

// SaleContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SaleContractTransactorRaw struct {
	Contract *SaleContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSaleContract creates a new instance of SaleContract, bound to a specific deployed contract.
func NewSaleContract(address common.Address, backend bind.ContractBackend) (*SaleContract, error) {
	contract, err := bindSaleContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SaleContract{SaleContractCaller: SaleContractCaller{contract: contract}, SaleContractTransactor: SaleContractTransactor{contract: contract}, SaleContractFilterer: SaleContractFilterer{contract: contract}}, nil
}

// NewSaleContractCaller creates a new read-only instance of SaleContract, bound to a specific deployed contract.
func NewSaleContractCaller(address common.Address, caller bind.ContractCaller) (*SaleContractCaller, error) {
	contract, err := bindSaleContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SaleContractCaller{contract: contract}, nil
}

// NewSaleContractTransactor creates a new write-only instance of SaleContract, bound to a specific deployed contract.
func NewSaleContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SaleContractTransactor, error) {
	contract, err := bindSaleContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SaleContractTransactor{contract: contract}, nil
}

// NewSaleContractFilterer creates a new log filterer instance of SaleContract, bound to a specific deployed contract.
func NewSaleContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SaleContractFilterer, error) {
	contract, err := bindSaleContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SaleContractFilterer{contract: contract}, nil
}

// bindSaleContract binds a generic wrapper to an already deployed contract.
func bindSaleContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SaleContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SaleContract *SaleContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SaleContract.Contract.SaleContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SaleContract *SaleContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleContract.Contract.SaleContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SaleContract *SaleContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SaleContract.Contract.SaleContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SaleContract *SaleContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SaleContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SaleContract *SaleContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SaleContract *SaleContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SaleContract.Contract.contract.Transact(opts, method, params...)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(tokenId uint256) constant returns(bool)
func (_SaleContract *SaleContractCaller) Exists(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SaleContract.contract.Call(opts, out, "exists", tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(tokenId uint256) constant returns(bool)
func (_SaleContract *SaleContractSession) Exists(tokenId *big.Int) (bool, error) {
	return _SaleContract.Contract.Exists(&_SaleContract.CallOpts, tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(tokenId uint256) constant returns(bool)
func (_SaleContract *SaleContractCallerSession) Exists(tokenId *big.Int) (bool, error) {
	return _SaleContract.Contract.Exists(&_SaleContract.CallOpts, tokenId)
}

// GetBuyer is a free data retrieval call binding the contract method 0x5bf608b8.
//
// Solidity: function getBuyer(tokenId uint256) constant returns(address)
func (_SaleContract *SaleContractCaller) GetBuyer(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SaleContract.contract.Call(opts, out, "getBuyer", tokenId)
	return *ret0, err
}

// GetBuyer is a free data retrieval call binding the contract method 0x5bf608b8.
//
// Solidity: function getBuyer(tokenId uint256) constant returns(address)
func (_SaleContract *SaleContractSession) GetBuyer(tokenId *big.Int) (common.Address, error) {
	return _SaleContract.Contract.GetBuyer(&_SaleContract.CallOpts, tokenId)
}

// GetBuyer is a free data retrieval call binding the contract method 0x5bf608b8.
//
// Solidity: function getBuyer(tokenId uint256) constant returns(address)
func (_SaleContract *SaleContractCallerSession) GetBuyer(tokenId *big.Int) (common.Address, error) {
	return _SaleContract.Contract.GetBuyer(&_SaleContract.CallOpts, tokenId)
}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(tokenId uint256) constant returns(uint256)
func (_SaleContract *SaleContractCaller) GetPrice(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SaleContract.contract.Call(opts, out, "getPrice", tokenId)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(tokenId uint256) constant returns(uint256)
func (_SaleContract *SaleContractSession) GetPrice(tokenId *big.Int) (*big.Int, error) {
	return _SaleContract.Contract.GetPrice(&_SaleContract.CallOpts, tokenId)
}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(tokenId uint256) constant returns(uint256)
func (_SaleContract *SaleContractCallerSession) GetPrice(tokenId *big.Int) (*big.Int, error) {
	return _SaleContract.Contract.GetPrice(&_SaleContract.CallOpts, tokenId)
}

// GetSeller is a free data retrieval call binding the contract method 0xd6a9de51.
//
// Solidity: function getSeller(tokenId uint256) constant returns(address)
func (_SaleContract *SaleContractCaller) GetSeller(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SaleContract.contract.Call(opts, out, "getSeller", tokenId)
	return *ret0, err
}

// GetSeller is a free data retrieval call binding the contract method 0xd6a9de51.
//
// Solidity: function getSeller(tokenId uint256) constant returns(address)
func (_SaleContract *SaleContractSession) GetSeller(tokenId *big.Int) (common.Address, error) {
	return _SaleContract.Contract.GetSeller(&_SaleContract.CallOpts, tokenId)
}

// GetSeller is a free data retrieval call binding the contract method 0xd6a9de51.
//
// Solidity: function getSeller(tokenId uint256) constant returns(address)
func (_SaleContract *SaleContractCallerSession) GetSeller(tokenId *big.Int) (common.Address, error) {
	return _SaleContract.Contract.GetSeller(&_SaleContract.CallOpts, tokenId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(tokenId uint256) returns()
func (_SaleContract *SaleContractTransactor) Cancel(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _SaleContract.contract.Transact(opts, "cancel", tokenId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(tokenId uint256) returns()
func (_SaleContract *SaleContractSession) Cancel(tokenId *big.Int) (*types.Transaction, error) {
	return _SaleContract.Contract.Cancel(&_SaleContract.TransactOpts, tokenId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(tokenId uint256) returns()
func (_SaleContract *SaleContractTransactorSession) Cancel(tokenId *big.Int) (*types.Transaction, error) {
	return _SaleContract.Contract.Cancel(&_SaleContract.TransactOpts, tokenId)
}

// Confirm is a paid mutator transaction binding the contract method 0xba0179b5.
//
// Solidity: function confirm(tokenId uint256) returns()
func (_SaleContract *SaleContractTransactor) Confirm(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _SaleContract.contract.Transact(opts, "confirm", tokenId)
}

// Confirm is a paid mutator transaction binding the contract method 0xba0179b5.
//
// Solidity: function confirm(tokenId uint256) returns()
func (_SaleContract *SaleContractSession) Confirm(tokenId *big.Int) (*types.Transaction, error) {
	return _SaleContract.Contract.Confirm(&_SaleContract.TransactOpts, tokenId)
}

// Confirm is a paid mutator transaction binding the contract method 0xba0179b5.
//
// Solidity: function confirm(tokenId uint256) returns()
func (_SaleContract *SaleContractTransactorSession) Confirm(tokenId *big.Int) (*types.Transaction, error) {
	return _SaleContract.Contract.Confirm(&_SaleContract.TransactOpts, tokenId)
}

// Create is a paid mutator transaction binding the contract method 0x23cdf365.
//
// Solidity: function create(tokenId uint256, price uint256, buyer address) returns()
func (_SaleContract *SaleContractTransactor) Create(opts *bind.TransactOpts, tokenId *big.Int, price *big.Int, buyer common.Address) (*types.Transaction, error) {
	return _SaleContract.contract.Transact(opts, "create", tokenId, price, buyer)
}

// Create is a paid mutator transaction binding the contract method 0x23cdf365.
//
// Solidity: function create(tokenId uint256, price uint256, buyer address) returns()
func (_SaleContract *SaleContractSession) Create(tokenId *big.Int, price *big.Int, buyer common.Address) (*types.Transaction, error) {
	return _SaleContract.Contract.Create(&_SaleContract.TransactOpts, tokenId, price, buyer)
}

// Create is a paid mutator transaction binding the contract method 0x23cdf365.
//
// Solidity: function create(tokenId uint256, price uint256, buyer address) returns()
func (_SaleContract *SaleContractTransactorSession) Create(tokenId *big.Int, price *big.Int, buyer common.Address) (*types.Transaction, error) {
	return _SaleContract.Contract.Create(&_SaleContract.TransactOpts, tokenId, price, buyer)
}

// SaleContractCancelIterator is returned from FilterCancel and is used to iterate over the raw logs and unpacked data for Cancel events raised by the SaleContract contract.
type SaleContractCancelIterator struct {
	Event *SaleContractCancel // Event containing the contract specifics and raw log

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
func (it *SaleContractCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleContractCancel)
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
		it.Event = new(SaleContractCancel)
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
func (it *SaleContractCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleContractCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleContractCancel represents a Cancel event raised by the SaleContract contract.
type SaleContractCancel struct {
	TokenId *big.Int
	Dci     string
	Seller  common.Address
	Buyer   common.Address
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCancel is a free log retrieval operation binding the contract event 0x44e2a04ad594fb3b55cdf6a56531388a6ead5e22127700b9e7e85498510d32f0.
//
// Solidity: e Cancel(tokenId indexed uint256, dci string, seller address, buyer address, price uint256)
func (_SaleContract *SaleContractFilterer) FilterCancel(opts *bind.FilterOpts, tokenId []*big.Int) (*SaleContractCancelIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SaleContract.contract.FilterLogs(opts, "Cancel", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SaleContractCancelIterator{contract: _SaleContract.contract, event: "Cancel", logs: logs, sub: sub}, nil
}

// WatchCancel is a free log subscription operation binding the contract event 0x44e2a04ad594fb3b55cdf6a56531388a6ead5e22127700b9e7e85498510d32f0.
//
// Solidity: e Cancel(tokenId indexed uint256, dci string, seller address, buyer address, price uint256)
func (_SaleContract *SaleContractFilterer) WatchCancel(opts *bind.WatchOpts, sink chan<- *SaleContractCancel, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SaleContract.contract.WatchLogs(opts, "Cancel", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleContractCancel)
				if err := _SaleContract.contract.UnpackLog(event, "Cancel", log); err != nil {
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

// SaleContractConfirmIterator is returned from FilterConfirm and is used to iterate over the raw logs and unpacked data for Confirm events raised by the SaleContract contract.
type SaleContractConfirmIterator struct {
	Event *SaleContractConfirm // Event containing the contract specifics and raw log

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
func (it *SaleContractConfirmIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleContractConfirm)
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
		it.Event = new(SaleContractConfirm)
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
func (it *SaleContractConfirmIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleContractConfirmIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleContractConfirm represents a Confirm event raised by the SaleContract contract.
type SaleContractConfirm struct {
	TokenId *big.Int
	Dci     string
	Seller  common.Address
	Buyer   common.Address
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConfirm is a free log retrieval operation binding the contract event 0x4bf5cf8e0c546146096c5ee70738e1612ac80ba22ebbe1e06cfb156c3a7c77a9.
//
// Solidity: e Confirm(tokenId indexed uint256, dci string, seller address, buyer address, price uint256)
func (_SaleContract *SaleContractFilterer) FilterConfirm(opts *bind.FilterOpts, tokenId []*big.Int) (*SaleContractConfirmIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SaleContract.contract.FilterLogs(opts, "Confirm", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SaleContractConfirmIterator{contract: _SaleContract.contract, event: "Confirm", logs: logs, sub: sub}, nil
}

// WatchConfirm is a free log subscription operation binding the contract event 0x4bf5cf8e0c546146096c5ee70738e1612ac80ba22ebbe1e06cfb156c3a7c77a9.
//
// Solidity: e Confirm(tokenId indexed uint256, dci string, seller address, buyer address, price uint256)
func (_SaleContract *SaleContractFilterer) WatchConfirm(opts *bind.WatchOpts, sink chan<- *SaleContractConfirm, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SaleContract.contract.WatchLogs(opts, "Confirm", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleContractConfirm)
				if err := _SaleContract.contract.UnpackLog(event, "Confirm", log); err != nil {
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

// SaleContractCreateIterator is returned from FilterCreate and is used to iterate over the raw logs and unpacked data for Create events raised by the SaleContract contract.
type SaleContractCreateIterator struct {
	Event *SaleContractCreate // Event containing the contract specifics and raw log

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
func (it *SaleContractCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleContractCreate)
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
		it.Event = new(SaleContractCreate)
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
func (it *SaleContractCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleContractCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleContractCreate represents a Create event raised by the SaleContract contract.
type SaleContractCreate struct {
	TokenId *big.Int
	Dci     string
	Seller  common.Address
	Buyer   common.Address
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreate is a free log retrieval operation binding the contract event 0x0048209bd46e9c7af682cae0cfbd459137e1a1d6cfe0b90f9767f3ef1c6f6bf0.
//
// Solidity: e Create(tokenId indexed uint256, dci string, seller address, buyer address, price uint256)
func (_SaleContract *SaleContractFilterer) FilterCreate(opts *bind.FilterOpts, tokenId []*big.Int) (*SaleContractCreateIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SaleContract.contract.FilterLogs(opts, "Create", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SaleContractCreateIterator{contract: _SaleContract.contract, event: "Create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0x0048209bd46e9c7af682cae0cfbd459137e1a1d6cfe0b90f9767f3ef1c6f6bf0.
//
// Solidity: e Create(tokenId indexed uint256, dci string, seller address, buyer address, price uint256)
func (_SaleContract *SaleContractFilterer) WatchCreate(opts *bind.WatchOpts, sink chan<- *SaleContractCreate, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SaleContract.contract.WatchLogs(opts, "Create", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleContractCreate)
				if err := _SaleContract.contract.UnpackLog(event, "Create", log); err != nil {
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
