package model

type PageInfo struct {
	TotalCount int64 `json:"totalCount,omitempty"`
	Offset     int   `json:"offset,omitempty"`
	Limit      int   `json:"limit,omitempty"`
	HasNext    bool  `json:"hasNext,omitempty"`
}
