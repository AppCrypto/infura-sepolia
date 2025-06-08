// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hello

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

// HelloMetaData contains all meta data concerning the Hello contract.
var HelloMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getMsg\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_msg\",\"type\":\"string\"}],\"name\":\"setMsg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506106328061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c8063b5fdeb2314610038578063c4784fd414610056575b5f5ffd5b610040610072565b60405161004d9190610183565b60405180910390f35b610070600480360381019061006b91906102e0565b610101565b005b60605f805461008090610354565b80601f01602080910402602001604051908101604052809291908181526020018280546100ac90610354565b80156100f75780601f106100ce576101008083540402835291602001916100f7565b820191905f5260205f20905b8154815290600101906020018083116100da57829003601f168201915b5050505050905090565b805f908161010f919061052d565b5050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61015582610113565b61015f818561011d565b935061016f81856020860161012d565b6101788161013b565b840191505092915050565b5f6020820190508181035f83015261019b818461014b565b905092915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6101f28261013b565b810181811067ffffffffffffffff82111715610211576102106101bc565b5b80604052505050565b5f6102236101a3565b905061022f82826101e9565b919050565b5f67ffffffffffffffff82111561024e5761024d6101bc565b5b6102578261013b565b9050602081019050919050565b828183375f83830152505050565b5f61028461027f84610234565b61021a565b9050828152602081018484840111156102a05761029f6101b8565b5b6102ab848285610264565b509392505050565b5f82601f8301126102c7576102c66101b4565b5b81356102d7848260208601610272565b91505092915050565b5f602082840312156102f5576102f46101ac565b5b5f82013567ffffffffffffffff811115610312576103116101b0565b5b61031e848285016102b3565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061036b57607f821691505b60208210810361037e5761037d610327565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026103e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826103a5565b6103ea86836103a5565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61042e61042961042484610402565b61040b565b610402565b9050919050565b5f819050919050565b61044783610414565b61045b61045382610435565b8484546103b1565b825550505050565b5f5f905090565b610472610463565b61047d81848461043e565b505050565b5b818110156104a0576104955f8261046a565b600181019050610483565b5050565b601f8211156104e5576104b681610384565b6104bf84610396565b810160208510156104ce578190505b6104e26104da85610396565b830182610482565b50505b505050565b5f82821c905092915050565b5f6105055f19846008026104ea565b1980831691505092915050565b5f61051d83836104f6565b9150826002028217905092915050565b61053682610113565b67ffffffffffffffff81111561054f5761054e6101bc565b5b6105598254610354565b6105648282856104a4565b5f60209050601f831160018114610595575f8415610583578287015190505b61058d8582610512565b8655506105f4565b601f1984166105a386610384565b5f5b828110156105ca578489015182556001820191506020850194506020810190506105a5565b868310156105e757848901516105e3601f8916826104f6565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220e3c299d21bdf497380bf38c0e4ee2bd3891e6ca278accfbcd53d68ba7b11af2564736f6c634300081c0033",
}

// HelloABI is the input ABI used to generate the binding from.
// Deprecated: Use HelloMetaData.ABI instead.
var HelloABI = HelloMetaData.ABI

// HelloBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HelloMetaData.Bin instead.
var HelloBin = HelloMetaData.Bin

// DeployHello deploys a new Ethereum contract, binding an instance of Hello to it.
func DeployHello(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Hello, error) {
	parsed, err := HelloMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HelloBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Hello{HelloCaller: HelloCaller{contract: contract}, HelloTransactor: HelloTransactor{contract: contract}, HelloFilterer: HelloFilterer{contract: contract}}, nil
}

// Hello is an auto generated Go binding around an Ethereum contract.
type Hello struct {
	HelloCaller     // Read-only binding to the contract
	HelloTransactor // Write-only binding to the contract
	HelloFilterer   // Log filterer for contract events
}

// HelloCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelloCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelloTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelloFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelloSession struct {
	Contract     *Hello            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelloCallerSession struct {
	Contract *HelloCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HelloTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelloTransactorSession struct {
	Contract     *HelloTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelloRaw struct {
	Contract *Hello // Generic contract binding to access the raw methods on
}

// HelloCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelloCallerRaw struct {
	Contract *HelloCaller // Generic read-only contract binding to access the raw methods on
}

// HelloTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelloTransactorRaw struct {
	Contract *HelloTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHello creates a new instance of Hello, bound to a specific deployed contract.
func NewHello(address common.Address, backend bind.ContractBackend) (*Hello, error) {
	contract, err := bindHello(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hello{HelloCaller: HelloCaller{contract: contract}, HelloTransactor: HelloTransactor{contract: contract}, HelloFilterer: HelloFilterer{contract: contract}}, nil
}

// NewHelloCaller creates a new read-only instance of Hello, bound to a specific deployed contract.
func NewHelloCaller(address common.Address, caller bind.ContractCaller) (*HelloCaller, error) {
	contract, err := bindHello(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelloCaller{contract: contract}, nil
}

// NewHelloTransactor creates a new write-only instance of Hello, bound to a specific deployed contract.
func NewHelloTransactor(address common.Address, transactor bind.ContractTransactor) (*HelloTransactor, error) {
	contract, err := bindHello(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelloTransactor{contract: contract}, nil
}

// NewHelloFilterer creates a new log filterer instance of Hello, bound to a specific deployed contract.
func NewHelloFilterer(address common.Address, filterer bind.ContractFilterer) (*HelloFilterer, error) {
	contract, err := bindHello(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelloFilterer{contract: contract}, nil
}

// bindHello binds a generic wrapper to an already deployed contract.
func bindHello(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HelloMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hello *HelloRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hello.Contract.HelloCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hello *HelloRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hello.Contract.HelloTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hello *HelloRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hello.Contract.HelloTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hello *HelloCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hello.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hello *HelloTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hello.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hello *HelloTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hello.Contract.contract.Transact(opts, method, params...)
}

// GetMsg is a free data retrieval call binding the contract method 0xb5fdeb23.
//
// Solidity: function getMsg() view returns(string)
func (_Hello *HelloCaller) GetMsg(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hello.contract.Call(opts, &out, "getMsg")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMsg is a free data retrieval call binding the contract method 0xb5fdeb23.
//
// Solidity: function getMsg() view returns(string)
func (_Hello *HelloSession) GetMsg() (string, error) {
	return _Hello.Contract.GetMsg(&_Hello.CallOpts)
}

// GetMsg is a free data retrieval call binding the contract method 0xb5fdeb23.
//
// Solidity: function getMsg() view returns(string)
func (_Hello *HelloCallerSession) GetMsg() (string, error) {
	return _Hello.Contract.GetMsg(&_Hello.CallOpts)
}

// SetMsg is a paid mutator transaction binding the contract method 0xc4784fd4.
//
// Solidity: function setMsg(string _msg) returns()
func (_Hello *HelloTransactor) SetMsg(opts *bind.TransactOpts, _msg string) (*types.Transaction, error) {
	return _Hello.contract.Transact(opts, "setMsg", _msg)
}

// SetMsg is a paid mutator transaction binding the contract method 0xc4784fd4.
//
// Solidity: function setMsg(string _msg) returns()
func (_Hello *HelloSession) SetMsg(_msg string) (*types.Transaction, error) {
	return _Hello.Contract.SetMsg(&_Hello.TransactOpts, _msg)
}

// SetMsg is a paid mutator transaction binding the contract method 0xc4784fd4.
//
// Solidity: function setMsg(string _msg) returns()
func (_Hello *HelloTransactorSession) SetMsg(_msg string) (*types.Transaction, error) {
	return _Hello.Contract.SetMsg(&_Hello.TransactOpts, _msg)
}
