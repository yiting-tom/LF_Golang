package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// Address is a value object for address of a User
type Address struct {
	City   string
	County string
	Street string
	Detail string
}

// User is an entity
type User struct {
	Id      uuid.UUID
	Name    string
	Address *Address
}

// NewUser creates a new user with auto-generated uuid as id
func NewUser(name string, address *Address) (*User, error) {
	// generate uuid
	newId, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate uuid: %w", err)
	}

	return &User{
		Id:      newId,
		Name:    name,
		Address: address,
	}, nil
}

// SetAddress sets the address of the user
func (u *User) SetAddress(address *Address) {
	u.Address = address
}

// Equals checks if two users are equal
func (u *User) Equals(other *User) bool {
	return u.Id == other.Id
}

// String returns the string representation of the user
func (u *User) String() string {
	m, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	return string(m)
}

func main() {
	// create a user
	user1, _ := NewUser("John", &Address{
		City:   "New York",
		County: "New York",
		Street: "Wall Street",
		Detail: "No. 1",
	})

	user2, _ := NewUser("John", &Address{
		City:   "New York",
		County: "New York",
		Street: "Wall Street",
		Detail: "No. 1",
	})

	// print user
	fmt.Printf("user1: %s\nuser2: %s\n", user1.String(), user2.String())
	fmt.Printf("user1 equal to user2: %v",user1.Equals(user2))
}