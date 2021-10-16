package graph

import (
	"github.com/ducnguyen96/reddit-clone/graph/services/auth_services"
	"github.com/ducnguyen96/reddit-clone/graph/services/user_services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	UerService *user_services.UserService
	AuthService *auth_services.AuthService
}