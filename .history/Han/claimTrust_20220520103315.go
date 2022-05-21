package Han

import (
	contract "dapp/Connect"
	"github.com/gin-gonic/gin"
	"math/big"
	"fmt"
)

func ClaimTrust(c *gin.Context){
	ins := contract.Getsmartcontract()
	Txopts := contract.GetTxopts()
	times := c.PostForm("timestap")
	sign:=c.PostForm("Sign")
	taskname := c.PostForm("taskname")
	n := new(big.Int)
	n, ok := n.SetString(times, 10)
	if !ok {
		fmt.Println("SetString: error")
		return
	}
	contract.ClaimTrust(ins,Txopts,n,sign,taskname)
  
}