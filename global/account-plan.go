package global

// AccountPlan is piyuo service plan, used in account
//
type AccountPlan int

const (
	// AccountPlanFree is free Plan, free for 1 year
	//
	AccountPlanFree AccountPlan = 0

	// AccountPlanStandard is standard Plan
	//
	AccountPlanStandard = 1

	// AccountPlanBusiness is business plan
	//
	AccountPlanBusiness = 2
)
