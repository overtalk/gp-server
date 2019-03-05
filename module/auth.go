package module

// Auth : User identity authentication module
type Auth interface {
	Login(data ...interface{}) interface{}
	Logout(data ...interface{}) interface{}
}
