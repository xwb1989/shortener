package storage

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
)

type mockEncoder struct {
	cnt uint64
}

func (encoder *mockEncoder) Encode(s string) uint64 {
	ret := encoder.cnt
	encoder.cnt++
	return ret
}

func (*mockEncoder) StringToKey(s string) (uint64, error) {
	i, _ := strconv.Atoi(s)
	return uint64(i), nil
}

func (*mockEncoder) KeyToString(i uint64) string {
	return strconv.FormatUint(i, 10)
}

func TestStorage(t *testing.T) {
	Convey("Simple memory storage should be able to...", t, func() {
		s := NewMemMap(&mockEncoder{})
		n := 100
		Convey("should be able to write then read", func() {
			for i := 0; i < n; i++ {
				key, err := s.Write(fmt.Sprintf("url%d", i))
				So(err, ShouldBeNil)
				So(key, ShouldEqual, fmt.Sprint(i))
			}

			for i := 0; i < n; i++ {
				val, err := s.Read(fmt.Sprintf("%d", i))
				So(err, ShouldBeNil)
				So(val, ShouldEqual, fmt.Sprintf("url%d", i))
			}

			Convey("and shall return error if key does not exist", func() {
				key := fmt.Sprintf("%d", n+1)
				val, err := s.Read(key)
				So(val, ShouldBeEmpty)
				So(err, ShouldResemble, InvalidKeyError(key))
			})
		})
	})
}
