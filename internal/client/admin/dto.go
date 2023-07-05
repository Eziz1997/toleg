package admin

type RegisterDoRequest struct {
	Amount string `json:"amount"`
}

type RegisterDoResponse struct {
	OrderId      string `json:"orderId"`
	FormUrl      string `json:"formUrl"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type RegisterDoRequestToBank struct {
	UserName           string      `json:"userName"`
	Password           string      `json:"password"`
	OrderNumber        int         `json:"orderNumber"`
	Amount             string      `json:"amount"`
	Currency           string      `json:"currency"` // 934 (manat)
	ReturnUrl          string      `json:"returnUrl"`
	FailUrl            string      `json:"failUrl"`
	Description        string      `json:"description"`
	Language           string      `json:"language"` // ru
	PageView           string      `json:"pageView"` // default = DESKTOP (MOBILE)
	ClientId           string      `json:"clientId"`
	MerchantLogin      string      `json:"merchantLogin"`
	JsonParams         interface{} `json:"jsonParams"`
	SessionTimeoutSecs int         `json:"sessionTimeoutSecs"` // seconds (default = 1200)
	ExpirationDate     string      `json:"expirationDate"`     // Формат: yyyy-MM-ddTHH:mm:ss
	BindingId          string      `json:"bindingId"`
}
