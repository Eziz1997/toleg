package admin

type Merchant struct {
	OrderNumber int64  `json:"orderNumber"`
	Currency    string `json:"currency"`
	Language    string `json:"language"`
	Password    string `json:"password"`
	ReturnUrl   string `json:"returnUrl"`
	UserName    string `json:"userName"`
}
