package menu

import (
	// "fmt"
	"go-StellarWallet/publicdefine"
	"os"
)

type ExitApp struct {
	MenuSubItem
}

func (this *ExitApp) InitExitApp(parent MenuSubItemInterface, key string) {
	this.MenuSubItem.InitMenu(key)
	this.parentItem = parent
	this.MenuSubItem.Exec = this.execute

	this.MenuSubItem.title = []string{
		publicdefine.L_Chinese: "退出钱包程序",
		publicdefine.L_English: "Exit",
	}
}

func (this *ExitApp) execute(isSync bool) {
	os.Exit(0)
}
