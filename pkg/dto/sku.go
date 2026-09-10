package dto

type SKUListItem struct {
		ID          			uint           			`json:"id"`
    Code        			string         			`json:"code"`
    Price 			      float64         		`json:"price"`
    Stock             int                 `json:"stock"`
}
