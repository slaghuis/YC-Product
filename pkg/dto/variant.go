package dto

type OptionItem struct {
	OptionId     uint    `json:"option_id"`
	Value        string  `json:"value"`
}

type VariantGroup struct {
	VariantId   uint  					`json:"variant_id"`
	VariantName string  				`json:"variant_name"`
	Options     []OptionItem    `json:"options"`
}

type Variant struct {
	  VariantID   uint            `json:"id"`
	  Name        string          `json:"name"`
	  ProductID   uint            `json:"product_id"`
	  Options     []OptionItem    `json:"options"`
}


//TIDY THS UP.  VARINAT RESPONSE SHOULD BE vARIANTgROUP
type ProductVariantDetail struct {
  ID  			uint    					`json:"id"`
  Name 			string   					`json:"name"`
  Variants  []VariantResponse `json:"variants"`
}
