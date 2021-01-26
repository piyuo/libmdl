package global

import (
	"context"
	"strings"
	"time"

	"github.com/piyuo/libsrv/data"
	"github.com/piyuo/libsrv/env"
	"github.com/piyuo/libsrv/identifier"
	"github.com/pkg/errors"
)

// User represent single user, ID is serial id to keep it short
//
type User struct {
	data.DomainObject

	// Status user status
	//
	Status UserStatus

	// Email is unique in User table, user need use email to login to their store,when user choose to leave this account permanently. Email will be empty so another account can create new user with this Email
	//
	Email string

	// BackupEmail used when user can't access their email service, they can choose send email to BackupEmail
	//
	BackupEmail string

	// FirstName is user first name
	//
	FirstName string

	// LastName is user last name
	//
	LastName string

	// Token keep all refresh token id for search
	//
	Tokens []string

	// RefreshTokens keep issued RefreshToken
	//
	RefreshTokens map[string]*RefreshToken

	// Logins latest 5 login record
	//
	Logins []*Login

	// Type is user type in system, like UserTypeAdministrator
	//
	Type UserType

	// StoreRoles is a map define user role in store
	//
	//  StoreRoles["storeID1"]=["ManagerID"]
	//  StoreRoles["storeID2"]=["ReaderID"]
	//
	StoreRoles map[string]int32

	// LocationRoles is a map define user role in location
	//
	//  LocationRoles["locationID1"]=["ManagerID"]
	//  LocationRoles["locationID2"]=["ReaderID"]
	//
	LocationRoles map[string]int32
}

// UserTable return user table
//
//	table := db.UserTable()
//
func (c *Global) UserTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "User",
		Factory: func() data.Object {
			return &User{}
		},
	}
}

// RemoveAllUser remove all user
//
//	err := RemoveAllUser(ctx)
//
func (c *Global) RemoveAllUser(ctx context.Context) error {
	return c.UserTable().Clear(ctx)
}

// IsEmailTaken return true if email registered in global database
//
//	registered, err := IsEmailTaken(ctx, g, "a@b.c")
//
func (c *Global) IsEmailTaken(ctx context.Context, email string) (bool, error) {
	found, err := c.UserTable().Query().Where("Email", "==", email).Where("Type", "==", UserTypeOwner).IsExist(ctx)
	if err != nil {
		return false, errors.Wrap(err, "failed to find user by email: "+email)
	}
	if found {
		return true, nil
	}
	return false, nil
}

// GetUserByRefreshToken get user from refresh token that login need
//
func (c *Global) GetUserByRefreshToken(ctx context.Context, userID, refreshTokenID string) (*User, error) {
	iUser, err := c.UserTable().Query().Where("ID", "==", userID).Where("Tokens", "array-contains", refreshTokenID).GetFirstObject(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user, userID: "+userID+", refreshTokenID:"+refreshTokenID)
	}
	if iUser == nil {
		return nil, nil // possible user already removed
	}
	return iUser.(*User), nil
}

// GetUserByID get user from id that login need
//
func (c *Global) GetUserByID(ctx context.Context, userID string) (*User, error) {
	iUser, err := c.UserTable().Get(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by id: "+userID)
	}
	if iUser == nil {
		return nil, nil // possible user already removed
	}
	return iUser.(*User), nil
}

// GetUserByEmail get user from email that login need
//
func (c *Global) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	iUser, err := c.UserTable().Find(ctx, "Email", "==", email)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by email: "+email)
	}
	if iUser == nil {
		return nil, nil // possible user already removed
	}
	return iUser.(*User), nil
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

// RefreshToken let user login using refresh token
//
type RefreshToken struct {

	// IP is user ip the token belong to, user can have multiple refresh token in different IP
	//
	IP string

	// Agent is user agent id from request user agent
	//
	Agent string

	// Expired time
	//
	Expired time.Time
}

// Login is user login record
//
type Login struct {

	// IP is user ip the token belong to, user can have multiple refresh token in different IP
	//
	IP string

	// When the user perform login
	//
	When time.Time

	// Agent is user agent
	//
	Agent string
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
