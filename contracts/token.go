// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TobaTokenABI is the input ABI used to generate the binding from.
const TobaTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"mintedAmount\",\"type\":\"uint256\"}],\"name\":\"mintToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"}]"

// TobaTokenBin is the compiled bytecode used for deploying new contracts.
const TobaTokenBin = `0x60606040526003805460ff19166012179055341561001c57600080fd5b604051610bdc380380610bdc833981016040528080519190602001805182019190602001805160008054600160a060020a033316600160a060020a03199091168117825560035460ff16600a0a870260048190559082526005602052604090912055909101905060018280516100969291602001906100b3565b5060028180516100aa9291602001906100b3565b5050505061014e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100f457805160ff1916838001178555610121565b82800160010185558215610121579182015b82811115610121578251825591602001919060010190610106565b5061012d929150610131565b5090565b61014b91905b8082111561012d5760008155600101610137565b90565b610a7f8061015d6000396000f300606060405236156100d85763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100dd578063095ea7b31461016757806318160ddd1461019d57806323b872dd146101c2578063313ce567146101ea57806342966c681461021357806370a082311461022957806379c650681461024857806379cc67901461026c5780638da5cb5b1461028e57806395d89b41146102bd578063a9059cbb146102d0578063cae9ca51146102f2578063dd62ed3e14610357578063f2fde38b1461037c575b600080fd5b34156100e857600080fd5b6100f061039b565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561012c578082015183820152602001610114565b50505050905090810190601f1680156101595780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561017257600080fd5b610189600160a060020a0360043516602435610439565b604051901515815260200160405180910390f35b34156101a857600080fd5b6101b0610469565b60405190815260200160405180910390f35b34156101cd57600080fd5b610189600160a060020a036004358116906024351660443561046f565b34156101f557600080fd5b6101fd6104e6565b60405160ff909116815260200160405180910390f35b341561021e57600080fd5b6101896004356104ef565b341561023457600080fd5b6101b0600160a060020a036004351661057a565b341561025357600080fd5b61026a600160a060020a036004351660243561058c565b005b341561027757600080fd5b610189600160a060020a036004351660243561064f565b341561029957600080fd5b6102a161072b565b604051600160a060020a03909116815260200160405180910390f35b34156102c857600080fd5b6100f061073a565b34156102db57600080fd5b61026a600160a060020a03600435166024356107a5565b34156102fd57600080fd5b61018960048035600160a060020a03169060248035919060649060443590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506107b495505050505050565b341561036257600080fd5b6101b0600160a060020a03600435811690602435166108e6565b341561038757600080fd5b61026a600160a060020a0360043516610903565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104315780601f1061040657610100808354040283529160200191610431565b820191906000526020600020905b81548152906001019060200180831161041457829003601f168201915b505050505081565b600160a060020a033381166000908152600660209081526040808320938616835292905220819055600192915050565b60045481565b600160a060020a038084166000908152600660209081526040808320339094168352929052908120548211156104a457600080fd5b600160a060020a03808516600090815260066020908152604080832033909416835292905220805483900390556104dc84848461094d565b5060019392505050565b60035460ff1681565b600160a060020a0333166000908152600560205260408120548290101561051557600080fd5b600160a060020a03331660008181526005602052604090819020805485900390556004805485900390557fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca59084905190815260200160405180910390a2506001919050565b60056020526000908152604090205481565b60005433600160a060020a039081169116146105a757600080fd5b600160a060020a0380831660009081526005602052604080822080548501905560048054850190558154909216917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9084905190815260200160405180910390a3600054600160a060020a0380841691167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405190815260200160405180910390a35050565b600160a060020a0382166000908152600560205260408120548290101561067557600080fd5b600160a060020a03808416600090815260066020908152604080832033909416835292905220548211156106a857600080fd5b600160a060020a038084166000818152600560209081526040808320805488900390556006825280832033909516835293905282902080548590039055600480548590039055907fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca59084905190815260200160405180910390a250600192915050565b600054600160a060020a031681565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104315780601f1061040657610100808354040283529160200191610431565b6107b033838361094d565b5050565b6000836107c18185610439565b156108de5780600160a060020a0316638f4ffcb1338630876040518563ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018085600160a060020a0316600160a060020a0316815260200184815260200183600160a060020a0316600160a060020a0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561087757808201518382015260200161085f565b50505050905090810190601f1680156108a45780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b15156108c557600080fd5b6102c65a03f115156108d657600080fd5b505050600191505b509392505050565b600660209081526000928352604080842090915290825290205481565b60005433600160a060020a0390811691161461091e57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000600160a060020a038316151561096457600080fd5b600160a060020a0384166000908152600560205260409020548290101561098a57600080fd5b600160a060020a038316600090815260056020526040902054828101116109b057600080fd5b50600160a060020a0380831660008181526005602052604080822080549488168084528284208054888103909155938590528154870190915591909301927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a3600160a060020a03808416600090815260056020526040808220549287168252902054018114610a4d57fe5b505050505600a165627a7a72305820e184f1240777a0990e28d067f1752ca536c06b580ea78ec9726c20c1034589250029`

// DeployTobaToken deploys a new Ethereum contract, binding an instance of TobaToken to it.
func DeployTobaToken(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, tokenName string, tokenSymbol string) (common.Address, *types.Transaction, *TobaToken, error) {
	parsed, err := abi.JSON(strings.NewReader(TobaTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TobaTokenBin), backend, initialSupply, tokenName, tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TobaToken{TobaTokenCaller: TobaTokenCaller{contract: contract}, TobaTokenTransactor: TobaTokenTransactor{contract: contract}}, nil
}

// TobaToken is an auto generated Go binding around an Ethereum contract.
type TobaToken struct {
	TobaTokenCaller     // Read-only binding to the contract
	TobaTokenTransactor // Write-only binding to the contract
}

// TobaTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TobaTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TobaTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TobaTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TobaTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TobaTokenSession struct {
	Contract     *TobaToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TobaTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TobaTokenCallerSession struct {
	Contract *TobaTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TobaTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TobaTokenTransactorSession struct {
	Contract     *TobaTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TobaTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TobaTokenRaw struct {
	Contract *TobaToken // Generic contract binding to access the raw methods on
}

// TobaTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TobaTokenCallerRaw struct {
	Contract *TobaTokenCaller // Generic read-only contract binding to access the raw methods on
}

// TobaTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TobaTokenTransactorRaw struct {
	Contract *TobaTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTobaToken creates a new instance of TobaToken, bound to a specific deployed contract.
func NewTobaToken(address common.Address, backend bind.ContractBackend) (*TobaToken, error) {
	contract, err := bindTobaToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TobaToken{TobaTokenCaller: TobaTokenCaller{contract: contract}, TobaTokenTransactor: TobaTokenTransactor{contract: contract}}, nil
}

// NewTobaTokenCaller creates a new read-only instance of TobaToken, bound to a specific deployed contract.
func NewTobaTokenCaller(address common.Address, caller bind.ContractCaller) (*TobaTokenCaller, error) {
	contract, err := bindTobaToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TobaTokenCaller{contract: contract}, nil
}

// NewTobaTokenTransactor creates a new write-only instance of TobaToken, bound to a specific deployed contract.
func NewTobaTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TobaTokenTransactor, error) {
	contract, err := bindTobaToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TobaTokenTransactor{contract: contract}, nil
}

// bindTobaToken binds a generic wrapper to an already deployed contract.
func bindTobaToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TobaTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TobaToken *TobaTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TobaToken.Contract.TobaTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TobaToken *TobaTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TobaToken.Contract.TobaTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TobaToken *TobaTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TobaToken.Contract.TobaTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TobaToken *TobaTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TobaToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TobaToken *TobaTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TobaToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TobaToken *TobaTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TobaToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_TobaToken *TobaTokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TobaToken.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_TobaToken *TobaTokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TobaToken.Contract.Allowance(&_TobaToken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_TobaToken *TobaTokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TobaToken.Contract.Allowance(&_TobaToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_TobaToken *TobaTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TobaToken.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_TobaToken *TobaTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _TobaToken.Contract.BalanceOf(&_TobaToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_TobaToken *TobaTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _TobaToken.Contract.BalanceOf(&_TobaToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_TobaToken *TobaTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _TobaToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_TobaToken *TobaTokenSession) Decimals() (uint8, error) {
	return _TobaToken.Contract.Decimals(&_TobaToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_TobaToken *TobaTokenCallerSession) Decimals() (uint8, error) {
	return _TobaToken.Contract.Decimals(&_TobaToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TobaToken *TobaTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TobaToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TobaToken *TobaTokenSession) Name() (string, error) {
	return _TobaToken.Contract.Name(&_TobaToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TobaToken *TobaTokenCallerSession) Name() (string, error) {
	return _TobaToken.Contract.Name(&_TobaToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TobaToken *TobaTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TobaToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TobaToken *TobaTokenSession) Owner() (common.Address, error) {
	return _TobaToken.Contract.Owner(&_TobaToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TobaToken *TobaTokenCallerSession) Owner() (common.Address, error) {
	return _TobaToken.Contract.Owner(&_TobaToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TobaToken *TobaTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TobaToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TobaToken *TobaTokenSession) Symbol() (string, error) {
	return _TobaToken.Contract.Symbol(&_TobaToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TobaToken *TobaTokenCallerSession) Symbol() (string, error) {
	return _TobaToken.Contract.Symbol(&_TobaToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TobaToken *TobaTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TobaToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TobaToken *TobaTokenSession) TotalSupply() (*big.Int, error) {
	return _TobaToken.Contract.TotalSupply(&_TobaToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TobaToken *TobaTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _TobaToken.Contract.TotalSupply(&_TobaToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_TobaToken *TobaTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TobaToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_TobaToken *TobaTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TobaToken.Contract.Approve(&_TobaToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_TobaToken *TobaTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TobaToken.Contract.Approve(&_TobaToken.TransactOpts, _spender, _value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_TobaToken *TobaTokenTransactor) ApproveAndCall(opts *bind.TransactOpts, _spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _TobaToken.contract.Transact(opts, "approveAndCall", _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_TobaToken *TobaTokenSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _TobaToken.Contract.ApproveAndCall(&_TobaToken.TransactOpts, _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_TobaToken *TobaTokenTransactorSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _TobaToken.Contract.ApproveAndCall(&_TobaToken.TransactOpts, _spender, _value, _extraData)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns(success bool)
func (_TobaToken *TobaTokenTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _TobaToken.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns(success bool)
func (_TobaToken *TobaTokenSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _TobaToken.Contract.Burn(&_TobaToken.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns(success bool)
func (_TobaToken *TobaTokenTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _TobaToken.Contract.Burn(&_TobaToken.TransactOpts, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(_from address, _value uint256) returns(success bool)
func (_TobaToken *TobaTokenTransactor) BurnFrom(opts *bind.TransactOpts, _from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TobaToken.contract.Transact(opts, "burnFrom", _from, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(_from address, _value uint256) returns(success bool)
func (_TobaToken *TobaTokenSession) BurnFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TobaToken.Contract.BurnFrom(&_TobaToken.TransactOpts, _from, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(_from address, _value uint256) returns(success bool)
func (_TobaToken *TobaTokenTransactorSession) BurnFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TobaToken.Contract.BurnFrom(&_TobaToken.TransactOpts, _from, _value)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(target address, mintedAmount uint256) returns()
func (_TobaToken *TobaTokenTransactor) MintToken(opts *bind.TransactOpts, target common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _TobaToken.contract.Transact(opts, "mintToken", target, mintedAmount)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(target address, mintedAmount uint256) returns()
func (_TobaToken *TobaTokenSession) MintToken(target common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _TobaToken.Contract.MintToken(&_TobaToken.TransactOpts, target, mintedAmount)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(target address, mintedAmount uint256) returns()
func (_TobaToken *TobaTokenTransactorSession) MintToken(target common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _TobaToken.Contract.MintToken(&_TobaToken.TransactOpts, target, mintedAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_TobaToken *TobaTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TobaToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_TobaToken *TobaTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TobaToken.Contract.Transfer(&_TobaToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_TobaToken *TobaTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TobaToken.Contract.Transfer(&_TobaToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_TobaToken *TobaTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TobaToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_TobaToken *TobaTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TobaToken.Contract.TransferFrom(&_TobaToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_TobaToken *TobaTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TobaToken.Contract.TransferFrom(&_TobaToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TobaToken *TobaTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TobaToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TobaToken *TobaTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TobaToken.Contract.TransferOwnership(&_TobaToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TobaToken *TobaTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TobaToken.Contract.TransferOwnership(&_TobaToken.TransactOpts, newOwner)
}

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OwnedBin is the compiled bytecode used for deploying new contracts.
const OwnedBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a031990911617905561015d8061003b6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610047578063f2fde38b1461008357600080fd5b341561005257600080fd5b61005a6100b1565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561008e57600080fd5b6100af73ffffffffffffffffffffffffffffffffffffffff600435166100cd565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146100f557600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff929092169190911790555600a165627a7a7230582052e190e768dd91be2f6b6b471c8969981648cd41ef105302773287067afa5ceb0029`

// DeployOwned deploys a new Ethereum contract, binding an instance of Owned to it.
func DeployOwned(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Owned, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}}, nil
}

// Owned is an auto generated Go binding around an Ethereum contract.
type Owned struct {
	OwnedCaller     // Read-only binding to the contract
	OwnedTransactor // Write-only binding to the contract
}

// OwnedCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnedSession struct {
	Contract     *Owned            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnedCallerSession struct {
	Contract *OwnedCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnedTransactorSession struct {
	Contract     *OwnedTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnedRaw struct {
	Contract *Owned // Generic contract binding to access the raw methods on
}

// OwnedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnedCallerRaw struct {
	Contract *OwnedCaller // Generic read-only contract binding to access the raw methods on
}

// OwnedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnedTransactorRaw struct {
	Contract *OwnedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwned creates a new instance of Owned, bound to a specific deployed contract.
func NewOwned(address common.Address, backend bind.ContractBackend) (*Owned, error) {
	contract, err := bindOwned(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}}, nil
}

// NewOwnedCaller creates a new read-only instance of Owned, bound to a specific deployed contract.
func NewOwnedCaller(address common.Address, caller bind.ContractCaller) (*OwnedCaller, error) {
	contract, err := bindOwned(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedCaller{contract: contract}, nil
}

// NewOwnedTransactor creates a new write-only instance of Owned, bound to a specific deployed contract.
func NewOwnedTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnedTransactor, error) {
	contract, err := bindOwned(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &OwnedTransactor{contract: contract}, nil
}

// bindOwned binds a generic wrapper to an already deployed contract.
func bindOwned(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.OwnedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Owned.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCallerSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Owned *OwnedTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Owned *OwnedSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.TransferOwnership(&_Owned.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Owned *OwnedTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.TransferOwnership(&_Owned.TransactOpts, newOwner)
}

// TokenRecipientABI is the input ABI used to generate the binding from.
const TokenRecipientABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"receiveApproval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TokenRecipientBin is the compiled bytecode used for deploying new contracts.
const TokenRecipientBin = `0x`

// DeployTokenRecipient deploys a new Ethereum contract, binding an instance of TokenRecipient to it.
func DeployTokenRecipient(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenRecipient, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenRecipientABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenRecipientBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenRecipient{TokenRecipientCaller: TokenRecipientCaller{contract: contract}, TokenRecipientTransactor: TokenRecipientTransactor{contract: contract}}, nil
}

// TokenRecipient is an auto generated Go binding around an Ethereum contract.
type TokenRecipient struct {
	TokenRecipientCaller     // Read-only binding to the contract
	TokenRecipientTransactor // Write-only binding to the contract
}

// TokenRecipientCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenRecipientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRecipientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenRecipientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRecipientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenRecipientSession struct {
	Contract     *TokenRecipient   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRecipientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenRecipientCallerSession struct {
	Contract *TokenRecipientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TokenRecipientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenRecipientTransactorSession struct {
	Contract     *TokenRecipientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TokenRecipientRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRecipientRaw struct {
	Contract *TokenRecipient // Generic contract binding to access the raw methods on
}

// TokenRecipientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenRecipientCallerRaw struct {
	Contract *TokenRecipientCaller // Generic read-only contract binding to access the raw methods on
}

// TokenRecipientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenRecipientTransactorRaw struct {
	Contract *TokenRecipientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenRecipient creates a new instance of TokenRecipient, bound to a specific deployed contract.
func NewTokenRecipient(address common.Address, backend bind.ContractBackend) (*TokenRecipient, error) {
	contract, err := bindTokenRecipient(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenRecipient{TokenRecipientCaller: TokenRecipientCaller{contract: contract}, TokenRecipientTransactor: TokenRecipientTransactor{contract: contract}}, nil
}

// NewTokenRecipientCaller creates a new read-only instance of TokenRecipient, bound to a specific deployed contract.
func NewTokenRecipientCaller(address common.Address, caller bind.ContractCaller) (*TokenRecipientCaller, error) {
	contract, err := bindTokenRecipient(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRecipientCaller{contract: contract}, nil
}

// NewTokenRecipientTransactor creates a new write-only instance of TokenRecipient, bound to a specific deployed contract.
func NewTokenRecipientTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenRecipientTransactor, error) {
	contract, err := bindTokenRecipient(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TokenRecipientTransactor{contract: contract}, nil
}

// bindTokenRecipient binds a generic wrapper to an already deployed contract.
func bindTokenRecipient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenRecipientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRecipient *TokenRecipientRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenRecipient.Contract.TokenRecipientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRecipient *TokenRecipientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRecipient.Contract.TokenRecipientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRecipient *TokenRecipientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRecipient.Contract.TokenRecipientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRecipient *TokenRecipientCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenRecipient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRecipient *TokenRecipientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRecipient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRecipient *TokenRecipientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRecipient.Contract.contract.Transact(opts, method, params...)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(_from address, _value uint256, _token address, _extraData bytes) returns()
func (_TokenRecipient *TokenRecipientTransactor) ReceiveApproval(opts *bind.TransactOpts, _from common.Address, _value *big.Int, _token common.Address, _extraData []byte) (*types.Transaction, error) {
	return _TokenRecipient.contract.Transact(opts, "receiveApproval", _from, _value, _token, _extraData)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(_from address, _value uint256, _token address, _extraData bytes) returns()
func (_TokenRecipient *TokenRecipientSession) ReceiveApproval(_from common.Address, _value *big.Int, _token common.Address, _extraData []byte) (*types.Transaction, error) {
	return _TokenRecipient.Contract.ReceiveApproval(&_TokenRecipient.TransactOpts, _from, _value, _token, _extraData)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(_from address, _value uint256, _token address, _extraData bytes) returns()
func (_TokenRecipient *TokenRecipientTransactorSession) ReceiveApproval(_from common.Address, _value *big.Int, _token common.Address, _extraData []byte) (*types.Transaction, error) {
	return _TokenRecipient.Contract.ReceiveApproval(&_TokenRecipient.TransactOpts, _from, _value, _token, _extraData)
}
