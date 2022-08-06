package admin

import (
	"github.com/google/uuid"
)

// Admin is the domain model for an admin.
type Admin struct {
	ID    uuid.UUID
	Name  string
	Email string
	Role  RoleState
}

// RoleState is the state of an admin.
type RoleState struct {
	s string
}

var (
	Root      = RoleState{"root"}
	Marketing = RoleState{"marketing"}
	PM        = RoleState{"pm"}
	UI        = RoleState{"ui"}
	Dev       = RoleState{"dev"}
)

// AdminRepository is the interface for the AdminRepository defined at infrastructure's persistence.
type AdminRepository interface {
	GetById(id uuid.UUID) (*Admin, error)
	Create(admin *Admin) error
	Update(id uuid.UUID, admin *Admin) error
	Delete(id uuid.UUID) error
}
