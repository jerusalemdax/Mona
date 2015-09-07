package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	_ "github.com/jerusalemdax/Mona/routers"
	"strings"
)

func init() {
	langTypes := strings.Split(beego.AppConfig.String("lang"), "|")

	for _, lang := range langTypes {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "lang/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file:", err)
			return
		}
	}
}

func main() {
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}
