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

// IERC1155BridgeERC1155Deposit is an auto generated low-level Go binding around an user-defined struct.
type IERC1155BridgeERC1155Deposit struct {
	Nonce         *big.Int
	From          common.Address
	To            common.Address
	OriginalToken common.Address
	TokenId       *big.Int
	Amount        *big.Int
	TokenURI      string
	Canceler      common.Address
}

// IERC1155BridgeTokenDescription is an auto generated low-level Go binding around an user-defined struct.
type IERC1155BridgeTokenDescription struct {
	OriginalToken common.Address
	Uri           string
}

// IERC1155BridgeMetaData contains all meta data concerning the IERC1155Bridge contract.
var IERC1155BridgeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"cancelDeposit\",\"inputs\":[{\"name\":\"erc1155Deposit\",\"type\":\"tuple\",\"internalType\":\"structIERC1155Bridge.ERC1155Deposit\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"canceler\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"claimee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimDeposit\",\"inputs\":[{\"name\":\"erc1155Deposit\",\"type\":\"tuple\",\"internalType\":\"structIERC1155Bridge.ERC1155Deposit\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"canceler\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployCounterpartToken\",\"inputs\":[{\"name\":\"tokenDesc\",\"type\":\"tuple\",\"internalType\":\"structIERC1155Bridge.TokenDescription\",\"components\":[{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"deployedToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"canceler\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCounterpartToken\",\"inputs\":[{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDepositId\",\"inputs\":[{\"name\":\"erc1155Deposit\",\"type\":\"tuple\",\"internalType\":\"structIERC1155Bridge.ERC1155Deposit\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"canceler\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getTokenDescriptionId\",\"inputs\":[{\"name\":\"tokenDesc\",\"type\":\"tuple\",\"internalType\":\"structIERC1155Bridge.TokenDescription\",\"components\":[{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"processed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recordTokenDescription\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CounterpartTokenDeployed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"description\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC1155Bridge.TokenDescription\",\"components\":[{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"deployedToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositCancelled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"claimee\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositClaimed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"deposit\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC1155Bridge.ERC1155Deposit\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"canceler\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositMade\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"deposit\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC1155Bridge.ERC1155Deposit\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"canceler\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenDescriptionRecorded\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"description\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC1155Bridge.TokenDescription\",\"components\":[{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CounterpartTokenAlreadyDeployed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCanceler\",\"inputs\":[]}]",
}

// IERC1155BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC1155BridgeMetaData.ABI instead.
var IERC1155BridgeABI = IERC1155BridgeMetaData.ABI

// IERC1155Bridge is an auto generated Go binding around an Ethereum contract.
type IERC1155Bridge struct {
	IERC1155BridgeCaller     // Read-only binding to the contract
	IERC1155BridgeTransactor // Write-only binding to the contract
	IERC1155BridgeFilterer   // Log filterer for contract events
}

// IERC1155BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1155BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1155BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC1155BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC1155BridgeSession struct {
	Contract     *IERC1155Bridge   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC1155BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC1155BridgeCallerSession struct {
	Contract *IERC1155BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IERC1155BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC1155BridgeTransactorSession struct {
	Contract     *IERC1155BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IERC1155BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC1155BridgeRaw struct {
	Contract *IERC1155Bridge // Generic contract binding to access the raw methods on
}

// IERC1155BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC1155BridgeCallerRaw struct {
	Contract *IERC1155BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// IERC1155BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC1155BridgeTransactorRaw struct {
	Contract *IERC1155BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC1155Bridge creates a new instance of IERC1155Bridge, bound to a specific deployed contract.
func NewIERC1155Bridge(address common.Address, backend bind.ContractBackend) (*IERC1155Bridge, error) {
	contract, err := bindIERC1155Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1155Bridge{IERC1155BridgeCaller: IERC1155BridgeCaller{contract: contract}, IERC1155BridgeTransactor: IERC1155BridgeTransactor{contract: contract}, IERC1155BridgeFilterer: IERC1155BridgeFilterer{contract: contract}}, nil
}

// NewIERC1155BridgeCaller creates a new read-only instance of IERC1155Bridge, bound to a specific deployed contract.
func NewIERC1155BridgeCaller(address common.Address, caller bind.ContractCaller) (*IERC1155BridgeCaller, error) {
	contract, err := bindIERC1155Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155BridgeCaller{contract: contract}, nil
}

// NewIERC1155BridgeTransactor creates a new write-only instance of IERC1155Bridge, bound to a specific deployed contract.
func NewIERC1155BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC1155BridgeTransactor, error) {
	contract, err := bindIERC1155Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155BridgeTransactor{contract: contract}, nil
}

// NewIERC1155BridgeFilterer creates a new log filterer instance of IERC1155Bridge, bound to a specific deployed contract.
func NewIERC1155BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC1155BridgeFilterer, error) {
	contract, err := bindIERC1155Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1155BridgeFilterer{contract: contract}, nil
}

// bindIERC1155Bridge binds a generic wrapper to an already deployed contract.
func bindIERC1155Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC1155BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1155Bridge *IERC1155BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1155Bridge.Contract.IERC1155BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1155Bridge *IERC1155BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1155Bridge.Contract.IERC1155BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1155Bridge *IERC1155BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1155Bridge.Contract.IERC1155BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1155Bridge *IERC1155BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1155Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1155Bridge *IERC1155BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1155Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1155Bridge *IERC1155BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1155Bridge.Contract.contract.Transact(opts, method, params...)
}

// GetCounterpartToken is a free data retrieval call binding the contract method 0xf859cca6.
//
// Solidity: function getCounterpartToken(address originalToken) view returns(address)
func (_IERC1155Bridge *IERC1155BridgeCaller) GetCounterpartToken(opts *bind.CallOpts, originalToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _IERC1155Bridge.contract.Call(opts, &out, "getCounterpartToken", originalToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCounterpartToken is a free data retrieval call binding the contract method 0xf859cca6.
//
// Solidity: function getCounterpartToken(address originalToken) view returns(address)
func (_IERC1155Bridge *IERC1155BridgeSession) GetCounterpartToken(originalToken common.Address) (common.Address, error) {
	return _IERC1155Bridge.Contract.GetCounterpartToken(&_IERC1155Bridge.CallOpts, originalToken)
}

// GetCounterpartToken is a free data retrieval call binding the contract method 0xf859cca6.
//
// Solidity: function getCounterpartToken(address originalToken) view returns(address)
func (_IERC1155Bridge *IERC1155BridgeCallerSession) GetCounterpartToken(originalToken common.Address) (common.Address, error) {
	return _IERC1155Bridge.Contract.GetCounterpartToken(&_IERC1155Bridge.CallOpts, originalToken)
}

// GetDepositId is a free data retrieval call binding the contract method 0xa7fc189b.
//
// Solidity: function getDepositId((uint256,address,address,address,uint256,uint256,string,address) erc1155Deposit) pure returns(bytes32 id)
func (_IERC1155Bridge *IERC1155BridgeCaller) GetDepositId(opts *bind.CallOpts, erc1155Deposit IERC1155BridgeERC1155Deposit) ([32]byte, error) {
	var out []interface{}
	err := _IERC1155Bridge.contract.Call(opts, &out, "getDepositId", erc1155Deposit)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDepositId is a free data retrieval call binding the contract method 0xa7fc189b.
//
// Solidity: function getDepositId((uint256,address,address,address,uint256,uint256,string,address) erc1155Deposit) pure returns(bytes32 id)
func (_IERC1155Bridge *IERC1155BridgeSession) GetDepositId(erc1155Deposit IERC1155BridgeERC1155Deposit) ([32]byte, error) {
	return _IERC1155Bridge.Contract.GetDepositId(&_IERC1155Bridge.CallOpts, erc1155Deposit)
}

// GetDepositId is a free data retrieval call binding the contract method 0xa7fc189b.
//
// Solidity: function getDepositId((uint256,address,address,address,uint256,uint256,string,address) erc1155Deposit) pure returns(bytes32 id)
func (_IERC1155Bridge *IERC1155BridgeCallerSession) GetDepositId(erc1155Deposit IERC1155BridgeERC1155Deposit) ([32]byte, error) {
	return _IERC1155Bridge.Contract.GetDepositId(&_IERC1155Bridge.CallOpts, erc1155Deposit)
}

// GetTokenDescriptionId is a free data retrieval call binding the contract method 0xb28e9415.
//
// Solidity: function getTokenDescriptionId((address,string) tokenDesc) pure returns(bytes32 id)
func (_IERC1155Bridge *IERC1155BridgeCaller) GetTokenDescriptionId(opts *bind.CallOpts, tokenDesc IERC1155BridgeTokenDescription) ([32]byte, error) {
	var out []interface{}
	err := _IERC1155Bridge.contract.Call(opts, &out, "getTokenDescriptionId", tokenDesc)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTokenDescriptionId is a free data retrieval call binding the contract method 0xb28e9415.
//
// Solidity: function getTokenDescriptionId((address,string) tokenDesc) pure returns(bytes32 id)
func (_IERC1155Bridge *IERC1155BridgeSession) GetTokenDescriptionId(tokenDesc IERC1155BridgeTokenDescription) ([32]byte, error) {
	return _IERC1155Bridge.Contract.GetTokenDescriptionId(&_IERC1155Bridge.CallOpts, tokenDesc)
}

// GetTokenDescriptionId is a free data retrieval call binding the contract method 0xb28e9415.
//
// Solidity: function getTokenDescriptionId((address,string) tokenDesc) pure returns(bytes32 id)
func (_IERC1155Bridge *IERC1155BridgeCallerSession) GetTokenDescriptionId(tokenDesc IERC1155BridgeTokenDescription) ([32]byte, error) {
	return _IERC1155Bridge.Contract.GetTokenDescriptionId(&_IERC1155Bridge.CallOpts, tokenDesc)
}

// Processed is a free data retrieval call binding the contract method 0xc1f0808a.
//
// Solidity: function processed(bytes32 id) view returns(bool)
func (_IERC1155Bridge *IERC1155BridgeCaller) Processed(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _IERC1155Bridge.contract.Call(opts, &out, "processed", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Processed is a free data retrieval call binding the contract method 0xc1f0808a.
//
// Solidity: function processed(bytes32 id) view returns(bool)
func (_IERC1155Bridge *IERC1155BridgeSession) Processed(id [32]byte) (bool, error) {
	return _IERC1155Bridge.Contract.Processed(&_IERC1155Bridge.CallOpts, id)
}

// Processed is a free data retrieval call binding the contract method 0xc1f0808a.
//
// Solidity: function processed(bytes32 id) view returns(bool)
func (_IERC1155Bridge *IERC1155BridgeCallerSession) Processed(id [32]byte) (bool, error) {
	return _IERC1155Bridge.Contract.Processed(&_IERC1155Bridge.CallOpts, id)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x8aabac1e.
//
// Solidity: function cancelDeposit((uint256,address,address,address,uint256,uint256,string,address) erc1155Deposit, address claimee, uint256 height, bytes proof) returns()
func (_IERC1155Bridge *IERC1155BridgeTransactor) CancelDeposit(opts *bind.TransactOpts, erc1155Deposit IERC1155BridgeERC1155Deposit, claimee common.Address, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC1155Bridge.contract.Transact(opts, "cancelDeposit", erc1155Deposit, claimee, height, proof)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x8aabac1e.
//
// Solidity: function cancelDeposit((uint256,address,address,address,uint256,uint256,string,address) erc1155Deposit, address claimee, uint256 height, bytes proof) returns()
func (_IERC1155Bridge *IERC1155BridgeSession) CancelDeposit(erc1155Deposit IERC1155BridgeERC1155Deposit, claimee common.Address, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC1155Bridge.Contract.CancelDeposit(&_IERC1155Bridge.TransactOpts, erc1155Deposit, claimee, height, proof)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x8aabac1e.
//
// Solidity: function cancelDeposit((uint256,address,address,address,uint256,uint256,string,address) erc1155Deposit, address claimee, uint256 height, bytes proof) returns()
func (_IERC1155Bridge *IERC1155BridgeTransactorSession) CancelDeposit(erc1155Deposit IERC1155BridgeERC1155Deposit, claimee common.Address, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC1155Bridge.Contract.CancelDeposit(&_IERC1155Bridge.TransactOpts, erc1155Deposit, claimee, height, proof)
}

// ClaimDeposit is a paid mutator transaction binding the contract method 0xe03756b9.
//
// Solidity: function claimDeposit((uint256,address,address,address,uint256,uint256,string,address) erc1155Deposit, uint256 height, bytes proof) returns()
func (_IERC1155Bridge *IERC1155BridgeTransactor) ClaimDeposit(opts *bind.TransactOpts, erc1155Deposit IERC1155BridgeERC1155Deposit, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC1155Bridge.contract.Transact(opts, "claimDeposit", erc1155Deposit, height, proof)
}

// ClaimDeposit is a paid mutator transaction binding the contract method 0xe03756b9.
//
// Solidity: function claimDeposit((uint256,address,address,address,uint256,uint256,string,address) erc1155Deposit, uint256 height, bytes proof) returns()
func (_IERC1155Bridge *IERC1155BridgeSession) ClaimDeposit(erc1155Deposit IERC1155BridgeERC1155Deposit, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC1155Bridge.Contract.ClaimDeposit(&_IERC1155Bridge.TransactOpts, erc1155Deposit, height, proof)
}

// ClaimDeposit is a paid mutator transaction binding the contract method 0xe03756b9.
//
// Solidity: function claimDeposit((uint256,address,address,address,uint256,uint256,string,address) erc1155Deposit, uint256 height, bytes proof) returns()
func (_IERC1155Bridge *IERC1155BridgeTransactorSession) ClaimDeposit(erc1155Deposit IERC1155BridgeERC1155Deposit, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC1155Bridge.Contract.ClaimDeposit(&_IERC1155Bridge.TransactOpts, erc1155Deposit, height, proof)
}

// DeployCounterpartToken is a paid mutator transaction binding the contract method 0xde3cda43.
//
// Solidity: function deployCounterpartToken((address,string) tokenDesc, uint256 height, bytes proof) returns(address deployedToken)
func (_IERC1155Bridge *IERC1155BridgeTransactor) DeployCounterpartToken(opts *bind.TransactOpts, tokenDesc IERC1155BridgeTokenDescription, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC1155Bridge.contract.Transact(opts, "deployCounterpartToken", tokenDesc, height, proof)
}

// DeployCounterpartToken is a paid mutator transaction binding the contract method 0xde3cda43.
//
// Solidity: function deployCounterpartToken((address,string) tokenDesc, uint256 height, bytes proof) returns(address deployedToken)
func (_IERC1155Bridge *IERC1155BridgeSession) DeployCounterpartToken(tokenDesc IERC1155BridgeTokenDescription, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC1155Bridge.Contract.DeployCounterpartToken(&_IERC1155Bridge.TransactOpts, tokenDesc, height, proof)
}

// DeployCounterpartToken is a paid mutator transaction binding the contract method 0xde3cda43.
//
// Solidity: function deployCounterpartToken((address,string) tokenDesc, uint256 height, bytes proof) returns(address deployedToken)
func (_IERC1155Bridge *IERC1155BridgeTransactorSession) DeployCounterpartToken(tokenDesc IERC1155BridgeTokenDescription, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC1155Bridge.Contract.DeployCounterpartToken(&_IERC1155Bridge.TransactOpts, tokenDesc, height, proof)
}

// Deposit is a paid mutator transaction binding the contract method 0xa3ed9cd0.
//
// Solidity: function deposit(address to, address localToken, uint256 tokenId, uint256 amount, address canceler) returns(bytes32 id)
func (_IERC1155Bridge *IERC1155BridgeTransactor) Deposit(opts *bind.TransactOpts, to common.Address, localToken common.Address, tokenId *big.Int, amount *big.Int, canceler common.Address) (*types.Transaction, error) {
	return _IERC1155Bridge.contract.Transact(opts, "deposit", to, localToken, tokenId, amount, canceler)
}

// Deposit is a paid mutator transaction binding the contract method 0xa3ed9cd0.
//
// Solidity: function deposit(address to, address localToken, uint256 tokenId, uint256 amount, address canceler) returns(bytes32 id)
func (_IERC1155Bridge *IERC1155BridgeSession) Deposit(to common.Address, localToken common.Address, tokenId *big.Int, amount *big.Int, canceler common.Address) (*types.Transaction, error) {
	return _IERC1155Bridge.Contract.Deposit(&_IERC1155Bridge.TransactOpts, to, localToken, tokenId, amount, canceler)
}

// Deposit is a paid mutator transaction binding the contract method 0xa3ed9cd0.
//
// Solidity: function deposit(address to, address localToken, uint256 tokenId, uint256 amount, address canceler) returns(bytes32 id)
func (_IERC1155Bridge *IERC1155BridgeTransactorSession) Deposit(to common.Address, localToken common.Address, tokenId *big.Int, amount *big.Int, canceler common.Address) (*types.Transaction, error) {
	return _IERC1155Bridge.Contract.Deposit(&_IERC1155Bridge.TransactOpts, to, localToken, tokenId, amount, canceler)
}

// RecordTokenDescription is a paid mutator transaction binding the contract method 0x174c049b.
//
// Solidity: function recordTokenDescription(address token) returns(bytes32 id)
func (_IERC1155Bridge *IERC1155BridgeTransactor) RecordTokenDescription(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _IERC1155Bridge.contract.Transact(opts, "recordTokenDescription", token)
}

// RecordTokenDescription is a paid mutator transaction binding the contract method 0x174c049b.
//
// Solidity: function recordTokenDescription(address token) returns(bytes32 id)
func (_IERC1155Bridge *IERC1155BridgeSession) RecordTokenDescription(token common.Address) (*types.Transaction, error) {
	return _IERC1155Bridge.Contract.RecordTokenDescription(&_IERC1155Bridge.TransactOpts, token)
}

// RecordTokenDescription is a paid mutator transaction binding the contract method 0x174c049b.
//
// Solidity: function recordTokenDescription(address token) returns(bytes32 id)
func (_IERC1155Bridge *IERC1155BridgeTransactorSession) RecordTokenDescription(token common.Address) (*types.Transaction, error) {
	return _IERC1155Bridge.Contract.RecordTokenDescription(&_IERC1155Bridge.TransactOpts, token)
}

// IERC1155BridgeCounterpartTokenDeployedIterator is returned from FilterCounterpartTokenDeployed and is used to iterate over the raw logs and unpacked data for CounterpartTokenDeployed events raised by the IERC1155Bridge contract.
type IERC1155BridgeCounterpartTokenDeployedIterator struct {
	Event *IERC1155BridgeCounterpartTokenDeployed // Event containing the contract specifics and raw log

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
func (it *IERC1155BridgeCounterpartTokenDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155BridgeCounterpartTokenDeployed)
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
		it.Event = new(IERC1155BridgeCounterpartTokenDeployed)
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
func (it *IERC1155BridgeCounterpartTokenDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155BridgeCounterpartTokenDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155BridgeCounterpartTokenDeployed represents a CounterpartTokenDeployed event raised by the IERC1155Bridge contract.
type IERC1155BridgeCounterpartTokenDeployed struct {
	Id            [32]byte
	Description   IERC1155BridgeTokenDescription
	DeployedToken common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCounterpartTokenDeployed is a free log retrieval operation binding the contract event 0xd9575e3205554fcebf0e69fefa7bdc6815981959cdf095373da5e9e393b9c833.
//
// Solidity: event CounterpartTokenDeployed(bytes32 indexed id, (address,string) description, address indexed deployedToken)
func (_IERC1155Bridge *IERC1155BridgeFilterer) FilterCounterpartTokenDeployed(opts *bind.FilterOpts, id [][32]byte, deployedToken []common.Address) (*IERC1155BridgeCounterpartTokenDeployedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var deployedTokenRule []interface{}
	for _, deployedTokenItem := range deployedToken {
		deployedTokenRule = append(deployedTokenRule, deployedTokenItem)
	}

	logs, sub, err := _IERC1155Bridge.contract.FilterLogs(opts, "CounterpartTokenDeployed", idRule, deployedTokenRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155BridgeCounterpartTokenDeployedIterator{contract: _IERC1155Bridge.contract, event: "CounterpartTokenDeployed", logs: logs, sub: sub}, nil
}

// WatchCounterpartTokenDeployed is a free log subscription operation binding the contract event 0xd9575e3205554fcebf0e69fefa7bdc6815981959cdf095373da5e9e393b9c833.
//
// Solidity: event CounterpartTokenDeployed(bytes32 indexed id, (address,string) description, address indexed deployedToken)
func (_IERC1155Bridge *IERC1155BridgeFilterer) WatchCounterpartTokenDeployed(opts *bind.WatchOpts, sink chan<- *IERC1155BridgeCounterpartTokenDeployed, id [][32]byte, deployedToken []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var deployedTokenRule []interface{}
	for _, deployedTokenItem := range deployedToken {
		deployedTokenRule = append(deployedTokenRule, deployedTokenItem)
	}

	logs, sub, err := _IERC1155Bridge.contract.WatchLogs(opts, "CounterpartTokenDeployed", idRule, deployedTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155BridgeCounterpartTokenDeployed)
				if err := _IERC1155Bridge.contract.UnpackLog(event, "CounterpartTokenDeployed", log); err != nil {
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

// ParseCounterpartTokenDeployed is a log parse operation binding the contract event 0xd9575e3205554fcebf0e69fefa7bdc6815981959cdf095373da5e9e393b9c833.
//
// Solidity: event CounterpartTokenDeployed(bytes32 indexed id, (address,string) description, address indexed deployedToken)
func (_IERC1155Bridge *IERC1155BridgeFilterer) ParseCounterpartTokenDeployed(log types.Log) (*IERC1155BridgeCounterpartTokenDeployed, error) {
	event := new(IERC1155BridgeCounterpartTokenDeployed)
	if err := _IERC1155Bridge.contract.UnpackLog(event, "CounterpartTokenDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155BridgeDepositCancelledIterator is returned from FilterDepositCancelled and is used to iterate over the raw logs and unpacked data for DepositCancelled events raised by the IERC1155Bridge contract.
type IERC1155BridgeDepositCancelledIterator struct {
	Event *IERC1155BridgeDepositCancelled // Event containing the contract specifics and raw log

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
func (it *IERC1155BridgeDepositCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155BridgeDepositCancelled)
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
		it.Event = new(IERC1155BridgeDepositCancelled)
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
func (it *IERC1155BridgeDepositCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155BridgeDepositCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155BridgeDepositCancelled represents a DepositCancelled event raised by the IERC1155Bridge contract.
type IERC1155BridgeDepositCancelled struct {
	Id      [32]byte
	Claimee common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositCancelled is a free log retrieval operation binding the contract event 0xda0bb473fd2c6fd2fcab08cf8a9cee668e53fc46468ffeabaaa147e0a3c226f3.
//
// Solidity: event DepositCancelled(bytes32 indexed id, address claimee)
func (_IERC1155Bridge *IERC1155BridgeFilterer) FilterDepositCancelled(opts *bind.FilterOpts, id [][32]byte) (*IERC1155BridgeDepositCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC1155Bridge.contract.FilterLogs(opts, "DepositCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155BridgeDepositCancelledIterator{contract: _IERC1155Bridge.contract, event: "DepositCancelled", logs: logs, sub: sub}, nil
}

// WatchDepositCancelled is a free log subscription operation binding the contract event 0xda0bb473fd2c6fd2fcab08cf8a9cee668e53fc46468ffeabaaa147e0a3c226f3.
//
// Solidity: event DepositCancelled(bytes32 indexed id, address claimee)
func (_IERC1155Bridge *IERC1155BridgeFilterer) WatchDepositCancelled(opts *bind.WatchOpts, sink chan<- *IERC1155BridgeDepositCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC1155Bridge.contract.WatchLogs(opts, "DepositCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155BridgeDepositCancelled)
				if err := _IERC1155Bridge.contract.UnpackLog(event, "DepositCancelled", log); err != nil {
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
func (_IERC1155Bridge *IERC1155BridgeFilterer) ParseDepositCancelled(log types.Log) (*IERC1155BridgeDepositCancelled, error) {
	event := new(IERC1155BridgeDepositCancelled)
	if err := _IERC1155Bridge.contract.UnpackLog(event, "DepositCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155BridgeDepositClaimedIterator is returned from FilterDepositClaimed and is used to iterate over the raw logs and unpacked data for DepositClaimed events raised by the IERC1155Bridge contract.
type IERC1155BridgeDepositClaimedIterator struct {
	Event *IERC1155BridgeDepositClaimed // Event containing the contract specifics and raw log

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
func (it *IERC1155BridgeDepositClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155BridgeDepositClaimed)
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
		it.Event = new(IERC1155BridgeDepositClaimed)
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
func (it *IERC1155BridgeDepositClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155BridgeDepositClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155BridgeDepositClaimed represents a DepositClaimed event raised by the IERC1155Bridge contract.
type IERC1155BridgeDepositClaimed struct {
	Id      [32]byte
	Deposit IERC1155BridgeERC1155Deposit
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositClaimed is a free log retrieval operation binding the contract event 0x2ef9280ff669717264dbcd2b0192bf5dff064034f0eebed9d5305c87f12e220d.
//
// Solidity: event DepositClaimed(bytes32 indexed id, (uint256,address,address,address,uint256,uint256,string,address) deposit)
func (_IERC1155Bridge *IERC1155BridgeFilterer) FilterDepositClaimed(opts *bind.FilterOpts, id [][32]byte) (*IERC1155BridgeDepositClaimedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC1155Bridge.contract.FilterLogs(opts, "DepositClaimed", idRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155BridgeDepositClaimedIterator{contract: _IERC1155Bridge.contract, event: "DepositClaimed", logs: logs, sub: sub}, nil
}

// WatchDepositClaimed is a free log subscription operation binding the contract event 0x2ef9280ff669717264dbcd2b0192bf5dff064034f0eebed9d5305c87f12e220d.
//
// Solidity: event DepositClaimed(bytes32 indexed id, (uint256,address,address,address,uint256,uint256,string,address) deposit)
func (_IERC1155Bridge *IERC1155BridgeFilterer) WatchDepositClaimed(opts *bind.WatchOpts, sink chan<- *IERC1155BridgeDepositClaimed, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC1155Bridge.contract.WatchLogs(opts, "DepositClaimed", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155BridgeDepositClaimed)
				if err := _IERC1155Bridge.contract.UnpackLog(event, "DepositClaimed", log); err != nil {
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

// ParseDepositClaimed is a log parse operation binding the contract event 0x2ef9280ff669717264dbcd2b0192bf5dff064034f0eebed9d5305c87f12e220d.
//
// Solidity: event DepositClaimed(bytes32 indexed id, (uint256,address,address,address,uint256,uint256,string,address) deposit)
func (_IERC1155Bridge *IERC1155BridgeFilterer) ParseDepositClaimed(log types.Log) (*IERC1155BridgeDepositClaimed, error) {
	event := new(IERC1155BridgeDepositClaimed)
	if err := _IERC1155Bridge.contract.UnpackLog(event, "DepositClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155BridgeDepositMadeIterator is returned from FilterDepositMade and is used to iterate over the raw logs and unpacked data for DepositMade events raised by the IERC1155Bridge contract.
type IERC1155BridgeDepositMadeIterator struct {
	Event *IERC1155BridgeDepositMade // Event containing the contract specifics and raw log

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
func (it *IERC1155BridgeDepositMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155BridgeDepositMade)
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
		it.Event = new(IERC1155BridgeDepositMade)
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
func (it *IERC1155BridgeDepositMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155BridgeDepositMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155BridgeDepositMade represents a DepositMade event raised by the IERC1155Bridge contract.
type IERC1155BridgeDepositMade struct {
	Id         [32]byte
	Deposit    IERC1155BridgeERC1155Deposit
	LocalToken common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositMade is a free log retrieval operation binding the contract event 0xabd8a33deca12ccb386a79788b1eed86b2a5f1b546d159792becb0663ec996d5.
//
// Solidity: event DepositMade(bytes32 indexed id, (uint256,address,address,address,uint256,uint256,string,address) deposit, address localToken)
func (_IERC1155Bridge *IERC1155BridgeFilterer) FilterDepositMade(opts *bind.FilterOpts, id [][32]byte) (*IERC1155BridgeDepositMadeIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC1155Bridge.contract.FilterLogs(opts, "DepositMade", idRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155BridgeDepositMadeIterator{contract: _IERC1155Bridge.contract, event: "DepositMade", logs: logs, sub: sub}, nil
}

// WatchDepositMade is a free log subscription operation binding the contract event 0xabd8a33deca12ccb386a79788b1eed86b2a5f1b546d159792becb0663ec996d5.
//
// Solidity: event DepositMade(bytes32 indexed id, (uint256,address,address,address,uint256,uint256,string,address) deposit, address localToken)
func (_IERC1155Bridge *IERC1155BridgeFilterer) WatchDepositMade(opts *bind.WatchOpts, sink chan<- *IERC1155BridgeDepositMade, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC1155Bridge.contract.WatchLogs(opts, "DepositMade", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155BridgeDepositMade)
				if err := _IERC1155Bridge.contract.UnpackLog(event, "DepositMade", log); err != nil {
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

// ParseDepositMade is a log parse operation binding the contract event 0xabd8a33deca12ccb386a79788b1eed86b2a5f1b546d159792becb0663ec996d5.
//
// Solidity: event DepositMade(bytes32 indexed id, (uint256,address,address,address,uint256,uint256,string,address) deposit, address localToken)
func (_IERC1155Bridge *IERC1155BridgeFilterer) ParseDepositMade(log types.Log) (*IERC1155BridgeDepositMade, error) {
	event := new(IERC1155BridgeDepositMade)
	if err := _IERC1155Bridge.contract.UnpackLog(event, "DepositMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155BridgeTokenDescriptionRecordedIterator is returned from FilterTokenDescriptionRecorded and is used to iterate over the raw logs and unpacked data for TokenDescriptionRecorded events raised by the IERC1155Bridge contract.
type IERC1155BridgeTokenDescriptionRecordedIterator struct {
	Event *IERC1155BridgeTokenDescriptionRecorded // Event containing the contract specifics and raw log

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
func (it *IERC1155BridgeTokenDescriptionRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155BridgeTokenDescriptionRecorded)
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
		it.Event = new(IERC1155BridgeTokenDescriptionRecorded)
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
func (it *IERC1155BridgeTokenDescriptionRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155BridgeTokenDescriptionRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155BridgeTokenDescriptionRecorded represents a TokenDescriptionRecorded event raised by the IERC1155Bridge contract.
type IERC1155BridgeTokenDescriptionRecorded struct {
	Id          [32]byte
	Description IERC1155BridgeTokenDescription
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenDescriptionRecorded is a free log retrieval operation binding the contract event 0xbe24fa65b4cdefaf43e2a1145efd257b3b0dc8fcba9139fa8f48df7e56fd9414.
//
// Solidity: event TokenDescriptionRecorded(bytes32 indexed id, (address,string) description)
func (_IERC1155Bridge *IERC1155BridgeFilterer) FilterTokenDescriptionRecorded(opts *bind.FilterOpts, id [][32]byte) (*IERC1155BridgeTokenDescriptionRecordedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC1155Bridge.contract.FilterLogs(opts, "TokenDescriptionRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155BridgeTokenDescriptionRecordedIterator{contract: _IERC1155Bridge.contract, event: "TokenDescriptionRecorded", logs: logs, sub: sub}, nil
}

// WatchTokenDescriptionRecorded is a free log subscription operation binding the contract event 0xbe24fa65b4cdefaf43e2a1145efd257b3b0dc8fcba9139fa8f48df7e56fd9414.
//
// Solidity: event TokenDescriptionRecorded(bytes32 indexed id, (address,string) description)
func (_IERC1155Bridge *IERC1155BridgeFilterer) WatchTokenDescriptionRecorded(opts *bind.WatchOpts, sink chan<- *IERC1155BridgeTokenDescriptionRecorded, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC1155Bridge.contract.WatchLogs(opts, "TokenDescriptionRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155BridgeTokenDescriptionRecorded)
				if err := _IERC1155Bridge.contract.UnpackLog(event, "TokenDescriptionRecorded", log); err != nil {
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

// ParseTokenDescriptionRecorded is a log parse operation binding the contract event 0xbe24fa65b4cdefaf43e2a1145efd257b3b0dc8fcba9139fa8f48df7e56fd9414.
//
// Solidity: event TokenDescriptionRecorded(bytes32 indexed id, (address,string) description)
func (_IERC1155Bridge *IERC1155BridgeFilterer) ParseTokenDescriptionRecorded(log types.Log) (*IERC1155BridgeTokenDescriptionRecorded, error) {
	event := new(IERC1155BridgeTokenDescriptionRecorded)
	if err := _IERC1155Bridge.contract.UnpackLog(event, "TokenDescriptionRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
