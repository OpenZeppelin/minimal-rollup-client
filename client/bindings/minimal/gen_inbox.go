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

// IInboxPublicationHeader is an auto generated low-level Go binding around an user-defined struct.
type IInboxPublicationHeader struct {
	Id             *big.Int
	PrevHash       [32]byte
	Publisher      common.Address
	Timestamp      *big.Int
	BlockNumber    *big.Int
	AttributesHash [32]byte
}

// IInboxMetaData contains all meta data concerning the IInbox contract.
var IInboxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getNextPublicationId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPublicationHash\",\"inputs\":[{\"name\":\"idx\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publish\",\"inputs\":[{\"name\":\"nBlobs\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"anchorBlockId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validateHeader\",\"inputs\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structIInbox.PublicationHeader\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"prevHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"publisher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"attributesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Published\",\"inputs\":[{\"name\":\"pubHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"header\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIInbox.PublicationHeader\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"prevHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"publisher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"attributesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"attributes\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false}]",
}

// IInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use IInboxMetaData.ABI instead.
var IInboxABI = IInboxMetaData.ABI

// IInbox is an auto generated Go binding around an Ethereum contract.
type IInbox struct {
	IInboxCaller     // Read-only binding to the contract
	IInboxTransactor // Write-only binding to the contract
	IInboxFilterer   // Log filterer for contract events
}

// IInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInboxSession struct {
	Contract     *IInbox           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInboxCallerSession struct {
	Contract *IInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInboxTransactorSession struct {
	Contract     *IInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInboxRaw struct {
	Contract *IInbox // Generic contract binding to access the raw methods on
}

// IInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInboxCallerRaw struct {
	Contract *IInboxCaller // Generic read-only contract binding to access the raw methods on
}

// IInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInboxTransactorRaw struct {
	Contract *IInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInbox creates a new instance of IInbox, bound to a specific deployed contract.
func NewIInbox(address common.Address, backend bind.ContractBackend) (*IInbox, error) {
	contract, err := bindIInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInbox{IInboxCaller: IInboxCaller{contract: contract}, IInboxTransactor: IInboxTransactor{contract: contract}, IInboxFilterer: IInboxFilterer{contract: contract}}, nil
}

// NewIInboxCaller creates a new read-only instance of IInbox, bound to a specific deployed contract.
func NewIInboxCaller(address common.Address, caller bind.ContractCaller) (*IInboxCaller, error) {
	contract, err := bindIInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInboxCaller{contract: contract}, nil
}

// NewIInboxTransactor creates a new write-only instance of IInbox, bound to a specific deployed contract.
func NewIInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*IInboxTransactor, error) {
	contract, err := bindIInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInboxTransactor{contract: contract}, nil
}

// NewIInboxFilterer creates a new log filterer instance of IInbox, bound to a specific deployed contract.
func NewIInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*IInboxFilterer, error) {
	contract, err := bindIInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInboxFilterer{contract: contract}, nil
}

// bindIInbox binds a generic wrapper to an already deployed contract.
func bindIInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IInboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInbox *IInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInbox.Contract.IInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInbox *IInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInbox.Contract.IInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInbox *IInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInbox.Contract.IInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInbox *IInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInbox *IInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInbox *IInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInbox.Contract.contract.Transact(opts, method, params...)
}

// GetNextPublicationId is a free data retrieval call binding the contract method 0x416d86de.
//
// Solidity: function getNextPublicationId() view returns(uint256)
func (_IInbox *IInboxCaller) GetNextPublicationId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IInbox.contract.Call(opts, &out, "getNextPublicationId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextPublicationId is a free data retrieval call binding the contract method 0x416d86de.
//
// Solidity: function getNextPublicationId() view returns(uint256)
func (_IInbox *IInboxSession) GetNextPublicationId() (*big.Int, error) {
	return _IInbox.Contract.GetNextPublicationId(&_IInbox.CallOpts)
}

// GetNextPublicationId is a free data retrieval call binding the contract method 0x416d86de.
//
// Solidity: function getNextPublicationId() view returns(uint256)
func (_IInbox *IInboxCallerSession) GetNextPublicationId() (*big.Int, error) {
	return _IInbox.Contract.GetNextPublicationId(&_IInbox.CallOpts)
}

// GetPublicationHash is a free data retrieval call binding the contract method 0xef7922e3.
//
// Solidity: function getPublicationHash(uint256 idx) view returns(bytes32)
func (_IInbox *IInboxCaller) GetPublicationHash(opts *bind.CallOpts, idx *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IInbox.contract.Call(opts, &out, "getPublicationHash", idx)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPublicationHash is a free data retrieval call binding the contract method 0xef7922e3.
//
// Solidity: function getPublicationHash(uint256 idx) view returns(bytes32)
func (_IInbox *IInboxSession) GetPublicationHash(idx *big.Int) ([32]byte, error) {
	return _IInbox.Contract.GetPublicationHash(&_IInbox.CallOpts, idx)
}

// GetPublicationHash is a free data retrieval call binding the contract method 0xef7922e3.
//
// Solidity: function getPublicationHash(uint256 idx) view returns(bytes32)
func (_IInbox *IInboxCallerSession) GetPublicationHash(idx *big.Int) ([32]byte, error) {
	return _IInbox.Contract.GetPublicationHash(&_IInbox.CallOpts, idx)
}

// ValidateHeader is a free data retrieval call binding the contract method 0x62c0e745.
//
// Solidity: function validateHeader((uint256,bytes32,address,uint256,uint256,bytes32) header) view returns(bool)
func (_IInbox *IInboxCaller) ValidateHeader(opts *bind.CallOpts, header IInboxPublicationHeader) (bool, error) {
	var out []interface{}
	err := _IInbox.contract.Call(opts, &out, "validateHeader", header)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateHeader is a free data retrieval call binding the contract method 0x62c0e745.
//
// Solidity: function validateHeader((uint256,bytes32,address,uint256,uint256,bytes32) header) view returns(bool)
func (_IInbox *IInboxSession) ValidateHeader(header IInboxPublicationHeader) (bool, error) {
	return _IInbox.Contract.ValidateHeader(&_IInbox.CallOpts, header)
}

// ValidateHeader is a free data retrieval call binding the contract method 0x62c0e745.
//
// Solidity: function validateHeader((uint256,bytes32,address,uint256,uint256,bytes32) header) view returns(bool)
func (_IInbox *IInboxCallerSession) ValidateHeader(header IInboxPublicationHeader) (bool, error) {
	return _IInbox.Contract.ValidateHeader(&_IInbox.CallOpts, header)
}

// Publish is a paid mutator transaction binding the contract method 0x4039cc61.
//
// Solidity: function publish(uint256 nBlobs, uint64 anchorBlockId) returns()
func (_IInbox *IInboxTransactor) Publish(opts *bind.TransactOpts, nBlobs *big.Int, anchorBlockId uint64) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "publish", nBlobs, anchorBlockId)
}

// Publish is a paid mutator transaction binding the contract method 0x4039cc61.
//
// Solidity: function publish(uint256 nBlobs, uint64 anchorBlockId) returns()
func (_IInbox *IInboxSession) Publish(nBlobs *big.Int, anchorBlockId uint64) (*types.Transaction, error) {
	return _IInbox.Contract.Publish(&_IInbox.TransactOpts, nBlobs, anchorBlockId)
}

// Publish is a paid mutator transaction binding the contract method 0x4039cc61.
//
// Solidity: function publish(uint256 nBlobs, uint64 anchorBlockId) returns()
func (_IInbox *IInboxTransactorSession) Publish(nBlobs *big.Int, anchorBlockId uint64) (*types.Transaction, error) {
	return _IInbox.Contract.Publish(&_IInbox.TransactOpts, nBlobs, anchorBlockId)
}

// IInboxPublishedIterator is returned from FilterPublished and is used to iterate over the raw logs and unpacked data for Published events raised by the IInbox contract.
type IInboxPublishedIterator struct {
	Event *IInboxPublished // Event containing the contract specifics and raw log

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
func (it *IInboxPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IInboxPublished)
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
		it.Event = new(IInboxPublished)
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
func (it *IInboxPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IInboxPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IInboxPublished represents a Published event raised by the IInbox contract.
type IInboxPublished struct {
	PubHash    [32]byte
	Header     IInboxPublicationHeader
	Attributes [][]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPublished is a free log retrieval operation binding the contract event 0x09a62876c6768f7c1c3233bafb09528ae9ff0fb0287ca35592ece85e856b31b2.
//
// Solidity: event Published(bytes32 indexed pubHash, (uint256,bytes32,address,uint256,uint256,bytes32) header, bytes[] attributes)
func (_IInbox *IInboxFilterer) FilterPublished(opts *bind.FilterOpts, pubHash [][32]byte) (*IInboxPublishedIterator, error) {

	var pubHashRule []interface{}
	for _, pubHashItem := range pubHash {
		pubHashRule = append(pubHashRule, pubHashItem)
	}

	logs, sub, err := _IInbox.contract.FilterLogs(opts, "Published", pubHashRule)
	if err != nil {
		return nil, err
	}
	return &IInboxPublishedIterator{contract: _IInbox.contract, event: "Published", logs: logs, sub: sub}, nil
}

// WatchPublished is a free log subscription operation binding the contract event 0x09a62876c6768f7c1c3233bafb09528ae9ff0fb0287ca35592ece85e856b31b2.
//
// Solidity: event Published(bytes32 indexed pubHash, (uint256,bytes32,address,uint256,uint256,bytes32) header, bytes[] attributes)
func (_IInbox *IInboxFilterer) WatchPublished(opts *bind.WatchOpts, sink chan<- *IInboxPublished, pubHash [][32]byte) (event.Subscription, error) {

	var pubHashRule []interface{}
	for _, pubHashItem := range pubHash {
		pubHashRule = append(pubHashRule, pubHashItem)
	}

	logs, sub, err := _IInbox.contract.WatchLogs(opts, "Published", pubHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IInboxPublished)
				if err := _IInbox.contract.UnpackLog(event, "Published", log); err != nil {
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

// ParsePublished is a log parse operation binding the contract event 0x09a62876c6768f7c1c3233bafb09528ae9ff0fb0287ca35592ece85e856b31b2.
//
// Solidity: event Published(bytes32 indexed pubHash, (uint256,bytes32,address,uint256,uint256,bytes32) header, bytes[] attributes)
func (_IInbox *IInboxFilterer) ParsePublished(log types.Log) (*IInboxPublished, error) {
	event := new(IInboxPublished)
	if err := _IInbox.contract.UnpackLog(event, "Published", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
