package model

import (
	"testing"
)

func TestModelUser(t *testing.T) {
	tv := "0"
	user := &User{
		ID:         "1",
		Firstname:  "Lucas",
		Lastname:   "Laurentino",
		Email:      "laurentino14@outlook.com",
		Password:   "123",
		Cellphone:  "123",
		TokenUser:  &tv,
		Enrollment: nil,
	}
	ID := "1"
	Firstname := "Lucas"
	Lastname := "Laurentino"
	Email := "laurentino14@outlook.com"
	Password := "123"
	Cellphone := "123"
	TokenUser := &tv
	Enrollment := []*Enrollment{}

	if user.ID != ID {
		t.Error("Expected: ", ID, "Got:", user.ID)
	}
	if user.Firstname != Firstname {
		t.Error("Expected: ", Firstname, "Got:", user.Firstname)
	}
	if user.Lastname != Lastname {
		t.Error("Expected: ", Lastname, "Got:", user.Lastname)
	}
	if user.Email != Email {
		t.Error("Expected: ", Email, "Got:", user.Email)
	}
	if user.Password != Password {
		t.Error("Expected: ", Password, "Got:", user.Password)
	}
	if user.Cellphone != Cellphone {
		t.Error("Expected: ", Cellphone, "Got:", user.Cellphone)
	}
	if user.TokenUser != TokenUser {
		t.Error("Expected: ", TokenUser, "Got:", user.TokenUser)
	}
	if user.Enrollment != nil {
		t.Error("Expected: ", Enrollment, "Got:", user.Enrollment)
	}
}
