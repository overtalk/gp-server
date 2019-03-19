package module

// BackStageManage : the backstage administration  module
type BackStageManage interface {
	// user manage
	GetUsers(args map[string]interface{}) interface{}
	AddUsers(args map[string]interface{}) interface{}
	UpdateUsers(args map[string]interface{}) interface{}
	DelUsers(args map[string]interface{}) interface{}

	// problems manage
	GetProblems(args map[string]interface{}) interface{}
	GetProblemByID(args map[string]interface{}) interface{}
	AddProblem(args map[string]interface{}) interface{}
	EditProblem(args map[string]interface{}) interface{}
}
