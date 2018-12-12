package utils

const (
	PageSize = 15
)

// Pagination represents pagination info.
type Page struct {
	Current int `json:"current"`
	Size    int `json:"size"`
	Total   int `json:"total"`
}

// NewPage creates a new pagination with the specified current page num, page size, window size and record count.
func NewPage(currentPageNum, pageSize, recordCount int) *Page {

	return &Page{
		Current: currentPageNum,
		Size:    pageSize,
		Total:   recordCount,
	}
}
