package menu

import (
	"github.com/Ledgercn/go-StellarWallet/publicdefine"
)

const (
	BACK_TO_MENU_FLAG = 99999
)

type ReturnParentMenu struct {
	MenuSubItem
	infoStrings []map[int]string
}

func (this *ReturnParentMenu) InitReturnParentMenu(parent MenuSubItemInterface, key string) {
	this.MenuSubItem.InitMenu(key)
	this.parentItem = parent
	this.MenuSubItem.Exec = this.execute

	this.MenuSubItem.title = []string{
		publicdefine.L_Chinese: "返回上一级",
		publicdefine.L_English: "Go back",
	}
}

func (this *ReturnParentMenu) execute(isSync bool) {
	if !isSync {
		this.ASyncChan <- BACK_TO_MENU_FLAG
	}
}
