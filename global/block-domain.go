package global

import (
	"github.com/piyuo/libsrv/db"
)

// BlockDomain id in block domain can't be register
//
type BlockDomain struct {
	db.Entity
}

func (c *BlockDomain) Factory() db.Object {
	return &BlockDomain{}
}

func (c *BlockDomain) Collection() string {
	return "BlockDomain"
}
