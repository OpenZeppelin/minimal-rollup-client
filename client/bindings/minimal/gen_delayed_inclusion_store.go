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

// DelayedInclusionStoreMetaData contains all meta data concerning the DelayedInclusionStore contract.
var DelayedInclusionStoreMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"blobRefRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBlobRefRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inclusionDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publishDelayed\",\"inputs\":[{\"name\":\"blobIndices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DelayedInclusionProcessed\",\"inputs\":[{\"name\":\"inclusionsList\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structIDelayedInclusionStore.Inclusion[]\",\"components\":[{\"name\":\"blobRefHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelayedInclusionStored\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dueInclusion\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structDelayedInclusionStore.DueInclusion\",\"components\":[{\"name\":\"blobRefHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"due\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ZeroBlobRefRegistry\",\"inputs\":[]}]",
}

// DelayedInclusionStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use DelayedInclusionStoreMetaData.ABI instead.
var DelayedInclusionStoreABI = DelayedInclusionStoreMetaData.ABI

// DelayedInclusionStore is an auto generated Go binding around an Ethereum contract.
type DelayedInclusionStore struct {
	DelayedInclusionStoreCaller     // Read-only binding to the contract
	DelayedInclusionStoreTransactor // Write-only binding to the contract
	DelayedInclusionStoreFilterer   // Log filterer for contract events
}

// DelayedInclusionStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelayedInclusionStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedInclusionStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelayedInclusionStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedInclusionStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelayedInclusionStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedInclusionStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelayedInclusionStoreSession struct {
	Contract     *DelayedInclusionStore // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DelayedInclusionStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelayedInclusionStoreCallerSession struct {
	Contract *DelayedInclusionStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// DelayedInclusionStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelayedInclusionStoreTransactorSession struct {
	Contract     *DelayedInclusionStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// DelayedInclusionStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelayedInclusionStoreRaw struct {
	Contract *DelayedInclusionStore // Generic contract binding to access the raw methods on
}

// DelayedInclusionStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelayedInclusionStoreCallerRaw struct {
	Contract *DelayedInclusionStoreCaller // Generic read-only contract binding to access the raw methods on
}

// DelayedInclusionStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelayedInclusionStoreTransactorRaw struct {
	Contract *DelayedInclusionStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelayedInclusionStore creates a new instance of DelayedInclusionStore, bound to a specific deployed contract.
func NewDelayedInclusionStore(address common.Address, backend bind.ContractBackend) (*DelayedInclusionStore, error) {
	contract, err := bindDelayedInclusionStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DelayedInclusionStore{DelayedInclusionStoreCaller: DelayedInclusionStoreCaller{contract: contract}, DelayedInclusionStoreTransactor: DelayedInclusionStoreTransactor{contract: contract}, DelayedInclusionStoreFilterer: DelayedInclusionStoreFilterer{contract: contract}}, nil
}

// NewDelayedInclusionStoreCaller creates a new read-only instance of DelayedInclusionStore, bound to a specific deployed contract.
func NewDelayedInclusionStoreCaller(address common.Address, caller bind.ContractCaller) (*DelayedInclusionStoreCaller, error) {
	contract, err := bindDelayedInclusionStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelayedInclusionStoreCaller{contract: contract}, nil
}

// NewDelayedInclusionStoreTransactor creates a new write-only instance of DelayedInclusionStore, bound to a specific deployed contract.
func NewDelayedInclusionStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*DelayedInclusionStoreTransactor, error) {
	contract, err := bindDelayedInclusionStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelayedInclusionStoreTransactor{contract: contract}, nil
}

// NewDelayedInclusionStoreFilterer creates a new log filterer instance of DelayedInclusionStore, bound to a specific deployed contract.
func NewDelayedInclusionStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*DelayedInclusionStoreFilterer, error) {
	contract, err := bindDelayedInclusionStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelayedInclusionStoreFilterer{contract: contract}, nil
}

// bindDelayedInclusionStore binds a generic wrapper to an already deployed contract.
func bindDelayedInclusionStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DelayedInclusionStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelayedInclusionStore *DelayedInclusionStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelayedInclusionStore.Contract.DelayedInclusionStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelayedInclusionStore *DelayedInclusionStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedInclusionStore.Contract.DelayedInclusionStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelayedInclusionStore *DelayedInclusionStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelayedInclusionStore.Contract.DelayedInclusionStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelayedInclusionStore *DelayedInclusionStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelayedInclusionStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelayedInclusionStore *DelayedInclusionStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedInclusionStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelayedInclusionStore *DelayedInclusionStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelayedInclusionStore.Contract.contract.Transact(opts, method, params...)
}

// BlobRefRegistry is a free data retrieval call binding the contract method 0x74eb5c62.
//
// Solidity: function blobRefRegistry() view returns(address)
func (_DelayedInclusionStore *DelayedInclusionStoreCaller) BlobRefRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelayedInclusionStore.contract.Call(opts, &out, "blobRefRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlobRefRegistry is a free data retrieval call binding the contract method 0x74eb5c62.
//
// Solidity: function blobRefRegistry() view returns(address)
func (_DelayedInclusionStore *DelayedInclusionStoreSession) BlobRefRegistry() (common.Address, error) {
	return _DelayedInclusionStore.Contract.BlobRefRegistry(&_DelayedInclusionStore.CallOpts)
}

// BlobRefRegistry is a free data retrieval call binding the contract method 0x74eb5c62.
//
// Solidity: function blobRefRegistry() view returns(address)
func (_DelayedInclusionStore *DelayedInclusionStoreCallerSession) BlobRefRegistry() (common.Address, error) {
	return _DelayedInclusionStore.Contract.BlobRefRegistry(&_DelayedInclusionStore.CallOpts)
}

// InclusionDelay is a free data retrieval call binding the contract method 0xf765b4c3.
//
// Solidity: function inclusionDelay() view returns(uint256)
func (_DelayedInclusionStore *DelayedInclusionStoreCaller) InclusionDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelayedInclusionStore.contract.Call(opts, &out, "inclusionDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InclusionDelay is a free data retrieval call binding the contract method 0xf765b4c3.
//
// Solidity: function inclusionDelay() view returns(uint256)
func (_DelayedInclusionStore *DelayedInclusionStoreSession) InclusionDelay() (*big.Int, error) {
	return _DelayedInclusionStore.Contract.InclusionDelay(&_DelayedInclusionStore.CallOpts)
}

// InclusionDelay is a free data retrieval call binding the contract method 0xf765b4c3.
//
// Solidity: function inclusionDelay() view returns(uint256)
func (_DelayedInclusionStore *DelayedInclusionStoreCallerSession) InclusionDelay() (*big.Int, error) {
	return _DelayedInclusionStore.Contract.InclusionDelay(&_DelayedInclusionStore.CallOpts)
}

// PublishDelayed is a paid mutator transaction binding the contract method 0xd1831f60.
//
// Solidity: function publishDelayed(uint256[] blobIndices) returns()
func (_DelayedInclusionStore *DelayedInclusionStoreTransactor) PublishDelayed(opts *bind.TransactOpts, blobIndices []*big.Int) (*types.Transaction, error) {
	return _DelayedInclusionStore.contract.Transact(opts, "publishDelayed", blobIndices)
}

// PublishDelayed is a paid mutator transaction binding the contract method 0xd1831f60.
//
// Solidity: function publishDelayed(uint256[] blobIndices) returns()
func (_DelayedInclusionStore *DelayedInclusionStoreSession) PublishDelayed(blobIndices []*big.Int) (*types.Transaction, error) {
	return _DelayedInclusionStore.Contract.PublishDelayed(&_DelayedInclusionStore.TransactOpts, blobIndices)
}

// PublishDelayed is a paid mutator transaction binding the contract method 0xd1831f60.
//
// Solidity: function publishDelayed(uint256[] blobIndices) returns()
func (_DelayedInclusionStore *DelayedInclusionStoreTransactorSession) PublishDelayed(blobIndices []*big.Int) (*types.Transaction, error) {
	return _DelayedInclusionStore.Contract.PublishDelayed(&_DelayedInclusionStore.TransactOpts, blobIndices)
}

// DelayedInclusionStoreDelayedInclusionProcessedIterator is returned from FilterDelayedInclusionProcessed and is used to iterate over the raw logs and unpacked data for DelayedInclusionProcessed events raised by the DelayedInclusionStore contract.
type DelayedInclusionStoreDelayedInclusionProcessedIterator struct {
	Event *DelayedInclusionStoreDelayedInclusionProcessed // Event containing the contract specifics and raw log

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
func (it *DelayedInclusionStoreDelayedInclusionProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedInclusionStoreDelayedInclusionProcessed)
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
		it.Event = new(DelayedInclusionStoreDelayedInclusionProcessed)
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
func (it *DelayedInclusionStoreDelayedInclusionProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedInclusionStoreDelayedInclusionProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedInclusionStoreDelayedInclusionProcessed represents a DelayedInclusionProcessed event raised by the DelayedInclusionStore contract.
type DelayedInclusionStoreDelayedInclusionProcessed struct {
	InclusionsList []IDelayedInclusionStoreInclusion
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDelayedInclusionProcessed is a free log retrieval operation binding the contract event 0x7008b97c3a37d78dd1f286f5b621404472aee22210b1c02406212f3ee87cfe70.
//
// Solidity: event DelayedInclusionProcessed((bytes32)[] inclusionsList)
func (_DelayedInclusionStore *DelayedInclusionStoreFilterer) FilterDelayedInclusionProcessed(opts *bind.FilterOpts) (*DelayedInclusionStoreDelayedInclusionProcessedIterator, error) {

	logs, sub, err := _DelayedInclusionStore.contract.FilterLogs(opts, "DelayedInclusionProcessed")
	if err != nil {
		return nil, err
	}
	return &DelayedInclusionStoreDelayedInclusionProcessedIterator{contract: _DelayedInclusionStore.contract, event: "DelayedInclusionProcessed", logs: logs, sub: sub}, nil
}

// WatchDelayedInclusionProcessed is a free log subscription operation binding the contract event 0x7008b97c3a37d78dd1f286f5b621404472aee22210b1c02406212f3ee87cfe70.
//
// Solidity: event DelayedInclusionProcessed((bytes32)[] inclusionsList)
func (_DelayedInclusionStore *DelayedInclusionStoreFilterer) WatchDelayedInclusionProcessed(opts *bind.WatchOpts, sink chan<- *DelayedInclusionStoreDelayedInclusionProcessed) (event.Subscription, error) {

	logs, sub, err := _DelayedInclusionStore.contract.WatchLogs(opts, "DelayedInclusionProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedInclusionStoreDelayedInclusionProcessed)
				if err := _DelayedInclusionStore.contract.UnpackLog(event, "DelayedInclusionProcessed", log); err != nil {
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
func (_DelayedInclusionStore *DelayedInclusionStoreFilterer) ParseDelayedInclusionProcessed(log types.Log) (*DelayedInclusionStoreDelayedInclusionProcessed, error) {
	event := new(DelayedInclusionStoreDelayedInclusionProcessed)
	if err := _DelayedInclusionStore.contract.UnpackLog(event, "DelayedInclusionProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedInclusionStoreDelayedInclusionStoredIterator is returned from FilterDelayedInclusionStored and is used to iterate over the raw logs and unpacked data for DelayedInclusionStored events raised by the DelayedInclusionStore contract.
type DelayedInclusionStoreDelayedInclusionStoredIterator struct {
	Event *DelayedInclusionStoreDelayedInclusionStored // Event containing the contract specifics and raw log

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
func (it *DelayedInclusionStoreDelayedInclusionStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedInclusionStoreDelayedInclusionStored)
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
		it.Event = new(DelayedInclusionStoreDelayedInclusionStored)
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
func (it *DelayedInclusionStoreDelayedInclusionStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedInclusionStoreDelayedInclusionStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedInclusionStoreDelayedInclusionStored represents a DelayedInclusionStored event raised by the DelayedInclusionStore contract.
type DelayedInclusionStoreDelayedInclusionStored struct {
	Sender       common.Address
	DueInclusion DelayedInclusionStoreDueInclusion
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelayedInclusionStored is a free log retrieval operation binding the contract event 0x6d87b8f648fd37d1ff7a2ecf57c51a3ac9830fc10bfc716f48223cb19de5b052.
//
// Solidity: event DelayedInclusionStored(address indexed sender, (bytes32,uint256) dueInclusion)
func (_DelayedInclusionStore *DelayedInclusionStoreFilterer) FilterDelayedInclusionStored(opts *bind.FilterOpts, sender []common.Address) (*DelayedInclusionStoreDelayedInclusionStoredIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DelayedInclusionStore.contract.FilterLogs(opts, "DelayedInclusionStored", senderRule)
	if err != nil {
		return nil, err
	}
	return &DelayedInclusionStoreDelayedInclusionStoredIterator{contract: _DelayedInclusionStore.contract, event: "DelayedInclusionStored", logs: logs, sub: sub}, nil
}

// WatchDelayedInclusionStored is a free log subscription operation binding the contract event 0x6d87b8f648fd37d1ff7a2ecf57c51a3ac9830fc10bfc716f48223cb19de5b052.
//
// Solidity: event DelayedInclusionStored(address indexed sender, (bytes32,uint256) dueInclusion)
func (_DelayedInclusionStore *DelayedInclusionStoreFilterer) WatchDelayedInclusionStored(opts *bind.WatchOpts, sink chan<- *DelayedInclusionStoreDelayedInclusionStored, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DelayedInclusionStore.contract.WatchLogs(opts, "DelayedInclusionStored", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedInclusionStoreDelayedInclusionStored)
				if err := _DelayedInclusionStore.contract.UnpackLog(event, "DelayedInclusionStored", log); err != nil {
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

// ParseDelayedInclusionStored is a log parse operation binding the contract event 0x6d87b8f648fd37d1ff7a2ecf57c51a3ac9830fc10bfc716f48223cb19de5b052.
//
// Solidity: event DelayedInclusionStored(address indexed sender, (bytes32,uint256) dueInclusion)
func (_DelayedInclusionStore *DelayedInclusionStoreFilterer) ParseDelayedInclusionStored(log types.Log) (*DelayedInclusionStoreDelayedInclusionStored, error) {
	event := new(DelayedInclusionStoreDelayedInclusionStored)
	if err := _DelayedInclusionStore.contract.UnpackLog(event, "DelayedInclusionStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
