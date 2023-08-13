package schema

type ProkerBodyReq struct {
	Name             string `validate:"required" json:"name"`
	Image            string `validate:"required" json:"image"`
	Description      string `validate:"required" json:"description"`
	ShortDescription string `validate:"required" json:"short_description"`
	TimeLineImage    string `validate:"" json:"time_line_image"`
	OrganizationId   string `validate:"" json:"organization_id"`
}
