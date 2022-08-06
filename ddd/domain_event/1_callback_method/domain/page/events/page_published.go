package events

import (
	"time"

	"github.com/google/uuid"
)

// DomainEvent is the base struct for all domain events.
type DomainEvent struct {
	occurredAt time.Time
	name       string
	version    string
}

// PagePublished is the domain event for a page being published.
type PagePublished struct {
	DomainEvent
	PageId  string
	AdminId uuid.UUID
}

// NewPagePublished returns a new PagePublished domain event.
// the occurredAt field is set to the current time.
func NewPagePublished(pageId string, adminId uuid.UUID, version string) *PagePublished {
	return &PagePublished{
		DomainEvent: DomainEvent{
			occurredAt: time.Now(),
			name:       "PagePublished",
			version:    version,
		},
		PageId:  pageId,
		AdminId: adminId,
	}
}
