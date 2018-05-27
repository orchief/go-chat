package controllers

import (
	"github.com/gorilla/websocket"
	"net/http"
	"github.com/astaxie/beego"
	"WebIm/models"
	"encoding/json"
)

type WebSocketController struct {
	baseController
}

func (this *WebSocketController) Get() {
	uname := this.GetString("uname")
	if len(uname) == 0{
		this.Redirect("/", 302)
		return
	}
	this.TplName = "websocket.html"
	this.Data["IsWebSocket"] = true
	this.Data["UserName"] = uname
}

func (this *WebSocketController) Join() {
	uname := this.GetString("uname")
	if len(uname) == 0{
		this.Redirect("/", 302)
		return
	}
	w := websocket.Upgrader{}
	ws, err := w.Upgrade(this.Ctx.ResponseWriter,this.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "not a websocket handshake", 400)
		return
	} else if err != nil{
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	Join(uname, ws)
	defer Leave(uname)

	for  {
		_, p, err := ws.ReadMessage()
		if err != nil{
			return
		}
		publish <- newEvent(models.EVENT_MESSAGE, uname, string(p))
	}
}

func broadcastWebSocket(event models.Event)  {
	data, err := json.Marshal(event)
	if err != nil{
		beego.Error("fail to marshal event:", err)
		return
	}

	for sub := subcribers.Front(); sub != nil; sub = sub.Next(){
		ws := sub.Value.(Subscriber).Conn
		if ws != nil{
			if ws.WriteMessage(websocket.TextMessage, data) != nil{
				unsubscribe <- sub.Value.(Subscriber).Name
			}
		}
	}
}