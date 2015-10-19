package menu

import (
	"fmt"
	"github.com/jojopoper/go-StellarWallet/publicdefine"
)

const (
	SA_INFO_MEMO = iota
)

type SoftwareAbout struct {
	MenuSubItem
	infoStrings []map[int]string
}

func (this *SoftwareAbout) InitAbout(parent MenuSubItemInterface, key string) {
	this.MenuSubItem.InitMenu(key)
	this.parentItem = parent
	this.MenuSubItem.Exec = this.execute

	this.MenuSubItem.title = []string{
		publicdefine.L_Chinese: "关于",
		publicdefine.L_English: "About",
	}

	this.infoStrings = []map[int]string{
		publicdefine.L_Chinese: map[int]string{
			SA_INFO_MEMO: "   软件版本 : 1.0.0.20151018\r\n" +
				"   开发团队 : http://www.ledgercn.com\r\n" +
				"   钱包源码 : https://www.github.com/jojopoper/go-StellarWallet\r\n" +
				" 我们的QQ群 : 204444685\r\n" +
				" 支持和打赏 : gBgfNY9Mzpx1Fbz9btGuY4RLxSmkE3PeYC\r\n",
		},
		publicdefine.L_English: map[int]string{
			SA_INFO_MEMO: "     Wallet Version : 1.0.0.20151018\r\n" +
				"           Our team : http://www.ledgercn.com\r\n" +
				"        Source code : https://www.github.com/jojopoper/go-StellarWallet\r\n" +
				"       Our QQ group : 204444685\r\n" +
				" Support and reward : gBgfNY9Mzpx1Fbz9btGuY4RLxSmkE3PeYC\r\n",
		},
	}

}

func (this *SoftwareAbout) execute(isSync bool) {
	fmt.Println("")
	fmt.Println(this.infoStrings[this.languageIndex][SA_INFO_MEMO])
	fmt.Println("")

	var input string

	fmt.Scanf("%s\n", &input)

	if !isSync {
		this.ASyncChan <- 0
	}
}
