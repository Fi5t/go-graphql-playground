package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/Fi5t/go-graphql-playground/graph/generated"
	"github.com/Fi5t/go-graphql-playground/graph/model"
)

func (r *queryResolver) Todos(ctx context.Context, filter *model.TodosFilter, limit *int, offset *int) ([]*model.Todo, error) {
	var allTodos []*model.Todo

	for i := 0; i < 20; i++ {
		todo := model.Todo{
			ID:   strconv.Itoa(i),
			Text: fmt.Sprintf("Todo #%d", i),
		}

		allTodos = append(allTodos, &todo)
	}

	var result []*model.Todo

	if filter != nil {
		for _, todo := range allTodos {
			if filter.ID == todo.ID {
				result = append(result, todo)
			}
		}
	} else {
		result = allTodos

		if limit != nil && offset != nil {
			start := *offset
			end := *limit + *offset

			if end > len(allTodos) {
				end = len(allTodos)
			}

			return result[start:end], nil
		}
	}

	return result, nil
}

func (r *queryResolver) UsersConnection(ctx context.Context, first *int, after *string) (*model.UsersConnection, error) {
	var allUsers []*model.User

	for i := 0; i < 20; i++ {
		id := strconv.Itoa(i)

		user := model.User{
			ID:        id,
			FirstName: "Simple",
			LastName:  "User",
		}

		allUsers = append(allUsers, &user)
	}

	from := 0

	if after != nil {
		b, err := base64.StdEncoding.DecodeString(*after)

		if err != nil {
			return nil, err
		}

		i, err := strconv.Atoi(strings.TrimPrefix(string(b), "cursor"))

		if err != nil {
			return nil, err
		}

		from = i
	}

	to := len(allUsers)

	if first != nil {
		to = from + *first

		if to > len(allUsers) {
			to = len(allUsers)
		}
	}

	return &model.UsersConnection{
		Users: allUsers,
		From:  from,
		To:    to,
	}, nil
}

func (r *usersConnectionResolver) Edges(ctx context.Context, obj *model.UsersConnection) ([]*model.UsersEdge, error) {
	edges := make([]*model.UsersEdge, obj.To-obj.From)

	for i := range edges {
		edges[i] = &model.UsersEdge{
			Node:   obj.Users[obj.From+i],
			Cursor: model.EncodeCursor(obj.From + i),
		}
	}

	return edges, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// UsersConnection returns generated.UsersConnectionResolver implementation.
func (r *Resolver) UsersConnection() generated.UsersConnectionResolver {
	return &usersConnectionResolver{r}
}

type queryResolver struct{ *Resolver }
type usersConnectionResolver struct{ *Resolver }
