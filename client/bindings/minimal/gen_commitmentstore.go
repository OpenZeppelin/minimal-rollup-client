// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package minimal

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ICommitmentStoreMetaData contains all meta data concerning the ICommitmentStore contract.
var ICommitmentStoreMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"commitmentAt\",\"inputs\":[{\"name\":\"source\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"storeCommitment\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CommitmentStored\",\"inputs\":[{\"name\":\"source\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"height\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"commitment\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false}]",
}

// ICommitmentStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use ICommitmentStoreMetaData.ABI instead.
var ICommitmentStoreABI = ICommitmentStoreMetaData.ABI

// ICommitmentStore is an auto generated Go binding around an Ethereum contract.
type ICommitmentStore struct {
	ICommitmentStoreCaller     // Read-only binding to the contract
	ICommitmentStoreTransactor // Write-only binding to the contract
	ICommitmentStoreFilterer   // Log filterer for contract events
}

// ICommitmentStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICommitmentStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICommitmentStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICommitmentStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICommitmentStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICommitmentStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICommitmentStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICommitmentStoreSession struct {
	Contract     *ICommitmentStore // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICommitmentStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICommitmentStoreCallerSession struct {
	Contract *ICommitmentStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ICommitmentStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICommitmentStoreTransactorSession struct {
	Contract     *ICommitmentStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ICommitmentStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICommitmentStoreRaw struct {
	Contract *ICommitmentStore // Generic contract binding to access the raw methods on
}

// ICommitmentStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICommitmentStoreCallerRaw struct {
	Contract *ICommitmentStoreCaller // Generic read-only contract binding to access the raw methods on
}

// ICommitmentStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICommitmentStoreTransactorRaw struct {
	Contract *ICommitmentStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICommitmentStore creates a new instance of ICommitmentStore, bound to a specific deployed contract.
func NewICommitmentStore(address common.Address, backend bind.ContractBackend) (*ICommitmentStore, error) {
	contract, err := bindICommitmentStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICommitmentStore{ICommitmentStoreCaller: ICommitmentStoreCaller{contract: contract}, ICommitmentStoreTransactor: ICommitmentStoreTransactor{contract: contract}, ICommitmentStoreFilterer: ICommitmentStoreFilterer{contract: contract}}, nil
}

// NewICommitmentStoreCaller creates a new read-only instance of ICommitmentStore, bound to a specific deployed contract.
func NewICommitmentStoreCaller(address common.Address, caller bind.ContractCaller) (*ICommitmentStoreCaller, error) {
	contract, err := bindICommitmentStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICommitmentStoreCaller{contract: contract}, nil
}

// NewICommitmentStoreTransactor creates a new write-only instance of ICommitmentStore, bound to a specific deployed contract.
func NewICommitmentStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*ICommitmentStoreTransactor, error) {
	contract, err := bindICommitmentStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICommitmentStoreTransactor{contract: contract}, nil
}

// NewICommitmentStoreFilterer creates a new log filterer instance of ICommitmentStore, bound to a specific deployed contract.
func NewICommitmentStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*ICommitmentStoreFilterer, error) {
	contract, err := bindICommitmentStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICommitmentStoreFilterer{contract: contract}, nil
}

// bindICommitmentStore binds a generic wrapper to an already deployed contract.
func bindICommitmentStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICommitmentStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICommitmentStore *ICommitmentStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICommitmentStore.Contract.ICommitmentStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICommitmentStore *ICommitmentStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICommitmentStore.Contract.ICommitmentStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICommitmentStore *ICommitmentStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICommitmentStore.Contract.ICommitmentStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICommitmentStore *ICommitmentStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICommitmentStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICommitmentStore *ICommitmentStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICommitmentStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICommitmentStore *ICommitmentStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICommitmentStore.Contract.contract.Transact(opts, method, params...)
}

// CommitmentAt is a free data retrieval call binding the contract method 0xb4c80294.
//
// Solidity: function commitmentAt(address source, uint256 height) view returns(bytes32 commitment)
func (_ICommitmentStore *ICommitmentStoreCaller) CommitmentAt(opts *bind.CallOpts, source common.Address, height *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ICommitmentStore.contract.Call(opts, &out, "commitmentAt", source, height)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CommitmentAt is a free data retrieval call binding the contract method 0xb4c80294.
//
// Solidity: function commitmentAt(address source, uint256 height) view returns(bytes32 commitment)
func (_ICommitmentStore *ICommitmentStoreSession) CommitmentAt(source common.Address, height *big.Int) ([32]byte, error) {
	return _ICommitmentStore.Contract.CommitmentAt(&_ICommitmentStore.CallOpts, source, height)
}

// CommitmentAt is a free data retrieval call binding the contract method 0xb4c80294.
//
// Solidity: function commitmentAt(address source, uint256 height) view returns(bytes32 commitment)
func (_ICommitmentStore *ICommitmentStoreCallerSession) CommitmentAt(source common.Address, height *big.Int) ([32]byte, error) {
	return _ICommitmentStore.Contract.CommitmentAt(&_ICommitmentStore.CallOpts, source, height)
}

// StoreCommitment is a paid mutator transaction binding the contract method 0x7f0d2379.
//
// Solidity: function storeCommitment(uint256 height, bytes32 commitment) returns()
func (_ICommitmentStore *ICommitmentStoreTransactor) StoreCommitment(opts *bind.TransactOpts, height *big.Int, commitment [32]byte) (*types.Transaction, error) {
	return _ICommitmentStore.contract.Transact(opts, "storeCommitment", height, commitment)
}

// StoreCommitment is a paid mutator transaction binding the contract method 0x7f0d2379.
//
// Solidity: function storeCommitment(uint256 height, bytes32 commitment) returns()
func (_ICommitmentStore *ICommitmentStoreSession) StoreCommitment(height *big.Int, commitment [32]byte) (*types.Transaction, error) {
	return _ICommitmentStore.Contract.StoreCommitment(&_ICommitmentStore.TransactOpts, height, commitment)
}

// StoreCommitment is a paid mutator transaction binding the contract method 0x7f0d2379.
//
// Solidity: function storeCommitment(uint256 height, bytes32 commitment) returns()
func (_ICommitmentStore *ICommitmentStoreTransactorSession) StoreCommitment(height *big.Int, commitment [32]byte) (*types.Transaction, error) {
	return _ICommitmentStore.Contract.StoreCommitment(&_ICommitmentStore.TransactOpts, height, commitment)
}

// ICommitmentStoreCommitmentStoredIterator is returned from FilterCommitmentStored and is used to iterate over the raw logs and unpacked data for CommitmentStored events raised by the ICommitmentStore contract.
type ICommitmentStoreCommitmentStoredIterator struct {
	Event *ICommitmentStoreCommitmentStored // Event containing the contract specifics and raw log

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
func (it *ICommitmentStoreCommitmentStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICommitmentStoreCommitmentStored)
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
		it.Event = new(ICommitmentStoreCommitmentStored)
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
func (it *ICommitmentStoreCommitmentStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICommitmentStoreCommitmentStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICommitmentStoreCommitmentStored represents a CommitmentStored event raised by the ICommitmentStore contract.
type ICommitmentStoreCommitmentStored struct {
	Source     common.Address
	Height     *big.Int
	Commitment [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommitmentStored is a free log retrieval operation binding the contract event 0x18e9b16cc8f48e098f39d821fe5c11eb557c2d91f1e058759165a1ce1266e28b.
//
// Solidity: event CommitmentStored(address indexed source, uint256 indexed height, bytes32 commitment)
func (_ICommitmentStore *ICommitmentStoreFilterer) FilterCommitmentStored(opts *bind.FilterOpts, source []common.Address, height []*big.Int) (*ICommitmentStoreCommitmentStoredIterator, error) {

	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}
	var heightRule []interface{}
	for _, heightItem := range height {
		heightRule = append(heightRule, heightItem)
	}

	logs, sub, err := _ICommitmentStore.contract.FilterLogs(opts, "CommitmentStored", sourceRule, heightRule)
	if err != nil {
		return nil, err
	}
	return &ICommitmentStoreCommitmentStoredIterator{contract: _ICommitmentStore.contract, event: "CommitmentStored", logs: logs, sub: sub}, nil
}

// WatchCommitmentStored is a free log subscription operation binding the contract event 0x18e9b16cc8f48e098f39d821fe5c11eb557c2d91f1e058759165a1ce1266e28b.
//
// Solidity: event CommitmentStored(address indexed source, uint256 indexed height, bytes32 commitment)
func (_ICommitmentStore *ICommitmentStoreFilterer) WatchCommitmentStored(opts *bind.WatchOpts, sink chan<- *ICommitmentStoreCommitmentStored, source []common.Address, height []*big.Int) (event.Subscription, error) {

	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}
	var heightRule []interface{}
	for _, heightItem := range height {
		heightRule = append(heightRule, heightItem)
	}

	logs, sub, err := _ICommitmentStore.contract.WatchLogs(opts, "CommitmentStored", sourceRule, heightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICommitmentStoreCommitmentStored)
				if err := _ICommitmentStore.contract.UnpackLog(event, "CommitmentStored", log); err != nil {
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

// ParseCommitmentStored is a log parse operation binding the contract event 0x18e9b16cc8f48e098f39d821fe5c11eb557c2d91f1e058759165a1ce1266e28b.
//
// Solidity: event CommitmentStored(address indexed source, uint256 indexed height, bytes32 commitment)
func (_ICommitmentStore *ICommitmentStoreFilterer) ParseCommitmentStored(log types.Log) (*ICommitmentStoreCommitmentStored, error) {
	event := new(ICommitmentStoreCommitmentStored)
	if err := _ICommitmentStore.contract.UnpackLog(event, "CommitmentStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
