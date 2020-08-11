package acl

// Rule is Casbin rule
//
type Rule struct {

	// Sub is subject
	//
	Sub string

	// Obj is object
	//
	Obj string

	// Act is action
	//
	Act string
}
