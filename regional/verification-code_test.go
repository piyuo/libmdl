package regional

import (
	"context"
	"testing"
	"time"

	"github.com/piyuo/libsrv/data"
	"github.com/stretchr/testify/assert"
)

func TestVerificationCodeCRUD(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	r, err := New(ctx)
	assert.Nil(err)
	defer r.Close()

	//create
	err = r.CreateVerificationCode(ctx, "verification@code.com", "123456")
	assert.Nil(err)

	//get
	found, code, err := r.GetVerificationCode(ctx, "notExist@code.com") // wrong email
	assert.Nil(err)
	assert.False(found)
	assert.Empty(code)

	found, code, err = r.GetVerificationCode(ctx, "verification@code.com") // right email
	assert.Nil(err)
	assert.True(found)
	assert.Equal("123456", code)

	//confirm
	found, confirm, err := r.ConfirmVerificationCode(ctx, "verification@code.com", "111111") // wrong code
	assert.Nil(err)
	assert.True(found)
	assert.False(confirm)

	found, confirm, err = r.ConfirmVerificationCode(ctx, "verification@code.com", "123456") // right code
	assert.Nil(err)
	assert.True(found)
	assert.True(confirm)

	// verification code should be removed after confirm
	found, code, err = r.GetVerificationCode(ctx, "verification@code.com") // right email
	assert.Nil(err)
	assert.False(found)
}

func TestVerificationCodeCleanup(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	r, err := New(ctx)
	assert.Nil(err)
	defer r.Close()

	//create record need to be remove
	expired := &VerificationCode{
		BaseObject: data.BaseObject{
			ID:         "expired",
			CreateTime: time.Now().Add(time.Duration(-2) * time.Hour).UTC(),
		},
	}
	err = r.VerificationCodeTable().Set(ctx, expired)
	assert.Nil(err)

	valid := &VerificationCode{
		BaseObject: data.BaseObject{
			ID: "valid",
		},
	}
	err = r.VerificationCodeTable().Set(ctx, valid)
	assert.Nil(err)
	defer r.VerificationCodeTable().DeleteObject(ctx, valid)

	// before cleanup
	obj, err := r.VerificationCodeTable().Get(ctx, expired.ID)
	assert.Nil(err)
	assert.NotNil(obj)
	obj, err = r.VerificationCodeTable().Get(ctx, valid.ID)
	assert.Nil(err)
	assert.NotNil(obj)

	// cleanup
	err = r.RemoveUnusedVerificationCode(ctx)
	assert.Nil(err)

	// after cleanup
	obj, err = r.VerificationCodeTable().Get(ctx, expired.ID)
	assert.Nil(err)
	assert.Nil(obj)
	obj, err = r.VerificationCodeTable().Get(ctx, valid.ID)
	assert.Nil(err)
	assert.NotNil(obj)
}
