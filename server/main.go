package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"

	"kokoichi206-sandbox/gh-actions-screen/server/handler"
	"kokoichi206-sandbox/gh-actions-screen/server/util"
)

func main() {
	uuid, _ := util.GetUuid()
	fmt.Println(uuid)

	server := handler.NewServer()

	http.Handle("/ws/actions", websocket.Handler(server.HandleActions))

	http.ListenAndServe(":3333", nil)
}
