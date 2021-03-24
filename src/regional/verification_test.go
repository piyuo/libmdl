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
	err := CreateVerification(ctx, "verification@code.com", "123456")
	assert.Nil(err)

	//get
	found, code, err := GetVerification(ctx, "notExist@code.com") // wrong email
	assert.Nil(err)
	assert.False(found)
	assert.Empty(code)

	found, code, err = GetVerification(ctx, "verification@code.com") // right email
	assert.Nil(err)
	assert.True(found)
	assert.Equal("123456", code)

	//confirm
	found, confirm, err := ConfirmVerification(ctx, "verification@code.com", "111111") // wrong code
	assert.Nil(err)
	assert.True(found)
	assert.False(confirm)

	found, confirm, err = ConfirmVerification(ctx, "verification@code.com", "123456") // right code
	assert.Nil(err)
	assert.True(found)
	assert.True(confirm)

	// verification code should be delete after confirm
	found, code, err = GetVerification(ctx, "verification@code.com") // right email
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
	expired := &Verification{}
	expired.SetID("expired")
	expired.SetCreateTime(time.Now().Add(time.Duration(-2) * time.Hour).UTC())
	err = client.Set(ctx, expired)
	assert.Nil(err)

	valid := &Verification{}
	valid.SetID("valid")

	err = client.Set(ctx, valid)
	assert.Nil(err)
	defer client.Delete(ctx, valid)

	// before cleanup
	obj, err := client.Get(ctx, &Verification{}, expired.ID())
	assert.Nil(err)
	assert.NotNil(obj)
	obj, err = client.Get(ctx, &Verification{}, valid.ID())
	assert.Nil(err)
	assert.NotNil(obj)

	// cleanup
	done, err := DeleteUnusedVerification(ctx, 1000)
	assert.Nil(err)
	assert.True(done)

	// after cleanup
	obj, err = client.Get(ctx, &Verification{}, expired.ID())
	assert.Nil(err)
	assert.Nil(obj)
	obj, err = client.Get(ctx, &Verification{}, valid.ID())
	assert.Nil(err)
	assert.NotNil(obj)
}
