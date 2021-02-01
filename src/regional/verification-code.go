package regional

import (
	"context"
	"fmt"
	"time"

	"github.com/piyuo/libsrv/src/crypto"
	"github.com/piyuo/libsrv/src/data"
	"github.com/piyuo/libsrv/src/util"
	"github.com/pkg/errors"
)

// VerificationCode keep verification code
//
type VerificationCode struct {
	data.BaseObject

	// Hash is code hash with salt, we do not store code only hash is enough
	//
	Hash uint32

	// Crypted code
	//
	Crypted string
}

// VerificationCodeTable return VerificationCode table
//
//	table := VerificationCodeTable(r)
//
func (c *Regional) VerificationCodeTable() *data.Table {
	return &data.Table{
		Connection: c.Connection,
		TableName:  "VerificationCode",
		Factory: func() data.Object {
			return &VerificationCode{}
		},
	}
}

// CreateVerificationCode create verification code
//
//	err := CreateVerificationCode(ctx,"a@b.c","123456")
//
func (c *Regional) CreateVerificationCode(ctx context.Context, email, code string) error {
	hash := util.StringHash(code)
	crypted, err := crypto.Encrypt(code)
	if err != nil {
		return errors.Wrap(err, "failed to Encrypt code: "+code)
	}

	vc := &VerificationCode{
		BaseObject: data.BaseObject{
			ID: email,
		},
		Hash:    hash,
		Crypted: crypted,
	}
	if err := c.VerificationCodeTable().Set(ctx, vc); err != nil {
		return errors.Wrap(err, "failed to Set verification code")
	}
	return nil
}

// GetVerificationCode get verification code from database for resend, return found,code
//
//	found,code,err := GetVerificationCode(ctx, "a@b.c")
//
func (c *Regional) GetVerificationCode(ctx context.Context, email string) (bool, string, error) {
	obj, err := c.VerificationCodeTable().Get(ctx, email)
	if err != nil {
		return false, "", errors.Wrap(err, "failed to Get data")
	}

	if obj != nil {
		vc := obj.(*VerificationCode)
		code, err := crypto.Decrypt(vc.Crypted)
		if err != nil {
			return false, "", errors.Wrap(err, "failed to Decrypt: "+vc.Crypted)
		}
		return true, code, nil
	}
	return false, "", nil
}

// ConfirmVerificationCode return found and confirm of a verification code
//
//	found,confirm, err := ConfirmVerificationCode(ctx, "a@b.c", "123456")
//
func (c *Regional) ConfirmVerificationCode(ctx context.Context, email, code string) (bool, bool, error) {

	obj, err := c.VerificationCodeTable().Get(ctx, email)
	if err != nil {
		return false, false, errors.Wrap(err, "failed to Get data")
	}
	if obj == nil {
		//verification code not exist, maybe removed after 30 min
		return false, false, nil
	}

	hash := util.StringHash(code)
	vc := obj.(*VerificationCode)
	if vc.Hash != hash {
		//user input code is not match
		return true, false, nil
	}

	//remove code after confirm
	if err := c.RemoveVerificationCode(ctx, email); err != nil {
		return false, false, err
	}
	return true, true, nil
}

// RemoveVerificationCode remove verification code
//
//	err := RemoveVerificationCode(ctx, "a@b.c")
//
func (c *Regional) RemoveVerificationCode(ctx context.Context, email string) error {

	if err := c.VerificationCodeTable().Delete(ctx, email); err != nil {
		return err
	}
	return nil
}

// RemoveAllVerificationCode remove all verification code
//
//	err := RemoveAllVerificationCode(ctx)
//
func (c *Regional) RemoveAllVerificationCode(ctx context.Context) error {
	return c.VerificationCodeTable().Clear(ctx)
}

// RemoveUnusedVerificationCode cleanup verification created more than 1 hour
//
//	err := RemoveUnusedVerificationCode(ctx)
//
func (c *Regional) RemoveUnusedVerificationCode(ctx context.Context) error {
	// verification code only valid for 1 hour.
	deadline := time.Now().Add(time.Duration(-1) * time.Hour).UTC()
	count, err := c.VerificationCodeTable().Query().Where("CreateTime", "<", deadline).Clear(ctx)
	if count > 0 {
		fmt.Printf("cleanup %v Job\n", count)
	}
	return err
}
