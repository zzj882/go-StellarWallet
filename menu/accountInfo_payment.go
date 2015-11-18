package menu

import (
	// "encoding/json"
	"fmt"
	"github.com/Ledgercn/ConsoleColor"
	"github.com/Ledgercn/go-StellarWallet/publicdefine"
	"github.com/howeyc/gopass"
	"github.com/stellar/go-stellar-base/keypair"
	"net/url"
	"strconv"
	"strings"
)

const (
	AIP_INFO_INPUT_PUBLIC_ADDR = iota
	AIP_INFO_INPUT_PRIVATE_SEED
	AIP_INFO_INPUT_DESTINATION
	AIP_INFO_INPUT_AMOUNT
	AIP_INFO_CHECK_SOURCE_ACCOUNT
	AIP_INFO_SOURCE_ACCOUNT_NOT_EXIST
	AIP_INFO_SEED_AND_ADDR_IS_NOT_PAIR
	AIP_INFO_CREDIT_IS_LOW
	AIP_INFO_DEST_ADDR_NOT_EXIST
	AIP_INFO_PAYMENT_ABORT
	AIP_INFO_PAYMENT_ABORT_LESS20
	AIP_INFO_CREATE_DEST_ADDR
	AIP_INFO_CREATE_DEST_ADDR_SUCCESS
	AIP_INFO_SENDING
	AIP_INFO_SEND_ERROR
	AIP_INFO_SEND_COMPLETE
	AIP_INFO_CHECK_TRANSACTION
	AIP_INFO_CHECK_TRANSACTION_ERROR
	AIP_INFO_CHECK_TRANSACTION_SUCCESS
	AIP_INFO_ADDR_FORMAT_ERR
)

type AccountInfoPayment struct {
	MenuSubItem
	infoStrings []map[int]string
}

func (this *AccountInfoPayment) InitAccInfoPayment(parent MenuSubItemInterface, key string) {
	this.MenuSubItem.InitMenu(key)
	this.parentItem = parent
	this.MenuSubItem.Exec = this.execute

	this.MenuSubItem.title = []string{
		publicdefine.L_Chinese: "支付",
		publicdefine.L_English: "Payment",
	}

	this.infoStrings = []map[int]string{
		publicdefine.L_Chinese: map[int]string{
			AIP_INFO_INPUT_PUBLIC_ADDR:         " 请输入源账号Public地址(G...)\t: ",
			AIP_INFO_INPUT_PRIVATE_SEED:        " 请输入源账号Private Seed(S...)\t: ",
			AIP_INFO_INPUT_DESTINATION:         " 请输入接收账户Public地址(G...)\t: ",
			AIP_INFO_INPUT_AMOUNT:              " 请输入发送金额: ",
			AIP_INFO_CHECK_SOURCE_ACCOUNT:      " 正在检查账户有效性....",
			AIP_INFO_SOURCE_ACCOUNT_NOT_EXIST:  " ** 输入的账户[%s]不存在，请确认!",
			AIP_INFO_SEED_AND_ADDR_IS_NOT_PAIR: " ** 输入的Private Seed与Public地址不匹配！",
			AIP_INFO_CREDIT_IS_LOW:             " ** 账户[%s]余额不足，余额为[%s]",
			AIP_INFO_DEST_ADDR_NOT_EXIST:       " 目标账户[ %s ]不存在，需要创建请输入 yes，否则按回车结束操作: \r\n>",
			AIP_INFO_PAYMENT_ABORT:             " ** 支付流程终止！",
			AIP_INFO_PAYMENT_ABORT_LESS20:      " ** 支付流程终止！(新建账户最少需要 20 Lumens)",
			AIP_INFO_CREATE_DEST_ADDR:          " 正在创建账户[ %s ]....\r\n",
			AIP_INFO_CREATE_DEST_ADDR_SUCCESS:  " 创建账户成功!",
			AIP_INFO_SENDING:                   " 正在发送....",
			AIP_INFO_SEND_ERROR:                " 发送过程中发生错误! ",
			AIP_INFO_SEND_COMPLETE:             " 发送完成，检查发送结果....",
			AIP_INFO_CHECK_TRANSACTION:         " 检查发送的有效性....",
			AIP_INFO_CHECK_TRANSACTION_ERROR:   " 发送失败! ",
			AIP_INFO_CHECK_TRANSACTION_SUCCESS: " 发送成功! ",
			AIP_INFO_ADDR_FORMAT_ERR:           "\r\n ** 你的输入无效\r\n",
		},
		publicdefine.L_English: map[int]string{
			AIP_INFO_INPUT_PUBLIC_ADDR:         " Please input source Public address(G...)     : ",
			AIP_INFO_INPUT_PRIVATE_SEED:        " Please input source Private Seed(S...)       : ",
			AIP_INFO_INPUT_DESTINATION:         " Please input destination Public address(G...): ",
			AIP_INFO_INPUT_AMOUNT:              " Please input amount : ",
			AIP_INFO_CHECK_SOURCE_ACCOUNT:      " Checking source account ....",
			AIP_INFO_SOURCE_ACCOUNT_NOT_EXIST:  " ** Source account[%s] is not exist! ",
			AIP_INFO_SEED_AND_ADDR_IS_NOT_PAIR: " ** Private Seed and Public address does not match!",
			AIP_INFO_CREDIT_IS_LOW:             " ** Account[%s] credit is low，Balance = [%s]",
			AIP_INFO_DEST_ADDR_NOT_EXIST:       " Destation address [ %s ] is not Exist，if you need to create this account, input yes + enter, otherwise press the Enter to terminate this payment : \r\n>",
			AIP_INFO_PAYMENT_ABORT:             " ** Payment process is terminated！",
			AIP_INFO_PAYMENT_ABORT_LESS20:      " ** Payment process is terminated！(Create new account at least 20 Lumens)",
			AIP_INFO_CREATE_DEST_ADDR:          " Creating account [ %s ]....\r\n",
			AIP_INFO_CREATE_DEST_ADDR_SUCCESS:  " Create account success!",
			AIP_INFO_SENDING:                   " sending ....",
			AIP_INFO_SEND_ERROR:                " Send error! ",
			AIP_INFO_SEND_COMPLETE:             " Send complete, check send results....",
			AIP_INFO_CHECK_TRANSACTION:         " Check the validity of the transmission....",
			AIP_INFO_CHECK_TRANSACTION_ERROR:   " Transaction error! ",
			AIP_INFO_CHECK_TRANSACTION_SUCCESS: " Send success! ",
			AIP_INFO_ADDR_FORMAT_ERR:           "\r\n ** Your input is invalid\r\n",
		},
	}

}

func (this *AccountInfoPayment) execute(isSync bool) {
	fmt.Println("")
	var isError bool = true
	var srcAddr, srcSeed, destAddr string
	var amount float64

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
			amount = this.input_Amount()
			isError = amount == 0
		case 4:
			this.beginSend(srcAddr, srcSeed, destAddr, amount)
		}

		if isError {
			ConsoleColor.Println(ConsoleColor.C_RED,
				this.infoStrings[this.languageIndex][AIP_INFO_ADDR_FORMAT_ERR])
			// fmt.Println(this.infoStrings[this.languageIndex][AIP_INFO_ADDR_FORMAT_ERR])
			break
		}
	}

	if !isSync {
		this.ASyncChan <- 0
	}
}

func (this *AccountInfoPayment) input_SrcAddr() string {
	fmt.Printf(this.infoStrings[this.languageIndex][AIP_INFO_INPUT_PUBLIC_ADDR])

	var input string

	_, err := fmt.Scanf("%s\n", &input)
	if err == nil {
		if publicdefine.VerifyGAddress(input) == nil {
			return input
		}
	}
	return ""
}

func (this *AccountInfoPayment) input_SrcSeed() string {
	fmt.Printf(this.infoStrings[this.languageIndex][AIP_INFO_INPUT_PRIVATE_SEED])

	input := gopass.GetPasswdMasked()
	if publicdefine.VerifySAddress(string(input)) == nil {
		return string(input)
	}
	return ""
}

func (this *AccountInfoPayment) input_DestAddr() string {
	fmt.Printf(this.infoStrings[this.languageIndex][AIP_INFO_INPUT_DESTINATION])

	var input string

	_, err := fmt.Scanf("%s\n", &input)
	if err == nil {
		if publicdefine.VerifyGAddress(input) == nil {
			return input
		}
	}
	return ""
}

func (this *AccountInfoPayment) input_Amount() float64 {
	fmt.Printf(this.infoStrings[this.languageIndex][AIP_INFO_INPUT_AMOUNT])

	var input string

	_, err := fmt.Scanf("%s\n", &input)
	if err == nil {
		amount, b := publicdefine.IsFLOAT64(input)
		if b && amount > 0 {
			return amount
		}
	}
	return 0
}

func (this *AccountInfoPayment) beginSend(srcAddr, srcSeed, destAddr string, amount float64) {
	ConsoleColor.Println(ConsoleColor.C_BLUE,
		this.infoStrings[this.languageIndex][AIP_INFO_CHECK_SOURCE_ACCOUNT])
	// fmt.Println(this.infoStrings[this.languageIndex][AIP_INFO_CHECK_SOURCE_ACCOUNT])
	// 先检查账户是不是存在并且余额是不是够
	ret := this.checkSourceAddr(srcAddr, amount)
	if ret == nil {
		return
	}
	// 检查Seed和Public Addr是不是匹配
	if this.checkSeed(srcSeed, srcAddr) == false {
		return
	}

	// 检查目标账户是否存在，不存在需要先建立账户
	if this.checkPublicAddrExist(destAddr) == nil {
		// 如果目标账户不存在，需要按照建立账户的方式进行
		if this.create_account(ret, srcSeed, destAddr, amount) == false {
			return
		}
	} else {
		// 如果已经存在，就按照正常流程付款
		ConsoleColor.Println(ConsoleColor.C_BLUE,
			this.infoStrings[this.languageIndex][AIP_INFO_SENDING])
		// fmt.Println(this.infoStrings[this.languageIndex][AIP_INFO_SENDING])

		// 签名支付
		payment := this.pay(ret, srcSeed, destAddr, amount)
		if payment == nil {
			ConsoleColor.Println(ConsoleColor.C_RED,
				this.infoStrings[this.languageIndex][AIP_INFO_SEND_ERROR])
			// fmt.Println(this.infoStrings[this.languageIndex][AIP_INFO_SEND_ERROR])
			return
		}

		ConsoleColor.Println(ConsoleColor.C_BLUE,
			this.infoStrings[this.languageIndex][AIP_INFO_SEND_COMPLETE])
		// fmt.Println(this.infoStrings[this.languageIndex][AIP_INFO_SEND_COMPLETE])

		// 检查transaction hash是否生效
		if len(payment.ResultHash) == 0 {
			ConsoleColor.Println(ConsoleColor.C_RED,
				this.infoStrings[this.languageIndex][AIP_INFO_CHECK_TRANSACTION_ERROR])
			// fmt.Println(this.infoStrings[this.languageIndex][AIP_INFO_CHECK_TRANSACTION_ERROR])
			return
		}
	}

	ConsoleColor.Println(ConsoleColor.C_GREEN,
		this.infoStrings[this.languageIndex][AIP_INFO_CHECK_TRANSACTION_SUCCESS])
	// fmt.Println(this.infoStrings[this.languageIndex][AIP_INFO_CHECK_TRANSACTION_SUCCESS])
}

func (this *AccountInfoPayment) checkSourceAddr(addr string, amount float64) *publicdefine.StellarAccInfoDef {
	ret := this.checkPublicAddrExist(addr)
	if ret == nil {
		ConsoleColor.Printf(ConsoleColor.C_RED,
			this.infoStrings[this.languageIndex][AIP_INFO_SOURCE_ACCOUNT_NOT_EXIST]+"\r\n", addr)
	} else {
		balance, _ := strconv.ParseFloat(ret.Balance, 64)
		// 每个账户至少要保留20个币
		if balance < amount || balance-amount < 20 {
			ConsoleColor.Printf(ConsoleColor.C_RED,
				this.infoStrings[this.languageIndex][AIP_INFO_CREDIT_IS_LOW]+"\r\n", addr, ret.Balance)
			// fmt.Printf(this.infoStrings[this.languageIndex][AIP_INFO_CREDIT_IS_LOW]+"\r\n", addr, ret.Balance)
		} else {
			return ret
		}

	}
	return nil
}

func (this *AccountInfoPayment) checkSeed(seed, srcAddr string) bool {
	pk, err := keypair.Parse(seed)
	if err == nil {
		if pk.Address() == srcAddr {
			return true
		}
		ConsoleColor.Printf(ConsoleColor.C_RED,
			this.infoStrings[this.languageIndex][AIP_INFO_SEED_AND_ADDR_IS_NOT_PAIR]+"\r\n")
		// fmt.Printf(this.infoStrings[this.languageIndex][AIP_INFO_SEED_AND_ADDR_IS_NOT_PAIR] + "\r\n")
	} else {
		ConsoleColor.Println(ConsoleColor.C_RED, err)
		// fmt.Println(err)
	}
	return false
}

func (this *AccountInfoPayment) checkPublicAddrExist(addr string) *publicdefine.StellarAccInfoDef {
	reqUrl := publicdefine.STELLAR_DEFAULT_NETWORK + publicdefine.STELLAR_NETWORK_ACCOUNTS + "/" + addr
	resMap, err := this.httpget(reqUrl)

	if err == nil {
		ret := &publicdefine.StellarAccInfoDef{}
		ret.PutMapBody(addr, resMap)
		if ret.IsExist() {
			return ret
		}
	} else {
		ConsoleColor.Println(ConsoleColor.C_RED, err)
		// fmt.Println(err)
	}
	return nil
}

func (this *AccountInfoPayment) create_account(src *publicdefine.StellarAccInfoDef,
	srcSeed, destAddr string, amount float64) bool {

	if len(this.inputConfirm(destAddr)) == 0 {
		ConsoleColor.Println(ConsoleColor.C_RED,
			this.infoStrings[this.languageIndex][AIP_INFO_PAYMENT_ABORT])
		return false
	}

	ConsoleColor.Printf(ConsoleColor.C_BLUE,
		this.infoStrings[this.languageIndex][AIP_INFO_CREATE_DEST_ADDR], destAddr)

	if amount < 20 {
		ConsoleColor.Println(ConsoleColor.C_RED,
			this.infoStrings[this.languageIndex][AIP_INFO_PAYMENT_ABORT_LESS20])
		return false
	}

	cAcc := publicdefine.StellarAccountCreateInfo{
		SrcInfo:    src,
		Amount:     amount,
		Destinaton: destAddr,
	}

	signed := cAcc.GetSigned(srcSeed)

	if len(signed) > 0 {
		data := "tx=" + url.QueryEscape(signed)

		postUrl := publicdefine.STELLAR_DEFAULT_NETWORK + publicdefine.STELLAR_NETWORK_TRANSACTIONS
		ret, err := this.httppost_form(postUrl, data)
		// ret, err := this.httppost_json(postUrl, data)
		if err == nil {
			cAcc.PutResult(ret)
			if len(cAcc.ResultHash) > 0 {
				ConsoleColor.Println(ConsoleColor.C_BLUE,
					this.infoStrings[this.languageIndex][AIP_INFO_CREATE_DEST_ADDR_SUCCESS])
				return true
			}
		}
		ConsoleColor.Println(ConsoleColor.C_RED, err)
		// fmt.Println(err)
	}
	ConsoleColor.Println(ConsoleColor.C_RED,
		this.infoStrings[this.languageIndex][AIP_INFO_CHECK_TRANSACTION_ERROR])
	return false
}

func (this *AccountInfoPayment) inputConfirm(addr string) string {
	fmt.Println("")
	ConsoleColor.Printf(ConsoleColor.C_YELLOW,
		this.infoStrings[this.languageIndex][AIP_INFO_DEST_ADDR_NOT_EXIST], addr)

	var input string

	_, err := fmt.Scanf("%s\n", &input)
	if err == nil {
		if strings.Trim(input, " ") == "yes" {
			return input
		}
	}
	return ""

}

func (this *AccountInfoPayment) pay(src *publicdefine.StellarAccInfoDef,
	srcSeed, destAddr string, amount float64) *publicdefine.StellarPaymentInfo {
	payment := &publicdefine.StellarPaymentInfo{
		SrcInfo:    src,
		Amount:     amount,
		Destinaton: destAddr,
	}

	signed := payment.GetSigned(srcSeed)

	if len(signed) > 0 {
		data := "tx=" + url.QueryEscape(signed)
		// data, err := json.Marshal(map[string]interface{}{
		// 	"tx": signed,
		// })

		postUrl := publicdefine.STELLAR_DEFAULT_NETWORK + publicdefine.STELLAR_NETWORK_TRANSACTIONS
		ret, err := this.httppost_form(postUrl, data)
		// ret, err := this.httppost_json(postUrl, data)
		if err == nil {
			payment.PutResult(ret)
			return payment
		}
		ConsoleColor.Println(ConsoleColor.C_RED, err)
		// fmt.Println(err)
	}
	return nil
}
