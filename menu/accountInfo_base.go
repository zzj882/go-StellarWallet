package menu

import (
	"fmt"
	"github.com/jojopoper/ConsoleColor"
	"github.com/jojopoper/go-StellarWallet/publicdefine"
	"github.com/stellar/go-stellar-base/strkey"
	"strings"
)

const (
	AIB_INFO_INPUT_ADDR = iota
	AIB_INFO_ADDR_FORMAT_ERR
)

type AccountInfoBase struct {
	MenuSubItem
	infoStrings []map[int]string
}

func (this *AccountInfoBase) InitAccInfoBase(parent MenuSubItemInterface, key string) {
	this.MenuSubItem.InitMenu(key)
	this.parentItem = parent
	this.MenuSubItem.Exec = this.execute

	this.MenuSubItem.title = []string{
		publicdefine.L_Chinese: "账户基础信息",
		publicdefine.L_English: "Account basic Informations",
	}

	this.infoStrings = []map[int]string{
		publicdefine.L_Chinese: map[int]string{
			AIB_INFO_INPUT_ADDR:      " 请输入要查询的账户地址，多个账户地址以逗号分隔，不允许有空格\r\n",
			AIB_INFO_ADDR_FORMAT_ERR: "\r\n ** 输入的地址无效 [%s]\r\n",
		},
		publicdefine.L_English: map[int]string{
			AIB_INFO_INPUT_ADDR:      " Please enter the account address you want to query, and the number of the account address is separated by a comma, and space is not allow\r\n",
			AIB_INFO_ADDR_FORMAT_ERR: "\r\n ** Stellar address is invalid [%s]\r\n",
		},
	}

}

func (this *AccountInfoBase) execute(isSync bool) {
	fmt.Println(this.infoStrings[this.languageIndex][AIB_INFO_INPUT_ADDR])

	var input string

	_, err := fmt.Scanf("%s\n", &input)
	if err == nil {
		addrs := strings.Split(input, ",")
		addrs = this.checkAddrs(addrs)
		fmt.Println("\r\n")
		length := len(addrs)
		for i := 0; i < length; i++ {
			ret, err := this.getAccountInfo(addrs[i])
			if err == nil {
				ConsoleColor.Println(ConsoleColor.C_GREEN, " ..........................................................................  ")

				ConsoleColor.Println(ConsoleColor.C_BLUE, ret.ToString(), "\r\n\r\n")
			}
		}
	}
	if !isSync {
		this.ASyncChan <- 0
	}
}

func (this *AccountInfoBase) checkAddrs(addrs []string) []string {
	ret := make([]string, 0)
	for _, itm := range addrs {
		tmp := strings.TrimFunc(itm, func(r rune) bool {
			return r == '\n' || r == '\r' || r == ' '
		})
		if this.verifyAddress(tmp) == nil {
			ret = append(ret, tmp)
		} else {
			ConsoleColor.Printf(ConsoleColor.C_RED, this.infoStrings[this.languageIndex][AIB_INFO_ADDR_FORMAT_ERR],
				itm)
			// fmt.Printf(this.infoStrings[this.languageIndex][AIB_INFO_ADDR_FORMAT_ERR],
			// 	itm)
		}
	}
	return ret
}

func (this *AccountInfoBase) verifyAddress(addr string) error {
	_, err := strkey.Decode(strkey.VersionByteAccountID, addr)
	return err
}

func (this *AccountInfoBase) getAccountInfo(addr string) (ret *publicdefine.StellarAccInfoDef, err error) {
	reqUrl := publicdefine.STELLAR_DEFAULT_NETWORK + publicdefine.STELLAR_NETWORK_ACCOUNTS + "/" + addr
	resMap, err := publicdefine.HttpGet(reqUrl)

	if err == nil {
		ret = &publicdefine.StellarAccInfoDef{}
		ret.PutMapBody(addr, resMap)
	}
	return ret, err
}
