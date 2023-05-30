package health

type Readiness struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
