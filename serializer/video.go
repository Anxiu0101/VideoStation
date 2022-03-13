package serializer

type Video struct {
	ID       uint   `json:"id" form:"id" example:"1"`
	Title    string `json:"title" form:"title" example:"My first step in go learning"`
	Status   string `json:"status" form:"status"`
	CreateAt int64  `json:"create_at" form:"create_at"`
}
