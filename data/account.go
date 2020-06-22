package account

import (
	data "github.com/piyuo/libsrv/data"
)

// Account is user account
type Account struct {
	data.DBObject
	Email     string
	FirstName string
	LastName  string
}

// AccountFactory provide function to create instance
var AccountFactory = func() data.Object {
	return new(Account)
}

// Class is object database name
func (c *Account) Class() string {
	return "Account"
}
