package blog

import (
	"strconv"
	"strings"

	"github.com/jerusalemdax/Mona/models"
)

type MainController struct {
	baseController
}

//首页, 只显示前N条
func (this *MainController) Index() {
	var list []*models.Post
	query := new(models.Post).Query().Filter("status", 0).Filter("urltype", 0)
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-istop", "-views").Limit(this.pagesize, (this.page-1)*this.pagesize).All(&list)
	}
	this.Data["list"] = list

	this.Data["pagebar"] = models.NewPager(int64(this.page), int64(count), int64(this.pagesize), "/index%d.html").ToString()
	this.setHeadMetas()
	this.display("index")
}

//留404页面
func (this *MainController) Go404() {
	this.setHeadMetas("Sorry 404页面没找到")
	this.display("404")
}

func (this *MainController) About() {
	this.setHeadMetas("关于Judy")
	this.right = ""
	this.display("about")
}

//文章显示
func (this *MainController) Show() {
	var (
		post *models.Post = new(models.Post)
		err  error
	)
	urlname := this.Ctx.Input.Param(":urlname")
	if urlname != "" {
		post.Urlname = urlname
		err = post.Read("urlname")
	} else {
		id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
		post.Id = int64(id)
		err = post.Read()
	}
	if err != nil || post.Status != 0 {
		this.Redirect("/404.html", 302)
	}
	post.Views++
	post.Update("Views")
	models.Cache.Delete("hotblog")
	post.Content = strings.Replace(post.Content, "_ueditor_page_break_tag_", "", -1)
	pre, next := post.GetPreAndNext()
	this.Data["post"] = post
	this.Data["pre"] = pre
	this.Data["next"] = next
	this.Data["smalltitle"] = "文章内容"
	if urlname == "about.html" {
		this.Data["smalltitle"] = "关于我"
	}

	this.setHeadMetas(post.Title, strings.Trim(post.Tags, ","), post.Title)
	this.display("article")
}

//blog分页显示
func (this *MainController) BlogList() {
	var list []*models.Post
	query := new(models.Post).Query().Filter("status", 0).Filter("urltype", 0)
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-istop", "-posttime").Limit(this.pagesize, (this.page-1)*this.pagesize).All(&list)
	}
	this.Data["list"] = list
	this.Data["pagebar"] = models.NewPager(int64(this.page), int64(count), int64(this.pagesize), "/life%d.html").ToString()
	this.setHeadMetas("点点滴滴")
	this.right = ""
	this.display("life")
}

func (this *MainController) Mood() {
	var list []*models.Mood
	query := new(models.Mood).Query()
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-posttime").Limit(this.pagesize, (this.page-1)*this.pagesize).All(&list)
	}
	this.Data["list"] = list
	this.setHeadMetas("碎言碎语")
	this.right = ""
	this.Data["pagebar"] = models.NewPager(int64(this.page), int64(count), int64(this.pagesize), "/mood%d.html").ToString()
	this.display("mood")
}

func (this *MainController) Album() {
	pagesize, _ := strconv.Atoi(this.getOption("albumsize"))
	if pagesize < 1 {
		pagesize = 12
	}
	var list []*models.Album
	query := new(models.Album).Query().Filter("ishide", 0)
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-rank", "-posttime").Limit(pagesize, (this.page-1)*pagesize).All(&list)
	}
	this.setHeadMetas("美丽瞬间")
	this.right = ""
	this.Data["list"] = list
	this.Data["pagebar"] = models.NewPager(int64(this.page), int64(count), int64(pagesize), "/album%d.html").ToString()
	this.display("album")
}

func (this *MainController) Book() {
	this.setHeadMetas("游客足迹")
	this.right = "about.html"
	this.display("book")
}
