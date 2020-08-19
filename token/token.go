package token

import (
	"context"
	"time"

	"github.com/piyuo/libsrv/session"
	"github.com/piyuo/libsrv/token"
)

// KeyAccountID is where the account id locate in token
//
const KeyAccountID = "a"

// KeyUserID is where the user id locate in token
//
const KeyUserID = "u"

// KeyRefreshTokenID is where the refresh token id locate in token
//
const KeyRefreshTokenID = "r"

// AccessTokenDuration define duration of AccessToken
//
const AccessTokenDuration = 60 // 60 min

// RefreshTokenDuration define a 10 year refresh token
//
const RefreshTokenDuration = 10 // 10 year

// WriteAccessToken create token string to client side
//
//	tokenStr,err := WriteAccessToken(ctx,accountID,userID)
//
func WriteAccessToken(ctx context.Context, accountID, userID string) (string, error) {
	accessToken := token.NewToken()
	accessToken.Set(KeyAccountID, accountID)
	accessToken.Set(KeyUserID, userID)
	expired := time.Now().UTC().Add(AccessTokenDuration * time.Minute)
	return accessToken.ToString(expired)
}

// ReadAccessToken return account id and user id from string, set current context user id from user id
//
//	ctx,accountID,userID,isExpired,err := ReadAccessToken(ctx,"token")
//
func ReadAccessToken(ctx context.Context, crypted string) (context.Context, string, string, bool, error) {
	accessToken, isExpired, err := token.FromString(crypted)
	if err != nil {
		return ctx, "", "", false, err
	}
	if isExpired {
		return ctx, "", "", true, nil
	}
	accountID := accessToken.Get(KeyAccountID)
	userID := accessToken.Get(KeyUserID)
	ctx = session.SetUserID(ctx, userID)
	return ctx, accountID, userID, false, nil
}

// WriteRefreshToken create refresh token string from user id and refresh token id
//
//	tokenStr,err := WriteRefreshToken(userID,refreshTokenID)
//
func WriteRefreshToken(userID, refreshTokenID string) (string, error) {
	refreshToken := token.NewToken()
	refreshToken.Set(KeyUserID, userID)
	refreshToken.Set(KeyRefreshTokenID, refreshTokenID)
	expired := time.Now().UTC().AddDate(RefreshTokenDuration, 0, 0) // 10 year
	return refreshToken.ToString(expired)
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
