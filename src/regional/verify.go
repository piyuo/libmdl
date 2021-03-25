package regional

import (
	"context"
	"time"

	"github.com/piyuo/libsrv/src/crypto"
	"github.com/piyuo/libsrv/src/db"
	"github.com/piyuo/libsrv/src/util"
	"github.com/pkg/errors"
)

// Verify keep verification code
//
type Verify struct {
	db.Entity

	// Hash is code hash with salt, we do not store code only hash is enough
	//
	Hash uint32 `firestore:"Hash,omitempty"`

	// Crypted code
	//
	Crypted string `firestore:"Crypted,omitempty"`
}

// Factory create a empty object, return object must be nil safe, no nil in any field
//
func (c *Verify) Factory() db.Object {
	return &Verify{}
}

// Collection return the name in database
//
func (c *Verify) Collection() string {
	return "Verify"
}

// CreateVerify create verification code
//
//	err := CreateVerify(ctx,"a@b.c","123456")
//
func CreateVerify(ctx context.Context, email, code string) error {
	client, err := RegionalClient(ctx)
	if err != nil {
		return err
	}

	hash := util.StringHash(code)
	crypted, err := crypto.Encrypt(code)
	if err != nil {
		return errors.Wrap(err, "encrypt "+code)
	}

	vc := &Verify{
		Hash:    hash,
		Crypted: crypted,
	}
	vc.SetID(email)
	if err := client.Set(ctx, vc); err != nil {
		return errors.Wrapf(err, "set verify %v,%v,%v", email, hash, crypted)
	}
	return nil
}

// GetVerify get verification code from database for resend, return found,code
//
//	found,code,err := GetVerify(ctx, "a@b.c")
//
func GetVerify(ctx context.Context, email string) (bool, string, error) {
	client, err := RegionalClient(ctx)
	if err != nil {
		return false, "", err
	}

	obj, err := client.Get(ctx, &Verify{}, email)
	if err != nil {
		return false, "", errors.Wrapf(err, "get verify %v", email)
	}

	if obj != nil {
		vc := obj.(*Verify)
		code, err := crypto.Decrypt(vc.Crypted)
		if err != nil {
			return false, "", errors.Wrap(err, "decrypt "+vc.Crypted)
		}
		return true, code, nil
	}
	return false, "", nil
}

// ConfirmVerify return found and confirm of a verify code
//
//	found,confirm, err := ConfirmVerify(ctx, "a@b.c", "123456")
//
func ConfirmVerify(ctx context.Context, email, code string) (bool, bool, error) {
	client, err := RegionalClient(ctx)
	if err != nil {
		return false, false, err
	}

	obj, err := client.Get(ctx, &Verify{}, email)
	if err != nil {
		return false, false, errors.Wrap(err, "get verify "+email)
	}
	if obj == nil {
		//verification code not exist, maybe removed after 30 min
		return false, false, nil
	}

	hash := util.StringHash(code)
	vc := obj.(*Verify)
	if vc.Hash != hash {
		//user input code is not match
		return true, false, nil
	}

	//remove code after confirm
	if err := DeleteVerify(ctx, email); err != nil {
		return false, false, err
	}
	return true, true, nil
}

// DeleteVerify remove verify code
//
//	err := DeleteVerify(ctx, "a@b.c")
//
func DeleteVerify(ctx context.Context, email string) error {
	client, err := RegionalClient(ctx)
	if err != nil {
		return err
	}

	v := &Verify{}
	v.SetID(email)
	if err := client.Delete(ctx, v); err != nil {
		return err
	}
	return nil
}

// DeleteUnusedVerify cleanup verify code created more than 1 hour
//
//	err := DeleteUnusedVerify(ctx)
//
func DeleteUnusedVerify(ctx context.Context, max int) (bool, error) {
	client, err := RegionalClient(ctx)
	if err != nil {
		return false, err
	}

	// verification code only valid for 1 hour.
	deadline := time.Now().Add(time.Duration(-1) * time.Hour).UTC()
	return client.Query(&Verify{}).Where("CreateTime", "<", deadline).Delete(ctx, max)
}
