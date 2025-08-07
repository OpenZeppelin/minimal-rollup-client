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

// IInboxPublicationHeader is an auto generated low-level Go binding around an user-defined struct.
type IInboxPublicationHeader struct {
	Id             *big.Int
	PrevHash       [32]byte
	Publisher      common.Address
	Timestamp      *big.Int
	BlockNumber    *big.Int
	AttributesHash [32]byte
}

// IProverManagerMetaData contains all meta data concerning the IProverManager contract.
var IProverManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"bid\",\"inputs\":[{\"name\":\"offeredFee\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimProvingVacancy\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"evictProver\",\"inputs\":[{\"name\":\"publicationHeader\",\"type\":\"tuple\",\"internalType\":\"structIInbox.PublicationHeader\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"prevHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"publisher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"attributesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"exit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizePastPeriod\",\"inputs\":[{\"name\":\"periodId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"provenPublication\",\"type\":\"tuple\",\"internalType\":\"structIInbox.PublicationHeader\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"prevHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"publisher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"attributesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"prove\",\"inputs\":[{\"name\":\"start\",\"type\":\"tuple\",\"internalType\":\"structICheckpointTracker.Checkpoint\",\"components\":[{\"name\":\"publicationId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"totalDelayedPublications\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"end\",\"type\":\"tuple\",\"internalType\":\"structICheckpointTracker.Checkpoint\",\"components\":[{\"name\":\"publicationId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"totalDelayedPublications\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"firstPub\",\"type\":\"tuple\",\"internalType\":\"structIInbox.PublicationHeader\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"prevHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"publisher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"attributesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"lastPub\",\"type\":\"tuple\",\"internalType\":\"structIInbox.PublicationHeader\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"prevHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"publisher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"attributesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"periodId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"NewPeriod\",\"inputs\":[{\"name\":\"periodId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProverEvicted\",\"inputs\":[{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"evictor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"periodEnd\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"livenessBond\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProverExited\",\"inputs\":[{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"periodEnd\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"provingDeadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProverOffer\",\"inputs\":[{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"periodId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"stake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProverVacancyClaimed\",\"inputs\":[{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"periodId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"stake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// IProverManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IProverManagerMetaData.ABI instead.
var IProverManagerABI = IProverManagerMetaData.ABI

// IProverManager is an auto generated Go binding around an Ethereum contract.
type IProverManager struct {
	IProverManagerCaller     // Read-only binding to the contract
	IProverManagerTransactor // Write-only binding to the contract
	IProverManagerFilterer   // Log filterer for contract events
}

// IProverManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IProverManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProverManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IProverManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProverManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IProverManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProverManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IProverManagerSession struct {
	Contract     *IProverManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IProverManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IProverManagerCallerSession struct {
	Contract *IProverManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IProverManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IProverManagerTransactorSession struct {
	Contract     *IProverManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IProverManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IProverManagerRaw struct {
	Contract *IProverManager // Generic contract binding to access the raw methods on
}

// IProverManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IProverManagerCallerRaw struct {
	Contract *IProverManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IProverManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IProverManagerTransactorRaw struct {
	Contract *IProverManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIProverManager creates a new instance of IProverManager, bound to a specific deployed contract.
func NewIProverManager(address common.Address, backend bind.ContractBackend) (*IProverManager, error) {
	contract, err := bindIProverManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IProverManager{IProverManagerCaller: IProverManagerCaller{contract: contract}, IProverManagerTransactor: IProverManagerTransactor{contract: contract}, IProverManagerFilterer: IProverManagerFilterer{contract: contract}}, nil
}

// NewIProverManagerCaller creates a new read-only instance of IProverManager, bound to a specific deployed contract.
func NewIProverManagerCaller(address common.Address, caller bind.ContractCaller) (*IProverManagerCaller, error) {
	contract, err := bindIProverManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IProverManagerCaller{contract: contract}, nil
}

// NewIProverManagerTransactor creates a new write-only instance of IProverManager, bound to a specific deployed contract.
func NewIProverManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IProverManagerTransactor, error) {
	contract, err := bindIProverManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IProverManagerTransactor{contract: contract}, nil
}

// NewIProverManagerFilterer creates a new log filterer instance of IProverManager, bound to a specific deployed contract.
func NewIProverManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IProverManagerFilterer, error) {
	contract, err := bindIProverManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IProverManagerFilterer{contract: contract}, nil
}

// bindIProverManager binds a generic wrapper to an already deployed contract.
func bindIProverManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IProverManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProverManager *IProverManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IProverManager.Contract.IProverManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProverManager *IProverManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProverManager.Contract.IProverManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProverManager *IProverManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProverManager.Contract.IProverManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProverManager *IProverManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IProverManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProverManager *IProverManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProverManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProverManager *IProverManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProverManager.Contract.contract.Transact(opts, method, params...)
}

// Bid is a paid mutator transaction binding the contract method 0x9d8f9379.
//
// Solidity: function bid(uint96 offeredFee) returns()
func (_IProverManager *IProverManagerTransactor) Bid(opts *bind.TransactOpts, offeredFee *big.Int) (*types.Transaction, error) {
	return _IProverManager.contract.Transact(opts, "bid", offeredFee)
}

// Bid is a paid mutator transaction binding the contract method 0x9d8f9379.
//
// Solidity: function bid(uint96 offeredFee) returns()
func (_IProverManager *IProverManagerSession) Bid(offeredFee *big.Int) (*types.Transaction, error) {
	return _IProverManager.Contract.Bid(&_IProverManager.TransactOpts, offeredFee)
}

// Bid is a paid mutator transaction binding the contract method 0x9d8f9379.
//
// Solidity: function bid(uint96 offeredFee) returns()
func (_IProverManager *IProverManagerTransactorSession) Bid(offeredFee *big.Int) (*types.Transaction, error) {
	return _IProverManager.Contract.Bid(&_IProverManager.TransactOpts, offeredFee)
}

// ClaimProvingVacancy is a paid mutator transaction binding the contract method 0x2cf10a14.
//
// Solidity: function claimProvingVacancy(uint96 fee) returns()
func (_IProverManager *IProverManagerTransactor) ClaimProvingVacancy(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _IProverManager.contract.Transact(opts, "claimProvingVacancy", fee)
}

// ClaimProvingVacancy is a paid mutator transaction binding the contract method 0x2cf10a14.
//
// Solidity: function claimProvingVacancy(uint96 fee) returns()
func (_IProverManager *IProverManagerSession) ClaimProvingVacancy(fee *big.Int) (*types.Transaction, error) {
	return _IProverManager.Contract.ClaimProvingVacancy(&_IProverManager.TransactOpts, fee)
}

// ClaimProvingVacancy is a paid mutator transaction binding the contract method 0x2cf10a14.
//
// Solidity: function claimProvingVacancy(uint96 fee) returns()
func (_IProverManager *IProverManagerTransactorSession) ClaimProvingVacancy(fee *big.Int) (*types.Transaction, error) {
	return _IProverManager.Contract.ClaimProvingVacancy(&_IProverManager.TransactOpts, fee)
}

// EvictProver is a paid mutator transaction binding the contract method 0x1cf44b92.
//
// Solidity: function evictProver((uint256,bytes32,address,uint256,uint256,bytes32) publicationHeader) returns()
func (_IProverManager *IProverManagerTransactor) EvictProver(opts *bind.TransactOpts, publicationHeader IInboxPublicationHeader) (*types.Transaction, error) {
	return _IProverManager.contract.Transact(opts, "evictProver", publicationHeader)
}

// EvictProver is a paid mutator transaction binding the contract method 0x1cf44b92.
//
// Solidity: function evictProver((uint256,bytes32,address,uint256,uint256,bytes32) publicationHeader) returns()
func (_IProverManager *IProverManagerSession) EvictProver(publicationHeader IInboxPublicationHeader) (*types.Transaction, error) {
	return _IProverManager.Contract.EvictProver(&_IProverManager.TransactOpts, publicationHeader)
}

// EvictProver is a paid mutator transaction binding the contract method 0x1cf44b92.
//
// Solidity: function evictProver((uint256,bytes32,address,uint256,uint256,bytes32) publicationHeader) returns()
func (_IProverManager *IProverManagerTransactorSession) EvictProver(publicationHeader IInboxPublicationHeader) (*types.Transaction, error) {
	return _IProverManager.Contract.EvictProver(&_IProverManager.TransactOpts, publicationHeader)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_IProverManager *IProverManagerTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProverManager.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_IProverManager *IProverManagerSession) Exit() (*types.Transaction, error) {
	return _IProverManager.Contract.Exit(&_IProverManager.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_IProverManager *IProverManagerTransactorSession) Exit() (*types.Transaction, error) {
	return _IProverManager.Contract.Exit(&_IProverManager.TransactOpts)
}

// FinalizePastPeriod is a paid mutator transaction binding the contract method 0x0f83c9ba.
//
// Solidity: function finalizePastPeriod(uint256 periodId, (uint256,bytes32,address,uint256,uint256,bytes32) provenPublication) returns()
func (_IProverManager *IProverManagerTransactor) FinalizePastPeriod(opts *bind.TransactOpts, periodId *big.Int, provenPublication IInboxPublicationHeader) (*types.Transaction, error) {
	return _IProverManager.contract.Transact(opts, "finalizePastPeriod", periodId, provenPublication)
}

// FinalizePastPeriod is a paid mutator transaction binding the contract method 0x0f83c9ba.
//
// Solidity: function finalizePastPeriod(uint256 periodId, (uint256,bytes32,address,uint256,uint256,bytes32) provenPublication) returns()
func (_IProverManager *IProverManagerSession) FinalizePastPeriod(periodId *big.Int, provenPublication IInboxPublicationHeader) (*types.Transaction, error) {
	return _IProverManager.Contract.FinalizePastPeriod(&_IProverManager.TransactOpts, periodId, provenPublication)
}

// FinalizePastPeriod is a paid mutator transaction binding the contract method 0x0f83c9ba.
//
// Solidity: function finalizePastPeriod(uint256 periodId, (uint256,bytes32,address,uint256,uint256,bytes32) provenPublication) returns()
func (_IProverManager *IProverManagerTransactorSession) FinalizePastPeriod(periodId *big.Int, provenPublication IInboxPublicationHeader) (*types.Transaction, error) {
	return _IProverManager.Contract.FinalizePastPeriod(&_IProverManager.TransactOpts, periodId, provenPublication)
}

// Prove is a paid mutator transaction binding the contract method 0xd2e5eac5.
//
// Solidity: function prove((uint256,bytes32,uint256) start, (uint256,bytes32,uint256) end, (uint256,bytes32,address,uint256,uint256,bytes32) firstPub, (uint256,bytes32,address,uint256,uint256,bytes32) lastPub, bytes proof, uint256 periodId) returns()
func (_IProverManager *IProverManagerTransactor) Prove(opts *bind.TransactOpts, start ICheckpointTrackerCheckpoint, end ICheckpointTrackerCheckpoint, firstPub IInboxPublicationHeader, lastPub IInboxPublicationHeader, proof []byte, periodId *big.Int) (*types.Transaction, error) {
	return _IProverManager.contract.Transact(opts, "prove", start, end, firstPub, lastPub, proof, periodId)
}

// Prove is a paid mutator transaction binding the contract method 0xd2e5eac5.
//
// Solidity: function prove((uint256,bytes32,uint256) start, (uint256,bytes32,uint256) end, (uint256,bytes32,address,uint256,uint256,bytes32) firstPub, (uint256,bytes32,address,uint256,uint256,bytes32) lastPub, bytes proof, uint256 periodId) returns()
func (_IProverManager *IProverManagerSession) Prove(start ICheckpointTrackerCheckpoint, end ICheckpointTrackerCheckpoint, firstPub IInboxPublicationHeader, lastPub IInboxPublicationHeader, proof []byte, periodId *big.Int) (*types.Transaction, error) {
	return _IProverManager.Contract.Prove(&_IProverManager.TransactOpts, start, end, firstPub, lastPub, proof, periodId)
}

// Prove is a paid mutator transaction binding the contract method 0xd2e5eac5.
//
// Solidity: function prove((uint256,bytes32,uint256) start, (uint256,bytes32,uint256) end, (uint256,bytes32,address,uint256,uint256,bytes32) firstPub, (uint256,bytes32,address,uint256,uint256,bytes32) lastPub, bytes proof, uint256 periodId) returns()
func (_IProverManager *IProverManagerTransactorSession) Prove(start ICheckpointTrackerCheckpoint, end ICheckpointTrackerCheckpoint, firstPub IInboxPublicationHeader, lastPub IInboxPublicationHeader, proof []byte, periodId *big.Int) (*types.Transaction, error) {
	return _IProverManager.Contract.Prove(&_IProverManager.TransactOpts, start, end, firstPub, lastPub, proof, periodId)
}

// IProverManagerNewPeriodIterator is returned from FilterNewPeriod and is used to iterate over the raw logs and unpacked data for NewPeriod events raised by the IProverManager contract.
type IProverManagerNewPeriodIterator struct {
	Event *IProverManagerNewPeriod // Event containing the contract specifics and raw log

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
func (it *IProverManagerNewPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProverManagerNewPeriod)
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
		it.Event = new(IProverManagerNewPeriod)
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
func (it *IProverManagerNewPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProverManagerNewPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProverManagerNewPeriod represents a NewPeriod event raised by the IProverManager contract.
type IProverManagerNewPeriod struct {
	PeriodId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewPeriod is a free log retrieval operation binding the contract event 0x61a611267e7ed28f8a566b021b9ac3ccc3985343a31971a180d01a57f63f3380.
//
// Solidity: event NewPeriod(uint256 periodId)
func (_IProverManager *IProverManagerFilterer) FilterNewPeriod(opts *bind.FilterOpts) (*IProverManagerNewPeriodIterator, error) {

	logs, sub, err := _IProverManager.contract.FilterLogs(opts, "NewPeriod")
	if err != nil {
		return nil, err
	}
	return &IProverManagerNewPeriodIterator{contract: _IProverManager.contract, event: "NewPeriod", logs: logs, sub: sub}, nil
}

// WatchNewPeriod is a free log subscription operation binding the contract event 0x61a611267e7ed28f8a566b021b9ac3ccc3985343a31971a180d01a57f63f3380.
//
// Solidity: event NewPeriod(uint256 periodId)
func (_IProverManager *IProverManagerFilterer) WatchNewPeriod(opts *bind.WatchOpts, sink chan<- *IProverManagerNewPeriod) (event.Subscription, error) {

	logs, sub, err := _IProverManager.contract.WatchLogs(opts, "NewPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProverManagerNewPeriod)
				if err := _IProverManager.contract.UnpackLog(event, "NewPeriod", log); err != nil {
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

// ParseNewPeriod is a log parse operation binding the contract event 0x61a611267e7ed28f8a566b021b9ac3ccc3985343a31971a180d01a57f63f3380.
//
// Solidity: event NewPeriod(uint256 periodId)
func (_IProverManager *IProverManagerFilterer) ParseNewPeriod(log types.Log) (*IProverManagerNewPeriod, error) {
	event := new(IProverManagerNewPeriod)
	if err := _IProverManager.contract.UnpackLog(event, "NewPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProverManagerProverEvictedIterator is returned from FilterProverEvicted and is used to iterate over the raw logs and unpacked data for ProverEvicted events raised by the IProverManager contract.
type IProverManagerProverEvictedIterator struct {
	Event *IProverManagerProverEvicted // Event containing the contract specifics and raw log

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
func (it *IProverManagerProverEvictedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProverManagerProverEvicted)
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
		it.Event = new(IProverManagerProverEvicted)
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
func (it *IProverManagerProverEvictedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProverManagerProverEvictedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProverManagerProverEvicted represents a ProverEvicted event raised by the IProverManager contract.
type IProverManagerProverEvicted struct {
	Prover       common.Address
	Evictor      common.Address
	PeriodEnd    *big.Int
	LivenessBond *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProverEvicted is a free log retrieval operation binding the contract event 0xc533214a7b99f5a45801bfa3d169dfaded72dcb302e35975809d247ba05c58bb.
//
// Solidity: event ProverEvicted(address indexed prover, address indexed evictor, uint256 periodEnd, uint256 livenessBond)
func (_IProverManager *IProverManagerFilterer) FilterProverEvicted(opts *bind.FilterOpts, prover []common.Address, evictor []common.Address) (*IProverManagerProverEvictedIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var evictorRule []interface{}
	for _, evictorItem := range evictor {
		evictorRule = append(evictorRule, evictorItem)
	}

	logs, sub, err := _IProverManager.contract.FilterLogs(opts, "ProverEvicted", proverRule, evictorRule)
	if err != nil {
		return nil, err
	}
	return &IProverManagerProverEvictedIterator{contract: _IProverManager.contract, event: "ProverEvicted", logs: logs, sub: sub}, nil
}

// WatchProverEvicted is a free log subscription operation binding the contract event 0xc533214a7b99f5a45801bfa3d169dfaded72dcb302e35975809d247ba05c58bb.
//
// Solidity: event ProverEvicted(address indexed prover, address indexed evictor, uint256 periodEnd, uint256 livenessBond)
func (_IProverManager *IProverManagerFilterer) WatchProverEvicted(opts *bind.WatchOpts, sink chan<- *IProverManagerProverEvicted, prover []common.Address, evictor []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var evictorRule []interface{}
	for _, evictorItem := range evictor {
		evictorRule = append(evictorRule, evictorItem)
	}

	logs, sub, err := _IProverManager.contract.WatchLogs(opts, "ProverEvicted", proverRule, evictorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProverManagerProverEvicted)
				if err := _IProverManager.contract.UnpackLog(event, "ProverEvicted", log); err != nil {
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

// ParseProverEvicted is a log parse operation binding the contract event 0xc533214a7b99f5a45801bfa3d169dfaded72dcb302e35975809d247ba05c58bb.
//
// Solidity: event ProverEvicted(address indexed prover, address indexed evictor, uint256 periodEnd, uint256 livenessBond)
func (_IProverManager *IProverManagerFilterer) ParseProverEvicted(log types.Log) (*IProverManagerProverEvicted, error) {
	event := new(IProverManagerProverEvicted)
	if err := _IProverManager.contract.UnpackLog(event, "ProverEvicted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProverManagerProverExitedIterator is returned from FilterProverExited and is used to iterate over the raw logs and unpacked data for ProverExited events raised by the IProverManager contract.
type IProverManagerProverExitedIterator struct {
	Event *IProverManagerProverExited // Event containing the contract specifics and raw log

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
func (it *IProverManagerProverExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProverManagerProverExited)
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
		it.Event = new(IProverManagerProverExited)
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
func (it *IProverManagerProverExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProverManagerProverExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProverManagerProverExited represents a ProverExited event raised by the IProverManager contract.
type IProverManagerProverExited struct {
	Prover          common.Address
	PeriodEnd       *big.Int
	ProvingDeadline *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProverExited is a free log retrieval operation binding the contract event 0xbf851044e869817ee35a381b9c3aca782413defdf564b86c4543e08dfc0fb312.
//
// Solidity: event ProverExited(address indexed prover, uint256 periodEnd, uint256 provingDeadline)
func (_IProverManager *IProverManagerFilterer) FilterProverExited(opts *bind.FilterOpts, prover []common.Address) (*IProverManagerProverExitedIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IProverManager.contract.FilterLogs(opts, "ProverExited", proverRule)
	if err != nil {
		return nil, err
	}
	return &IProverManagerProverExitedIterator{contract: _IProverManager.contract, event: "ProverExited", logs: logs, sub: sub}, nil
}

// WatchProverExited is a free log subscription operation binding the contract event 0xbf851044e869817ee35a381b9c3aca782413defdf564b86c4543e08dfc0fb312.
//
// Solidity: event ProverExited(address indexed prover, uint256 periodEnd, uint256 provingDeadline)
func (_IProverManager *IProverManagerFilterer) WatchProverExited(opts *bind.WatchOpts, sink chan<- *IProverManagerProverExited, prover []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IProverManager.contract.WatchLogs(opts, "ProverExited", proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProverManagerProverExited)
				if err := _IProverManager.contract.UnpackLog(event, "ProverExited", log); err != nil {
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

// ParseProverExited is a log parse operation binding the contract event 0xbf851044e869817ee35a381b9c3aca782413defdf564b86c4543e08dfc0fb312.
//
// Solidity: event ProverExited(address indexed prover, uint256 periodEnd, uint256 provingDeadline)
func (_IProverManager *IProverManagerFilterer) ParseProverExited(log types.Log) (*IProverManagerProverExited, error) {
	event := new(IProverManagerProverExited)
	if err := _IProverManager.contract.UnpackLog(event, "ProverExited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProverManagerProverOfferIterator is returned from FilterProverOffer and is used to iterate over the raw logs and unpacked data for ProverOffer events raised by the IProverManager contract.
type IProverManagerProverOfferIterator struct {
	Event *IProverManagerProverOffer // Event containing the contract specifics and raw log

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
func (it *IProverManagerProverOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProverManagerProverOffer)
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
		it.Event = new(IProverManagerProverOffer)
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
func (it *IProverManagerProverOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProverManagerProverOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProverManagerProverOffer represents a ProverOffer event raised by the IProverManager contract.
type IProverManagerProverOffer struct {
	Prover   common.Address
	PeriodId *big.Int
	Fee      *big.Int
	Stake    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProverOffer is a free log retrieval operation binding the contract event 0xf45ff6059124bdda46a19f3767ce6fd2eea57c96302c9819423af738b9b31191.
//
// Solidity: event ProverOffer(address indexed prover, uint256 periodId, uint256 fee, uint256 stake)
func (_IProverManager *IProverManagerFilterer) FilterProverOffer(opts *bind.FilterOpts, prover []common.Address) (*IProverManagerProverOfferIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IProverManager.contract.FilterLogs(opts, "ProverOffer", proverRule)
	if err != nil {
		return nil, err
	}
	return &IProverManagerProverOfferIterator{contract: _IProverManager.contract, event: "ProverOffer", logs: logs, sub: sub}, nil
}

// WatchProverOffer is a free log subscription operation binding the contract event 0xf45ff6059124bdda46a19f3767ce6fd2eea57c96302c9819423af738b9b31191.
//
// Solidity: event ProverOffer(address indexed prover, uint256 periodId, uint256 fee, uint256 stake)
func (_IProverManager *IProverManagerFilterer) WatchProverOffer(opts *bind.WatchOpts, sink chan<- *IProverManagerProverOffer, prover []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IProverManager.contract.WatchLogs(opts, "ProverOffer", proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProverManagerProverOffer)
				if err := _IProverManager.contract.UnpackLog(event, "ProverOffer", log); err != nil {
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

// ParseProverOffer is a log parse operation binding the contract event 0xf45ff6059124bdda46a19f3767ce6fd2eea57c96302c9819423af738b9b31191.
//
// Solidity: event ProverOffer(address indexed prover, uint256 periodId, uint256 fee, uint256 stake)
func (_IProverManager *IProverManagerFilterer) ParseProverOffer(log types.Log) (*IProverManagerProverOffer, error) {
	event := new(IProverManagerProverOffer)
	if err := _IProverManager.contract.UnpackLog(event, "ProverOffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProverManagerProverVacancyClaimedIterator is returned from FilterProverVacancyClaimed and is used to iterate over the raw logs and unpacked data for ProverVacancyClaimed events raised by the IProverManager contract.
type IProverManagerProverVacancyClaimedIterator struct {
	Event *IProverManagerProverVacancyClaimed // Event containing the contract specifics and raw log

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
func (it *IProverManagerProverVacancyClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProverManagerProverVacancyClaimed)
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
		it.Event = new(IProverManagerProverVacancyClaimed)
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
func (it *IProverManagerProverVacancyClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProverManagerProverVacancyClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProverManagerProverVacancyClaimed represents a ProverVacancyClaimed event raised by the IProverManager contract.
type IProverManagerProverVacancyClaimed struct {
	Prover   common.Address
	PeriodId *big.Int
	Fee      *big.Int
	Stake    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProverVacancyClaimed is a free log retrieval operation binding the contract event 0xf37ada07dc97ef546a7a49418ad56d0479a1a5a8068580ec5948fb232f112f35.
//
// Solidity: event ProverVacancyClaimed(address indexed prover, uint256 periodId, uint256 fee, uint256 stake)
func (_IProverManager *IProverManagerFilterer) FilterProverVacancyClaimed(opts *bind.FilterOpts, prover []common.Address) (*IProverManagerProverVacancyClaimedIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IProverManager.contract.FilterLogs(opts, "ProverVacancyClaimed", proverRule)
	if err != nil {
		return nil, err
	}
	return &IProverManagerProverVacancyClaimedIterator{contract: _IProverManager.contract, event: "ProverVacancyClaimed", logs: logs, sub: sub}, nil
}

// WatchProverVacancyClaimed is a free log subscription operation binding the contract event 0xf37ada07dc97ef546a7a49418ad56d0479a1a5a8068580ec5948fb232f112f35.
//
// Solidity: event ProverVacancyClaimed(address indexed prover, uint256 periodId, uint256 fee, uint256 stake)
func (_IProverManager *IProverManagerFilterer) WatchProverVacancyClaimed(opts *bind.WatchOpts, sink chan<- *IProverManagerProverVacancyClaimed, prover []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IProverManager.contract.WatchLogs(opts, "ProverVacancyClaimed", proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProverManagerProverVacancyClaimed)
				if err := _IProverManager.contract.UnpackLog(event, "ProverVacancyClaimed", log); err != nil {
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

// ParseProverVacancyClaimed is a log parse operation binding the contract event 0xf37ada07dc97ef546a7a49418ad56d0479a1a5a8068580ec5948fb232f112f35.
//
// Solidity: event ProverVacancyClaimed(address indexed prover, uint256 periodId, uint256 fee, uint256 stake)
func (_IProverManager *IProverManagerFilterer) ParseProverVacancyClaimed(log types.Log) (*IProverManagerProverVacancyClaimed, error) {
	event := new(IProverManagerProverVacancyClaimed)
	if err := _IProverManager.contract.UnpackLog(event, "ProverVacancyClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
