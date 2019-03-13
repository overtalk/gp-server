package module

// UserManage : User manage module
type UserManage interface {
	GetUsers(args map[string]interface{}) interface{}
	AddUsers(args map[string]interface{}) interface{}
	UpdateUsers(args map[string]interface{}) interface{}
	DelUsers(args map[string]interface{}) interface{}
}
