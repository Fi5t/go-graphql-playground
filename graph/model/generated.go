// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type PageInfo struct {
	StartCursor string `json:"startCursor"`
	EndCursor   string `json:"endCursor"`
	HasNextPage bool   `json:"hasNextPage"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type TodosFilter struct {
	ID string `json:"id"`
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UsersEdge struct {
	Node   *User  `json:"node"`
	Cursor string `json:"cursor"`
}