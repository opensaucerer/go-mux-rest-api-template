package typing

type Health struct {
	Version     string `json:"version"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
}

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
