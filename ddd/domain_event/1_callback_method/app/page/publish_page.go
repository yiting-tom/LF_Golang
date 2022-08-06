package page

import (
	"github.com/google/uuid"
	"github.com/yiting-tom/LF_Golang/ddd/domain_event/1_callback_method/domain/admin"
	"github.com/yiting-tom/LF_Golang/ddd/domain_event/1_callback_method/domain/page"
	"github.com/yiting-tom/LF_Golang/ddd/domain_event/1_callback_method/domain/page/events"
)

// PublishPageInput is the input for the PublishPage.
type PublishPageInput struct {
	PageId  string
	AdminId uuid.UUID
}

// PublishPageOutput is the output for the PublishPage.
type PublishPageOutput struct {
	err error
	// and other outputs
}

// PublishPage is the application service for publishing a page.
type PublishPage struct {
	pageRepository  page.PageRepository
	adminRepository admin.AdminRepository
	// just for example
	publishBucket interface{}
}

// NewPublishPage returns a new PublishPage application service.
func NewPublishPage(pageRepository page.PageRepository, adminRepository admin.AdminRepository) *PublishPage {
	return &PublishPage{
		pageRepository:  pageRepository,
		adminRepository: adminRepository,
		publishBucket:   nil,
	}
}

// Execute is the entry point for the PublishPage application service.
func (s *PublishPage) Execute(input PublishPageInput) PublishPageOutput {
	// 1. get the page from the repository
	page, err := s.pageRepository.GetById(input.PageId)
	if err != nil {
		return PublishPageOutput{err: err}
	}

	page.AddPagePublishedHandler(s.OnPagePublished)
	page.Publish(input.AdminId)

	// go func() {
	// 	s.publishBucket.Publish(page)
	// }()

	return PublishPageOutput{
		err: nil,
	}
}

func (s *PublishPage) OnPagePublished(event *events.PagePublished) {
	// do something
}
