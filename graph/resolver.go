package graph

import "github.com/Fi5t/go-graphql-playground/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	todos []*model.Todo
	users []*model.User
}
