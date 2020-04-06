package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func Websocket(w http.ResponseWriter,r *http.Request)  {
	websocketConn,err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return  true
		},
	}).Upgrade(w,r,nil)
	if err != nil{
		log.Fatal("升级websocket出错:",err.Error())
	}
	fmt.Println(websocketConn)
	fmt.Println("success connect ws")

	for  {
		//获取客户端发送的消息
		_, data, e := websocketConn.ReadMessage()
		if e != nil{
			log.Println("can't receive:",err.Error())
			break
		}
		log.Printf("[ws]<=%s\n", data)

		fmt.Println("sending to client: ",data)
		err := websocketConn.WriteMessage(websocket.TextMessage, data)
		if err != nil{
			fmt.Println("send message err:",err.Error())
			break
		}
	}

}

func main()  {

	http.HandleFunc("/",Websocket)
	err := http.ListenAndServe(":1234", nil)
	if err != nil{
		log.Fatal("监听http服务出错：",err.Error())
	}
}
