// go:generate go run github.com/99designs/gqlgen generate
package graph

import (
	"github.com/NhanThanhHuynh/gqlgen1/graph/model"
	"github.com/NhanThanhHuynh/gqlgen1/graph/postgres"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	tododetail []*model.TodoDetail `json:"todos,omitempty" :"tododetail" :"tododetail" :"tododetail"`
	todo       model.Todo          `:"todo" :"todo" :"todo"`
	user       model.User          `:"user" :"user"`
	UserRepo   postgres.UsersRepo  `:"user_repo"`
}
