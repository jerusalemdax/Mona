package routers

import (
	"github.com/jerusalemdax/Mona/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
