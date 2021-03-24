package global

import (
	"time"

	"github.com/piyuo/libsrv/src/db"
)

// Bill represent bill
//
type Bill struct {
	db.Model

	// BeginDate is billing period date begin
	//
	BeginDate time.Time `firestore:"BeginDate,omitempty"`

	// EndDate is billing period date end
	//
	EndDate time.Time `firestore:"EndDate,omitempty"`

	// Email is owner email where bill send
	//
	Email string `firestore:"Email,omitempty"`

	// FirstName is owner first name where bill send
	//
	FirstName string `firestore:"FirstName,omitempty"`

	// LastName is owner last name where bill send
	//
	LastName string `firestore:"LastName,omitempty"`

	// Plan is account servie plan
	//
	Plan AccountPlan `firestore:"Plan,omitempty"`

	// Currency is plan fee currency
	//
	Currency string `firestore:"Currency,omitempty"`

	// Plan is account servie plan
	//
	PlanServiceFee int64 `firestore:"PlanServiceFee,omitempty"`

	//additional billable item
}

func (c *Bill) Factory() db.Object {
	return &Bill{}
}

func (c *Bill) Collection() string {
	return "Bill"
}
