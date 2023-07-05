package admin

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"toleg/internal/appresult"
	"toleg/internal/handlers"
	"toleg/pkg/logging"
)

const (
	registerDoURL = "/register-do"
	resultURL     = "/result-register-do"
)

type handler struct {
	logger     *logging.Logger
	repository Repository
}

func NewHandler(repository Repository, logger *logging.Logger) handlers.Handler {
	return &handler{
		repository: repository,
		logger:     logger,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc(registerDoURL, appresult.Middleware(h.RegisterDo)).Methods("POST")
	router.HandleFunc(resultURL, appresult.Middleware(h.ResultRegister)).Methods("GET")
}

func (h *handler) ResultRegister(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("---->>", r.URL.Query())

	var jsonMap map[string]interface{}

	payload := url.Values{}
	payload.Add("orderId", r.URL.Query().Get("orderId"))
	payload.Add("currency", "934")
	payload.Add("language", "ru")
	payload.Add("userName", "102211004516")
	payload.Add("password", "Jt3dsRvgTPdfPfA")

	fmt.Println("Hello World!")

	req, err := http.NewRequest("POST", "https://mpi.gov.tm/payment/rest/getOrderStatus.do?"+payload.Encode(), nil)

	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	res, errR := client.Do(req)
	if errR != nil {
		//h.logger.Error("res pay request-de error bbar", errR)
		fmt.Println(errR)
	}

	cBody, err := io.ReadAll(res.Body)
	res.Body.Close()

	Resp := string(cBody)

	json.Unmarshal([]byte(Resp), &jsonMap)

	OrderStatus, _ := strconv.Atoi(fmt.Sprintf("%v", jsonMap[("OrderStatus")]))

	fmt.Println("assssssss", string(cBody))

	fmt.Println("assssssss", OrderStatus)

	successResult := appresult.Success
	successResult.Data = jsonMap // "result order status: " + strconv.Itoa(OrderStatus)

	w.Header().Add(appresult.HeaderContentTypeJson())
	err = json.NewEncoder(w).Encode(successResult)
	if err != nil {
		return err
	}

	return nil
}

func (h *handler) RegisterDo(w http.ResponseWriter, r *http.Request) error {

	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		return appresult.ErrMissingParam
	}

	registerDoRequest := RegisterDoRequest{}
	errData := json.Unmarshal(body, &registerDoRequest)

	if errData != nil {
		return appresult.ErrMissingParam
	}

	// post to bank
	successResult := appresult.Success
	successResult.Data = postToBank(registerDoRequest.Amount)

	w.Header().Add(appresult.HeaderContentTypeJson())
	err := json.NewEncoder(w).Encode(successResult)
	if err != nil {
		return err
	}

	return nil
}

func postToBank(amount string) RegisterDoResponse {
	postUrl := "https://mpi.gov.tm/payment/rest/register.do?"

	//registerDoRequestToBank := RegisterDoRequestToBank{
	//	UserName:    "102211004516",
	//	Password:    "Jt3dsRvgTPdfPfA",
	//	ReturnUrl:   "https://salamnews.tm",
	//	PageView:    "MOBILE",
	//	Description: "nailidir bir desc",
	//	Language:    "ru",
	//	Currency:    "934",
	//	Amount:      amount,
	//	OrderNumber: 1,
	//}
	fmt.Println("sssss")
	params := url.Values{}
	params.Add("orderNumber", "9358")
	params.Add("amount", amount)

	params.Add("currency", "934")
	params.Add("language", "ru")
	params.Add("password", "Jt3dsRvgTPdfPfA")
	params.Add("returnUrl", "http://localhost:3003/api/altynasyr/result-register-do")

	params.Add("userName", "102211004516")
	params.Add("description", "-ddddddd-")

	resp, err := http.PostForm(postUrl, params)
	if err != nil {
		fmt.Printf("Request Failed: %s", err)
	}
	fmt.Println("sssss")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// Log the request body
	bodyString := string(body)
	fmt.Println(bodyString)
	// Unmarshal result
	result := RegisterDoResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("Reading body failed: %s", err)
	}

	return result
}
