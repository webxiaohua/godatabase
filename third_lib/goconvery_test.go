package main

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAssertion(t *testing.T){
	convey.Convey("test",t, func() {
		convey.So(1,convey.ShouldEqual,2)
	})
}
