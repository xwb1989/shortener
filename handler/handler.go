package handler

import (
	"fmt"
	"github.com/xwb1989/shortener/encoder"
	"github.com/xwb1989/shortener/storage"
	"log"
	"net/http"
)

const (
	UrlParamName string = "url"
)

func Encode(s storage.Writer, e encoder.Encoder) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if url := r.PostFormValue(UrlParamName); url != "" {
			key := e.Encode(url)
			err := s.Write(key, url)
			if err != nil {
				msg := fmt.Sprint("unable to write to storage: ", url)
				log.Println(msg)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(msg))
			} else {
				_, err = w.Write([]byte(key))
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			msg := fmt.Sprintf("invalid url: %s", url)
			w.Write([]byte(msg))
		}
	}
	return http.HandlerFunc(handler)
}

func Decode(s storage.Reader) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if key := r.URL.Path; key != "" {
			log.Println("decoding:", key)
			res, err := s.Read(key)
			if err != nil {
				msg := fmt.Sprintf("invalid short url: %s", key)
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(msg))
			} else {
				w.Write([]byte(res))
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			msg := fmt.Sprintf("invalid short url: %s.", key)
			w.Write([]byte(msg))
		}
	}
	return http.HandlerFunc(handler)
}

func Redirect(s storage.Reader) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if key := r.URL.Path; key != "" {
			res, err := s.Read(key)
			if err != nil {
				msg := fmt.Sprintf("invalid short url: %s", key)
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(msg))
			} else {
				http.Redirect(w, r, res, http.StatusTemporaryRedirect)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			msg := fmt.Sprintf("invalid short url: %s.", key)
			w.Write([]byte(msg))
		}
	}
	return http.HandlerFunc(handler)
}
