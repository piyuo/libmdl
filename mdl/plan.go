package mdl

// Plan is piyuo service plan, used in account
//
type Plan int

const (
	// PlanFree free Plan, free
	//
	PlanFree Plan = 0

	// PlanBasic basic Plan, 0.99/month
	//
	PlanBasic Plan = 1

	// PlanPremium basic Plan, 9.99/month
	//
	PlanPremium Plan = 2

	// PlanBusiness business plan, 99.99/month
	//
	PlanBusiness Plan = 3

	// PlanEnterprise enterprise Plan, 999.99/month
	//
	PlanEnterprise Plan = 6

	// PlanPartner partner Plan, 99,999.99/month
	//
	PlanPartner Plan = 6
)
