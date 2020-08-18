package global

import (
	"strings"
	"time"

	"github.com/piyuo/libsrv/data"
	"github.com/piyuo/libsrv/identifier"
)


// User represent single user, ID is serial id to keep it short
//
type User struct {
	data.BaseObject

	// AccountID indicate user belong to which account
	//
	AccountID string

	// Region meas user belong to which data center
	//
	Region string

	// Email is unique in User table, user need use email to login to their store
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

	// LocationID mean user belong to which location
	//
	LocationID string

	// Roles is user belong which policy role
	//
	Roles []string

	// Token keep all refresh token id for search
	//
	Tokens []string

	// RefreshTokens keep issued RefreshToken
	//
	RefreshTokens map[string]*RefreshToken

	// Logins latest 5 login record
	//
	Logins map[string]string

	// Policy is Casbin Policy
	//
	StorePolicy string

	// Roles keep custom roles
	//
	StoreRoles map[string]string
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

	// Expired time
	//
	Agent time.Time
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
