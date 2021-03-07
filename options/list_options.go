package options

// ListOptions type
type ListOptions struct {
	All    bool   `json:"all"`
	Search string `json:"search"`
	Page   int    `json:"page"`
	Size   int    `json:"size"`
}
