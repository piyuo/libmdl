package regional

import (
	"time"

	data "github.com/piyuo/libsrv/data"
)

// Counters is collection of regional usage counters
//
type Counters struct {
	data.Counters `firestore:"-"`
}

// Timezone using Los Angeles, CA, US timezone
//	counter := regional.StoreCounter()
//
func (c *Counters) Timezone() *time.Location {
	return time.FixedZone("PDT", -25200) //california timezone
}

// StoreCounter return store counter
//
//	counter := regional.StoreCounter()
//
func (c *Counters) StoreCounter() data.Counter {
	return c.Counter("Store", 1000, data.DateHierarchyNone) // 1,000 shard, safe concurrent use is is 100
}

// LocationCounter return location counter
//
//	counter := regional.LocationCounter()
//
func (c *Counters) LocationCounter() data.Counter {
	return c.Counter("Location", 1000, data.DateHierarchyNone) // 10,00 shard, safe concurrent use is 100
}
