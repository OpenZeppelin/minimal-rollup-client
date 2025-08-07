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

// IERC721BridgeERC721Deposit is an auto generated low-level Go binding around an user-defined struct.
type IERC721BridgeERC721Deposit struct {
	Nonce         *big.Int
	From          common.Address
	To            common.Address
	OriginalToken common.Address
	TokenId       *big.Int
	TokenURI      string
	Canceler      common.Address
}

// IERC721BridgeTokenDescription is an auto generated low-level Go binding around an user-defined struct.
type IERC721BridgeTokenDescription struct {
	OriginalToken common.Address
	Name          string
	Symbol        string
}

// IERC721BridgeMetaData contains all meta data concerning the IERC721Bridge contract.
var IERC721BridgeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"cancelDeposit\",\"inputs\":[{\"name\":\"erc721Deposit\",\"type\":\"tuple\",\"internalType\":\"structIERC721Bridge.ERC721Deposit\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"canceler\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"claimee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimDeposit\",\"inputs\":[{\"name\":\"erc721Deposit\",\"type\":\"tuple\",\"internalType\":\"structIERC721Bridge.ERC721Deposit\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"canceler\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployCounterpartToken\",\"inputs\":[{\"name\":\"tokenDesc\",\"type\":\"tuple\",\"internalType\":\"structIERC721Bridge.TokenDescription\",\"components\":[{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"deployedToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"canceler\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCounterpartToken\",\"inputs\":[{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDepositId\",\"inputs\":[{\"name\":\"erc721Deposit\",\"type\":\"tuple\",\"internalType\":\"structIERC721Bridge.ERC721Deposit\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"canceler\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getTokenDescriptionId\",\"inputs\":[{\"name\":\"tokenDesc\",\"type\":\"tuple\",\"internalType\":\"structIERC721Bridge.TokenDescription\",\"components\":[{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"processed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recordTokenDescription\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CounterpartTokenDeployed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"description\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC721Bridge.TokenDescription\",\"components\":[{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"deployedToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositCancelled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"claimee\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositClaimed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"deposit\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC721Bridge.ERC721Deposit\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"canceler\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositMade\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"deposit\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC721Bridge.ERC721Deposit\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"canceler\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenDescriptionRecorded\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"description\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC721Bridge.TokenDescription\",\"components\":[{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CounterpartTokenAlreadyDeployed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCanceler\",\"inputs\":[]}]",
}

// IERC721BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC721BridgeMetaData.ABI instead.
var IERC721BridgeABI = IERC721BridgeMetaData.ABI

// IERC721Bridge is an auto generated Go binding around an Ethereum contract.
type IERC721Bridge struct {
	IERC721BridgeCaller     // Read-only binding to the contract
	IERC721BridgeTransactor // Write-only binding to the contract
	IERC721BridgeFilterer   // Log filterer for contract events
}

// IERC721BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721BridgeSession struct {
	Contract     *IERC721Bridge    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721BridgeCallerSession struct {
	Contract *IERC721BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IERC721BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721BridgeTransactorSession struct {
	Contract     *IERC721BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IERC721BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721BridgeRaw struct {
	Contract *IERC721Bridge // Generic contract binding to access the raw methods on
}

// IERC721BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721BridgeCallerRaw struct {
	Contract *IERC721BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721BridgeTransactorRaw struct {
	Contract *IERC721BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Bridge creates a new instance of IERC721Bridge, bound to a specific deployed contract.
func NewIERC721Bridge(address common.Address, backend bind.ContractBackend) (*IERC721Bridge, error) {
	contract, err := bindIERC721Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Bridge{IERC721BridgeCaller: IERC721BridgeCaller{contract: contract}, IERC721BridgeTransactor: IERC721BridgeTransactor{contract: contract}, IERC721BridgeFilterer: IERC721BridgeFilterer{contract: contract}}, nil
}

// NewIERC721BridgeCaller creates a new read-only instance of IERC721Bridge, bound to a specific deployed contract.
func NewIERC721BridgeCaller(address common.Address, caller bind.ContractCaller) (*IERC721BridgeCaller, error) {
	contract, err := bindIERC721Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721BridgeCaller{contract: contract}, nil
}

// NewIERC721BridgeTransactor creates a new write-only instance of IERC721Bridge, bound to a specific deployed contract.
func NewIERC721BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721BridgeTransactor, error) {
	contract, err := bindIERC721Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721BridgeTransactor{contract: contract}, nil
}

// NewIERC721BridgeFilterer creates a new log filterer instance of IERC721Bridge, bound to a specific deployed contract.
func NewIERC721BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721BridgeFilterer, error) {
	contract, err := bindIERC721Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721BridgeFilterer{contract: contract}, nil
}

// bindIERC721Bridge binds a generic wrapper to an already deployed contract.
func bindIERC721Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC721BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Bridge *IERC721BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Bridge.Contract.IERC721BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Bridge *IERC721BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Bridge.Contract.IERC721BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Bridge *IERC721BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Bridge.Contract.IERC721BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Bridge *IERC721BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Bridge *IERC721BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Bridge *IERC721BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Bridge.Contract.contract.Transact(opts, method, params...)
}

// GetCounterpartToken is a free data retrieval call binding the contract method 0xf859cca6.
//
// Solidity: function getCounterpartToken(address originalToken) view returns(address)
func (_IERC721Bridge *IERC721BridgeCaller) GetCounterpartToken(opts *bind.CallOpts, originalToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _IERC721Bridge.contract.Call(opts, &out, "getCounterpartToken", originalToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCounterpartToken is a free data retrieval call binding the contract method 0xf859cca6.
//
// Solidity: function getCounterpartToken(address originalToken) view returns(address)
func (_IERC721Bridge *IERC721BridgeSession) GetCounterpartToken(originalToken common.Address) (common.Address, error) {
	return _IERC721Bridge.Contract.GetCounterpartToken(&_IERC721Bridge.CallOpts, originalToken)
}

// GetCounterpartToken is a free data retrieval call binding the contract method 0xf859cca6.
//
// Solidity: function getCounterpartToken(address originalToken) view returns(address)
func (_IERC721Bridge *IERC721BridgeCallerSession) GetCounterpartToken(originalToken common.Address) (common.Address, error) {
	return _IERC721Bridge.Contract.GetCounterpartToken(&_IERC721Bridge.CallOpts, originalToken)
}

// GetDepositId is a free data retrieval call binding the contract method 0x61423679.
//
// Solidity: function getDepositId((uint256,address,address,address,uint256,string,address) erc721Deposit) pure returns(bytes32 id)
func (_IERC721Bridge *IERC721BridgeCaller) GetDepositId(opts *bind.CallOpts, erc721Deposit IERC721BridgeERC721Deposit) ([32]byte, error) {
	var out []interface{}
	err := _IERC721Bridge.contract.Call(opts, &out, "getDepositId", erc721Deposit)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDepositId is a free data retrieval call binding the contract method 0x61423679.
//
// Solidity: function getDepositId((uint256,address,address,address,uint256,string,address) erc721Deposit) pure returns(bytes32 id)
func (_IERC721Bridge *IERC721BridgeSession) GetDepositId(erc721Deposit IERC721BridgeERC721Deposit) ([32]byte, error) {
	return _IERC721Bridge.Contract.GetDepositId(&_IERC721Bridge.CallOpts, erc721Deposit)
}

// GetDepositId is a free data retrieval call binding the contract method 0x61423679.
//
// Solidity: function getDepositId((uint256,address,address,address,uint256,string,address) erc721Deposit) pure returns(bytes32 id)
func (_IERC721Bridge *IERC721BridgeCallerSession) GetDepositId(erc721Deposit IERC721BridgeERC721Deposit) ([32]byte, error) {
	return _IERC721Bridge.Contract.GetDepositId(&_IERC721Bridge.CallOpts, erc721Deposit)
}

// GetTokenDescriptionId is a free data retrieval call binding the contract method 0xa1665cd1.
//
// Solidity: function getTokenDescriptionId((address,string,string) tokenDesc) pure returns(bytes32 id)
func (_IERC721Bridge *IERC721BridgeCaller) GetTokenDescriptionId(opts *bind.CallOpts, tokenDesc IERC721BridgeTokenDescription) ([32]byte, error) {
	var out []interface{}
	err := _IERC721Bridge.contract.Call(opts, &out, "getTokenDescriptionId", tokenDesc)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTokenDescriptionId is a free data retrieval call binding the contract method 0xa1665cd1.
//
// Solidity: function getTokenDescriptionId((address,string,string) tokenDesc) pure returns(bytes32 id)
func (_IERC721Bridge *IERC721BridgeSession) GetTokenDescriptionId(tokenDesc IERC721BridgeTokenDescription) ([32]byte, error) {
	return _IERC721Bridge.Contract.GetTokenDescriptionId(&_IERC721Bridge.CallOpts, tokenDesc)
}

// GetTokenDescriptionId is a free data retrieval call binding the contract method 0xa1665cd1.
//
// Solidity: function getTokenDescriptionId((address,string,string) tokenDesc) pure returns(bytes32 id)
func (_IERC721Bridge *IERC721BridgeCallerSession) GetTokenDescriptionId(tokenDesc IERC721BridgeTokenDescription) ([32]byte, error) {
	return _IERC721Bridge.Contract.GetTokenDescriptionId(&_IERC721Bridge.CallOpts, tokenDesc)
}

// Processed is a free data retrieval call binding the contract method 0xc1f0808a.
//
// Solidity: function processed(bytes32 id) view returns(bool)
func (_IERC721Bridge *IERC721BridgeCaller) Processed(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _IERC721Bridge.contract.Call(opts, &out, "processed", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Processed is a free data retrieval call binding the contract method 0xc1f0808a.
//
// Solidity: function processed(bytes32 id) view returns(bool)
func (_IERC721Bridge *IERC721BridgeSession) Processed(id [32]byte) (bool, error) {
	return _IERC721Bridge.Contract.Processed(&_IERC721Bridge.CallOpts, id)
}

// Processed is a free data retrieval call binding the contract method 0xc1f0808a.
//
// Solidity: function processed(bytes32 id) view returns(bool)
func (_IERC721Bridge *IERC721BridgeCallerSession) Processed(id [32]byte) (bool, error) {
	return _IERC721Bridge.Contract.Processed(&_IERC721Bridge.CallOpts, id)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x27f8d16c.
//
// Solidity: function cancelDeposit((uint256,address,address,address,uint256,string,address) erc721Deposit, address claimee, uint256 height, bytes proof) returns()
func (_IERC721Bridge *IERC721BridgeTransactor) CancelDeposit(opts *bind.TransactOpts, erc721Deposit IERC721BridgeERC721Deposit, claimee common.Address, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC721Bridge.contract.Transact(opts, "cancelDeposit", erc721Deposit, claimee, height, proof)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x27f8d16c.
//
// Solidity: function cancelDeposit((uint256,address,address,address,uint256,string,address) erc721Deposit, address claimee, uint256 height, bytes proof) returns()
func (_IERC721Bridge *IERC721BridgeSession) CancelDeposit(erc721Deposit IERC721BridgeERC721Deposit, claimee common.Address, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC721Bridge.Contract.CancelDeposit(&_IERC721Bridge.TransactOpts, erc721Deposit, claimee, height, proof)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x27f8d16c.
//
// Solidity: function cancelDeposit((uint256,address,address,address,uint256,string,address) erc721Deposit, address claimee, uint256 height, bytes proof) returns()
func (_IERC721Bridge *IERC721BridgeTransactorSession) CancelDeposit(erc721Deposit IERC721BridgeERC721Deposit, claimee common.Address, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC721Bridge.Contract.CancelDeposit(&_IERC721Bridge.TransactOpts, erc721Deposit, claimee, height, proof)
}

// ClaimDeposit is a paid mutator transaction binding the contract method 0x41669e8b.
//
// Solidity: function claimDeposit((uint256,address,address,address,uint256,string,address) erc721Deposit, uint256 height, bytes proof) returns()
func (_IERC721Bridge *IERC721BridgeTransactor) ClaimDeposit(opts *bind.TransactOpts, erc721Deposit IERC721BridgeERC721Deposit, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC721Bridge.contract.Transact(opts, "claimDeposit", erc721Deposit, height, proof)
}

// ClaimDeposit is a paid mutator transaction binding the contract method 0x41669e8b.
//
// Solidity: function claimDeposit((uint256,address,address,address,uint256,string,address) erc721Deposit, uint256 height, bytes proof) returns()
func (_IERC721Bridge *IERC721BridgeSession) ClaimDeposit(erc721Deposit IERC721BridgeERC721Deposit, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC721Bridge.Contract.ClaimDeposit(&_IERC721Bridge.TransactOpts, erc721Deposit, height, proof)
}

// ClaimDeposit is a paid mutator transaction binding the contract method 0x41669e8b.
//
// Solidity: function claimDeposit((uint256,address,address,address,uint256,string,address) erc721Deposit, uint256 height, bytes proof) returns()
func (_IERC721Bridge *IERC721BridgeTransactorSession) ClaimDeposit(erc721Deposit IERC721BridgeERC721Deposit, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC721Bridge.Contract.ClaimDeposit(&_IERC721Bridge.TransactOpts, erc721Deposit, height, proof)
}

// DeployCounterpartToken is a paid mutator transaction binding the contract method 0x0ed9fedf.
//
// Solidity: function deployCounterpartToken((address,string,string) tokenDesc, uint256 height, bytes proof) returns(address deployedToken)
func (_IERC721Bridge *IERC721BridgeTransactor) DeployCounterpartToken(opts *bind.TransactOpts, tokenDesc IERC721BridgeTokenDescription, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC721Bridge.contract.Transact(opts, "deployCounterpartToken", tokenDesc, height, proof)
}

// DeployCounterpartToken is a paid mutator transaction binding the contract method 0x0ed9fedf.
//
// Solidity: function deployCounterpartToken((address,string,string) tokenDesc, uint256 height, bytes proof) returns(address deployedToken)
func (_IERC721Bridge *IERC721BridgeSession) DeployCounterpartToken(tokenDesc IERC721BridgeTokenDescription, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC721Bridge.Contract.DeployCounterpartToken(&_IERC721Bridge.TransactOpts, tokenDesc, height, proof)
}

// DeployCounterpartToken is a paid mutator transaction binding the contract method 0x0ed9fedf.
//
// Solidity: function deployCounterpartToken((address,string,string) tokenDesc, uint256 height, bytes proof) returns(address deployedToken)
func (_IERC721Bridge *IERC721BridgeTransactorSession) DeployCounterpartToken(tokenDesc IERC721BridgeTokenDescription, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC721Bridge.Contract.DeployCounterpartToken(&_IERC721Bridge.TransactOpts, tokenDesc, height, proof)
}

// Deposit is a paid mutator transaction binding the contract method 0x3bc1f1ed.
//
// Solidity: function deposit(address to, address localToken, uint256 tokenId, address canceler) returns(bytes32 id)
func (_IERC721Bridge *IERC721BridgeTransactor) Deposit(opts *bind.TransactOpts, to common.Address, localToken common.Address, tokenId *big.Int, canceler common.Address) (*types.Transaction, error) {
	return _IERC721Bridge.contract.Transact(opts, "deposit", to, localToken, tokenId, canceler)
}

// Deposit is a paid mutator transaction binding the contract method 0x3bc1f1ed.
//
// Solidity: function deposit(address to, address localToken, uint256 tokenId, address canceler) returns(bytes32 id)
func (_IERC721Bridge *IERC721BridgeSession) Deposit(to common.Address, localToken common.Address, tokenId *big.Int, canceler common.Address) (*types.Transaction, error) {
	return _IERC721Bridge.Contract.Deposit(&_IERC721Bridge.TransactOpts, to, localToken, tokenId, canceler)
}

// Deposit is a paid mutator transaction binding the contract method 0x3bc1f1ed.
//
// Solidity: function deposit(address to, address localToken, uint256 tokenId, address canceler) returns(bytes32 id)
func (_IERC721Bridge *IERC721BridgeTransactorSession) Deposit(to common.Address, localToken common.Address, tokenId *big.Int, canceler common.Address) (*types.Transaction, error) {
	return _IERC721Bridge.Contract.Deposit(&_IERC721Bridge.TransactOpts, to, localToken, tokenId, canceler)
}

// RecordTokenDescription is a paid mutator transaction binding the contract method 0x174c049b.
//
// Solidity: function recordTokenDescription(address token) returns(bytes32 id)
func (_IERC721Bridge *IERC721BridgeTransactor) RecordTokenDescription(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _IERC721Bridge.contract.Transact(opts, "recordTokenDescription", token)
}

// RecordTokenDescription is a paid mutator transaction binding the contract method 0x174c049b.
//
// Solidity: function recordTokenDescription(address token) returns(bytes32 id)
func (_IERC721Bridge *IERC721BridgeSession) RecordTokenDescription(token common.Address) (*types.Transaction, error) {
	return _IERC721Bridge.Contract.RecordTokenDescription(&_IERC721Bridge.TransactOpts, token)
}

// RecordTokenDescription is a paid mutator transaction binding the contract method 0x174c049b.
//
// Solidity: function recordTokenDescription(address token) returns(bytes32 id)
func (_IERC721Bridge *IERC721BridgeTransactorSession) RecordTokenDescription(token common.Address) (*types.Transaction, error) {
	return _IERC721Bridge.Contract.RecordTokenDescription(&_IERC721Bridge.TransactOpts, token)
}

// IERC721BridgeCounterpartTokenDeployedIterator is returned from FilterCounterpartTokenDeployed and is used to iterate over the raw logs and unpacked data for CounterpartTokenDeployed events raised by the IERC721Bridge contract.
type IERC721BridgeCounterpartTokenDeployedIterator struct {
	Event *IERC721BridgeCounterpartTokenDeployed // Event containing the contract specifics and raw log

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
func (it *IERC721BridgeCounterpartTokenDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721BridgeCounterpartTokenDeployed)
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
		it.Event = new(IERC721BridgeCounterpartTokenDeployed)
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
func (it *IERC721BridgeCounterpartTokenDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721BridgeCounterpartTokenDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721BridgeCounterpartTokenDeployed represents a CounterpartTokenDeployed event raised by the IERC721Bridge contract.
type IERC721BridgeCounterpartTokenDeployed struct {
	Id            [32]byte
	Description   IERC721BridgeTokenDescription
	DeployedToken common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCounterpartTokenDeployed is a free log retrieval operation binding the contract event 0x1414a634efcbec1806fc85612221c1a1b5093bed109ff11d696fcd371a77c53b.
//
// Solidity: event CounterpartTokenDeployed(bytes32 indexed id, (address,string,string) description, address indexed deployedToken)
func (_IERC721Bridge *IERC721BridgeFilterer) FilterCounterpartTokenDeployed(opts *bind.FilterOpts, id [][32]byte, deployedToken []common.Address) (*IERC721BridgeCounterpartTokenDeployedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var deployedTokenRule []interface{}
	for _, deployedTokenItem := range deployedToken {
		deployedTokenRule = append(deployedTokenRule, deployedTokenItem)
	}

	logs, sub, err := _IERC721Bridge.contract.FilterLogs(opts, "CounterpartTokenDeployed", idRule, deployedTokenRule)
	if err != nil {
		return nil, err
	}
	return &IERC721BridgeCounterpartTokenDeployedIterator{contract: _IERC721Bridge.contract, event: "CounterpartTokenDeployed", logs: logs, sub: sub}, nil
}

// WatchCounterpartTokenDeployed is a free log subscription operation binding the contract event 0x1414a634efcbec1806fc85612221c1a1b5093bed109ff11d696fcd371a77c53b.
//
// Solidity: event CounterpartTokenDeployed(bytes32 indexed id, (address,string,string) description, address indexed deployedToken)
func (_IERC721Bridge *IERC721BridgeFilterer) WatchCounterpartTokenDeployed(opts *bind.WatchOpts, sink chan<- *IERC721BridgeCounterpartTokenDeployed, id [][32]byte, deployedToken []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var deployedTokenRule []interface{}
	for _, deployedTokenItem := range deployedToken {
		deployedTokenRule = append(deployedTokenRule, deployedTokenItem)
	}

	logs, sub, err := _IERC721Bridge.contract.WatchLogs(opts, "CounterpartTokenDeployed", idRule, deployedTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721BridgeCounterpartTokenDeployed)
				if err := _IERC721Bridge.contract.UnpackLog(event, "CounterpartTokenDeployed", log); err != nil {
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

// ParseCounterpartTokenDeployed is a log parse operation binding the contract event 0x1414a634efcbec1806fc85612221c1a1b5093bed109ff11d696fcd371a77c53b.
//
// Solidity: event CounterpartTokenDeployed(bytes32 indexed id, (address,string,string) description, address indexed deployedToken)
func (_IERC721Bridge *IERC721BridgeFilterer) ParseCounterpartTokenDeployed(log types.Log) (*IERC721BridgeCounterpartTokenDeployed, error) {
	event := new(IERC721BridgeCounterpartTokenDeployed)
	if err := _IERC721Bridge.contract.UnpackLog(event, "CounterpartTokenDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721BridgeDepositCancelledIterator is returned from FilterDepositCancelled and is used to iterate over the raw logs and unpacked data for DepositCancelled events raised by the IERC721Bridge contract.
type IERC721BridgeDepositCancelledIterator struct {
	Event *IERC721BridgeDepositCancelled // Event containing the contract specifics and raw log

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
func (it *IERC721BridgeDepositCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721BridgeDepositCancelled)
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
		it.Event = new(IERC721BridgeDepositCancelled)
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
func (it *IERC721BridgeDepositCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721BridgeDepositCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721BridgeDepositCancelled represents a DepositCancelled event raised by the IERC721Bridge contract.
type IERC721BridgeDepositCancelled struct {
	Id      [32]byte
	Claimee common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositCancelled is a free log retrieval operation binding the contract event 0xda0bb473fd2c6fd2fcab08cf8a9cee668e53fc46468ffeabaaa147e0a3c226f3.
//
// Solidity: event DepositCancelled(bytes32 indexed id, address claimee)
func (_IERC721Bridge *IERC721BridgeFilterer) FilterDepositCancelled(opts *bind.FilterOpts, id [][32]byte) (*IERC721BridgeDepositCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC721Bridge.contract.FilterLogs(opts, "DepositCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &IERC721BridgeDepositCancelledIterator{contract: _IERC721Bridge.contract, event: "DepositCancelled", logs: logs, sub: sub}, nil
}

// WatchDepositCancelled is a free log subscription operation binding the contract event 0xda0bb473fd2c6fd2fcab08cf8a9cee668e53fc46468ffeabaaa147e0a3c226f3.
//
// Solidity: event DepositCancelled(bytes32 indexed id, address claimee)
func (_IERC721Bridge *IERC721BridgeFilterer) WatchDepositCancelled(opts *bind.WatchOpts, sink chan<- *IERC721BridgeDepositCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC721Bridge.contract.WatchLogs(opts, "DepositCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721BridgeDepositCancelled)
				if err := _IERC721Bridge.contract.UnpackLog(event, "DepositCancelled", log); err != nil {
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

// ParseDepositCancelled is a log parse operation binding the contract event 0xda0bb473fd2c6fd2fcab08cf8a9cee668e53fc46468ffeabaaa147e0a3c226f3.
//
// Solidity: event DepositCancelled(bytes32 indexed id, address claimee)
func (_IERC721Bridge *IERC721BridgeFilterer) ParseDepositCancelled(log types.Log) (*IERC721BridgeDepositCancelled, error) {
	event := new(IERC721BridgeDepositCancelled)
	if err := _IERC721Bridge.contract.UnpackLog(event, "DepositCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721BridgeDepositClaimedIterator is returned from FilterDepositClaimed and is used to iterate over the raw logs and unpacked data for DepositClaimed events raised by the IERC721Bridge contract.
type IERC721BridgeDepositClaimedIterator struct {
	Event *IERC721BridgeDepositClaimed // Event containing the contract specifics and raw log

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
func (it *IERC721BridgeDepositClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721BridgeDepositClaimed)
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
		it.Event = new(IERC721BridgeDepositClaimed)
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
func (it *IERC721BridgeDepositClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721BridgeDepositClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721BridgeDepositClaimed represents a DepositClaimed event raised by the IERC721Bridge contract.
type IERC721BridgeDepositClaimed struct {
	Id      [32]byte
	Deposit IERC721BridgeERC721Deposit
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositClaimed is a free log retrieval operation binding the contract event 0x60a60d4fd4b57ba9a4df27c2b19ed9e778d4d7820219fe321dbbe73a4c3824ab.
//
// Solidity: event DepositClaimed(bytes32 indexed id, (uint256,address,address,address,uint256,string,address) deposit)
func (_IERC721Bridge *IERC721BridgeFilterer) FilterDepositClaimed(opts *bind.FilterOpts, id [][32]byte) (*IERC721BridgeDepositClaimedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC721Bridge.contract.FilterLogs(opts, "DepositClaimed", idRule)
	if err != nil {
		return nil, err
	}
	return &IERC721BridgeDepositClaimedIterator{contract: _IERC721Bridge.contract, event: "DepositClaimed", logs: logs, sub: sub}, nil
}

// WatchDepositClaimed is a free log subscription operation binding the contract event 0x60a60d4fd4b57ba9a4df27c2b19ed9e778d4d7820219fe321dbbe73a4c3824ab.
//
// Solidity: event DepositClaimed(bytes32 indexed id, (uint256,address,address,address,uint256,string,address) deposit)
func (_IERC721Bridge *IERC721BridgeFilterer) WatchDepositClaimed(opts *bind.WatchOpts, sink chan<- *IERC721BridgeDepositClaimed, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC721Bridge.contract.WatchLogs(opts, "DepositClaimed", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721BridgeDepositClaimed)
				if err := _IERC721Bridge.contract.UnpackLog(event, "DepositClaimed", log); err != nil {
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

// ParseDepositClaimed is a log parse operation binding the contract event 0x60a60d4fd4b57ba9a4df27c2b19ed9e778d4d7820219fe321dbbe73a4c3824ab.
//
// Solidity: event DepositClaimed(bytes32 indexed id, (uint256,address,address,address,uint256,string,address) deposit)
func (_IERC721Bridge *IERC721BridgeFilterer) ParseDepositClaimed(log types.Log) (*IERC721BridgeDepositClaimed, error) {
	event := new(IERC721BridgeDepositClaimed)
	if err := _IERC721Bridge.contract.UnpackLog(event, "DepositClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721BridgeDepositMadeIterator is returned from FilterDepositMade and is used to iterate over the raw logs and unpacked data for DepositMade events raised by the IERC721Bridge contract.
type IERC721BridgeDepositMadeIterator struct {
	Event *IERC721BridgeDepositMade // Event containing the contract specifics and raw log

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
func (it *IERC721BridgeDepositMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721BridgeDepositMade)
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
		it.Event = new(IERC721BridgeDepositMade)
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
func (it *IERC721BridgeDepositMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721BridgeDepositMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721BridgeDepositMade represents a DepositMade event raised by the IERC721Bridge contract.
type IERC721BridgeDepositMade struct {
	Id         [32]byte
	Deposit    IERC721BridgeERC721Deposit
	LocalToken common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositMade is a free log retrieval operation binding the contract event 0x04d8a5951bc9705bab0b27bd21d440db2dadf11180beea04f624e33d03c26089.
//
// Solidity: event DepositMade(bytes32 indexed id, (uint256,address,address,address,uint256,string,address) deposit, address localToken)
func (_IERC721Bridge *IERC721BridgeFilterer) FilterDepositMade(opts *bind.FilterOpts, id [][32]byte) (*IERC721BridgeDepositMadeIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC721Bridge.contract.FilterLogs(opts, "DepositMade", idRule)
	if err != nil {
		return nil, err
	}
	return &IERC721BridgeDepositMadeIterator{contract: _IERC721Bridge.contract, event: "DepositMade", logs: logs, sub: sub}, nil
}

// WatchDepositMade is a free log subscription operation binding the contract event 0x04d8a5951bc9705bab0b27bd21d440db2dadf11180beea04f624e33d03c26089.
//
// Solidity: event DepositMade(bytes32 indexed id, (uint256,address,address,address,uint256,string,address) deposit, address localToken)
func (_IERC721Bridge *IERC721BridgeFilterer) WatchDepositMade(opts *bind.WatchOpts, sink chan<- *IERC721BridgeDepositMade, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC721Bridge.contract.WatchLogs(opts, "DepositMade", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721BridgeDepositMade)
				if err := _IERC721Bridge.contract.UnpackLog(event, "DepositMade", log); err != nil {
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

// ParseDepositMade is a log parse operation binding the contract event 0x04d8a5951bc9705bab0b27bd21d440db2dadf11180beea04f624e33d03c26089.
//
// Solidity: event DepositMade(bytes32 indexed id, (uint256,address,address,address,uint256,string,address) deposit, address localToken)
func (_IERC721Bridge *IERC721BridgeFilterer) ParseDepositMade(log types.Log) (*IERC721BridgeDepositMade, error) {
	event := new(IERC721BridgeDepositMade)
	if err := _IERC721Bridge.contract.UnpackLog(event, "DepositMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721BridgeTokenDescriptionRecordedIterator is returned from FilterTokenDescriptionRecorded and is used to iterate over the raw logs and unpacked data for TokenDescriptionRecorded events raised by the IERC721Bridge contract.
type IERC721BridgeTokenDescriptionRecordedIterator struct {
	Event *IERC721BridgeTokenDescriptionRecorded // Event containing the contract specifics and raw log

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
func (it *IERC721BridgeTokenDescriptionRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721BridgeTokenDescriptionRecorded)
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
		it.Event = new(IERC721BridgeTokenDescriptionRecorded)
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
func (it *IERC721BridgeTokenDescriptionRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721BridgeTokenDescriptionRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721BridgeTokenDescriptionRecorded represents a TokenDescriptionRecorded event raised by the IERC721Bridge contract.
type IERC721BridgeTokenDescriptionRecorded struct {
	Id          [32]byte
	Description IERC721BridgeTokenDescription
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenDescriptionRecorded is a free log retrieval operation binding the contract event 0x36d047aebc6c72a7c2640d5ddcd7c7f1114bf1895a49acf9905480ba312c3227.
//
// Solidity: event TokenDescriptionRecorded(bytes32 indexed id, (address,string,string) description)
func (_IERC721Bridge *IERC721BridgeFilterer) FilterTokenDescriptionRecorded(opts *bind.FilterOpts, id [][32]byte) (*IERC721BridgeTokenDescriptionRecordedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC721Bridge.contract.FilterLogs(opts, "TokenDescriptionRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return &IERC721BridgeTokenDescriptionRecordedIterator{contract: _IERC721Bridge.contract, event: "TokenDescriptionRecorded", logs: logs, sub: sub}, nil
}

// WatchTokenDescriptionRecorded is a free log subscription operation binding the contract event 0x36d047aebc6c72a7c2640d5ddcd7c7f1114bf1895a49acf9905480ba312c3227.
//
// Solidity: event TokenDescriptionRecorded(bytes32 indexed id, (address,string,string) description)
func (_IERC721Bridge *IERC721BridgeFilterer) WatchTokenDescriptionRecorded(opts *bind.WatchOpts, sink chan<- *IERC721BridgeTokenDescriptionRecorded, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC721Bridge.contract.WatchLogs(opts, "TokenDescriptionRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721BridgeTokenDescriptionRecorded)
				if err := _IERC721Bridge.contract.UnpackLog(event, "TokenDescriptionRecorded", log); err != nil {
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

// ParseTokenDescriptionRecorded is a log parse operation binding the contract event 0x36d047aebc6c72a7c2640d5ddcd7c7f1114bf1895a49acf9905480ba312c3227.
//
// Solidity: event TokenDescriptionRecorded(bytes32 indexed id, (address,string,string) description)
func (_IERC721Bridge *IERC721BridgeFilterer) ParseTokenDescriptionRecorded(log types.Log) (*IERC721BridgeTokenDescriptionRecorded, error) {
	event := new(IERC721BridgeTokenDescriptionRecorded)
	if err := _IERC721Bridge.contract.UnpackLog(event, "TokenDescriptionRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
