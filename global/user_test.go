package global

import (
	"strconv"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRefreshToken(t *testing.T) {
	Convey("should set/get refresh tokens", t, func() {
		user := &User{
			Tokens:        []string{},
			RefreshTokens: map[string]*RefreshToken{},
		}
		expired := time.Now().UTC().AddDate(1, 0, 0) // 1 year
		id := user.AddRefreshToken("agent", "1.1.1.1", expired)
		So(user.Tokens[0], ShouldEqual, id)

		//same agent should reuse token
		id2 := user.AddRefreshToken("agent", "1.1.1.1", expired)
		So(len(user.Tokens), ShouldEqual, 1)
		So(id2, ShouldEqual, id)

		//same agent should reuse token
		token := user.GetRefreshTokenByID(id)
		So(token, ShouldNotBeNil)
		So(token.IP, ShouldEqual, "1.1.1.1")
	})
}

func TestExpiredRefreshToken(t *testing.T) {
	Convey("should check expired on get", t, func() {
		user := &User{
			Tokens:        []string{},
			RefreshTokens: map[string]*RefreshToken{},
		}
		expired := time.Now().UTC().Add(-10 * time.Second)

		id := user.AddRefreshToken("agent", "1.1.1.1", expired)
		So(len(user.Tokens), ShouldEqual, 1)

		//token expire will be remove automatically
		token := user.GetRefreshTokenByID(id)
		So(token, ShouldBeNil)
		So(len(user.Tokens), ShouldEqual, 0)

		expired = time.Now().UTC().Add(-10 * time.Second)
		id = user.AddRefreshToken("agent", "1.1.1.1", expired)
		So(len(user.Tokens), ShouldEqual, 1)

		//token expire will be remove automatically
		token, id = user.GetRefreshTokenByAgent("agent")
		So(token, ShouldBeNil)
		So(len(user.Tokens), ShouldEqual, 0)
	})
}

func TestOnlyKeep10RefreshToken(t *testing.T) {
	Convey("should only keep 10 refresh token", t, func() {
		user := &User{
			Tokens:        []string{},
			RefreshTokens: map[string]*RefreshToken{},
		}
		expired := time.Now().UTC().Add(100 * time.Second)

		for i := 0; i < 15; i++ {
			id := user.AddRefreshToken("agent"+strconv.Itoa(i), "1.1.1.1", expired)
			So(id, ShouldNotBeEmpty)
		}

		So(len(user.Tokens), ShouldEqual, 10)
		So(len(user.RefreshTokens), ShouldEqual, 10)

		for _, id := range user.Tokens {
			token := user.GetRefreshTokenByID(id)
			So(token, ShouldNotBeNil)
		}
	})
}
