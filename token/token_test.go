package token

import (
	"context"
	"testing"

	"github.com/piyuo/libsrv/env"
	"github.com/stretchr/testify/assert"
)

func TestReadWriteAccessToken(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	ctx := context.Background()
	accessExpired := DefaultAccessTokenExpired()
	assert.False(accessExpired.IsZero())

	tokenStr, expired, err := WriteAccessToken(ctx, "account1", "user1", accessExpired)
	assert.Nil(err)
	assert.False(expired.IsZero())
	assert.NotEmpty(tokenStr)

	ctx, accountID, userID, isExpired, err := ReadAccessToken(ctx, tokenStr)
	assert.Nil(err)
	assert.Equal("user1", env.GetUserID(ctx))
	assert.Equal("account1", env.GetAccountID(ctx))
	assert.Equal("user1", userID)
	assert.Equal("account1", accountID)
	assert.False(isExpired)

	//test invalid token
	ctx = env.SetUserID(ctx, "")
	ctx = env.SetAccountID(ctx, "")
	ctx, accountID, userID, isExpired, err = ReadAccessToken(ctx, "invalid")
	assert.NotNil(err)
	assert.Empty(env.GetUserID(ctx))
	assert.Empty(env.GetAccountID(ctx))
	assert.Empty(accountID)
	assert.Empty(userID)
	assert.False(isExpired)
}
func TestReadWriteRefreshToken(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	refreshExpired := DefaultRefreshTokenExpired()
	assert.False(refreshExpired.IsZero())

	tokenStr, expired, err := WriteRefreshToken("user1", "rt1", refreshExpired)
	assert.Nil(err)
	assert.NotEmpty(tokenStr)
	assert.False(expired.IsZero())

	userID, refreshTokenID, isExpired, err := ReadRefreshToken(tokenStr)
	assert.Nil(err)
	assert.Equal("user1", userID)
	assert.Equal("rt1", refreshTokenID)
	assert.False(isExpired)

	//test invalid token
	userID, refreshTokenID, isExpired, err = ReadRefreshToken("invalid")
	assert.NotNil(err)
	assert.Empty(userID)
	assert.Empty(refreshTokenID)
	assert.False(isExpired)

}
