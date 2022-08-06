package page

import (
	"database/sql"

	"github.com/yiting-tom/LF_Golang/ddd/domain_event/1_callback_method/domain/page"
)

type SqlPageRepository struct {
	db *sql.DB
	// and other fields
}

// GetById returns a Page Domain Model by its id.
func (r *SqlPageRepository) GetById(id string) (*page.Page, error) {
	// do something
	return nil, nil
}

// Save saves a Page Domain Model.
func (r *SqlPageRepository) Create(page *page.Page) error {
	// do something
	return nil
}

// Remove removes a Page Domain Model.
func (r *SqlPageRepository) Remove(id string) error {
	// do something
	return nil
}

// Update updates a Page Domain Model.
func (r *SqlPageRepository) Update(page *page.Page) error {
	// do something
	return nil
}
