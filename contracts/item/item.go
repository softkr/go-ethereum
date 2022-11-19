// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package item

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
const ItemContractBin = `0x608060405234801561001057600080fd5b5061153c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100c6576000357c0100000000000000000000000000000000000000000000000000000000900480636b8ff5741161008e5780636b8ff5741461042a578063b7760c8f14610447578063c41a360a14610473578063d48e638a14610490578063db157339146104ad578063febe4909146104ca576100c6565b8063081812fc146100cb5780631e7663bc146101045780633bb3a24d146101bc5780634f0cd27b1461024e5780635d28560a14610274575b600080fd5b6100e8600480360360208110156100e157600080fd5b50356104f6565b60408051600160a060020a039092168252519081900360200190f35b6101aa6004803603602081101561011a57600080fd5b81019060208101813564010000000081111561013557600080fd5b82018360208201111561014757600080fd5b8035906020019184600183028401116401000000008311171561016957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610561945050505050565b60408051918252519081900360200190f35b6101d9600480360360208110156101d257600080fd5b50356105c9565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102135781810151838201526020016101fb565b50505050905090810190601f1680156102405780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101aa6004803603602081101561026457600080fd5b5035600160a060020a03166106bb565b6104286004803603606081101561028a57600080fd5b8101906020810181356401000000008111156102a557600080fd5b8201836020820111156102b757600080fd5b803590602001918460018302840111640100000000831117156102d957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561032c57600080fd5b82018360208201111561033e57600080fd5b8035906020019184600183028401116401000000008311171561036057600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156103b357600080fd5b8201836020820111156103c557600080fd5b803590602001918460018302840111640100000000831117156103e757600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061072e945050505050565b005b6101d96004803603602081101561044057600080fd5b5035610a61565b6104286004803603604081101561045d57600080fd5b5080359060200135600160a060020a0316610b2f565b6100e86004803603602081101561048957600080fd5b5035610db1565b6100e8600480360360208110156104a657600080fd5b5035610e48565b6101d9600480360360208110156104c357600080fd5b5035610eb6565b610428600480360360408110156104e057600080fd5b5080359060200135600160a060020a0316610f6d565b6000610501826110d0565b1515610545576040805160e560020a62461bcd02815260206004820152601660248201526000805160206114aa833981519152604482015290519081900360640190fd5b50600090815260066020526040902054600160a060020a031690565b60006004826040518082805190602001908083835b602083106105955780518252601f199092019160209182019101610576565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054949350505050565b60606105d4826110d0565b1515610618576040805160e560020a62461bcd02815260206004820152601660248201526000805160206114aa833981519152604482015290519081900360640190fd5b6000828152600560209081526040918290206002908101805484516001821615610100026000190190911692909204601f8101849004840283018401909452838252909290918301828280156106af5780601f10610684576101008083540402835291602001916106af565b820191906000526020600020905b81548152906001019060200180831161069257829003601f168201915b50505050509050919050565b6000600160a060020a03821615156107075760405160e560020a62461bcd0281526004018080602001828103825260268152602001806114846026913960400191505060405180910390fd5b600160a060020a0382166000908152600160205260409020610728906110e3565b92915050565b61073f61073a84610561565b6110d0565b15610794576040805160e560020a62461bcd02815260206004820152601360248201527f4974656d3a206578697374656e74206974656d00000000000000000000000000604482015290519081900360640190fd5b61079e60006110ee565b60006107aa60006110f7565b3360009081526001602052604090209091506107cc908263ffffffff6110fb16565b506107df6002823363ffffffff61110e16565b50806004856040518082805190602001908083835b602083106108135780518252601f1990920191602091820191016107f4565b51815160209384036101000a60001901801990921691161790529201948552506040805194859003820185209590955560808401855288845283810188905283850187905233606085015260008681526005825294909420835180519495919461088394508593509101906113e8565b50602082810151805161089c92600185019201906113e8565b50604082015180516108b89160028401916020909101906113e8565b5060608201518160030160006101000a815481600160a060020a030219169083600160a060020a031602179055509050508033600160a060020a03167fefebf1e9cc6cf1906f62a4ed1545a6c54e027262b597e06b8e00abc4a3d8f50186868660405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561095d578181015183820152602001610945565b50505050905090810190601f16801561098a5780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b838110156109bd5781810151838201526020016109a5565b50505050905090810190601f1680156109ea5780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b83811015610a1d578181015183820152602001610a05565b50505050905090810190601f168015610a4a5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a350505050565b6060610a6c826110d0565b1515610ab0576040805160e560020a62461bcd02815260206004820152601660248201526000805160206114aa833981519152604482015290519081900360640190fd5b600560008381526020019081526020016000206001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106af5780601f10610684576101008083540402835291602001916106af565b600160a060020a0381161515610b795760405160e560020a62461bcd0281526004018080602001828103825260228152602001806114ef6022913960400191505060405180910390fd5b610b82826110d0565b1515610bc6576040805160e560020a62461bcd02815260206004820152601660248201526000805160206114aa833981519152604482015290519081900360640190fd5b60408051808201909152601681526000805160206114aa8339815191526020820152600090610bff90600290859063ffffffff61112c16565b9050610c0a836104f6565b600160a060020a031633600160a060020a03161480610c31575033600160a060020a038216145b1515610c715760405160e560020a62461bcd0281526004018080602001828103825260258152602001806114ca6025913960400191505060405180910390fd5b600160a060020a0381166000908152600160205260409020610c99908463ffffffff61113916565b50600160a060020a0382166000908152600160205260409020610cc2908463ffffffff6110fb16565b50610cd56002848463ffffffff61110e16565b506000838152600660205260409020805473ffffffffffffffffffffffffffffffffffffffff1916905582600160a060020a038381169083167fcd6e659e4c2e75c3bfe47fecaccf39aeb368116a0ee52afb532e07f6cba6c0d1610d3884610eb6565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610d72578181015183820152602001610d5a565b50505050905090810190601f168015610d9f5780820380516001836020036101000a031916815260200191505b509250505060405180910390a4505050565b6000610dbc826110d0565b1515610e00576040805160e560020a62461bcd02815260206004820152601660248201526000805160206114aa833981519152604482015290519081900360640190fd5b60408051808201909152601781527f4974656d3a206e6f6e6578697374656e7420746f6b656e000000000000000000602082015261072890600290849063ffffffff61112c16565b6000610e53826110d0565b1515610e97576040805160e560020a62461bcd02815260206004820152601660248201526000805160206114aa833981519152604482015290519081900360640190fd5b50600090815260056020526040902060030154600160a060020a031690565b6060610ec1826110d0565b1515610f05576040805160e560020a62461bcd02815260206004820152601660248201526000805160206114aa833981519152604482015290519081900360640190fd5b60008281526005602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156106af5780601f10610684576101008083540402835291602001916106af565b60408051808201909152601681526000805160206114aa8339815191526020820152600090610fa690600290859063ffffffff61112c16565b9050600160a060020a03828116908216141561100c576040805160e560020a62461bcd02815260206004820152601f60248201527f4974656d3a20617070726f76616c20746f2063757272656e74206f776e657200604482015290519081900360640190fd5b33600160a060020a0382161461106c576040805160e560020a62461bcd02815260206004820152601960248201527f4974656d3a2063616c6c6572206973206e6f74206f776e657200000000000000604482015290519081900360640190fd5b6000838152600660205260409020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384169081179091558390337ffa07013316165a7c424e6be19a166ba4784477c405b341b4bb535985391e7baa610d3884610eb6565b600061072860028363ffffffff61114516565b6000610728826110f7565b80546001019055565b5490565b60006111078383611151565b9392505050565b60006111248484600160a060020a03851661119d565b949350505050565b6000611124848484611237565b60006111078383611306565b600061110783836113d0565b600061115d83836113d0565b151561119557508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610728565b506000610728565b6000828152600184016020526040812054801515611204575050604080518082018252838152602080820184815286546001818101895560008981528481209551600290930290950191825591519082015586548684528188019092529290912055611107565b845483908690600019840190811061121857fe5b9060005260206000209060020201600101819055506000915050611107565b6000828152600184016020526040812054828115156112d75760405160e560020a62461bcd0281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561129c578181015183820152602001611284565b50505050905090810190601f1680156112c95780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b508454859060001983019081106112ea57fe5b9060005260206000209060020201600101549150509392505050565b600081815260018301602052604081205480156113c6578354600019808301919081019060009087908390811061133957fe5b9060005260206000200154905080876000018481548110151561135857fe5b600091825260208083209091019290925582815260018981019092526040902090840190558654879080151561138a57fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050610728565b6000915050610728565b60009081526001919091016020526040902054151590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061142957805160ff1916838001178555611456565b82800160010185558215611456579182015b8281111561145657825182559160200191906001019061143b565b50611462929150611466565b5090565b61148091905b80821115611462576000815560010161146c565b9056fe4974656d3a20636f756e7420717565727920666f7220746865207a65726f20616464726573734974656d3a206e6f6e6578697374656e74206974656d000000000000000000004974656d3a2063616c6c6572206973206e6f74206f776e6572206f7220617070726f76616c4974656d3a207472616e7366657220746f20746865207a65726f2061646472657373a165627a7a72305820ab290ee2be4486a352cd492c2f868d5c49537f1db67e8086e30a2a6dddeb8e0e0029`

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
