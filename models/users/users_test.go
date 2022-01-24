package users_test

import (
	"GWI_assingment/platform2.0-go-challenge/models/users"
	"testing"
)

func TestGetUserByIdSuccess(t *testing.T) {
	users.LoadUsers("./users.json")
	existingUser := users.User{Id: 1}
	ok := existingUser.GetUserById()
	if !ok {
		t.Fail()
	}
}

func TestGetUserByIdFailure(t *testing.T) {
	users.LoadUsers("./users.json")
	nonExUser := users.User{Id: 1213}
	ok := nonExUser.GetUserById()
	if ok {
		t.Fail()
	}
}

func TestCheckIdInSliceSuccess(t *testing.T) {
	users.LoadUsers("./users.json")
	existingUser := users.User{Id: 1}
	ok := existingUser.CheckIdInSlice()
	if !ok {
		t.Fail()
	}
}

func TestCheckIdInSliceFaiure(t *testing.T) {
	users.LoadUsers("./users.json")
	existingUser := users.User{Id: 1213123}
	ok := existingUser.CheckIdInSlice()
	if ok {
		t.Fail()
	}
}
