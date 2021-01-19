package token

import (
	"context"
	"strconv"
	"time"

	"github.com/piyuo/libsrv/env"
	"github.com/piyuo/libsrv/token"
	"github.com/pkg/errors"
)

// KeyAccountID is where the account id locate in token
//
const KeyAccountID = "a"

// KeyUserID is where the user id locate in token
//
const KeyUserID = "u"

// KeyExtendCount is how many time this access token has been extend
//
const KeyExtendCount = "e"

// KeyRefreshTokenID is where the refresh token id locate in token
//
const KeyRefreshTokenID = "r"

// AccessTokenDuration define duration of AccessToken
//
const AccessTokenDuration = 60 // 60 min

// RefreshTokenDuration define a 10 year refresh token
//
const RefreshTokenDuration = 10 // 10 year

// DefaultAccessTokenExpired return default access token expired time, default is 60 minutes
//
//	refreshExpired := DefaultAccessTokenExpired()// 60 minutes
//
func DefaultAccessTokenExpired() time.Time {
	return time.Now().UTC().Add(AccessTokenDuration * time.Minute)
}

// WriteAccessToken create token string to client side
//
// 	accessExpired := DefaultAccessTokenExpired()
//	tokenStr,expired,err := WriteAccessToken(ctx,accountID,userID,0,accessExpired)
//
func WriteAccessToken(ctx context.Context, accountID, userID string, extendCount int, expired time.Time) (string, time.Time, error) {
	accessToken := token.NewToken()
	accessToken.Set(KeyAccountID, accountID)
	accessToken.Set(KeyUserID, userID)
	accessToken.Set(KeyExtendCount, strconv.Itoa(extendCount))

	token, err := accessToken.ToString(expired)
	if err != nil {
		return "", time.Time{}, errors.Wrap(err, "failed to create access token for user: "+userID)
	}
	return token, expired, nil
}

// ReadAccessToken return account id and user id from string, set current context user id from user id, extendCount is how many time this access token has been extend
//
//	ctx,accountID,userID,isExpired,extendCount,err := ReadAccessToken(ctx,"token")
//
func ReadAccessToken(ctx context.Context, crypted string) (context.Context, string, string, bool, int, error) {
	accessToken, isExpired, err := token.FromString(crypted)
	if err != nil {
		return ctx, "", "", false, 0, err
	}
	if isExpired {
		return ctx, "", "", true, 0, nil
	}

	accountID := accessToken.Get(KeyAccountID)
	userID := accessToken.Get(KeyUserID)
	iExtendCount := accessToken.Get(KeyExtendCount)
	extendCount, err := strconv.Atoi(iExtendCount)
	if err != nil {
		return ctx, "", "", false, 0, err
	}
	ctx = env.SetUserID(ctx, userID)
	ctx = env.SetAccountID(ctx, accountID)
	return ctx, accountID, userID, false, extendCount, nil
}

// DefaultRefreshTokenExpired return default refresh token expired time, default is 10 years
//
//	refreshExpired := DefaultRefreshTokenExpired()// 10 year
//
func DefaultRefreshTokenExpired() time.Time {
	return time.Now().UTC().AddDate(RefreshTokenDuration, 0, 0)
}

// WriteRefreshToken create refresh token string from user id and refresh token id
//
//	refreshExpired := DefaultRefreshTokenExpired()
//	tokenStr,expired,err := WriteRefreshToken(userID,refreshTokenID,refreshExpired)
//
func WriteRefreshToken(userID, refreshTokenID string, expired time.Time) (string, time.Time, error) {
	refreshToken := token.NewToken()
	refreshToken.Set(KeyUserID, userID)
	refreshToken.Set(KeyRefreshTokenID, refreshTokenID)

	token, err := refreshToken.ToString(expired)
	if err != nil {
		return "", time.Time{}, errors.Wrap(err, "failed to create refresh token for user: "+userID)
	}
	return token, expired, nil
}

// ReadRefreshToken return user id, refresh token id from string
//
//	userID,refreshTokenID,isExpired,err := ReadRefreshToken(ctx,"token")
//
func ReadRefreshToken(crypted string) (string, string, bool, error) {
	refreshToken, isExpired, err := token.FromString(crypted)
	if err != nil {
		return "", "", false, err
	}
	if isExpired {
		return "", "", true, nil
	}
	userID := refreshToken.Get(KeyUserID)
	refreshTokenID := refreshToken.Get(KeyRefreshTokenID)
	return userID, refreshTokenID, false, nil
}
