package schema

type ProkerBodyReq struct {
	Name          string `validate:"required" json:"name"`
	Image         string `validate:"required" json:"image"`
	Description   string `validate:"required" json:"description"`
	TimeLineImage string `validate:"required" json:"time_line_image"`
}
