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

// IBlobRefRegistryBlobRef is an auto generated low-level Go binding around an user-defined struct.
type IBlobRefRegistryBlobRef struct {
	BlockNumber *big.Int
	Blobhashes  [][32]byte
}

// IBlobRefRegistryMetaData contains all meta data concerning the IBlobRefRegistry contract.
var IBlobRefRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getRef\",\"inputs\":[{\"name\":\"blobIndices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBlobRefRegistry.BlobRef\",\"components\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blobhashes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRefRegistered\",\"inputs\":[{\"name\":\"refHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerRef\",\"inputs\":[{\"name\":\"blobIndices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"refHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ref\",\"type\":\"tuple\",\"internalType\":\"structIBlobRefRegistry.BlobRef\",\"components\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blobhashes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Registered\",\"inputs\":[{\"name\":\"refHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ref\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIBlobRefRegistry.BlobRef\",\"components\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blobhashes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"anonymous\":false}]",
}

// IBlobRefRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IBlobRefRegistryMetaData.ABI instead.
var IBlobRefRegistryABI = IBlobRefRegistryMetaData.ABI

// IBlobRefRegistry is an auto generated Go binding around an Ethereum contract.
type IBlobRefRegistry struct {
	IBlobRefRegistryCaller     // Read-only binding to the contract
	IBlobRefRegistryTransactor // Write-only binding to the contract
	IBlobRefRegistryFilterer   // Log filterer for contract events
}

// IBlobRefRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBlobRefRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlobRefRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBlobRefRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlobRefRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBlobRefRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlobRefRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBlobRefRegistrySession struct {
	Contract     *IBlobRefRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBlobRefRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBlobRefRegistryCallerSession struct {
	Contract *IBlobRefRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IBlobRefRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBlobRefRegistryTransactorSession struct {
	Contract     *IBlobRefRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IBlobRefRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBlobRefRegistryRaw struct {
	Contract *IBlobRefRegistry // Generic contract binding to access the raw methods on
}

// IBlobRefRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBlobRefRegistryCallerRaw struct {
	Contract *IBlobRefRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IBlobRefRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBlobRefRegistryTransactorRaw struct {
	Contract *IBlobRefRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBlobRefRegistry creates a new instance of IBlobRefRegistry, bound to a specific deployed contract.
func NewIBlobRefRegistry(address common.Address, backend bind.ContractBackend) (*IBlobRefRegistry, error) {
	contract, err := bindIBlobRefRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBlobRefRegistry{IBlobRefRegistryCaller: IBlobRefRegistryCaller{contract: contract}, IBlobRefRegistryTransactor: IBlobRefRegistryTransactor{contract: contract}, IBlobRefRegistryFilterer: IBlobRefRegistryFilterer{contract: contract}}, nil
}

// NewIBlobRefRegistryCaller creates a new read-only instance of IBlobRefRegistry, bound to a specific deployed contract.
func NewIBlobRefRegistryCaller(address common.Address, caller bind.ContractCaller) (*IBlobRefRegistryCaller, error) {
	contract, err := bindIBlobRefRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBlobRefRegistryCaller{contract: contract}, nil
}

// NewIBlobRefRegistryTransactor creates a new write-only instance of IBlobRefRegistry, bound to a specific deployed contract.
func NewIBlobRefRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IBlobRefRegistryTransactor, error) {
	contract, err := bindIBlobRefRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBlobRefRegistryTransactor{contract: contract}, nil
}

// NewIBlobRefRegistryFilterer creates a new log filterer instance of IBlobRefRegistry, bound to a specific deployed contract.
func NewIBlobRefRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IBlobRefRegistryFilterer, error) {
	contract, err := bindIBlobRefRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBlobRefRegistryFilterer{contract: contract}, nil
}

// bindIBlobRefRegistry binds a generic wrapper to an already deployed contract.
func bindIBlobRefRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBlobRefRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlobRefRegistry *IBlobRefRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBlobRefRegistry.Contract.IBlobRefRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlobRefRegistry *IBlobRefRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlobRefRegistry.Contract.IBlobRefRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlobRefRegistry *IBlobRefRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlobRefRegistry.Contract.IBlobRefRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlobRefRegistry *IBlobRefRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBlobRefRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlobRefRegistry *IBlobRefRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlobRefRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlobRefRegistry *IBlobRefRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlobRefRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetRef is a free data retrieval call binding the contract method 0xfe56ee8d.
//
// Solidity: function getRef(uint256[] blobIndices) view returns((uint256,bytes32[]))
func (_IBlobRefRegistry *IBlobRefRegistryCaller) GetRef(opts *bind.CallOpts, blobIndices []*big.Int) (IBlobRefRegistryBlobRef, error) {
	var out []interface{}
	err := _IBlobRefRegistry.contract.Call(opts, &out, "getRef", blobIndices)

	if err != nil {
		return *new(IBlobRefRegistryBlobRef), err
	}

	out0 := *abi.ConvertType(out[0], new(IBlobRefRegistryBlobRef)).(*IBlobRefRegistryBlobRef)

	return out0, err

}

// GetRef is a free data retrieval call binding the contract method 0xfe56ee8d.
//
// Solidity: function getRef(uint256[] blobIndices) view returns((uint256,bytes32[]))
func (_IBlobRefRegistry *IBlobRefRegistrySession) GetRef(blobIndices []*big.Int) (IBlobRefRegistryBlobRef, error) {
	return _IBlobRefRegistry.Contract.GetRef(&_IBlobRefRegistry.CallOpts, blobIndices)
}

// GetRef is a free data retrieval call binding the contract method 0xfe56ee8d.
//
// Solidity: function getRef(uint256[] blobIndices) view returns((uint256,bytes32[]))
func (_IBlobRefRegistry *IBlobRefRegistryCallerSession) GetRef(blobIndices []*big.Int) (IBlobRefRegistryBlobRef, error) {
	return _IBlobRefRegistry.Contract.GetRef(&_IBlobRefRegistry.CallOpts, blobIndices)
}

// IsRefRegistered is a free data retrieval call binding the contract method 0x4f202962.
//
// Solidity: function isRefRegistered(bytes32 refHash) view returns(bool)
func (_IBlobRefRegistry *IBlobRefRegistryCaller) IsRefRegistered(opts *bind.CallOpts, refHash [32]byte) (bool, error) {
	var out []interface{}
	err := _IBlobRefRegistry.contract.Call(opts, &out, "isRefRegistered", refHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRefRegistered is a free data retrieval call binding the contract method 0x4f202962.
//
// Solidity: function isRefRegistered(bytes32 refHash) view returns(bool)
func (_IBlobRefRegistry *IBlobRefRegistrySession) IsRefRegistered(refHash [32]byte) (bool, error) {
	return _IBlobRefRegistry.Contract.IsRefRegistered(&_IBlobRefRegistry.CallOpts, refHash)
}

// IsRefRegistered is a free data retrieval call binding the contract method 0x4f202962.
//
// Solidity: function isRefRegistered(bytes32 refHash) view returns(bool)
func (_IBlobRefRegistry *IBlobRefRegistryCallerSession) IsRefRegistered(refHash [32]byte) (bool, error) {
	return _IBlobRefRegistry.Contract.IsRefRegistered(&_IBlobRefRegistry.CallOpts, refHash)
}

// RegisterRef is a paid mutator transaction binding the contract method 0x0fe0bf4f.
//
// Solidity: function registerRef(uint256[] blobIndices) returns(bytes32 refHash, (uint256,bytes32[]) ref)
func (_IBlobRefRegistry *IBlobRefRegistryTransactor) RegisterRef(opts *bind.TransactOpts, blobIndices []*big.Int) (*types.Transaction, error) {
	return _IBlobRefRegistry.contract.Transact(opts, "registerRef", blobIndices)
}

// RegisterRef is a paid mutator transaction binding the contract method 0x0fe0bf4f.
//
// Solidity: function registerRef(uint256[] blobIndices) returns(bytes32 refHash, (uint256,bytes32[]) ref)
func (_IBlobRefRegistry *IBlobRefRegistrySession) RegisterRef(blobIndices []*big.Int) (*types.Transaction, error) {
	return _IBlobRefRegistry.Contract.RegisterRef(&_IBlobRefRegistry.TransactOpts, blobIndices)
}

// RegisterRef is a paid mutator transaction binding the contract method 0x0fe0bf4f.
//
// Solidity: function registerRef(uint256[] blobIndices) returns(bytes32 refHash, (uint256,bytes32[]) ref)
func (_IBlobRefRegistry *IBlobRefRegistryTransactorSession) RegisterRef(blobIndices []*big.Int) (*types.Transaction, error) {
	return _IBlobRefRegistry.Contract.RegisterRef(&_IBlobRefRegistry.TransactOpts, blobIndices)
}

// IBlobRefRegistryRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the IBlobRefRegistry contract.
type IBlobRefRegistryRegisteredIterator struct {
	Event *IBlobRefRegistryRegistered // Event containing the contract specifics and raw log

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
func (it *IBlobRefRegistryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlobRefRegistryRegistered)
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
		it.Event = new(IBlobRefRegistryRegistered)
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
func (it *IBlobRefRegistryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlobRefRegistryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlobRefRegistryRegistered represents a Registered event raised by the IBlobRefRegistry contract.
type IBlobRefRegistryRegistered struct {
	RefHash [32]byte
	Ref     IBlobRefRegistryBlobRef
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0xed198782f987144c59354d8218a3d68b5bc3c98d665723a6932495218ccba366.
//
// Solidity: event Registered(bytes32 indexed refHash, (uint256,bytes32[]) ref)
func (_IBlobRefRegistry *IBlobRefRegistryFilterer) FilterRegistered(opts *bind.FilterOpts, refHash [][32]byte) (*IBlobRefRegistryRegisteredIterator, error) {

	var refHashRule []interface{}
	for _, refHashItem := range refHash {
		refHashRule = append(refHashRule, refHashItem)
	}

	logs, sub, err := _IBlobRefRegistry.contract.FilterLogs(opts, "Registered", refHashRule)
	if err != nil {
		return nil, err
	}
	return &IBlobRefRegistryRegisteredIterator{contract: _IBlobRefRegistry.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0xed198782f987144c59354d8218a3d68b5bc3c98d665723a6932495218ccba366.
//
// Solidity: event Registered(bytes32 indexed refHash, (uint256,bytes32[]) ref)
func (_IBlobRefRegistry *IBlobRefRegistryFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *IBlobRefRegistryRegistered, refHash [][32]byte) (event.Subscription, error) {

	var refHashRule []interface{}
	for _, refHashItem := range refHash {
		refHashRule = append(refHashRule, refHashItem)
	}

	logs, sub, err := _IBlobRefRegistry.contract.WatchLogs(opts, "Registered", refHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlobRefRegistryRegistered)
				if err := _IBlobRefRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0xed198782f987144c59354d8218a3d68b5bc3c98d665723a6932495218ccba366.
//
// Solidity: event Registered(bytes32 indexed refHash, (uint256,bytes32[]) ref)
func (_IBlobRefRegistry *IBlobRefRegistryFilterer) ParseRegistered(log types.Log) (*IBlobRefRegistryRegistered, error) {
	event := new(IBlobRefRegistryRegistered)
	if err := _IBlobRefRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
