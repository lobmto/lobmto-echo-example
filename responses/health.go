package responses

type HealthResponse struct {
	Status string `json:"status"`
}

func NewHealthResponse() *HealthResponse {
	return &HealthResponse{
		Status: "UP",
	}
}
