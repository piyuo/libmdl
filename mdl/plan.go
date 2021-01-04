package mdl

// Plan is piyuo service plan, used in account
//
type Plan int

const (
	// PlanFree is free Plan, free for 1 year
	//
	PlanFree Plan = 0

	// PlanStandard is standard Plan
	//
	PlanStandard Plan = 1

	// PlanBusiness is business plan
	//
	PlanBusiness Plan = 2
)
