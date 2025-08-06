// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// Todolistv2MetaData contains all meta data concerning the Todolistv2 contract.
var Todolistv2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assignee\",\"type\":\"address\"}],\"name\":\"TaskAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumTodolistV2.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assignee\",\"type\":\"address\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"TaskDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumTodolistV2.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"TaskStatusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"name\":\"TaskUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_taskid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_assignee\",\"type\":\"address\"}],\"name\":\"assignTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_assignee\",\"type\":\"address\"}],\"name\":\"createTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_taskid\",\"type\":\"uint256\"}],\"name\":\"deleteTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assignee\",\"type\":\"address\"}],\"name\":\"getAssignedTaskCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_taskid\",\"type\":\"uint256\"}],\"name\":\"getTask\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"enumTodolistV2.Status\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assignee\",\"type\":\"address\"}],\"name\":\"getTasksByAssignee\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskcount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"enumTodolistV2.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"assignee\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_taskid\",\"type\":\"uint256\"},{\"internalType\":\"enumTodolistV2.Status\",\"name\":\"_newStatus\",\"type\":\"uint8\"}],\"name\":\"updateStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_taskid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_newcontent\",\"type\":\"string\"}],\"name\":\"updateTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Todolistv2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Todolistv2MetaData.ABI instead.
var Todolistv2ABI = Todolistv2MetaData.ABI

// Todolistv2 is an auto generated Go binding around an Ethereum contract.
type Todolistv2 struct {
	Todolistv2Caller     // Read-only binding to the contract
	Todolistv2Transactor // Write-only binding to the contract
	Todolistv2Filterer   // Log filterer for contract events
}

// Todolistv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Todolistv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Todolistv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Todolistv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Todolistv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Todolistv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Todolistv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Todolistv2Session struct {
	Contract     *Todolistv2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Todolistv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Todolistv2CallerSession struct {
	Contract *Todolistv2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Todolistv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Todolistv2TransactorSession struct {
	Contract     *Todolistv2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Todolistv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Todolistv2Raw struct {
	Contract *Todolistv2 // Generic contract binding to access the raw methods on
}

// Todolistv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Todolistv2CallerRaw struct {
	Contract *Todolistv2Caller // Generic read-only contract binding to access the raw methods on
}

// Todolistv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Todolistv2TransactorRaw struct {
	Contract *Todolistv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTodolistv2 creates a new instance of Todolistv2, bound to a specific deployed contract.
func NewTodolistv2(address common.Address, backend bind.ContractBackend) (*Todolistv2, error) {
	contract, err := bindTodolistv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Todolistv2{Todolistv2Caller: Todolistv2Caller{contract: contract}, Todolistv2Transactor: Todolistv2Transactor{contract: contract}, Todolistv2Filterer: Todolistv2Filterer{contract: contract}}, nil
}

// NewTodolistv2Caller creates a new read-only instance of Todolistv2, bound to a specific deployed contract.
func NewTodolistv2Caller(address common.Address, caller bind.ContractCaller) (*Todolistv2Caller, error) {
	contract, err := bindTodolistv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Todolistv2Caller{contract: contract}, nil
}

// NewTodolistv2Transactor creates a new write-only instance of Todolistv2, bound to a specific deployed contract.
func NewTodolistv2Transactor(address common.Address, transactor bind.ContractTransactor) (*Todolistv2Transactor, error) {
	contract, err := bindTodolistv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Todolistv2Transactor{contract: contract}, nil
}

// NewTodolistv2Filterer creates a new log filterer instance of Todolistv2, bound to a specific deployed contract.
func NewTodolistv2Filterer(address common.Address, filterer bind.ContractFilterer) (*Todolistv2Filterer, error) {
	contract, err := bindTodolistv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Todolistv2Filterer{contract: contract}, nil
}

// bindTodolistv2 binds a generic wrapper to an already deployed contract.
func bindTodolistv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Todolistv2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todolistv2 *Todolistv2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todolistv2.Contract.Todolistv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todolistv2 *Todolistv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todolistv2.Contract.Todolistv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todolistv2 *Todolistv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todolistv2.Contract.Todolistv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todolistv2 *Todolistv2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todolistv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todolistv2 *Todolistv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todolistv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todolistv2 *Todolistv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todolistv2.Contract.contract.Transact(opts, method, params...)
}

// GetAssignedTaskCount is a free data retrieval call binding the contract method 0xa1045d4d.
//
// Solidity: function getAssignedTaskCount(address _assignee) view returns(uint256)
func (_Todolistv2 *Todolistv2Caller) GetAssignedTaskCount(opts *bind.CallOpts, _assignee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Todolistv2.contract.Call(opts, &out, "getAssignedTaskCount", _assignee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssignedTaskCount is a free data retrieval call binding the contract method 0xa1045d4d.
//
// Solidity: function getAssignedTaskCount(address _assignee) view returns(uint256)
func (_Todolistv2 *Todolistv2Session) GetAssignedTaskCount(_assignee common.Address) (*big.Int, error) {
	return _Todolistv2.Contract.GetAssignedTaskCount(&_Todolistv2.CallOpts, _assignee)
}

// GetAssignedTaskCount is a free data retrieval call binding the contract method 0xa1045d4d.
//
// Solidity: function getAssignedTaskCount(address _assignee) view returns(uint256)
func (_Todolistv2 *Todolistv2CallerSession) GetAssignedTaskCount(_assignee common.Address) (*big.Int, error) {
	return _Todolistv2.Contract.GetAssignedTaskCount(&_Todolistv2.CallOpts, _assignee)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Todolistv2 *Todolistv2Caller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Todolistv2.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Todolistv2 *Todolistv2Session) GetCount() (*big.Int, error) {
	return _Todolistv2.Contract.GetCount(&_Todolistv2.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Todolistv2 *Todolistv2CallerSession) GetCount() (*big.Int, error) {
	return _Todolistv2.Contract.GetCount(&_Todolistv2.CallOpts)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 _taskid) view returns(string, uint8, address)
func (_Todolistv2 *Todolistv2Caller) GetTask(opts *bind.CallOpts, _taskid *big.Int) (string, uint8, common.Address, error) {
	var out []interface{}
	err := _Todolistv2.contract.Call(opts, &out, "getTask", _taskid)

	if err != nil {
		return *new(string), *new(uint8), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return out0, out1, out2, err

}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 _taskid) view returns(string, uint8, address)
func (_Todolistv2 *Todolistv2Session) GetTask(_taskid *big.Int) (string, uint8, common.Address, error) {
	return _Todolistv2.Contract.GetTask(&_Todolistv2.CallOpts, _taskid)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 _taskid) view returns(string, uint8, address)
func (_Todolistv2 *Todolistv2CallerSession) GetTask(_taskid *big.Int) (string, uint8, common.Address, error) {
	return _Todolistv2.Contract.GetTask(&_Todolistv2.CallOpts, _taskid)
}

// GetTasksByAssignee is a free data retrieval call binding the contract method 0x8db6c002.
//
// Solidity: function getTasksByAssignee(address _assignee) view returns(uint256[])
func (_Todolistv2 *Todolistv2Caller) GetTasksByAssignee(opts *bind.CallOpts, _assignee common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Todolistv2.contract.Call(opts, &out, "getTasksByAssignee", _assignee)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTasksByAssignee is a free data retrieval call binding the contract method 0x8db6c002.
//
// Solidity: function getTasksByAssignee(address _assignee) view returns(uint256[])
func (_Todolistv2 *Todolistv2Session) GetTasksByAssignee(_assignee common.Address) ([]*big.Int, error) {
	return _Todolistv2.Contract.GetTasksByAssignee(&_Todolistv2.CallOpts, _assignee)
}

// GetTasksByAssignee is a free data retrieval call binding the contract method 0x8db6c002.
//
// Solidity: function getTasksByAssignee(address _assignee) view returns(uint256[])
func (_Todolistv2 *Todolistv2CallerSession) GetTasksByAssignee(_assignee common.Address) ([]*big.Int, error) {
	return _Todolistv2.Contract.GetTasksByAssignee(&_Todolistv2.CallOpts, _assignee)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todolistv2 *Todolistv2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Todolistv2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todolistv2 *Todolistv2Session) Owner() (common.Address, error) {
	return _Todolistv2.Contract.Owner(&_Todolistv2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todolistv2 *Todolistv2CallerSession) Owner() (common.Address, error) {
	return _Todolistv2.Contract.Owner(&_Todolistv2.CallOpts)
}

// Taskcount is a free data retrieval call binding the contract method 0x8bd48089.
//
// Solidity: function taskcount() view returns(uint256)
func (_Todolistv2 *Todolistv2Caller) Taskcount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Todolistv2.contract.Call(opts, &out, "taskcount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Taskcount is a free data retrieval call binding the contract method 0x8bd48089.
//
// Solidity: function taskcount() view returns(uint256)
func (_Todolistv2 *Todolistv2Session) Taskcount() (*big.Int, error) {
	return _Todolistv2.Contract.Taskcount(&_Todolistv2.CallOpts)
}

// Taskcount is a free data retrieval call binding the contract method 0x8bd48089.
//
// Solidity: function taskcount() view returns(uint256)
func (_Todolistv2 *Todolistv2CallerSession) Taskcount() (*big.Int, error) {
	return _Todolistv2.Contract.Taskcount(&_Todolistv2.CallOpts)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 id, string content, uint8 status, address assignee)
func (_Todolistv2 *Todolistv2Caller) Tasks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id       *big.Int
	Content  string
	Status   uint8
	Assignee common.Address
}, error) {
	var out []interface{}
	err := _Todolistv2.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		Id       *big.Int
		Content  string
		Status   uint8
		Assignee common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Content = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Assignee = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 id, string content, uint8 status, address assignee)
func (_Todolistv2 *Todolistv2Session) Tasks(arg0 *big.Int) (struct {
	Id       *big.Int
	Content  string
	Status   uint8
	Assignee common.Address
}, error) {
	return _Todolistv2.Contract.Tasks(&_Todolistv2.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 id, string content, uint8 status, address assignee)
func (_Todolistv2 *Todolistv2CallerSession) Tasks(arg0 *big.Int) (struct {
	Id       *big.Int
	Content  string
	Status   uint8
	Assignee common.Address
}, error) {
	return _Todolistv2.Contract.Tasks(&_Todolistv2.CallOpts, arg0)
}

// AssignTask is a paid mutator transaction binding the contract method 0x5293ee81.
//
// Solidity: function assignTask(uint256 _taskid, address _assignee) returns()
func (_Todolistv2 *Todolistv2Transactor) AssignTask(opts *bind.TransactOpts, _taskid *big.Int, _assignee common.Address) (*types.Transaction, error) {
	return _Todolistv2.contract.Transact(opts, "assignTask", _taskid, _assignee)
}

// AssignTask is a paid mutator transaction binding the contract method 0x5293ee81.
//
// Solidity: function assignTask(uint256 _taskid, address _assignee) returns()
func (_Todolistv2 *Todolistv2Session) AssignTask(_taskid *big.Int, _assignee common.Address) (*types.Transaction, error) {
	return _Todolistv2.Contract.AssignTask(&_Todolistv2.TransactOpts, _taskid, _assignee)
}

// AssignTask is a paid mutator transaction binding the contract method 0x5293ee81.
//
// Solidity: function assignTask(uint256 _taskid, address _assignee) returns()
func (_Todolistv2 *Todolistv2TransactorSession) AssignTask(_taskid *big.Int, _assignee common.Address) (*types.Transaction, error) {
	return _Todolistv2.Contract.AssignTask(&_Todolistv2.TransactOpts, _taskid, _assignee)
}

// CreateTask is a paid mutator transaction binding the contract method 0x71a244b3.
//
// Solidity: function createTask(string _content, address _assignee) returns()
func (_Todolistv2 *Todolistv2Transactor) CreateTask(opts *bind.TransactOpts, _content string, _assignee common.Address) (*types.Transaction, error) {
	return _Todolistv2.contract.Transact(opts, "createTask", _content, _assignee)
}

// CreateTask is a paid mutator transaction binding the contract method 0x71a244b3.
//
// Solidity: function createTask(string _content, address _assignee) returns()
func (_Todolistv2 *Todolistv2Session) CreateTask(_content string, _assignee common.Address) (*types.Transaction, error) {
	return _Todolistv2.Contract.CreateTask(&_Todolistv2.TransactOpts, _content, _assignee)
}

// CreateTask is a paid mutator transaction binding the contract method 0x71a244b3.
//
// Solidity: function createTask(string _content, address _assignee) returns()
func (_Todolistv2 *Todolistv2TransactorSession) CreateTask(_content string, _assignee common.Address) (*types.Transaction, error) {
	return _Todolistv2.Contract.CreateTask(&_Todolistv2.TransactOpts, _content, _assignee)
}

// DeleteTask is a paid mutator transaction binding the contract method 0x560f3192.
//
// Solidity: function deleteTask(uint256 _taskid) returns()
func (_Todolistv2 *Todolistv2Transactor) DeleteTask(opts *bind.TransactOpts, _taskid *big.Int) (*types.Transaction, error) {
	return _Todolistv2.contract.Transact(opts, "deleteTask", _taskid)
}

// DeleteTask is a paid mutator transaction binding the contract method 0x560f3192.
//
// Solidity: function deleteTask(uint256 _taskid) returns()
func (_Todolistv2 *Todolistv2Session) DeleteTask(_taskid *big.Int) (*types.Transaction, error) {
	return _Todolistv2.Contract.DeleteTask(&_Todolistv2.TransactOpts, _taskid)
}

// DeleteTask is a paid mutator transaction binding the contract method 0x560f3192.
//
// Solidity: function deleteTask(uint256 _taskid) returns()
func (_Todolistv2 *Todolistv2TransactorSession) DeleteTask(_taskid *big.Int) (*types.Transaction, error) {
	return _Todolistv2.Contract.DeleteTask(&_Todolistv2.TransactOpts, _taskid)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Todolistv2 *Todolistv2Transactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Todolistv2.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Todolistv2 *Todolistv2Session) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _Todolistv2.Contract.Initialize(&_Todolistv2.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Todolistv2 *Todolistv2TransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _Todolistv2.Contract.Initialize(&_Todolistv2.TransactOpts, _owner)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x3a1b3d31.
//
// Solidity: function updateStatus(uint256 _taskid, uint8 _newStatus) returns()
func (_Todolistv2 *Todolistv2Transactor) UpdateStatus(opts *bind.TransactOpts, _taskid *big.Int, _newStatus uint8) (*types.Transaction, error) {
	return _Todolistv2.contract.Transact(opts, "updateStatus", _taskid, _newStatus)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x3a1b3d31.
//
// Solidity: function updateStatus(uint256 _taskid, uint8 _newStatus) returns()
func (_Todolistv2 *Todolistv2Session) UpdateStatus(_taskid *big.Int, _newStatus uint8) (*types.Transaction, error) {
	return _Todolistv2.Contract.UpdateStatus(&_Todolistv2.TransactOpts, _taskid, _newStatus)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x3a1b3d31.
//
// Solidity: function updateStatus(uint256 _taskid, uint8 _newStatus) returns()
func (_Todolistv2 *Todolistv2TransactorSession) UpdateStatus(_taskid *big.Int, _newStatus uint8) (*types.Transaction, error) {
	return _Todolistv2.Contract.UpdateStatus(&_Todolistv2.TransactOpts, _taskid, _newStatus)
}

// UpdateTask is a paid mutator transaction binding the contract method 0x86533f6c.
//
// Solidity: function updateTask(uint256 _taskid, string _newcontent) returns()
func (_Todolistv2 *Todolistv2Transactor) UpdateTask(opts *bind.TransactOpts, _taskid *big.Int, _newcontent string) (*types.Transaction, error) {
	return _Todolistv2.contract.Transact(opts, "updateTask", _taskid, _newcontent)
}

// UpdateTask is a paid mutator transaction binding the contract method 0x86533f6c.
//
// Solidity: function updateTask(uint256 _taskid, string _newcontent) returns()
func (_Todolistv2 *Todolistv2Session) UpdateTask(_taskid *big.Int, _newcontent string) (*types.Transaction, error) {
	return _Todolistv2.Contract.UpdateTask(&_Todolistv2.TransactOpts, _taskid, _newcontent)
}

// UpdateTask is a paid mutator transaction binding the contract method 0x86533f6c.
//
// Solidity: function updateTask(uint256 _taskid, string _newcontent) returns()
func (_Todolistv2 *Todolistv2TransactorSession) UpdateTask(_taskid *big.Int, _newcontent string) (*types.Transaction, error) {
	return _Todolistv2.Contract.UpdateTask(&_Todolistv2.TransactOpts, _taskid, _newcontent)
}

// Todolistv2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Todolistv2 contract.
type Todolistv2InitializedIterator struct {
	Event *Todolistv2Initialized // Event containing the contract specifics and raw log

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
func (it *Todolistv2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Todolistv2Initialized)
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
		it.Event = new(Todolistv2Initialized)
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
func (it *Todolistv2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Todolistv2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Todolistv2Initialized represents a Initialized event raised by the Todolistv2 contract.
type Todolistv2Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Todolistv2 *Todolistv2Filterer) FilterInitialized(opts *bind.FilterOpts) (*Todolistv2InitializedIterator, error) {

	logs, sub, err := _Todolistv2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Todolistv2InitializedIterator{contract: _Todolistv2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Todolistv2 *Todolistv2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Todolistv2Initialized) (event.Subscription, error) {

	logs, sub, err := _Todolistv2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Todolistv2Initialized)
				if err := _Todolistv2.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Todolistv2 *Todolistv2Filterer) ParseInitialized(log types.Log) (*Todolistv2Initialized, error) {
	event := new(Todolistv2Initialized)
	if err := _Todolistv2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Todolistv2TaskAssignedIterator is returned from FilterTaskAssigned and is used to iterate over the raw logs and unpacked data for TaskAssigned events raised by the Todolistv2 contract.
type Todolistv2TaskAssignedIterator struct {
	Event *Todolistv2TaskAssigned // Event containing the contract specifics and raw log

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
func (it *Todolistv2TaskAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Todolistv2TaskAssigned)
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
		it.Event = new(Todolistv2TaskAssigned)
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
func (it *Todolistv2TaskAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Todolistv2TaskAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Todolistv2TaskAssigned represents a TaskAssigned event raised by the Todolistv2 contract.
type Todolistv2TaskAssigned struct {
	Id       *big.Int
	Assignee common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTaskAssigned is a free log retrieval operation binding the contract event 0x52476d55ecef5cf13caa64038f297fe6bbf865d9584a98b8722a15a6d5db128f.
//
// Solidity: event TaskAssigned(uint256 id, address assignee)
func (_Todolistv2 *Todolistv2Filterer) FilterTaskAssigned(opts *bind.FilterOpts) (*Todolistv2TaskAssignedIterator, error) {

	logs, sub, err := _Todolistv2.contract.FilterLogs(opts, "TaskAssigned")
	if err != nil {
		return nil, err
	}
	return &Todolistv2TaskAssignedIterator{contract: _Todolistv2.contract, event: "TaskAssigned", logs: logs, sub: sub}, nil
}

// WatchTaskAssigned is a free log subscription operation binding the contract event 0x52476d55ecef5cf13caa64038f297fe6bbf865d9584a98b8722a15a6d5db128f.
//
// Solidity: event TaskAssigned(uint256 id, address assignee)
func (_Todolistv2 *Todolistv2Filterer) WatchTaskAssigned(opts *bind.WatchOpts, sink chan<- *Todolistv2TaskAssigned) (event.Subscription, error) {

	logs, sub, err := _Todolistv2.contract.WatchLogs(opts, "TaskAssigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Todolistv2TaskAssigned)
				if err := _Todolistv2.contract.UnpackLog(event, "TaskAssigned", log); err != nil {
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

// ParseTaskAssigned is a log parse operation binding the contract event 0x52476d55ecef5cf13caa64038f297fe6bbf865d9584a98b8722a15a6d5db128f.
//
// Solidity: event TaskAssigned(uint256 id, address assignee)
func (_Todolistv2 *Todolistv2Filterer) ParseTaskAssigned(log types.Log) (*Todolistv2TaskAssigned, error) {
	event := new(Todolistv2TaskAssigned)
	if err := _Todolistv2.contract.UnpackLog(event, "TaskAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Todolistv2TaskCreatedIterator is returned from FilterTaskCreated and is used to iterate over the raw logs and unpacked data for TaskCreated events raised by the Todolistv2 contract.
type Todolistv2TaskCreatedIterator struct {
	Event *Todolistv2TaskCreated // Event containing the contract specifics and raw log

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
func (it *Todolistv2TaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Todolistv2TaskCreated)
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
		it.Event = new(Todolistv2TaskCreated)
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
func (it *Todolistv2TaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Todolistv2TaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Todolistv2TaskCreated represents a TaskCreated event raised by the Todolistv2 contract.
type Todolistv2TaskCreated struct {
	Id       *big.Int
	Content  string
	Status   uint8
	Assignee common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0x6c2a584b3d703f6052727cdfa09d068ea155d5e23f8ae17b6737660658d57a01.
//
// Solidity: event TaskCreated(uint256 id, string content, uint8 status, address assignee)
func (_Todolistv2 *Todolistv2Filterer) FilterTaskCreated(opts *bind.FilterOpts) (*Todolistv2TaskCreatedIterator, error) {

	logs, sub, err := _Todolistv2.contract.FilterLogs(opts, "TaskCreated")
	if err != nil {
		return nil, err
	}
	return &Todolistv2TaskCreatedIterator{contract: _Todolistv2.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0x6c2a584b3d703f6052727cdfa09d068ea155d5e23f8ae17b6737660658d57a01.
//
// Solidity: event TaskCreated(uint256 id, string content, uint8 status, address assignee)
func (_Todolistv2 *Todolistv2Filterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *Todolistv2TaskCreated) (event.Subscription, error) {

	logs, sub, err := _Todolistv2.contract.WatchLogs(opts, "TaskCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Todolistv2TaskCreated)
				if err := _Todolistv2.contract.UnpackLog(event, "TaskCreated", log); err != nil {
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

// ParseTaskCreated is a log parse operation binding the contract event 0x6c2a584b3d703f6052727cdfa09d068ea155d5e23f8ae17b6737660658d57a01.
//
// Solidity: event TaskCreated(uint256 id, string content, uint8 status, address assignee)
func (_Todolistv2 *Todolistv2Filterer) ParseTaskCreated(log types.Log) (*Todolistv2TaskCreated, error) {
	event := new(Todolistv2TaskCreated)
	if err := _Todolistv2.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Todolistv2TaskDeletedIterator is returned from FilterTaskDeleted and is used to iterate over the raw logs and unpacked data for TaskDeleted events raised by the Todolistv2 contract.
type Todolistv2TaskDeletedIterator struct {
	Event *Todolistv2TaskDeleted // Event containing the contract specifics and raw log

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
func (it *Todolistv2TaskDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Todolistv2TaskDeleted)
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
		it.Event = new(Todolistv2TaskDeleted)
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
func (it *Todolistv2TaskDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Todolistv2TaskDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Todolistv2TaskDeleted represents a TaskDeleted event raised by the Todolistv2 contract.
type Todolistv2TaskDeleted struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTaskDeleted is a free log retrieval operation binding the contract event 0xd078b251d950cc55c44340be1db732337ef4939a76e1367ee271ad2cb19c46af.
//
// Solidity: event TaskDeleted(uint256 id)
func (_Todolistv2 *Todolistv2Filterer) FilterTaskDeleted(opts *bind.FilterOpts) (*Todolistv2TaskDeletedIterator, error) {

	logs, sub, err := _Todolistv2.contract.FilterLogs(opts, "TaskDeleted")
	if err != nil {
		return nil, err
	}
	return &Todolistv2TaskDeletedIterator{contract: _Todolistv2.contract, event: "TaskDeleted", logs: logs, sub: sub}, nil
}

// WatchTaskDeleted is a free log subscription operation binding the contract event 0xd078b251d950cc55c44340be1db732337ef4939a76e1367ee271ad2cb19c46af.
//
// Solidity: event TaskDeleted(uint256 id)
func (_Todolistv2 *Todolistv2Filterer) WatchTaskDeleted(opts *bind.WatchOpts, sink chan<- *Todolistv2TaskDeleted) (event.Subscription, error) {

	logs, sub, err := _Todolistv2.contract.WatchLogs(opts, "TaskDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Todolistv2TaskDeleted)
				if err := _Todolistv2.contract.UnpackLog(event, "TaskDeleted", log); err != nil {
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

// ParseTaskDeleted is a log parse operation binding the contract event 0xd078b251d950cc55c44340be1db732337ef4939a76e1367ee271ad2cb19c46af.
//
// Solidity: event TaskDeleted(uint256 id)
func (_Todolistv2 *Todolistv2Filterer) ParseTaskDeleted(log types.Log) (*Todolistv2TaskDeleted, error) {
	event := new(Todolistv2TaskDeleted)
	if err := _Todolistv2.contract.UnpackLog(event, "TaskDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Todolistv2TaskStatusUpdatedIterator is returned from FilterTaskStatusUpdated and is used to iterate over the raw logs and unpacked data for TaskStatusUpdated events raised by the Todolistv2 contract.
type Todolistv2TaskStatusUpdatedIterator struct {
	Event *Todolistv2TaskStatusUpdated // Event containing the contract specifics and raw log

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
func (it *Todolistv2TaskStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Todolistv2TaskStatusUpdated)
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
		it.Event = new(Todolistv2TaskStatusUpdated)
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
func (it *Todolistv2TaskStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Todolistv2TaskStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Todolistv2TaskStatusUpdated represents a TaskStatusUpdated event raised by the Todolistv2 contract.
type Todolistv2TaskStatusUpdated struct {
	Id     *big.Int
	Status uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTaskStatusUpdated is a free log retrieval operation binding the contract event 0x9a4f521400d1b2f610a5c4845f80b4d79e3ec73d18e977ac9e0b0e6e741ec85f.
//
// Solidity: event TaskStatusUpdated(uint256 id, uint8 status)
func (_Todolistv2 *Todolistv2Filterer) FilterTaskStatusUpdated(opts *bind.FilterOpts) (*Todolistv2TaskStatusUpdatedIterator, error) {

	logs, sub, err := _Todolistv2.contract.FilterLogs(opts, "TaskStatusUpdated")
	if err != nil {
		return nil, err
	}
	return &Todolistv2TaskStatusUpdatedIterator{contract: _Todolistv2.contract, event: "TaskStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchTaskStatusUpdated is a free log subscription operation binding the contract event 0x9a4f521400d1b2f610a5c4845f80b4d79e3ec73d18e977ac9e0b0e6e741ec85f.
//
// Solidity: event TaskStatusUpdated(uint256 id, uint8 status)
func (_Todolistv2 *Todolistv2Filterer) WatchTaskStatusUpdated(opts *bind.WatchOpts, sink chan<- *Todolistv2TaskStatusUpdated) (event.Subscription, error) {

	logs, sub, err := _Todolistv2.contract.WatchLogs(opts, "TaskStatusUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Todolistv2TaskStatusUpdated)
				if err := _Todolistv2.contract.UnpackLog(event, "TaskStatusUpdated", log); err != nil {
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

// ParseTaskStatusUpdated is a log parse operation binding the contract event 0x9a4f521400d1b2f610a5c4845f80b4d79e3ec73d18e977ac9e0b0e6e741ec85f.
//
// Solidity: event TaskStatusUpdated(uint256 id, uint8 status)
func (_Todolistv2 *Todolistv2Filterer) ParseTaskStatusUpdated(log types.Log) (*Todolistv2TaskStatusUpdated, error) {
	event := new(Todolistv2TaskStatusUpdated)
	if err := _Todolistv2.contract.UnpackLog(event, "TaskStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Todolistv2TaskUpdatedIterator is returned from FilterTaskUpdated and is used to iterate over the raw logs and unpacked data for TaskUpdated events raised by the Todolistv2 contract.
type Todolistv2TaskUpdatedIterator struct {
	Event *Todolistv2TaskUpdated // Event containing the contract specifics and raw log

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
func (it *Todolistv2TaskUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Todolistv2TaskUpdated)
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
		it.Event = new(Todolistv2TaskUpdated)
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
func (it *Todolistv2TaskUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Todolistv2TaskUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Todolistv2TaskUpdated represents a TaskUpdated event raised by the Todolistv2 contract.
type Todolistv2TaskUpdated struct {
	Id      *big.Int
	Content string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTaskUpdated is a free log retrieval operation binding the contract event 0xde038a66d64727fc277f5f5f3d50b98d205e4ae705bfb8ac824cea4877b9be49.
//
// Solidity: event TaskUpdated(uint256 id, string content)
func (_Todolistv2 *Todolistv2Filterer) FilterTaskUpdated(opts *bind.FilterOpts) (*Todolistv2TaskUpdatedIterator, error) {

	logs, sub, err := _Todolistv2.contract.FilterLogs(opts, "TaskUpdated")
	if err != nil {
		return nil, err
	}
	return &Todolistv2TaskUpdatedIterator{contract: _Todolistv2.contract, event: "TaskUpdated", logs: logs, sub: sub}, nil
}

// WatchTaskUpdated is a free log subscription operation binding the contract event 0xde038a66d64727fc277f5f5f3d50b98d205e4ae705bfb8ac824cea4877b9be49.
//
// Solidity: event TaskUpdated(uint256 id, string content)
func (_Todolistv2 *Todolistv2Filterer) WatchTaskUpdated(opts *bind.WatchOpts, sink chan<- *Todolistv2TaskUpdated) (event.Subscription, error) {

	logs, sub, err := _Todolistv2.contract.WatchLogs(opts, "TaskUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Todolistv2TaskUpdated)
				if err := _Todolistv2.contract.UnpackLog(event, "TaskUpdated", log); err != nil {
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

// ParseTaskUpdated is a log parse operation binding the contract event 0xde038a66d64727fc277f5f5f3d50b98d205e4ae705bfb8ac824cea4877b9be49.
//
// Solidity: event TaskUpdated(uint256 id, string content)
func (_Todolistv2 *Todolistv2Filterer) ParseTaskUpdated(log types.Log) (*Todolistv2TaskUpdated, error) {
	event := new(Todolistv2TaskUpdated)
	if err := _Todolistv2.contract.UnpackLog(event, "TaskUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
