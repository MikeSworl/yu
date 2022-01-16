package main

import (
	"github.com/sirupsen/logrus"
	"github.com/yu-org/yu/apps/asset"
	"github.com/yu-org/yu/apps/poa"
	"github.com/yu-org/yu/core/startup"
	"os"
	"strconv"
)

func main() {

	idxStr := os.Args[1]
	idx, err := strconv.Atoi(idxStr)
	if err != nil {
		panic(err)
	}

	myPubkey, myPrivkey, validatorsAddrs := poa.InitKeypair(idx)

	logrus.Info("My Address is ", myPubkey.Address().String())
	startup.StartUp(poa.NewPoa(myPubkey, myPrivkey, validatorsAddrs), asset.NewAsset("YuCoin"))
}
