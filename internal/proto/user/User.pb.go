// Code generated by protoc-gen-go-custom. DO NOT EDIT.

package user

import "fmt"

type Person struct {
	Id     int32  `json:"id" sql:"id"`
	Name   string `json:"name" sql:"name"`
	Email  string `json:"email" sql:"email"`
	Class  string `json:"class" sql:"class"`
	Travel int32  `json:"travel" sql:"travel"`
}

const (
	PAKISTAN = 0
	INDIA    = 1
	UAE      = 2
	USA      = 3
)

func (p Person) ValidateEmail() error {
	if len(p.Email) < 5 {
		return fmt.Errorf("email length should be greater than zero")
	}
	return nil
}
