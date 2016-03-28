package handler

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/xwb1989/shortener/storage"
	"net/http"
	"strings"
	"testing"
)

type mockStorage struct{}

func (*mockStorage) Read(key int) (string, error) {
	if key == -1 {
		return "", storage.InvalidKeyError()
	}
	return "a_valid_url", nil
}

func (*mockStorage) Write(k int, v string) error {
	if k == -1 {
		return errors.New("unable to write to the storage")
	} else {
		return nil
	}
}

type mockEncoder struct{}

func (*mockEncoder) Encode(s string) int {
	if strings.Contains(s, "invalid") {
		return -1
	} else {
		return 0
	}
}

func (*mockEncoder) ToString(i int) string {
	if i == -1 {
		return "invalid"
	} else {
		return "encoded"
	}
}

func (*mockEncoder) FromString(s string) int {
	if strings.Contains(s, "invalid") {
		return -1
	}
	return 0
}

type mockResponseWriter struct {
	received string
	status   int
}

func (*mockResponseWriter) Header() http.Header {
	return http.Header{}
}

func (w *mockResponseWriter) Write(in []byte) (int, error) {
	w.received = string(in)
	if w.status == 0 {
		w.status = http.StatusOK
	}
	return len(in), nil
}

func (w *mockResponseWriter) WriteHeader(status int) {
	w.status = status
}

func TestHandler(t *testing.T) {
	Convey("With router, storage, and encoder...", t, func() {
		storage := &mockStorage{}
		encoder := &mockEncoder{}
		router := httprouter.New()
		Convey("we can serve incoming encoding request", func() {
			encode := Shorten(storage, encoder)
			router.POST("/:url", encode)

			writer := &mockResponseWriter{}

			request, _ := http.NewRequest(http.MethodPost, "/an_valid_url", nil)

			// serve the request
			router.ServeHTTP(writer, request)

			// check response
			So(writer.status, ShouldEqual, http.StatusOK)
			So(writer.received, ShouldContainSubstring, "encoded")

			Convey("and get error if url is empty", func() {
				request.URL.Path = "/"
				router.ServeHTTP(writer, request)
				So(writer.status, ShouldEqual, http.StatusNotFound)
			})
			Convey("and get 500 if unable to write to storage", func() {
				request.URL.Path = "/an_invalid_url"
				router.ServeHTTP(writer, request)
				So(writer.status, ShouldEqual, http.StatusInternalServerError)
			})
		})
		Convey("we can also serve decoding request", func() {
			decode := Redirect(storage, encoder)
			router.GET("/:url", decode)

			writer := &mockResponseWriter{}

			request, _ := http.NewRequest(http.MethodGet, "/a_valid_url", nil)
			router.ServeHTTP(writer, request)
			So(writer.status, ShouldEqual, http.StatusTemporaryRedirect)
			So(writer.received, ShouldContainSubstring, "href=\"/a_valid_url\"")
			Convey("and get 400 if key is empty", func() {
				request.URL.Path = "/"
				router.ServeHTTP(writer, request)
				So(writer.status, ShouldEqual, http.StatusNotFound)
			})
			Convey("and get 404 if key is invalid", func() {
				request.URL.Path = "/an_invalid_string"
				router.ServeHTTP(writer, request)
				So(writer.status, ShouldEqual, http.StatusNotFound)
				So(writer.received, ShouldContainSubstring, "unable to get url for key an_invalid_string")
			})
		})
	})
}
