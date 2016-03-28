package handler

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/xwb1989/shortener/encoder"
	"github.com/xwb1989/shortener/storage"
	"log"
	"net/http"
)

const (
	UrlParamName string = "url"
)

func Shorten(s storage.Writer, e encoder.Encoder) httprouter.Handle {
	handler := func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		url := params.ByName(UrlParamName)
		key := e.Encode(url)
		err := s.Write(key, url)
		if err != nil {
			msg := fmt.Sprint("unable to write %s to storage: ", url, err.Error())
			log.Println(msg)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(msg))
		} else {
			_, err = w.Write([]byte(e.ToString(key)))
		}
	}
	return httprouter.Handle(handler)
}

func Redirect(s storage.Reader, e encoder.Encoder) httprouter.Handle {
	handler := func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		key := params.ByName(UrlParamName)
		res, err := s.Read(e.FromString(key))
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
