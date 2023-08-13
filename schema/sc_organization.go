package schema

type OrganizationBodyReq struct {
	Name        string `validate:"required" json:"name"`
	Logo        string `validate:"required" json:"logo"`
	Description string `validate:"required" json:"description"`
}
