package token

import (
	"context"
	"testing"
	"time"

	"github.com/piyuo/libsrv/session"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAccessToken(t *testing.T) {
	Convey("should read and write access token", t, func() {
		ctx := context.Background()
		accessExpired := time.Now().UTC().Add(AccessTokenDuration * time.Minute)
		tokenStr, expired, err := WriteAccessToken(ctx, "account1", "user1", 1, accessExpired)
		So(err, ShouldBeNil)
		So(expired.IsZero(), ShouldBeFalse)
		So(tokenStr, ShouldNotBeEmpty)

		ctx, accountID, userID, isExpired, extendCount, err := ReadAccessToken(ctx, tokenStr)
		So(err, ShouldBeNil)
		So(session.GetUserID(ctx), ShouldEqual, "user1")
		So(accountID, ShouldEqual, "account1")
		So(userID, ShouldEqual, "user1")
		So(isExpired, ShouldBeFalse)
		So(extendCount, ShouldEqual, 1)

		//test invalid token
		ctx = session.SetUserID(ctx, "")
		ctx, accountID, userID, isExpired, extendCount, err = ReadAccessToken(ctx, "invalid")
		So(err, ShouldNotBeNil)
		So(session.GetUserID(ctx), ShouldBeEmpty)
		So(accountID, ShouldBeEmpty)
		So(userID, ShouldBeEmpty)
		So(isExpired, ShouldBeFalse)
		So(extendCount, ShouldEqual, 0)
	})

	Convey("should read and write refresh token", t, func() {
		refreshExpired := time.Now().UTC().AddDate(RefreshTokenDuration, 0, 0) // 10 year
		tokenStr, expired, err := WriteRefreshToken("user1", "rt1", refreshExpired)
		So(err, ShouldBeNil)
		So(tokenStr, ShouldNotBeEmpty)
		So(expired.IsZero(), ShouldBeFalse)

		userID, refreshTokenID, isExpired, err := ReadRefreshToken(tokenStr)
		So(err, ShouldBeNil)
		So(userID, ShouldEqual, "user1")
		So(refreshTokenID, ShouldEqual, "rt1")
		So(isExpired, ShouldBeFalse)

		//test invalid token
		userID, refreshTokenID, isExpired, err = ReadRefreshToken("invalid")
		So(err, ShouldNotBeNil)
		So(userID, ShouldBeEmpty)
		So(refreshTokenID, ShouldBeEmpty)
		So(isExpired, ShouldBeFalse)
	})

}
