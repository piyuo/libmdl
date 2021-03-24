package regional

import (
	"github.com/piyuo/libsrv/src/db"
)

// Menu represent menu in location
//
type Menu struct {
	db.Model
}

func (c *Menu) Factory() db.Object {
	return &Menu{}
}

func (c *Menu) Collection() string {
	return "Menu"
}
