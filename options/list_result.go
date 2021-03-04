package options

// ListResult type
type ListResult struct {
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}
