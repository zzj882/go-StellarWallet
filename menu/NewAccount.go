package menu

import (
	"fmt"
	"github.com/jojopoper/ConsoleColor"
	"github.com/stellar/go-stellar-base"
	"go-StellarWallet/publicdefine"
	"os"
	"time"
)

const (
	CA_INFO_SECRET_SEED = iota
	CA_INFO_PUBLIC_ADDR
	CA_INFO_MEMO_TEXT
	CA_INFO_MEMO_SAVEFILE
	CA_INFO_MEMO_SAVEFILE_ERR
)

type CreateAccount struct {
	MenuSubItem
	infoStrings []map[int]string
}

func (this *CreateAccount) InitCreator(parent MenuSubItemInterface, key string) {
	this.MenuSubItem.InitMenu(key)
	this.parentItem = parent
	this.MenuSubItem.Exec = this.execute

	this.MenuSubItem.title = []string{
		publicdefine.L_Chinese: "创建账户",
		publicdefine.L_English: "Create new account",
	}
	this.infoStrings = []map[int]string{
		publicdefine.L_Chinese: map[int]string{
			CA_INFO_SECRET_SEED:       " Secret seed:",
			CA_INFO_PUBLIC_ADDR:       " Public:",
			CA_INFO_MEMO_TEXT:         " 需要保存账户信息到文件请输入s，否则输入任意键返回菜单: ",
			CA_INFO_MEMO_SAVEFILE:     " 保存账户信息完成，请妥善保存好文件 %s",
			CA_INFO_MEMO_SAVEFILE_ERR: " 保存账户信息失败，错误信息: ",
		},
		publicdefine.L_English: map[int]string{
			CA_INFO_SECRET_SEED:       " Secret seed:",
			CA_INFO_PUBLIC_ADDR:       " Public:",
			CA_INFO_MEMO_TEXT:         " If you need to save account informations to a file then press s, or press any key to return menu: ",
			CA_INFO_MEMO_SAVEFILE:     " Save account information is complete，please keep safe the file : %s",
			CA_INFO_MEMO_SAVEFILE_ERR: " Save account information is failure, the error message: ",
		},
	}
}

func (this *CreateAccount) execute(isSync bool) {
	var input string
	b, err := publicdefine.SafeRandomBytes(32)
	if err == nil {
		seed, err := stellarbase.NewRawSeed(b)
		if err == nil {
			pubKey, priKey, err := stellarbase.GenerateKeyFromRawSeed(seed)

			if err == nil {
				fmt.Printf("\r\n"+this.infoStrings[this.languageIndex][CA_INFO_SECRET_SEED]+" %s\r\n", priKey.Seed())
				fmt.Printf(this.infoStrings[this.languageIndex][CA_INFO_PUBLIC_ADDR]+" %s\r\n", pubKey.Address())
				fmt.Printf("\r\n" + this.infoStrings[this.languageIndex][CA_INFO_MEMO_TEXT])
				fmt.Scanf("%s\n", &input)

				if input == "s" {
					err = this.savefile("account_info.txt", priKey.Seed(), pubKey.Address())
					if err == nil {
						ConsoleColor.Printf(ConsoleColor.C_YELLOW,
							"\r\n"+this.infoStrings[this.languageIndex][CA_INFO_MEMO_SAVEFILE]+"\r\n\r\n", "account_info.txt")
						// fmt.Printf("\r\n"+this.infoStrings[this.languageIndex][CA_INFO_MEMO_SAVEFILE]+"\r\n\r\n", "account_info.txt")
					} else {
						ConsoleColor.Println(ConsoleColor.C_RED,
							"\r\n", this.infoStrings[this.languageIndex][CA_INFO_MEMO_SAVEFILE_ERR], "\r\n", err, "\r\n\r\n")
						// fmt.Println("\r\n", this.infoStrings[this.languageIndex][CA_INFO_MEMO_SAVEFILE_ERR], "\r\n", err, "\r\n\r\n")
					}
				}
			} else {
				fmt.Println(err.Error())
				fmt.Scanf("%s\n", &input)
			}
		} else {
			fmt.Println(err.Error())
			fmt.Scanf("%s\n", &input)
		}
	} else {
		fmt.Println(err.Error())
		fmt.Scanf("%s\n", &input)
	}
	if !isSync {
		this.ASyncChan <- 0
	}
}

func (this *CreateAccount) savefile(filepath, seed, addr string) error {
	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeType)
	if err != nil {
		return err
	}
	defer f.Close()
	sav := fmt.Sprintf("\r\n================= %s =================\r\n", time.Now().Format("2006-01-02 15:04:05"))
	sav += fmt.Sprintf(this.infoStrings[this.languageIndex][CA_INFO_SECRET_SEED]+" %s\r\n", seed)
	sav += fmt.Sprintf(this.infoStrings[this.languageIndex][CA_INFO_PUBLIC_ADDR]+" %s\r\n\r\n", addr)
	_, err = f.WriteString(sav)
	return err
}
