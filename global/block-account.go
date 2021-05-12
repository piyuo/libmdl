package global

import (
	"github.com/piyuo/libsrv/db"
)

// BlockAccount id in block account can't be register
//
type BlockAccount struct {
	db.Model
}

func (c *BlockAccount) Factory() db.Object {
	return &BlockAccount{}
}

func (c *BlockAccount) Collection() string {
	return "BlockAccount"
}
