package pkg_requests

type PageRequest struct {
	Search        string `query:"search"`
	PageNumber    uint   `query:"page_number"`
	PageSize      uint   `query:"page_size"`
	SortBy        string `query:"sort_by"`
	SortDirection string `query:"sort_direction"`
}
