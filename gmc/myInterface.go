package gmc

type myDB interface {
	CreatUser(name string ,pass string) bool
	CheckStats(name string ) int
	AddMoney(name string , money int) bool
	SubMoney(name string , money int) bool

}


