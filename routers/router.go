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

	beego.Router("/article/:id:int", &blog.MainController{}, "*:Show")
	beego.Router("/article/:urlname(.+)", &blog.MainController{}, "*:Show")

	beego.Router("/category/:name(.+?)", &blog.MainController{}, "*:Category")
	beego.Router("/category/:name(.+?)/page/:page:int", &blog.MainController{}, "*:Category")

	beego.Router("/:urlname(.+)", &blog.MainController{}, "*:Show")

	//后台路由
	beego.Router("/admin", &admin.IndexController{}, "*:Index")
	beego.Router("/admin/login", &admin.AccountController{}, "*:Login")
	beego.Router("/admin/logout", &admin.AccountController{}, "*:Logout")
	beego.Router("/admin/account/profile", &admin.AccountController{}, "*:Profile")

	beego.Router("/admin/system/setting", &admin.SystemController{}, "*:Setting")

	//内容管理
	beego.Router("/admin/article/list", &admin.ArticleController{}, "*:List")
	beego.Router("/admin/article/add", &admin.ArticleController{}, "*:Add")
	beego.Router("/admin/article/edit", &admin.ArticleController{}, "*:Edit")
	beego.Router("/admin/article/save", &admin.ArticleController{}, "post:Save")
	beego.Router("/admin/article/delete", &admin.ArticleController{}, "*:Delete")
	beego.Router("/admin/article/batch", &admin.ArticleController{}, "*:Batch")
	beego.Router("/admin/article/upload", &admin.ArticleController{}, "*:Upload")
	beego.Router("/admin/tag", &admin.TagController{}, "*:Index")

	beego.Router("/admin/mood/add", &admin.MoodController{}, "*:Add")
	beego.Router("/admin/mood/list", &admin.MoodController{}, "*:List")
	beego.Router("/admin/mood/delete", &admin.MoodController{}, "*:Delete")
}
