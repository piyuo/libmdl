package global

import (
	"context"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/piyuo/libmdl/src/token"
	"github.com/piyuo/libsrv/src/env"
	"github.com/stretchr/testify/assert"
)

func TestIsEmailCanOpenAccount(t *testing.T) {
	assert := assert.New(t)

	ctx := context.Background()
	g, err := New(ctx)
	assert.Nil(err)
	defer g.Close()

	//not taken
	taken, err := g.IsEmailCanOpenAccount(ctx, "access@taken.email")
	assert.Nil(err)
	assert.False(taken)

	//add user
	user := &User{
		Type:  UserTypeOwner,
		Email: "access@taken.email",
	}
	g.UserTable().Set(ctx, user)
	defer g.UserTable().DeleteObject(ctx, user)

	//taken
	taken, err = g.IsEmailCanOpenAccount(ctx, "access@taken.email")
	assert.Nil(err)
	assert.True(taken)
}

func TestIsEmailExist(t *testing.T) {
	assert := assert.New(t)

	ctx := context.Background()
	g, err := New(ctx)
	assert.Nil(err)
	defer g.Close()

	//not taken
	taken, err := g.IsEmailExist(ctx, "email@exist.com")
	assert.Nil(err)
	assert.False(taken)

	//add user
	user := &User{
		Email: "email@exist.com",
	}
	g.UserTable().Set(ctx, user)
	defer g.UserTable().DeleteObject(ctx, user)

	//taken
	taken, err = g.IsEmailExist(ctx, "email@exist.com")
	assert.Nil(err)
	assert.True(taken)
}

func TestGetUserByRefreshToken(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	g, err := New(ctx)
	assert.Nil(err)
	defer g.Close()

	// not exist
	user, err := g.GetUserByRefreshToken(ctx, "notExist", "notExist")
	assert.Nil(err)
	assert.Nil(user)

	//add user & refreshToken
	user = &User{
		Status:        UserStatusActive,
		Tokens:        []string{},
		RefreshTokens: map[string]*RefreshToken{},
		Logins:        []*Login{},
	}
	refreshTokenID := user.AddRefreshToken("agent", "::1", time.Now().UTC().Add(10*time.Minute))
	refreshExpired := token.DefaultRefreshTokenExpired()
	refreshToken, _, err := token.WriteRefreshToken(user.ID, refreshTokenID, refreshExpired)
	assert.NotEmpty(refreshToken)
	g.UserTable().Set(ctx, user)
	defer g.UserTable().DeleteObject(ctx, user)

	// found
	user, err = g.GetUserByRefreshToken(ctx, user.ID, refreshTokenID)
	assert.Nil(err)
	assert.NotNil(user)
}

func TestGetUserByID(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	g, err := New(ctx)
	assert.Nil(err)
	defer g.Close()

	// not exist
	user, err := g.GetUserByID(ctx, "notExist")
	assert.Nil(err)
	assert.Nil(user)

	//add user
	user = &User{}
	g.UserTable().Set(ctx, user)
	defer g.UserTable().DeleteObject(ctx, user)

	// found
	user, err = g.GetUserByID(ctx, user.ID)
	assert.Nil(err)
	assert.NotNil(user)
}

func TestGetUserByEmail(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	g, err := New(ctx)
	assert.Nil(err)
	defer g.Close()

	// not exist
	user, err := g.GetUserByEmail(ctx, "not@exist.mail")
	assert.Nil(err)
	assert.Nil(user)

	//add user
	user = &User{
		Email: "get@user.byEmail",
	}
	g.UserTable().Set(ctx, user)
	defer g.UserTable().DeleteObject(ctx, user)

	user, err = g.GetUserByEmail(ctx, "get@user.byEmail")
	assert.Nil(err)
	assert.NotNil(user)
}

func TestLogins(t *testing.T) {
	assert := assert.New(t)

	user := &User{
		Logins: []*Login{},
	}
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "[::1]:80"
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/546.10 (KHTML, like Gecko) Version/6.0 Mobile/7E18WD Safari/8536.25")
	ctx := env.SetRequest(context.Background(), req)

	user.AddLogin(ctx)
	assert.Equal("::1", user.Logins[0].IP)
	assert.NotEmpty(user.Logins[0].Agent)
	assert.False(user.Logins[0].When.IsZero())

	for i := 0; i < 12; i++ {
		req.RemoteAddr = "[::" + strconv.Itoa(i) + "]:80"
		user.AddLogin(ctx)
	}
	assert.Equal(10, len(user.Logins))
	assert.Equal("::11", user.Logins[0].IP) // latest record first
}

func TestRefreshTokens(t *testing.T) {
	assert := assert.New(t)
	user := &User{
		Tokens:        []string{},
		RefreshTokens: map[string]*RefreshToken{},
	}
	expired := time.Now().UTC().AddDate(1, 0, 0) // 1 year
	id := user.AddRefreshToken("agent", "1.1.1.1", expired)
	assert.Equal(id, user.Tokens[0])

	//same agent should reuse token
	id2 := user.AddRefreshToken("agent", "1.1.1.1", expired)
	assert.Equal(1, len(user.Tokens))
	assert.Equal(id, id2)

	//same agent should reuse token
	token := user.GetRefreshTokenByID(id)
	assert.NotNil(token)
	assert.Equal("1.1.1.1", token.IP)
}

func TestExpiredRefreshToken(t *testing.T) {
	assert := assert.New(t)
	user := &User{
		Tokens:        []string{},
		RefreshTokens: map[string]*RefreshToken{},
	}
	expired := time.Now().UTC().Add(-10 * time.Second)

	id := user.AddRefreshToken("agent", "1.1.1.1", expired)
	assert.Equal(1, len(user.Tokens))

	//token expire will be remove automatically
	token := user.GetRefreshTokenByID(id)
	assert.Nil(token)
	assert.Equal(0, len(user.Tokens))

	expired = time.Now().UTC().Add(-10 * time.Second)
	id = user.AddRefreshToken("agent", "1.1.1.1", expired)
	assert.Equal(1, len(user.Tokens))

	//token expire will be remove automatically
	token, id = user.GetRefreshTokenByAgent("agent")
	assert.Nil(token)
	assert.Equal(0, len(user.Tokens))
}

func TestOnlyKeep10RefreshToken(t *testing.T) {
	assert := assert.New(t)
	user := &User{
		Tokens:        []string{},
		RefreshTokens: map[string]*RefreshToken{},
	}
	expired := time.Now().UTC().Add(100 * time.Second)

	for i := 0; i < 15; i++ {
		id := user.AddRefreshToken("agent"+strconv.Itoa(i), "1.1.1.1", expired)
		assert.NotEmpty(id)
	}

	assert.Equal(10, len(user.Tokens))
	assert.Equal(10, len(user.RefreshTokens))

	for _, id := range user.Tokens {
		token := user.GetRefreshTokenByID(id)
		assert.NotNil(token)
	}
}
