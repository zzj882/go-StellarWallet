package menu

import (
	"fmt"
	"github.com/jojopoper/go-StellarWallet/publicdefine"
)

type MenuSubItem struct {
	title         []string
	Exec          func(isSync bool)
	subItems      []MenuSubItemInterface
	parentItem    MenuSubItemInterface
	keyPath       string
	ASyncChan     chan int
	languageIndex int
	inputMemo     []string
}

type MenuSubItemInterface interface {
	InitMenu(key string)
	SetLanguageType(langType int)
	GetTitle(langType int) string
	HasTitle() bool
	SetTitle(langType int, t string)
	GetSubItems() []MenuSubItemInterface
	AddSubItem(itm MenuSubItemInterface) int
	GetParentItem() MenuSubItemInterface
	SetParentItem(p MenuSubItemInterface)
	GetKeyPath() string
	SetKeyPath(kp string)
	GetTitlePath(langType int) string
	ExecuteFunc(isSync bool)
	ExecFlag() int
	PrintSubmenu()
	GetInputMemo(langType int) string
}

func (this *MenuSubItem) InitMenu(key string) {
	this.ASyncChan = make(chan int)
	this.title = make([]string, 2)
	this.subItems = make([]MenuSubItemInterface, 0)
	this.keyPath = key
	this.inputMemo = []string{
		publicdefine.L_Chinese: "请选择菜单列表项目对应的数字并回车: ",
		publicdefine.L_English: "Select the number of items on the menu list item and press enter: ",
	}
}

func (this *MenuSubItem) SetLanguageType(langType int) {
	this.languageIndex = langType
	for _, sub := range this.subItems {
		sub.SetLanguageType(langType)
	}
}

func (this *MenuSubItem) GetTitle(langType int) string {
	if this.HasTitle() {
		return this.title[langType]
	}
	return ""
}

func (this *MenuSubItem) HasTitle() bool {
	return len(this.title) > 0
}

func (this *MenuSubItem) SetTitle(langType int, t string) {
	this.title[langType] = t
}

func (this *MenuSubItem) GetSubItems() []MenuSubItemInterface {
	return this.subItems
}

func (this *MenuSubItem) AddSubItem(itm MenuSubItemInterface) int {
	length := len(this.subItems)
	itm.SetParentItem(this)
	itm.SetKeyPath(fmt.Sprintf("%s.%d", this.keyPath, length))
	this.subItems = append(this.subItems, itm)
	return length
}

func (this *MenuSubItem) GetParentItem() MenuSubItemInterface {
	return this.parentItem
}

func (this *MenuSubItem) SetParentItem(p MenuSubItemInterface) {
	this.parentItem = p
}

func (this *MenuSubItem) GetKeyPath() string {
	return this.keyPath
}

func (this *MenuSubItem) SetKeyPath(kp string) {
	this.keyPath = kp
}

func (this *MenuSubItem) GetTitlePath(langType int) (ret string) {
	if this.parentItem == nil {
		ret = this.title[langType]
	} else {
		ret = this.parentItem.GetTitlePath(langType) + " > " + this.title[langType]
	}
	return ret
}

func (this *MenuSubItem) ExecuteFunc(isSync bool) {
	if this.Exec != nil {
		if isSync {
			this.Exec(isSync)
		} else {
			go this.Exec(isSync)
		}
	}
}

func (this *MenuSubItem) ExecFlag() int {
	return <-this.ASyncChan
}

func (this *MenuSubItem) PrintSubmenu() {
	length := len(this.subItems)
	for i := 0; i < length; i++ {
		fmt.Printf(" %d.\t%s\r\n", i+1, this.subItems[i].GetTitle(this.languageIndex))
	}
}

func (this *MenuSubItem) GetInputMemo(langType int) string {
	return this.inputMemo[langType]
}
