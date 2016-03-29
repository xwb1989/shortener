package encoder

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"strconv"
	"testing"
)

func TestIncrementer(t *testing.T) {
	Convey("incrementer should return incremental key for each call", t, func() {
		encoder := NewIncrementalEncoder(0)
		var i uint64
		for i = math.MaxUint64 - 1000; i < math.MaxUint64; i++ {
			key := encoder.Encode("a")
			So(encoder.KeyToString(key), ShouldEqual, strconv.FormatUint(key, 36))
			So(encoder.StringToKey(encoder.KeyToString(key)), ShouldEqual, key)
		}
	})
}
