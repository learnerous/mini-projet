package rest

type ResponseShape struct {
	Code    string   `json:"code"`
	Data    any      `json:"data"`
	Fields  []string `json:"fields"`
	Context string   `json:"context"`
}
