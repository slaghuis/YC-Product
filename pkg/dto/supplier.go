package dto
//Data Transfer Objects

type SupplierBasic struct {
    ID              uint    `json:"id"`
    Name            string  `json:"name"`
    RepName         string  `json:"rep_name"`
    LogoPath        string  `json:"logo_path"`
    Phone           string  `json:"phone"`
    PhoneVerified   bool    `json:"phone_verified"`
    Email           string  `json:"email"`
    SMS             bool    `json:"sms"`
    WhatsApp        bool    `json:"whatsapp"`
    EmailNotify     bool    `json:"email_notify"`
    AppPush         bool    `json:"app_push"`
    ProductCount    int     `json:"product_count"`
    FilterSelected  bool    `json:"filter_selected"`   //Go will fill this with a false by default
}

type SupplierMerchantListItem struct {
		ID          			uint           			`json:"id"`
    Name        			string         			`json:"name"`
    RepName 			    string         			`json:"rep_name"`
		LogoPath     		  string 							`json:"logo_path"`
    ProductCount      int                 `json:"product_count"`
		Likes             int                 `json:"like_count"`
		Reviews           int                 `json:"review_count"`
		Sales							int                 `json:"sales"`
		Earnings					float64							`json:"earnings"`
}
