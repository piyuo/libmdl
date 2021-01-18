package token

import (
	"context"
	"testing"

	"github.com/piyuo/libsrv/session"
	"github.com/stretchr/testify/assert"
)

func TestReadWriteAccessToken(t *testing.T) {
	assert := assert.New(t)

	ctx := context.Background()
	accessExpired := DefaultAccessTokenExpired()
	assert.False(accessExpired.IsZero())

	tokenStr, expired, err := WriteAccessToken(ctx, "account1", "user1", 1, accessExpired)
	assert.Nil(err)
	assert.False(expired.IsZero())
	assert.NotEmpty(tokenStr)

	ctx, accountID, userID, isExpired, extendCount, err := ReadAccessToken(ctx, tokenStr)
	assert.Nil(err)
	assert.Equal("user1", session.GetUserID(ctx))
	assert.Equal("account1", accountID)
	assert.Equal("user1", userID)
	assert.False(isExpired)
	assert.Equal(1, extendCount)

	//test invalid token
	ctx = session.SetUserID(ctx, "")
	ctx, accountID, userID, isExpired, extendCount, err = ReadAccessToken(ctx, "invalid")
	assert.NotNil(err)
	assert.Empty(session.GetUserID(ctx))
	assert.Empty(accountID)
	assert.Empty(userID)
	assert.False(isExpired)
	assert.Equal(0, extendCount)

}
func TestReadWriteRefreshToken(t *testing.T) {
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
