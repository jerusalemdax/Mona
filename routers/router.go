package routers

import (
	"github.com/astaxie/beego"
	"github.com/jerusalemdax/Mona/controllers/admin"
	"github.com/jerusalemdax/Mona/controllers/blog"
)

func init() {
	//前台路由
	beego.Router("/", &blog.MainController{}, "*:Index")
	beego.Router("/404", &blog.MainController{}, "*:Go404")
	beego.Router("/index:page:int", &blog.MainController{}, "*:Index")
	beego.Router("/about", &blog.MainController{}, "*:About")

	beego.Router("/life:page:int", &blog.MainController{}, "*:BlogList")
	beego.Router("/life", &blog.MainController{}, "*:BlogList")

	beego.Router("/mood", &blog.MainController{}, "*:Mood")
	beego.Router("/mood:page:int", &blog.MainController{}, "*:Mood")

	beego.Router("/album", &blog.MainController{}, "*:Album")
	beego.Router("/album:page:int", &blog.MainController{}, "*:Album")

	beego.Router("/book", &blog.MainController{}, "*:Book")

	beego.Router("/article/:id:int", &blog.MainController{}, "*:Show")
	beego.Router("/article/:urlname(.+)", &blog.MainController{}, "*:Show")

	beego.Router("/category/:name(.+?)", &blog.MainController{}, "*:Category")
	beego.Router("/category/:name(.+?)/page/:page:int", &blog.MainController{}, "*:Category")

	//照片展示
	beego.Router("/photo", &blog.MainController{}, "*:Photo")
	beego.Router("/photo:page:int", &blog.MainController{}, "*:Photo")

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

	//用户管理
	beego.Router("/admin/user/list", &admin.UserController{}, "*:List")
	beego.Router("/admin/user/add", &admin.UserController{}, "*:Add")
	beego.Router("/admin/user/edit", &admin.UserController{}, "*:Edit")
	beego.Router("/admin/user/delete", &admin.UserController{}, "*:Delete")

	//相册管理
	beego.Router("/admin/album/add", &admin.AlbumController{}, "*:Add")
	beego.Router("/admin/album/list", &admin.AlbumController{}, "*:List")
	beego.Router("/admin/album/edit", &admin.AlbumController{}, "*:Edit")
	beego.Router("/admin/album/delete", &admin.AlbumController{}, "*:Delete")

	//照片管理
	beego.Router("/admin/photo/list", &admin.PhotoController{}, "*:List")
	beego.Router("/admin/photo/cover", &admin.PhotoController{}, "*:Cover")
	beego.Router("/admin/photo/delete", &admin.PhotoController{}, "*:Delete")
	beego.Router("/admin/photo/uploadphoto", &admin.PhotoController{}, "*:UploadPhoto")

	//友链管理
	beego.Router("/admin/link/list", &admin.LinkController{}, "*:List")
	beego.Router("/admin/link/add", &admin.LinkController{}, "*:Add")
	beego.Router("/admin/link/edit", &admin.LinkController{}, "*:Edit")
	beego.Router("/admin/link/delete", &admin.LinkController{}, "*:Delete")

}
