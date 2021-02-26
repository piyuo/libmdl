package global

import (
	"context"
	"fmt"
	"strings"

	"github.com/piyuo/libsrv/src/data"
	"github.com/pkg/errors"
)

// DomainName keep all registered domain name
//
type DomainName struct {
	data.DomainObject
}

// DomainNameTable return DomainName table
//
//	table := db.UserTable()
//
func (c *Global) DomainNameTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "DomainName",
		Factory: func() data.Object {
			return &DomainName{}
		},
	}
}

// RemoveAllDomainName remove all domain name
//
//	err := RemoveAllDomainName(ctx)
//
func (c *Global) RemoveAllDomainName(ctx context.Context) error {
	return c.DomainNameTable().Clear(ctx)
}

// CreateDomainName create domain name
//
//	err := CreateDomainName(ctx,"a@b.c")
//
func (c *Global) CreateDomainName(ctx context.Context, domainName, accountID string) error {
	d := &DomainName{}
	d.SetAccountID(accountID)
	d.SetID(strings.ToLower(domainName))
	if err := c.DomainNameTable().Set(ctx, d); err != nil {
		return errors.Wrap(err, "failed to set data")
	}
	return nil
}

// IsDomainNameTaken return true if domain already taken
//
//	taken, err := IsDomainNameTaken(ctx, "a@b.c")
//
func (c *Global) IsDomainNameTaken(ctx context.Context, domainName string) (bool, error) {
	return c.DomainNameTable().IsExists(ctx, strings.ToLower(domainName))
}

// RemoveDomainName remove domain name
//
//	err := RemoveDomainName(ctx,"a@b.c")
//
func (c *Global) RemoveDomainName(ctx context.Context, domainName string) error {
	if err := c.DomainNameTable().Delete(ctx, strings.ToLower(domainName)); err != nil {
		return errors.Wrap(err, "failed to delete data")
	}
	return nil
}

// RemoveDomainNameByAccountID remove domain name by accountID
//
//	err := RemoveDomainName(ctx,"accountID")
//
func (c *Global) RemoveDomainNameByAccountID(ctx context.Context, accountID string) error {
	count, err := c.DomainNameTable().Query().Where("AccountID", "==", accountID).Clear(ctx)
	if count > 0 {
		fmt.Printf("remove %v DomainName\n", count)
	}
	return err
}
