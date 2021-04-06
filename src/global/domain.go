package global

import (
	"context"
	"strings"

	"github.com/piyuo/libsrv/src/db"
	"github.com/piyuo/libsrv/src/log"
	"github.com/pkg/errors"
)

// Domain keep all registered domain name
//
type Domain struct {
	db.Model
}

// Factory create a empty object, return object must be nil safe, no nil in any field
//
func (c *Domain) Factory() db.Object {
	return &Domain{}
}

// Collection return the name in database
//
func (c *Domain) Collection() string {
	return "Domain"
}

// CreateDomain create domain name, must run in transaction
//
//	err := CreateDomain(ctx,tx,"a@b.c")
//
func CreateDomain(ctx context.Context, tx db.Transaction, domainName, accountID string) error {
	d := &Domain{}
	d.SetAccountID(accountID)
	d.SetID(strings.ToLower(domainName))
	if err := tx.Set(ctx, d); err != nil {
		return errors.Wrap(err, "new domain")
	}
	return nil
}

// IsDomainTaken return true if domain already taken
//
//	taken, err := IsDomainTaken(ctx, "a@b.c")
//
func IsDomainTaken(ctx context.Context, domainName string) (bool, error) {
	client, err := Client(ctx)
	if err != nil {
		return false, err
	}
	return client.Exists(ctx, &Domain{}, strings.ToLower(domainName))
}

// DeleteDomain remove domain name
//
//	err := DeleteDomain(ctx,"a@b.c")
//
func DeleteDomain(ctx context.Context, domainName string) error {
	client, err := Client(ctx)
	if err != nil {
		return err
	}
	d := &Domain{}
	d.SetID(strings.ToLower(domainName))
	if err := client.Delete(ctx, d); err != nil {
		return errors.Wrap(err, "delete domain")
	}
	return nil
}

// DeleteDomainByAccountID remove domain name by accountID
//
//	err := DeleteDomainByAccountID(ctx,"accountID")
//
func DeleteDomainByAccountID(ctx context.Context, accountID string) error {
	client, err := Client(ctx)
	if err != nil {
		return err
	}
	done, numDeleted, err := client.Query(&Domain{}).Where("AccountID", "==", accountID).Delete(ctx, 100)
	if done {
		log.Info(ctx, "del %v domain ", numDeleted)
		return err
	}
	log.Warn(ctx, "del domain not done, only del %v", numDeleted)
	return err
}
