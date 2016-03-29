package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/xwb1989/shortener/handler"
	"github.com/xwb1989/shortener/storage"
	"github.com/xwb1989/shortener/storage/encoder"
	"log"
	"net/http"
)

const (
	Port int = 8090
)

func main() {
	e := encoder.NewIncrementalEncoder(0)
	s := storage.NewMemMap(e)

	router := httprouter.New()
	router.POST("/encode", handler.Shorten(s))
	router.GET("/:url", handler.Redirect(s))

	log.Println("Server is up and running at port", Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Port), router))
}
