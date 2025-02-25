// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package internal

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

// AntCirculatingSupplyMetaData contains all meta data concerning the AntCirculatingSupply contract.
var AntCirculatingSupplyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_shareholders\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_emissionsReserve\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_emissionsService\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_foundationLP\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_foundation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_foundationNodeRewards\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_foundationReserve\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_emaidAirdropper\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ANT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"emaidAirdropper\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"emissionsReserve\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"emissionsService\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"foundation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"foundationLP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"foundationNodeRewards\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"foundationReserve\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEmaidAirdropper\",\"inputs\":[{\"name\":\"_emaidAirdropper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEmissionsReserve\",\"inputs\":[{\"name\":\"_emissionsReserve\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEmissionsService\",\"inputs\":[{\"name\":\"_emissionsService\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFoundation\",\"inputs\":[{\"name\":\"_foundation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFoundationLP\",\"inputs\":[{\"name\":\"_foundationLP\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFoundationNodeRewards\",\"inputs\":[{\"name\":\"_foundationNodeRewards\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFoundationReserve\",\"inputs\":[{\"name\":\"_foundationReserve\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setShareholders\",\"inputs\":[{\"name\":\"_shareholders\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shareholders\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalCirculatingSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// AntCirculatingSupplyABI is the input ABI used to generate the binding from.
// Deprecated: Use AntCirculatingSupplyMetaData.ABI instead.
var AntCirculatingSupplyABI = AntCirculatingSupplyMetaData.ABI

// AntCirculatingSupply is an auto generated Go binding around an Ethereum contract.
type AntCirculatingSupply struct {
	AntCirculatingSupplyCaller     // Read-only binding to the contract
	AntCirculatingSupplyTransactor // Write-only binding to the contract
	AntCirculatingSupplyFilterer   // Log filterer for contract events
}

// AntCirculatingSupplyCaller is an auto generated read-only Go binding around an Ethereum contract.
type AntCirculatingSupplyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AntCirculatingSupplyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AntCirculatingSupplyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AntCirculatingSupplyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AntCirculatingSupplyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AntCirculatingSupplySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AntCirculatingSupplySession struct {
	Contract     *AntCirculatingSupply // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AntCirculatingSupplyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AntCirculatingSupplyCallerSession struct {
	Contract *AntCirculatingSupplyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// AntCirculatingSupplyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AntCirculatingSupplyTransactorSession struct {
	Contract     *AntCirculatingSupplyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// AntCirculatingSupplyRaw is an auto generated low-level Go binding around an Ethereum contract.
type AntCirculatingSupplyRaw struct {
	Contract *AntCirculatingSupply // Generic contract binding to access the raw methods on
}

// AntCirculatingSupplyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AntCirculatingSupplyCallerRaw struct {
	Contract *AntCirculatingSupplyCaller // Generic read-only contract binding to access the raw methods on
}

// AntCirculatingSupplyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AntCirculatingSupplyTransactorRaw struct {
	Contract *AntCirculatingSupplyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAntCirculatingSupply creates a new instance of AntCirculatingSupply, bound to a specific deployed contract.
func NewAntCirculatingSupply(address common.Address, backend bind.ContractBackend) (*AntCirculatingSupply, error) {
	contract, err := bindAntCirculatingSupply(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AntCirculatingSupply{AntCirculatingSupplyCaller: AntCirculatingSupplyCaller{contract: contract}, AntCirculatingSupplyTransactor: AntCirculatingSupplyTransactor{contract: contract}, AntCirculatingSupplyFilterer: AntCirculatingSupplyFilterer{contract: contract}}, nil
}

// NewAntCirculatingSupplyCaller creates a new read-only instance of AntCirculatingSupply, bound to a specific deployed contract.
func NewAntCirculatingSupplyCaller(address common.Address, caller bind.ContractCaller) (*AntCirculatingSupplyCaller, error) {
	contract, err := bindAntCirculatingSupply(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AntCirculatingSupplyCaller{contract: contract}, nil
}

// NewAntCirculatingSupplyTransactor creates a new write-only instance of AntCirculatingSupply, bound to a specific deployed contract.
func NewAntCirculatingSupplyTransactor(address common.Address, transactor bind.ContractTransactor) (*AntCirculatingSupplyTransactor, error) {
	contract, err := bindAntCirculatingSupply(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AntCirculatingSupplyTransactor{contract: contract}, nil
}

// NewAntCirculatingSupplyFilterer creates a new log filterer instance of AntCirculatingSupply, bound to a specific deployed contract.
func NewAntCirculatingSupplyFilterer(address common.Address, filterer bind.ContractFilterer) (*AntCirculatingSupplyFilterer, error) {
	contract, err := bindAntCirculatingSupply(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AntCirculatingSupplyFilterer{contract: contract}, nil
}

// bindAntCirculatingSupply binds a generic wrapper to an already deployed contract.
func bindAntCirculatingSupply(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AntCirculatingSupplyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AntCirculatingSupply *AntCirculatingSupplyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AntCirculatingSupply.Contract.AntCirculatingSupplyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AntCirculatingSupply *AntCirculatingSupplyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.AntCirculatingSupplyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AntCirculatingSupply *AntCirculatingSupplyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.AntCirculatingSupplyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AntCirculatingSupply *AntCirculatingSupplyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AntCirculatingSupply.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AntCirculatingSupply *AntCirculatingSupplyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AntCirculatingSupply *AntCirculatingSupplyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.contract.Transact(opts, method, params...)
}

// ANT is a free data retrieval call binding the contract method 0x65595c61.
//
// Solidity: function ANT() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCaller) ANT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AntCirculatingSupply.contract.Call(opts, &out, "ANT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ANT is a free data retrieval call binding the contract method 0x65595c61.
//
// Solidity: function ANT() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplySession) ANT() (common.Address, error) {
	return _AntCirculatingSupply.Contract.ANT(&_AntCirculatingSupply.CallOpts)
}

// ANT is a free data retrieval call binding the contract method 0x65595c61.
//
// Solidity: function ANT() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCallerSession) ANT() (common.Address, error) {
	return _AntCirculatingSupply.Contract.ANT(&_AntCirculatingSupply.CallOpts)
}

// EmaidAirdropper is a free data retrieval call binding the contract method 0xe1b64ffb.
//
// Solidity: function emaidAirdropper() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCaller) EmaidAirdropper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AntCirculatingSupply.contract.Call(opts, &out, "emaidAirdropper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmaidAirdropper is a free data retrieval call binding the contract method 0xe1b64ffb.
//
// Solidity: function emaidAirdropper() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplySession) EmaidAirdropper() (common.Address, error) {
	return _AntCirculatingSupply.Contract.EmaidAirdropper(&_AntCirculatingSupply.CallOpts)
}

// EmaidAirdropper is a free data retrieval call binding the contract method 0xe1b64ffb.
//
// Solidity: function emaidAirdropper() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCallerSession) EmaidAirdropper() (common.Address, error) {
	return _AntCirculatingSupply.Contract.EmaidAirdropper(&_AntCirculatingSupply.CallOpts)
}

// EmissionsReserve is a free data retrieval call binding the contract method 0x6a47c7bb.
//
// Solidity: function emissionsReserve() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCaller) EmissionsReserve(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AntCirculatingSupply.contract.Call(opts, &out, "emissionsReserve")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmissionsReserve is a free data retrieval call binding the contract method 0x6a47c7bb.
//
// Solidity: function emissionsReserve() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplySession) EmissionsReserve() (common.Address, error) {
	return _AntCirculatingSupply.Contract.EmissionsReserve(&_AntCirculatingSupply.CallOpts)
}

// EmissionsReserve is a free data retrieval call binding the contract method 0x6a47c7bb.
//
// Solidity: function emissionsReserve() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCallerSession) EmissionsReserve() (common.Address, error) {
	return _AntCirculatingSupply.Contract.EmissionsReserve(&_AntCirculatingSupply.CallOpts)
}

// EmissionsService is a free data retrieval call binding the contract method 0x07ca6836.
//
// Solidity: function emissionsService() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCaller) EmissionsService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AntCirculatingSupply.contract.Call(opts, &out, "emissionsService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmissionsService is a free data retrieval call binding the contract method 0x07ca6836.
//
// Solidity: function emissionsService() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplySession) EmissionsService() (common.Address, error) {
	return _AntCirculatingSupply.Contract.EmissionsService(&_AntCirculatingSupply.CallOpts)
}

// EmissionsService is a free data retrieval call binding the contract method 0x07ca6836.
//
// Solidity: function emissionsService() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCallerSession) EmissionsService() (common.Address, error) {
	return _AntCirculatingSupply.Contract.EmissionsService(&_AntCirculatingSupply.CallOpts)
}

// Foundation is a free data retrieval call binding the contract method 0x41fbb050.
//
// Solidity: function foundation() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCaller) Foundation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AntCirculatingSupply.contract.Call(opts, &out, "foundation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Foundation is a free data retrieval call binding the contract method 0x41fbb050.
//
// Solidity: function foundation() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplySession) Foundation() (common.Address, error) {
	return _AntCirculatingSupply.Contract.Foundation(&_AntCirculatingSupply.CallOpts)
}

// Foundation is a free data retrieval call binding the contract method 0x41fbb050.
//
// Solidity: function foundation() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCallerSession) Foundation() (common.Address, error) {
	return _AntCirculatingSupply.Contract.Foundation(&_AntCirculatingSupply.CallOpts)
}

// FoundationLP is a free data retrieval call binding the contract method 0x47192aff.
//
// Solidity: function foundationLP() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCaller) FoundationLP(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AntCirculatingSupply.contract.Call(opts, &out, "foundationLP")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FoundationLP is a free data retrieval call binding the contract method 0x47192aff.
//
// Solidity: function foundationLP() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplySession) FoundationLP() (common.Address, error) {
	return _AntCirculatingSupply.Contract.FoundationLP(&_AntCirculatingSupply.CallOpts)
}

// FoundationLP is a free data retrieval call binding the contract method 0x47192aff.
//
// Solidity: function foundationLP() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCallerSession) FoundationLP() (common.Address, error) {
	return _AntCirculatingSupply.Contract.FoundationLP(&_AntCirculatingSupply.CallOpts)
}

// FoundationNodeRewards is a free data retrieval call binding the contract method 0xf4aba97d.
//
// Solidity: function foundationNodeRewards() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCaller) FoundationNodeRewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AntCirculatingSupply.contract.Call(opts, &out, "foundationNodeRewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FoundationNodeRewards is a free data retrieval call binding the contract method 0xf4aba97d.
//
// Solidity: function foundationNodeRewards() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplySession) FoundationNodeRewards() (common.Address, error) {
	return _AntCirculatingSupply.Contract.FoundationNodeRewards(&_AntCirculatingSupply.CallOpts)
}

// FoundationNodeRewards is a free data retrieval call binding the contract method 0xf4aba97d.
//
// Solidity: function foundationNodeRewards() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCallerSession) FoundationNodeRewards() (common.Address, error) {
	return _AntCirculatingSupply.Contract.FoundationNodeRewards(&_AntCirculatingSupply.CallOpts)
}

// FoundationReserve is a free data retrieval call binding the contract method 0x603066a4.
//
// Solidity: function foundationReserve() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCaller) FoundationReserve(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AntCirculatingSupply.contract.Call(opts, &out, "foundationReserve")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FoundationReserve is a free data retrieval call binding the contract method 0x603066a4.
//
// Solidity: function foundationReserve() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplySession) FoundationReserve() (common.Address, error) {
	return _AntCirculatingSupply.Contract.FoundationReserve(&_AntCirculatingSupply.CallOpts)
}

// FoundationReserve is a free data retrieval call binding the contract method 0x603066a4.
//
// Solidity: function foundationReserve() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCallerSession) FoundationReserve() (common.Address, error) {
	return _AntCirculatingSupply.Contract.FoundationReserve(&_AntCirculatingSupply.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AntCirculatingSupply.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplySession) Owner() (common.Address, error) {
	return _AntCirculatingSupply.Contract.Owner(&_AntCirculatingSupply.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCallerSession) Owner() (common.Address, error) {
	return _AntCirculatingSupply.Contract.Owner(&_AntCirculatingSupply.CallOpts)
}

// Shareholders is a free data retrieval call binding the contract method 0x3723bc0e.
//
// Solidity: function shareholders() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCaller) Shareholders(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AntCirculatingSupply.contract.Call(opts, &out, "shareholders")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Shareholders is a free data retrieval call binding the contract method 0x3723bc0e.
//
// Solidity: function shareholders() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplySession) Shareholders() (common.Address, error) {
	return _AntCirculatingSupply.Contract.Shareholders(&_AntCirculatingSupply.CallOpts)
}

// Shareholders is a free data retrieval call binding the contract method 0x3723bc0e.
//
// Solidity: function shareholders() view returns(address)
func (_AntCirculatingSupply *AntCirculatingSupplyCallerSession) Shareholders() (common.Address, error) {
	return _AntCirculatingSupply.Contract.Shareholders(&_AntCirculatingSupply.CallOpts)
}

// TotalCirculatingSupply is a free data retrieval call binding the contract method 0x5ee0ce31.
//
// Solidity: function totalCirculatingSupply() view returns(uint256)
func (_AntCirculatingSupply *AntCirculatingSupplyCaller) TotalCirculatingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AntCirculatingSupply.contract.Call(opts, &out, "totalCirculatingSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCirculatingSupply is a free data retrieval call binding the contract method 0x5ee0ce31.
//
// Solidity: function totalCirculatingSupply() view returns(uint256)
func (_AntCirculatingSupply *AntCirculatingSupplySession) TotalCirculatingSupply() (*big.Int, error) {
	return _AntCirculatingSupply.Contract.TotalCirculatingSupply(&_AntCirculatingSupply.CallOpts)
}

// TotalCirculatingSupply is a free data retrieval call binding the contract method 0x5ee0ce31.
//
// Solidity: function totalCirculatingSupply() view returns(uint256)
func (_AntCirculatingSupply *AntCirculatingSupplyCallerSession) TotalCirculatingSupply() (*big.Int, error) {
	return _AntCirculatingSupply.Contract.TotalCirculatingSupply(&_AntCirculatingSupply.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AntCirculatingSupply.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AntCirculatingSupply *AntCirculatingSupplySession) RenounceOwnership() (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.RenounceOwnership(&_AntCirculatingSupply.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.RenounceOwnership(&_AntCirculatingSupply.TransactOpts)
}

// SetEmaidAirdropper is a paid mutator transaction binding the contract method 0x6937672c.
//
// Solidity: function setEmaidAirdropper(address _emaidAirdropper) returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactor) SetEmaidAirdropper(opts *bind.TransactOpts, _emaidAirdropper common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.contract.Transact(opts, "setEmaidAirdropper", _emaidAirdropper)
}

// SetEmaidAirdropper is a paid mutator transaction binding the contract method 0x6937672c.
//
// Solidity: function setEmaidAirdropper(address _emaidAirdropper) returns()
func (_AntCirculatingSupply *AntCirculatingSupplySession) SetEmaidAirdropper(_emaidAirdropper common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.SetEmaidAirdropper(&_AntCirculatingSupply.TransactOpts, _emaidAirdropper)
}

// SetEmaidAirdropper is a paid mutator transaction binding the contract method 0x6937672c.
//
// Solidity: function setEmaidAirdropper(address _emaidAirdropper) returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactorSession) SetEmaidAirdropper(_emaidAirdropper common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.SetEmaidAirdropper(&_AntCirculatingSupply.TransactOpts, _emaidAirdropper)
}

// SetEmissionsReserve is a paid mutator transaction binding the contract method 0x12885a9c.
//
// Solidity: function setEmissionsReserve(address _emissionsReserve) returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactor) SetEmissionsReserve(opts *bind.TransactOpts, _emissionsReserve common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.contract.Transact(opts, "setEmissionsReserve", _emissionsReserve)
}

// SetEmissionsReserve is a paid mutator transaction binding the contract method 0x12885a9c.
//
// Solidity: function setEmissionsReserve(address _emissionsReserve) returns()
func (_AntCirculatingSupply *AntCirculatingSupplySession) SetEmissionsReserve(_emissionsReserve common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.SetEmissionsReserve(&_AntCirculatingSupply.TransactOpts, _emissionsReserve)
}

// SetEmissionsReserve is a paid mutator transaction binding the contract method 0x12885a9c.
//
// Solidity: function setEmissionsReserve(address _emissionsReserve) returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactorSession) SetEmissionsReserve(_emissionsReserve common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.SetEmissionsReserve(&_AntCirculatingSupply.TransactOpts, _emissionsReserve)
}

// SetEmissionsService is a paid mutator transaction binding the contract method 0x7228d170.
//
// Solidity: function setEmissionsService(address _emissionsService) returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactor) SetEmissionsService(opts *bind.TransactOpts, _emissionsService common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.contract.Transact(opts, "setEmissionsService", _emissionsService)
}

// SetEmissionsService is a paid mutator transaction binding the contract method 0x7228d170.
//
// Solidity: function setEmissionsService(address _emissionsService) returns()
func (_AntCirculatingSupply *AntCirculatingSupplySession) SetEmissionsService(_emissionsService common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.SetEmissionsService(&_AntCirculatingSupply.TransactOpts, _emissionsService)
}

// SetEmissionsService is a paid mutator transaction binding the contract method 0x7228d170.
//
// Solidity: function setEmissionsService(address _emissionsService) returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactorSession) SetEmissionsService(_emissionsService common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.SetEmissionsService(&_AntCirculatingSupply.TransactOpts, _emissionsService)
}

// SetFoundation is a paid mutator transaction binding the contract method 0xdb3543f5.
//
// Solidity: function setFoundation(address _foundation) returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactor) SetFoundation(opts *bind.TransactOpts, _foundation common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.contract.Transact(opts, "setFoundation", _foundation)
}

// SetFoundation is a paid mutator transaction binding the contract method 0xdb3543f5.
//
// Solidity: function setFoundation(address _foundation) returns()
func (_AntCirculatingSupply *AntCirculatingSupplySession) SetFoundation(_foundation common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.SetFoundation(&_AntCirculatingSupply.TransactOpts, _foundation)
}

// SetFoundation is a paid mutator transaction binding the contract method 0xdb3543f5.
//
// Solidity: function setFoundation(address _foundation) returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactorSession) SetFoundation(_foundation common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.SetFoundation(&_AntCirculatingSupply.TransactOpts, _foundation)
}

// SetFoundationLP is a paid mutator transaction binding the contract method 0x4c9b028c.
//
// Solidity: function setFoundationLP(address _foundationLP) returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactor) SetFoundationLP(opts *bind.TransactOpts, _foundationLP common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.contract.Transact(opts, "setFoundationLP", _foundationLP)
}

// SetFoundationLP is a paid mutator transaction binding the contract method 0x4c9b028c.
//
// Solidity: function setFoundationLP(address _foundationLP) returns()
func (_AntCirculatingSupply *AntCirculatingSupplySession) SetFoundationLP(_foundationLP common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.SetFoundationLP(&_AntCirculatingSupply.TransactOpts, _foundationLP)
}

// SetFoundationLP is a paid mutator transaction binding the contract method 0x4c9b028c.
//
// Solidity: function setFoundationLP(address _foundationLP) returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactorSession) SetFoundationLP(_foundationLP common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.SetFoundationLP(&_AntCirculatingSupply.TransactOpts, _foundationLP)
}

// SetFoundationNodeRewards is a paid mutator transaction binding the contract method 0xc702e930.
//
// Solidity: function setFoundationNodeRewards(address _foundationNodeRewards) returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactor) SetFoundationNodeRewards(opts *bind.TransactOpts, _foundationNodeRewards common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.contract.Transact(opts, "setFoundationNodeRewards", _foundationNodeRewards)
}

// SetFoundationNodeRewards is a paid mutator transaction binding the contract method 0xc702e930.
//
// Solidity: function setFoundationNodeRewards(address _foundationNodeRewards) returns()
func (_AntCirculatingSupply *AntCirculatingSupplySession) SetFoundationNodeRewards(_foundationNodeRewards common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.SetFoundationNodeRewards(&_AntCirculatingSupply.TransactOpts, _foundationNodeRewards)
}

// SetFoundationNodeRewards is a paid mutator transaction binding the contract method 0xc702e930.
//
// Solidity: function setFoundationNodeRewards(address _foundationNodeRewards) returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactorSession) SetFoundationNodeRewards(_foundationNodeRewards common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.SetFoundationNodeRewards(&_AntCirculatingSupply.TransactOpts, _foundationNodeRewards)
}

// SetFoundationReserve is a paid mutator transaction binding the contract method 0x0b241883.
//
// Solidity: function setFoundationReserve(address _foundationReserve) returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactor) SetFoundationReserve(opts *bind.TransactOpts, _foundationReserve common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.contract.Transact(opts, "setFoundationReserve", _foundationReserve)
}

// SetFoundationReserve is a paid mutator transaction binding the contract method 0x0b241883.
//
// Solidity: function setFoundationReserve(address _foundationReserve) returns()
func (_AntCirculatingSupply *AntCirculatingSupplySession) SetFoundationReserve(_foundationReserve common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.SetFoundationReserve(&_AntCirculatingSupply.TransactOpts, _foundationReserve)
}

// SetFoundationReserve is a paid mutator transaction binding the contract method 0x0b241883.
//
// Solidity: function setFoundationReserve(address _foundationReserve) returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactorSession) SetFoundationReserve(_foundationReserve common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.SetFoundationReserve(&_AntCirculatingSupply.TransactOpts, _foundationReserve)
}

// SetShareholders is a paid mutator transaction binding the contract method 0x60a7b7fb.
//
// Solidity: function setShareholders(address _shareholders) returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactor) SetShareholders(opts *bind.TransactOpts, _shareholders common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.contract.Transact(opts, "setShareholders", _shareholders)
}

// SetShareholders is a paid mutator transaction binding the contract method 0x60a7b7fb.
//
// Solidity: function setShareholders(address _shareholders) returns()
func (_AntCirculatingSupply *AntCirculatingSupplySession) SetShareholders(_shareholders common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.SetShareholders(&_AntCirculatingSupply.TransactOpts, _shareholders)
}

// SetShareholders is a paid mutator transaction binding the contract method 0x60a7b7fb.
//
// Solidity: function setShareholders(address _shareholders) returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactorSession) SetShareholders(_shareholders common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.SetShareholders(&_AntCirculatingSupply.TransactOpts, _shareholders)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AntCirculatingSupply *AntCirculatingSupplySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.TransferOwnership(&_AntCirculatingSupply.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AntCirculatingSupply *AntCirculatingSupplyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AntCirculatingSupply.Contract.TransferOwnership(&_AntCirculatingSupply.TransactOpts, newOwner)
}

// AntCirculatingSupplyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AntCirculatingSupply contract.
type AntCirculatingSupplyOwnershipTransferredIterator struct {
	Event *AntCirculatingSupplyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AntCirculatingSupplyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AntCirculatingSupplyOwnershipTransferred)
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
		it.Event = new(AntCirculatingSupplyOwnershipTransferred)
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
func (it *AntCirculatingSupplyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AntCirculatingSupplyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AntCirculatingSupplyOwnershipTransferred represents a OwnershipTransferred event raised by the AntCirculatingSupply contract.
type AntCirculatingSupplyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AntCirculatingSupply *AntCirculatingSupplyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AntCirculatingSupplyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AntCirculatingSupply.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AntCirculatingSupplyOwnershipTransferredIterator{contract: _AntCirculatingSupply.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AntCirculatingSupply *AntCirculatingSupplyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AntCirculatingSupplyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AntCirculatingSupply.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AntCirculatingSupplyOwnershipTransferred)
				if err := _AntCirculatingSupply.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AntCirculatingSupply *AntCirculatingSupplyFilterer) ParseOwnershipTransferred(log types.Log) (*AntCirculatingSupplyOwnershipTransferred, error) {
	event := new(AntCirculatingSupplyOwnershipTransferred)
	if err := _AntCirculatingSupply.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
