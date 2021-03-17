package options

//Result type
type Result struct {
	Data interface{} `json:"data"`
}

// ListResult type
type ListResult struct {
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}
