package interfaces

type ResponseBody struct {
	Message string         `json:"message"`
	Data    map[string]any `json:"data"`
}
