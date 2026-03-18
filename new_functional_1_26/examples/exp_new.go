package examples

import (
	"fmt"
	"reflect"
	"time"
)

type UpdateUserRequest struct {
	Name     *string    `json:"name,omitempty"`
	Email    *string    `json:"email,omitempty"`
	Age      *int       `json:"age,omitempty"`
	IsActive *bool      `json:"is_active,omitempty"`
	UpdateAt *time.Time `json:"update_at,omitempty"`
}

func ptr[T any](v T) *T { return &v }

func printUsrAndFieldTypes(usr UpdateUserRequest) {
	fmt.Printf("usr: %#v\n", usr)
	fmt.Printf("type(usr): %T\n", usr)

	v := reflect.ValueOf(usr)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		fv := v.Field(i)
		fmt.Printf("%s: type=%s value=%#v\n", f.Name, f.Type, fv.Interface())
	}
}

func PtrNew() {
	usr := UpdateUserRequest{
		Name:     ptr("John"),
		Age:      ptr(int(25)),
		IsActive: ptr(true),
		UpdateAt: ptr(time.Now()),
	}

	fmt.Println("Go before 1.26")
	printUsrAndFieldTypes(usr)

	// Go 1.26:
	usr = UpdateUserRequest{
		Name:     new("Smith"),
		Age:      new(int(25)),
		IsActive: new(false),
		UpdateAt: new(time.Now()),
	}

	fmt.Println("Go 1.26")
	printUsrAndFieldTypes(usr)
}
