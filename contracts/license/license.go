// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package license

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
const ItemContractBin = `0x608060405234801561001057600080fd5b50611422806100206000396000f3fe608060405234801561001057600080fd5b50600436106100c6576000357c0100000000000000000000000000000000000000000000000000000000900480636b8ff5741161008e5780636b8ff5741461042a578063b7760c8f14610447578063c41a360a14610473578063d48e638a14610490578063db157339146104ad578063febe4909146104ca576100c6565b8063081812fc146100cb5780631e7663bc146101045780633bb3a24d146101bc5780634f0cd27b1461024e5780635d28560a14610274575b600080fd5b6100e8600480360360208110156100e157600080fd5b50356104f6565b60408051600160a060020a039092168252519081900360200190f35b6101aa6004803603602081101561011a57600080fd5b81019060208101813564010000000081111561013557600080fd5b82018360208201111561014757600080fd5b8035906020019184600183028401116401000000008311171561016957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610561945050505050565b60408051918252519081900360200190f35b6101d9600480360360208110156101d257600080fd5b50356105c9565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102135781810151838201526020016101fb565b50505050905090810190601f1680156102405780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101aa6004803603602081101561026457600080fd5b5035600160a060020a03166106bb565b6104286004803603606081101561028a57600080fd5b8101906020810181356401000000008111156102a557600080fd5b8201836020820111156102b757600080fd5b803590602001918460018302840111640100000000831117156102d957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561032c57600080fd5b82018360208201111561033e57600080fd5b8035906020019184600183028401116401000000008311171561036057600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156103b357600080fd5b8201836020820111156103c557600080fd5b803590602001918460018302840111640100000000831117156103e757600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061072e945050505050565b005b6101d96004803603602081101561044057600080fd5b5035610a80565b6104286004803603604081101561045d57600080fd5b5080359060200135600160a060020a0316610b4e565b6100e86004803603602081101561048957600080fd5b5035610db3565b6100e8600480360360208110156104a657600080fd5b5035610e1e565b6101d9600480360360208110156104c357600080fd5b5035610e8c565b610428600480360360408110156104e057600080fd5b5080359060200135600160a060020a0316610f43565b60006105018261114c565b1515610545576040805160e560020a62461bcd0281526020600482015260166024820152600080516020611390833981519152604482015290519081900360640190fd5b50600090815260056020526040902054600160a060020a031690565b60006003826040518082805190602001908083835b602083106105955780518252601f199092019160209182019101610576565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054949350505050565b60606105d48261114c565b1515610618576040805160e560020a62461bcd0281526020600482015260166024820152600080516020611390833981519152604482015290519081900360640190fd5b6000828152600460209081526040918290206002908101805484516001821615610100026000190190911692909204601f8101849004840283018401909452838252909290918301828280156106af5780601f10610684576101008083540402835291602001916106af565b820191906000526020600020905b81548152906001019060200180831161069257829003601f168201915b50505050509050919050565b6000600160a060020a03821615156107075760405160e560020a62461bcd02815260040180806020018281038252602681526020018061136a6026913960400191505060405180910390fd5b600160a060020a038216600090815260016020526040902061072890611169565b92915050565b61073f61073a84610561565b61114c565b15610794576040805160e560020a62461bcd02815260206004820152601360248201527f4974656d3a206578697374656e74206974656d00000000000000000000000000604482015290519081900360640190fd5b61079e6000611174565b60006107aa600061117d565b3360009081526001602052604090209091506107cc908263ffffffff61118116565b50600081815260026020908152604091829020805473ffffffffffffffffffffffffffffffffffffffff191633179055905185518392600392889290918291908401908083835b602083106108325780518252601f199092019160209182019101610813565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382018520959095556080840185528884528381018890528385018790523360608501526000868152600482529490942083518051949591946108a294508593509101906112ce565b5060208281015180516108bb92600185019201906112ce565b50604082015180516108d79160028401916020909101906112ce565b5060608201518160030160006101000a815481600160a060020a030219169083600160a060020a031602179055509050508033600160a060020a03167fefebf1e9cc6cf1906f62a4ed1545a6c54e027262b597e06b8e00abc4a3d8f50186868660405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561097c578181015183820152602001610964565b50505050905090810190601f1680156109a95780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b838110156109dc5781810151838201526020016109c4565b50505050905090810190601f168015610a095780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b83811015610a3c578181015183820152602001610a24565b50505050905090810190601f168015610a695780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a350505050565b6060610a8b8261114c565b1515610acf576040805160e560020a62461bcd0281526020600482015260166024820152600080516020611390833981519152604482015290519081900360640190fd5b600460008381526020019081526020016000206001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106af5780601f10610684576101008083540402835291602001916106af565b600160a060020a0381161515610b985760405160e560020a62461bcd0281526004018080602001828103825260228152602001806113d56022913960400191505060405180910390fd5b610ba18261114c565b1515610be5576040805160e560020a62461bcd0281526020600482015260166024820152600080516020611390833981519152604482015290519081900360640190fd5b600082815260026020526040902054600160a060020a0316610c06836104f6565b600160a060020a031633600160a060020a03161480610c2d575033600160a060020a038216145b1515610c6d5760405160e560020a62461bcd0281526004018080602001828103825260258152602001806113b06025913960400191505060405180910390fd5b600160a060020a0381166000908152600160205260409020610c95908463ffffffff61119416565b50600160a060020a0382166000908152600160205260409020610cbe908463ffffffff61118116565b5060008381526002602090815260408083208054600160a060020a0380881673ffffffffffffffffffffffffffffffffffffffff1992831681179093556005909452919093208054909116905584919083167fcd6e659e4c2e75c3bfe47fecaccf39aeb368116a0ee52afb532e07f6cba6c0d1610d3a84610e8c565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610d74578181015183820152602001610d5c565b50505050905090810190601f168015610da15780820380516001836020036101000a031916815260200191505b509250505060405180910390a4505050565b6000610dbe8261114c565b1515610e02576040805160e560020a62461bcd0281526020600482015260166024820152600080516020611390833981519152604482015290519081900360640190fd5b50600090815260026020526040902054600160a060020a031690565b6000610e298261114c565b1515610e6d576040805160e560020a62461bcd0281526020600482015260166024820152600080516020611390833981519152604482015290519081900360640190fd5b50600090815260046020526040902060030154600160a060020a031690565b6060610e978261114c565b1515610edb576040805160e560020a62461bcd0281526020600482015260166024820152600080516020611390833981519152604482015290519081900360640190fd5b60008281526004602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156106af5780601f10610684576101008083540402835291602001916106af565b610f4c8261114c565b1515610f90576040805160e560020a62461bcd0281526020600482015260166024820152600080516020611390833981519152604482015290519081900360640190fd5b600082815260026020526040902054600160a060020a0382811691161415611002576040805160e560020a62461bcd02815260206004820152601f60248201527f4974656d3a20617070726f76616c20746f2063757272656e74206f776e657200604482015290519081900360640190fd5b600082815260026020526040902054600160a060020a03163314611070576040805160e560020a62461bcd02815260206004820152601960248201527f4974656d3a2063616c6c6572206973206e6f74206f776e657200000000000000604482015290519081900360640190fd5b6000828152600560205260409020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383169081179091558290337ffa07013316165a7c424e6be19a166ba4784477c405b341b4bb535985391e7baa6110d484610e8c565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561110e5781810151838201526020016110f6565b50505050905090810190601f16801561113b5780820380516001836020036101000a031916815260200191505b509250505060405180910390a45050565b600090815260026020526040902054600160a060020a0316151590565b60006107288261117d565b80546001019055565b5490565b600061118d83836111a0565b9392505050565b600061118d83836111ec565b60006111ac83836112b6565b15156111e457508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610728565b506000610728565b600081815260018301602052604081205480156112ac578354600019808301919081019060009087908390811061121f57fe5b9060005260206000200154905080876000018481548110151561123e57fe5b600091825260208083209091019290925582815260018981019092526040902090840190558654879080151561127057fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050610728565b6000915050610728565b60009081526001919091016020526040902054151590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061130f57805160ff191683800117855561133c565b8280016001018555821561133c579182015b8281111561133c578251825591602001919060010190611321565b5061134892915061134c565b5090565b61136691905b808211156113485760008155600101611352565b9056fe4974656d3a20636f756e7420717565727920666f7220746865207a65726f20616464726573734974656d3a206e6f6e6578697374656e74206974656d000000000000000000004974656d3a2063616c6c6572206973206e6f74206f776e6572206f7220617070726f76616c4974656d3a207472616e7366657220746f20746865207a65726f2061646472657373a165627a7a7230582045f890e92f7a84ef04d8297bd5abbadfc4c8fe5c59cd17dd887898c8a0055c180029`

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

// LicenseContractABI is the input ABI used to generate the binding from.
const LicenseContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"hasValidLicense\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"},{\"name\":\"count\",\"type\":\"uint32\"}],\"name\":\"create\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getDeadline\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getDuration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"finish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getSeller\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"itemContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"dci\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"Create\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"dci\",\"type\":\"string\"}],\"name\":\"Finish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"recordId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"dci\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"begin\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"Buy\",\"type\":\"event\"}]"

// LicenseContractBin is the compiled bytecode used for deploying new contracts.
const LicenseContractBin = `0x608060405234801561001057600080fd5b506040516020806116798339810180604052602081101561003057600080fd5b505160008054600160a060020a039092166101000261010060a860020a031960ff1990931660011792909216919091179055611608806100716000396000f3fe608060405260043610610098576000357c010000000000000000000000000000000000000000000000000000000090048063d353a1cb1161006b578063d353a1cb14610185578063d6a9de51146101af578063d96a094a146101f5578063e75722301461021257610098565b80631c48fc7d1461009d578063315ce431146100db578063828622751461011f5780638bb126a71461015b575b600080fd5b3480156100a957600080fd5b506100c7600480360360208110156100c057600080fd5b503561023c565b604080519115158252519081900360200190f35b3480156100e757600080fd5b5061011d600480360360808110156100fe57600080fd5b508035906020810135906040810135906060013563ffffffff166102a3565b005b34801561012b57600080fd5b506101496004803603602081101561014257600080fd5b50356105ce565b60408051918252519081900360200190f35b34801561016757600080fd5b506101496004803603602081101561017e57600080fd5b5035610655565b34801561019157600080fd5b5061011d600480360360208110156101a857600080fd5b50356106ba565b3480156101bb57600080fd5b506101d9600480360360208110156101d257600080fd5b5035610b86565b60408051600160a060020a039092168252519081900360200190f35b61011d6004803603602081101561020b57600080fd5b5035610bf1565b34801561021e57600080fd5b506101496004803603602081101561023557600080fd5b503561120d565b600061024782611273565b151561028b576040805160e560020a62461bcd02815260206004820152601960248201526000805160206115bd833981519152604482015290519081900360640190fd5b61029d6102988333611290565b611315565b92915050565b6102ac84611273565b15610301576040805160e560020a62461bcd02815260206004820152601960248201527f4c6963656e73653a206578697374656e74206c6963656e736500000000000000604482015290519081900360640190fd5b60408051608081018252338152602080820186815282840186815263ffffffff8681166060860190815260008b81526001958690528781209651875473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039182161788559451958701959095559151600286015590516003909401805463ffffffff19169490911693909317909255805483517fb7760c8f0000000000000000000000000000000000000000000000000000000081526004810189905230602482015293516101009091049092169263b7760c8f926044808301939282900301818387803b1580156103ef57600080fd5b505af1158015610403573d6000803e3d6000fd5b505060008054604080517fdb157339000000000000000000000000000000000000000000000000000000008152600481018a905290518995507f52ba39039f10a9e70334986a650e823646b5562a9726319428dc963670996b3e9450610100909204600160a060020a03169263db15733992602480840193829003018186803b15801561048f57600080fd5b505afa1580156104a3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156104cc57600080fd5b8101908080516401000000008111156104e457600080fd5b820160208101848111156104f757600080fd5b815164010000000081118282018710171561051157600080fd5b505092919050505033868686604051808060200186600160a060020a0316600160a060020a031681526020018581526020018481526020018363ffffffff168152602001828103825287818151815260200191508051906020019080838360005b8381101561058a578181015183820152602001610572565b50505050905090810190601f1680156105b75780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a250505050565b60006105d982611273565b151561061d576040805160e560020a62461bcd02815260206004820152601960248201526000805160206115bd833981519152604482015290519081900360640190fd5b50336000908152600560209081526040808320938352928152828220546003825283832063ffffffff90911683529052206004015490565b600061066082611273565b15156106a4576040805160e560020a62461bcd02815260206004820152601960248201526000805160206115bd833981519152604482015290519081900360640190fd5b5060009081526001602052604090206002015490565b6106c381611273565b1515610719576040805160e560020a62461bcd02815260206004820152601c60248201527f4c6963656e73653a206e6f6e6578697374656e74206c6963656e736500000000604482015290519081900360640190fd5b6107216114d0565b61072a8261131e565b8051909150600160a060020a0316331461078e576040805160e560020a62461bcd02815260206004820152601e60248201527f4c6963656e7365203a2063616c6c6572206973206e6f742073656c6c65720000604482015290519081900360640190fd5b60008281526004602090815260409182902080548351818402810184019094528084526060939283018282801561081057602002820191906000526020600020906000905b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116107d35790505b50939450600093505050505b8151811015610908576000828281518110151561083557fe5b602090810290910181015160008781526003808452604080832063ffffffff85168452855291829020825160a0810184528154600160a060020a03908116825260018301541695810195909552600281015492850192909252810154606084015260040154608083015291506108aa90611315565b156108ff576040805160e560020a62461bcd02815260206004820152601e60248201527f4c6963656e73653a206578697374656e742076616c6964207265636f72640000604482015290519081900360640190fd5b5060010161081c565b506000838152600160208181526040808420805473ffffffffffffffffffffffffffffffffffffffff19168155928301849055600283018490556003909201805463ffffffff19169055600490528120610961916114f7565b600083815260026020526040808220805463ffffffff19169055815481517fb7760c8f000000000000000000000000000000000000000000000000000000008152600481018790523360248201529151610100909104600160a060020a03169263b7760c8f926044808201939182900301818387803b1580156109e357600080fd5b505af11580156109f7573d6000803e3d6000fd5b505060008054604080517fdb1573390000000000000000000000000000000000000000000000000000000081526004810189905290518895507f762a3b15ea494b75373369b598c062b25af5e229f043d2d1fda080ae0f5130e09450610100909204600160a060020a03169263db15733992602480840193829003018186803b158015610a8357600080fd5b505afa158015610a97573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610ac057600080fd5b810190808051640100000000811115610ad857600080fd5b82016020810184811115610aeb57600080fd5b8151640100000000811182820187101715610b0557600080fd5b50506040805160208082528351818301528351939650909450849350908301919085019080838360005b83811015610b47578181015183820152602001610b2f565b50505050905090810190601f168015610b745780820380516001836020036101000a031916815260200191505b509250505060405180910390a2505050565b6000610b9182611273565b1515610bd5576040805160e560020a62461bcd02815260206004820152601960248201526000805160206115bd833981519152604482015290519081900360640190fd5b50600090815260016020526040902054600160a060020a031690565b60005460ff161515610c4d576040805160e560020a62461bcd02815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6000805460ff19169055610c6081611273565b1515610cb6576040805160e560020a62461bcd02815260206004820152601c60248201527f4c6963656e73653a206e6f6e6578697374656e74206c6963656e736500000000604482015290519081900360640190fd5b610cbe6114d0565b610cc78261131e565b8051909150600160a060020a0316331415610d2c576040805160e560020a62461bcd02815260206004820152601960248201527f4c6963656e73653a2063616c6c65722069732073656c6c657200000000000000604482015290519081900360640190fd5b6020810151341015610d88576040805160e560020a62461bcd02815260206004820152601960248201527f4c6963656e73653a206e6f7420656e6f7567682076616c756500000000000000604482015290519081900360640190fd5b6000816060015163ffffffff161115610e0d57600082815260046020526040902054606082015163ffffffff168110610e0b576040805160e560020a62461bcd02815260206004820152601560248201527f4c6963656e73653a206e6f2072656d61696e6465720000000000000000000000604482015290519081900360640190fd5b505b600082815260026020526040808220805463ffffffff198116600163ffffffff92831690810183169190911790925591840151909291610e4f91429161137916565b9050610e5961151f565b60a0604051908101604052808560000151600160a060020a0316815260200133600160a060020a0316815260200185602001518152602001428152602001838152509050806003600087815260200190815260200160002060008563ffffffff1663ffffffff16815260200190815260200160002060008201518160000160006101000a815481600160a060020a030219169083600160a060020a0316021790555060208201518160010160006101000a815481600160a060020a030219169083600160a060020a031602179055506040820151816002015560608201518160030155608082015181600401559050506004600086815260200190815260200160002083908060018154018082558091505090600182039060005260206000209060089182820401919006600402909192909190916101000a81548163ffffffff021916908363ffffffff16021790555050826005600033600160a060020a0316600160a060020a03168152602001908152602001600020600087815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff16021790555061102484602001518560000151600160a060020a03166113dd90919063ffffffff16565b8263ffffffff16857f090dcca045911a4fb20577b7b3dd06f88d7fe807cf32e62985e2c1312bdf674c600060019054906101000a9004600160a060020a0316600160a060020a031663db157339896040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082815260200191505060006040518083038186803b1580156110c257600080fd5b505afa1580156110d6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156110ff57600080fd5b81019080805164010000000081111561111757600080fd5b8201602081018481111561112a57600080fd5b815164010000000081118282018710171561114457600080fd5b50509291905050503388602001514288604051808060200186600160a060020a0316600160a060020a03168152602001858152602001848152602001838152602001828103825287818151815260200191508051906020019080838360005b838110156111bb5781810151838201526020016111a3565b50505050905090810190601f1680156111e85780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a350506000805460ff19166001179055505050565b600061121882611273565b151561125c576040805160e560020a62461bcd02815260206004820152601960248201526000805160206115bd833981519152604482015290519081900360640190fd5b506000908152600160208190526040909120015490565b600090815260016020526040902054600160a060020a0316151590565b61129861151f565b50600160a060020a038082166000908152600560209081526040808320868452825280832054600380845282852063ffffffff909216855290835292819020815160a08101835281548616815260018201549095169285019290925260028201549084015290810154606083015260040154608082015292915050565b60800151421090565b6113266114d0565b5060009081526001602081815260409283902083516080810185528154600160a060020a03168152928101549183019190915260028101549282019290925260039091015463ffffffff16606082015290565b6000828201838110156113d6576040805160e560020a62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b3031811115611436576040805160e560020a62461bcd02815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e6365000000604482015290519081900360640190fd5b604051600090600160a060020a0384169083908381818185875af1925050503d8060008114611481576040519150601f19603f3d011682016040523d82523d6000602084013e611486565b606091505b505090508015156114cb5760405160e560020a62461bcd02815260040180806020018281038252603a815260200180611583603a913960400191505060405180910390fd5b505050565b60408051608081018252600080825260208201819052918101829052606081019190915290565b50805460008255600701600890049060005260206000209081019061151c9190611561565b50565b60a0604051908101604052806000600160a060020a031681526020016000600160a060020a031681526020016000815260200160008152602001600081525090565b61157f91905b8082111561157b5760008155600101611567565b5090565b9056fe416464726573733a20756e61626c6520746f2073656e642076616c75652c20726563697069656e74206d617920686176652072657665727465644c6963656e73653a206e6f6e6578697374656e74206974656d00000000000000a165627a7a723058205de072540b60d50edf4d671d17d2744d50dc791b8e4e285c2c049379445f5df90029`

// DeployLicenseContract deploys a new Ethereum contract, binding an instance of LicenseContract to it.
func DeployLicenseContract(auth *bind.TransactOpts, backend bind.ContractBackend, itemContract common.Address) (common.Address, *types.Transaction, *LicenseContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LicenseContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LicenseContractBin), backend, itemContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LicenseContract{LicenseContractCaller: LicenseContractCaller{contract: contract}, LicenseContractTransactor: LicenseContractTransactor{contract: contract}, LicenseContractFilterer: LicenseContractFilterer{contract: contract}}, nil
}

// LicenseContract is an auto generated Go binding around an Ethereum contract.
type LicenseContract struct {
	LicenseContractCaller     // Read-only binding to the contract
	LicenseContractTransactor // Write-only binding to the contract
	LicenseContractFilterer   // Log filterer for contract events
}

// LicenseContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type LicenseContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LicenseContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LicenseContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LicenseContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LicenseContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LicenseContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LicenseContractSession struct {
	Contract     *LicenseContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LicenseContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LicenseContractCallerSession struct {
	Contract *LicenseContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// LicenseContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LicenseContractTransactorSession struct {
	Contract     *LicenseContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// LicenseContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type LicenseContractRaw struct {
	Contract *LicenseContract // Generic contract binding to access the raw methods on
}

// LicenseContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LicenseContractCallerRaw struct {
	Contract *LicenseContractCaller // Generic read-only contract binding to access the raw methods on
}

// LicenseContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LicenseContractTransactorRaw struct {
	Contract *LicenseContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLicenseContract creates a new instance of LicenseContract, bound to a specific deployed contract.
func NewLicenseContract(address common.Address, backend bind.ContractBackend) (*LicenseContract, error) {
	contract, err := bindLicenseContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LicenseContract{LicenseContractCaller: LicenseContractCaller{contract: contract}, LicenseContractTransactor: LicenseContractTransactor{contract: contract}, LicenseContractFilterer: LicenseContractFilterer{contract: contract}}, nil
}

// NewLicenseContractCaller creates a new read-only instance of LicenseContract, bound to a specific deployed contract.
func NewLicenseContractCaller(address common.Address, caller bind.ContractCaller) (*LicenseContractCaller, error) {
	contract, err := bindLicenseContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LicenseContractCaller{contract: contract}, nil
}

// NewLicenseContractTransactor creates a new write-only instance of LicenseContract, bound to a specific deployed contract.
func NewLicenseContractTransactor(address common.Address, transactor bind.ContractTransactor) (*LicenseContractTransactor, error) {
	contract, err := bindLicenseContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LicenseContractTransactor{contract: contract}, nil
}

// NewLicenseContractFilterer creates a new log filterer instance of LicenseContract, bound to a specific deployed contract.
func NewLicenseContractFilterer(address common.Address, filterer bind.ContractFilterer) (*LicenseContractFilterer, error) {
	contract, err := bindLicenseContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LicenseContractFilterer{contract: contract}, nil
}

// bindLicenseContract binds a generic wrapper to an already deployed contract.
func bindLicenseContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LicenseContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LicenseContract *LicenseContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LicenseContract.Contract.LicenseContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LicenseContract *LicenseContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LicenseContract.Contract.LicenseContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LicenseContract *LicenseContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LicenseContract.Contract.LicenseContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LicenseContract *LicenseContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LicenseContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LicenseContract *LicenseContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LicenseContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LicenseContract *LicenseContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LicenseContract.Contract.contract.Transact(opts, method, params...)
}

// GetDeadline is a free data retrieval call binding the contract method 0x82862275.
//
// Solidity: function getDeadline(tokenId uint256) constant returns(uint256)
func (_LicenseContract *LicenseContractCaller) GetDeadline(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LicenseContract.contract.Call(opts, out, "getDeadline", tokenId)
	return *ret0, err
}

// GetDeadline is a free data retrieval call binding the contract method 0x82862275.
//
// Solidity: function getDeadline(tokenId uint256) constant returns(uint256)
func (_LicenseContract *LicenseContractSession) GetDeadline(tokenId *big.Int) (*big.Int, error) {
	return _LicenseContract.Contract.GetDeadline(&_LicenseContract.CallOpts, tokenId)
}

// GetDeadline is a free data retrieval call binding the contract method 0x82862275.
//
// Solidity: function getDeadline(tokenId uint256) constant returns(uint256)
func (_LicenseContract *LicenseContractCallerSession) GetDeadline(tokenId *big.Int) (*big.Int, error) {
	return _LicenseContract.Contract.GetDeadline(&_LicenseContract.CallOpts, tokenId)
}

// GetDuration is a free data retrieval call binding the contract method 0x8bb126a7.
//
// Solidity: function getDuration(tokenId uint256) constant returns(uint256)
func (_LicenseContract *LicenseContractCaller) GetDuration(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LicenseContract.contract.Call(opts, out, "getDuration", tokenId)
	return *ret0, err
}

// GetDuration is a free data retrieval call binding the contract method 0x8bb126a7.
//
// Solidity: function getDuration(tokenId uint256) constant returns(uint256)
func (_LicenseContract *LicenseContractSession) GetDuration(tokenId *big.Int) (*big.Int, error) {
	return _LicenseContract.Contract.GetDuration(&_LicenseContract.CallOpts, tokenId)
}

// GetDuration is a free data retrieval call binding the contract method 0x8bb126a7.
//
// Solidity: function getDuration(tokenId uint256) constant returns(uint256)
func (_LicenseContract *LicenseContractCallerSession) GetDuration(tokenId *big.Int) (*big.Int, error) {
	return _LicenseContract.Contract.GetDuration(&_LicenseContract.CallOpts, tokenId)
}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(tokenId uint256) constant returns(uint256)
func (_LicenseContract *LicenseContractCaller) GetPrice(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LicenseContract.contract.Call(opts, out, "getPrice", tokenId)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(tokenId uint256) constant returns(uint256)
func (_LicenseContract *LicenseContractSession) GetPrice(tokenId *big.Int) (*big.Int, error) {
	return _LicenseContract.Contract.GetPrice(&_LicenseContract.CallOpts, tokenId)
}

// GetPrice is a free data retrieval call binding the contract method 0xe7572230.
//
// Solidity: function getPrice(tokenId uint256) constant returns(uint256)
func (_LicenseContract *LicenseContractCallerSession) GetPrice(tokenId *big.Int) (*big.Int, error) {
	return _LicenseContract.Contract.GetPrice(&_LicenseContract.CallOpts, tokenId)
}

// GetSeller is a free data retrieval call binding the contract method 0xd6a9de51.
//
// Solidity: function getSeller(tokenId uint256) constant returns(address)
func (_LicenseContract *LicenseContractCaller) GetSeller(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LicenseContract.contract.Call(opts, out, "getSeller", tokenId)
	return *ret0, err
}

// GetSeller is a free data retrieval call binding the contract method 0xd6a9de51.
//
// Solidity: function getSeller(tokenId uint256) constant returns(address)
func (_LicenseContract *LicenseContractSession) GetSeller(tokenId *big.Int) (common.Address, error) {
	return _LicenseContract.Contract.GetSeller(&_LicenseContract.CallOpts, tokenId)
}

// GetSeller is a free data retrieval call binding the contract method 0xd6a9de51.
//
// Solidity: function getSeller(tokenId uint256) constant returns(address)
func (_LicenseContract *LicenseContractCallerSession) GetSeller(tokenId *big.Int) (common.Address, error) {
	return _LicenseContract.Contract.GetSeller(&_LicenseContract.CallOpts, tokenId)
}

// HasValidLicense is a free data retrieval call binding the contract method 0x1c48fc7d.
//
// Solidity: function hasValidLicense(tokenId uint256) constant returns(bool)
func (_LicenseContract *LicenseContractCaller) HasValidLicense(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LicenseContract.contract.Call(opts, out, "hasValidLicense", tokenId)
	return *ret0, err
}

// HasValidLicense is a free data retrieval call binding the contract method 0x1c48fc7d.
//
// Solidity: function hasValidLicense(tokenId uint256) constant returns(bool)
func (_LicenseContract *LicenseContractSession) HasValidLicense(tokenId *big.Int) (bool, error) {
	return _LicenseContract.Contract.HasValidLicense(&_LicenseContract.CallOpts, tokenId)
}

// HasValidLicense is a free data retrieval call binding the contract method 0x1c48fc7d.
//
// Solidity: function hasValidLicense(tokenId uint256) constant returns(bool)
func (_LicenseContract *LicenseContractCallerSession) HasValidLicense(tokenId *big.Int) (bool, error) {
	return _LicenseContract.Contract.HasValidLicense(&_LicenseContract.CallOpts, tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(tokenId uint256) returns()
func (_LicenseContract *LicenseContractTransactor) Buy(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _LicenseContract.contract.Transact(opts, "buy", tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(tokenId uint256) returns()
func (_LicenseContract *LicenseContractSession) Buy(tokenId *big.Int) (*types.Transaction, error) {
	return _LicenseContract.Contract.Buy(&_LicenseContract.TransactOpts, tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(tokenId uint256) returns()
func (_LicenseContract *LicenseContractTransactorSession) Buy(tokenId *big.Int) (*types.Transaction, error) {
	return _LicenseContract.Contract.Buy(&_LicenseContract.TransactOpts, tokenId)
}

// Create is a paid mutator transaction binding the contract method 0x315ce431.
//
// Solidity: function create(tokenId uint256, price uint256, duration uint256, count uint32) returns()
func (_LicenseContract *LicenseContractTransactor) Create(opts *bind.TransactOpts, tokenId *big.Int, price *big.Int, duration *big.Int, count uint32) (*types.Transaction, error) {
	return _LicenseContract.contract.Transact(opts, "create", tokenId, price, duration, count)
}

// Create is a paid mutator transaction binding the contract method 0x315ce431.
//
// Solidity: function create(tokenId uint256, price uint256, duration uint256, count uint32) returns()
func (_LicenseContract *LicenseContractSession) Create(tokenId *big.Int, price *big.Int, duration *big.Int, count uint32) (*types.Transaction, error) {
	return _LicenseContract.Contract.Create(&_LicenseContract.TransactOpts, tokenId, price, duration, count)
}

// Create is a paid mutator transaction binding the contract method 0x315ce431.
//
// Solidity: function create(tokenId uint256, price uint256, duration uint256, count uint32) returns()
func (_LicenseContract *LicenseContractTransactorSession) Create(tokenId *big.Int, price *big.Int, duration *big.Int, count uint32) (*types.Transaction, error) {
	return _LicenseContract.Contract.Create(&_LicenseContract.TransactOpts, tokenId, price, duration, count)
}

// Finish is a paid mutator transaction binding the contract method 0xd353a1cb.
//
// Solidity: function finish(tokenId uint256) returns()
func (_LicenseContract *LicenseContractTransactor) Finish(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _LicenseContract.contract.Transact(opts, "finish", tokenId)
}

// Finish is a paid mutator transaction binding the contract method 0xd353a1cb.
//
// Solidity: function finish(tokenId uint256) returns()
func (_LicenseContract *LicenseContractSession) Finish(tokenId *big.Int) (*types.Transaction, error) {
	return _LicenseContract.Contract.Finish(&_LicenseContract.TransactOpts, tokenId)
}

// Finish is a paid mutator transaction binding the contract method 0xd353a1cb.
//
// Solidity: function finish(tokenId uint256) returns()
func (_LicenseContract *LicenseContractTransactorSession) Finish(tokenId *big.Int) (*types.Transaction, error) {
	return _LicenseContract.Contract.Finish(&_LicenseContract.TransactOpts, tokenId)
}

// LicenseContractBuyIterator is returned from FilterBuy and is used to iterate over the raw logs and unpacked data for Buy events raised by the LicenseContract contract.
type LicenseContractBuyIterator struct {
	Event *LicenseContractBuy // Event containing the contract specifics and raw log

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
func (it *LicenseContractBuyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenseContractBuy)
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
		it.Event = new(LicenseContractBuy)
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
func (it *LicenseContractBuyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenseContractBuyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenseContractBuy represents a Buy event raised by the LicenseContract contract.
type LicenseContractBuy struct {
	TokenId  *big.Int
	RecordId *big.Int
	Dci      string
	Buyer    common.Address
	Price    *big.Int
	Begin    *big.Int
	Deadline *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBuy is a free log retrieval operation binding the contract event 0x090dcca045911a4fb20577b7b3dd06f88d7fe807cf32e62985e2c1312bdf674c.
//
// Solidity: e Buy(tokenId indexed uint256, recordId indexed uint256, dci string, buyer address, price uint256, begin uint256, deadline uint256)
func (_LicenseContract *LicenseContractFilterer) FilterBuy(opts *bind.FilterOpts, tokenId []*big.Int, recordId []*big.Int) (*LicenseContractBuyIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var recordIdRule []interface{}
	for _, recordIdItem := range recordId {
		recordIdRule = append(recordIdRule, recordIdItem)
	}

	logs, sub, err := _LicenseContract.contract.FilterLogs(opts, "Buy", tokenIdRule, recordIdRule)
	if err != nil {
		return nil, err
	}
	return &LicenseContractBuyIterator{contract: _LicenseContract.contract, event: "Buy", logs: logs, sub: sub}, nil
}

// WatchBuy is a free log subscription operation binding the contract event 0x090dcca045911a4fb20577b7b3dd06f88d7fe807cf32e62985e2c1312bdf674c.
//
// Solidity: e Buy(tokenId indexed uint256, recordId indexed uint256, dci string, buyer address, price uint256, begin uint256, deadline uint256)
func (_LicenseContract *LicenseContractFilterer) WatchBuy(opts *bind.WatchOpts, sink chan<- *LicenseContractBuy, tokenId []*big.Int, recordId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var recordIdRule []interface{}
	for _, recordIdItem := range recordId {
		recordIdRule = append(recordIdRule, recordIdItem)
	}

	logs, sub, err := _LicenseContract.contract.WatchLogs(opts, "Buy", tokenIdRule, recordIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenseContractBuy)
				if err := _LicenseContract.contract.UnpackLog(event, "Buy", log); err != nil {
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

// LicenseContractCreateIterator is returned from FilterCreate and is used to iterate over the raw logs and unpacked data for Create events raised by the LicenseContract contract.
type LicenseContractCreateIterator struct {
	Event *LicenseContractCreate // Event containing the contract specifics and raw log

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
func (it *LicenseContractCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenseContractCreate)
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
		it.Event = new(LicenseContractCreate)
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
func (it *LicenseContractCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenseContractCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenseContractCreate represents a Create event raised by the LicenseContract contract.
type LicenseContractCreate struct {
	TokenId  *big.Int
	Dci      string
	Seller   common.Address
	Price    *big.Int
	Duration *big.Int
	Count    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCreate is a free log retrieval operation binding the contract event 0x52ba39039f10a9e70334986a650e823646b5562a9726319428dc963670996b3e.
//
// Solidity: e Create(tokenId indexed uint256, dci string, seller address, price uint256, duration uint256, count uint256)
func (_LicenseContract *LicenseContractFilterer) FilterCreate(opts *bind.FilterOpts, tokenId []*big.Int) (*LicenseContractCreateIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LicenseContract.contract.FilterLogs(opts, "Create", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LicenseContractCreateIterator{contract: _LicenseContract.contract, event: "Create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0x52ba39039f10a9e70334986a650e823646b5562a9726319428dc963670996b3e.
//
// Solidity: e Create(tokenId indexed uint256, dci string, seller address, price uint256, duration uint256, count uint256)
func (_LicenseContract *LicenseContractFilterer) WatchCreate(opts *bind.WatchOpts, sink chan<- *LicenseContractCreate, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LicenseContract.contract.WatchLogs(opts, "Create", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenseContractCreate)
				if err := _LicenseContract.contract.UnpackLog(event, "Create", log); err != nil {
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

// LicenseContractFinishIterator is returned from FilterFinish and is used to iterate over the raw logs and unpacked data for Finish events raised by the LicenseContract contract.
type LicenseContractFinishIterator struct {
	Event *LicenseContractFinish // Event containing the contract specifics and raw log

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
func (it *LicenseContractFinishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenseContractFinish)
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
		it.Event = new(LicenseContractFinish)
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
func (it *LicenseContractFinishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenseContractFinishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenseContractFinish represents a Finish event raised by the LicenseContract contract.
type LicenseContractFinish struct {
	TokenId *big.Int
	Dci     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinish is a free log retrieval operation binding the contract event 0x762a3b15ea494b75373369b598c062b25af5e229f043d2d1fda080ae0f5130e0.
//
// Solidity: e Finish(tokenId indexed uint256, dci string)
func (_LicenseContract *LicenseContractFilterer) FilterFinish(opts *bind.FilterOpts, tokenId []*big.Int) (*LicenseContractFinishIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LicenseContract.contract.FilterLogs(opts, "Finish", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LicenseContractFinishIterator{contract: _LicenseContract.contract, event: "Finish", logs: logs, sub: sub}, nil
}

// WatchFinish is a free log subscription operation binding the contract event 0x762a3b15ea494b75373369b598c062b25af5e229f043d2d1fda080ae0f5130e0.
//
// Solidity: e Finish(tokenId indexed uint256, dci string)
func (_LicenseContract *LicenseContractFilterer) WatchFinish(opts *bind.WatchOpts, sink chan<- *LicenseContractFinish, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LicenseContract.contract.WatchLogs(opts, "Finish", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenseContractFinish)
				if err := _LicenseContract.contract.UnpackLog(event, "Finish", log); err != nil {
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
