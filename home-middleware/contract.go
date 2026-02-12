// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// AdvancedHomeAutomationMetaData contains all meta data concerning the AdvancedHomeAutomation contract.
var AdvancedHomeAutomationMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GUEST_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ROOM_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SUPER_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accessRules\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"fromTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addSuperAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRoom\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_ip\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"defineDevice\",\"inputs\":[{\"name\":\"_roomId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_pin\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_type\",\"type\":\"uint8\",\"internalType\":\"enumAdvancedHomeAutomation.DeviceType\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDeviceStatus\",\"inputs\":[{\"name\":\"_rId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_dId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantAccess\",\"inputs\":[{\"name\":\"_roomId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_end\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operateDevice\",\"inputs\":[{\"name\":\"_roomId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_deviceId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeSuperAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeAccess\",\"inputs\":[{\"name\":\"_roomId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"roomCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rooms\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"espIP\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"deviceCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AccessUpdated\",\"inputs\":[{\"name\":\"roomId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoomCreated\",\"inputs\":[{\"name\":\"roomId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StateChanged\",\"inputs\":[{\"name\":\"roomId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"deviceId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// AdvancedHomeAutomationABI is the input ABI used to generate the binding from.
// Deprecated: Use AdvancedHomeAutomationMetaData.ABI instead.
var AdvancedHomeAutomationABI = AdvancedHomeAutomationMetaData.ABI

// AdvancedHomeAutomation is an auto generated Go binding around an Ethereum contract.
type AdvancedHomeAutomation struct {
	AdvancedHomeAutomationCaller     // Read-only binding to the contract
	AdvancedHomeAutomationTransactor // Write-only binding to the contract
	AdvancedHomeAutomationFilterer   // Log filterer for contract events
}

// AdvancedHomeAutomationCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdvancedHomeAutomationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdvancedHomeAutomationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdvancedHomeAutomationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdvancedHomeAutomationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdvancedHomeAutomationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdvancedHomeAutomationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdvancedHomeAutomationSession struct {
	Contract     *AdvancedHomeAutomation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AdvancedHomeAutomationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdvancedHomeAutomationCallerSession struct {
	Contract *AdvancedHomeAutomationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// AdvancedHomeAutomationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdvancedHomeAutomationTransactorSession struct {
	Contract     *AdvancedHomeAutomationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// AdvancedHomeAutomationRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdvancedHomeAutomationRaw struct {
	Contract *AdvancedHomeAutomation // Generic contract binding to access the raw methods on
}

// AdvancedHomeAutomationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdvancedHomeAutomationCallerRaw struct {
	Contract *AdvancedHomeAutomationCaller // Generic read-only contract binding to access the raw methods on
}

// AdvancedHomeAutomationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdvancedHomeAutomationTransactorRaw struct {
	Contract *AdvancedHomeAutomationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdvancedHomeAutomation creates a new instance of AdvancedHomeAutomation, bound to a specific deployed contract.
func NewAdvancedHomeAutomation(address common.Address, backend bind.ContractBackend) (*AdvancedHomeAutomation, error) {
	contract, err := bindAdvancedHomeAutomation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdvancedHomeAutomation{AdvancedHomeAutomationCaller: AdvancedHomeAutomationCaller{contract: contract}, AdvancedHomeAutomationTransactor: AdvancedHomeAutomationTransactor{contract: contract}, AdvancedHomeAutomationFilterer: AdvancedHomeAutomationFilterer{contract: contract}}, nil
}

// NewAdvancedHomeAutomationCaller creates a new read-only instance of AdvancedHomeAutomation, bound to a specific deployed contract.
func NewAdvancedHomeAutomationCaller(address common.Address, caller bind.ContractCaller) (*AdvancedHomeAutomationCaller, error) {
	contract, err := bindAdvancedHomeAutomation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdvancedHomeAutomationCaller{contract: contract}, nil
}

// NewAdvancedHomeAutomationTransactor creates a new write-only instance of AdvancedHomeAutomation, bound to a specific deployed contract.
func NewAdvancedHomeAutomationTransactor(address common.Address, transactor bind.ContractTransactor) (*AdvancedHomeAutomationTransactor, error) {
	contract, err := bindAdvancedHomeAutomation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdvancedHomeAutomationTransactor{contract: contract}, nil
}

// NewAdvancedHomeAutomationFilterer creates a new log filterer instance of AdvancedHomeAutomation, bound to a specific deployed contract.
func NewAdvancedHomeAutomationFilterer(address common.Address, filterer bind.ContractFilterer) (*AdvancedHomeAutomationFilterer, error) {
	contract, err := bindAdvancedHomeAutomation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdvancedHomeAutomationFilterer{contract: contract}, nil
}

// bindAdvancedHomeAutomation binds a generic wrapper to an already deployed contract.
func bindAdvancedHomeAutomation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AdvancedHomeAutomationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdvancedHomeAutomation *AdvancedHomeAutomationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdvancedHomeAutomation.Contract.AdvancedHomeAutomationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdvancedHomeAutomation *AdvancedHomeAutomationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.AdvancedHomeAutomationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdvancedHomeAutomation *AdvancedHomeAutomationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.AdvancedHomeAutomationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdvancedHomeAutomation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdvancedHomeAutomation.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AdvancedHomeAutomation.Contract.DEFAULTADMINROLE(&_AdvancedHomeAutomation.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AdvancedHomeAutomation.Contract.DEFAULTADMINROLE(&_AdvancedHomeAutomation.CallOpts)
}

// GUESTROLE is a free data retrieval call binding the contract method 0x5c7e8863.
//
// Solidity: function GUEST_ROLE() view returns(bytes32)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCaller) GUESTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdvancedHomeAutomation.contract.Call(opts, &out, "GUEST_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GUESTROLE is a free data retrieval call binding the contract method 0x5c7e8863.
//
// Solidity: function GUEST_ROLE() view returns(bytes32)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) GUESTROLE() ([32]byte, error) {
	return _AdvancedHomeAutomation.Contract.GUESTROLE(&_AdvancedHomeAutomation.CallOpts)
}

// GUESTROLE is a free data retrieval call binding the contract method 0x5c7e8863.
//
// Solidity: function GUEST_ROLE() view returns(bytes32)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCallerSession) GUESTROLE() ([32]byte, error) {
	return _AdvancedHomeAutomation.Contract.GUESTROLE(&_AdvancedHomeAutomation.CallOpts)
}

// ROOMADMINROLE is a free data retrieval call binding the contract method 0x8437db89.
//
// Solidity: function ROOM_ADMIN_ROLE() view returns(bytes32)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCaller) ROOMADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdvancedHomeAutomation.contract.Call(opts, &out, "ROOM_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROOMADMINROLE is a free data retrieval call binding the contract method 0x8437db89.
//
// Solidity: function ROOM_ADMIN_ROLE() view returns(bytes32)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) ROOMADMINROLE() ([32]byte, error) {
	return _AdvancedHomeAutomation.Contract.ROOMADMINROLE(&_AdvancedHomeAutomation.CallOpts)
}

// ROOMADMINROLE is a free data retrieval call binding the contract method 0x8437db89.
//
// Solidity: function ROOM_ADMIN_ROLE() view returns(bytes32)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCallerSession) ROOMADMINROLE() ([32]byte, error) {
	return _AdvancedHomeAutomation.Contract.ROOMADMINROLE(&_AdvancedHomeAutomation.CallOpts)
}

// SUPERADMINROLE is a free data retrieval call binding the contract method 0x4460bdd6.
//
// Solidity: function SUPER_ADMIN_ROLE() view returns(bytes32)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCaller) SUPERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AdvancedHomeAutomation.contract.Call(opts, &out, "SUPER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SUPERADMINROLE is a free data retrieval call binding the contract method 0x4460bdd6.
//
// Solidity: function SUPER_ADMIN_ROLE() view returns(bytes32)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) SUPERADMINROLE() ([32]byte, error) {
	return _AdvancedHomeAutomation.Contract.SUPERADMINROLE(&_AdvancedHomeAutomation.CallOpts)
}

// SUPERADMINROLE is a free data retrieval call binding the contract method 0x4460bdd6.
//
// Solidity: function SUPER_ADMIN_ROLE() view returns(bytes32)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCallerSession) SUPERADMINROLE() ([32]byte, error) {
	return _AdvancedHomeAutomation.Contract.SUPERADMINROLE(&_AdvancedHomeAutomation.CallOpts)
}

// AccessRules is a free data retrieval call binding the contract method 0xfa525fb5.
//
// Solidity: function accessRules(uint256 , address ) view returns(uint256 fromTimestamp, uint256 toTimestamp, bool isActive)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCaller) AccessRules(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	FromTimestamp *big.Int
	ToTimestamp   *big.Int
	IsActive      bool
}, error) {
	var out []interface{}
	err := _AdvancedHomeAutomation.contract.Call(opts, &out, "accessRules", arg0, arg1)

	outstruct := new(struct {
		FromTimestamp *big.Int
		ToTimestamp   *big.Int
		IsActive      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ToTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// AccessRules is a free data retrieval call binding the contract method 0xfa525fb5.
//
// Solidity: function accessRules(uint256 , address ) view returns(uint256 fromTimestamp, uint256 toTimestamp, bool isActive)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) AccessRules(arg0 *big.Int, arg1 common.Address) (struct {
	FromTimestamp *big.Int
	ToTimestamp   *big.Int
	IsActive      bool
}, error) {
	return _AdvancedHomeAutomation.Contract.AccessRules(&_AdvancedHomeAutomation.CallOpts, arg0, arg1)
}

// AccessRules is a free data retrieval call binding the contract method 0xfa525fb5.
//
// Solidity: function accessRules(uint256 , address ) view returns(uint256 fromTimestamp, uint256 toTimestamp, bool isActive)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCallerSession) AccessRules(arg0 *big.Int, arg1 common.Address) (struct {
	FromTimestamp *big.Int
	ToTimestamp   *big.Int
	IsActive      bool
}, error) {
	return _AdvancedHomeAutomation.Contract.AccessRules(&_AdvancedHomeAutomation.CallOpts, arg0, arg1)
}

// GetDeviceStatus is a free data retrieval call binding the contract method 0x1d16f5c9.
//
// Solidity: function getDeviceStatus(uint256 _rId, uint256 _dId) view returns(uint256)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCaller) GetDeviceStatus(opts *bind.CallOpts, _rId *big.Int, _dId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AdvancedHomeAutomation.contract.Call(opts, &out, "getDeviceStatus", _rId, _dId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeviceStatus is a free data retrieval call binding the contract method 0x1d16f5c9.
//
// Solidity: function getDeviceStatus(uint256 _rId, uint256 _dId) view returns(uint256)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) GetDeviceStatus(_rId *big.Int, _dId *big.Int) (*big.Int, error) {
	return _AdvancedHomeAutomation.Contract.GetDeviceStatus(&_AdvancedHomeAutomation.CallOpts, _rId, _dId)
}

// GetDeviceStatus is a free data retrieval call binding the contract method 0x1d16f5c9.
//
// Solidity: function getDeviceStatus(uint256 _rId, uint256 _dId) view returns(uint256)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCallerSession) GetDeviceStatus(_rId *big.Int, _dId *big.Int) (*big.Int, error) {
	return _AdvancedHomeAutomation.Contract.GetDeviceStatus(&_AdvancedHomeAutomation.CallOpts, _rId, _dId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AdvancedHomeAutomation.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AdvancedHomeAutomation.Contract.GetRoleAdmin(&_AdvancedHomeAutomation.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AdvancedHomeAutomation.Contract.GetRoleAdmin(&_AdvancedHomeAutomation.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AdvancedHomeAutomation.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AdvancedHomeAutomation.Contract.HasRole(&_AdvancedHomeAutomation.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AdvancedHomeAutomation.Contract.HasRole(&_AdvancedHomeAutomation.CallOpts, role, account)
}

// RoomCount is a free data retrieval call binding the contract method 0xdf93a4e3.
//
// Solidity: function roomCount() view returns(uint256)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCaller) RoomCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdvancedHomeAutomation.contract.Call(opts, &out, "roomCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoomCount is a free data retrieval call binding the contract method 0xdf93a4e3.
//
// Solidity: function roomCount() view returns(uint256)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) RoomCount() (*big.Int, error) {
	return _AdvancedHomeAutomation.Contract.RoomCount(&_AdvancedHomeAutomation.CallOpts)
}

// RoomCount is a free data retrieval call binding the contract method 0xdf93a4e3.
//
// Solidity: function roomCount() view returns(uint256)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCallerSession) RoomCount() (*big.Int, error) {
	return _AdvancedHomeAutomation.Contract.RoomCount(&_AdvancedHomeAutomation.CallOpts)
}

// Rooms is a free data retrieval call binding the contract method 0x1bae0ac8.
//
// Solidity: function rooms(uint256 ) view returns(string name, string espIP, uint256 deviceCount, bool exists)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCaller) Rooms(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name        string
	EspIP       string
	DeviceCount *big.Int
	Exists      bool
}, error) {
	var out []interface{}
	err := _AdvancedHomeAutomation.contract.Call(opts, &out, "rooms", arg0)

	outstruct := new(struct {
		Name        string
		EspIP       string
		DeviceCount *big.Int
		Exists      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.EspIP = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.DeviceCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Exists = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Rooms is a free data retrieval call binding the contract method 0x1bae0ac8.
//
// Solidity: function rooms(uint256 ) view returns(string name, string espIP, uint256 deviceCount, bool exists)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) Rooms(arg0 *big.Int) (struct {
	Name        string
	EspIP       string
	DeviceCount *big.Int
	Exists      bool
}, error) {
	return _AdvancedHomeAutomation.Contract.Rooms(&_AdvancedHomeAutomation.CallOpts, arg0)
}

// Rooms is a free data retrieval call binding the contract method 0x1bae0ac8.
//
// Solidity: function rooms(uint256 ) view returns(string name, string espIP, uint256 deviceCount, bool exists)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCallerSession) Rooms(arg0 *big.Int) (struct {
	Name        string
	EspIP       string
	DeviceCount *big.Int
	Exists      bool
}, error) {
	return _AdvancedHomeAutomation.Contract.Rooms(&_AdvancedHomeAutomation.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AdvancedHomeAutomation.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AdvancedHomeAutomation.Contract.SupportsInterface(&_AdvancedHomeAutomation.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AdvancedHomeAutomation.Contract.SupportsInterface(&_AdvancedHomeAutomation.CallOpts, interfaceId)
}

// AddSuperAdmin is a paid mutator transaction binding the contract method 0xb3292ff0.
//
// Solidity: function addSuperAdmin(address newAdmin) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactor) AddSuperAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.contract.Transact(opts, "addSuperAdmin", newAdmin)
}

// AddSuperAdmin is a paid mutator transaction binding the contract method 0xb3292ff0.
//
// Solidity: function addSuperAdmin(address newAdmin) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) AddSuperAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.AddSuperAdmin(&_AdvancedHomeAutomation.TransactOpts, newAdmin)
}

// AddSuperAdmin is a paid mutator transaction binding the contract method 0xb3292ff0.
//
// Solidity: function addSuperAdmin(address newAdmin) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactorSession) AddSuperAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.AddSuperAdmin(&_AdvancedHomeAutomation.TransactOpts, newAdmin)
}

// CreateRoom is a paid mutator transaction binding the contract method 0xfa68daad.
//
// Solidity: function createRoom(string _name, string _ip) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactor) CreateRoom(opts *bind.TransactOpts, _name string, _ip string) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.contract.Transact(opts, "createRoom", _name, _ip)
}

// CreateRoom is a paid mutator transaction binding the contract method 0xfa68daad.
//
// Solidity: function createRoom(string _name, string _ip) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) CreateRoom(_name string, _ip string) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.CreateRoom(&_AdvancedHomeAutomation.TransactOpts, _name, _ip)
}

// CreateRoom is a paid mutator transaction binding the contract method 0xfa68daad.
//
// Solidity: function createRoom(string _name, string _ip) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactorSession) CreateRoom(_name string, _ip string) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.CreateRoom(&_AdvancedHomeAutomation.TransactOpts, _name, _ip)
}

// DefineDevice is a paid mutator transaction binding the contract method 0x29e9c277.
//
// Solidity: function defineDevice(uint256 _roomId, string _name, uint256 _pin, uint8 _type) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactor) DefineDevice(opts *bind.TransactOpts, _roomId *big.Int, _name string, _pin *big.Int, _type uint8) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.contract.Transact(opts, "defineDevice", _roomId, _name, _pin, _type)
}

// DefineDevice is a paid mutator transaction binding the contract method 0x29e9c277.
//
// Solidity: function defineDevice(uint256 _roomId, string _name, uint256 _pin, uint8 _type) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) DefineDevice(_roomId *big.Int, _name string, _pin *big.Int, _type uint8) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.DefineDevice(&_AdvancedHomeAutomation.TransactOpts, _roomId, _name, _pin, _type)
}

// DefineDevice is a paid mutator transaction binding the contract method 0x29e9c277.
//
// Solidity: function defineDevice(uint256 _roomId, string _name, uint256 _pin, uint8 _type) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactorSession) DefineDevice(_roomId *big.Int, _name string, _pin *big.Int, _type uint8) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.DefineDevice(&_AdvancedHomeAutomation.TransactOpts, _roomId, _name, _pin, _type)
}

// GrantAccess is a paid mutator transaction binding the contract method 0xdf399508.
//
// Solidity: function grantAccess(uint256 _roomId, address _user, uint256 _start, uint256 _end, bytes32 _role) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactor) GrantAccess(opts *bind.TransactOpts, _roomId *big.Int, _user common.Address, _start *big.Int, _end *big.Int, _role [32]byte) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.contract.Transact(opts, "grantAccess", _roomId, _user, _start, _end, _role)
}

// GrantAccess is a paid mutator transaction binding the contract method 0xdf399508.
//
// Solidity: function grantAccess(uint256 _roomId, address _user, uint256 _start, uint256 _end, bytes32 _role) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) GrantAccess(_roomId *big.Int, _user common.Address, _start *big.Int, _end *big.Int, _role [32]byte) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.GrantAccess(&_AdvancedHomeAutomation.TransactOpts, _roomId, _user, _start, _end, _role)
}

// GrantAccess is a paid mutator transaction binding the contract method 0xdf399508.
//
// Solidity: function grantAccess(uint256 _roomId, address _user, uint256 _start, uint256 _end, bytes32 _role) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactorSession) GrantAccess(_roomId *big.Int, _user common.Address, _start *big.Int, _end *big.Int, _role [32]byte) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.GrantAccess(&_AdvancedHomeAutomation.TransactOpts, _roomId, _user, _start, _end, _role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.GrantRole(&_AdvancedHomeAutomation.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.GrantRole(&_AdvancedHomeAutomation.TransactOpts, role, account)
}

// OperateDevice is a paid mutator transaction binding the contract method 0x4115369e.
//
// Solidity: function operateDevice(uint256 _roomId, uint256 _deviceId, uint256 _value) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactor) OperateDevice(opts *bind.TransactOpts, _roomId *big.Int, _deviceId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.contract.Transact(opts, "operateDevice", _roomId, _deviceId, _value)
}

// OperateDevice is a paid mutator transaction binding the contract method 0x4115369e.
//
// Solidity: function operateDevice(uint256 _roomId, uint256 _deviceId, uint256 _value) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) OperateDevice(_roomId *big.Int, _deviceId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.OperateDevice(&_AdvancedHomeAutomation.TransactOpts, _roomId, _deviceId, _value)
}

// OperateDevice is a paid mutator transaction binding the contract method 0x4115369e.
//
// Solidity: function operateDevice(uint256 _roomId, uint256 _deviceId, uint256 _value) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactorSession) OperateDevice(_roomId *big.Int, _deviceId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.OperateDevice(&_AdvancedHomeAutomation.TransactOpts, _roomId, _deviceId, _value)
}

// RemoveSuperAdmin is a paid mutator transaction binding the contract method 0x4902e4aa.
//
// Solidity: function removeSuperAdmin(address admin) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactor) RemoveSuperAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.contract.Transact(opts, "removeSuperAdmin", admin)
}

// RemoveSuperAdmin is a paid mutator transaction binding the contract method 0x4902e4aa.
//
// Solidity: function removeSuperAdmin(address admin) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) RemoveSuperAdmin(admin common.Address) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.RemoveSuperAdmin(&_AdvancedHomeAutomation.TransactOpts, admin)
}

// RemoveSuperAdmin is a paid mutator transaction binding the contract method 0x4902e4aa.
//
// Solidity: function removeSuperAdmin(address admin) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactorSession) RemoveSuperAdmin(admin common.Address) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.RemoveSuperAdmin(&_AdvancedHomeAutomation.TransactOpts, admin)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.RenounceRole(&_AdvancedHomeAutomation.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.RenounceRole(&_AdvancedHomeAutomation.TransactOpts, role, callerConfirmation)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x6af54a26.
//
// Solidity: function revokeAccess(uint256 _roomId, address _user, bytes32 _role) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactor) RevokeAccess(opts *bind.TransactOpts, _roomId *big.Int, _user common.Address, _role [32]byte) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.contract.Transact(opts, "revokeAccess", _roomId, _user, _role)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x6af54a26.
//
// Solidity: function revokeAccess(uint256 _roomId, address _user, bytes32 _role) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) RevokeAccess(_roomId *big.Int, _user common.Address, _role [32]byte) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.RevokeAccess(&_AdvancedHomeAutomation.TransactOpts, _roomId, _user, _role)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x6af54a26.
//
// Solidity: function revokeAccess(uint256 _roomId, address _user, bytes32 _role) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactorSession) RevokeAccess(_roomId *big.Int, _user common.Address, _role [32]byte) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.RevokeAccess(&_AdvancedHomeAutomation.TransactOpts, _roomId, _user, _role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.RevokeRole(&_AdvancedHomeAutomation.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AdvancedHomeAutomation *AdvancedHomeAutomationTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AdvancedHomeAutomation.Contract.RevokeRole(&_AdvancedHomeAutomation.TransactOpts, role, account)
}

// AdvancedHomeAutomationAccessUpdatedIterator is returned from FilterAccessUpdated and is used to iterate over the raw logs and unpacked data for AccessUpdated events raised by the AdvancedHomeAutomation contract.
type AdvancedHomeAutomationAccessUpdatedIterator struct {
	Event *AdvancedHomeAutomationAccessUpdated // Event containing the contract specifics and raw log

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
func (it *AdvancedHomeAutomationAccessUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedHomeAutomationAccessUpdated)
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
		it.Event = new(AdvancedHomeAutomationAccessUpdated)
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
func (it *AdvancedHomeAutomationAccessUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedHomeAutomationAccessUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedHomeAutomationAccessUpdated represents a AccessUpdated event raised by the AdvancedHomeAutomation contract.
type AdvancedHomeAutomationAccessUpdated struct {
	RoomId *big.Int
	User   common.Address
	From   *big.Int
	To     *big.Int
	Role   [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAccessUpdated is a free log retrieval operation binding the contract event 0xc761ffe24f37f69b18d4ab526796cab2020e38194ecc4598faa2c6ef62204e26.
//
// Solidity: event AccessUpdated(uint256 indexed roomId, address indexed user, uint256 from, uint256 to, bytes32 role)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationFilterer) FilterAccessUpdated(opts *bind.FilterOpts, roomId []*big.Int, user []common.Address) (*AdvancedHomeAutomationAccessUpdatedIterator, error) {

	var roomIdRule []interface{}
	for _, roomIdItem := range roomId {
		roomIdRule = append(roomIdRule, roomIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AdvancedHomeAutomation.contract.FilterLogs(opts, "AccessUpdated", roomIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedHomeAutomationAccessUpdatedIterator{contract: _AdvancedHomeAutomation.contract, event: "AccessUpdated", logs: logs, sub: sub}, nil
}

// WatchAccessUpdated is a free log subscription operation binding the contract event 0xc761ffe24f37f69b18d4ab526796cab2020e38194ecc4598faa2c6ef62204e26.
//
// Solidity: event AccessUpdated(uint256 indexed roomId, address indexed user, uint256 from, uint256 to, bytes32 role)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationFilterer) WatchAccessUpdated(opts *bind.WatchOpts, sink chan<- *AdvancedHomeAutomationAccessUpdated, roomId []*big.Int, user []common.Address) (event.Subscription, error) {

	var roomIdRule []interface{}
	for _, roomIdItem := range roomId {
		roomIdRule = append(roomIdRule, roomIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AdvancedHomeAutomation.contract.WatchLogs(opts, "AccessUpdated", roomIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedHomeAutomationAccessUpdated)
				if err := _AdvancedHomeAutomation.contract.UnpackLog(event, "AccessUpdated", log); err != nil {
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

// ParseAccessUpdated is a log parse operation binding the contract event 0xc761ffe24f37f69b18d4ab526796cab2020e38194ecc4598faa2c6ef62204e26.
//
// Solidity: event AccessUpdated(uint256 indexed roomId, address indexed user, uint256 from, uint256 to, bytes32 role)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationFilterer) ParseAccessUpdated(log types.Log) (*AdvancedHomeAutomationAccessUpdated, error) {
	event := new(AdvancedHomeAutomationAccessUpdated)
	if err := _AdvancedHomeAutomation.contract.UnpackLog(event, "AccessUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdvancedHomeAutomationRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AdvancedHomeAutomation contract.
type AdvancedHomeAutomationRoleAdminChangedIterator struct {
	Event *AdvancedHomeAutomationRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AdvancedHomeAutomationRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedHomeAutomationRoleAdminChanged)
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
		it.Event = new(AdvancedHomeAutomationRoleAdminChanged)
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
func (it *AdvancedHomeAutomationRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedHomeAutomationRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedHomeAutomationRoleAdminChanged represents a RoleAdminChanged event raised by the AdvancedHomeAutomation contract.
type AdvancedHomeAutomationRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AdvancedHomeAutomationRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AdvancedHomeAutomation.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedHomeAutomationRoleAdminChangedIterator{contract: _AdvancedHomeAutomation.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AdvancedHomeAutomationRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AdvancedHomeAutomation.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedHomeAutomationRoleAdminChanged)
				if err := _AdvancedHomeAutomation.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationFilterer) ParseRoleAdminChanged(log types.Log) (*AdvancedHomeAutomationRoleAdminChanged, error) {
	event := new(AdvancedHomeAutomationRoleAdminChanged)
	if err := _AdvancedHomeAutomation.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdvancedHomeAutomationRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AdvancedHomeAutomation contract.
type AdvancedHomeAutomationRoleGrantedIterator struct {
	Event *AdvancedHomeAutomationRoleGranted // Event containing the contract specifics and raw log

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
func (it *AdvancedHomeAutomationRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedHomeAutomationRoleGranted)
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
		it.Event = new(AdvancedHomeAutomationRoleGranted)
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
func (it *AdvancedHomeAutomationRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedHomeAutomationRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedHomeAutomationRoleGranted represents a RoleGranted event raised by the AdvancedHomeAutomation contract.
type AdvancedHomeAutomationRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AdvancedHomeAutomationRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AdvancedHomeAutomation.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedHomeAutomationRoleGrantedIterator{contract: _AdvancedHomeAutomation.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AdvancedHomeAutomationRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AdvancedHomeAutomation.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedHomeAutomationRoleGranted)
				if err := _AdvancedHomeAutomation.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationFilterer) ParseRoleGranted(log types.Log) (*AdvancedHomeAutomationRoleGranted, error) {
	event := new(AdvancedHomeAutomationRoleGranted)
	if err := _AdvancedHomeAutomation.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdvancedHomeAutomationRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AdvancedHomeAutomation contract.
type AdvancedHomeAutomationRoleRevokedIterator struct {
	Event *AdvancedHomeAutomationRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AdvancedHomeAutomationRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedHomeAutomationRoleRevoked)
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
		it.Event = new(AdvancedHomeAutomationRoleRevoked)
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
func (it *AdvancedHomeAutomationRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedHomeAutomationRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedHomeAutomationRoleRevoked represents a RoleRevoked event raised by the AdvancedHomeAutomation contract.
type AdvancedHomeAutomationRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AdvancedHomeAutomationRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AdvancedHomeAutomation.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedHomeAutomationRoleRevokedIterator{contract: _AdvancedHomeAutomation.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AdvancedHomeAutomationRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AdvancedHomeAutomation.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedHomeAutomationRoleRevoked)
				if err := _AdvancedHomeAutomation.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationFilterer) ParseRoleRevoked(log types.Log) (*AdvancedHomeAutomationRoleRevoked, error) {
	event := new(AdvancedHomeAutomationRoleRevoked)
	if err := _AdvancedHomeAutomation.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdvancedHomeAutomationRoomCreatedIterator is returned from FilterRoomCreated and is used to iterate over the raw logs and unpacked data for RoomCreated events raised by the AdvancedHomeAutomation contract.
type AdvancedHomeAutomationRoomCreatedIterator struct {
	Event *AdvancedHomeAutomationRoomCreated // Event containing the contract specifics and raw log

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
func (it *AdvancedHomeAutomationRoomCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedHomeAutomationRoomCreated)
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
		it.Event = new(AdvancedHomeAutomationRoomCreated)
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
func (it *AdvancedHomeAutomationRoomCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedHomeAutomationRoomCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedHomeAutomationRoomCreated represents a RoomCreated event raised by the AdvancedHomeAutomation contract.
type AdvancedHomeAutomationRoomCreated struct {
	RoomId *big.Int
	Name   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRoomCreated is a free log retrieval operation binding the contract event 0x37e1e413eec6f4b5fb3342770b4c3e7ad2b61d2300fe066b16ddb42eb3d9438b.
//
// Solidity: event RoomCreated(uint256 indexed roomId, string name)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationFilterer) FilterRoomCreated(opts *bind.FilterOpts, roomId []*big.Int) (*AdvancedHomeAutomationRoomCreatedIterator, error) {

	var roomIdRule []interface{}
	for _, roomIdItem := range roomId {
		roomIdRule = append(roomIdRule, roomIdItem)
	}

	logs, sub, err := _AdvancedHomeAutomation.contract.FilterLogs(opts, "RoomCreated", roomIdRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedHomeAutomationRoomCreatedIterator{contract: _AdvancedHomeAutomation.contract, event: "RoomCreated", logs: logs, sub: sub}, nil
}

// WatchRoomCreated is a free log subscription operation binding the contract event 0x37e1e413eec6f4b5fb3342770b4c3e7ad2b61d2300fe066b16ddb42eb3d9438b.
//
// Solidity: event RoomCreated(uint256 indexed roomId, string name)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationFilterer) WatchRoomCreated(opts *bind.WatchOpts, sink chan<- *AdvancedHomeAutomationRoomCreated, roomId []*big.Int) (event.Subscription, error) {

	var roomIdRule []interface{}
	for _, roomIdItem := range roomId {
		roomIdRule = append(roomIdRule, roomIdItem)
	}

	logs, sub, err := _AdvancedHomeAutomation.contract.WatchLogs(opts, "RoomCreated", roomIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedHomeAutomationRoomCreated)
				if err := _AdvancedHomeAutomation.contract.UnpackLog(event, "RoomCreated", log); err != nil {
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

// ParseRoomCreated is a log parse operation binding the contract event 0x37e1e413eec6f4b5fb3342770b4c3e7ad2b61d2300fe066b16ddb42eb3d9438b.
//
// Solidity: event RoomCreated(uint256 indexed roomId, string name)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationFilterer) ParseRoomCreated(log types.Log) (*AdvancedHomeAutomationRoomCreated, error) {
	event := new(AdvancedHomeAutomationRoomCreated)
	if err := _AdvancedHomeAutomation.contract.UnpackLog(event, "RoomCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdvancedHomeAutomationStateChangedIterator is returned from FilterStateChanged and is used to iterate over the raw logs and unpacked data for StateChanged events raised by the AdvancedHomeAutomation contract.
type AdvancedHomeAutomationStateChangedIterator struct {
	Event *AdvancedHomeAutomationStateChanged // Event containing the contract specifics and raw log

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
func (it *AdvancedHomeAutomationStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedHomeAutomationStateChanged)
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
		it.Event = new(AdvancedHomeAutomationStateChanged)
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
func (it *AdvancedHomeAutomationStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedHomeAutomationStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedHomeAutomationStateChanged represents a StateChanged event raised by the AdvancedHomeAutomation contract.
type AdvancedHomeAutomationStateChanged struct {
	RoomId   *big.Int
	DeviceId *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStateChanged is a free log retrieval operation binding the contract event 0x4e783b715efde58d097be2fe2f2e2bc0bb1df07fd682aeab1d8b4dc063ffdd9a.
//
// Solidity: event StateChanged(uint256 roomId, uint256 deviceId, uint256 newValue)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationFilterer) FilterStateChanged(opts *bind.FilterOpts) (*AdvancedHomeAutomationStateChangedIterator, error) {

	logs, sub, err := _AdvancedHomeAutomation.contract.FilterLogs(opts, "StateChanged")
	if err != nil {
		return nil, err
	}
	return &AdvancedHomeAutomationStateChangedIterator{contract: _AdvancedHomeAutomation.contract, event: "StateChanged", logs: logs, sub: sub}, nil
}

// WatchStateChanged is a free log subscription operation binding the contract event 0x4e783b715efde58d097be2fe2f2e2bc0bb1df07fd682aeab1d8b4dc063ffdd9a.
//
// Solidity: event StateChanged(uint256 roomId, uint256 deviceId, uint256 newValue)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationFilterer) WatchStateChanged(opts *bind.WatchOpts, sink chan<- *AdvancedHomeAutomationStateChanged) (event.Subscription, error) {

	logs, sub, err := _AdvancedHomeAutomation.contract.WatchLogs(opts, "StateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedHomeAutomationStateChanged)
				if err := _AdvancedHomeAutomation.contract.UnpackLog(event, "StateChanged", log); err != nil {
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

// ParseStateChanged is a log parse operation binding the contract event 0x4e783b715efde58d097be2fe2f2e2bc0bb1df07fd682aeab1d8b4dc063ffdd9a.
//
// Solidity: event StateChanged(uint256 roomId, uint256 deviceId, uint256 newValue)
func (_AdvancedHomeAutomation *AdvancedHomeAutomationFilterer) ParseStateChanged(log types.Log) (*AdvancedHomeAutomationStateChanged, error) {
	event := new(AdvancedHomeAutomationStateChanged)
	if err := _AdvancedHomeAutomation.contract.UnpackLog(event, "StateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
