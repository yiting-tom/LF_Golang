package page

import (
	"time"

	"github.com/google/uuid"
	"github.com/yiting-tom/LF_Golang/ddd/domain_event/1_callback_method/domain/page/events"
)

// Page is the domain model for a page.
type Page struct {
	ID                    string
	Title                 string
	Content               string
	State                 PageState
	CreatedAt             time.Time
	UpdatedAt             time.Time
	CreatedBy             uuid.UUID
	UpdatedBy             uuid.UUID
	PagePublishedHandlers []PagePublishedHandler
}

// PageState is the state of a page.
type PageState struct {
	s string
}

var (
	Draft       = PageState{"draft"}
	Published   = PageState{"published"}
	Unpublished = PageState{"unpublished"}
	Archived    = PageState{"archived"}
)

// PageRepository is the interface for the PageRepository defined at infrastructure's persistence.
type PageRepository interface {
	GetById(id string) (*Page, error)
	Create(page *Page) error
	Update(pageId string, page *Page) error
	Remove(id string) error
}

// PagePublishedHandler is a function which is called when a page is published.
type PagePublishedHandler func(*events.PagePublished)

// AddPagePublishedHandler adds a PagePublishedHandler to the list of PagePublishedHandlers.
func (p *Page) AddPagePublishedHandler(handler PagePublishedHandler) {
	p.PagePublishedHandlers = append(p.PagePublishedHandlers, handler)
}

func (p *Page) Publish(adminId uuid.UUID) {
	event := events.NewPagePublished(p.ID, adminId, "1.0")
	for _, handler := range p.PagePublishedHandlers {
		handler(event)
	}
}
