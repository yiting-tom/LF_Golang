package order

// ApplicationService is the interface for the application service
type ApplicationService interface {
    // Execute is the entry point for the application service
    Execute(ApplicationServiceInput) (*ApplicationServiceOutput, error)
}
// ApplicationServiceInput is the input for the ApplicationService
type ApplicationServiceInput struct {}
// ApplicationServiceOutput is the output for the ApplicationService
type ApplicationServiceOutput struct {}


// RegisterMemberInput is the input for the RegisterMember application service
type RegisterMemberInput struct {
    ApplicationServiceInput
    Name string
    Email string                    
    Password string
}

// RegisterMemberOutput is the output for the RegisterMember application service
type RegisterMemberOutput struct {
    ApplicationServiceOutput
    Success bool
    Member *MemberDto
    ErrorMessage string
}

// MemberDto is the data transfer object for the Member
type MemberDto struct {
    Name string
    Email string
}

// RegisterMember is the application service for registering a new member
type RegisterMember struct {
    ApplicationService
    // should have repository and other fields
    // but for the sake of the example we will leave it empty
    // repo *MemberRepository
}

// NewRegisterMember returns a new RegisterMember application service
func NewRegisterMember() *RegisterMember {
    return &RegisterMember{}
}

// Execute executes the application service
func (r *RegisterMember) Execute(input RegisterMemberInput) *RegisterMemberOutput {
    // do what should be done to register the member
    // it would operates the domain model and the repository
    return &RegisterMemberOutput{
        Success: true,
        Member: &MemberDto{
            Name: input.Name,
            Email: input.Email,
        },
        ErrorMessage: "",
    }
}