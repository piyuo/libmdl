package regional

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestVerificationCodeCRUD(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ctx := context.Background()

	//create
	err := CreateVerificationCode(ctx, "verification@code.com", "123456")
	assert.Nil(err)

	//get
	found, code, err := GetVerificationCode(ctx, "notExist@code.com") // wrong email
	assert.Nil(err)
	assert.False(found)
	assert.Empty(code)

	found, code, err = GetVerificationCode(ctx, "verification@code.com") // right email
	assert.Nil(err)
	assert.True(found)
	assert.Equal("123456", code)

	//confirm
	found, confirm, err := ConfirmVerificationCode(ctx, "verification@code.com", "111111") // wrong code
	assert.Nil(err)
	assert.True(found)
	assert.False(confirm)

	found, confirm, err = ConfirmVerificationCode(ctx, "verification@code.com", "123456") // right code
	assert.Nil(err)
	assert.True(found)
	assert.True(confirm)

	// verification code should be delete after confirm
	found, code, err = GetVerificationCode(ctx, "verification@code.com") // right email
	assert.Nil(err)
	assert.False(found)
}

func TestVerificationCodeCleanup(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ctx := context.Background()
	client, err := RegionalClient(ctx)
	assert.Nil(err)

	//create record need to be remove
	expired := &VerificationCode{}
	expired.SetID("expired")
	expired.SetCreateTime(time.Now().Add(time.Duration(-2) * time.Hour).UTC())
	err = client.Set(ctx, expired)
	assert.Nil(err)

	valid := &VerificationCode{}
	valid.SetID("valid")

	err = client.Set(ctx, valid)
	assert.Nil(err)
	defer client.Delete(ctx, valid)

	// before cleanup
	obj, err := client.Get(ctx, &VerificationCode{}, expired.ID())
	assert.Nil(err)
	assert.NotNil(obj)
	obj, err = client.Get(ctx, &VerificationCode{}, valid.ID())
	assert.Nil(err)
	assert.NotNil(obj)

	// cleanup
	done, err := DeleteUnusedVerificationCode(ctx, 1000)
	assert.Nil(err)
	assert.True(done)

	// after cleanup
	obj, err = client.Get(ctx, &VerificationCode{}, expired.ID())
	assert.Nil(err)
	assert.Nil(obj)
	obj, err = client.Get(ctx, &VerificationCode{}, valid.ID())
	assert.Nil(err)
	assert.NotNil(obj)
}
