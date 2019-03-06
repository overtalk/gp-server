package module

// Auth : User identity authentication module
type Auth interface {
	Login(args map[string]interface{}) interface{}
	Logout(args map[string]interface{}) interface{}
}
