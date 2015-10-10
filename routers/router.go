package routers

import (
	"github.com/astaxie/beego"
	"github.com/jerusalemdax/Mona/controllers/admin"
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

	beego.Router("/mood.html", &blog.MainController{}, "*:Mood")
	beego.Router("/mood:page:int.html", &blog.MainController{}, "*:Mood")

	beego.Router("/album.html", &blog.MainController{}, "*:Album")
	beego.Router("/album:page:int.html", &blog.MainController{}, "*:Album")

	beego.Router("/book.html", &blog.MainController{}, "*:Book")

	beego.Router("/:urlname(.+)", &blog.MainController{}, "*:Show")

	//后台路由
	beego.Router("/admin", &admin.IndexController{}, "*:Index")
	beego.Router("/admin/login", &admin.AccountController{}, "*:Login")
	beego.Router("/admin/logout", &admin.AccountController{}, "*:Logout")
	beego.Router("/admin/account/profile", &admin.AccountController{}, "*:Profile")
}
