package results

type PaginationRequest struct {
	Page         int    `form:"page" validate:"required,min=1" binding:"required"`              // Required, must be at least 1
	PageSize     int    `form:"page_size" validate:"required,min=1,max=100" binding:"required"` // Required, minimum 1, maximum 100
	SortBy       string `form:"sort_by" validate:"omitempty"`                                   // Optional, must be one of "title", "date", or "author"
	SortOrder    string `form:"sort_order" validate:"omitempty,oneof=asc desc"`                 // Optional, must be "asc" or "desc"
	FilterSearch string `form:"filter_search" validate:"omitempty,max=100"`                     // Optional, maximum 100 characters
}

type Pagination struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalCount int `json:"total_count"`
}

type PaginationResponse struct {
	IsNext     bool `json:"is_next"`     // IsNext Indicates if there's a next page
	IsPrev     bool `json:"is_prev"`     // IsPrev Indicates if there's a previous page
	Page       int  `json:"page"`        // Page Current page number
	PageSize   int  `json:"page_size"`   // PageSize Number of items per page
	Total      int  `json:"total"`       // Total number of items
	TotalPages int  `json:"total_pages"` // Total number of pages
}
