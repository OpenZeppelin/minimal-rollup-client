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

// CheckpointTrackerMetaData contains all meta data concerning the CheckpointTracker contract.
var CheckpointTrackerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_genesis\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_inbox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_verifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_commitmentStore\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"commitmentStore\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractICommitmentStore\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIInbox\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proveTransition\",\"inputs\":[{\"name\":\"start\",\"type\":\"tuple\",\"internalType\":\"structICheckpointTracker.Checkpoint\",\"components\":[{\"name\":\"publicationId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"totalDelayedPublications\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"end\",\"type\":\"tuple\",\"internalType\":\"structICheckpointTracker.Checkpoint\",\"components\":[{\"name\":\"publicationId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"totalDelayedPublications\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"numPublications\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numDelayedPublications\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provenPublicationId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proverManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateProverManager\",\"inputs\":[{\"name\":\"_proverManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CommitmentSaved\",\"inputs\":[{\"name\":\"publicationId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"commitment\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProverManagerUpdated\",\"inputs\":[{\"name\":\"proverManager\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EndPublicationNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExcessiveDelayedPublications\",\"inputs\":[{\"name\":\"delayedCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStartPublication\",\"inputs\":[{\"name\":\"startPublicationId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"latestProvenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlyProverManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroEndCommitment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroGenesisCommitment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroStartCommitment\",\"inputs\":[]}]",
}

// CheckpointTrackerABI is the input ABI used to generate the binding from.
// Deprecated: Use CheckpointTrackerMetaData.ABI instead.
var CheckpointTrackerABI = CheckpointTrackerMetaData.ABI

// CheckpointTracker is an auto generated Go binding around an Ethereum contract.
type CheckpointTracker struct {
	CheckpointTrackerCaller     // Read-only binding to the contract
	CheckpointTrackerTransactor // Write-only binding to the contract
	CheckpointTrackerFilterer   // Log filterer for contract events
}

// CheckpointTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CheckpointTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CheckpointTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CheckpointTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CheckpointTrackerSession struct {
	Contract     *CheckpointTracker // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CheckpointTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CheckpointTrackerCallerSession struct {
	Contract *CheckpointTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CheckpointTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CheckpointTrackerTransactorSession struct {
	Contract     *CheckpointTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CheckpointTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CheckpointTrackerRaw struct {
	Contract *CheckpointTracker // Generic contract binding to access the raw methods on
}

// CheckpointTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CheckpointTrackerCallerRaw struct {
	Contract *CheckpointTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// CheckpointTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CheckpointTrackerTransactorRaw struct {
	Contract *CheckpointTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCheckpointTracker creates a new instance of CheckpointTracker, bound to a specific deployed contract.
func NewCheckpointTracker(address common.Address, backend bind.ContractBackend) (*CheckpointTracker, error) {
	contract, err := bindCheckpointTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CheckpointTracker{CheckpointTrackerCaller: CheckpointTrackerCaller{contract: contract}, CheckpointTrackerTransactor: CheckpointTrackerTransactor{contract: contract}, CheckpointTrackerFilterer: CheckpointTrackerFilterer{contract: contract}}, nil
}

// NewCheckpointTrackerCaller creates a new read-only instance of CheckpointTracker, bound to a specific deployed contract.
func NewCheckpointTrackerCaller(address common.Address, caller bind.ContractCaller) (*CheckpointTrackerCaller, error) {
	contract, err := bindCheckpointTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CheckpointTrackerCaller{contract: contract}, nil
}

// NewCheckpointTrackerTransactor creates a new write-only instance of CheckpointTracker, bound to a specific deployed contract.
func NewCheckpointTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*CheckpointTrackerTransactor, error) {
	contract, err := bindCheckpointTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CheckpointTrackerTransactor{contract: contract}, nil
}

// NewCheckpointTrackerFilterer creates a new log filterer instance of CheckpointTracker, bound to a specific deployed contract.
func NewCheckpointTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*CheckpointTrackerFilterer, error) {
	contract, err := bindCheckpointTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CheckpointTrackerFilterer{contract: contract}, nil
}

// bindCheckpointTracker binds a generic wrapper to an already deployed contract.
func bindCheckpointTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CheckpointTrackerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CheckpointTracker *CheckpointTrackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CheckpointTracker.Contract.CheckpointTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CheckpointTracker *CheckpointTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CheckpointTracker.Contract.CheckpointTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CheckpointTracker *CheckpointTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CheckpointTracker.Contract.CheckpointTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CheckpointTracker *CheckpointTrackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CheckpointTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CheckpointTracker *CheckpointTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CheckpointTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CheckpointTracker *CheckpointTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CheckpointTracker.Contract.contract.Transact(opts, method, params...)
}

// CommitmentStore is a free data retrieval call binding the contract method 0xeee67125.
//
// Solidity: function commitmentStore() view returns(address)
func (_CheckpointTracker *CheckpointTrackerCaller) CommitmentStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CheckpointTracker.contract.Call(opts, &out, "commitmentStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CommitmentStore is a free data retrieval call binding the contract method 0xeee67125.
//
// Solidity: function commitmentStore() view returns(address)
func (_CheckpointTracker *CheckpointTrackerSession) CommitmentStore() (common.Address, error) {
	return _CheckpointTracker.Contract.CommitmentStore(&_CheckpointTracker.CallOpts)
}

// CommitmentStore is a free data retrieval call binding the contract method 0xeee67125.
//
// Solidity: function commitmentStore() view returns(address)
func (_CheckpointTracker *CheckpointTrackerCallerSession) CommitmentStore() (common.Address, error) {
	return _CheckpointTracker.Contract.CommitmentStore(&_CheckpointTracker.CallOpts)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_CheckpointTracker *CheckpointTrackerCaller) Inbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CheckpointTracker.contract.Call(opts, &out, "inbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_CheckpointTracker *CheckpointTrackerSession) Inbox() (common.Address, error) {
	return _CheckpointTracker.Contract.Inbox(&_CheckpointTracker.CallOpts)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_CheckpointTracker *CheckpointTrackerCallerSession) Inbox() (common.Address, error) {
	return _CheckpointTracker.Contract.Inbox(&_CheckpointTracker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CheckpointTracker *CheckpointTrackerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CheckpointTracker.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CheckpointTracker *CheckpointTrackerSession) Owner() (common.Address, error) {
	return _CheckpointTracker.Contract.Owner(&_CheckpointTracker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CheckpointTracker *CheckpointTrackerCallerSession) Owner() (common.Address, error) {
	return _CheckpointTracker.Contract.Owner(&_CheckpointTracker.CallOpts)
}

// ProvenPublicationId is a free data retrieval call binding the contract method 0x3f558a2a.
//
// Solidity: function provenPublicationId() view returns(uint256)
func (_CheckpointTracker *CheckpointTrackerCaller) ProvenPublicationId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CheckpointTracker.contract.Call(opts, &out, "provenPublicationId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProvenPublicationId is a free data retrieval call binding the contract method 0x3f558a2a.
//
// Solidity: function provenPublicationId() view returns(uint256)
func (_CheckpointTracker *CheckpointTrackerSession) ProvenPublicationId() (*big.Int, error) {
	return _CheckpointTracker.Contract.ProvenPublicationId(&_CheckpointTracker.CallOpts)
}

// ProvenPublicationId is a free data retrieval call binding the contract method 0x3f558a2a.
//
// Solidity: function provenPublicationId() view returns(uint256)
func (_CheckpointTracker *CheckpointTrackerCallerSession) ProvenPublicationId() (*big.Int, error) {
	return _CheckpointTracker.Contract.ProvenPublicationId(&_CheckpointTracker.CallOpts)
}

// ProverManager is a free data retrieval call binding the contract method 0x23f9b902.
//
// Solidity: function proverManager() view returns(address)
func (_CheckpointTracker *CheckpointTrackerCaller) ProverManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CheckpointTracker.contract.Call(opts, &out, "proverManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProverManager is a free data retrieval call binding the contract method 0x23f9b902.
//
// Solidity: function proverManager() view returns(address)
func (_CheckpointTracker *CheckpointTrackerSession) ProverManager() (common.Address, error) {
	return _CheckpointTracker.Contract.ProverManager(&_CheckpointTracker.CallOpts)
}

// ProverManager is a free data retrieval call binding the contract method 0x23f9b902.
//
// Solidity: function proverManager() view returns(address)
func (_CheckpointTracker *CheckpointTrackerCallerSession) ProverManager() (common.Address, error) {
	return _CheckpointTracker.Contract.ProverManager(&_CheckpointTracker.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_CheckpointTracker *CheckpointTrackerCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CheckpointTracker.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_CheckpointTracker *CheckpointTrackerSession) Verifier() (common.Address, error) {
	return _CheckpointTracker.Contract.Verifier(&_CheckpointTracker.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_CheckpointTracker *CheckpointTrackerCallerSession) Verifier() (common.Address, error) {
	return _CheckpointTracker.Contract.Verifier(&_CheckpointTracker.CallOpts)
}

// ProveTransition is a paid mutator transaction binding the contract method 0xc1ce1eb9.
//
// Solidity: function proveTransition((uint256,bytes32,uint256) start, (uint256,bytes32,uint256) end, bytes proof) returns(uint256 numPublications, uint256 numDelayedPublications)
func (_CheckpointTracker *CheckpointTrackerTransactor) ProveTransition(opts *bind.TransactOpts, start ICheckpointTrackerCheckpoint, end ICheckpointTrackerCheckpoint, proof []byte) (*types.Transaction, error) {
	return _CheckpointTracker.contract.Transact(opts, "proveTransition", start, end, proof)
}

// ProveTransition is a paid mutator transaction binding the contract method 0xc1ce1eb9.
//
// Solidity: function proveTransition((uint256,bytes32,uint256) start, (uint256,bytes32,uint256) end, bytes proof) returns(uint256 numPublications, uint256 numDelayedPublications)
func (_CheckpointTracker *CheckpointTrackerSession) ProveTransition(start ICheckpointTrackerCheckpoint, end ICheckpointTrackerCheckpoint, proof []byte) (*types.Transaction, error) {
	return _CheckpointTracker.Contract.ProveTransition(&_CheckpointTracker.TransactOpts, start, end, proof)
}

// ProveTransition is a paid mutator transaction binding the contract method 0xc1ce1eb9.
//
// Solidity: function proveTransition((uint256,bytes32,uint256) start, (uint256,bytes32,uint256) end, bytes proof) returns(uint256 numPublications, uint256 numDelayedPublications)
func (_CheckpointTracker *CheckpointTrackerTransactorSession) ProveTransition(start ICheckpointTrackerCheckpoint, end ICheckpointTrackerCheckpoint, proof []byte) (*types.Transaction, error) {
	return _CheckpointTracker.Contract.ProveTransition(&_CheckpointTracker.TransactOpts, start, end, proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CheckpointTracker *CheckpointTrackerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CheckpointTracker.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CheckpointTracker *CheckpointTrackerSession) RenounceOwnership() (*types.Transaction, error) {
	return _CheckpointTracker.Contract.RenounceOwnership(&_CheckpointTracker.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CheckpointTracker *CheckpointTrackerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CheckpointTracker.Contract.RenounceOwnership(&_CheckpointTracker.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CheckpointTracker *CheckpointTrackerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CheckpointTracker.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CheckpointTracker *CheckpointTrackerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CheckpointTracker.Contract.TransferOwnership(&_CheckpointTracker.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CheckpointTracker *CheckpointTrackerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CheckpointTracker.Contract.TransferOwnership(&_CheckpointTracker.TransactOpts, newOwner)
}

// UpdateProverManager is a paid mutator transaction binding the contract method 0xcaaa9a98.
//
// Solidity: function updateProverManager(address _proverManager) returns()
func (_CheckpointTracker *CheckpointTrackerTransactor) UpdateProverManager(opts *bind.TransactOpts, _proverManager common.Address) (*types.Transaction, error) {
	return _CheckpointTracker.contract.Transact(opts, "updateProverManager", _proverManager)
}

// UpdateProverManager is a paid mutator transaction binding the contract method 0xcaaa9a98.
//
// Solidity: function updateProverManager(address _proverManager) returns()
func (_CheckpointTracker *CheckpointTrackerSession) UpdateProverManager(_proverManager common.Address) (*types.Transaction, error) {
	return _CheckpointTracker.Contract.UpdateProverManager(&_CheckpointTracker.TransactOpts, _proverManager)
}

// UpdateProverManager is a paid mutator transaction binding the contract method 0xcaaa9a98.
//
// Solidity: function updateProverManager(address _proverManager) returns()
func (_CheckpointTracker *CheckpointTrackerTransactorSession) UpdateProverManager(_proverManager common.Address) (*types.Transaction, error) {
	return _CheckpointTracker.Contract.UpdateProverManager(&_CheckpointTracker.TransactOpts, _proverManager)
}

// CheckpointTrackerCommitmentSavedIterator is returned from FilterCommitmentSaved and is used to iterate over the raw logs and unpacked data for CommitmentSaved events raised by the CheckpointTracker contract.
type CheckpointTrackerCommitmentSavedIterator struct {
	Event *CheckpointTrackerCommitmentSaved // Event containing the contract specifics and raw log

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
func (it *CheckpointTrackerCommitmentSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CheckpointTrackerCommitmentSaved)
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
		it.Event = new(CheckpointTrackerCommitmentSaved)
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
func (it *CheckpointTrackerCommitmentSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CheckpointTrackerCommitmentSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CheckpointTrackerCommitmentSaved represents a CommitmentSaved event raised by the CheckpointTracker contract.
type CheckpointTrackerCommitmentSaved struct {
	PublicationId *big.Int
	Commitment    [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCommitmentSaved is a free log retrieval operation binding the contract event 0x42898ee6cff40a090090a9f2cf484511de9d1b1b735396cd2307646d2fb6f890.
//
// Solidity: event CommitmentSaved(uint256 indexed publicationId, bytes32 commitment)
func (_CheckpointTracker *CheckpointTrackerFilterer) FilterCommitmentSaved(opts *bind.FilterOpts, publicationId []*big.Int) (*CheckpointTrackerCommitmentSavedIterator, error) {

	var publicationIdRule []interface{}
	for _, publicationIdItem := range publicationId {
		publicationIdRule = append(publicationIdRule, publicationIdItem)
	}

	logs, sub, err := _CheckpointTracker.contract.FilterLogs(opts, "CommitmentSaved", publicationIdRule)
	if err != nil {
		return nil, err
	}
	return &CheckpointTrackerCommitmentSavedIterator{contract: _CheckpointTracker.contract, event: "CommitmentSaved", logs: logs, sub: sub}, nil
}

// WatchCommitmentSaved is a free log subscription operation binding the contract event 0x42898ee6cff40a090090a9f2cf484511de9d1b1b735396cd2307646d2fb6f890.
//
// Solidity: event CommitmentSaved(uint256 indexed publicationId, bytes32 commitment)
func (_CheckpointTracker *CheckpointTrackerFilterer) WatchCommitmentSaved(opts *bind.WatchOpts, sink chan<- *CheckpointTrackerCommitmentSaved, publicationId []*big.Int) (event.Subscription, error) {

	var publicationIdRule []interface{}
	for _, publicationIdItem := range publicationId {
		publicationIdRule = append(publicationIdRule, publicationIdItem)
	}

	logs, sub, err := _CheckpointTracker.contract.WatchLogs(opts, "CommitmentSaved", publicationIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CheckpointTrackerCommitmentSaved)
				if err := _CheckpointTracker.contract.UnpackLog(event, "CommitmentSaved", log); err != nil {
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
func (_CheckpointTracker *CheckpointTrackerFilterer) ParseCommitmentSaved(log types.Log) (*CheckpointTrackerCommitmentSaved, error) {
	event := new(CheckpointTrackerCommitmentSaved)
	if err := _CheckpointTracker.contract.UnpackLog(event, "CommitmentSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CheckpointTrackerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CheckpointTracker contract.
type CheckpointTrackerOwnershipTransferredIterator struct {
	Event *CheckpointTrackerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CheckpointTrackerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CheckpointTrackerOwnershipTransferred)
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
		it.Event = new(CheckpointTrackerOwnershipTransferred)
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
func (it *CheckpointTrackerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CheckpointTrackerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CheckpointTrackerOwnershipTransferred represents a OwnershipTransferred event raised by the CheckpointTracker contract.
type CheckpointTrackerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CheckpointTracker *CheckpointTrackerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CheckpointTrackerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CheckpointTracker.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CheckpointTrackerOwnershipTransferredIterator{contract: _CheckpointTracker.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CheckpointTracker *CheckpointTrackerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CheckpointTrackerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CheckpointTracker.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CheckpointTrackerOwnershipTransferred)
				if err := _CheckpointTracker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CheckpointTracker *CheckpointTrackerFilterer) ParseOwnershipTransferred(log types.Log) (*CheckpointTrackerOwnershipTransferred, error) {
	event := new(CheckpointTrackerOwnershipTransferred)
	if err := _CheckpointTracker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CheckpointTrackerProverManagerUpdatedIterator is returned from FilterProverManagerUpdated and is used to iterate over the raw logs and unpacked data for ProverManagerUpdated events raised by the CheckpointTracker contract.
type CheckpointTrackerProverManagerUpdatedIterator struct {
	Event *CheckpointTrackerProverManagerUpdated // Event containing the contract specifics and raw log

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
func (it *CheckpointTrackerProverManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CheckpointTrackerProverManagerUpdated)
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
		it.Event = new(CheckpointTrackerProverManagerUpdated)
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
func (it *CheckpointTrackerProverManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CheckpointTrackerProverManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CheckpointTrackerProverManagerUpdated represents a ProverManagerUpdated event raised by the CheckpointTracker contract.
type CheckpointTrackerProverManagerUpdated struct {
	ProverManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProverManagerUpdated is a free log retrieval operation binding the contract event 0x8ad478da7b4df3c373f46444bff2373d0bbacce5c8d7aaf53d925dfd2552253d.
//
// Solidity: event ProverManagerUpdated(address indexed proverManager)
func (_CheckpointTracker *CheckpointTrackerFilterer) FilterProverManagerUpdated(opts *bind.FilterOpts, proverManager []common.Address) (*CheckpointTrackerProverManagerUpdatedIterator, error) {

	var proverManagerRule []interface{}
	for _, proverManagerItem := range proverManager {
		proverManagerRule = append(proverManagerRule, proverManagerItem)
	}

	logs, sub, err := _CheckpointTracker.contract.FilterLogs(opts, "ProverManagerUpdated", proverManagerRule)
	if err != nil {
		return nil, err
	}
	return &CheckpointTrackerProverManagerUpdatedIterator{contract: _CheckpointTracker.contract, event: "ProverManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchProverManagerUpdated is a free log subscription operation binding the contract event 0x8ad478da7b4df3c373f46444bff2373d0bbacce5c8d7aaf53d925dfd2552253d.
//
// Solidity: event ProverManagerUpdated(address indexed proverManager)
func (_CheckpointTracker *CheckpointTrackerFilterer) WatchProverManagerUpdated(opts *bind.WatchOpts, sink chan<- *CheckpointTrackerProverManagerUpdated, proverManager []common.Address) (event.Subscription, error) {

	var proverManagerRule []interface{}
	for _, proverManagerItem := range proverManager {
		proverManagerRule = append(proverManagerRule, proverManagerItem)
	}

	logs, sub, err := _CheckpointTracker.contract.WatchLogs(opts, "ProverManagerUpdated", proverManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CheckpointTrackerProverManagerUpdated)
				if err := _CheckpointTracker.contract.UnpackLog(event, "ProverManagerUpdated", log); err != nil {
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

// ParseProverManagerUpdated is a log parse operation binding the contract event 0x8ad478da7b4df3c373f46444bff2373d0bbacce5c8d7aaf53d925dfd2552253d.
//
// Solidity: event ProverManagerUpdated(address indexed proverManager)
func (_CheckpointTracker *CheckpointTrackerFilterer) ParseProverManagerUpdated(log types.Log) (*CheckpointTrackerProverManagerUpdated, error) {
	event := new(CheckpointTrackerProverManagerUpdated)
	if err := _CheckpointTracker.contract.UnpackLog(event, "ProverManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
