package http

//http request 工具函数
import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var httpProxyClient = &http.Client{
	Transport: &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return &url.URL{
				Scheme: "socks5",
				Host:   "127.0.0.1:1086"}, nil
		},
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
	},
	Timeout: 10 * time.Second,
}

func NewHttpRequest(reqType string, reqUrl string, postData string, requstHeaders map[string]string) ([]byte, error) {
	req, _ := http.NewRequest(reqType, reqUrl, strings.NewReader(postData))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36")

	if requstHeaders != nil {
		for k, v := range requstHeaders {
			req.Header.Add(k, v)
		}
	}

	resp, err := httpProxyClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("HttpStatusCode:%d ,Desc:%s", resp.StatusCode, string(bodyData)))
	}

	return bodyData, nil
}

func HttpGet(reqUrl string) (map[string]interface{}, error) {
	respData, err := NewHttpRequest("GET", reqUrl, "", nil)
	if err != nil {
		return nil, err
	}

	var bodyDataMap map[string]interface{}
	err = json.Unmarshal(respData, &bodyDataMap)
	if err != nil {
		log.Println(string(respData))
		return nil, err
	}
	return bodyDataMap, nil
}

func HttpGet2(reqUrl string, headers map[string]string) (map[string]interface{}, error) {
	if headers == nil {
		headers = map[string]string{}
	}
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	respData, err := NewHttpRequest("GET", reqUrl, "", headers)
	if err != nil {
		return nil, err
	}

	var bodyDataMap map[string]interface{}
	err = json.Unmarshal(respData, &bodyDataMap)
	if err != nil {
		log.Println("respData", string(respData))
		return nil, err
	}
	return bodyDataMap, nil
}

func HttpGet3(reqUrl string, headers map[string]string) ([]interface{}, error) {
	if headers == nil {
		headers = map[string]string{}
	}
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	respData, err := NewHttpRequest("GET", reqUrl, "", headers)
	if err != nil {
		return nil, err
	}

	var bodyDataMap []interface{}
	err = json.Unmarshal(respData, &bodyDataMap)
	if err != nil {
		log.Println("respData", string(respData))
		return nil, err
	}
	return bodyDataMap, nil
}

func HttpGet4(reqUrl string, headers map[string]string, result interface{}) error {
	if headers == nil {
		headers = map[string]string{}
	}
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	respData, err := NewHttpRequest("GET", reqUrl, "", headers)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respData, result)
	if err != nil {
		log.Printf("HttpGet4 - json.Unmarshal failed : %v, resp %s", err, string(respData))
		return err
	}

	return nil
}
func HttpGet5(reqUrl string, headers map[string]string) ([]byte, error) {
	if headers == nil {
		headers = map[string]string{}
	}
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	respData, err := NewHttpRequest("GET", reqUrl, "", headers)
	if err != nil {
		return nil, err
	}

	return respData, nil
}

func HttpPostForm(reqUrl string, postData url.Values) ([]byte, error) {
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded"}
	return NewHttpRequest("POST", reqUrl, postData.Encode(), headers)
}

func HttpPostForm2(reqUrl string, postData url.Values, headers map[string]string) ([]byte, error) {
	if headers == nil {
		headers = map[string]string{}
	}
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	return NewHttpRequest("POST", reqUrl, postData.Encode(), headers)
}

func HttpPostForm3(reqUrl string, postData string, headers map[string]string) ([]byte, error) {
	return NewHttpRequest("POST", reqUrl, postData, headers)
}

func HttpPostForm4(reqUrl string, postData map[string]string, headers map[string]string) ([]byte, error) {
	if headers == nil {
		headers = map[string]string{}
	}
	headers["Content-Type"] = "application/json"
	data, _ := json.Marshal(postData)
	return NewHttpRequest("POST", reqUrl, string(data), headers)
}

func HttpDeleteForm(reqUrl string, postData url.Values, headers map[string]string) ([]byte, error) {
	if headers == nil {
		headers = map[string]string{}
	}
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	return NewHttpRequest("DELETE", reqUrl, postData.Encode(), headers)
}
