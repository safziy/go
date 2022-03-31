package base_grammar

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDeferReturn1(t *testing.T) {
	Convey("TestDeferReturn1", t, func() {
		So(deferReturn1(), ShouldEqual, 1)
	})
}

func TestDeferReturn2(t *testing.T) {
	Convey("TestDeferReturn2", t, func() {
		So(deferReturn2(), ShouldEqual, 2)
	})
}

func TestDeferReturn3(t *testing.T) {
	Convey("TestDeferReturn3", t, func() {
		So(deferReturn3(), ShouldEqual, 1)
	})
}

func TestDeferReturn4(t *testing.T) {
	Convey("TestDeferReturn4", t, func() {
		So(deferReturn4(), ShouldEqual, 1)
	})
}
