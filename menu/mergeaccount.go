package menu

import (
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/jojopoper/ConsoleColor"
	"github.com/stellar/go-stellar-base"
	"github.com/stellar/go-stellar-base/strkey"
	"go-StellarWallet/publicdefine"
	"net/url"
	"strings"
)

const (
	MA_INFO_INPUT_SOURCE_ADDR = iota
	MA_INFO_INPUT_SOURCE_SEED
	MA_INFO_INPUT_DEST_ADDR
	MA_INFO_CHECK_SOURCE_ACCOUNT
	MA_INFO_SOURCE_ACCOUNT_NOT_EXIST
	MA_INFO_SEED_AND_ADDR_IS_NOT_PAIR
	MA_INFO_CONFIRM_INFOS
	MA_INFO_MERGING
	MA_INFO_MERGING_ERR
	MA_INFO_MERGE_COMLETE
	MA_INFO_MERGING_FIAL
	MA_INFO_ADDR_FORMAT_ERR
	MA_INFO_OPERATION_BREAK
)

type MergeAccount struct {
	MenuSubItem
	infoStrings []map[int]string
}

func (this *MergeAccount) InitMerge(parent MenuSubItemInterface, key string) {
	this.MenuSubItem.InitMenu(key)
	this.parentItem = parent
	this.MenuSubItem.Exec = this.execute

	this.MenuSubItem.title = []string{
		publicdefine.L_Chinese: "账户合并",
		publicdefine.L_English: "Merge account",
	}

	this.infoStrings = []map[int]string{
		publicdefine.L_Chinese: map[int]string{
			MA_INFO_INPUT_SOURCE_ADDR:         " 请输入需要合并的源账户Public地址(G...)  : ",
			MA_INFO_INPUT_SOURCE_SEED:         " 请输入需要合并的源账户Private Seed(S...): ",
			MA_INFO_INPUT_DEST_ADDR:           " 请输入需要合并到目标账户Public地址(G...): ",
			MA_INFO_CHECK_SOURCE_ACCOUNT:      " 正在检查账户有效性....",
			MA_INFO_SOURCE_ACCOUNT_NOT_EXIST:  " ** 输入的账户[%s]不存在，请确认！",
			MA_INFO_SEED_AND_ADDR_IS_NOT_PAIR: " ** 输入的Private Seed与Public地址不匹配！",
			MA_INFO_CONFIRM_INFOS: "\r\n ~~~~ 请确认如下信息 ~~~~\r\n\r\n " +
				"需要将{源账户}[ %s ]合并到{目标账户}[ %s ]，合并成功后，" +
				"{源账户}内的所有余额在扣除交易手续费之后，全部充值到{目标账户}，" +
				"{源账户}将不可使用！\r\n\r\n 确认以上操作请输入 yes ，否则按任意键退出操作: ",
			MA_INFO_MERGING:         " 发送账户合并请求...",
			MA_INFO_MERGING_ERR:     " 发送账户合并请求失败!",
			MA_INFO_MERGE_COMLETE:   " 账户合并请求完成",
			MA_INFO_MERGING_FIAL:    " 账户合并失败!",
			MA_INFO_OPERATION_BREAK: "\r\n === 合并操作已经被终止 ===\r\n",
			MA_INFO_ADDR_FORMAT_ERR: "\r\n ** 你的输入无效\r\n",
		},
		publicdefine.L_English: map[int]string{
			MA_INFO_INPUT_SOURCE_ADDR:         " Please enter the Public address for the merged Source Account(G...):\r\n",
			MA_INFO_INPUT_SOURCE_SEED:         " Please enter the Private Seed for the merged Source Account(S...)  :\r\n",
			MA_INFO_INPUT_DEST_ADDR:           " Please enter the Public address for need to merge the Target Account(G...):\r\n",
			MA_INFO_CHECK_SOURCE_ACCOUNT:      " Checking Aource Account....",
			MA_INFO_SOURCE_ACCOUNT_NOT_EXIST:  " ** Source account[%s] is not exist!",
			MA_INFO_SEED_AND_ADDR_IS_NOT_PAIR: " ** Private Seed and Public address does not match!",
			MA_INFO_CONFIRM_INFOS: "\r\n ~~~~ Please confirm the following information ~~~~\r\n\r\n" +
				" Will require {Source Accounts}[% s] the merger to {Target Account}[% s], " +
				"after the success of the merger. {Source Account} all balance after deducting transaction fees, " +
				"all the recharge to a {Target Account}. {Source Account} will not can use!\r\n\r\n " +
				"Confirm the above operation, please enter yes, or press any key to exit the operation: ",
			MA_INFO_MERGING:         " Send account merge request...",
			MA_INFO_MERGING_ERR:     " ** Send account merge request FAIL!",
			MA_INFO_MERGE_COMLETE:   " Account merge request complete",
			MA_INFO_MERGING_FIAL:    " Account merge FAIL!",
			MA_INFO_OPERATION_BREAK: "\r\n === Merge operation is break ===\r\n",
			MA_INFO_ADDR_FORMAT_ERR: "\r\n ** Your input is invalid\r\n",
		},
	}
}

func (this *MergeAccount) execute(isSync bool) {
	fmt.Println("")
	var isError bool = true
	var srcAddr, srcSeed, destAddr string
	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			srcAddr = this.input_SrcAddr()
			isError = len(srcAddr) == 0
		case 1:
			srcSeed = this.input_SrcSeed()
			isError = len(srcSeed) == 0
		case 2:
			destAddr = this.input_DestAddr()
			isError = len(destAddr) == 0
		case 3:
			ConsoleColor.Printf(ConsoleColor.C_BLUE,
				this.infoStrings[this.languageIndex][MA_INFO_CONFIRM_INFOS], srcAddr, destAddr)
			// fmt.Printf(this.infoStrings[this.languageIndex][MA_INFO_CONFIRM_INFOS], srcAddr, destAddr)
			confirm := this.input_Confirm()
			if confirm != "yes" {
				i = 5
				ConsoleColor.Println(ConsoleColor.C_YELLOW,
					this.infoStrings[this.languageIndex][MA_INFO_OPERATION_BREAK])
				// fmt.Println(this.infoStrings[this.languageIndex][MA_INFO_OPERATION_BREAK])
			}
		case 4:
			this.beginMerge(srcAddr, srcSeed, destAddr)
		}

		if isError {
			ConsoleColor.Println(ConsoleColor.C_RED,
				this.infoStrings[this.languageIndex][MA_INFO_ADDR_FORMAT_ERR])
			// fmt.Println(this.infoStrings[this.languageIndex][MA_INFO_ADDR_FORMAT_ERR])
			break
		}
	}

	if !isSync {
		this.ASyncChan <- 0
	}
}

func (this *MergeAccount) input_SrcAddr() string {
	fmt.Printf(this.infoStrings[this.languageIndex][MA_INFO_INPUT_SOURCE_ADDR])

	var input string

	_, err := fmt.Scanf("%s\n", &input)
	if err == nil {
		if this.verifyGAddress(input) == nil {
			return input
		}
	}
	return ""
}

func (this *MergeAccount) input_SrcSeed() string {
	fmt.Printf(this.infoStrings[this.languageIndex][MA_INFO_INPUT_SOURCE_SEED])

	input := gopass.GetPasswdMasked()
	if this.verifySAddress(string(input)) == nil {
		return string(input)
	}
	return ""
}

func (this *MergeAccount) input_DestAddr() string {
	fmt.Printf(this.infoStrings[this.languageIndex][MA_INFO_INPUT_DEST_ADDR])

	var input string

	_, err := fmt.Scanf("%s\n", &input)
	if err == nil {
		if this.verifyGAddress(input) == nil {
			return input
		}
	}
	return ""
}

func (this *MergeAccount) input_Confirm() string {
	var input string

	_, err := fmt.Scanf("%s\n", &input)
	if err == nil {
		return strings.ToLower(strings.Trim(input, " "))
	}
	return ""
}

func (this *MergeAccount) beginMerge(srcAddr, srcSeed, destAddr string) {
	ConsoleColor.Println(ConsoleColor.C_BLUE,
		this.infoStrings[this.languageIndex][MA_INFO_CHECK_SOURCE_ACCOUNT])
	// fmt.Println(this.infoStrings[this.languageIndex][MA_INFO_CHECK_SOURCE_ACCOUNT])
	// 先检查账户是不是存在
	srcInfo := this.checkSourceAddr(srcAddr)
	if srcInfo == nil {
		return
	}
	// 检查Seed和Public Addr是不是匹配
	if this.checkSeed(srcSeed, srcAddr) == false {
		return
	}

	ConsoleColor.Println(ConsoleColor.C_BLUE,
		this.infoStrings[this.languageIndex][MA_INFO_MERGING])
	// fmt.Println(this.infoStrings[this.languageIndex][MA_INFO_MERGING])

	mergeInfo := this.merging(srcInfo, srcSeed, destAddr)
	if mergeInfo == nil {
		ConsoleColor.Println(ConsoleColor.C_RED,
			this.infoStrings[this.languageIndex][MA_INFO_MERGING_ERR])
		// fmt.Println(this.infoStrings[this.languageIndex][MA_INFO_MERGING_ERR])
		return
	}

	// 检查transaction hash是否生效
	if len(mergeInfo.ResultHash) == 0 {
		ConsoleColor.Println(ConsoleColor.C_RED,
			this.infoStrings[this.languageIndex][MA_INFO_MERGING_FIAL])
		// fmt.Println(this.infoStrings[this.languageIndex][MA_INFO_MERGING_FIAL])
		return
	}

	ConsoleColor.Println(ConsoleColor.C_GREEN,
		this.infoStrings[this.languageIndex][MA_INFO_MERGE_COMLETE])
	// fmt.Println(this.infoStrings[this.languageIndex][MA_INFO_MERGE_COMLETE])

}

func (this *MergeAccount) checkSourceAddr(addr string) *publicdefine.StellarAccInfoDef {
	reqUrl := publicdefine.STELLAR_DEFAULT_NETWORK + publicdefine.STELLAR_NETWORK_ACCOUNTS + "/" + addr
	resMap, err := publicdefine.HttpGet(reqUrl)

	if err == nil {
		ret := &publicdefine.StellarAccInfoDef{}
		ret.PutMapBody(addr, resMap)
		if ret.IsExist() {
			return ret
		} else {
			ConsoleColor.Printf(ConsoleColor.C_RED,
				this.infoStrings[this.languageIndex][MA_INFO_SOURCE_ACCOUNT_NOT_EXIST]+"\r\n", addr)
			// fmt.Printf(this.infoStrings[this.languageIndex][MA_INFO_SOURCE_ACCOUNT_NOT_EXIST]+"\r\n", addr)
		}
	} else {
		ConsoleColor.Println(ConsoleColor.C_RED, err)
		// fmt.Println(err)
	}
	return nil
}

func (this *MergeAccount) checkSeed(seed, srcAddr string) bool {
	pubadd, _, err := stellarbase.GenerateKeyFromSeed(seed)
	if err == nil {
		if pubadd.Address() == srcAddr {
			return true
		}
		ConsoleColor.Printf(ConsoleColor.C_RED,
			this.infoStrings[this.languageIndex][MA_INFO_SEED_AND_ADDR_IS_NOT_PAIR]+"\r\n")
		// fmt.Printf(this.infoStrings[this.languageIndex][MA_INFO_SEED_AND_ADDR_IS_NOT_PAIR] + "\r\n")
	} else {
		ConsoleColor.Println(ConsoleColor.C_RED, err)
		// fmt.Println(err)
	}
	return false
}

func (this *MergeAccount) merging(srcInfo *publicdefine.StellarAccInfoDef, srcSeed, destAddr string) *publicdefine.StellarAccountMerge {
	mergeInfo := &publicdefine.StellarAccountMerge{
		SrcInfo:        srcInfo,
		DestPublicAddr: destAddr,
	}

	signed := mergeInfo.GetSigned(srcSeed)

	if len(signed) > 0 {
		data := "tx=" + url.QueryEscape(signed)
		// data, err := json.Marshal(map[string]interface{}{
		// 	"tx": signed,
		// })

		postUrl := publicdefine.STELLAR_DEFAULT_NETWORK + publicdefine.STELLAR_NETWORK_TRANSACTIONS
		ret, err := publicdefine.HttpPost_form(postUrl, data)
		if err == nil {
			mergeInfo.PutResult(ret)
			return mergeInfo
		}
		ConsoleColor.Println(ConsoleColor.C_RED, err)
		// fmt.Println(err)
	}
	return nil
}

func (this *MergeAccount) verifyGAddress(addr string) error {
	_, err := strkey.Decode(strkey.VersionByteAccountID, addr)
	return err
}

func (this *MergeAccount) verifySAddress(addr string) error {
	_, err := strkey.Decode(strkey.VersionByteSeed, addr)
	return err
}
