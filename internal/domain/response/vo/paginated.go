package aggregation

import "fmt"

type PaginatedResponse struct {
	CurrentPage  int         `json:"currentPage"`
	Data         interface{} `json:"data"`
	FirstPageURL string      `json:"firstPageURL"`
	From         int         `json:"from"`
	LastPage     int         `json:"lastPage"`
	LastPageURL  interface{} `json:"lastPageURL"`
	NextPageURL  interface{} `json:"nextPageURL"`
	Path         string      `json:"path"`
	PerPage      int         `json:"perPage"`
	PrevPageURL  interface{} `json:"prevPageURL"`
	To           int         `json:"to"`
	Total        int64       `json:"total"`
}

func NewPaginatedResponse(data interface{}, page int, pageSize int, total int64, path string) *PaginatedResponse {
	lastPage := int(total) / pageSize
	if int(total)%pageSize != 0 {
		lastPage++
	}

	return &PaginatedResponse{
		CurrentPage:  page,
		Data:         data,
		FirstPageURL: fmt.Sprintf("%s?page=1", path),
		From:         (page-1)*pageSize + 1,
		LastPage:     lastPage,
		LastPageURL:  fmt.Sprintf("%s?page=%d", path, lastPage),
		NextPageURL:  getNextPageURL(page, lastPage, path),
		Path:         path,
		PerPage:      pageSize,
		PrevPageURL:  getPrevPageURL(page, path),
		To:           min((page)*pageSize, int(total)),
		Total:        total,
	}
}

func getNextPageURL(currentPage, lastPage int, path string) interface{} {
	if currentPage < lastPage {
		return fmt.Sprintf("%s?page=%d", path, currentPage+1)
	}
	return nil
}

func getPrevPageURL(currentPage int, path string) interface{} {
	if currentPage > 1 {
		return fmt.Sprintf("%s?page=%d", path, currentPage-1)
	}
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
