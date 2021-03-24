package regional

import (
	"context"
	"time"

	"github.com/piyuo/libsrv/src/crypto"
	"github.com/piyuo/libsrv/src/db"
	"github.com/piyuo/libsrv/src/util"
	"github.com/pkg/errors"
)

// Verification keep verification code
//
type Verification struct {
	db.Entity

	// Hash is code hash with salt, we do not store code only hash is enough
	//
	Hash uint32 `firestore:"Hash,omitempty"`

	// Crypted code
	//
	Crypted string `firestore:"Crypted,omitempty"`
}

func (c *Verification) Factory() db.Object {
	return &Verification{}
}

func (c *Verification) Collection() string {
	return "Verification"
}

// CreateVerification create verification code
//
//	err := CreateVerification(ctx,"a@b.c","123456")
//
func CreateVerification(ctx context.Context, email, code string) error {
	client, err := RegionalClient(ctx)
	if err != nil {
		return err
	}

	hash := util.StringHash(code)
	crypted, err := crypto.Encrypt(code)
	if err != nil {
		return errors.Wrap(err, "encrypt "+code)
	}

	vc := &Verification{
		Hash:    hash,
		Crypted: crypted,
	}
	vc.SetID(email)
	if err := client.Set(ctx, vc); err != nil {
		return errors.Wrap(err, "failed to Set verification code")
	}
	return nil
}

// GetVerification get verification code from database for resend, return found,code
//
//	found,code,err := GetVerification(ctx, "a@b.c")
//
func GetVerification(ctx context.Context, email string) (bool, string, error) {
	client, err := RegionalClient(ctx)
	if err != nil {
		return false, "", err
	}

	obj, err := client.Get(ctx, &Verification{}, email)
	if err != nil {
		return false, "", errors.Wrap(err, "failed to Get data")
	}

	if obj != nil {
		vc := obj.(*Verification)
		code, err := crypto.Decrypt(vc.Crypted)
		if err != nil {
			return false, "", errors.Wrap(err, "failed to Decrypt: "+vc.Crypted)
		}
		return true, code, nil
	}
	return false, "", nil
}

// ConfirmVerification return found and confirm of a verification code
//
//	found,confirm, err := ConfirmVerification(ctx, "a@b.c", "123456")
//
func ConfirmVerification(ctx context.Context, email, code string) (bool, bool, error) {
	client, err := RegionalClient(ctx)
	if err != nil {
		return false, false, err
	}

	obj, err := client.Get(ctx, &Verification{}, email)
	if err != nil {
		return false, false, errors.Wrap(err, "failed to Get data")
	}
	if obj == nil {
		//verification code not exist, maybe removed after 30 min
		return false, false, nil
	}

	hash := util.StringHash(code)
	vc := obj.(*Verification)
	if vc.Hash != hash {
		//user input code is not match
		return true, false, nil
	}

	//remove code after confirm
	if err := DeleteVerification(ctx, email); err != nil {
		return false, false, err
	}
	return true, true, nil
}

// DeleteVerification remove verification code
//
//	err := DeleteVerification(ctx, "a@b.c")
//
func DeleteVerification(ctx context.Context, email string) error {
	client, err := RegionalClient(ctx)
	if err != nil {
		return err
	}

	v := &Verification{}
	v.SetID(email)
	if err := client.Delete(ctx, v); err != nil {
		return err
	}
	return nil
}

// DeleteUnusedVerification cleanup verification created more than 1 hour
//
//	err := DeleteUnusedVerification(ctx)
//
func DeleteUnusedVerification(ctx context.Context, max int) (bool, error) {
	client, err := RegionalClient(ctx)
	if err != nil {
		return false, err
	}

	// verification code only valid for 1 hour.
	deadline := time.Now().Add(time.Duration(-1) * time.Hour).UTC()
	return client.Query(&Verification{}).Where("CreateTime", "<", deadline).Delete(ctx, max)
}
