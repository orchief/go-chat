package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	_"WebIM/routers"
)

const APP_VER  = "1.0"

func main() {
	beego.Info(beego.BConfig.AppName, APP_VER)
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}