// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	"github.com/vntchain/go-vnt/accounts/abi"
	"github.com/vntchain/go-vnt/accounts/abi/bind"
	"github.com/vntchain/go-vnt/common"
	"github.com/vntchain/go-vnt/core/types"
)

// FIFSRegistrarABI is the input ABI used to generate the binding from.
const FIFSRegistrarABI = "[{\"name\":\"FIFSRegistrar\",\"constant\":false,\"inputs\":[{\"name\":\"ensAddr\",\"type\":\"address\",\"indexed\":false},{\"name\":\"node\",\"type\":\"string\",\"indexed\":false}],\"outputs\":[],\"type\":\"constructor\"},{\"name\":\"registernode\",\"constant\":false,\"inputs\":[{\"name\":\"subnode\",\"type\":\"string\",\"indexed\":false},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false}],\"outputs\":[],\"type\":\"function\"},{\"name\":\"setSubnodeOwner\",\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"string\",\"indexed\":false},{\"name\":\"label\",\"type\":\"string\",\"indexed\":false},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false}],\"outputs\":[],\"type\":\"call\"},{\"name\":\"owner\",\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"string\",\"indexed\":false}],\"outputs\":[{\"name\":\"output\",\"type\":\"address\",\"indexed\":false}],\"type\":\"call\"}]"

// FIFSRegistrarBin is the compiled bytecode used for deploying new contracts.
const FIFSRegistrarBin = `0x0161736db909140100789cbc576f8c1c6519ffbdef3befdceeceedde5e177abda370c35e91a2c7eef6fed64a812dd0722a85b4d682a6eecdedceee0ddd9b3966660fae86dbe3288880063f9898180b49a32144b1317cb04762fc66301ae317484c306af11f441244414210f3ccccdeeef640eb1fdc0ff3fedee7cfef799ef79977b2cf5b5bfe74b6ef1c0c6f810160138959d6c42c5f59c12c5bc1ac5c69ae349b98056679b3c99ab38c1e60cd59a5d924056f36c19f63bdc2b497d443e692e9fa006dd2475dcb378f5afefc6d8e65fba60b46e2be43a651d924558b95ca01c303a74da258a97cc25c9eb1ab0e04093233b6e55b46dd3a617eda702d63ae6e7a5002bf1b1cbb6cf890b4510edf5c1c874a503a77dba61b8ae54d77358c7a88e3074cffb069574c173d41369ee91f6eccd94ec5bc35f08809a9284a3cae48b6c818938281ab178b262bfef021456ba2b57c3fa6f51c55d50573c1719739b45269de34164b7386670a9628952a866f944cbb2278ba6296eb866b56f637ecb26f393692c9fd33fb0f1f326b96e7bb868b54af1b60d3a53cd097382d140eed32ecc33ea4ef5b5d5d0dd15a88d2f7d3aa1d646ca578e6a1d5557c27cb8b3d3a2fc68b489f245df1999658ed14a71f0894d031c503c7227436c50363edbb5cb0265b61cd91d5c0aa703ccb7784300a228a3dbad81c4414d52e3176868c3ad20f0651bf40cfcb159d17d377166377ee7be24beab480ce8bab6436c577e89c9cce843b5de7a1c54e9debfc2aa14f8b34bf3e70d695f443c495558a8f9171fa8bb44b4257d20faf76ec8bdf0cb441f29aceffa398080da24a74967e84d8745e2cdc199e8aa6ddc4515c8d6985c6d8e4d4aec9f1e95dbbc7a6a7a70a302a15d7f4bc0ed13d05d8a659d1c39712c533312d86d6efa7aa96ab98738d5ac9b2abce5e0128815ca1fb885e8c44765f8fd62701f05f00982620a4282880b8863c0e003847b7785c7e1280fa4b7a30fe0a800a192b524c90b1d532fe0b19177a28527f1422769a1e4f53a4b804701980af3280bd9a789bac8f314dd0f6d83a002d1340e2839a640067bd4446dcdf22d51b3fe180f67a1bbeb601937f25f83782290a3c09a08f22bc154647c05405f03d327ca7cdf44e9b29807d0113d806d500801451bddb514d9a349f6a05e12c8ad2cfa2d83f20123520d9f27a077ead8d33ad1eb05897fbf5007e4c665a87bbd6e11ee230cf649b6a1b51a55a541711d518e5cb818b07e834136ceb0e0007013c4ffb94f666b050b7b5c900d2d90f9047813a9c460f07f01480320762c739b06d891ef7d1e351522646c8712bd3ae0e562248065cdb02b887e01045481e2038f83572524e73e061caffe5c134fbedd04106900dfb3d868e32806a627f50866a0c7881f01f63e12b56229f31e5371cb89fe0a85a67404f6cab0f20cf809728c84898cc0865a0bd1bc07606575206418db7518d575c42a52d12fcd07905067d4e45d769c128d38d02becc9391c835ec9ae9d1ab9e01300ce02368ff4601e4a21e5cdb21a7164c01d80de01600b70238d2a17f34ba3adfa07715c03300d63bf4efa59b61c0ed0cf81c03ca0c30197092010fb2b61fc12d00e81a8e44fb9ba3b516ad53d15a8cd6f1285f978e07c0b701d0499feee0c5df370ec4989b73cd25d6cfae4865a448a5072e49f5b32303009f8048cd64ae1bdca37d4ce3db003189990ca06cc54c46a4480ac8be40a6ee80485dab691ad09363fd6c8bba2b03c4246889e74092d0e5ba41209163e1e6cac1eb06750dd024420171f586042d0bca00484e7458a42e65bb326473bb768706f4453a49ba74874e02fd1d64b245b645763864c8e1c840687e51982a657d710ead8c29c5addd5911cdc0a5883c3560db4450eb6086695a783443c9e0fc68b35b032ed988a351031ee8697df3eb966d36826ffeebd441c6de4e256961613f59fe8867ba5efe6ed3f20cebc4bc61d7f2373ae5c68269fb5ebee6e43db79caf59fe7c632e577616f24bb65f9e372c3b5f73ae5eb2fd7cd9b17dd728fb5edeb4bd8d5d3e870f86d769f88b0d1f58b2fdba35979ba712bafeeb94165db3ec2c2c5a7533570e5e50484eb58a17c1643c71f9bc2a1e6fae49a68a73d0390363909c2e807895312913773c7ebf64b59622f07c8331d99fd82b93cfca2bd53fcbf16be4874fc9b82a7eb7b2762db1ad4ba98ab799ce4587d33b8cc96462afccaa319d8c7fb5a21742e3a42ad6b82eb3eac765bf3a2ab34fc831754d0e9e92c92a19beb8526db39ee4ddac0f7262bd5c6e579f27db175674d92b9e623929d553e2cd7b8348bf06099f644470aaab9047f84621dbd57599d5e5153939b14e4e3f5fd1a51467584e15afdd7b4d8b468aa7d97a9080c25b24712659626fac160b0f2950d00741a448b15bc65f914c2ce76a32feb264e244eeb88cd764ea39c9bed2b6a7be88ad4c0e27627be536f531995993dbd59372f0ac1cfe911ccacb41d597dbd5676522f61999539f976339b93f27c7aa325e550baa985ba706ea3aef6d31d227488c5006a216f3e5b048be2499f07235392c7a09fa3942d5ae68e736453bf93ed1a42a76ea72f0ac7073225595c3e2aedbe590e8cb05ae527d560e8bc57579a94853277cc9d4c783d4f0989288eea1e7bbe5ba61d7f425d3f52cc7d677e70ab982bed3771bf6717d7c62d747a70a577da0b726ff4f2eca7f1717a6eda13c6fb8adbfa6701dc73f48b386e7bb965dc321f3ae86e59a283b76c50a6694d23ec7a963c1abe1b8b93c3e7d7c697cdeabc2b1ebcbe19ce48543138247b9e1baa6ed879a45c335163c945ab14ac682d3b07d342cdb1f9b9c42a96678a83b764d0f1e0ddbb36ab659d1adc8666a023718f5fa6d21cda6e1e9c8d8e4d461a36aded2a86fe01bada50d7cb831b7818b954af7e7075d83d632eec1093a9c62a5e222fa73fe02536c63c1643f636944f32cdb34caf2f3a758110db04ac7ec2adf6b6c55a389b527185663e19c1a0f47d4447b3ad5ce1f4c7bfb4ba5bb0d6fa15436eaf552d9775c2fb969b24c750f967d5d73e559213efbf92cd596dd93edb2cb8e66cb8eedf986ed67f7548dba678e662d7bb1e17bd93d6d97e89cb2a3597f7991045183b3645c31ef312b91f3bda31b3e14b7ed10be6c9bed8f8d66c32b40f18e6d980739b90daa35dbc1d959d385651ebdabff3a917690a02d1750eafba45e8d1ad299f7790dbdb0d4ffddbcebc69c59ff3f164a6f63f63de8fe67a575c56d87094417967657a6c756ff010000ffff`

// DeployFIFSRegistrar deploys a new VNT contract, binding an instance of FIFSRegistrar to it.
func DeployFIFSRegistrar(auth *bind.TransactOpts, backend bind.ContractBackend, ensAddr common.Address, node string) (common.Address, *types.Transaction, *FIFSRegistrar, error) {
	parsed, err := abi.JSON(strings.NewReader(FIFSRegistrarABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FIFSRegistrarBin), backend, ensAddr, node)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FIFSRegistrar{FIFSRegistrarCaller: FIFSRegistrarCaller{contract: contract}, FIFSRegistrarTransactor: FIFSRegistrarTransactor{contract: contract}, FIFSRegistrarFilterer: FIFSRegistrarFilterer{contract: contract}}, nil
}

// FIFSRegistrar is an auto generated Go binding around an VNT contract.
type FIFSRegistrar struct {
	FIFSRegistrarCaller     // Read-only binding to the contract
	FIFSRegistrarTransactor // Write-only binding to the contract
	FIFSRegistrarFilterer   // Log filterer for contract events
}

// FIFSRegistrarCaller is an auto generated read-only Go binding around an VNT contract.
type FIFSRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FIFSRegistrarTransactor is an auto generated write-only Go binding around an VNT contract.
type FIFSRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FIFSRegistrarFilterer is an auto generated log filtering Go binding around an VNT contract events.
type FIFSRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FIFSRegistrarSession is an auto generated Go binding around an VNT contract,
// with pre-set call and transact options.
type FIFSRegistrarSession struct {
	Contract     *FIFSRegistrar    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FIFSRegistrarCallerSession is an auto generated read-only Go binding around an VNT contract,
// with pre-set call options.
type FIFSRegistrarCallerSession struct {
	Contract *FIFSRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// FIFSRegistrarTransactorSession is an auto generated write-only Go binding around an VNT contract,
// with pre-set transact options.
type FIFSRegistrarTransactorSession struct {
	Contract     *FIFSRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// FIFSRegistrarRaw is an auto generated low-level Go binding around an VNT contract.
type FIFSRegistrarRaw struct {
	Contract *FIFSRegistrar // Generic contract binding to access the raw methods on
}

// FIFSRegistrarCallerRaw is an auto generated low-level read-only Go binding around an VNT contract.
type FIFSRegistrarCallerRaw struct {
	Contract *FIFSRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// FIFSRegistrarTransactorRaw is an auto generated low-level write-only Go binding around an VNT contract.
type FIFSRegistrarTransactorRaw struct {
	Contract *FIFSRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFIFSRegistrar creates a new instance of FIFSRegistrar, bound to a specific deployed contract.
func NewFIFSRegistrar(address common.Address, backend bind.ContractBackend) (*FIFSRegistrar, error) {
	contract, err := bindFIFSRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FIFSRegistrar{FIFSRegistrarCaller: FIFSRegistrarCaller{contract: contract}, FIFSRegistrarTransactor: FIFSRegistrarTransactor{contract: contract}, FIFSRegistrarFilterer: FIFSRegistrarFilterer{contract: contract}}, nil
}

// NewFIFSRegistrarCaller creates a new read-only instance of FIFSRegistrar, bound to a specific deployed contract.
func NewFIFSRegistrarCaller(address common.Address, caller bind.ContractCaller) (*FIFSRegistrarCaller, error) {
	contract, err := bindFIFSRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FIFSRegistrarCaller{contract: contract}, nil
}

// NewFIFSRegistrarTransactor creates a new write-only instance of FIFSRegistrar, bound to a specific deployed contract.
func NewFIFSRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*FIFSRegistrarTransactor, error) {
	contract, err := bindFIFSRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FIFSRegistrarTransactor{contract: contract}, nil
}

// NewFIFSRegistrarFilterer creates a new log filterer instance of FIFSRegistrar, bound to a specific deployed contract.
func NewFIFSRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*FIFSRegistrarFilterer, error) {
	contract, err := bindFIFSRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FIFSRegistrarFilterer{contract: contract}, nil
}

// bindFIFSRegistrar binds a generic wrapper to an already deployed contract.
func bindFIFSRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FIFSRegistrarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FIFSRegistrar *FIFSRegistrarRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FIFSRegistrar.Contract.FIFSRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FIFSRegistrar *FIFSRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.FIFSRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FIFSRegistrar *FIFSRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.FIFSRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FIFSRegistrar *FIFSRegistrarCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FIFSRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FIFSRegistrar *FIFSRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FIFSRegistrar *FIFSRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.contract.Transact(opts, method, params...)
}

// Register is a paid mutator transaction binding the contract method 0xd22057a9.
//
// Solidity: function register(subnode bytes32, owner address) returns()
func (_FIFSRegistrar *FIFSRegistrarTransactor) Register(opts *bind.TransactOpts, subnode string, owner common.Address) (*types.Transaction, error) {
	return _FIFSRegistrar.contract.Transact(opts, "registernode", subnode, owner)
}

// Register is a paid mutator transaction binding the contract method 0xd22057a9.
//
// Solidity: function register(subnode bytes32, owner address) returns()
func (_FIFSRegistrar *FIFSRegistrarSession) Register(subnode string, owner common.Address) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.Register(&_FIFSRegistrar.TransactOpts, subnode, owner)
}

// Register is a paid mutator transaction binding the contract method 0xd22057a9.
//
// Solidity: function register(subnode bytes32, owner address) returns()
func (_FIFSRegistrar *FIFSRegistrarTransactorSession) Register(subnode string, owner common.Address) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.Register(&_FIFSRegistrar.TransactOpts, subnode, owner)
}
