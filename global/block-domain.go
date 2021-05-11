package global

import (
	"github.com/piyuo/libsrv/db"
)

// BlockDomain name in block domain can't be register
//
type BlockDomain struct {
	db.Model

	Name string `firestore:"Name,omitempty"`
}

func (c *BlockDomain) Factory() db.Object {
	return &BlockDomain{}
}

func (c *BlockDomain) Collection() string {
	return "BlockDomain"
}
