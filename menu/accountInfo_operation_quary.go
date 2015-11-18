package menu

import (
	"fmt"
	"github.com/Ledgercn/ConsoleColor"
	"github.com/Ledgercn/go-StellarWallet/publicdefine"
	"strings"
)

const (
	AIOQ_INFO_INPUT_ADDR = iota
	AIOQ_INFO_QUARY_WAITING
	AIOQ_INFO_ADDR_FORMAT_ERR
	AIOQ_INFO_NEXT_RECORDS
)

type AccountInfoOperationQuary struct {
	MenuSubItem
	infoStrings []map[int]string
}

func (this *AccountInfoOperationQuary) InitAccInfoOperQuary(parent MenuSubItemInterface, key string) {
	this.MenuSubItem.InitMenu(key)
	this.parentItem = parent
	this.MenuSubItem.Exec = this.execute

	this.MenuSubItem.title = []string{
		publicdefine.L_Chinese: "账户操作查询",
		publicdefine.L_English: "Account Operations Quary",
	}

	this.infoStrings = []map[int]string{
		publicdefine.L_Chinese: map[int]string{
			AIOQ_INFO_INPUT_ADDR:      " 请输入要查询的账户地址\r\n > ",
			AIOQ_INFO_QUARY_WAITING:   " 正在查询请稍后...",
			AIOQ_INFO_ADDR_FORMAT_ERR: "\r\n ** 输入的地址无效 [%s]\r\n",
			AIOQ_INFO_NEXT_RECORDS:    " > 查看下10条操作，请输入n回车，结束请输入回车: ",
		},
		publicdefine.L_English: map[int]string{
			AIOQ_INFO_INPUT_ADDR:      " Please input the account address you want to query\r\n > ",
			AIOQ_INFO_QUARY_WAITING:   " Searching for a query...",
			AIOQ_INFO_ADDR_FORMAT_ERR: "\r\n ** Stellar address is invalid [%s]\r\n",
			AIOQ_INFO_NEXT_RECORDS:    " > Quary next 10 operations, input n + enter, or press enter to end quary: ",
		},
	}

}

func (this *AccountInfoOperationQuary) execute(isSync bool) {
	addr := this.input_addr()
	if len(addr) > 0 && publicdefine.VerifyGAddress(addr) == nil {
		this.quary(addr)
	} else {
		ConsoleColor.Printf(ConsoleColor.C_RED,
			"\r\n"+this.infoStrings[this.languageIndex][AIOQ_INFO_ADDR_FORMAT_ERR]+"\r\n\r\n", addr)
	}
	if !isSync {
		this.ASyncChan <- 0
	}
}

func (this *AccountInfoOperationQuary) input_addr() string {
	fmt.Printf(this.infoStrings[this.languageIndex][AIOQ_INFO_INPUT_ADDR])

	var input string

	_, err := fmt.Scanf("%s\n", &input)
	if err == nil {
		return strings.Trim(input, " ")
	}
	return ""
}

func (this *AccountInfoOperationQuary) quary(addr string) {
	result := &publicdefine.StellarAccOperationQuary{
		QuaryCursor: "",
		IsEnd:       false,
	}
	for {
		ConsoleColor.Print(ConsoleColor.C_BLUE,
			this.infoStrings[this.languageIndex][AIOQ_INFO_QUARY_WAITING])

		reqUrl := fmt.Sprintf("%s%s/%s/%s?order=desc&limit=%d&cursor=%s", publicdefine.STELLAR_DEFAULT_NETWORK,
			publicdefine.STELLAR_NETWORK_ACCOUNTS, addr,
			publicdefine.STELLAR_NETWORK_OPERATIONS, 10, result.QuaryCursor)
		resMap, err := this.httpget(reqUrl)

		if err == nil {
			result.PutMapBody(resMap)
			this.PrintResult(result)
		} else {
			ConsoleColor.Println(ConsoleColor.C_RED, err)
			break
		}

		if result.IsEnd {
			break
		}
		if this.input_next() != "n" {
			break
		}
	}
}
func (a *AccountInfoOperationQuary) PrintResult(r *publicdefine.StellarAccOperationQuary) {
	fmt.Print("\r")
	for i := 0; i < len(r.Records); i++ {
		ConsoleColor.Printf(ConsoleColor.C_GREEN,
			" %02d --------------------------------------------------------------------------------\r\n", i+1)

		ConsoleColor.Println(ConsoleColor.C_BLUE,
			r.Records[i].ToString(), "\r\n\r\n")
	}
}

func (this *AccountInfoOperationQuary) input_next() string {
	fmt.Printf(this.infoStrings[this.languageIndex][AIOQ_INFO_NEXT_RECORDS])

	var input string

	_, err := fmt.Scanf("%s\n", &input)
	if err == nil {
		return strings.Trim(input, " ")
	}
	return ""
}
