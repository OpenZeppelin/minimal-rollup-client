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

// BlobRefRegistryMetaData contains all meta data concerning the BlobRefRegistry contract.
var BlobRefRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getRef\",\"inputs\":[{\"name\":\"blobIndices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBlobRefRegistry.BlobRef\",\"components\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blobhashes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRefRegistered\",\"inputs\":[{\"name\":\"refHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerRef\",\"inputs\":[{\"name\":\"blobIndices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"refHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ref\",\"type\":\"tuple\",\"internalType\":\"structIBlobRefRegistry.BlobRef\",\"components\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blobhashes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Registered\",\"inputs\":[{\"name\":\"refHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ref\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIBlobRefRegistry.BlobRef\",\"components\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blobhashes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"anonymous\":false}]",
}

// BlobRefRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use BlobRefRegistryMetaData.ABI instead.
var BlobRefRegistryABI = BlobRefRegistryMetaData.ABI

// BlobRefRegistry is an auto generated Go binding around an Ethereum contract.
type BlobRefRegistry struct {
	BlobRefRegistryCaller     // Read-only binding to the contract
	BlobRefRegistryTransactor // Write-only binding to the contract
	BlobRefRegistryFilterer   // Log filterer for contract events
}

// BlobRefRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlobRefRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlobRefRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlobRefRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlobRefRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlobRefRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlobRefRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlobRefRegistrySession struct {
	Contract     *BlobRefRegistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlobRefRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlobRefRegistryCallerSession struct {
	Contract *BlobRefRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BlobRefRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlobRefRegistryTransactorSession struct {
	Contract     *BlobRefRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BlobRefRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlobRefRegistryRaw struct {
	Contract *BlobRefRegistry // Generic contract binding to access the raw methods on
}

// BlobRefRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlobRefRegistryCallerRaw struct {
	Contract *BlobRefRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// BlobRefRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlobRefRegistryTransactorRaw struct {
	Contract *BlobRefRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlobRefRegistry creates a new instance of BlobRefRegistry, bound to a specific deployed contract.
func NewBlobRefRegistry(address common.Address, backend bind.ContractBackend) (*BlobRefRegistry, error) {
	contract, err := bindBlobRefRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlobRefRegistry{BlobRefRegistryCaller: BlobRefRegistryCaller{contract: contract}, BlobRefRegistryTransactor: BlobRefRegistryTransactor{contract: contract}, BlobRefRegistryFilterer: BlobRefRegistryFilterer{contract: contract}}, nil
}

// NewBlobRefRegistryCaller creates a new read-only instance of BlobRefRegistry, bound to a specific deployed contract.
func NewBlobRefRegistryCaller(address common.Address, caller bind.ContractCaller) (*BlobRefRegistryCaller, error) {
	contract, err := bindBlobRefRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlobRefRegistryCaller{contract: contract}, nil
}

// NewBlobRefRegistryTransactor creates a new write-only instance of BlobRefRegistry, bound to a specific deployed contract.
func NewBlobRefRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*BlobRefRegistryTransactor, error) {
	contract, err := bindBlobRefRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlobRefRegistryTransactor{contract: contract}, nil
}

// NewBlobRefRegistryFilterer creates a new log filterer instance of BlobRefRegistry, bound to a specific deployed contract.
func NewBlobRefRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*BlobRefRegistryFilterer, error) {
	contract, err := bindBlobRefRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlobRefRegistryFilterer{contract: contract}, nil
}

// bindBlobRefRegistry binds a generic wrapper to an already deployed contract.
func bindBlobRefRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlobRefRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlobRefRegistry *BlobRefRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlobRefRegistry.Contract.BlobRefRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlobRefRegistry *BlobRefRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlobRefRegistry.Contract.BlobRefRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlobRefRegistry *BlobRefRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlobRefRegistry.Contract.BlobRefRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlobRefRegistry *BlobRefRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlobRefRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlobRefRegistry *BlobRefRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlobRefRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlobRefRegistry *BlobRefRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlobRefRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetRef is a free data retrieval call binding the contract method 0xfe56ee8d.
//
// Solidity: function getRef(uint256[] blobIndices) view returns((uint256,bytes32[]))
func (_BlobRefRegistry *BlobRefRegistryCaller) GetRef(opts *bind.CallOpts, blobIndices []*big.Int) (IBlobRefRegistryBlobRef, error) {
	var out []interface{}
	err := _BlobRefRegistry.contract.Call(opts, &out, "getRef", blobIndices)

	if err != nil {
		return *new(IBlobRefRegistryBlobRef), err
	}

	out0 := *abi.ConvertType(out[0], new(IBlobRefRegistryBlobRef)).(*IBlobRefRegistryBlobRef)

	return out0, err

}

// GetRef is a free data retrieval call binding the contract method 0xfe56ee8d.
//
// Solidity: function getRef(uint256[] blobIndices) view returns((uint256,bytes32[]))
func (_BlobRefRegistry *BlobRefRegistrySession) GetRef(blobIndices []*big.Int) (IBlobRefRegistryBlobRef, error) {
	return _BlobRefRegistry.Contract.GetRef(&_BlobRefRegistry.CallOpts, blobIndices)
}

// GetRef is a free data retrieval call binding the contract method 0xfe56ee8d.
//
// Solidity: function getRef(uint256[] blobIndices) view returns((uint256,bytes32[]))
func (_BlobRefRegistry *BlobRefRegistryCallerSession) GetRef(blobIndices []*big.Int) (IBlobRefRegistryBlobRef, error) {
	return _BlobRefRegistry.Contract.GetRef(&_BlobRefRegistry.CallOpts, blobIndices)
}

// IsRefRegistered is a free data retrieval call binding the contract method 0x4f202962.
//
// Solidity: function isRefRegistered(bytes32 refHash) view returns(bool)
func (_BlobRefRegistry *BlobRefRegistryCaller) IsRefRegistered(opts *bind.CallOpts, refHash [32]byte) (bool, error) {
	var out []interface{}
	err := _BlobRefRegistry.contract.Call(opts, &out, "isRefRegistered", refHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRefRegistered is a free data retrieval call binding the contract method 0x4f202962.
//
// Solidity: function isRefRegistered(bytes32 refHash) view returns(bool)
func (_BlobRefRegistry *BlobRefRegistrySession) IsRefRegistered(refHash [32]byte) (bool, error) {
	return _BlobRefRegistry.Contract.IsRefRegistered(&_BlobRefRegistry.CallOpts, refHash)
}

// IsRefRegistered is a free data retrieval call binding the contract method 0x4f202962.
//
// Solidity: function isRefRegistered(bytes32 refHash) view returns(bool)
func (_BlobRefRegistry *BlobRefRegistryCallerSession) IsRefRegistered(refHash [32]byte) (bool, error) {
	return _BlobRefRegistry.Contract.IsRefRegistered(&_BlobRefRegistry.CallOpts, refHash)
}

// RegisterRef is a paid mutator transaction binding the contract method 0x0fe0bf4f.
//
// Solidity: function registerRef(uint256[] blobIndices) returns(bytes32 refHash, (uint256,bytes32[]) ref)
func (_BlobRefRegistry *BlobRefRegistryTransactor) RegisterRef(opts *bind.TransactOpts, blobIndices []*big.Int) (*types.Transaction, error) {
	return _BlobRefRegistry.contract.Transact(opts, "registerRef", blobIndices)
}

// RegisterRef is a paid mutator transaction binding the contract method 0x0fe0bf4f.
//
// Solidity: function registerRef(uint256[] blobIndices) returns(bytes32 refHash, (uint256,bytes32[]) ref)
func (_BlobRefRegistry *BlobRefRegistrySession) RegisterRef(blobIndices []*big.Int) (*types.Transaction, error) {
	return _BlobRefRegistry.Contract.RegisterRef(&_BlobRefRegistry.TransactOpts, blobIndices)
}

// RegisterRef is a paid mutator transaction binding the contract method 0x0fe0bf4f.
//
// Solidity: function registerRef(uint256[] blobIndices) returns(bytes32 refHash, (uint256,bytes32[]) ref)
func (_BlobRefRegistry *BlobRefRegistryTransactorSession) RegisterRef(blobIndices []*big.Int) (*types.Transaction, error) {
	return _BlobRefRegistry.Contract.RegisterRef(&_BlobRefRegistry.TransactOpts, blobIndices)
}

// BlobRefRegistryRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the BlobRefRegistry contract.
type BlobRefRegistryRegisteredIterator struct {
	Event *BlobRefRegistryRegistered // Event containing the contract specifics and raw log

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
func (it *BlobRefRegistryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlobRefRegistryRegistered)
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
		it.Event = new(BlobRefRegistryRegistered)
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
func (it *BlobRefRegistryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlobRefRegistryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlobRefRegistryRegistered represents a Registered event raised by the BlobRefRegistry contract.
type BlobRefRegistryRegistered struct {
	RefHash [32]byte
	Ref     IBlobRefRegistryBlobRef
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0xed198782f987144c59354d8218a3d68b5bc3c98d665723a6932495218ccba366.
//
// Solidity: event Registered(bytes32 indexed refHash, (uint256,bytes32[]) ref)
func (_BlobRefRegistry *BlobRefRegistryFilterer) FilterRegistered(opts *bind.FilterOpts, refHash [][32]byte) (*BlobRefRegistryRegisteredIterator, error) {

	var refHashRule []interface{}
	for _, refHashItem := range refHash {
		refHashRule = append(refHashRule, refHashItem)
	}

	logs, sub, err := _BlobRefRegistry.contract.FilterLogs(opts, "Registered", refHashRule)
	if err != nil {
		return nil, err
	}
	return &BlobRefRegistryRegisteredIterator{contract: _BlobRefRegistry.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0xed198782f987144c59354d8218a3d68b5bc3c98d665723a6932495218ccba366.
//
// Solidity: event Registered(bytes32 indexed refHash, (uint256,bytes32[]) ref)
func (_BlobRefRegistry *BlobRefRegistryFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *BlobRefRegistryRegistered, refHash [][32]byte) (event.Subscription, error) {

	var refHashRule []interface{}
	for _, refHashItem := range refHash {
		refHashRule = append(refHashRule, refHashItem)
	}

	logs, sub, err := _BlobRefRegistry.contract.WatchLogs(opts, "Registered", refHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlobRefRegistryRegistered)
				if err := _BlobRefRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
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
func (_BlobRefRegistry *BlobRefRegistryFilterer) ParseRegistered(log types.Log) (*BlobRefRegistryRegistered, error) {
	event := new(BlobRefRegistryRegistered)
	if err := _BlobRefRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
