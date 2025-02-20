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

// Position is an auto generated low-level Go binding around an user-defined struct.
type Position struct {
	OracleKey  [32]byte
	Guess      *big.Int
	Long       bool
	Timeframe  *big.Int
	Timestamp  *big.Int
	IsOccupied bool
	Resolvable bool
}

// MasterMetaData contains all meta data concerning the Master contract.
var MasterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"oracle_addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"AIRQUALITY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"HOUR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REALESTATE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addCity\",\"inputs\":[{\"name\":\"city\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"realEstate\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"airQuality\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createPosition\",\"inputs\":[{\"name\":\"categoryId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cityName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeframeIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"long\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"currentTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fillActualValues\",\"inputs\":[{\"name\":\"playerAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"categoryId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cityName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeframeIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fillActualValuesAll\",\"inputs\":[{\"name\":\"categoryId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cityName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeframeIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fillActualValuesBatches\",\"inputs\":[{\"name\":\"categoryId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cityName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeframeIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actualValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActivePlayers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActualValue\",\"inputs\":[{\"name\":\"playerAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"categoryId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cityName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeframeIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCategoriesCount\",\"inputs\":[{\"name\":\"player\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"categoryId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCityCount\",\"inputs\":[{\"name\":\"player\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cityName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlayerPositions\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cityName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"categoryId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structPosition[]\",\"components\":[{\"name\":\"oracleKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"guess\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"long\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"timeframe\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isOccupied\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"resolvable\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlayerWinLoss\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"wins\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"losses\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"players\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"records\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"wins\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"losses\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initialized\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"numPositions\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolvePosition\",\"inputs\":[{\"name\":\"categoryId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cityName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeframeIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"timeframes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateTime\",\"inputs\":[{\"name\":\"newTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validCitiesAirQuality\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validCitiesRealEstate\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"PositionResolved\",\"inputs\":[{\"name\":\"player\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oracleKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"guess\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"long\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"timeframe\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"won\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"finalValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// MasterABI is the input ABI used to generate the binding from.
// Deprecated: Use MasterMetaData.ABI instead.
var MasterABI = MasterMetaData.ABI

// Master is an auto generated Go binding around an Ethereum contract.
type Master struct {
	MasterCaller     // Read-only binding to the contract
	MasterTransactor // Write-only binding to the contract
	MasterFilterer   // Log filterer for contract events
}

// MasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MasterSession struct {
	Contract     *Master           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MasterCallerSession struct {
	Contract *MasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MasterTransactorSession struct {
	Contract     *MasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MasterRaw struct {
	Contract *Master // Generic contract binding to access the raw methods on
}

// MasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MasterCallerRaw struct {
	Contract *MasterCaller // Generic read-only contract binding to access the raw methods on
}

// MasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MasterTransactorRaw struct {
	Contract *MasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMaster creates a new instance of Master, bound to a specific deployed contract.
func NewMaster(address common.Address, backend bind.ContractBackend) (*Master, error) {
	contract, err := bindMaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Master{MasterCaller: MasterCaller{contract: contract}, MasterTransactor: MasterTransactor{contract: contract}, MasterFilterer: MasterFilterer{contract: contract}}, nil
}

// NewMasterCaller creates a new read-only instance of Master, bound to a specific deployed contract.
func NewMasterCaller(address common.Address, caller bind.ContractCaller) (*MasterCaller, error) {
	contract, err := bindMaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MasterCaller{contract: contract}, nil
}

// NewMasterTransactor creates a new write-only instance of Master, bound to a specific deployed contract.
func NewMasterTransactor(address common.Address, transactor bind.ContractTransactor) (*MasterTransactor, error) {
	contract, err := bindMaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MasterTransactor{contract: contract}, nil
}

// NewMasterFilterer creates a new log filterer instance of Master, bound to a specific deployed contract.
func NewMasterFilterer(address common.Address, filterer bind.ContractFilterer) (*MasterFilterer, error) {
	contract, err := bindMaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MasterFilterer{contract: contract}, nil
}

// bindMaster binds a generic wrapper to an already deployed contract.
func bindMaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Master *MasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Master.Contract.MasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Master *MasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Master.Contract.MasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Master *MasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Master.Contract.MasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Master *MasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Master.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Master *MasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Master.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Master *MasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Master.Contract.contract.Transact(opts, method, params...)
}

// AIRQUALITY is a free data retrieval call binding the contract method 0xf2a311f9.
//
// Solidity: function AIRQUALITY() view returns(uint256)
func (_Master *MasterCaller) AIRQUALITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Master.contract.Call(opts, &out, "AIRQUALITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AIRQUALITY is a free data retrieval call binding the contract method 0xf2a311f9.
//
// Solidity: function AIRQUALITY() view returns(uint256)
func (_Master *MasterSession) AIRQUALITY() (*big.Int, error) {
	return _Master.Contract.AIRQUALITY(&_Master.CallOpts)
}

// AIRQUALITY is a free data retrieval call binding the contract method 0xf2a311f9.
//
// Solidity: function AIRQUALITY() view returns(uint256)
func (_Master *MasterCallerSession) AIRQUALITY() (*big.Int, error) {
	return _Master.Contract.AIRQUALITY(&_Master.CallOpts)
}

// HOUR is a free data retrieval call binding the contract method 0xa39dc9be.
//
// Solidity: function HOUR() view returns(uint256)
func (_Master *MasterCaller) HOUR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Master.contract.Call(opts, &out, "HOUR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HOUR is a free data retrieval call binding the contract method 0xa39dc9be.
//
// Solidity: function HOUR() view returns(uint256)
func (_Master *MasterSession) HOUR() (*big.Int, error) {
	return _Master.Contract.HOUR(&_Master.CallOpts)
}

// HOUR is a free data retrieval call binding the contract method 0xa39dc9be.
//
// Solidity: function HOUR() view returns(uint256)
func (_Master *MasterCallerSession) HOUR() (*big.Int, error) {
	return _Master.Contract.HOUR(&_Master.CallOpts)
}

// REALESTATE is a free data retrieval call binding the contract method 0x07f64618.
//
// Solidity: function REALESTATE() view returns(uint256)
func (_Master *MasterCaller) REALESTATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Master.contract.Call(opts, &out, "REALESTATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REALESTATE is a free data retrieval call binding the contract method 0x07f64618.
//
// Solidity: function REALESTATE() view returns(uint256)
func (_Master *MasterSession) REALESTATE() (*big.Int, error) {
	return _Master.Contract.REALESTATE(&_Master.CallOpts)
}

// REALESTATE is a free data retrieval call binding the contract method 0x07f64618.
//
// Solidity: function REALESTATE() view returns(uint256)
func (_Master *MasterCallerSession) REALESTATE() (*big.Int, error) {
	return _Master.Contract.REALESTATE(&_Master.CallOpts)
}

// CurrentTime is a free data retrieval call binding the contract method 0xd18e81b3.
//
// Solidity: function currentTime() view returns(uint256)
func (_Master *MasterCaller) CurrentTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Master.contract.Call(opts, &out, "currentTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTime is a free data retrieval call binding the contract method 0xd18e81b3.
//
// Solidity: function currentTime() view returns(uint256)
func (_Master *MasterSession) CurrentTime() (*big.Int, error) {
	return _Master.Contract.CurrentTime(&_Master.CallOpts)
}

// CurrentTime is a free data retrieval call binding the contract method 0xd18e81b3.
//
// Solidity: function currentTime() view returns(uint256)
func (_Master *MasterCallerSession) CurrentTime() (*big.Int, error) {
	return _Master.Contract.CurrentTime(&_Master.CallOpts)
}

// GetActivePlayers is a free data retrieval call binding the contract method 0x3c4f5a66.
//
// Solidity: function getActivePlayers() view returns(address[])
func (_Master *MasterCaller) GetActivePlayers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Master.contract.Call(opts, &out, "getActivePlayers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActivePlayers is a free data retrieval call binding the contract method 0x3c4f5a66.
//
// Solidity: function getActivePlayers() view returns(address[])
func (_Master *MasterSession) GetActivePlayers() ([]common.Address, error) {
	return _Master.Contract.GetActivePlayers(&_Master.CallOpts)
}

// GetActivePlayers is a free data retrieval call binding the contract method 0x3c4f5a66.
//
// Solidity: function getActivePlayers() view returns(address[])
func (_Master *MasterCallerSession) GetActivePlayers() ([]common.Address, error) {
	return _Master.Contract.GetActivePlayers(&_Master.CallOpts)
}

// GetActualValue is a free data retrieval call binding the contract method 0x04178b78.
//
// Solidity: function getActualValue(address playerAddr, uint256 categoryId, string cityName, uint256 timeframeIndex) view returns(uint256)
func (_Master *MasterCaller) GetActualValue(opts *bind.CallOpts, playerAddr common.Address, categoryId *big.Int, cityName string, timeframeIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Master.contract.Call(opts, &out, "getActualValue", playerAddr, categoryId, cityName, timeframeIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActualValue is a free data retrieval call binding the contract method 0x04178b78.
//
// Solidity: function getActualValue(address playerAddr, uint256 categoryId, string cityName, uint256 timeframeIndex) view returns(uint256)
func (_Master *MasterSession) GetActualValue(playerAddr common.Address, categoryId *big.Int, cityName string, timeframeIndex *big.Int) (*big.Int, error) {
	return _Master.Contract.GetActualValue(&_Master.CallOpts, playerAddr, categoryId, cityName, timeframeIndex)
}

// GetActualValue is a free data retrieval call binding the contract method 0x04178b78.
//
// Solidity: function getActualValue(address playerAddr, uint256 categoryId, string cityName, uint256 timeframeIndex) view returns(uint256)
func (_Master *MasterCallerSession) GetActualValue(playerAddr common.Address, categoryId *big.Int, cityName string, timeframeIndex *big.Int) (*big.Int, error) {
	return _Master.Contract.GetActualValue(&_Master.CallOpts, playerAddr, categoryId, cityName, timeframeIndex)
}

// GetCategoriesCount is a free data retrieval call binding the contract method 0xed3c1d69.
//
// Solidity: function getCategoriesCount(address player, uint256 categoryId) view returns(uint256)
func (_Master *MasterCaller) GetCategoriesCount(opts *bind.CallOpts, player common.Address, categoryId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Master.contract.Call(opts, &out, "getCategoriesCount", player, categoryId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCategoriesCount is a free data retrieval call binding the contract method 0xed3c1d69.
//
// Solidity: function getCategoriesCount(address player, uint256 categoryId) view returns(uint256)
func (_Master *MasterSession) GetCategoriesCount(player common.Address, categoryId *big.Int) (*big.Int, error) {
	return _Master.Contract.GetCategoriesCount(&_Master.CallOpts, player, categoryId)
}

// GetCategoriesCount is a free data retrieval call binding the contract method 0xed3c1d69.
//
// Solidity: function getCategoriesCount(address player, uint256 categoryId) view returns(uint256)
func (_Master *MasterCallerSession) GetCategoriesCount(player common.Address, categoryId *big.Int) (*big.Int, error) {
	return _Master.Contract.GetCategoriesCount(&_Master.CallOpts, player, categoryId)
}

// GetCityCount is a free data retrieval call binding the contract method 0x06722a06.
//
// Solidity: function getCityCount(address player, string cityName) view returns(uint256)
func (_Master *MasterCaller) GetCityCount(opts *bind.CallOpts, player common.Address, cityName string) (*big.Int, error) {
	var out []interface{}
	err := _Master.contract.Call(opts, &out, "getCityCount", player, cityName)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCityCount is a free data retrieval call binding the contract method 0x06722a06.
//
// Solidity: function getCityCount(address player, string cityName) view returns(uint256)
func (_Master *MasterSession) GetCityCount(player common.Address, cityName string) (*big.Int, error) {
	return _Master.Contract.GetCityCount(&_Master.CallOpts, player, cityName)
}

// GetCityCount is a free data retrieval call binding the contract method 0x06722a06.
//
// Solidity: function getCityCount(address player, string cityName) view returns(uint256)
func (_Master *MasterCallerSession) GetCityCount(player common.Address, cityName string) (*big.Int, error) {
	return _Master.Contract.GetCityCount(&_Master.CallOpts, player, cityName)
}

// GetPlayerPositions is a free data retrieval call binding the contract method 0xa0d029b9.
//
// Solidity: function getPlayerPositions(address addr, string cityName, uint256 categoryId) view returns((bytes32,uint256,bool,uint256,uint256,bool,bool)[])
func (_Master *MasterCaller) GetPlayerPositions(opts *bind.CallOpts, addr common.Address, cityName string, categoryId *big.Int) ([]Position, error) {
	var out []interface{}
	err := _Master.contract.Call(opts, &out, "getPlayerPositions", addr, cityName, categoryId)

	if err != nil {
		return *new([]Position), err
	}

	out0 := *abi.ConvertType(out[0], new([]Position)).(*[]Position)

	return out0, err

}

// GetPlayerPositions is a free data retrieval call binding the contract method 0xa0d029b9.
//
// Solidity: function getPlayerPositions(address addr, string cityName, uint256 categoryId) view returns((bytes32,uint256,bool,uint256,uint256,bool,bool)[])
func (_Master *MasterSession) GetPlayerPositions(addr common.Address, cityName string, categoryId *big.Int) ([]Position, error) {
	return _Master.Contract.GetPlayerPositions(&_Master.CallOpts, addr, cityName, categoryId)
}

// GetPlayerPositions is a free data retrieval call binding the contract method 0xa0d029b9.
//
// Solidity: function getPlayerPositions(address addr, string cityName, uint256 categoryId) view returns((bytes32,uint256,bool,uint256,uint256,bool,bool)[])
func (_Master *MasterCallerSession) GetPlayerPositions(addr common.Address, cityName string, categoryId *big.Int) ([]Position, error) {
	return _Master.Contract.GetPlayerPositions(&_Master.CallOpts, addr, cityName, categoryId)
}

// GetPlayerWinLoss is a free data retrieval call binding the contract method 0xb53affea.
//
// Solidity: function getPlayerWinLoss(address addr) view returns(uint256 wins, uint256 losses)
func (_Master *MasterCaller) GetPlayerWinLoss(opts *bind.CallOpts, addr common.Address) (struct {
	Wins   *big.Int
	Losses *big.Int
}, error) {
	var out []interface{}
	err := _Master.contract.Call(opts, &out, "getPlayerWinLoss", addr)

	outstruct := new(struct {
		Wins   *big.Int
		Losses *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Wins = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Losses = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPlayerWinLoss is a free data retrieval call binding the contract method 0xb53affea.
//
// Solidity: function getPlayerWinLoss(address addr) view returns(uint256 wins, uint256 losses)
func (_Master *MasterSession) GetPlayerWinLoss(addr common.Address) (struct {
	Wins   *big.Int
	Losses *big.Int
}, error) {
	return _Master.Contract.GetPlayerWinLoss(&_Master.CallOpts, addr)
}

// GetPlayerWinLoss is a free data retrieval call binding the contract method 0xb53affea.
//
// Solidity: function getPlayerWinLoss(address addr) view returns(uint256 wins, uint256 losses)
func (_Master *MasterCallerSession) GetPlayerWinLoss(addr common.Address) (struct {
	Wins   *big.Int
	Losses *big.Int
}, error) {
	return _Master.Contract.GetPlayerWinLoss(&_Master.CallOpts, addr)
}

// Players is a free data retrieval call binding the contract method 0xf71d96cb.
//
// Solidity: function players(uint256 ) view returns(address)
func (_Master *MasterCaller) Players(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Master.contract.Call(opts, &out, "players", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Players is a free data retrieval call binding the contract method 0xf71d96cb.
//
// Solidity: function players(uint256 ) view returns(address)
func (_Master *MasterSession) Players(arg0 *big.Int) (common.Address, error) {
	return _Master.Contract.Players(&_Master.CallOpts, arg0)
}

// Players is a free data retrieval call binding the contract method 0xf71d96cb.
//
// Solidity: function players(uint256 ) view returns(address)
func (_Master *MasterCallerSession) Players(arg0 *big.Int) (common.Address, error) {
	return _Master.Contract.Players(&_Master.CallOpts, arg0)
}

// Records is a free data retrieval call binding the contract method 0x469e9067.
//
// Solidity: function records(address ) view returns(uint256 wins, uint256 losses, bool initialized, uint256 numPositions)
func (_Master *MasterCaller) Records(opts *bind.CallOpts, arg0 common.Address) (struct {
	Wins         *big.Int
	Losses       *big.Int
	Initialized  bool
	NumPositions *big.Int
}, error) {
	var out []interface{}
	err := _Master.contract.Call(opts, &out, "records", arg0)

	outstruct := new(struct {
		Wins         *big.Int
		Losses       *big.Int
		Initialized  bool
		NumPositions *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Wins = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Losses = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Initialized = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.NumPositions = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Records is a free data retrieval call binding the contract method 0x469e9067.
//
// Solidity: function records(address ) view returns(uint256 wins, uint256 losses, bool initialized, uint256 numPositions)
func (_Master *MasterSession) Records(arg0 common.Address) (struct {
	Wins         *big.Int
	Losses       *big.Int
	Initialized  bool
	NumPositions *big.Int
}, error) {
	return _Master.Contract.Records(&_Master.CallOpts, arg0)
}

// Records is a free data retrieval call binding the contract method 0x469e9067.
//
// Solidity: function records(address ) view returns(uint256 wins, uint256 losses, bool initialized, uint256 numPositions)
func (_Master *MasterCallerSession) Records(arg0 common.Address) (struct {
	Wins         *big.Int
	Losses       *big.Int
	Initialized  bool
	NumPositions *big.Int
}, error) {
	return _Master.Contract.Records(&_Master.CallOpts, arg0)
}

// Timeframes is a free data retrieval call binding the contract method 0x5099b78d.
//
// Solidity: function timeframes(uint256 ) view returns(uint256)
func (_Master *MasterCaller) Timeframes(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Master.contract.Call(opts, &out, "timeframes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timeframes is a free data retrieval call binding the contract method 0x5099b78d.
//
// Solidity: function timeframes(uint256 ) view returns(uint256)
func (_Master *MasterSession) Timeframes(arg0 *big.Int) (*big.Int, error) {
	return _Master.Contract.Timeframes(&_Master.CallOpts, arg0)
}

// Timeframes is a free data retrieval call binding the contract method 0x5099b78d.
//
// Solidity: function timeframes(uint256 ) view returns(uint256)
func (_Master *MasterCallerSession) Timeframes(arg0 *big.Int) (*big.Int, error) {
	return _Master.Contract.Timeframes(&_Master.CallOpts, arg0)
}

// ValidCitiesAirQuality is a free data retrieval call binding the contract method 0x43bef3de.
//
// Solidity: function validCitiesAirQuality(string ) view returns(bool)
func (_Master *MasterCaller) ValidCitiesAirQuality(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Master.contract.Call(opts, &out, "validCitiesAirQuality", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidCitiesAirQuality is a free data retrieval call binding the contract method 0x43bef3de.
//
// Solidity: function validCitiesAirQuality(string ) view returns(bool)
func (_Master *MasterSession) ValidCitiesAirQuality(arg0 string) (bool, error) {
	return _Master.Contract.ValidCitiesAirQuality(&_Master.CallOpts, arg0)
}

// ValidCitiesAirQuality is a free data retrieval call binding the contract method 0x43bef3de.
//
// Solidity: function validCitiesAirQuality(string ) view returns(bool)
func (_Master *MasterCallerSession) ValidCitiesAirQuality(arg0 string) (bool, error) {
	return _Master.Contract.ValidCitiesAirQuality(&_Master.CallOpts, arg0)
}

// ValidCitiesRealEstate is a free data retrieval call binding the contract method 0x8c21945e.
//
// Solidity: function validCitiesRealEstate(string ) view returns(bool)
func (_Master *MasterCaller) ValidCitiesRealEstate(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Master.contract.Call(opts, &out, "validCitiesRealEstate", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidCitiesRealEstate is a free data retrieval call binding the contract method 0x8c21945e.
//
// Solidity: function validCitiesRealEstate(string ) view returns(bool)
func (_Master *MasterSession) ValidCitiesRealEstate(arg0 string) (bool, error) {
	return _Master.Contract.ValidCitiesRealEstate(&_Master.CallOpts, arg0)
}

// ValidCitiesRealEstate is a free data retrieval call binding the contract method 0x8c21945e.
//
// Solidity: function validCitiesRealEstate(string ) view returns(bool)
func (_Master *MasterCallerSession) ValidCitiesRealEstate(arg0 string) (bool, error) {
	return _Master.Contract.ValidCitiesRealEstate(&_Master.CallOpts, arg0)
}

// AddCity is a paid mutator transaction binding the contract method 0x9b262969.
//
// Solidity: function addCity(string city, bool realEstate, bool airQuality) returns()
func (_Master *MasterTransactor) AddCity(opts *bind.TransactOpts, city string, realEstate bool, airQuality bool) (*types.Transaction, error) {
	return _Master.contract.Transact(opts, "addCity", city, realEstate, airQuality)
}

// AddCity is a paid mutator transaction binding the contract method 0x9b262969.
//
// Solidity: function addCity(string city, bool realEstate, bool airQuality) returns()
func (_Master *MasterSession) AddCity(city string, realEstate bool, airQuality bool) (*types.Transaction, error) {
	return _Master.Contract.AddCity(&_Master.TransactOpts, city, realEstate, airQuality)
}

// AddCity is a paid mutator transaction binding the contract method 0x9b262969.
//
// Solidity: function addCity(string city, bool realEstate, bool airQuality) returns()
func (_Master *MasterTransactorSession) AddCity(city string, realEstate bool, airQuality bool) (*types.Transaction, error) {
	return _Master.Contract.AddCity(&_Master.TransactOpts, city, realEstate, airQuality)
}

// CreatePosition is a paid mutator transaction binding the contract method 0x66acf00a.
//
// Solidity: function createPosition(uint256 categoryId, string cityName, uint256 timeframeIndex, bool long) returns()
func (_Master *MasterTransactor) CreatePosition(opts *bind.TransactOpts, categoryId *big.Int, cityName string, timeframeIndex *big.Int, long bool) (*types.Transaction, error) {
	return _Master.contract.Transact(opts, "createPosition", categoryId, cityName, timeframeIndex, long)
}

// CreatePosition is a paid mutator transaction binding the contract method 0x66acf00a.
//
// Solidity: function createPosition(uint256 categoryId, string cityName, uint256 timeframeIndex, bool long) returns()
func (_Master *MasterSession) CreatePosition(categoryId *big.Int, cityName string, timeframeIndex *big.Int, long bool) (*types.Transaction, error) {
	return _Master.Contract.CreatePosition(&_Master.TransactOpts, categoryId, cityName, timeframeIndex, long)
}

// CreatePosition is a paid mutator transaction binding the contract method 0x66acf00a.
//
// Solidity: function createPosition(uint256 categoryId, string cityName, uint256 timeframeIndex, bool long) returns()
func (_Master *MasterTransactorSession) CreatePosition(categoryId *big.Int, cityName string, timeframeIndex *big.Int, long bool) (*types.Transaction, error) {
	return _Master.Contract.CreatePosition(&_Master.TransactOpts, categoryId, cityName, timeframeIndex, long)
}

// FillActualValues is a paid mutator transaction binding the contract method 0x31af5d9d.
//
// Solidity: function fillActualValues(address playerAddr, uint256 categoryId, string cityName, uint256 timeframeIndex, uint256 actualValue) returns()
func (_Master *MasterTransactor) FillActualValues(opts *bind.TransactOpts, playerAddr common.Address, categoryId *big.Int, cityName string, timeframeIndex *big.Int, actualValue *big.Int) (*types.Transaction, error) {
	return _Master.contract.Transact(opts, "fillActualValues", playerAddr, categoryId, cityName, timeframeIndex, actualValue)
}

// FillActualValues is a paid mutator transaction binding the contract method 0x31af5d9d.
//
// Solidity: function fillActualValues(address playerAddr, uint256 categoryId, string cityName, uint256 timeframeIndex, uint256 actualValue) returns()
func (_Master *MasterSession) FillActualValues(playerAddr common.Address, categoryId *big.Int, cityName string, timeframeIndex *big.Int, actualValue *big.Int) (*types.Transaction, error) {
	return _Master.Contract.FillActualValues(&_Master.TransactOpts, playerAddr, categoryId, cityName, timeframeIndex, actualValue)
}

// FillActualValues is a paid mutator transaction binding the contract method 0x31af5d9d.
//
// Solidity: function fillActualValues(address playerAddr, uint256 categoryId, string cityName, uint256 timeframeIndex, uint256 actualValue) returns()
func (_Master *MasterTransactorSession) FillActualValues(playerAddr common.Address, categoryId *big.Int, cityName string, timeframeIndex *big.Int, actualValue *big.Int) (*types.Transaction, error) {
	return _Master.Contract.FillActualValues(&_Master.TransactOpts, playerAddr, categoryId, cityName, timeframeIndex, actualValue)
}

// FillActualValuesAll is a paid mutator transaction binding the contract method 0x595e5ed8.
//
// Solidity: function fillActualValuesAll(uint256 categoryId, string cityName, uint256 timeframeIndex, uint256 actualValue) returns()
func (_Master *MasterTransactor) FillActualValuesAll(opts *bind.TransactOpts, categoryId *big.Int, cityName string, timeframeIndex *big.Int, actualValue *big.Int) (*types.Transaction, error) {
	return _Master.contract.Transact(opts, "fillActualValuesAll", categoryId, cityName, timeframeIndex, actualValue)
}

// FillActualValuesAll is a paid mutator transaction binding the contract method 0x595e5ed8.
//
// Solidity: function fillActualValuesAll(uint256 categoryId, string cityName, uint256 timeframeIndex, uint256 actualValue) returns()
func (_Master *MasterSession) FillActualValuesAll(categoryId *big.Int, cityName string, timeframeIndex *big.Int, actualValue *big.Int) (*types.Transaction, error) {
	return _Master.Contract.FillActualValuesAll(&_Master.TransactOpts, categoryId, cityName, timeframeIndex, actualValue)
}

// FillActualValuesAll is a paid mutator transaction binding the contract method 0x595e5ed8.
//
// Solidity: function fillActualValuesAll(uint256 categoryId, string cityName, uint256 timeframeIndex, uint256 actualValue) returns()
func (_Master *MasterTransactorSession) FillActualValuesAll(categoryId *big.Int, cityName string, timeframeIndex *big.Int, actualValue *big.Int) (*types.Transaction, error) {
	return _Master.Contract.FillActualValuesAll(&_Master.TransactOpts, categoryId, cityName, timeframeIndex, actualValue)
}

// FillActualValuesBatches is a paid mutator transaction binding the contract method 0xb8e795d6.
//
// Solidity: function fillActualValuesBatches(uint256 categoryId, string cityName, uint256 timeframeIndex, uint256 actualValue, uint256 startIndex, uint256 endIndex) returns()
func (_Master *MasterTransactor) FillActualValuesBatches(opts *bind.TransactOpts, categoryId *big.Int, cityName string, timeframeIndex *big.Int, actualValue *big.Int, startIndex *big.Int, endIndex *big.Int) (*types.Transaction, error) {
	return _Master.contract.Transact(opts, "fillActualValuesBatches", categoryId, cityName, timeframeIndex, actualValue, startIndex, endIndex)
}

// FillActualValuesBatches is a paid mutator transaction binding the contract method 0xb8e795d6.
//
// Solidity: function fillActualValuesBatches(uint256 categoryId, string cityName, uint256 timeframeIndex, uint256 actualValue, uint256 startIndex, uint256 endIndex) returns()
func (_Master *MasterSession) FillActualValuesBatches(categoryId *big.Int, cityName string, timeframeIndex *big.Int, actualValue *big.Int, startIndex *big.Int, endIndex *big.Int) (*types.Transaction, error) {
	return _Master.Contract.FillActualValuesBatches(&_Master.TransactOpts, categoryId, cityName, timeframeIndex, actualValue, startIndex, endIndex)
}

// FillActualValuesBatches is a paid mutator transaction binding the contract method 0xb8e795d6.
//
// Solidity: function fillActualValuesBatches(uint256 categoryId, string cityName, uint256 timeframeIndex, uint256 actualValue, uint256 startIndex, uint256 endIndex) returns()
func (_Master *MasterTransactorSession) FillActualValuesBatches(categoryId *big.Int, cityName string, timeframeIndex *big.Int, actualValue *big.Int, startIndex *big.Int, endIndex *big.Int) (*types.Transaction, error) {
	return _Master.Contract.FillActualValuesBatches(&_Master.TransactOpts, categoryId, cityName, timeframeIndex, actualValue, startIndex, endIndex)
}

// ResolvePosition is a paid mutator transaction binding the contract method 0x0d19e9a1.
//
// Solidity: function resolvePosition(uint256 categoryId, string cityName, uint256 timeframeIndex) returns()
func (_Master *MasterTransactor) ResolvePosition(opts *bind.TransactOpts, categoryId *big.Int, cityName string, timeframeIndex *big.Int) (*types.Transaction, error) {
	return _Master.contract.Transact(opts, "resolvePosition", categoryId, cityName, timeframeIndex)
}

// ResolvePosition is a paid mutator transaction binding the contract method 0x0d19e9a1.
//
// Solidity: function resolvePosition(uint256 categoryId, string cityName, uint256 timeframeIndex) returns()
func (_Master *MasterSession) ResolvePosition(categoryId *big.Int, cityName string, timeframeIndex *big.Int) (*types.Transaction, error) {
	return _Master.Contract.ResolvePosition(&_Master.TransactOpts, categoryId, cityName, timeframeIndex)
}

// ResolvePosition is a paid mutator transaction binding the contract method 0x0d19e9a1.
//
// Solidity: function resolvePosition(uint256 categoryId, string cityName, uint256 timeframeIndex) returns()
func (_Master *MasterTransactorSession) ResolvePosition(categoryId *big.Int, cityName string, timeframeIndex *big.Int) (*types.Transaction, error) {
	return _Master.Contract.ResolvePosition(&_Master.TransactOpts, categoryId, cityName, timeframeIndex)
}

// UpdateTime is a paid mutator transaction binding the contract method 0x6c59bd0c.
//
// Solidity: function updateTime(uint256 newTimestamp) returns()
func (_Master *MasterTransactor) UpdateTime(opts *bind.TransactOpts, newTimestamp *big.Int) (*types.Transaction, error) {
	return _Master.contract.Transact(opts, "updateTime", newTimestamp)
}

// UpdateTime is a paid mutator transaction binding the contract method 0x6c59bd0c.
//
// Solidity: function updateTime(uint256 newTimestamp) returns()
func (_Master *MasterSession) UpdateTime(newTimestamp *big.Int) (*types.Transaction, error) {
	return _Master.Contract.UpdateTime(&_Master.TransactOpts, newTimestamp)
}

// UpdateTime is a paid mutator transaction binding the contract method 0x6c59bd0c.
//
// Solidity: function updateTime(uint256 newTimestamp) returns()
func (_Master *MasterTransactorSession) UpdateTime(newTimestamp *big.Int) (*types.Transaction, error) {
	return _Master.Contract.UpdateTime(&_Master.TransactOpts, newTimestamp)
}

// MasterPositionResolvedIterator is returned from FilterPositionResolved and is used to iterate over the raw logs and unpacked data for PositionResolved events raised by the Master contract.
type MasterPositionResolvedIterator struct {
	Event *MasterPositionResolved // Event containing the contract specifics and raw log

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
func (it *MasterPositionResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterPositionResolved)
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
		it.Event = new(MasterPositionResolved)
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
func (it *MasterPositionResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterPositionResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterPositionResolved represents a PositionResolved event raised by the Master contract.
type MasterPositionResolved struct {
	Player     common.Address
	OracleKey  [32]byte
	Guess      *big.Int
	Long       bool
	Timeframe  *big.Int
	Timestamp  *big.Int
	Won        bool
	FinalValue *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPositionResolved is a free log retrieval operation binding the contract event 0x60b5c22d75b9304bd2d38e9c15505bf43a99e47c56443edc058f3ad90b445db0.
//
// Solidity: event PositionResolved(address indexed player, bytes32 indexed oracleKey, uint256 guess, bool long, uint256 timeframe, uint256 timestamp, bool won, uint256 finalValue)
func (_Master *MasterFilterer) FilterPositionResolved(opts *bind.FilterOpts, player []common.Address, oracleKey [][32]byte) (*MasterPositionResolvedIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var oracleKeyRule []interface{}
	for _, oracleKeyItem := range oracleKey {
		oracleKeyRule = append(oracleKeyRule, oracleKeyItem)
	}

	logs, sub, err := _Master.contract.FilterLogs(opts, "PositionResolved", playerRule, oracleKeyRule)
	if err != nil {
		return nil, err
	}
	return &MasterPositionResolvedIterator{contract: _Master.contract, event: "PositionResolved", logs: logs, sub: sub}, nil
}

// WatchPositionResolved is a free log subscription operation binding the contract event 0x60b5c22d75b9304bd2d38e9c15505bf43a99e47c56443edc058f3ad90b445db0.
//
// Solidity: event PositionResolved(address indexed player, bytes32 indexed oracleKey, uint256 guess, bool long, uint256 timeframe, uint256 timestamp, bool won, uint256 finalValue)
func (_Master *MasterFilterer) WatchPositionResolved(opts *bind.WatchOpts, sink chan<- *MasterPositionResolved, player []common.Address, oracleKey [][32]byte) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var oracleKeyRule []interface{}
	for _, oracleKeyItem := range oracleKey {
		oracleKeyRule = append(oracleKeyRule, oracleKeyItem)
	}

	logs, sub, err := _Master.contract.WatchLogs(opts, "PositionResolved", playerRule, oracleKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterPositionResolved)
				if err := _Master.contract.UnpackLog(event, "PositionResolved", log); err != nil {
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

// ParsePositionResolved is a log parse operation binding the contract event 0x60b5c22d75b9304bd2d38e9c15505bf43a99e47c56443edc058f3ad90b445db0.
//
// Solidity: event PositionResolved(address indexed player, bytes32 indexed oracleKey, uint256 guess, bool long, uint256 timeframe, uint256 timestamp, bool won, uint256 finalValue)
func (_Master *MasterFilterer) ParsePositionResolved(log types.Log) (*MasterPositionResolved, error) {
	event := new(MasterPositionResolved)
	if err := _Master.contract.UnpackLog(event, "PositionResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
