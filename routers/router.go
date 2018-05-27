package routers

import (
	"github.com/astaxie/beego"
	"WebIm/controllers"
)

func init() {
	beego.Router("/", &controllers.AppController{})
	beego.Router("/join", &controllers.AppController{}, "post:Join")

	beego.Router("/lp", &controllers.LongPollingController{}, "get:Join")
	beego.Router("/lp/post", &controllers.LongPollingController{})
	beego.Router("/lp/fetch", &controllers.LongPollingController{}, "get:Fetch")

	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")
}
