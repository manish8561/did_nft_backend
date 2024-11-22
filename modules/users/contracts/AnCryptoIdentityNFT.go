// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AnCryptoIdentityNFT

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
)

// AnCryptoIdentityNFTMetaData contains all meta data concerning the AnCryptoIdentityNFT contract.
var AnCryptoIdentityNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"AddressUsername\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"BurnUsername\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_fees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressUsername\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"check\",\"type\":\"bool\"}],\"name\":\"changeBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"name\":\"changeFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"changeManger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"check\",\"type\":\"bool\"}],\"name\":\"changeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"}],\"name\":\"getTokenURIFromUsernameTokenId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getUsernameFromTokenId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"}],\"name\":\"safeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testing\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tokenURIUsername\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"updateTokenUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"}],\"name\":\"updateUsername\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"usernameAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"validateAlphabeticAndNumeric\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AnCryptoIdentityNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use AnCryptoIdentityNFTMetaData.ABI instead.
var AnCryptoIdentityNFTABI = AnCryptoIdentityNFTMetaData.ABI

// AnCryptoIdentityNFT is an auto generated Go binding around an Ethereum contract.
type AnCryptoIdentityNFT struct {
	AnCryptoIdentityNFTCaller     // Read-only binding to the contract
	AnCryptoIdentityNFTTransactor // Write-only binding to the contract
	AnCryptoIdentityNFTFilterer   // Log filterer for contract events
}

// AnCryptoIdentityNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type AnCryptoIdentityNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnCryptoIdentityNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AnCryptoIdentityNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnCryptoIdentityNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AnCryptoIdentityNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnCryptoIdentityNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AnCryptoIdentityNFTSession struct {
	Contract     *AnCryptoIdentityNFT // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AnCryptoIdentityNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AnCryptoIdentityNFTCallerSession struct {
	Contract *AnCryptoIdentityNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// AnCryptoIdentityNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AnCryptoIdentityNFTTransactorSession struct {
	Contract     *AnCryptoIdentityNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AnCryptoIdentityNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type AnCryptoIdentityNFTRaw struct {
	Contract *AnCryptoIdentityNFT // Generic contract binding to access the raw methods on
}

// AnCryptoIdentityNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AnCryptoIdentityNFTCallerRaw struct {
	Contract *AnCryptoIdentityNFTCaller // Generic read-only contract binding to access the raw methods on
}

// AnCryptoIdentityNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AnCryptoIdentityNFTTransactorRaw struct {
	Contract *AnCryptoIdentityNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAnCryptoIdentityNFT creates a new instance of AnCryptoIdentityNFT, bound to a specific deployed contract.
func NewAnCryptoIdentityNFT(address common.Address, backend bind.ContractBackend) (*AnCryptoIdentityNFT, error) {
	contract, err := bindAnCryptoIdentityNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AnCryptoIdentityNFT{AnCryptoIdentityNFTCaller: AnCryptoIdentityNFTCaller{contract: contract}, AnCryptoIdentityNFTTransactor: AnCryptoIdentityNFTTransactor{contract: contract}, AnCryptoIdentityNFTFilterer: AnCryptoIdentityNFTFilterer{contract: contract}}, nil
}

// NewAnCryptoIdentityNFTCaller creates a new read-only instance of AnCryptoIdentityNFT, bound to a specific deployed contract.
func NewAnCryptoIdentityNFTCaller(address common.Address, caller bind.ContractCaller) (*AnCryptoIdentityNFTCaller, error) {
	contract, err := bindAnCryptoIdentityNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AnCryptoIdentityNFTCaller{contract: contract}, nil
}

// NewAnCryptoIdentityNFTTransactor creates a new write-only instance of AnCryptoIdentityNFT, bound to a specific deployed contract.
func NewAnCryptoIdentityNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*AnCryptoIdentityNFTTransactor, error) {
	contract, err := bindAnCryptoIdentityNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AnCryptoIdentityNFTTransactor{contract: contract}, nil
}

// NewAnCryptoIdentityNFTFilterer creates a new log filterer instance of AnCryptoIdentityNFT, bound to a specific deployed contract.
func NewAnCryptoIdentityNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*AnCryptoIdentityNFTFilterer, error) {
	contract, err := bindAnCryptoIdentityNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AnCryptoIdentityNFTFilterer{contract: contract}, nil
}

// bindAnCryptoIdentityNFT binds a generic wrapper to an already deployed contract.
func bindAnCryptoIdentityNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AnCryptoIdentityNFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AnCryptoIdentityNFT.Contract.AnCryptoIdentityNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.AnCryptoIdentityNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.AnCryptoIdentityNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AnCryptoIdentityNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.contract.Transact(opts, method, params...)
}

// Fees is a free data retrieval call binding the contract method 0xd212a69a.
//
// Solidity: function _fees() view returns(uint256)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) Fees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "_fees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fees is a free data retrieval call binding the contract method 0xd212a69a.
//
// Solidity: function _fees() view returns(uint256)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) Fees() (*big.Int, error) {
	return _AnCryptoIdentityNFT.Contract.Fees(&_AnCryptoIdentityNFT.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0xd212a69a.
//
// Solidity: function _fees() view returns(uint256)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) Fees() (*big.Int, error) {
	return _AnCryptoIdentityNFT.Contract.Fees(&_AnCryptoIdentityNFT.CallOpts)
}

// AddressUsername is a free data retrieval call binding the contract method 0xf9c1ed02.
//
// Solidity: function addressUsername(address ) view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) AddressUsername(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "addressUsername", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AddressUsername is a free data retrieval call binding the contract method 0xf9c1ed02.
//
// Solidity: function addressUsername(address ) view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) AddressUsername(arg0 common.Address) (string, error) {
	return _AnCryptoIdentityNFT.Contract.AddressUsername(&_AnCryptoIdentityNFT.CallOpts, arg0)
}

// AddressUsername is a free data retrieval call binding the contract method 0xf9c1ed02.
//
// Solidity: function addressUsername(address ) view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) AddressUsername(arg0 common.Address) (string, error) {
	return _AnCryptoIdentityNFT.Contract.AddressUsername(&_AnCryptoIdentityNFT.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _AnCryptoIdentityNFT.Contract.BalanceOf(&_AnCryptoIdentityNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _AnCryptoIdentityNFT.Contract.BalanceOf(&_AnCryptoIdentityNFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _AnCryptoIdentityNFT.Contract.GetApproved(&_AnCryptoIdentityNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _AnCryptoIdentityNFT.Contract.GetApproved(&_AnCryptoIdentityNFT.CallOpts, tokenId)
}

// GetTokenURIFromUsernameTokenId is a free data retrieval call binding the contract method 0x80440574.
//
// Solidity: function getTokenURIFromUsernameTokenId(string username) view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) GetTokenURIFromUsernameTokenId(opts *bind.CallOpts, username string) (string, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "getTokenURIFromUsernameTokenId", username)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTokenURIFromUsernameTokenId is a free data retrieval call binding the contract method 0x80440574.
//
// Solidity: function getTokenURIFromUsernameTokenId(string username) view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) GetTokenURIFromUsernameTokenId(username string) (string, error) {
	return _AnCryptoIdentityNFT.Contract.GetTokenURIFromUsernameTokenId(&_AnCryptoIdentityNFT.CallOpts, username)
}

// GetTokenURIFromUsernameTokenId is a free data retrieval call binding the contract method 0x80440574.
//
// Solidity: function getTokenURIFromUsernameTokenId(string username) view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) GetTokenURIFromUsernameTokenId(username string) (string, error) {
	return _AnCryptoIdentityNFT.Contract.GetTokenURIFromUsernameTokenId(&_AnCryptoIdentityNFT.CallOpts, username)
}

// GetUsernameFromTokenId is a free data retrieval call binding the contract method 0x5b082438.
//
// Solidity: function getUsernameFromTokenId(uint256 tokenId) view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) GetUsernameFromTokenId(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "getUsernameFromTokenId", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUsernameFromTokenId is a free data retrieval call binding the contract method 0x5b082438.
//
// Solidity: function getUsernameFromTokenId(uint256 tokenId) view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) GetUsernameFromTokenId(tokenId *big.Int) (string, error) {
	return _AnCryptoIdentityNFT.Contract.GetUsernameFromTokenId(&_AnCryptoIdentityNFT.CallOpts, tokenId)
}

// GetUsernameFromTokenId is a free data retrieval call binding the contract method 0x5b082438.
//
// Solidity: function getUsernameFromTokenId(uint256 tokenId) view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) GetUsernameFromTokenId(tokenId *big.Int) (string, error) {
	return _AnCryptoIdentityNFT.Contract.GetUsernameFromTokenId(&_AnCryptoIdentityNFT.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _AnCryptoIdentityNFT.Contract.IsApprovedForAll(&_AnCryptoIdentityNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _AnCryptoIdentityNFT.Contract.IsApprovedForAll(&_AnCryptoIdentityNFT.CallOpts, owner, operator)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) Manager() (common.Address, error) {
	return _AnCryptoIdentityNFT.Contract.Manager(&_AnCryptoIdentityNFT.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) Manager() (common.Address, error) {
	return _AnCryptoIdentityNFT.Contract.Manager(&_AnCryptoIdentityNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) Name() (string, error) {
	return _AnCryptoIdentityNFT.Contract.Name(&_AnCryptoIdentityNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) Name() (string, error) {
	return _AnCryptoIdentityNFT.Contract.Name(&_AnCryptoIdentityNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) Owner() (common.Address, error) {
	return _AnCryptoIdentityNFT.Contract.Owner(&_AnCryptoIdentityNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) Owner() (common.Address, error) {
	return _AnCryptoIdentityNFT.Contract.Owner(&_AnCryptoIdentityNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _AnCryptoIdentityNFT.Contract.OwnerOf(&_AnCryptoIdentityNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _AnCryptoIdentityNFT.Contract.OwnerOf(&_AnCryptoIdentityNFT.CallOpts, tokenId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) ProxiableUUID() ([32]byte, error) {
	return _AnCryptoIdentityNFT.Contract.ProxiableUUID(&_AnCryptoIdentityNFT.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) ProxiableUUID() ([32]byte, error) {
	return _AnCryptoIdentityNFT.Contract.ProxiableUUID(&_AnCryptoIdentityNFT.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AnCryptoIdentityNFT.Contract.SupportsInterface(&_AnCryptoIdentityNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AnCryptoIdentityNFT.Contract.SupportsInterface(&_AnCryptoIdentityNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) Symbol() (string, error) {
	return _AnCryptoIdentityNFT.Contract.Symbol(&_AnCryptoIdentityNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) Symbol() (string, error) {
	return _AnCryptoIdentityNFT.Contract.Symbol(&_AnCryptoIdentityNFT.CallOpts)
}

// Testing is a free data retrieval call binding the contract method 0x8d03b102.
//
// Solidity: function testing() pure returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) Testing(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "testing")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Testing is a free data retrieval call binding the contract method 0x8d03b102.
//
// Solidity: function testing() pure returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) Testing() (string, error) {
	return _AnCryptoIdentityNFT.Contract.Testing(&_AnCryptoIdentityNFT.CallOpts)
}

// Testing is a free data retrieval call binding the contract method 0x8d03b102.
//
// Solidity: function testing() pure returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) Testing() (string, error) {
	return _AnCryptoIdentityNFT.Contract.Testing(&_AnCryptoIdentityNFT.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _AnCryptoIdentityNFT.Contract.TokenByIndex(&_AnCryptoIdentityNFT.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _AnCryptoIdentityNFT.Contract.TokenByIndex(&_AnCryptoIdentityNFT.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _AnCryptoIdentityNFT.Contract.TokenOfOwnerByIndex(&_AnCryptoIdentityNFT.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _AnCryptoIdentityNFT.Contract.TokenOfOwnerByIndex(&_AnCryptoIdentityNFT.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _AnCryptoIdentityNFT.Contract.TokenURI(&_AnCryptoIdentityNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _AnCryptoIdentityNFT.Contract.TokenURI(&_AnCryptoIdentityNFT.CallOpts, tokenId)
}

// TokenURIUsername is a free data retrieval call binding the contract method 0x164e46e8.
//
// Solidity: function tokenURIUsername(string ) view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) TokenURIUsername(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "tokenURIUsername", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURIUsername is a free data retrieval call binding the contract method 0x164e46e8.
//
// Solidity: function tokenURIUsername(string ) view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) TokenURIUsername(arg0 string) (string, error) {
	return _AnCryptoIdentityNFT.Contract.TokenURIUsername(&_AnCryptoIdentityNFT.CallOpts, arg0)
}

// TokenURIUsername is a free data retrieval call binding the contract method 0x164e46e8.
//
// Solidity: function tokenURIUsername(string ) view returns(string)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) TokenURIUsername(arg0 string) (string, error) {
	return _AnCryptoIdentityNFT.Contract.TokenURIUsername(&_AnCryptoIdentityNFT.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) TotalSupply() (*big.Int, error) {
	return _AnCryptoIdentityNFT.Contract.TotalSupply(&_AnCryptoIdentityNFT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) TotalSupply() (*big.Int, error) {
	return _AnCryptoIdentityNFT.Contract.TotalSupply(&_AnCryptoIdentityNFT.CallOpts)
}

// UsernameAddress is a free data retrieval call binding the contract method 0xf8831954.
//
// Solidity: function usernameAddress(string ) view returns(address)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) UsernameAddress(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "usernameAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsernameAddress is a free data retrieval call binding the contract method 0xf8831954.
//
// Solidity: function usernameAddress(string ) view returns(address)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) UsernameAddress(arg0 string) (common.Address, error) {
	return _AnCryptoIdentityNFT.Contract.UsernameAddress(&_AnCryptoIdentityNFT.CallOpts, arg0)
}

// UsernameAddress is a free data retrieval call binding the contract method 0xf8831954.
//
// Solidity: function usernameAddress(string ) view returns(address)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) UsernameAddress(arg0 string) (common.Address, error) {
	return _AnCryptoIdentityNFT.Contract.UsernameAddress(&_AnCryptoIdentityNFT.CallOpts, arg0)
}

// ValidateAlphabeticAndNumeric is a free data retrieval call binding the contract method 0xd685c3bd.
//
// Solidity: function validateAlphabeticAndNumeric(string str) pure returns(bool)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCaller) ValidateAlphabeticAndNumeric(opts *bind.CallOpts, str string) (bool, error) {
	var out []interface{}
	err := _AnCryptoIdentityNFT.contract.Call(opts, &out, "validateAlphabeticAndNumeric", str)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateAlphabeticAndNumeric is a free data retrieval call binding the contract method 0xd685c3bd.
//
// Solidity: function validateAlphabeticAndNumeric(string str) pure returns(bool)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) ValidateAlphabeticAndNumeric(str string) (bool, error) {
	return _AnCryptoIdentityNFT.Contract.ValidateAlphabeticAndNumeric(&_AnCryptoIdentityNFT.CallOpts, str)
}

// ValidateAlphabeticAndNumeric is a free data retrieval call binding the contract method 0xd685c3bd.
//
// Solidity: function validateAlphabeticAndNumeric(string str) pure returns(bool)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTCallerSession) ValidateAlphabeticAndNumeric(str string) (bool, error) {
	return _AnCryptoIdentityNFT.Contract.ValidateAlphabeticAndNumeric(&_AnCryptoIdentityNFT.CallOpts, str)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.Approve(&_AnCryptoIdentityNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.Approve(&_AnCryptoIdentityNFT.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.Burn(&_AnCryptoIdentityNFT.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.Burn(&_AnCryptoIdentityNFT.TransactOpts, tokenId)
}

// ChangeBurn is a paid mutator transaction binding the contract method 0x7e0c00b2.
//
// Solidity: function changeBurn(bool check) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactor) ChangeBurn(opts *bind.TransactOpts, check bool) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.contract.Transact(opts, "changeBurn", check)
}

// ChangeBurn is a paid mutator transaction binding the contract method 0x7e0c00b2.
//
// Solidity: function changeBurn(bool check) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) ChangeBurn(check bool) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.ChangeBurn(&_AnCryptoIdentityNFT.TransactOpts, check)
}

// ChangeBurn is a paid mutator transaction binding the contract method 0x7e0c00b2.
//
// Solidity: function changeBurn(bool check) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorSession) ChangeBurn(check bool) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.ChangeBurn(&_AnCryptoIdentityNFT.TransactOpts, check)
}

// ChangeFees is a paid mutator transaction binding the contract method 0x6cda375b.
//
// Solidity: function changeFees(uint256 fees) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactor) ChangeFees(opts *bind.TransactOpts, fees *big.Int) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.contract.Transact(opts, "changeFees", fees)
}

// ChangeFees is a paid mutator transaction binding the contract method 0x6cda375b.
//
// Solidity: function changeFees(uint256 fees) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) ChangeFees(fees *big.Int) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.ChangeFees(&_AnCryptoIdentityNFT.TransactOpts, fees)
}

// ChangeFees is a paid mutator transaction binding the contract method 0x6cda375b.
//
// Solidity: function changeFees(uint256 fees) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorSession) ChangeFees(fees *big.Int) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.ChangeFees(&_AnCryptoIdentityNFT.TransactOpts, fees)
}

// ChangeManger is a paid mutator transaction binding the contract method 0x96754863.
//
// Solidity: function changeManger(address _manager) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactor) ChangeManger(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.contract.Transact(opts, "changeManger", _manager)
}

// ChangeManger is a paid mutator transaction binding the contract method 0x96754863.
//
// Solidity: function changeManger(address _manager) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) ChangeManger(_manager common.Address) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.ChangeManger(&_AnCryptoIdentityNFT.TransactOpts, _manager)
}

// ChangeManger is a paid mutator transaction binding the contract method 0x96754863.
//
// Solidity: function changeManger(address _manager) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorSession) ChangeManger(_manager common.Address) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.ChangeManger(&_AnCryptoIdentityNFT.TransactOpts, _manager)
}

// ChangeTransferFrom is a paid mutator transaction binding the contract method 0x75a1f000.
//
// Solidity: function changeTransferFrom(bool check) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactor) ChangeTransferFrom(opts *bind.TransactOpts, check bool) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.contract.Transact(opts, "changeTransferFrom", check)
}

// ChangeTransferFrom is a paid mutator transaction binding the contract method 0x75a1f000.
//
// Solidity: function changeTransferFrom(bool check) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) ChangeTransferFrom(check bool) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.ChangeTransferFrom(&_AnCryptoIdentityNFT.TransactOpts, check)
}

// ChangeTransferFrom is a paid mutator transaction binding the contract method 0x75a1f000.
//
// Solidity: function changeTransferFrom(bool check) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorSession) ChangeTransferFrom(check bool) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.ChangeTransferFrom(&_AnCryptoIdentityNFT.TransactOpts, check)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address _manager) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, _manager common.Address) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.contract.Transact(opts, "initialize", name, symbol, _manager)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address _manager) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) Initialize(name string, symbol string, _manager common.Address) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.Initialize(&_AnCryptoIdentityNFT.TransactOpts, name, symbol, _manager)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address _manager) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorSession) Initialize(name string, symbol string, _manager common.Address) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.Initialize(&_AnCryptoIdentityNFT.TransactOpts, name, symbol, _manager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.RenounceOwnership(&_AnCryptoIdentityNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.RenounceOwnership(&_AnCryptoIdentityNFT.TransactOpts)
}

// SafeMint is a paid mutator transaction binding the contract method 0xeca81d42.
//
// Solidity: function safeMint(address to, string uri, string username) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactor) SafeMint(opts *bind.TransactOpts, to common.Address, uri string, username string) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.contract.Transact(opts, "safeMint", to, uri, username)
}

// SafeMint is a paid mutator transaction binding the contract method 0xeca81d42.
//
// Solidity: function safeMint(address to, string uri, string username) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) SafeMint(to common.Address, uri string, username string) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.SafeMint(&_AnCryptoIdentityNFT.TransactOpts, to, uri, username)
}

// SafeMint is a paid mutator transaction binding the contract method 0xeca81d42.
//
// Solidity: function safeMint(address to, string uri, string username) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorSession) SafeMint(to common.Address, uri string, username string) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.SafeMint(&_AnCryptoIdentityNFT.TransactOpts, to, uri, username)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.SafeTransferFrom(&_AnCryptoIdentityNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.SafeTransferFrom(&_AnCryptoIdentityNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.SafeTransferFrom0(&_AnCryptoIdentityNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.SafeTransferFrom0(&_AnCryptoIdentityNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.SetApprovalForAll(&_AnCryptoIdentityNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.SetApprovalForAll(&_AnCryptoIdentityNFT.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.TransferFrom(&_AnCryptoIdentityNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.TransferFrom(&_AnCryptoIdentityNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.TransferOwnership(&_AnCryptoIdentityNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.TransferOwnership(&_AnCryptoIdentityNFT.TransactOpts, newOwner)
}

// UpdateTokenUri is a paid mutator transaction binding the contract method 0xfd34019a.
//
// Solidity: function updateTokenUri(string username, string uri) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactor) UpdateTokenUri(opts *bind.TransactOpts, username string, uri string) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.contract.Transact(opts, "updateTokenUri", username, uri)
}

// UpdateTokenUri is a paid mutator transaction binding the contract method 0xfd34019a.
//
// Solidity: function updateTokenUri(string username, string uri) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) UpdateTokenUri(username string, uri string) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.UpdateTokenUri(&_AnCryptoIdentityNFT.TransactOpts, username, uri)
}

// UpdateTokenUri is a paid mutator transaction binding the contract method 0xfd34019a.
//
// Solidity: function updateTokenUri(string username, string uri) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorSession) UpdateTokenUri(username string, uri string) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.UpdateTokenUri(&_AnCryptoIdentityNFT.TransactOpts, username, uri)
}

// UpdateUsername is a paid mutator transaction binding the contract method 0xc96cea70.
//
// Solidity: function updateUsername(string username) payable returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactor) UpdateUsername(opts *bind.TransactOpts, username string) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.contract.Transact(opts, "updateUsername", username)
}

// UpdateUsername is a paid mutator transaction binding the contract method 0xc96cea70.
//
// Solidity: function updateUsername(string username) payable returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) UpdateUsername(username string) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.UpdateUsername(&_AnCryptoIdentityNFT.TransactOpts, username)
}

// UpdateUsername is a paid mutator transaction binding the contract method 0xc96cea70.
//
// Solidity: function updateUsername(string username) payable returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorSession) UpdateUsername(username string) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.UpdateUsername(&_AnCryptoIdentityNFT.TransactOpts, username)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.UpgradeTo(&_AnCryptoIdentityNFT.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.UpgradeTo(&_AnCryptoIdentityNFT.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.UpgradeToAndCall(&_AnCryptoIdentityNFT.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.UpgradeToAndCall(&_AnCryptoIdentityNFT.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTSession) Receive() (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.Receive(&_AnCryptoIdentityNFT.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTTransactorSession) Receive() (*types.Transaction, error) {
	return _AnCryptoIdentityNFT.Contract.Receive(&_AnCryptoIdentityNFT.TransactOpts)
}

// AnCryptoIdentityNFTAddressUsernameIterator is returned from FilterAddressUsername and is used to iterate over the raw logs and unpacked data for AddressUsername events raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTAddressUsernameIterator struct {
	Event *AnCryptoIdentityNFTAddressUsername // Event containing the contract specifics and raw log

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
func (it *AnCryptoIdentityNFTAddressUsernameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnCryptoIdentityNFTAddressUsername)
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
		it.Event = new(AnCryptoIdentityNFTAddressUsername)
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
func (it *AnCryptoIdentityNFTAddressUsernameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnCryptoIdentityNFTAddressUsernameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnCryptoIdentityNFTAddressUsername represents a AddressUsername event raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTAddressUsername struct {
	To       common.Address
	Username common.Hash
	TokenId  *big.Int
	TokenURI string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddressUsername is a free log retrieval operation binding the contract event 0xb5e14b6423ef08d2dfd10a8d560e429adbdf970ea9c0b426297d7e55ed2f562b.
//
// Solidity: event AddressUsername(address indexed to, string indexed username, uint256 indexed tokenId, string tokenURI)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) FilterAddressUsername(opts *bind.FilterOpts, to []common.Address, username []string, tokenId []*big.Int) (*AnCryptoIdentityNFTAddressUsernameIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AnCryptoIdentityNFT.contract.FilterLogs(opts, "AddressUsername", toRule, usernameRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AnCryptoIdentityNFTAddressUsernameIterator{contract: _AnCryptoIdentityNFT.contract, event: "AddressUsername", logs: logs, sub: sub}, nil
}

// WatchAddressUsername is a free log subscription operation binding the contract event 0xb5e14b6423ef08d2dfd10a8d560e429adbdf970ea9c0b426297d7e55ed2f562b.
//
// Solidity: event AddressUsername(address indexed to, string indexed username, uint256 indexed tokenId, string tokenURI)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) WatchAddressUsername(opts *bind.WatchOpts, sink chan<- *AnCryptoIdentityNFTAddressUsername, to []common.Address, username []string, tokenId []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AnCryptoIdentityNFT.contract.WatchLogs(opts, "AddressUsername", toRule, usernameRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnCryptoIdentityNFTAddressUsername)
				if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "AddressUsername", log); err != nil {
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

// ParseAddressUsername is a log parse operation binding the contract event 0xb5e14b6423ef08d2dfd10a8d560e429adbdf970ea9c0b426297d7e55ed2f562b.
//
// Solidity: event AddressUsername(address indexed to, string indexed username, uint256 indexed tokenId, string tokenURI)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) ParseAddressUsername(log types.Log) (*AnCryptoIdentityNFTAddressUsername, error) {
	event := new(AnCryptoIdentityNFTAddressUsername)
	if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "AddressUsername", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnCryptoIdentityNFTAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTAdminChangedIterator struct {
	Event *AnCryptoIdentityNFTAdminChanged // Event containing the contract specifics and raw log

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
func (it *AnCryptoIdentityNFTAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnCryptoIdentityNFTAdminChanged)
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
		it.Event = new(AnCryptoIdentityNFTAdminChanged)
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
func (it *AnCryptoIdentityNFTAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnCryptoIdentityNFTAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnCryptoIdentityNFTAdminChanged represents a AdminChanged event raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AnCryptoIdentityNFTAdminChangedIterator, error) {

	logs, sub, err := _AnCryptoIdentityNFT.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AnCryptoIdentityNFTAdminChangedIterator{contract: _AnCryptoIdentityNFT.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AnCryptoIdentityNFTAdminChanged) (event.Subscription, error) {

	logs, sub, err := _AnCryptoIdentityNFT.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnCryptoIdentityNFTAdminChanged)
				if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) ParseAdminChanged(log types.Log) (*AnCryptoIdentityNFTAdminChanged, error) {
	event := new(AnCryptoIdentityNFTAdminChanged)
	if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnCryptoIdentityNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTApprovalIterator struct {
	Event *AnCryptoIdentityNFTApproval // Event containing the contract specifics and raw log

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
func (it *AnCryptoIdentityNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnCryptoIdentityNFTApproval)
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
		it.Event = new(AnCryptoIdentityNFTApproval)
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
func (it *AnCryptoIdentityNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnCryptoIdentityNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnCryptoIdentityNFTApproval represents a Approval event raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*AnCryptoIdentityNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AnCryptoIdentityNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AnCryptoIdentityNFTApprovalIterator{contract: _AnCryptoIdentityNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AnCryptoIdentityNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AnCryptoIdentityNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnCryptoIdentityNFTApproval)
				if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) ParseApproval(log types.Log) (*AnCryptoIdentityNFTApproval, error) {
	event := new(AnCryptoIdentityNFTApproval)
	if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnCryptoIdentityNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTApprovalForAllIterator struct {
	Event *AnCryptoIdentityNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *AnCryptoIdentityNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnCryptoIdentityNFTApprovalForAll)
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
		it.Event = new(AnCryptoIdentityNFTApprovalForAll)
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
func (it *AnCryptoIdentityNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnCryptoIdentityNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnCryptoIdentityNFTApprovalForAll represents a ApprovalForAll event raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*AnCryptoIdentityNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AnCryptoIdentityNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &AnCryptoIdentityNFTApprovalForAllIterator{contract: _AnCryptoIdentityNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *AnCryptoIdentityNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AnCryptoIdentityNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnCryptoIdentityNFTApprovalForAll)
				if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) ParseApprovalForAll(log types.Log) (*AnCryptoIdentityNFTApprovalForAll, error) {
	event := new(AnCryptoIdentityNFTApprovalForAll)
	if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnCryptoIdentityNFTBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTBatchMetadataUpdateIterator struct {
	Event *AnCryptoIdentityNFTBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *AnCryptoIdentityNFTBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnCryptoIdentityNFTBatchMetadataUpdate)
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
		it.Event = new(AnCryptoIdentityNFTBatchMetadataUpdate)
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
func (it *AnCryptoIdentityNFTBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnCryptoIdentityNFTBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnCryptoIdentityNFTBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*AnCryptoIdentityNFTBatchMetadataUpdateIterator, error) {

	logs, sub, err := _AnCryptoIdentityNFT.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &AnCryptoIdentityNFTBatchMetadataUpdateIterator{contract: _AnCryptoIdentityNFT.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *AnCryptoIdentityNFTBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _AnCryptoIdentityNFT.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnCryptoIdentityNFTBatchMetadataUpdate)
				if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) ParseBatchMetadataUpdate(log types.Log) (*AnCryptoIdentityNFTBatchMetadataUpdate, error) {
	event := new(AnCryptoIdentityNFTBatchMetadataUpdate)
	if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnCryptoIdentityNFTBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTBeaconUpgradedIterator struct {
	Event *AnCryptoIdentityNFTBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *AnCryptoIdentityNFTBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnCryptoIdentityNFTBeaconUpgraded)
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
		it.Event = new(AnCryptoIdentityNFTBeaconUpgraded)
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
func (it *AnCryptoIdentityNFTBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnCryptoIdentityNFTBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnCryptoIdentityNFTBeaconUpgraded represents a BeaconUpgraded event raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*AnCryptoIdentityNFTBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _AnCryptoIdentityNFT.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &AnCryptoIdentityNFTBeaconUpgradedIterator{contract: _AnCryptoIdentityNFT.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *AnCryptoIdentityNFTBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _AnCryptoIdentityNFT.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnCryptoIdentityNFTBeaconUpgraded)
				if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) ParseBeaconUpgraded(log types.Log) (*AnCryptoIdentityNFTBeaconUpgraded, error) {
	event := new(AnCryptoIdentityNFTBeaconUpgraded)
	if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnCryptoIdentityNFTBurnUsernameIterator is returned from FilterBurnUsername and is used to iterate over the raw logs and unpacked data for BurnUsername events raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTBurnUsernameIterator struct {
	Event *AnCryptoIdentityNFTBurnUsername // Event containing the contract specifics and raw log

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
func (it *AnCryptoIdentityNFTBurnUsernameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnCryptoIdentityNFTBurnUsername)
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
		it.Event = new(AnCryptoIdentityNFTBurnUsername)
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
func (it *AnCryptoIdentityNFTBurnUsernameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnCryptoIdentityNFTBurnUsernameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnCryptoIdentityNFTBurnUsername represents a BurnUsername event raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTBurnUsername struct {
	To       common.Address
	Username common.Hash
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBurnUsername is a free log retrieval operation binding the contract event 0x056a1aa01da06f3501b430a5ebc166606fdc68356fe0c7bbd058d1b11e11705e.
//
// Solidity: event BurnUsername(address indexed to, string indexed username, uint256 indexed tokenId)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) FilterBurnUsername(opts *bind.FilterOpts, to []common.Address, username []string, tokenId []*big.Int) (*AnCryptoIdentityNFTBurnUsernameIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AnCryptoIdentityNFT.contract.FilterLogs(opts, "BurnUsername", toRule, usernameRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AnCryptoIdentityNFTBurnUsernameIterator{contract: _AnCryptoIdentityNFT.contract, event: "BurnUsername", logs: logs, sub: sub}, nil
}

// WatchBurnUsername is a free log subscription operation binding the contract event 0x056a1aa01da06f3501b430a5ebc166606fdc68356fe0c7bbd058d1b11e11705e.
//
// Solidity: event BurnUsername(address indexed to, string indexed username, uint256 indexed tokenId)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) WatchBurnUsername(opts *bind.WatchOpts, sink chan<- *AnCryptoIdentityNFTBurnUsername, to []common.Address, username []string, tokenId []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AnCryptoIdentityNFT.contract.WatchLogs(opts, "BurnUsername", toRule, usernameRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnCryptoIdentityNFTBurnUsername)
				if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "BurnUsername", log); err != nil {
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

// ParseBurnUsername is a log parse operation binding the contract event 0x056a1aa01da06f3501b430a5ebc166606fdc68356fe0c7bbd058d1b11e11705e.
//
// Solidity: event BurnUsername(address indexed to, string indexed username, uint256 indexed tokenId)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) ParseBurnUsername(log types.Log) (*AnCryptoIdentityNFTBurnUsername, error) {
	event := new(AnCryptoIdentityNFTBurnUsername)
	if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "BurnUsername", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnCryptoIdentityNFTInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTInitializedIterator struct {
	Event *AnCryptoIdentityNFTInitialized // Event containing the contract specifics and raw log

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
func (it *AnCryptoIdentityNFTInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnCryptoIdentityNFTInitialized)
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
		it.Event = new(AnCryptoIdentityNFTInitialized)
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
func (it *AnCryptoIdentityNFTInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnCryptoIdentityNFTInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnCryptoIdentityNFTInitialized represents a Initialized event raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) FilterInitialized(opts *bind.FilterOpts) (*AnCryptoIdentityNFTInitializedIterator, error) {

	logs, sub, err := _AnCryptoIdentityNFT.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AnCryptoIdentityNFTInitializedIterator{contract: _AnCryptoIdentityNFT.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AnCryptoIdentityNFTInitialized) (event.Subscription, error) {

	logs, sub, err := _AnCryptoIdentityNFT.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnCryptoIdentityNFTInitialized)
				if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) ParseInitialized(log types.Log) (*AnCryptoIdentityNFTInitialized, error) {
	event := new(AnCryptoIdentityNFTInitialized)
	if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnCryptoIdentityNFTMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTMetadataUpdateIterator struct {
	Event *AnCryptoIdentityNFTMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *AnCryptoIdentityNFTMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnCryptoIdentityNFTMetadataUpdate)
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
		it.Event = new(AnCryptoIdentityNFTMetadataUpdate)
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
func (it *AnCryptoIdentityNFTMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnCryptoIdentityNFTMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnCryptoIdentityNFTMetadataUpdate represents a MetadataUpdate event raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*AnCryptoIdentityNFTMetadataUpdateIterator, error) {

	logs, sub, err := _AnCryptoIdentityNFT.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &AnCryptoIdentityNFTMetadataUpdateIterator{contract: _AnCryptoIdentityNFT.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *AnCryptoIdentityNFTMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _AnCryptoIdentityNFT.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnCryptoIdentityNFTMetadataUpdate)
				if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) ParseMetadataUpdate(log types.Log) (*AnCryptoIdentityNFTMetadataUpdate, error) {
	event := new(AnCryptoIdentityNFTMetadataUpdate)
	if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnCryptoIdentityNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTOwnershipTransferredIterator struct {
	Event *AnCryptoIdentityNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AnCryptoIdentityNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnCryptoIdentityNFTOwnershipTransferred)
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
		it.Event = new(AnCryptoIdentityNFTOwnershipTransferred)
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
func (it *AnCryptoIdentityNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnCryptoIdentityNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnCryptoIdentityNFTOwnershipTransferred represents a OwnershipTransferred event raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AnCryptoIdentityNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AnCryptoIdentityNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AnCryptoIdentityNFTOwnershipTransferredIterator{contract: _AnCryptoIdentityNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AnCryptoIdentityNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AnCryptoIdentityNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnCryptoIdentityNFTOwnershipTransferred)
				if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) ParseOwnershipTransferred(log types.Log) (*AnCryptoIdentityNFTOwnershipTransferred, error) {
	event := new(AnCryptoIdentityNFTOwnershipTransferred)
	if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnCryptoIdentityNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTTransferIterator struct {
	Event *AnCryptoIdentityNFTTransfer // Event containing the contract specifics and raw log

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
func (it *AnCryptoIdentityNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnCryptoIdentityNFTTransfer)
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
		it.Event = new(AnCryptoIdentityNFTTransfer)
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
func (it *AnCryptoIdentityNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnCryptoIdentityNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnCryptoIdentityNFTTransfer represents a Transfer event raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*AnCryptoIdentityNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AnCryptoIdentityNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AnCryptoIdentityNFTTransferIterator{contract: _AnCryptoIdentityNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AnCryptoIdentityNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _AnCryptoIdentityNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnCryptoIdentityNFTTransfer)
				if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) ParseTransfer(log types.Log) (*AnCryptoIdentityNFTTransfer, error) {
	event := new(AnCryptoIdentityNFTTransfer)
	if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AnCryptoIdentityNFTUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTUpgradedIterator struct {
	Event *AnCryptoIdentityNFTUpgraded // Event containing the contract specifics and raw log

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
func (it *AnCryptoIdentityNFTUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnCryptoIdentityNFTUpgraded)
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
		it.Event = new(AnCryptoIdentityNFTUpgraded)
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
func (it *AnCryptoIdentityNFTUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnCryptoIdentityNFTUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnCryptoIdentityNFTUpgraded represents a Upgraded event raised by the AnCryptoIdentityNFT contract.
type AnCryptoIdentityNFTUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AnCryptoIdentityNFTUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AnCryptoIdentityNFT.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AnCryptoIdentityNFTUpgradedIterator{contract: _AnCryptoIdentityNFT.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AnCryptoIdentityNFTUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AnCryptoIdentityNFT.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnCryptoIdentityNFTUpgraded)
				if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AnCryptoIdentityNFT *AnCryptoIdentityNFTFilterer) ParseUpgraded(log types.Log) (*AnCryptoIdentityNFTUpgraded, error) {
	event := new(AnCryptoIdentityNFTUpgraded)
	if err := _AnCryptoIdentityNFT.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
