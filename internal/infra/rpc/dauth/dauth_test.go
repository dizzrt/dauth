package dauth

import (
	"context"
	"fmt"
	"testing"

	"github.com/dizzrt/dauth/api/gen/identity"
)

func TestGetUser(t *testing.T) {
	resp, err := GetUser(context.Background(), &identity.GetUserRequest{
		Id: 10000,
	})

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)
}
