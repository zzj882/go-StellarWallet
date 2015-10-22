package menu

import (
	"fmt"
	"github.com/jojopoper/go-StellarWallet/publicdefine"
)

type AccountInfo struct {
	MenuSubItem
	infoStrings []map[int]string
}

func (this *AccountInfo) InitAccInfo(parent MenuSubItemInterface, key string) {
	this.MenuSubItem.InitMenu(key)
	this.parentItem = parent
	this.MenuSubItem.Exec = this.execute

	this.MenuSubItem.title = []string{
		publicdefine.L_Chinese: "账户信息",
		publicdefine.L_English: "Account Informations",
	}

	this.infoStrings = []map[int]string{
		publicdefine.L_Chinese: map[int]string{},
		publicdefine.L_English: map[int]string{},
	}

	baseinfo := AccountInfoBase{}
	baseinfo.InitAccInfoBase(this, "1")
	this.AddSubItem(&baseinfo)

	payinfo := AccountInfoPayment{}
	payinfo.InitAccInfoPayment(this, "1")
	this.AddSubItem(&payinfo)

	operquary := AccountInfoOperationQuary{}
	operquary.InitAccInfoOperQuary(this, "1")
	this.AddSubItem(&operquary)

	returnParent := ReturnParentMenu{}
	returnParent.InitReturnParentMenu(this, "1")
	this.AddSubItem(&returnParent)

	exitapp := ExitApp{}
	exitapp.InitExitApp(this, "1")
	this.AddSubItem(&exitapp)
}

func (this *AccountInfo) execute(isSync bool) {
	for {
		fmt.Println("\n\n", this.GetTitlePath(this.languageIndex), "\r\n")
		this.PrintSubmenu()
		fmt.Printf("\n %s", this.GetInputMemo(this.languageIndex))

		var input string

		_, err := fmt.Scanf("%s\n", &input)
		if err == nil {
			selectIndex, b := publicdefine.IsNumber(input)
			if b {
				if selectIndex <= len(this.subItems) && selectIndex >= 0 {
					this.subItems[selectIndex-1].ExecuteFunc(false)
					ret := this.subItems[selectIndex-1].ExecFlag()
					if ret == BACK_TO_MENU_FLAG {
						break
					}
				}
			}
		} else {
			fmt.Println(err)
		}
	}
	if !isSync {
		this.ASyncChan <- 1
	}
}
