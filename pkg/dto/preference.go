package dto

type Filter struct {
    FilterType string        `json:"filterType"`
    Values     []string      `json:"values,omitempty"`
    Range      *RangeFilter  `json:"range,omitempty"`
    Metadata   *FilterMeta   `json:"metadata,omitempty"`
}

type RangeFilter struct {
    Min      float64 `json:"min"`
    Max      float64 `json:"max"`
    Currency string  `json:"currency"`
}

type FilterMeta struct {
    Priority  int  `json:"priority,omitempty"`
    Exclusive bool `json:"exclusive,omitempty"`
}

type Preferences struct {
    Filters []Filter `json:"filters"`
}

type SortOrder struct {
    SortOrder string `json:"sort_order"`
}
