 # GP Server Part

 ## Overview
+ all modules are defined as interfaces for easy replacement
+ Web server, judgment server separation
+ docker deployment available

## Module（./module）
 
### Basic Module
- Gateway（gate.go）
- Configuration (config.go)
- Database（db.go）
- Distributed Shared Cache（cache.go）
 
### Business Logic Module
- User Authentication (auth.go)
- Administrator Manage (manage.go)


## Module Implementation（./service）

### Gateway（package gate）
- stateless
- TLS available
- all services are registered by each module

### Configuration (package config)
- configuration source separation

### Database（package db）
- xorm frame + mysql

### Distributed shared cache（package cache）
- redis