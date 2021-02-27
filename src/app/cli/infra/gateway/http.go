package gateway

import (
	"bytes"
	"io/ioutil"
	httpAgent "net/http"
)

type (
	http struct {
		headers map[string]string
		baseUrl string
	}
)

func NewHttp(baseUrl string, headers map[string]string) http {
	return http{
		headers: headers,
		baseUrl: baseUrl,
	}
}

func (e http) Get(uri string, params map[string]string) ([]byte, error) {
	fullURL := e.baseUrl + "/" + uri
	req, _ := httpAgent.NewRequest("GET", fullURL, nil)

	if len(e.headers) > 0 {
		setQueryString(req, e.headers)
	}

	if len(params) > 0 {
		setQueryString(req, params)
	}

	client := new(httpAgent.Client)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return b, err
	}

	return b, nil
}

func (e http) Post(uri string, body string, params map[string]string) ([]byte, error) {
	fullURL := e.baseUrl + "/" + uri
	// req, _ := httpAgent.NewRequest("GET", fullURL, nil)
	req, _ := httpAgent.NewRequest("POST", fullURL, bytes.NewBuffer([]byte(body)))

	if len(e.headers) > 0 {
		setQueryString(req, e.headers)
	}

	if len(params) > 0 {
		setQueryString(req, params)
	}

	client := new(httpAgent.Client)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return b, err
	}

	return b, nil
}

// func (e ecroboRepository) CreateOrders(orders []domain.Order) (domain.OrderResponse, error) {
// 	res := domain.OrderResponse{}
// 	o := postOrdersResponse{}
// 	apiPath := "order/"
// 	fullURL := fmt.Sprintf("%s/%s", e.URL, apiPath)
// 	payload, err := e.convertToPostOrdersPayload(orders)
// 	if err != nil {
// 		return res, err
// 	}

// 	values, err := json.Marshal(payload)
// 	if err != nil {
// 		return res, err
// 	}ggo

// 	req, _ := http.NewRequest("POST", fullURL, bytes.NewBuffer(values))
// 	req.Header.Set("X-API-KEY", e.AccessKey)
// 	req.Header.Set("Content-Type", "application/json")
// 	client := new(http.Client)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return res, err
// 	}
// 	defer resp.Body.Close()
// 	b, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return res, err
// 	}
// 	log.Printf(string(b))
// 	if err := json.Unmarshal(b, &o); err != nil {
// 		return res, err
// 	}
// 	res = domain.OrderResponse{
// 		Result:        o.Result,
// 		TransactionID: o.TransactionID,
// 		TaskID:        o.TaskID,
// 	}

// 	return res, nil
// }

func setQueryString(req *httpAgent.Request, params map[string]string) {
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
}

func setHeaders(req *httpAgent.Request, params map[string]string) {
	for k, v := range params {
		req.Header.Set(k, v)
	}
}
