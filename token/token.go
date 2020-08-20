package token

import (
	"context"
	"strconv"
	"time"

	"github.com/piyuo/libsrv/session"
	"github.com/piyuo/libsrv/token"
	"github.com/pkg/errors"
)

// KeyAccountID is where the account id locate in token
//
const KeyAccountID = "a"

// KeyUserID is where the user id locate in token
//
const KeyUserID = "u"

// KeyAccessIP is the IP when issue access token
//
const KeyAccessIP = "i"

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

// WriteAccessToken create token string to client side
//
//	tokenStr,err := WriteAccessToken(ctx,accountID,userID)
//
func WriteAccessToken(ctx context.Context, accountID, userID string, extendCount int) (string, time.Time, error) {
	accessToken := token.NewToken()
	accessToken.Set(KeyAccountID, accountID)
	accessToken.Set(KeyUserID, userID)
	accessToken.Set(KeyAccessIP, session.GetIP(ctx))
	accessToken.Set(KeyExtendCount, strconv.Itoa(extendCount))
	expired := time.Now().UTC().Add(AccessTokenDuration * time.Minute)
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
	accessIP := accessToken.Get(KeyAccessIP)
	if accessIP != session.GetIP(ctx) { // if ip has been changed. assess token is expired
		return ctx, "", "", true, 0, nil
	}

	accountID := accessToken.Get(KeyAccountID)
	userID := accessToken.Get(KeyUserID)
	iExtendCount := accessToken.Get(KeyExtendCount)
	extendCount, err := strconv.Atoi(iExtendCount)
	if err != nil {
		return ctx, "", "", false, 0, err
	}
	ctx = session.SetUserID(ctx, userID)
	return ctx, accountID, userID, false, extendCount, nil
}

// WriteRefreshToken create refresh token string from user id and refresh token id
//
//	tokenStr,expired,err := WriteRefreshToken(userID,refreshTokenID)
//
func WriteRefreshToken(userID, refreshTokenID string) (string, time.Time, error) {
	refreshToken := token.NewToken()
	refreshToken.Set(KeyUserID, userID)
	refreshToken.Set(KeyRefreshTokenID, refreshTokenID)
	expired := time.Now().UTC().AddDate(RefreshTokenDuration, 0, 0) // 10 year

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
