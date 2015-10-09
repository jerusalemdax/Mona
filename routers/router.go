package routers

import (
	"github.com/astaxie/beego"
	"github.com/jerusalemdax/Mona/controllers/blog"
)

func init() {
	//前台路由
	beego.Router("/", &blog.MainController{}, "*:Index")
	beego.Router("/404.html", &blog.MainController{}, "*:Go404")
	beego.Router("/index:page:int.html", &blog.MainController{}, "*:Index")
	beego.Router("/about.html", &blog.MainController{}, "*:About")

	beego.Router("/life:page:int.html", &blog.MainController{}, "*:BlogList")
	beego.Router("/life.html", &blog.MainController{}, "*:BlogList")

	beego.Router("/:urlname(.+)", &blog.MainController{}, "*:Show")
}
