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

// IDelayedInclusionStoreInclusion is an auto generated low-level Go binding around an user-defined struct.
type IDelayedInclusionStoreInclusion struct {
	BlobRefHash [32]byte
}

// IDelayedInclusionStoreMetaData contains all meta data concerning the IDelayedInclusionStore contract.
var IDelayedInclusionStoreMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"publishDelayed\",\"inputs\":[{\"name\":\"blobIndices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DelayedInclusionProcessed\",\"inputs\":[{\"name\":\"inclusionsList\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structIDelayedInclusionStore.Inclusion[]\",\"components\":[{\"name\":\"blobRefHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ZeroBlobRefRegistry\",\"inputs\":[]}]",
}

// IDelayedInclusionStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use IDelayedInclusionStoreMetaData.ABI instead.
var IDelayedInclusionStoreABI = IDelayedInclusionStoreMetaData.ABI

// IDelayedInclusionStore is an auto generated Go binding around an Ethereum contract.
type IDelayedInclusionStore struct {
	IDelayedInclusionStoreCaller     // Read-only binding to the contract
	IDelayedInclusionStoreTransactor // Write-only binding to the contract
	IDelayedInclusionStoreFilterer   // Log filterer for contract events
}

// IDelayedInclusionStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDelayedInclusionStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelayedInclusionStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDelayedInclusionStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelayedInclusionStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDelayedInclusionStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelayedInclusionStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDelayedInclusionStoreSession struct {
	Contract     *IDelayedInclusionStore // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IDelayedInclusionStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDelayedInclusionStoreCallerSession struct {
	Contract *IDelayedInclusionStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IDelayedInclusionStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDelayedInclusionStoreTransactorSession struct {
	Contract     *IDelayedInclusionStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IDelayedInclusionStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDelayedInclusionStoreRaw struct {
	Contract *IDelayedInclusionStore // Generic contract binding to access the raw methods on
}

// IDelayedInclusionStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDelayedInclusionStoreCallerRaw struct {
	Contract *IDelayedInclusionStoreCaller // Generic read-only contract binding to access the raw methods on
}

// IDelayedInclusionStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDelayedInclusionStoreTransactorRaw struct {
	Contract *IDelayedInclusionStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDelayedInclusionStore creates a new instance of IDelayedInclusionStore, bound to a specific deployed contract.
func NewIDelayedInclusionStore(address common.Address, backend bind.ContractBackend) (*IDelayedInclusionStore, error) {
	contract, err := bindIDelayedInclusionStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDelayedInclusionStore{IDelayedInclusionStoreCaller: IDelayedInclusionStoreCaller{contract: contract}, IDelayedInclusionStoreTransactor: IDelayedInclusionStoreTransactor{contract: contract}, IDelayedInclusionStoreFilterer: IDelayedInclusionStoreFilterer{contract: contract}}, nil
}

// NewIDelayedInclusionStoreCaller creates a new read-only instance of IDelayedInclusionStore, bound to a specific deployed contract.
func NewIDelayedInclusionStoreCaller(address common.Address, caller bind.ContractCaller) (*IDelayedInclusionStoreCaller, error) {
	contract, err := bindIDelayedInclusionStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDelayedInclusionStoreCaller{contract: contract}, nil
}

// NewIDelayedInclusionStoreTransactor creates a new write-only instance of IDelayedInclusionStore, bound to a specific deployed contract.
func NewIDelayedInclusionStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*IDelayedInclusionStoreTransactor, error) {
	contract, err := bindIDelayedInclusionStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDelayedInclusionStoreTransactor{contract: contract}, nil
}

// NewIDelayedInclusionStoreFilterer creates a new log filterer instance of IDelayedInclusionStore, bound to a specific deployed contract.
func NewIDelayedInclusionStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*IDelayedInclusionStoreFilterer, error) {
	contract, err := bindIDelayedInclusionStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDelayedInclusionStoreFilterer{contract: contract}, nil
}

// bindIDelayedInclusionStore binds a generic wrapper to an already deployed contract.
func bindIDelayedInclusionStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IDelayedInclusionStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDelayedInclusionStore *IDelayedInclusionStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDelayedInclusionStore.Contract.IDelayedInclusionStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDelayedInclusionStore *IDelayedInclusionStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDelayedInclusionStore.Contract.IDelayedInclusionStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDelayedInclusionStore *IDelayedInclusionStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDelayedInclusionStore.Contract.IDelayedInclusionStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDelayedInclusionStore *IDelayedInclusionStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDelayedInclusionStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDelayedInclusionStore *IDelayedInclusionStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDelayedInclusionStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDelayedInclusionStore *IDelayedInclusionStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDelayedInclusionStore.Contract.contract.Transact(opts, method, params...)
}

// PublishDelayed is a paid mutator transaction binding the contract method 0xd1831f60.
//
// Solidity: function publishDelayed(uint256[] blobIndices) returns()
func (_IDelayedInclusionStore *IDelayedInclusionStoreTransactor) PublishDelayed(opts *bind.TransactOpts, blobIndices []*big.Int) (*types.Transaction, error) {
	return _IDelayedInclusionStore.contract.Transact(opts, "publishDelayed", blobIndices)
}

// PublishDelayed is a paid mutator transaction binding the contract method 0xd1831f60.
//
// Solidity: function publishDelayed(uint256[] blobIndices) returns()
func (_IDelayedInclusionStore *IDelayedInclusionStoreSession) PublishDelayed(blobIndices []*big.Int) (*types.Transaction, error) {
	return _IDelayedInclusionStore.Contract.PublishDelayed(&_IDelayedInclusionStore.TransactOpts, blobIndices)
}

// PublishDelayed is a paid mutator transaction binding the contract method 0xd1831f60.
//
// Solidity: function publishDelayed(uint256[] blobIndices) returns()
func (_IDelayedInclusionStore *IDelayedInclusionStoreTransactorSession) PublishDelayed(blobIndices []*big.Int) (*types.Transaction, error) {
	return _IDelayedInclusionStore.Contract.PublishDelayed(&_IDelayedInclusionStore.TransactOpts, blobIndices)
}

// IDelayedInclusionStoreDelayedInclusionProcessedIterator is returned from FilterDelayedInclusionProcessed and is used to iterate over the raw logs and unpacked data for DelayedInclusionProcessed events raised by the IDelayedInclusionStore contract.
type IDelayedInclusionStoreDelayedInclusionProcessedIterator struct {
	Event *IDelayedInclusionStoreDelayedInclusionProcessed // Event containing the contract specifics and raw log

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
func (it *IDelayedInclusionStoreDelayedInclusionProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDelayedInclusionStoreDelayedInclusionProcessed)
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
		it.Event = new(IDelayedInclusionStoreDelayedInclusionProcessed)
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
func (it *IDelayedInclusionStoreDelayedInclusionProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDelayedInclusionStoreDelayedInclusionProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDelayedInclusionStoreDelayedInclusionProcessed represents a DelayedInclusionProcessed event raised by the IDelayedInclusionStore contract.
type IDelayedInclusionStoreDelayedInclusionProcessed struct {
	InclusionsList []IDelayedInclusionStoreInclusion
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDelayedInclusionProcessed is a free log retrieval operation binding the contract event 0x7008b97c3a37d78dd1f286f5b621404472aee22210b1c02406212f3ee87cfe70.
//
// Solidity: event DelayedInclusionProcessed((bytes32)[] inclusionsList)
func (_IDelayedInclusionStore *IDelayedInclusionStoreFilterer) FilterDelayedInclusionProcessed(opts *bind.FilterOpts) (*IDelayedInclusionStoreDelayedInclusionProcessedIterator, error) {

	logs, sub, err := _IDelayedInclusionStore.contract.FilterLogs(opts, "DelayedInclusionProcessed")
	if err != nil {
		return nil, err
	}
	return &IDelayedInclusionStoreDelayedInclusionProcessedIterator{contract: _IDelayedInclusionStore.contract, event: "DelayedInclusionProcessed", logs: logs, sub: sub}, nil
}

// WatchDelayedInclusionProcessed is a free log subscription operation binding the contract event 0x7008b97c3a37d78dd1f286f5b621404472aee22210b1c02406212f3ee87cfe70.
//
// Solidity: event DelayedInclusionProcessed((bytes32)[] inclusionsList)
func (_IDelayedInclusionStore *IDelayedInclusionStoreFilterer) WatchDelayedInclusionProcessed(opts *bind.WatchOpts, sink chan<- *IDelayedInclusionStoreDelayedInclusionProcessed) (event.Subscription, error) {

	logs, sub, err := _IDelayedInclusionStore.contract.WatchLogs(opts, "DelayedInclusionProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDelayedInclusionStoreDelayedInclusionProcessed)
				if err := _IDelayedInclusionStore.contract.UnpackLog(event, "DelayedInclusionProcessed", log); err != nil {
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

// ParseDelayedInclusionProcessed is a log parse operation binding the contract event 0x7008b97c3a37d78dd1f286f5b621404472aee22210b1c02406212f3ee87cfe70.
//
// Solidity: event DelayedInclusionProcessed((bytes32)[] inclusionsList)
func (_IDelayedInclusionStore *IDelayedInclusionStoreFilterer) ParseDelayedInclusionProcessed(log types.Log) (*IDelayedInclusionStoreDelayedInclusionProcessed, error) {
	event := new(IDelayedInclusionStoreDelayedInclusionProcessed)
	if err := _IDelayedInclusionStore.contract.UnpackLog(event, "DelayedInclusionProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
