package Han

import (
	contract "dapp/Connect"

	"math/big"
	// "github.com/ethereum/go-ethereum/common"
	"fmt"

	"github.com/gin-gonic/gin"
)
func Cancel(c *gin.Context){
	ins :=contract.Getsmartcontract()

	times:=c.PostForm("timestamp")
	n := new(big.Int)
    n, ok := n.SetString(times, 10)
    if !ok {
        fmt.Println("SetString: error")
        return
    }
	userlist:=contract.Getuser(ins,n,adress,head)
	fmt.Print(userlist)
	tohtml(c,userlist)
}