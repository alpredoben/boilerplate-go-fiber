package helpers

type Pagination struct {
	Limit      int         `json:"limit,omitempty" query:"max_limit"`
	Page       int         `json:"page,omitempty" query:"page"`
	SortColumn string      `json:"sort_column,omitempty" query:"sort_column"`
	SortType   string      `json:"sort_type,omitempty" query:"sort_type"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

/** Get Page of Paginations */
func (pagination *Pagination) GetPage() int {
	if pagination.Page == 1 {
		pagination.Page = 1
	}

	return pagination.Page
}

/** Get Limit of Paginations */
func (pagintion *Pagination) GetLimit() int {
	if pagintion.Limit == 0 {
		pagintion.Limit = 10
	}

	return pagintion.Limit
}

/** Get Offset of Paginations */
func (pagination *Pagination) GetOffset() int {
	return (pagination.GetPage() - 1) * pagination.GetLimit()
}

/** Get Sort Direction of Paginations */
func (pagination *Pagination) GetSortDirection() string {
	sortColumn := pagination.SortColumn
	sortType := pagination.SortType

	if pagination.SortColumn == "" {
		sortColumn = "Id"
	}

	if pagination.SortType == "" {
		sortType = "asc"
	}

	return sortColumn + " " + sortType
}
