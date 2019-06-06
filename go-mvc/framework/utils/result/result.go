package result

type Result struct {
	State   bool        `json:"State"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data,omitempty"`
}
