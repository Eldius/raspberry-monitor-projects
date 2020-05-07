/*
Package handlers is where I will put the handlers
*/
package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/websocket"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Index is the handler for index path
func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}
