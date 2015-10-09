package main

import (
	"github.com/astaxie/beego"
	_ "github.com/jerusalemdax/Mona/models"
	_ "github.com/jerusalemdax/Mona/routers"
)

func main() {
	beego.Run()
}
