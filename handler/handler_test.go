package handler

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/xwb1989/shortener/storage"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

type mockStorage struct{}

func (*mockStorage) Read(s string) (string, error) {
	if strings.Contains(s, "valid") {
		return "a_valid_url", nil
	}
	return "", storage.InvalidKeyError()
}

func (*mockStorage) Write(k, v string) error {
	if strings.Contains(k, "valid") {
		return nil
	} else {
		return errors.New("unable to write to the storage")
	}
}

type mockEncoder struct{}

func (*mockEncoder) Encode(s string) string {
	return "encoded" + s
}

func (*mockEncoder) Decode(s string) string {
	return "decoded" + s
}

type mockResponseWriter struct {
	received string
	status   int
}

func (*mockResponseWriter) Header() http.Header {
	return nil
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
	Convey("With mock up storage and encoder...", t, func() {
		storage := &mockStorage{}
		encoder := &mockEncoder{}
		Convey("we can serve incoming encoding request", func() {
			encode := Encode(storage, encoder)

			request, _ := http.NewRequest(http.MethodPost, "localhost", nil)
			request.PostForm = url.Values{"url": {"a_valid_string"}}
			writer := &mockResponseWriter{}
			encode.ServeHTTP(writer, request)

			So(writer.status, ShouldEqual, http.StatusOK)
			So(writer.received, ShouldContainSubstring, "encoded")
		})
	})
}
