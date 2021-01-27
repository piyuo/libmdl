package global

import (
	"context"

	"github.com/piyuo/libsrv/data"
	"github.com/pkg/errors"
)

// DomainName keep all registered domain name
//
type DomainName struct {
	data.BaseObject
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
func (c *Global) CreateDomainName(ctx context.Context, domainName string) error {
	d := &DomainName{}
	d.SetID(domainName)
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
	return c.DomainNameTable().Exist(ctx, domainName)
}

// RemoveDomainName remove domain name
//
//	err := RemoveDomainName(ctx,"a@b.c")
//
func (c *Global) RemoveDomainName(ctx context.Context, domainName string) error {
	if err := c.DomainNameTable().Delete(ctx, domainName); err != nil {
		return errors.Wrap(err, "failed to delete data")
	}
	return nil
}
