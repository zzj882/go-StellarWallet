package main

import (
	"errors"
	"fmt"
	"github.com/jojopoper/ConsoleColor"
	"go-StellarWallet/menu"
	"go-StellarWallet/publicdefine"
	"strconv"
)

func main() {
	currLanguage, errMsg := selectLanguage(3)
	if errMsg == nil {
		menu.MainMenuInstace.SetLanguageType(currLanguage)
		menu.MainMenuInstace.ExecuteFunc(true)
	}
}

func selectLanguage(maxRetry int) (int, error) {
	for i := 0; i < maxRetry; i++ {
		fmt.Printf("选择语言(select language):\r\n中文输入 %d 回车，For English press %d + Enter: ",
			publicdefine.L_Chinese+1, publicdefine.L_English+1)
		var input string
		_, err := fmt.Scanf("%s\n", &input)
		if err != nil {
			fmt.Println(err)
			fmt.Scanf("%s\n", &input)
			return -1, err
		}

		switch input {
		case strconv.Itoa(publicdefine.L_Chinese + 1):
			return publicdefine.L_Chinese, nil
		case strconv.Itoa(publicdefine.L_English + 1):
			return publicdefine.L_English, nil
		default:
			ConsoleColor.Println(ConsoleColor.C_RED, "语言选择错误(Language selection error)\r\n")
			// fmt.Println("语言选择错误(Language selection error)\r\n")
		}
	}
	ConsoleColor.Println(ConsoleColor.C_YELLOW, "ByeBye!\r\n")
	return -1, errors.New("Language selection error")
}
