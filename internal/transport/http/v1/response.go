package v1

type ResponseMessage struct {
	ID      int64       `json:"id,omitempty"`
	Status  int         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
