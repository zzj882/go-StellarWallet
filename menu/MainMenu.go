package menu

import (
	"fmt"
	// "github.com/jojopoper/go-StellarWallet/createAccount"
	// "github.com/jojopoper/ConsoleColor"
	"github.com/jojopoper/go-StellarWallet/publicdefine"
)

type MenuInfo struct {
	MenuSubItem
	currentLevel  int
	WelcomeString []string
}

var MainMenuInstace *MenuInfo

func init() {
	MainMenuInstace = new(MenuInfo)
	MainMenuInstace.MenuSubItem.InitMenu("0")
	MainMenuInstace.currentLevel = 0
	MainMenuInstace.MenuSubItem.title = []string{
		publicdefine.L_Chinese: "菜单",
		publicdefine.L_English: "Menu",
	}
	MainMenuInstace.WelcomeString = []string{
		publicdefine.L_Chinese: " ##     欢迎使用恒星币钱包，请选择您需要的功能     ##\n",
		publicdefine.L_English: " ##       Welcome to use the stellar wallet        ##\n" +
			" ##      please choose the function you need       ##\n",
	}
	MainMenuInstace.MenuSubItem.Exec = MainMenuInstace.Execute

	creatAcc := CreateAccount{}
	creatAcc.InitCreator(MainMenuInstace, "0")
	MainMenuInstace.MenuSubItem.AddSubItem(&creatAcc)

	accInfo := AccountInfo{}
	accInfo.InitAccInfo(MainMenuInstace, "0")
	MainMenuInstace.MenuSubItem.AddSubItem(&accInfo)

	mergeAcc := MergeAccount{}
	mergeAcc.InitMerge(MainMenuInstace, "0")
	MainMenuInstace.MenuSubItem.AddSubItem(&mergeAcc)

	about := SoftwareAbout{}
	about.InitAbout(MainMenuInstace, "0")
	MainMenuInstace.MenuSubItem.AddSubItem(&about)

	exitapp := ExitApp{}
	exitapp.InitExitApp(MainMenuInstace, "0")
	MainMenuInstace.MenuSubItem.AddSubItem(&exitapp)
}

func (this *MenuInfo) Execute(isSync bool) {
	for {
		fmt.Println("\r\n******************************************************")
		fmt.Println(this.getWelcomeString(this.languageIndex))
		fmt.Println(" " + this.GetTitle(this.languageIndex))
		this.PrintSubmenu()
		fmt.Printf("\n %s", this.GetInputMemo(this.languageIndex))

		var input string

		_, err := fmt.Scanf("%s\n", &input)
		if err == nil {
			selectIndex, b := publicdefine.IsNumber(input)
			if b {
				if selectIndex <= len(this.subItems) && selectIndex > 0 {
					this.subItems[selectIndex-1].ExecuteFunc(false)
					this.subItems[selectIndex-1].ExecFlag()
				}
			}
		}
	}
}

func (this *MenuInfo) getWelcomeString(langType int) string {
	return this.WelcomeString[langType]
}
