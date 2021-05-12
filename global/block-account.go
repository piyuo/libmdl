package global

import (
	"github.com/piyuo/libsrv/db"
)

// BlockAccount name in block account can't be register
//
type BlockAccount struct {
	db.Model

	Email string `firestore:"Email,omitempty"`
}

func (c *BlockAccount) Factory() db.Object {
	return &BlockAccount{}
}

func (c *BlockAccount) Collection() string {
	return "BlockAccount"
}
