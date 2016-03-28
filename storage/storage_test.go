package storage

import (
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
)

func TestStorage(t *testing.T) {
	Convey("Simple memory storage should be able to...", t, func() {
		s := NewMemMap()
		n := 100
		Convey("should be able to write then read", func() {
			for i := 0; i < n; i++ {
				err := s.Write(i, strconv.Itoa(i))
				So(err, ShouldBeNil)
			}

			for i := 0; i < n; i++ {
				val, err := s.Read(i)
				So(err, ShouldBeNil)
				j, _ := strconv.Atoi(val)
				So(j, ShouldEqual, i)
			}

			Convey("and shall return error if key does not exist", func() {
				_, err := s.Read(n + 1)
				So(err, ShouldResemble, InvalidKeyError())
			})
		})
	})
}
