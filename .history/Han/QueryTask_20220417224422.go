package Han
import (
	contract "dapp/Connect"
	// "math/big"
	// "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"fmt"
)
func Query(c *gin.Context){
	ins :=contract.Getsmartcontract()
	Txopts :=contract.GetTxopts()
	head,_:=contract.GetBlockNumber()
	_,adress:=contract.Getaccout()
	contract.GetTasklist(ins,times,adress,head)
}