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

// ICheckpointTrackerCheckpoint is an auto generated low-level Go binding around an user-defined struct.
type ICheckpointTrackerCheckpoint struct {
	PublicationId            *big.Int
	Commitment               [32]byte
	TotalDelayedPublications *big.Int
}

// ICheckpointTrackerMetaData contains all meta data concerning the ICheckpointTracker contract.
var ICheckpointTrackerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"proveTransition\",\"inputs\":[{\"name\":\"start\",\"type\":\"tuple\",\"internalType\":\"structICheckpointTracker.Checkpoint\",\"components\":[{\"name\":\"publicationId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"totalDelayedPublications\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"end\",\"type\":\"tuple\",\"internalType\":\"structICheckpointTracker.Checkpoint\",\"components\":[{\"name\":\"publicationId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"totalDelayedPublications\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provenPublicationId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CommitmentSaved\",\"inputs\":[{\"name\":\"publicationId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"commitment\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EndPublicationNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExcessiveDelayedPublications\",\"inputs\":[{\"name\":\"delayedCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStartPublication\",\"inputs\":[{\"name\":\"startPublicationId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"latestProvenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlyProverManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroEndCommitment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroGenesisCommitment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroStartCommitment\",\"inputs\":[]}]",
}

// ICheckpointTrackerABI is the input ABI used to generate the binding from.
// Deprecated: Use ICheckpointTrackerMetaData.ABI instead.
var ICheckpointTrackerABI = ICheckpointTrackerMetaData.ABI

// ICheckpointTracker is an auto generated Go binding around an Ethereum contract.
type ICheckpointTracker struct {
	ICheckpointTrackerCaller     // Read-only binding to the contract
	ICheckpointTrackerTransactor // Write-only binding to the contract
	ICheckpointTrackerFilterer   // Log filterer for contract events
}

// ICheckpointTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICheckpointTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckpointTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICheckpointTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckpointTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICheckpointTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICheckpointTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICheckpointTrackerSession struct {
	Contract     *ICheckpointTracker // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ICheckpointTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICheckpointTrackerCallerSession struct {
	Contract *ICheckpointTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ICheckpointTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICheckpointTrackerTransactorSession struct {
	Contract     *ICheckpointTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ICheckpointTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICheckpointTrackerRaw struct {
	Contract *ICheckpointTracker // Generic contract binding to access the raw methods on
}

// ICheckpointTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICheckpointTrackerCallerRaw struct {
	Contract *ICheckpointTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// ICheckpointTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICheckpointTrackerTransactorRaw struct {
	Contract *ICheckpointTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICheckpointTracker creates a new instance of ICheckpointTracker, bound to a specific deployed contract.
func NewICheckpointTracker(address common.Address, backend bind.ContractBackend) (*ICheckpointTracker, error) {
	contract, err := bindICheckpointTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICheckpointTracker{ICheckpointTrackerCaller: ICheckpointTrackerCaller{contract: contract}, ICheckpointTrackerTransactor: ICheckpointTrackerTransactor{contract: contract}, ICheckpointTrackerFilterer: ICheckpointTrackerFilterer{contract: contract}}, nil
}

// NewICheckpointTrackerCaller creates a new read-only instance of ICheckpointTracker, bound to a specific deployed contract.
func NewICheckpointTrackerCaller(address common.Address, caller bind.ContractCaller) (*ICheckpointTrackerCaller, error) {
	contract, err := bindICheckpointTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICheckpointTrackerCaller{contract: contract}, nil
}

// NewICheckpointTrackerTransactor creates a new write-only instance of ICheckpointTracker, bound to a specific deployed contract.
func NewICheckpointTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*ICheckpointTrackerTransactor, error) {
	contract, err := bindICheckpointTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICheckpointTrackerTransactor{contract: contract}, nil
}

// NewICheckpointTrackerFilterer creates a new log filterer instance of ICheckpointTracker, bound to a specific deployed contract.
func NewICheckpointTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*ICheckpointTrackerFilterer, error) {
	contract, err := bindICheckpointTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICheckpointTrackerFilterer{contract: contract}, nil
}

// bindICheckpointTracker binds a generic wrapper to an already deployed contract.
func bindICheckpointTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICheckpointTrackerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICheckpointTracker *ICheckpointTrackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICheckpointTracker.Contract.ICheckpointTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICheckpointTracker *ICheckpointTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICheckpointTracker.Contract.ICheckpointTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICheckpointTracker *ICheckpointTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICheckpointTracker.Contract.ICheckpointTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICheckpointTracker *ICheckpointTrackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICheckpointTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICheckpointTracker *ICheckpointTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICheckpointTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICheckpointTracker *ICheckpointTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICheckpointTracker.Contract.contract.Transact(opts, method, params...)
}

// ProvenPublicationId is a free data retrieval call binding the contract method 0x3f558a2a.
//
// Solidity: function provenPublicationId() view returns(uint256)
func (_ICheckpointTracker *ICheckpointTrackerCaller) ProvenPublicationId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ICheckpointTracker.contract.Call(opts, &out, "provenPublicationId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProvenPublicationId is a free data retrieval call binding the contract method 0x3f558a2a.
//
// Solidity: function provenPublicationId() view returns(uint256)
func (_ICheckpointTracker *ICheckpointTrackerSession) ProvenPublicationId() (*big.Int, error) {
	return _ICheckpointTracker.Contract.ProvenPublicationId(&_ICheckpointTracker.CallOpts)
}

// ProvenPublicationId is a free data retrieval call binding the contract method 0x3f558a2a.
//
// Solidity: function provenPublicationId() view returns(uint256)
func (_ICheckpointTracker *ICheckpointTrackerCallerSession) ProvenPublicationId() (*big.Int, error) {
	return _ICheckpointTracker.Contract.ProvenPublicationId(&_ICheckpointTracker.CallOpts)
}

// ProveTransition is a paid mutator transaction binding the contract method 0xc1ce1eb9.
//
// Solidity: function proveTransition((uint256,bytes32,uint256) start, (uint256,bytes32,uint256) end, bytes proof) returns(uint256, uint256)
func (_ICheckpointTracker *ICheckpointTrackerTransactor) ProveTransition(opts *bind.TransactOpts, start ICheckpointTrackerCheckpoint, end ICheckpointTrackerCheckpoint, proof []byte) (*types.Transaction, error) {
	return _ICheckpointTracker.contract.Transact(opts, "proveTransition", start, end, proof)
}

// ProveTransition is a paid mutator transaction binding the contract method 0xc1ce1eb9.
//
// Solidity: function proveTransition((uint256,bytes32,uint256) start, (uint256,bytes32,uint256) end, bytes proof) returns(uint256, uint256)
func (_ICheckpointTracker *ICheckpointTrackerSession) ProveTransition(start ICheckpointTrackerCheckpoint, end ICheckpointTrackerCheckpoint, proof []byte) (*types.Transaction, error) {
	return _ICheckpointTracker.Contract.ProveTransition(&_ICheckpointTracker.TransactOpts, start, end, proof)
}

// ProveTransition is a paid mutator transaction binding the contract method 0xc1ce1eb9.
//
// Solidity: function proveTransition((uint256,bytes32,uint256) start, (uint256,bytes32,uint256) end, bytes proof) returns(uint256, uint256)
func (_ICheckpointTracker *ICheckpointTrackerTransactorSession) ProveTransition(start ICheckpointTrackerCheckpoint, end ICheckpointTrackerCheckpoint, proof []byte) (*types.Transaction, error) {
	return _ICheckpointTracker.Contract.ProveTransition(&_ICheckpointTracker.TransactOpts, start, end, proof)
}

// ICheckpointTrackerCommitmentSavedIterator is returned from FilterCommitmentSaved and is used to iterate over the raw logs and unpacked data for CommitmentSaved events raised by the ICheckpointTracker contract.
type ICheckpointTrackerCommitmentSavedIterator struct {
	Event *ICheckpointTrackerCommitmentSaved // Event containing the contract specifics and raw log

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
func (it *ICheckpointTrackerCommitmentSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICheckpointTrackerCommitmentSaved)
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
		it.Event = new(ICheckpointTrackerCommitmentSaved)
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
func (it *ICheckpointTrackerCommitmentSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICheckpointTrackerCommitmentSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICheckpointTrackerCommitmentSaved represents a CommitmentSaved event raised by the ICheckpointTracker contract.
type ICheckpointTrackerCommitmentSaved struct {
	PublicationId *big.Int
	Commitment    [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCommitmentSaved is a free log retrieval operation binding the contract event 0x42898ee6cff40a090090a9f2cf484511de9d1b1b735396cd2307646d2fb6f890.
//
// Solidity: event CommitmentSaved(uint256 indexed publicationId, bytes32 commitment)
func (_ICheckpointTracker *ICheckpointTrackerFilterer) FilterCommitmentSaved(opts *bind.FilterOpts, publicationId []*big.Int) (*ICheckpointTrackerCommitmentSavedIterator, error) {

	var publicationIdRule []interface{}
	for _, publicationIdItem := range publicationId {
		publicationIdRule = append(publicationIdRule, publicationIdItem)
	}

	logs, sub, err := _ICheckpointTracker.contract.FilterLogs(opts, "CommitmentSaved", publicationIdRule)
	if err != nil {
		return nil, err
	}
	return &ICheckpointTrackerCommitmentSavedIterator{contract: _ICheckpointTracker.contract, event: "CommitmentSaved", logs: logs, sub: sub}, nil
}

// WatchCommitmentSaved is a free log subscription operation binding the contract event 0x42898ee6cff40a090090a9f2cf484511de9d1b1b735396cd2307646d2fb6f890.
//
// Solidity: event CommitmentSaved(uint256 indexed publicationId, bytes32 commitment)
func (_ICheckpointTracker *ICheckpointTrackerFilterer) WatchCommitmentSaved(opts *bind.WatchOpts, sink chan<- *ICheckpointTrackerCommitmentSaved, publicationId []*big.Int) (event.Subscription, error) {

	var publicationIdRule []interface{}
	for _, publicationIdItem := range publicationId {
		publicationIdRule = append(publicationIdRule, publicationIdItem)
	}

	logs, sub, err := _ICheckpointTracker.contract.WatchLogs(opts, "CommitmentSaved", publicationIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICheckpointTrackerCommitmentSaved)
				if err := _ICheckpointTracker.contract.UnpackLog(event, "CommitmentSaved", log); err != nil {
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

// ParseCommitmentSaved is a log parse operation binding the contract event 0x42898ee6cff40a090090a9f2cf484511de9d1b1b735396cd2307646d2fb6f890.
//
// Solidity: event CommitmentSaved(uint256 indexed publicationId, bytes32 commitment)
func (_ICheckpointTracker *ICheckpointTrackerFilterer) ParseCommitmentSaved(log types.Log) (*ICheckpointTrackerCommitmentSaved, error) {
	event := new(ICheckpointTrackerCommitmentSaved)
	if err := _ICheckpointTracker.contract.UnpackLog(event, "CommitmentSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
