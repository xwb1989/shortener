package handler

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/xwb1989/shortener/storage"
	"log"
	"net/http"
)

const (
	UrlParamName string = "url"
)

func Shorten(s storage.Writer) httprouter.Handle {
	handler := func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		url := params.ByName(UrlParamName)
		key, err := s.Write(url)
		if err != nil {
			msg := fmt.Sprint("unable to write %s to storage: ", url, err.Error())
			log.Println(msg)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(msg))
		} else {
			_, err = w.Write([]byte(key))
		}
	}
	return httprouter.Handle(handler)
}

func Redirect(s storage.Reader) httprouter.Handle {
	handler := func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		key := params.ByName(UrlParamName)
		res, err := s.Read(key)
		if err != nil {
			msg := fmt.Sprintf("unable to get url for key %s: %s", key, err.Error())
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(msg))
		} else {
			http.Redirect(w, r, res, http.StatusTemporaryRedirect)
		}
	}
	return httprouter.Handle(handler)
}
