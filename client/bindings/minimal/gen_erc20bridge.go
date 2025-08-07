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

// IERC20BridgeERC20Deposit is an auto generated low-level Go binding around an user-defined struct.
type IERC20BridgeERC20Deposit struct {
	Nonce         *big.Int
	From          common.Address
	To            common.Address
	OriginalToken common.Address
	Amount        *big.Int
}

// IERC20BridgeTokenDescription is an auto generated low-level Go binding around an user-defined struct.
type IERC20BridgeTokenDescription struct {
	OriginalToken common.Address
	Name          string
	Symbol        string
	Decimals      uint8
}

// IERC20BridgeMetaData contains all meta data concerning the IERC20Bridge contract.
var IERC20BridgeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"claimDeposit\",\"inputs\":[{\"name\":\"erc20Deposit\",\"type\":\"tuple\",\"internalType\":\"structIERC20Bridge.ERC20Deposit\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployCounterpartToken\",\"inputs\":[{\"name\":\"tokenDesc\",\"type\":\"tuple\",\"internalType\":\"structIERC20Bridge.TokenDescription\",\"components\":[{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"deployedToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCounterpartToken\",\"inputs\":[{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDepositId\",\"inputs\":[{\"name\":\"erc20Deposit\",\"type\":\"tuple\",\"internalType\":\"structIERC20Bridge.ERC20Deposit\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getTokenDescriptionId\",\"inputs\":[{\"name\":\"tokenDesc\",\"type\":\"tuple\",\"internalType\":\"structIERC20Bridge.TokenDescription\",\"components\":[{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"processed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recordTokenDescription\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CounterpartTokenDeployed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"description\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC20Bridge.TokenDescription\",\"components\":[{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"deployedToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositClaimed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"deposit\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC20Bridge.ERC20Deposit\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositMade\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"deposit\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC20Bridge.ERC20Deposit\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenDescriptionRecorded\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"description\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIERC20Bridge.TokenDescription\",\"components\":[{\"name\":\"originalToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CounterpartTokenAlreadyDeployed\",\"inputs\":[]}]",
}

// IERC20BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20BridgeMetaData.ABI instead.
var IERC20BridgeABI = IERC20BridgeMetaData.ABI

// IERC20Bridge is an auto generated Go binding around an Ethereum contract.
type IERC20Bridge struct {
	IERC20BridgeCaller     // Read-only binding to the contract
	IERC20BridgeTransactor // Write-only binding to the contract
	IERC20BridgeFilterer   // Log filterer for contract events
}

// IERC20BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20BridgeSession struct {
	Contract     *IERC20Bridge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20BridgeCallerSession struct {
	Contract *IERC20BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IERC20BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20BridgeTransactorSession struct {
	Contract     *IERC20BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IERC20BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20BridgeRaw struct {
	Contract *IERC20Bridge // Generic contract binding to access the raw methods on
}

// IERC20BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20BridgeCallerRaw struct {
	Contract *IERC20BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20BridgeTransactorRaw struct {
	Contract *IERC20BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Bridge creates a new instance of IERC20Bridge, bound to a specific deployed contract.
func NewIERC20Bridge(address common.Address, backend bind.ContractBackend) (*IERC20Bridge, error) {
	contract, err := bindIERC20Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Bridge{IERC20BridgeCaller: IERC20BridgeCaller{contract: contract}, IERC20BridgeTransactor: IERC20BridgeTransactor{contract: contract}, IERC20BridgeFilterer: IERC20BridgeFilterer{contract: contract}}, nil
}

// NewIERC20BridgeCaller creates a new read-only instance of IERC20Bridge, bound to a specific deployed contract.
func NewIERC20BridgeCaller(address common.Address, caller bind.ContractCaller) (*IERC20BridgeCaller, error) {
	contract, err := bindIERC20Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20BridgeCaller{contract: contract}, nil
}

// NewIERC20BridgeTransactor creates a new write-only instance of IERC20Bridge, bound to a specific deployed contract.
func NewIERC20BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20BridgeTransactor, error) {
	contract, err := bindIERC20Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20BridgeTransactor{contract: contract}, nil
}

// NewIERC20BridgeFilterer creates a new log filterer instance of IERC20Bridge, bound to a specific deployed contract.
func NewIERC20BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20BridgeFilterer, error) {
	contract, err := bindIERC20Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20BridgeFilterer{contract: contract}, nil
}

// bindIERC20Bridge binds a generic wrapper to an already deployed contract.
func bindIERC20Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Bridge *IERC20BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Bridge.Contract.IERC20BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Bridge *IERC20BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Bridge.Contract.IERC20BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Bridge *IERC20BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Bridge.Contract.IERC20BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Bridge *IERC20BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Bridge *IERC20BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Bridge *IERC20BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Bridge.Contract.contract.Transact(opts, method, params...)
}

// GetCounterpartToken is a free data retrieval call binding the contract method 0xf859cca6.
//
// Solidity: function getCounterpartToken(address originalToken) view returns(address)
func (_IERC20Bridge *IERC20BridgeCaller) GetCounterpartToken(opts *bind.CallOpts, originalToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _IERC20Bridge.contract.Call(opts, &out, "getCounterpartToken", originalToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCounterpartToken is a free data retrieval call binding the contract method 0xf859cca6.
//
// Solidity: function getCounterpartToken(address originalToken) view returns(address)
func (_IERC20Bridge *IERC20BridgeSession) GetCounterpartToken(originalToken common.Address) (common.Address, error) {
	return _IERC20Bridge.Contract.GetCounterpartToken(&_IERC20Bridge.CallOpts, originalToken)
}

// GetCounterpartToken is a free data retrieval call binding the contract method 0xf859cca6.
//
// Solidity: function getCounterpartToken(address originalToken) view returns(address)
func (_IERC20Bridge *IERC20BridgeCallerSession) GetCounterpartToken(originalToken common.Address) (common.Address, error) {
	return _IERC20Bridge.Contract.GetCounterpartToken(&_IERC20Bridge.CallOpts, originalToken)
}

// GetDepositId is a free data retrieval call binding the contract method 0x0fb65108.
//
// Solidity: function getDepositId((uint256,address,address,address,uint256) erc20Deposit) pure returns(bytes32 id)
func (_IERC20Bridge *IERC20BridgeCaller) GetDepositId(opts *bind.CallOpts, erc20Deposit IERC20BridgeERC20Deposit) ([32]byte, error) {
	var out []interface{}
	err := _IERC20Bridge.contract.Call(opts, &out, "getDepositId", erc20Deposit)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDepositId is a free data retrieval call binding the contract method 0x0fb65108.
//
// Solidity: function getDepositId((uint256,address,address,address,uint256) erc20Deposit) pure returns(bytes32 id)
func (_IERC20Bridge *IERC20BridgeSession) GetDepositId(erc20Deposit IERC20BridgeERC20Deposit) ([32]byte, error) {
	return _IERC20Bridge.Contract.GetDepositId(&_IERC20Bridge.CallOpts, erc20Deposit)
}

// GetDepositId is a free data retrieval call binding the contract method 0x0fb65108.
//
// Solidity: function getDepositId((uint256,address,address,address,uint256) erc20Deposit) pure returns(bytes32 id)
func (_IERC20Bridge *IERC20BridgeCallerSession) GetDepositId(erc20Deposit IERC20BridgeERC20Deposit) ([32]byte, error) {
	return _IERC20Bridge.Contract.GetDepositId(&_IERC20Bridge.CallOpts, erc20Deposit)
}

// GetTokenDescriptionId is a free data retrieval call binding the contract method 0x1e0c35d7.
//
// Solidity: function getTokenDescriptionId((address,string,string,uint8) tokenDesc) pure returns(bytes32 id)
func (_IERC20Bridge *IERC20BridgeCaller) GetTokenDescriptionId(opts *bind.CallOpts, tokenDesc IERC20BridgeTokenDescription) ([32]byte, error) {
	var out []interface{}
	err := _IERC20Bridge.contract.Call(opts, &out, "getTokenDescriptionId", tokenDesc)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTokenDescriptionId is a free data retrieval call binding the contract method 0x1e0c35d7.
//
// Solidity: function getTokenDescriptionId((address,string,string,uint8) tokenDesc) pure returns(bytes32 id)
func (_IERC20Bridge *IERC20BridgeSession) GetTokenDescriptionId(tokenDesc IERC20BridgeTokenDescription) ([32]byte, error) {
	return _IERC20Bridge.Contract.GetTokenDescriptionId(&_IERC20Bridge.CallOpts, tokenDesc)
}

// GetTokenDescriptionId is a free data retrieval call binding the contract method 0x1e0c35d7.
//
// Solidity: function getTokenDescriptionId((address,string,string,uint8) tokenDesc) pure returns(bytes32 id)
func (_IERC20Bridge *IERC20BridgeCallerSession) GetTokenDescriptionId(tokenDesc IERC20BridgeTokenDescription) ([32]byte, error) {
	return _IERC20Bridge.Contract.GetTokenDescriptionId(&_IERC20Bridge.CallOpts, tokenDesc)
}

// Processed is a free data retrieval call binding the contract method 0xc1f0808a.
//
// Solidity: function processed(bytes32 id) view returns(bool)
func (_IERC20Bridge *IERC20BridgeCaller) Processed(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _IERC20Bridge.contract.Call(opts, &out, "processed", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Processed is a free data retrieval call binding the contract method 0xc1f0808a.
//
// Solidity: function processed(bytes32 id) view returns(bool)
func (_IERC20Bridge *IERC20BridgeSession) Processed(id [32]byte) (bool, error) {
	return _IERC20Bridge.Contract.Processed(&_IERC20Bridge.CallOpts, id)
}

// Processed is a free data retrieval call binding the contract method 0xc1f0808a.
//
// Solidity: function processed(bytes32 id) view returns(bool)
func (_IERC20Bridge *IERC20BridgeCallerSession) Processed(id [32]byte) (bool, error) {
	return _IERC20Bridge.Contract.Processed(&_IERC20Bridge.CallOpts, id)
}

// ClaimDeposit is a paid mutator transaction binding the contract method 0x38d3c8de.
//
// Solidity: function claimDeposit((uint256,address,address,address,uint256) erc20Deposit, uint256 height, bytes proof) returns()
func (_IERC20Bridge *IERC20BridgeTransactor) ClaimDeposit(opts *bind.TransactOpts, erc20Deposit IERC20BridgeERC20Deposit, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC20Bridge.contract.Transact(opts, "claimDeposit", erc20Deposit, height, proof)
}

// ClaimDeposit is a paid mutator transaction binding the contract method 0x38d3c8de.
//
// Solidity: function claimDeposit((uint256,address,address,address,uint256) erc20Deposit, uint256 height, bytes proof) returns()
func (_IERC20Bridge *IERC20BridgeSession) ClaimDeposit(erc20Deposit IERC20BridgeERC20Deposit, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC20Bridge.Contract.ClaimDeposit(&_IERC20Bridge.TransactOpts, erc20Deposit, height, proof)
}

// ClaimDeposit is a paid mutator transaction binding the contract method 0x38d3c8de.
//
// Solidity: function claimDeposit((uint256,address,address,address,uint256) erc20Deposit, uint256 height, bytes proof) returns()
func (_IERC20Bridge *IERC20BridgeTransactorSession) ClaimDeposit(erc20Deposit IERC20BridgeERC20Deposit, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC20Bridge.Contract.ClaimDeposit(&_IERC20Bridge.TransactOpts, erc20Deposit, height, proof)
}

// DeployCounterpartToken is a paid mutator transaction binding the contract method 0x0d3ecdf1.
//
// Solidity: function deployCounterpartToken((address,string,string,uint8) tokenDesc, uint256 height, bytes proof) returns(address deployedToken)
func (_IERC20Bridge *IERC20BridgeTransactor) DeployCounterpartToken(opts *bind.TransactOpts, tokenDesc IERC20BridgeTokenDescription, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC20Bridge.contract.Transact(opts, "deployCounterpartToken", tokenDesc, height, proof)
}

// DeployCounterpartToken is a paid mutator transaction binding the contract method 0x0d3ecdf1.
//
// Solidity: function deployCounterpartToken((address,string,string,uint8) tokenDesc, uint256 height, bytes proof) returns(address deployedToken)
func (_IERC20Bridge *IERC20BridgeSession) DeployCounterpartToken(tokenDesc IERC20BridgeTokenDescription, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC20Bridge.Contract.DeployCounterpartToken(&_IERC20Bridge.TransactOpts, tokenDesc, height, proof)
}

// DeployCounterpartToken is a paid mutator transaction binding the contract method 0x0d3ecdf1.
//
// Solidity: function deployCounterpartToken((address,string,string,uint8) tokenDesc, uint256 height, bytes proof) returns(address deployedToken)
func (_IERC20Bridge *IERC20BridgeTransactorSession) DeployCounterpartToken(tokenDesc IERC20BridgeTokenDescription, height *big.Int, proof []byte) (*types.Transaction, error) {
	return _IERC20Bridge.Contract.DeployCounterpartToken(&_IERC20Bridge.TransactOpts, tokenDesc, height, proof)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address to, address localToken, uint256 amount) returns(bytes32 id)
func (_IERC20Bridge *IERC20BridgeTransactor) Deposit(opts *bind.TransactOpts, to common.Address, localToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Bridge.contract.Transact(opts, "deposit", to, localToken, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address to, address localToken, uint256 amount) returns(bytes32 id)
func (_IERC20Bridge *IERC20BridgeSession) Deposit(to common.Address, localToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Bridge.Contract.Deposit(&_IERC20Bridge.TransactOpts, to, localToken, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address to, address localToken, uint256 amount) returns(bytes32 id)
func (_IERC20Bridge *IERC20BridgeTransactorSession) Deposit(to common.Address, localToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Bridge.Contract.Deposit(&_IERC20Bridge.TransactOpts, to, localToken, amount)
}

// RecordTokenDescription is a paid mutator transaction binding the contract method 0x174c049b.
//
// Solidity: function recordTokenDescription(address token) returns(bytes32 id)
func (_IERC20Bridge *IERC20BridgeTransactor) RecordTokenDescription(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _IERC20Bridge.contract.Transact(opts, "recordTokenDescription", token)
}

// RecordTokenDescription is a paid mutator transaction binding the contract method 0x174c049b.
//
// Solidity: function recordTokenDescription(address token) returns(bytes32 id)
func (_IERC20Bridge *IERC20BridgeSession) RecordTokenDescription(token common.Address) (*types.Transaction, error) {
	return _IERC20Bridge.Contract.RecordTokenDescription(&_IERC20Bridge.TransactOpts, token)
}

// RecordTokenDescription is a paid mutator transaction binding the contract method 0x174c049b.
//
// Solidity: function recordTokenDescription(address token) returns(bytes32 id)
func (_IERC20Bridge *IERC20BridgeTransactorSession) RecordTokenDescription(token common.Address) (*types.Transaction, error) {
	return _IERC20Bridge.Contract.RecordTokenDescription(&_IERC20Bridge.TransactOpts, token)
}

// IERC20BridgeCounterpartTokenDeployedIterator is returned from FilterCounterpartTokenDeployed and is used to iterate over the raw logs and unpacked data for CounterpartTokenDeployed events raised by the IERC20Bridge contract.
type IERC20BridgeCounterpartTokenDeployedIterator struct {
	Event *IERC20BridgeCounterpartTokenDeployed // Event containing the contract specifics and raw log

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
func (it *IERC20BridgeCounterpartTokenDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20BridgeCounterpartTokenDeployed)
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
		it.Event = new(IERC20BridgeCounterpartTokenDeployed)
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
func (it *IERC20BridgeCounterpartTokenDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20BridgeCounterpartTokenDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20BridgeCounterpartTokenDeployed represents a CounterpartTokenDeployed event raised by the IERC20Bridge contract.
type IERC20BridgeCounterpartTokenDeployed struct {
	Id            [32]byte
	Description   IERC20BridgeTokenDescription
	DeployedToken common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCounterpartTokenDeployed is a free log retrieval operation binding the contract event 0x7f7a756c4acefa321f2d647bc59682a7a9bd5f6c11e8395540a3cdf69cca609c.
//
// Solidity: event CounterpartTokenDeployed(bytes32 indexed id, (address,string,string,uint8) description, address indexed deployedToken)
func (_IERC20Bridge *IERC20BridgeFilterer) FilterCounterpartTokenDeployed(opts *bind.FilterOpts, id [][32]byte, deployedToken []common.Address) (*IERC20BridgeCounterpartTokenDeployedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var deployedTokenRule []interface{}
	for _, deployedTokenItem := range deployedToken {
		deployedTokenRule = append(deployedTokenRule, deployedTokenItem)
	}

	logs, sub, err := _IERC20Bridge.contract.FilterLogs(opts, "CounterpartTokenDeployed", idRule, deployedTokenRule)
	if err != nil {
		return nil, err
	}
	return &IERC20BridgeCounterpartTokenDeployedIterator{contract: _IERC20Bridge.contract, event: "CounterpartTokenDeployed", logs: logs, sub: sub}, nil
}

// WatchCounterpartTokenDeployed is a free log subscription operation binding the contract event 0x7f7a756c4acefa321f2d647bc59682a7a9bd5f6c11e8395540a3cdf69cca609c.
//
// Solidity: event CounterpartTokenDeployed(bytes32 indexed id, (address,string,string,uint8) description, address indexed deployedToken)
func (_IERC20Bridge *IERC20BridgeFilterer) WatchCounterpartTokenDeployed(opts *bind.WatchOpts, sink chan<- *IERC20BridgeCounterpartTokenDeployed, id [][32]byte, deployedToken []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var deployedTokenRule []interface{}
	for _, deployedTokenItem := range deployedToken {
		deployedTokenRule = append(deployedTokenRule, deployedTokenItem)
	}

	logs, sub, err := _IERC20Bridge.contract.WatchLogs(opts, "CounterpartTokenDeployed", idRule, deployedTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20BridgeCounterpartTokenDeployed)
				if err := _IERC20Bridge.contract.UnpackLog(event, "CounterpartTokenDeployed", log); err != nil {
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

// ParseCounterpartTokenDeployed is a log parse operation binding the contract event 0x7f7a756c4acefa321f2d647bc59682a7a9bd5f6c11e8395540a3cdf69cca609c.
//
// Solidity: event CounterpartTokenDeployed(bytes32 indexed id, (address,string,string,uint8) description, address indexed deployedToken)
func (_IERC20Bridge *IERC20BridgeFilterer) ParseCounterpartTokenDeployed(log types.Log) (*IERC20BridgeCounterpartTokenDeployed, error) {
	event := new(IERC20BridgeCounterpartTokenDeployed)
	if err := _IERC20Bridge.contract.UnpackLog(event, "CounterpartTokenDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20BridgeDepositClaimedIterator is returned from FilterDepositClaimed and is used to iterate over the raw logs and unpacked data for DepositClaimed events raised by the IERC20Bridge contract.
type IERC20BridgeDepositClaimedIterator struct {
	Event *IERC20BridgeDepositClaimed // Event containing the contract specifics and raw log

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
func (it *IERC20BridgeDepositClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20BridgeDepositClaimed)
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
		it.Event = new(IERC20BridgeDepositClaimed)
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
func (it *IERC20BridgeDepositClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20BridgeDepositClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20BridgeDepositClaimed represents a DepositClaimed event raised by the IERC20Bridge contract.
type IERC20BridgeDepositClaimed struct {
	Id      [32]byte
	Deposit IERC20BridgeERC20Deposit
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositClaimed is a free log retrieval operation binding the contract event 0x75ca861d9a16245a1c9752ca37595c5077f967df5bdb151249f969d5280d3d42.
//
// Solidity: event DepositClaimed(bytes32 indexed id, (uint256,address,address,address,uint256) deposit)
func (_IERC20Bridge *IERC20BridgeFilterer) FilterDepositClaimed(opts *bind.FilterOpts, id [][32]byte) (*IERC20BridgeDepositClaimedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC20Bridge.contract.FilterLogs(opts, "DepositClaimed", idRule)
	if err != nil {
		return nil, err
	}
	return &IERC20BridgeDepositClaimedIterator{contract: _IERC20Bridge.contract, event: "DepositClaimed", logs: logs, sub: sub}, nil
}

// WatchDepositClaimed is a free log subscription operation binding the contract event 0x75ca861d9a16245a1c9752ca37595c5077f967df5bdb151249f969d5280d3d42.
//
// Solidity: event DepositClaimed(bytes32 indexed id, (uint256,address,address,address,uint256) deposit)
func (_IERC20Bridge *IERC20BridgeFilterer) WatchDepositClaimed(opts *bind.WatchOpts, sink chan<- *IERC20BridgeDepositClaimed, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC20Bridge.contract.WatchLogs(opts, "DepositClaimed", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20BridgeDepositClaimed)
				if err := _IERC20Bridge.contract.UnpackLog(event, "DepositClaimed", log); err != nil {
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

// ParseDepositClaimed is a log parse operation binding the contract event 0x75ca861d9a16245a1c9752ca37595c5077f967df5bdb151249f969d5280d3d42.
//
// Solidity: event DepositClaimed(bytes32 indexed id, (uint256,address,address,address,uint256) deposit)
func (_IERC20Bridge *IERC20BridgeFilterer) ParseDepositClaimed(log types.Log) (*IERC20BridgeDepositClaimed, error) {
	event := new(IERC20BridgeDepositClaimed)
	if err := _IERC20Bridge.contract.UnpackLog(event, "DepositClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20BridgeDepositMadeIterator is returned from FilterDepositMade and is used to iterate over the raw logs and unpacked data for DepositMade events raised by the IERC20Bridge contract.
type IERC20BridgeDepositMadeIterator struct {
	Event *IERC20BridgeDepositMade // Event containing the contract specifics and raw log

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
func (it *IERC20BridgeDepositMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20BridgeDepositMade)
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
		it.Event = new(IERC20BridgeDepositMade)
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
func (it *IERC20BridgeDepositMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20BridgeDepositMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20BridgeDepositMade represents a DepositMade event raised by the IERC20Bridge contract.
type IERC20BridgeDepositMade struct {
	Id         [32]byte
	Deposit    IERC20BridgeERC20Deposit
	LocalToken common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositMade is a free log retrieval operation binding the contract event 0x1004487d12f34ac75f1b4c5a8f3e4192362030f580a7646f6345e41d2e7eb30c.
//
// Solidity: event DepositMade(bytes32 indexed id, (uint256,address,address,address,uint256) deposit, address localToken)
func (_IERC20Bridge *IERC20BridgeFilterer) FilterDepositMade(opts *bind.FilterOpts, id [][32]byte) (*IERC20BridgeDepositMadeIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC20Bridge.contract.FilterLogs(opts, "DepositMade", idRule)
	if err != nil {
		return nil, err
	}
	return &IERC20BridgeDepositMadeIterator{contract: _IERC20Bridge.contract, event: "DepositMade", logs: logs, sub: sub}, nil
}

// WatchDepositMade is a free log subscription operation binding the contract event 0x1004487d12f34ac75f1b4c5a8f3e4192362030f580a7646f6345e41d2e7eb30c.
//
// Solidity: event DepositMade(bytes32 indexed id, (uint256,address,address,address,uint256) deposit, address localToken)
func (_IERC20Bridge *IERC20BridgeFilterer) WatchDepositMade(opts *bind.WatchOpts, sink chan<- *IERC20BridgeDepositMade, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC20Bridge.contract.WatchLogs(opts, "DepositMade", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20BridgeDepositMade)
				if err := _IERC20Bridge.contract.UnpackLog(event, "DepositMade", log); err != nil {
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

// ParseDepositMade is a log parse operation binding the contract event 0x1004487d12f34ac75f1b4c5a8f3e4192362030f580a7646f6345e41d2e7eb30c.
//
// Solidity: event DepositMade(bytes32 indexed id, (uint256,address,address,address,uint256) deposit, address localToken)
func (_IERC20Bridge *IERC20BridgeFilterer) ParseDepositMade(log types.Log) (*IERC20BridgeDepositMade, error) {
	event := new(IERC20BridgeDepositMade)
	if err := _IERC20Bridge.contract.UnpackLog(event, "DepositMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20BridgeTokenDescriptionRecordedIterator is returned from FilterTokenDescriptionRecorded and is used to iterate over the raw logs and unpacked data for TokenDescriptionRecorded events raised by the IERC20Bridge contract.
type IERC20BridgeTokenDescriptionRecordedIterator struct {
	Event *IERC20BridgeTokenDescriptionRecorded // Event containing the contract specifics and raw log

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
func (it *IERC20BridgeTokenDescriptionRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20BridgeTokenDescriptionRecorded)
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
		it.Event = new(IERC20BridgeTokenDescriptionRecorded)
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
func (it *IERC20BridgeTokenDescriptionRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20BridgeTokenDescriptionRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20BridgeTokenDescriptionRecorded represents a TokenDescriptionRecorded event raised by the IERC20Bridge contract.
type IERC20BridgeTokenDescriptionRecorded struct {
	Id          [32]byte
	Description IERC20BridgeTokenDescription
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenDescriptionRecorded is a free log retrieval operation binding the contract event 0x4b8cb4433f787c2a1c1113d60b44206f33c34dd8dc3287f4e4d82f1623adf53e.
//
// Solidity: event TokenDescriptionRecorded(bytes32 indexed id, (address,string,string,uint8) description)
func (_IERC20Bridge *IERC20BridgeFilterer) FilterTokenDescriptionRecorded(opts *bind.FilterOpts, id [][32]byte) (*IERC20BridgeTokenDescriptionRecordedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC20Bridge.contract.FilterLogs(opts, "TokenDescriptionRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return &IERC20BridgeTokenDescriptionRecordedIterator{contract: _IERC20Bridge.contract, event: "TokenDescriptionRecorded", logs: logs, sub: sub}, nil
}

// WatchTokenDescriptionRecorded is a free log subscription operation binding the contract event 0x4b8cb4433f787c2a1c1113d60b44206f33c34dd8dc3287f4e4d82f1623adf53e.
//
// Solidity: event TokenDescriptionRecorded(bytes32 indexed id, (address,string,string,uint8) description)
func (_IERC20Bridge *IERC20BridgeFilterer) WatchTokenDescriptionRecorded(opts *bind.WatchOpts, sink chan<- *IERC20BridgeTokenDescriptionRecorded, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC20Bridge.contract.WatchLogs(opts, "TokenDescriptionRecorded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20BridgeTokenDescriptionRecorded)
				if err := _IERC20Bridge.contract.UnpackLog(event, "TokenDescriptionRecorded", log); err != nil {
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

// ParseTokenDescriptionRecorded is a log parse operation binding the contract event 0x4b8cb4433f787c2a1c1113d60b44206f33c34dd8dc3287f4e4d82f1623adf53e.
//
// Solidity: event TokenDescriptionRecorded(bytes32 indexed id, (address,string,string,uint8) description)
func (_IERC20Bridge *IERC20BridgeFilterer) ParseTokenDescriptionRecorded(log types.Log) (*IERC20BridgeTokenDescriptionRecorded, error) {
	event := new(IERC20BridgeTokenDescriptionRecorded)
	if err := _IERC20Bridge.contract.UnpackLog(event, "TokenDescriptionRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
