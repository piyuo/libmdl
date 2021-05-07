package global

import (
	"context"
	"strings"
	"time"

	"github.com/piyuo/libsrv/db"
	"github.com/piyuo/libsrv/env"
	"github.com/piyuo/libsrv/identifier"
	"github.com/pkg/errors"
)

// User represent single user, ID is serial id to keep it short
//
type User struct {
	db.Model

	// Status user status
	//
	Suspend bool `firestore:"Suspend,omitempty"`

	// Email is unique in User table, user need use email to login to their store,when user choose to leave this account permanently. Email will be empty so another account can create new user with this Email
	//
	Email string `firestore:"Email,omitempty"`

	// BackupEmail used when user can't access their email service, they can choose send email to BackupEmail
	//
	BackupEmail string `firestore:"BackupEmail,omitempty"`

	// FirstName is user first name
	//
	FirstName string `firestore:"FirstName,omitempty"`

	// LastName is user last name
	//
	LastName string `firestore:"LastName,omitempty"`

	// Token keep all refresh token id for search
	//
	Tokens []string `firestore:"Tokens,omitempty"`

	// RefreshTokens keep issued RefreshToken
	//
	RefreshTokens map[string]*RefreshToken `firestore:"RefreshTokens,omitempty"`

	// Logins latest 5 login record
	//
	Logins []*Login `firestore:"Logins,omitempty"`

	// Type is user type in system, like UserTypeAdministrator
	//
	Type UserType `firestore:"Type,omitempty"`

	// StoreRoles is a map define user role in store
	//
	//  StoreRoles["storeID1"]=["ManagerID"]
	//  StoreRoles["storeID2"]=["ReaderID"]
	//
	StoreRoles map[string]int32 `firestore:"StoreRoles,omitempty"`

	// LocationRoles is a map define user role in location
	//
	//  LocationRoles["locationID1"]=["ManagerID"]
	//  LocationRoles["locationID2"]=["ReaderID"]
	//
	LocationRoles map[string]int32 `firestore:"LocationRoles,omitempty"`
}

// RefreshToken let user login using refresh token
//
type RefreshToken struct {

	// IP is user ip the token belong to, user can have multiple refresh token in different IP
	//
	IP string `firestore:"IP,omitempty"`

	// Agent is user agent id from request user agent
	//
	Agent string `firestore:"Agent,omitempty"`

	// Expired time
	//
	Expired time.Time `firestore:"Expired,omitempty"`
}

// Login is user login record
//
type Login struct {

	// IP is user ip the token belong to, user can have multiple refresh token in different IP
	//
	IP string `firestore:"IP,omitempty"`

	// When the user perform login
	//
	When time.Time `firestore:"When,omitempty"`

	// Agent is user agent
	//
	Agent string `firestore:"Agent,omitempty"`
}

// Factory create a empty object, return object must be nil safe, no nil in any field
//
func (c *User) Factory() db.Object {
	return &User{
		Tokens:        []string{},
		RefreshTokens: map[string]*RefreshToken{},
		Logins:        []*Login{},
		StoreRoles:    map[string]int32{},
		LocationRoles: map[string]int32{},
	}
}

// Collection return the name in database
//
func (c *User) Collection() string {
	return "User"
}

// GetRefreshTokenByID return refresh token by id, return nil if not found
//
func (c *User) GetRefreshTokenByID(id string) *RefreshToken {
	token := c.RefreshTokens[id]
	if time.Now().UTC().Before(token.Expired) {
		return token
	}
	c.RemoveRefreshToken(id)
	return nil
}

// GetRefreshTokenByAgent return refresh token by agent, return nil if not found
//
func (c *User) GetRefreshTokenByAgent(agent string) (*RefreshToken, string) {
	a := strings.ToLower(agent)
	for id, token := range c.RefreshTokens {
		tokenAgent := strings.ToLower(token.Agent)
		if a == tokenAgent {
			if time.Now().UTC().Before(token.Expired) {
				return token, id
			}
			c.RemoveRefreshToken(id)
			return nil, ""
		}
	}
	return nil, ""
}

// AddRefreshToken add new refresh token, return token id
//
func (c *User) AddRefreshToken(agent, ip string, expired time.Time) string {

	token, id := c.GetRefreshTokenByAgent(agent)
	if token != nil {
		token.IP = ip
		token.Expired = expired
		return id
	}

	id = identifier.UUID()
	token = &RefreshToken{
		Agent:   agent,
		IP:      ip,
		Expired: expired,
	}
	c.RefreshTokens[id] = token
	c.Tokens = append(c.Tokens, id)
	c.CleanRefreshToken()
	return id
}

// RemoveRefreshToken remove refresh token by id
//
func (c *User) RemoveRefreshToken(id string) {
	delete(c.RefreshTokens, id)
	for i, ele := range c.Tokens {
		if ele == id {
			copy(c.Tokens[i:], c.Tokens[i+1:])    // Shift a[i+1:] left one index.
			c.Tokens[len(c.Tokens)-1] = ""        // Erase last element (write zero value).
			c.Tokens = c.Tokens[:len(c.Tokens)-1] // Truncate slice
			return
		}
	}
}

// AddLogin add new login record, only keep 5 record
//
func (c *User) AddLogin(ctx context.Context) {
	login := &Login{
		Agent: env.GetUserAgentID(ctx),
		IP:    env.GetIP(ctx),
		When:  time.Now().UTC(),
	}
	c.Logins = insertLogin(c.Logins, login, 0)
	if len(c.Logins) > 10 {
		c.Logins[len(c.Logins)-1] = nil       // Erase last element (write zero value).
		c.Logins = c.Logins[:len(c.Logins)-1] // Truncate slice
	}
}

func insertLogin(a []*Login, c *Login, i int) []*Login {
	return append(a[:i], append([]*Login{c}, a[i:]...)...)
}

// CleanRefreshToken keep only 10 token
//
func (c *User) CleanRefreshToken() {
	removeCount := 0
	tokenCount := len(c.Tokens)
	if tokenCount > 10 {
		removeCount = tokenCount - 10

		for i := 0; i < removeCount; i++ {
			delete(c.RefreshTokens, c.Tokens[0])

			copy(c.Tokens[0:], c.Tokens[1:])      // Shift a[i+1:] left one index.
			c.Tokens[len(c.Tokens)-1] = ""        // Erase last element (write zero value).
			c.Tokens = c.Tokens[:len(c.Tokens)-1] // Truncate slice
		}
	}
}

// IsEmailExist return true if email can be create account
//
//	found, err := IsEmailExist(ctx, "a@b.c")
//
func IsEmailExist(ctx context.Context, email string) (bool, error) {
	client, err := Client(ctx)
	if err != nil {
		return false, err
	}
	return client.Query(&User{}).Where("Email", "==", email).ReturnExists(ctx)
}

// GetUserByRefreshToken get user from refresh token that login need
//
func GetUserByRefreshToken(ctx context.Context, userID, refreshTokenID string) (*User, error) {
	client, err := Client(ctx)
	if err != nil {
		return nil, err
	}
	iUser, err := client.Query(&User{}).Where("ID", "==", userID).Where("Tokens", "array-contains", refreshTokenID).ReturnFirst(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "get user %v token %v", userID, refreshTokenID)
	}
	if iUser == nil {
		return nil, nil // possible user already removed
	}
	return iUser.(*User), nil
}

// GetUserByID get user from id that login need
//
func GetUserByID(ctx context.Context, userID string) (*User, error) {
	client, err := Client(ctx)
	if err != nil {
		return nil, err
	}

	iUser, err := client.Get(ctx, &User{}, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "get user %v", userID)
	}
	if iUser == nil {
		return nil, nil // possible user already removed
	}
	return iUser.(*User), nil
}

// GetUserByEmail get user from email that login need
//
func GetUserByEmail(ctx context.Context, email string) (*User, error) {
	client, err := Client(ctx)
	if err != nil {
		return nil, err
	}

	iUser, err := client.Query(&User{}).Where("Email", "==", email).ReturnFirst(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "get user %v", email)
	}
	if iUser == nil {
		return nil, nil // possible user already removed
	}
	return iUser.(*User), nil
}
