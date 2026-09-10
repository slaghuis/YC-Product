package dto

type ParameterListItem struct {
  ParameterName   	string					`json:"parameter_name"`
  CategoryId  			uint						`json:"category_id"`
  ParameterId       uint            `json:"parameter_id"`
  Markdown    			string          `json:"markdown"`
}
