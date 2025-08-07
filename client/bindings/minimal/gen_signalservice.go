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

// ISignalServiceMetaData contains all meta data concerning the ISignalService contract.
var ISignalServiceMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"isSignalStored\",\"inputs\":[{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sendSignal\",\"inputs\":[{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifySignal\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitmentPublisher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"SignalSent\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignalVerified\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CommitmentNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCommitment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StorageRootCommitmentNotSupported\",\"inputs\":[]}]",
}

// ISignalServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use ISignalServiceMetaData.ABI instead.
var ISignalServiceABI = ISignalServiceMetaData.ABI

// ISignalService is an auto generated Go binding around an Ethereum contract.
type ISignalService struct {
	ISignalServiceCaller     // Read-only binding to the contract
	ISignalServiceTransactor // Write-only binding to the contract
	ISignalServiceFilterer   // Log filterer for contract events
}

// ISignalServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISignalServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignalServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISignalServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignalServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISignalServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignalServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISignalServiceSession struct {
	Contract     *ISignalService   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISignalServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISignalServiceCallerSession struct {
	Contract *ISignalServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ISignalServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISignalServiceTransactorSession struct {
	Contract     *ISignalServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ISignalServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISignalServiceRaw struct {
	Contract *ISignalService // Generic contract binding to access the raw methods on
}

// ISignalServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISignalServiceCallerRaw struct {
	Contract *ISignalServiceCaller // Generic read-only contract binding to access the raw methods on
}

// ISignalServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISignalServiceTransactorRaw struct {
	Contract *ISignalServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISignalService creates a new instance of ISignalService, bound to a specific deployed contract.
func NewISignalService(address common.Address, backend bind.ContractBackend) (*ISignalService, error) {
	contract, err := bindISignalService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISignalService{ISignalServiceCaller: ISignalServiceCaller{contract: contract}, ISignalServiceTransactor: ISignalServiceTransactor{contract: contract}, ISignalServiceFilterer: ISignalServiceFilterer{contract: contract}}, nil
}

// NewISignalServiceCaller creates a new read-only instance of ISignalService, bound to a specific deployed contract.
func NewISignalServiceCaller(address common.Address, caller bind.ContractCaller) (*ISignalServiceCaller, error) {
	contract, err := bindISignalService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISignalServiceCaller{contract: contract}, nil
}

// NewISignalServiceTransactor creates a new write-only instance of ISignalService, bound to a specific deployed contract.
func NewISignalServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*ISignalServiceTransactor, error) {
	contract, err := bindISignalService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISignalServiceTransactor{contract: contract}, nil
}

// NewISignalServiceFilterer creates a new log filterer instance of ISignalService, bound to a specific deployed contract.
func NewISignalServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*ISignalServiceFilterer, error) {
	contract, err := bindISignalService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISignalServiceFilterer{contract: contract}, nil
}

// bindISignalService binds a generic wrapper to an already deployed contract.
func bindISignalService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISignalServiceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISignalService *ISignalServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISignalService.Contract.ISignalServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISignalService *ISignalServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISignalService.Contract.ISignalServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISignalService *ISignalServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISignalService.Contract.ISignalServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISignalService *ISignalServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISignalService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISignalService *ISignalServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISignalService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISignalService *ISignalServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISignalService.Contract.contract.Transact(opts, method, params...)
}

// IsSignalStored is a free data retrieval call binding the contract method 0xcc753ae5.
//
// Solidity: function isSignalStored(bytes32 value, address sender) view returns(bool)
func (_ISignalService *ISignalServiceCaller) IsSignalStored(opts *bind.CallOpts, value [32]byte, sender common.Address) (bool, error) {
	var out []interface{}
	err := _ISignalService.contract.Call(opts, &out, "isSignalStored", value, sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSignalStored is a free data retrieval call binding the contract method 0xcc753ae5.
//
// Solidity: function isSignalStored(bytes32 value, address sender) view returns(bool)
func (_ISignalService *ISignalServiceSession) IsSignalStored(value [32]byte, sender common.Address) (bool, error) {
	return _ISignalService.Contract.IsSignalStored(&_ISignalService.CallOpts, value, sender)
}

// IsSignalStored is a free data retrieval call binding the contract method 0xcc753ae5.
//
// Solidity: function isSignalStored(bytes32 value, address sender) view returns(bool)
func (_ISignalService *ISignalServiceCallerSession) IsSignalStored(value [32]byte, sender common.Address) (bool, error) {
	return _ISignalService.Contract.IsSignalStored(&_ISignalService.CallOpts, value, sender)
}

// SendSignal is a paid mutator transaction binding the contract method 0x66ca2bc0.
//
// Solidity: function sendSignal(bytes32 value) returns(bytes32 slot)
func (_ISignalService *ISignalServiceTransactor) SendSignal(opts *bind.TransactOpts, value [32]byte) (*types.Transaction, error) {
	return _ISignalService.contract.Transact(opts, "sendSignal", value)
}

// SendSignal is a paid mutator transaction binding the contract method 0x66ca2bc0.
//
// Solidity: function sendSignal(bytes32 value) returns(bytes32 slot)
func (_ISignalService *ISignalServiceSession) SendSignal(value [32]byte) (*types.Transaction, error) {
	return _ISignalService.Contract.SendSignal(&_ISignalService.TransactOpts, value)
}

// SendSignal is a paid mutator transaction binding the contract method 0x66ca2bc0.
//
// Solidity: function sendSignal(bytes32 value) returns(bytes32 slot)
func (_ISignalService *ISignalServiceTransactorSession) SendSignal(value [32]byte) (*types.Transaction, error) {
	return _ISignalService.Contract.SendSignal(&_ISignalService.TransactOpts, value)
}

// VerifySignal is a paid mutator transaction binding the contract method 0xa44c8ddd.
//
// Solidity: function verifySignal(uint256 height, address commitmentPublisher, address sender, bytes32 value, bytes proof) returns()
func (_ISignalService *ISignalServiceTransactor) VerifySignal(opts *bind.TransactOpts, height *big.Int, commitmentPublisher common.Address, sender common.Address, value [32]byte, proof []byte) (*types.Transaction, error) {
	return _ISignalService.contract.Transact(opts, "verifySignal", height, commitmentPublisher, sender, value, proof)
}

// VerifySignal is a paid mutator transaction binding the contract method 0xa44c8ddd.
//
// Solidity: function verifySignal(uint256 height, address commitmentPublisher, address sender, bytes32 value, bytes proof) returns()
func (_ISignalService *ISignalServiceSession) VerifySignal(height *big.Int, commitmentPublisher common.Address, sender common.Address, value [32]byte, proof []byte) (*types.Transaction, error) {
	return _ISignalService.Contract.VerifySignal(&_ISignalService.TransactOpts, height, commitmentPublisher, sender, value, proof)
}

// VerifySignal is a paid mutator transaction binding the contract method 0xa44c8ddd.
//
// Solidity: function verifySignal(uint256 height, address commitmentPublisher, address sender, bytes32 value, bytes proof) returns()
func (_ISignalService *ISignalServiceTransactorSession) VerifySignal(height *big.Int, commitmentPublisher common.Address, sender common.Address, value [32]byte, proof []byte) (*types.Transaction, error) {
	return _ISignalService.Contract.VerifySignal(&_ISignalService.TransactOpts, height, commitmentPublisher, sender, value, proof)
}

// ISignalServiceSignalSentIterator is returned from FilterSignalSent and is used to iterate over the raw logs and unpacked data for SignalSent events raised by the ISignalService contract.
type ISignalServiceSignalSentIterator struct {
	Event *ISignalServiceSignalSent // Event containing the contract specifics and raw log

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
func (it *ISignalServiceSignalSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISignalServiceSignalSent)
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
		it.Event = new(ISignalServiceSignalSent)
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
func (it *ISignalServiceSignalSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISignalServiceSignalSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISignalServiceSignalSent represents a SignalSent event raised by the ISignalService contract.
type ISignalServiceSignalSent struct {
	Sender common.Address
	Value  [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignalSent is a free log retrieval operation binding the contract event 0xf0958489d2a32db2b0faf27a72a31fdf28f2636ca5532f1b077ddc2f51b20d1d.
//
// Solidity: event SignalSent(address indexed sender, bytes32 value)
func (_ISignalService *ISignalServiceFilterer) FilterSignalSent(opts *bind.FilterOpts, sender []common.Address) (*ISignalServiceSignalSentIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ISignalService.contract.FilterLogs(opts, "SignalSent", senderRule)
	if err != nil {
		return nil, err
	}
	return &ISignalServiceSignalSentIterator{contract: _ISignalService.contract, event: "SignalSent", logs: logs, sub: sub}, nil
}

// WatchSignalSent is a free log subscription operation binding the contract event 0xf0958489d2a32db2b0faf27a72a31fdf28f2636ca5532f1b077ddc2f51b20d1d.
//
// Solidity: event SignalSent(address indexed sender, bytes32 value)
func (_ISignalService *ISignalServiceFilterer) WatchSignalSent(opts *bind.WatchOpts, sink chan<- *ISignalServiceSignalSent, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ISignalService.contract.WatchLogs(opts, "SignalSent", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISignalServiceSignalSent)
				if err := _ISignalService.contract.UnpackLog(event, "SignalSent", log); err != nil {
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

// ParseSignalSent is a log parse operation binding the contract event 0xf0958489d2a32db2b0faf27a72a31fdf28f2636ca5532f1b077ddc2f51b20d1d.
//
// Solidity: event SignalSent(address indexed sender, bytes32 value)
func (_ISignalService *ISignalServiceFilterer) ParseSignalSent(log types.Log) (*ISignalServiceSignalSent, error) {
	event := new(ISignalServiceSignalSent)
	if err := _ISignalService.contract.UnpackLog(event, "SignalSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISignalServiceSignalVerifiedIterator is returned from FilterSignalVerified and is used to iterate over the raw logs and unpacked data for SignalVerified events raised by the ISignalService contract.
type ISignalServiceSignalVerifiedIterator struct {
	Event *ISignalServiceSignalVerified // Event containing the contract specifics and raw log

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
func (it *ISignalServiceSignalVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISignalServiceSignalVerified)
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
		it.Event = new(ISignalServiceSignalVerified)
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
func (it *ISignalServiceSignalVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISignalServiceSignalVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISignalServiceSignalVerified represents a SignalVerified event raised by the ISignalService contract.
type ISignalServiceSignalVerified struct {
	Sender common.Address
	Value  [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignalVerified is a free log retrieval operation binding the contract event 0x3f7829d8adbe0b40c9345225ce3100f599fe4143b4de8053f8202c773a917532.
//
// Solidity: event SignalVerified(address indexed sender, bytes32 value)
func (_ISignalService *ISignalServiceFilterer) FilterSignalVerified(opts *bind.FilterOpts, sender []common.Address) (*ISignalServiceSignalVerifiedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ISignalService.contract.FilterLogs(opts, "SignalVerified", senderRule)
	if err != nil {
		return nil, err
	}
	return &ISignalServiceSignalVerifiedIterator{contract: _ISignalService.contract, event: "SignalVerified", logs: logs, sub: sub}, nil
}

// WatchSignalVerified is a free log subscription operation binding the contract event 0x3f7829d8adbe0b40c9345225ce3100f599fe4143b4de8053f8202c773a917532.
//
// Solidity: event SignalVerified(address indexed sender, bytes32 value)
func (_ISignalService *ISignalServiceFilterer) WatchSignalVerified(opts *bind.WatchOpts, sink chan<- *ISignalServiceSignalVerified, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ISignalService.contract.WatchLogs(opts, "SignalVerified", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISignalServiceSignalVerified)
				if err := _ISignalService.contract.UnpackLog(event, "SignalVerified", log); err != nil {
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

// ParseSignalVerified is a log parse operation binding the contract event 0x3f7829d8adbe0b40c9345225ce3100f599fe4143b4de8053f8202c773a917532.
//
// Solidity: event SignalVerified(address indexed sender, bytes32 value)
func (_ISignalService *ISignalServiceFilterer) ParseSignalVerified(log types.Log) (*ISignalServiceSignalVerified, error) {
	event := new(ISignalServiceSignalVerified)
	if err := _ISignalService.contract.UnpackLog(event, "SignalVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
