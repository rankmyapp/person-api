package entity

type DeletedResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

func NewDeletedResponse(id, object string) *DeletedResponse {
	return &DeletedResponse{
		ID:      id,
		Object:  object,
		Deleted: true,
	}
}
