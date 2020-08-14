package acl

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAccessToken(t *testing.T) {
	Convey("should create access token", t, func() {
		token, err := CreateAccessToken("12345678", "12345678")
		So(err, ShouldBeNil)
		So(token, ShouldNotBeEmpty)
	})
}
