package controllers

type ChatController struct {
	baseController
}

func (this *ChatController) Get() {
	this.TplNames = "chat.html"
}

func (this *ChatController) Join() {
	uname := this.GetString("uname")
	tech := this.GetString("tech")

	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	switch tech {
	case "longpolling":
		this.Redirect("/longpolling?uname="+uname, 302)
	case "websocket":
		this.Redirect("/websocket?uname="+uname, 302)
	default:
		this.Redirect("/", 302)
	}

	return
}
