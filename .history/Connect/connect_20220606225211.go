package Connect

import (
	"context"
	"crypto/ecdsa"
	contract2 "dapp/Smart2"
	contract "dapp/Smartgo"

	// "encoding"
	"bytes"
	"fmt"
	"io/ioutil"
	"math/big"

	// "unsafe"

	// "strconv"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	// "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	// "log"
	"os"

	"github.com/lithammer/fuzzysearch/fuzzy"
	// "fmt"
)

var (
	//本地geth地址
	adress = "http://localhost:8545"
	//本地账户地址
	privatekeyfile = ""
	//本地账户密码
	password = ""
	//合约地址
	contractadress  = "0x856E8800c0aF1bc786d4FD40efaC4f010FFC5E8f"
	contractadress2 = "0x078d51D7DE304A23c5FE169442B6486251b677F5"
	//读取用户keystore文件地址
	relativePath = "D:\\y\\geth\\node1\\nodedata\\keystore\\"
	//本地链chainID交易:修改为本地的chainID
	chainID = 9876678900
)

var client *ethclient.Client
var rDel *rpc.Client

//连接geth
func init() {
	rpcDel, err := rpc.Dial(adress)
	if err != nil {
		fmt.Println("连接geth====>", err)
		panic(err)
	} else {
		fmt.Println("geth连接成功*****============================***********")
	}
	rDel = rpcDel
	client = ethclient.NewClient(rpcDel)

	//fmt.Println(client)
}

/*
实例化合约
*/
func Getsmartcontract() *contract.TaskDeployerContract {
	ins, err := contract.NewTaskDeployerContract(common.HexToAddress(contractadress), client)
	if err != nil {
		panic(err)
	}
	return ins
}
func Getsmartcontract2() *contract2.Faucet {
	ins2, err := contract2.NewFaucet(common.HexToAddress(contractadress2), client)
	if err != nil {
		panic(err)
	}
	return ins2
}
func Getaccout() (*ecdsa.PrivateKey, common.Address) {
	file := privatekeyfile
	account, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	pwd := password
	key, err := keystore.DecryptKey(account, pwd)
	if err != nil {
		panic(err)
	}
	//fmt.Println(key.PrivateKey, key.Address)
	return key.PrivateKey, key.Address
}

//获取gasprice
func GetgasPrice() (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return big.NewInt(0), err
	} else {
		return gasPrice, nil
	}

}

//获取nonce
func Getnonce(address common.Address) (uint64, error) {
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return 0, err
	} else {
		return nonce, nil
	}
}

//获取区块数
func GetBlockNumber() (*types.Header, error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	//fmt.Println(header)
	return header, err
}

//获取区块的详细信息
func Getblockmessage(headr int64) (*big.Int, uint64) {
	blockNumber := big.NewInt(headr)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		panic(err)
	}
	blockNow := block.Number()
	timestamp := block.Time()
	fmt.Println(block.Number().Uint64())
	fmt.Println(block.Time())
	return blockNow, timestamp
}

//设置TransactOpts
func setopts(privateKey *ecdsa.PrivateKey, address common.Address) *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(chainID)))
	if err != nil {
		panic(err)
	}
	nonce, err := Getnonce(address)
	if err != nil {
		panic(err)
	}
	gasPrice, err := GetgasPrice()
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	return auth
}
func GetTxopts() *bind.TransactOpts {
	privateKey, publicKey := Getaccout()
	opts := setopts(privateKey, publicKey)
	return opts
}
func Getaccout2() (*ecdsa.PrivateKey, common.Address) {
	file := "D:\\y\\geth\\node1\\nodedata\\keystore\\UTC--2021-09-12T17-06-06.881126000Z--00dc6e8b60fa02a5d83e525bbef3240e8ea54dc5"
	account, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	pwd := "1111"
	key, err := keystore.DecryptKey(account, pwd)
	if err != nil {
		panic(err)
	}
	//fmt.Println(key.PrivateKey, key.Address)
	return key.PrivateKey, key.Address
}
func GerTxopts2() *bind.TransactOpts{
	privateKey, publicKey := Getaccout2()
	opts := setopts(privateKey, publicKey)
	return opts
}

/*
获取合约余额
*/
func GetcontractBanlance(ins *contract.TaskDeployerContract, address common.Address, header *types.Header) *big.Int {
	opts := bind.CallOpts{
		Pending:     true,
		From:        address,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
	balance, err := ins.GetBalanceOfContract(&opts)
	if err != nil {
		panic(err)
	}
	return balance
}

/*
获取任务列表函数
*/
func GetTasklist(ins *contract.TaskDeployerContract, timestap *big.Int, address common.Address, header *types.Header) struct {
	Taskname    string
	Sponsor     common.Address
	Beneficiary common.Address
	Category    string
	Amount      *big.Int
	Timestamp   *big.Int
	State       string
	LaunchTime  string
} {
	opts := bind.CallOpts{
		Pending:     true,
		From:        address,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
	Tasklist, err := ins.Tasklist(&opts, timestap)
	if err != nil {
		panic(err)
	}
	return Tasklist
}

/*
获取用户信息列表
*/
func Getuser(ins *contract.TaskDeployerContract, timestap *big.Int, address common.Address, header *types.Header) struct {
	Name        string
	Phonenumber string
	Studentid   *big.Int
	Password    string
} {
	opts := bind.CallOpts{
		Pending:     true,
		From:        address,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
	UserList, err := ins.Userlist(&opts, timestap)
	if err != nil {
		panic(err)
	}
	return UserList

}

/*
创建任务函数
*/
func CreatNewEvent(
	ins *contract.TaskDeployerContract,
	ops *bind.TransactOpts,
	Taskname string,
	Taskcatagory string,
	launchTime string,
	amount *big.Int,
) *types.Transaction {
	c := amount.String() + "000000000000000000"
	n := new(big.Int)
	n, ok := n.SetString(c, 10)
	if !ok {
		fmt.Println("SetString: error")
	}
	ops.Value = n
	a, err := ins.CreateNewEvent(ops, launchTime, Taskcatagory, Taskname, amount)
	if err != nil {
		fmt.Println("CreatNewEvent error ===>", err)
		panic(err)
	}
	return a
	// a:=receipt.Status
}
func QueryStatus(txHash common.Hash) int {
	rec, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		// panic(err)
		return (-1)
	}

	return (int(rec.Status))
}

/*
获取任务时间戳
*/
func Querytime(ins *contract.TaskDeployerContract,
	address common.Address,
	header *types.Header) *big.Int {
	opts := bind.CallOpts{
		Pending:     true,
		From:        address,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
	timestamp, err := ins.Query(&opts)
	if err != nil {
		panic(err)
	}
	return timestamp
}

//获取用户余额
func GetuserBanlance(ins *contract.TaskDeployerContract,
	address common.Address,
	header *types.Header) *big.Int {
	opts := bind.CallOpts{
		Pending:     true,
		From:        address,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
	balance, err := ins.GetBalanceOfUser(&opts)
	if err != nil {
		panic(err)
	}
	return balance
}

/*
取消任务函数
*/

func CancelEvent(
	ins *contract.TaskDeployerContract,
	ops *bind.TransactOpts,
	timestamp *big.Int,
	address common.Address) {
	ops.From = address
	_, err := ins.CancelEvent(ops, timestamp)
	if err != nil {
		panic(err)
	}
}

/*
接受任务函数
*/
func Confirmtask(
	ins *contract.TaskDeployerContract,
	timestamp *big.Int,
	ops *bind.TransactOpts,
	address common.Address) {
	ops.From = address
	_, err := ins.Confirmtask(ops, timestamp)
	if err != nil {
		panic(err)
	}
}

/*
发布任务者确认接受任务者完成任务
*/

func ClaimTrust(
	ins *contract.TaskDeployerContract,
	ops *bind.TransactOpts,
	timestamp *big.Int,
) int {
	_, err := ins.ClaimTrust(ops, timestamp)
	if err != nil {
		panic(err)
	}
	return 1
}

//获取交易的hash值
func Gettaskhash(ins *contract.TaskDeployerContract, address common.Address, header *types.Header, taskname string, timestamp string) [32]byte {
	opts := bind.CallOpts{
		Pending:     true,
		From:        address,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
	hash, err := ins.GetTxHash(&opts, taskname, timestamp)
	if err != nil {
		panic(hash)
	}
	fmt.Println("hash==>", hexutil.Encode(hash[:]))
	return hash
}

//获取发布任务的用户对当前任务的确认签名
func GetthistaskSign(PrivateKey *ecdsa.PrivateKey, hash [32]byte) []byte {
	signature, _ := crypto.Sign(hash[:], PrivateKey)
	fmt.Println("signature", hexutil.Encode(signature))
	return signature
}

/*
创建用户信息
*/
func Creatuser(
	ins *contract.TaskDeployerContract,
	opts *bind.TransactOpts,
	name string,
	phonenumber string,
	studentid *big.Int,
	password string) bool {
	_, err := ins.Createuser(opts, name, phonenumber, studentid, password)
	if err != nil {
		panic(err)
	}
	return true
}

/*
申请代币
*/
func Dotrasfer(ins2 *contract2.Faucet, opts *bind.TransactOpts, address common.Address) *types.Transaction {
	Trasfermessage, err := ins2.Dotransfer(opts, address)
	if err != nil {
		panic(err)
	}
	return Trasfermessage
}

//申请代币次数
func Getfrequency(ins2 *contract2.Faucet, address common.Address, header *types.Header) *big.Int {
	opts := bind.CallOpts{
		Pending:     true,
		From:        address,
		BlockNumber: header.Number,
		Context:     context.Background(),
	}
	f, err := ins2.Gettransfercount(&opts, address)
	if err != nil {
		panic(err)
	}
	return f
}

/*
用户修改信息
*/
//修改密码
func Userchangepsword(
	ins *contract.TaskDeployerContract,
	opts *bind.TransactOpts,
	password string,
	studentid *big.Int) bool {
	_, err := ins.Changepassword(opts, password, studentid)
	if err != nil {
		panic(err)
	}
	return true
}

//修改名称
func Userchangename(ins *contract.TaskDeployerContract,
	opts *bind.TransactOpts,
	name string,
	studentid *big.Int,
	password string) bool {
	_, err := ins.Changename(opts, name, studentid, password)
	if err != nil {
		panic(err)
	}
	return true
}

var newAccount string
var accounts []string

func CreatnewActogeth(pd string) string {
	// fmt.Print("why----->", rDel)
	err := rDel.Call(&newAccount, "personal_newAccount", pd)
	if err != nil {
		panic(err)
	}
	rDel.Call(&accounts, "personal_listAccounts")
	// if err!=nil {
	// 	panic(err)
	// }
	Transfer(accounts[0])
	fmt.Print(accounts)
	return newAccount
}

func Transfer(account string){
	_,b:=Getaccout()
    privateKey, err := crypto.HexToECDSA("00dc6e8b60fa02a5d83e525bbef3240e8ea54dc5")
    if err != nil {
        panic(err)
    }
	// privateKey,publicKey:=Getaccout2()
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
		panic("error casting public key to ECDSA")
    }
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        panic(err)
    }

    value := big.NewInt(1000000000000000000) // in wei (1 eth)
    gasLimit := uint64(21000)                // in units
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
		panic(err)
    }

    toAddress := common.
    var data []byte
    tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

    chainid, err := client.NetworkID(context.Background())
    if err != nil {
		panic(err)
    }

    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainid), privateKey)
    if err != nil {
        panic(err)
    }

    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        panic(err)
    }

    fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
//针对不同用户登入获取不同用户的信息来对交易签名
func Changeuser(ad string, pw string) {
	var FileInfo []os.FileInfo
	var err error

	if FileInfo, err = ioutil.ReadDir(relativePath); err != nil {
		fmt.Println("读取 keystore 文件夹出错")
		return
	}
	a := make([]string, 0)
	for _, fileInfo := range FileInfo {
		a = append(a, fileInfo.Name())
	}
	// ac:="5c595872e02b0613658036bdf5daa6d9f42954be"
	fmt.Print("keystore", a)
	matches2 := fuzzy.Find(ad[2:], a)
	fmt.Println("当前登入的用户为", ad)
	// print(relativePath+"//"+matches1[0])
	fmt.Println(".....", matches2)
	privatekeyfile = relativePath + "\\" + matches2[0]
	password = pw
	
	fmt.Println("privatekeyfile and pw:", privatekeyfile, pw)
	// matches2 = nil
}

//用户退出状态 文件处于空状态
func Userexit() {
	privatekeyfile = ""
	password = ""
}
func Get() (string, string) {
	return password, privatekeyfile
}

//注销用户
func Cancellation(ad string) string {
	var FileInfo []os.FileInfo
	var err error
	relativePath := "E://Block_chain//data//keystore"

	if FileInfo, err = ioutil.ReadDir(relativePath); err != nil {
		fmt.Println("读取 keystore 文件夹出错")
		return err.Error()
	}
	a := make([]string, 0)
	for _, fileInfo := range FileInfo {
		a = append(a, fileInfo.Name())
	}
	matches2 := fuzzy.Find(ad[2:], a)
	adfile := relativePath + "//" + matches2[0]
	err2 := os.Remove(adfile)
	if err2 != nil {
		panic(err2)
	} else {
		fmt.Println("用户注销完毕")
		matches2 = nil
		result := "成功注销"
		return result
	}

}

func Validation(hash []byte, signature []byte, publicKeyBytes []byte) bool {
	sigPublicKey, err := crypto.Ecrecover(hash, signature)
	if err != nil {
		panic(err)
	}
	fmt.Println("sigPublick", sigPublicKey)
	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches) // true
	return (matches)
}
