package model

import (
	"encoding/base64"
	"fmt"
)

type UsersConnection struct {
	Users []*User
	From  int
	To    int
}

func (u *UsersConnection) TotalCount() int {
	return len(u.Users)
}

func (u *UsersConnection) PageInfo() PageInfo {
	return PageInfo{
		StartCursor: EncodeCursor(u.From),
		EndCursor:   EncodeCursor(u.To - 1),
		HasNextPage: u.To < len(u.Users),
	}
}

func EncodeCursor(i int) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("cursor%d", i+1)))
}
