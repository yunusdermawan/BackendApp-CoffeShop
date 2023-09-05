package config

type Metas struct {
	Next  interface{} `json:"next"`
	Prev  interface{} `json:"prev"`
	Total int         `json:"total"`
}

type Result struct {
	Data    interface{} `json:"data, omitempty"`
	Meta    interface{} `json:"meta, omitempty"`
	Message interface{} `json:"message, omitempty"`
}
