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

// DelayedInclusionStoreDueInclusion is an auto generated low-level Go binding around an user-defined struct.
type DelayedInclusionStoreDueInclusion struct {
	BlobRefHash [32]byte
	Due         *big.Int
}

// IDelayedInclusionStoreInclusion is an auto generated low-level Go binding around an user-defined struct.
type IDelayedInclusionStoreInclusion struct {
	BlobRefHash [32]byte
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

// TaikoInboxMetaData contains all meta data concerning the TaikoInbox contract.
var TaikoInboxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_lookahead\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_blobRefRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_maxAnchorBlockIdOffset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_inclusionDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blobRefRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBlobRefRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextPublicationId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPublicationHash\",\"inputs\":[{\"name\":\"idx\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inclusionDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lookahead\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILookahead\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxAnchorBlockIdOffset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposerFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIProposerFees\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publish\",\"inputs\":[{\"name\":\"nBlobs\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"anchorBlockId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"publishDelayed\",\"inputs\":[{\"name\":\"blobIndices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateProposerFees\",\"inputs\":[{\"name\":\"_proposerFees\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validateHeader\",\"inputs\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structIInbox.PublicationHeader\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"prevHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"publisher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"attributesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DelayedInclusionProcessed\",\"inputs\":[{\"name\":\"inclusionsList\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structIDelayedInclusionStore.Inclusion[]\",\"components\":[{\"name\":\"blobRefHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelayedInclusionStored\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dueInclusion\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structDelayedInclusionStore.DueInclusion\",\"components\":[{\"name\":\"blobRefHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"due\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposerFeesInitialized\",\"inputs\":[{\"name\":\"proposerFees\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Published\",\"inputs\":[{\"name\":\"pubHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"header\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIInbox.PublicationHeader\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"prevHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"publisher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"attributesHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"attributes\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AnchorBlockTooOld\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlockhashUnavailable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotCurrentPreconfer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroBlobRefRegistry\",\"inputs\":[]}]",
}

// TaikoInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use TaikoInboxMetaData.ABI instead.
var TaikoInboxABI = TaikoInboxMetaData.ABI

// TaikoInbox is an auto generated Go binding around an Ethereum contract.
type TaikoInbox struct {
	TaikoInboxCaller     // Read-only binding to the contract
	TaikoInboxTransactor // Write-only binding to the contract
	TaikoInboxFilterer   // Log filterer for contract events
}

// TaikoInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaikoInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaikoInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaikoInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaikoInboxSession struct {
	Contract     *TaikoInbox       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaikoInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaikoInboxCallerSession struct {
	Contract *TaikoInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TaikoInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaikoInboxTransactorSession struct {
	Contract     *TaikoInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TaikoInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaikoInboxRaw struct {
	Contract *TaikoInbox // Generic contract binding to access the raw methods on
}

// TaikoInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaikoInboxCallerRaw struct {
	Contract *TaikoInboxCaller // Generic read-only contract binding to access the raw methods on
}

// TaikoInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaikoInboxTransactorRaw struct {
	Contract *TaikoInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaikoInbox creates a new instance of TaikoInbox, bound to a specific deployed contract.
func NewTaikoInbox(address common.Address, backend bind.ContractBackend) (*TaikoInbox, error) {
	contract, err := bindTaikoInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaikoInbox{TaikoInboxCaller: TaikoInboxCaller{contract: contract}, TaikoInboxTransactor: TaikoInboxTransactor{contract: contract}, TaikoInboxFilterer: TaikoInboxFilterer{contract: contract}}, nil
}

// NewTaikoInboxCaller creates a new read-only instance of TaikoInbox, bound to a specific deployed contract.
func NewTaikoInboxCaller(address common.Address, caller bind.ContractCaller) (*TaikoInboxCaller, error) {
	contract, err := bindTaikoInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaikoInboxCaller{contract: contract}, nil
}

// NewTaikoInboxTransactor creates a new write-only instance of TaikoInbox, bound to a specific deployed contract.
func NewTaikoInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*TaikoInboxTransactor, error) {
	contract, err := bindTaikoInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaikoInboxTransactor{contract: contract}, nil
}

// NewTaikoInboxFilterer creates a new log filterer instance of TaikoInbox, bound to a specific deployed contract.
func NewTaikoInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*TaikoInboxFilterer, error) {
	contract, err := bindTaikoInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaikoInboxFilterer{contract: contract}, nil
}

// bindTaikoInbox binds a generic wrapper to an already deployed contract.
func bindTaikoInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaikoInboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaikoInbox *TaikoInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaikoInbox.Contract.TaikoInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaikoInbox *TaikoInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoInbox.Contract.TaikoInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaikoInbox *TaikoInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaikoInbox.Contract.TaikoInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaikoInbox *TaikoInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaikoInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaikoInbox *TaikoInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaikoInbox *TaikoInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaikoInbox.Contract.contract.Transact(opts, method, params...)
}

// BlobRefRegistry is a free data retrieval call binding the contract method 0x74eb5c62.
//
// Solidity: function blobRefRegistry() view returns(address)
func (_TaikoInbox *TaikoInboxCaller) BlobRefRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaikoInbox.contract.Call(opts, &out, "blobRefRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlobRefRegistry is a free data retrieval call binding the contract method 0x74eb5c62.
//
// Solidity: function blobRefRegistry() view returns(address)
func (_TaikoInbox *TaikoInboxSession) BlobRefRegistry() (common.Address, error) {
	return _TaikoInbox.Contract.BlobRefRegistry(&_TaikoInbox.CallOpts)
}

// BlobRefRegistry is a free data retrieval call binding the contract method 0x74eb5c62.
//
// Solidity: function blobRefRegistry() view returns(address)
func (_TaikoInbox *TaikoInboxCallerSession) BlobRefRegistry() (common.Address, error) {
	return _TaikoInbox.Contract.BlobRefRegistry(&_TaikoInbox.CallOpts)
}

// GetNextPublicationId is a free data retrieval call binding the contract method 0x416d86de.
//
// Solidity: function getNextPublicationId() view returns(uint256)
func (_TaikoInbox *TaikoInboxCaller) GetNextPublicationId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaikoInbox.contract.Call(opts, &out, "getNextPublicationId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextPublicationId is a free data retrieval call binding the contract method 0x416d86de.
//
// Solidity: function getNextPublicationId() view returns(uint256)
func (_TaikoInbox *TaikoInboxSession) GetNextPublicationId() (*big.Int, error) {
	return _TaikoInbox.Contract.GetNextPublicationId(&_TaikoInbox.CallOpts)
}

// GetNextPublicationId is a free data retrieval call binding the contract method 0x416d86de.
//
// Solidity: function getNextPublicationId() view returns(uint256)
func (_TaikoInbox *TaikoInboxCallerSession) GetNextPublicationId() (*big.Int, error) {
	return _TaikoInbox.Contract.GetNextPublicationId(&_TaikoInbox.CallOpts)
}

// GetPublicationHash is a free data retrieval call binding the contract method 0xef7922e3.
//
// Solidity: function getPublicationHash(uint256 idx) view returns(bytes32)
func (_TaikoInbox *TaikoInboxCaller) GetPublicationHash(opts *bind.CallOpts, idx *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TaikoInbox.contract.Call(opts, &out, "getPublicationHash", idx)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPublicationHash is a free data retrieval call binding the contract method 0xef7922e3.
//
// Solidity: function getPublicationHash(uint256 idx) view returns(bytes32)
func (_TaikoInbox *TaikoInboxSession) GetPublicationHash(idx *big.Int) ([32]byte, error) {
	return _TaikoInbox.Contract.GetPublicationHash(&_TaikoInbox.CallOpts, idx)
}

// GetPublicationHash is a free data retrieval call binding the contract method 0xef7922e3.
//
// Solidity: function getPublicationHash(uint256 idx) view returns(bytes32)
func (_TaikoInbox *TaikoInboxCallerSession) GetPublicationHash(idx *big.Int) ([32]byte, error) {
	return _TaikoInbox.Contract.GetPublicationHash(&_TaikoInbox.CallOpts, idx)
}

// InclusionDelay is a free data retrieval call binding the contract method 0xf765b4c3.
//
// Solidity: function inclusionDelay() view returns(uint256)
func (_TaikoInbox *TaikoInboxCaller) InclusionDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaikoInbox.contract.Call(opts, &out, "inclusionDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InclusionDelay is a free data retrieval call binding the contract method 0xf765b4c3.
//
// Solidity: function inclusionDelay() view returns(uint256)
func (_TaikoInbox *TaikoInboxSession) InclusionDelay() (*big.Int, error) {
	return _TaikoInbox.Contract.InclusionDelay(&_TaikoInbox.CallOpts)
}

// InclusionDelay is a free data retrieval call binding the contract method 0xf765b4c3.
//
// Solidity: function inclusionDelay() view returns(uint256)
func (_TaikoInbox *TaikoInboxCallerSession) InclusionDelay() (*big.Int, error) {
	return _TaikoInbox.Contract.InclusionDelay(&_TaikoInbox.CallOpts)
}

// Lookahead is a free data retrieval call binding the contract method 0x3b5351e5.
//
// Solidity: function lookahead() view returns(address)
func (_TaikoInbox *TaikoInboxCaller) Lookahead(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaikoInbox.contract.Call(opts, &out, "lookahead")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lookahead is a free data retrieval call binding the contract method 0x3b5351e5.
//
// Solidity: function lookahead() view returns(address)
func (_TaikoInbox *TaikoInboxSession) Lookahead() (common.Address, error) {
	return _TaikoInbox.Contract.Lookahead(&_TaikoInbox.CallOpts)
}

// Lookahead is a free data retrieval call binding the contract method 0x3b5351e5.
//
// Solidity: function lookahead() view returns(address)
func (_TaikoInbox *TaikoInboxCallerSession) Lookahead() (common.Address, error) {
	return _TaikoInbox.Contract.Lookahead(&_TaikoInbox.CallOpts)
}

// MaxAnchorBlockIdOffset is a free data retrieval call binding the contract method 0xc65474c9.
//
// Solidity: function maxAnchorBlockIdOffset() view returns(uint256)
func (_TaikoInbox *TaikoInboxCaller) MaxAnchorBlockIdOffset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaikoInbox.contract.Call(opts, &out, "maxAnchorBlockIdOffset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAnchorBlockIdOffset is a free data retrieval call binding the contract method 0xc65474c9.
//
// Solidity: function maxAnchorBlockIdOffset() view returns(uint256)
func (_TaikoInbox *TaikoInboxSession) MaxAnchorBlockIdOffset() (*big.Int, error) {
	return _TaikoInbox.Contract.MaxAnchorBlockIdOffset(&_TaikoInbox.CallOpts)
}

// MaxAnchorBlockIdOffset is a free data retrieval call binding the contract method 0xc65474c9.
//
// Solidity: function maxAnchorBlockIdOffset() view returns(uint256)
func (_TaikoInbox *TaikoInboxCallerSession) MaxAnchorBlockIdOffset() (*big.Int, error) {
	return _TaikoInbox.Contract.MaxAnchorBlockIdOffset(&_TaikoInbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoInbox *TaikoInboxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaikoInbox.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoInbox *TaikoInboxSession) Owner() (common.Address, error) {
	return _TaikoInbox.Contract.Owner(&_TaikoInbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoInbox *TaikoInboxCallerSession) Owner() (common.Address, error) {
	return _TaikoInbox.Contract.Owner(&_TaikoInbox.CallOpts)
}

// ProposerFees is a free data retrieval call binding the contract method 0x31e0102e.
//
// Solidity: function proposerFees() view returns(address)
func (_TaikoInbox *TaikoInboxCaller) ProposerFees(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaikoInbox.contract.Call(opts, &out, "proposerFees")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposerFees is a free data retrieval call binding the contract method 0x31e0102e.
//
// Solidity: function proposerFees() view returns(address)
func (_TaikoInbox *TaikoInboxSession) ProposerFees() (common.Address, error) {
	return _TaikoInbox.Contract.ProposerFees(&_TaikoInbox.CallOpts)
}

// ProposerFees is a free data retrieval call binding the contract method 0x31e0102e.
//
// Solidity: function proposerFees() view returns(address)
func (_TaikoInbox *TaikoInboxCallerSession) ProposerFees() (common.Address, error) {
	return _TaikoInbox.Contract.ProposerFees(&_TaikoInbox.CallOpts)
}

// ValidateHeader is a free data retrieval call binding the contract method 0x62c0e745.
//
// Solidity: function validateHeader((uint256,bytes32,address,uint256,uint256,bytes32) header) view returns(bool)
func (_TaikoInbox *TaikoInboxCaller) ValidateHeader(opts *bind.CallOpts, header IInboxPublicationHeader) (bool, error) {
	var out []interface{}
	err := _TaikoInbox.contract.Call(opts, &out, "validateHeader", header)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateHeader is a free data retrieval call binding the contract method 0x62c0e745.
//
// Solidity: function validateHeader((uint256,bytes32,address,uint256,uint256,bytes32) header) view returns(bool)
func (_TaikoInbox *TaikoInboxSession) ValidateHeader(header IInboxPublicationHeader) (bool, error) {
	return _TaikoInbox.Contract.ValidateHeader(&_TaikoInbox.CallOpts, header)
}

// ValidateHeader is a free data retrieval call binding the contract method 0x62c0e745.
//
// Solidity: function validateHeader((uint256,bytes32,address,uint256,uint256,bytes32) header) view returns(bool)
func (_TaikoInbox *TaikoInboxCallerSession) ValidateHeader(header IInboxPublicationHeader) (bool, error) {
	return _TaikoInbox.Contract.ValidateHeader(&_TaikoInbox.CallOpts, header)
}

// Publish is a paid mutator transaction binding the contract method 0x4039cc61.
//
// Solidity: function publish(uint256 nBlobs, uint64 anchorBlockId) returns()
func (_TaikoInbox *TaikoInboxTransactor) Publish(opts *bind.TransactOpts, nBlobs *big.Int, anchorBlockId uint64) (*types.Transaction, error) {
	return _TaikoInbox.contract.Transact(opts, "publish", nBlobs, anchorBlockId)
}

// Publish is a paid mutator transaction binding the contract method 0x4039cc61.
//
// Solidity: function publish(uint256 nBlobs, uint64 anchorBlockId) returns()
func (_TaikoInbox *TaikoInboxSession) Publish(nBlobs *big.Int, anchorBlockId uint64) (*types.Transaction, error) {
	return _TaikoInbox.Contract.Publish(&_TaikoInbox.TransactOpts, nBlobs, anchorBlockId)
}

// Publish is a paid mutator transaction binding the contract method 0x4039cc61.
//
// Solidity: function publish(uint256 nBlobs, uint64 anchorBlockId) returns()
func (_TaikoInbox *TaikoInboxTransactorSession) Publish(nBlobs *big.Int, anchorBlockId uint64) (*types.Transaction, error) {
	return _TaikoInbox.Contract.Publish(&_TaikoInbox.TransactOpts, nBlobs, anchorBlockId)
}

// PublishDelayed is a paid mutator transaction binding the contract method 0xd1831f60.
//
// Solidity: function publishDelayed(uint256[] blobIndices) returns()
func (_TaikoInbox *TaikoInboxTransactor) PublishDelayed(opts *bind.TransactOpts, blobIndices []*big.Int) (*types.Transaction, error) {
	return _TaikoInbox.contract.Transact(opts, "publishDelayed", blobIndices)
}

// PublishDelayed is a paid mutator transaction binding the contract method 0xd1831f60.
//
// Solidity: function publishDelayed(uint256[] blobIndices) returns()
func (_TaikoInbox *TaikoInboxSession) PublishDelayed(blobIndices []*big.Int) (*types.Transaction, error) {
	return _TaikoInbox.Contract.PublishDelayed(&_TaikoInbox.TransactOpts, blobIndices)
}

// PublishDelayed is a paid mutator transaction binding the contract method 0xd1831f60.
//
// Solidity: function publishDelayed(uint256[] blobIndices) returns()
func (_TaikoInbox *TaikoInboxTransactorSession) PublishDelayed(blobIndices []*big.Int) (*types.Transaction, error) {
	return _TaikoInbox.Contract.PublishDelayed(&_TaikoInbox.TransactOpts, blobIndices)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoInbox *TaikoInboxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoInbox.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoInbox *TaikoInboxSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaikoInbox.Contract.RenounceOwnership(&_TaikoInbox.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoInbox *TaikoInboxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaikoInbox.Contract.RenounceOwnership(&_TaikoInbox.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoInbox *TaikoInboxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TaikoInbox.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoInbox *TaikoInboxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaikoInbox.Contract.TransferOwnership(&_TaikoInbox.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoInbox *TaikoInboxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaikoInbox.Contract.TransferOwnership(&_TaikoInbox.TransactOpts, newOwner)
}

// UpdateProposerFees is a paid mutator transaction binding the contract method 0x18891fcc.
//
// Solidity: function updateProposerFees(address _proposerFees) returns()
func (_TaikoInbox *TaikoInboxTransactor) UpdateProposerFees(opts *bind.TransactOpts, _proposerFees common.Address) (*types.Transaction, error) {
	return _TaikoInbox.contract.Transact(opts, "updateProposerFees", _proposerFees)
}

// UpdateProposerFees is a paid mutator transaction binding the contract method 0x18891fcc.
//
// Solidity: function updateProposerFees(address _proposerFees) returns()
func (_TaikoInbox *TaikoInboxSession) UpdateProposerFees(_proposerFees common.Address) (*types.Transaction, error) {
	return _TaikoInbox.Contract.UpdateProposerFees(&_TaikoInbox.TransactOpts, _proposerFees)
}

// UpdateProposerFees is a paid mutator transaction binding the contract method 0x18891fcc.
//
// Solidity: function updateProposerFees(address _proposerFees) returns()
func (_TaikoInbox *TaikoInboxTransactorSession) UpdateProposerFees(_proposerFees common.Address) (*types.Transaction, error) {
	return _TaikoInbox.Contract.UpdateProposerFees(&_TaikoInbox.TransactOpts, _proposerFees)
}

// TaikoInboxDelayedInclusionProcessedIterator is returned from FilterDelayedInclusionProcessed and is used to iterate over the raw logs and unpacked data for DelayedInclusionProcessed events raised by the TaikoInbox contract.
type TaikoInboxDelayedInclusionProcessedIterator struct {
	Event *TaikoInboxDelayedInclusionProcessed // Event containing the contract specifics and raw log

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
func (it *TaikoInboxDelayedInclusionProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoInboxDelayedInclusionProcessed)
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
		it.Event = new(TaikoInboxDelayedInclusionProcessed)
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
func (it *TaikoInboxDelayedInclusionProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoInboxDelayedInclusionProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoInboxDelayedInclusionProcessed represents a DelayedInclusionProcessed event raised by the TaikoInbox contract.
type TaikoInboxDelayedInclusionProcessed struct {
	InclusionsList []IDelayedInclusionStoreInclusion
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDelayedInclusionProcessed is a free log retrieval operation binding the contract event 0x7008b97c3a37d78dd1f286f5b621404472aee22210b1c02406212f3ee87cfe70.
//
// Solidity: event DelayedInclusionProcessed((bytes32)[] inclusionsList)
func (_TaikoInbox *TaikoInboxFilterer) FilterDelayedInclusionProcessed(opts *bind.FilterOpts) (*TaikoInboxDelayedInclusionProcessedIterator, error) {

	logs, sub, err := _TaikoInbox.contract.FilterLogs(opts, "DelayedInclusionProcessed")
	if err != nil {
		return nil, err
	}
	return &TaikoInboxDelayedInclusionProcessedIterator{contract: _TaikoInbox.contract, event: "DelayedInclusionProcessed", logs: logs, sub: sub}, nil
}

// WatchDelayedInclusionProcessed is a free log subscription operation binding the contract event 0x7008b97c3a37d78dd1f286f5b621404472aee22210b1c02406212f3ee87cfe70.
//
// Solidity: event DelayedInclusionProcessed((bytes32)[] inclusionsList)
func (_TaikoInbox *TaikoInboxFilterer) WatchDelayedInclusionProcessed(opts *bind.WatchOpts, sink chan<- *TaikoInboxDelayedInclusionProcessed) (event.Subscription, error) {

	logs, sub, err := _TaikoInbox.contract.WatchLogs(opts, "DelayedInclusionProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoInboxDelayedInclusionProcessed)
				if err := _TaikoInbox.contract.UnpackLog(event, "DelayedInclusionProcessed", log); err != nil {
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
func (_TaikoInbox *TaikoInboxFilterer) ParseDelayedInclusionProcessed(log types.Log) (*TaikoInboxDelayedInclusionProcessed, error) {
	event := new(TaikoInboxDelayedInclusionProcessed)
	if err := _TaikoInbox.contract.UnpackLog(event, "DelayedInclusionProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoInboxDelayedInclusionStoredIterator is returned from FilterDelayedInclusionStored and is used to iterate over the raw logs and unpacked data for DelayedInclusionStored events raised by the TaikoInbox contract.
type TaikoInboxDelayedInclusionStoredIterator struct {
	Event *TaikoInboxDelayedInclusionStored // Event containing the contract specifics and raw log

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
func (it *TaikoInboxDelayedInclusionStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoInboxDelayedInclusionStored)
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
		it.Event = new(TaikoInboxDelayedInclusionStored)
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
func (it *TaikoInboxDelayedInclusionStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoInboxDelayedInclusionStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoInboxDelayedInclusionStored represents a DelayedInclusionStored event raised by the TaikoInbox contract.
type TaikoInboxDelayedInclusionStored struct {
	Sender       common.Address
	DueInclusion DelayedInclusionStoreDueInclusion
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelayedInclusionStored is a free log retrieval operation binding the contract event 0x6d87b8f648fd37d1ff7a2ecf57c51a3ac9830fc10bfc716f48223cb19de5b052.
//
// Solidity: event DelayedInclusionStored(address indexed sender, (bytes32,uint256) dueInclusion)
func (_TaikoInbox *TaikoInboxFilterer) FilterDelayedInclusionStored(opts *bind.FilterOpts, sender []common.Address) (*TaikoInboxDelayedInclusionStoredIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TaikoInbox.contract.FilterLogs(opts, "DelayedInclusionStored", senderRule)
	if err != nil {
		return nil, err
	}
	return &TaikoInboxDelayedInclusionStoredIterator{contract: _TaikoInbox.contract, event: "DelayedInclusionStored", logs: logs, sub: sub}, nil
}

// WatchDelayedInclusionStored is a free log subscription operation binding the contract event 0x6d87b8f648fd37d1ff7a2ecf57c51a3ac9830fc10bfc716f48223cb19de5b052.
//
// Solidity: event DelayedInclusionStored(address indexed sender, (bytes32,uint256) dueInclusion)
func (_TaikoInbox *TaikoInboxFilterer) WatchDelayedInclusionStored(opts *bind.WatchOpts, sink chan<- *TaikoInboxDelayedInclusionStored, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TaikoInbox.contract.WatchLogs(opts, "DelayedInclusionStored", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoInboxDelayedInclusionStored)
				if err := _TaikoInbox.contract.UnpackLog(event, "DelayedInclusionStored", log); err != nil {
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
func (_TaikoInbox *TaikoInboxFilterer) ParseDelayedInclusionStored(log types.Log) (*TaikoInboxDelayedInclusionStored, error) {
	event := new(TaikoInboxDelayedInclusionStored)
	if err := _TaikoInbox.contract.UnpackLog(event, "DelayedInclusionStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoInboxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TaikoInbox contract.
type TaikoInboxOwnershipTransferredIterator struct {
	Event *TaikoInboxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TaikoInboxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoInboxOwnershipTransferred)
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
		it.Event = new(TaikoInboxOwnershipTransferred)
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
func (it *TaikoInboxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoInboxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoInboxOwnershipTransferred represents a OwnershipTransferred event raised by the TaikoInbox contract.
type TaikoInboxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaikoInbox *TaikoInboxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TaikoInboxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaikoInbox.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TaikoInboxOwnershipTransferredIterator{contract: _TaikoInbox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaikoInbox *TaikoInboxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TaikoInboxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaikoInbox.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoInboxOwnershipTransferred)
				if err := _TaikoInbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TaikoInbox *TaikoInboxFilterer) ParseOwnershipTransferred(log types.Log) (*TaikoInboxOwnershipTransferred, error) {
	event := new(TaikoInboxOwnershipTransferred)
	if err := _TaikoInbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoInboxProposerFeesInitializedIterator is returned from FilterProposerFeesInitialized and is used to iterate over the raw logs and unpacked data for ProposerFeesInitialized events raised by the TaikoInbox contract.
type TaikoInboxProposerFeesInitializedIterator struct {
	Event *TaikoInboxProposerFeesInitialized // Event containing the contract specifics and raw log

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
func (it *TaikoInboxProposerFeesInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoInboxProposerFeesInitialized)
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
		it.Event = new(TaikoInboxProposerFeesInitialized)
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
func (it *TaikoInboxProposerFeesInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoInboxProposerFeesInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoInboxProposerFeesInitialized represents a ProposerFeesInitialized event raised by the TaikoInbox contract.
type TaikoInboxProposerFeesInitialized struct {
	ProposerFees common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposerFeesInitialized is a free log retrieval operation binding the contract event 0x9913f003385a4cd958cd6f2b42ee8ec6bcaff577177a4dd9f3d2f77889055338.
//
// Solidity: event ProposerFeesInitialized(address proposerFees)
func (_TaikoInbox *TaikoInboxFilterer) FilterProposerFeesInitialized(opts *bind.FilterOpts) (*TaikoInboxProposerFeesInitializedIterator, error) {

	logs, sub, err := _TaikoInbox.contract.FilterLogs(opts, "ProposerFeesInitialized")
	if err != nil {
		return nil, err
	}
	return &TaikoInboxProposerFeesInitializedIterator{contract: _TaikoInbox.contract, event: "ProposerFeesInitialized", logs: logs, sub: sub}, nil
}

// WatchProposerFeesInitialized is a free log subscription operation binding the contract event 0x9913f003385a4cd958cd6f2b42ee8ec6bcaff577177a4dd9f3d2f77889055338.
//
// Solidity: event ProposerFeesInitialized(address proposerFees)
func (_TaikoInbox *TaikoInboxFilterer) WatchProposerFeesInitialized(opts *bind.WatchOpts, sink chan<- *TaikoInboxProposerFeesInitialized) (event.Subscription, error) {

	logs, sub, err := _TaikoInbox.contract.WatchLogs(opts, "ProposerFeesInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoInboxProposerFeesInitialized)
				if err := _TaikoInbox.contract.UnpackLog(event, "ProposerFeesInitialized", log); err != nil {
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

// ParseProposerFeesInitialized is a log parse operation binding the contract event 0x9913f003385a4cd958cd6f2b42ee8ec6bcaff577177a4dd9f3d2f77889055338.
//
// Solidity: event ProposerFeesInitialized(address proposerFees)
func (_TaikoInbox *TaikoInboxFilterer) ParseProposerFeesInitialized(log types.Log) (*TaikoInboxProposerFeesInitialized, error) {
	event := new(TaikoInboxProposerFeesInitialized)
	if err := _TaikoInbox.contract.UnpackLog(event, "ProposerFeesInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoInboxPublishedIterator is returned from FilterPublished and is used to iterate over the raw logs and unpacked data for Published events raised by the TaikoInbox contract.
type TaikoInboxPublishedIterator struct {
	Event *TaikoInboxPublished // Event containing the contract specifics and raw log

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
func (it *TaikoInboxPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoInboxPublished)
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
		it.Event = new(TaikoInboxPublished)
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
func (it *TaikoInboxPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoInboxPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoInboxPublished represents a Published event raised by the TaikoInbox contract.
type TaikoInboxPublished struct {
	PubHash    [32]byte
	Header     IInboxPublicationHeader
	Attributes [][]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPublished is a free log retrieval operation binding the contract event 0x09a62876c6768f7c1c3233bafb09528ae9ff0fb0287ca35592ece85e856b31b2.
//
// Solidity: event Published(bytes32 indexed pubHash, (uint256,bytes32,address,uint256,uint256,bytes32) header, bytes[] attributes)
func (_TaikoInbox *TaikoInboxFilterer) FilterPublished(opts *bind.FilterOpts, pubHash [][32]byte) (*TaikoInboxPublishedIterator, error) {

	var pubHashRule []interface{}
	for _, pubHashItem := range pubHash {
		pubHashRule = append(pubHashRule, pubHashItem)
	}

	logs, sub, err := _TaikoInbox.contract.FilterLogs(opts, "Published", pubHashRule)
	if err != nil {
		return nil, err
	}
	return &TaikoInboxPublishedIterator{contract: _TaikoInbox.contract, event: "Published", logs: logs, sub: sub}, nil
}

// WatchPublished is a free log subscription operation binding the contract event 0x09a62876c6768f7c1c3233bafb09528ae9ff0fb0287ca35592ece85e856b31b2.
//
// Solidity: event Published(bytes32 indexed pubHash, (uint256,bytes32,address,uint256,uint256,bytes32) header, bytes[] attributes)
func (_TaikoInbox *TaikoInboxFilterer) WatchPublished(opts *bind.WatchOpts, sink chan<- *TaikoInboxPublished, pubHash [][32]byte) (event.Subscription, error) {

	var pubHashRule []interface{}
	for _, pubHashItem := range pubHash {
		pubHashRule = append(pubHashRule, pubHashItem)
	}

	logs, sub, err := _TaikoInbox.contract.WatchLogs(opts, "Published", pubHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoInboxPublished)
				if err := _TaikoInbox.contract.UnpackLog(event, "Published", log); err != nil {
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
func (_TaikoInbox *TaikoInboxFilterer) ParsePublished(log types.Log) (*TaikoInboxPublished, error) {
	event := new(TaikoInboxPublished)
	if err := _TaikoInbox.contract.UnpackLog(event, "Published", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
