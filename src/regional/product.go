package regional

import "github.com/piyuo/libsrv/src/db"

// Product represent product in store
//
type Product struct {
	db.Model
}

func (c *Product) Factory() db.Object {
	return &Product{}
}

func (c *Product) Collection() string {
	return "Product"
}
