package schema

type CommentBodyReq struct {
	Description string `validate:"required" json:"description"`
}
