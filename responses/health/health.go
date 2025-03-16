package health

func NewGetHealthResponse() *GetHealthResponse {
	return &GetHealthResponse{
		Status: UP,
	}
}
