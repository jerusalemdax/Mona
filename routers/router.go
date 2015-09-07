package routers

import (
	"github.com/astaxie/beego"
	"github.com/jerusalemdax/Mona/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})

	//chat
	beego.Router("/chat", &controllers.ChatController{})
	beego.Router("/chat/join", &controllers.ChatController{}, "post:Join")

	//longpolling
	beego.Router("/longpolling", &controllers.LongPollingController{}, "get:Join")
	beego.Router("/longpolling/post", &controllers.LongPollingController{})
	beego.Router("/longpolling/fetch", &controllers.LongPollingController{}, "get:Fetch")

	// WebSocket.
	beego.Router("/websocket", &controllers.WebSocketController{})
	beego.Router("/websocket/join", &controllers.WebSocketController{}, "get:Join")
}
