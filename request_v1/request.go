package request_v1

import (
	"github.com/loudbund/go-json/json_v1"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// 接口请求1： Get接口请求
func Get(sUrl string, opt ...map[string]string) (int, string, error) {
	// 1、接口请求配置KeyReq
	var KeyReq *http.Request
	if req, err := http.NewRequest("GET", sUrl, nil); err != nil {
		log.Error(err)
		return 0, "", err
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		// req.Header.Set("Cookie", "name=anny")
		if len(opt) > 0 {
			for k, v := range opt[0] {
				req.Header.Set(k, v)
			}
		}
		KeyReq = req
	}

	// 2、发起请求获得结果KeyResponse
	var KeyResponse *http.Response
	if resp, err := (&http.Client{}).Do(KeyReq); err != nil {
		log.Error(err)
		return 0, "", err
	} else {
		defer func() { _ = resp.Body.Close() }()
		KeyResponse = resp
	}

	// 3、请求返回内容KeyBody
	KeyBody := make([]byte, 0)
	if body, err := ioutil.ReadAll(KeyResponse.Body); err != nil {
		log.Error(err)
		return -1, "", err
	} else {
		KeyBody = body
	}

	// 返回数据
	return KeyResponse.StatusCode, string(KeyBody), nil
}

// 接口请求2：Post的form格式参数请求
func PostForm(sUrl string, params map[string]string, opt ...map[string]string) (int, string, error) {
	// 1、参数拼接KeyP
	KeyP := url.Values{}
	for k, v := range params {
		KeyP.Set(k, v)
	}

	// 2、接口请求配置KeyReq
	var KeyReq *http.Request
	if req, err := http.NewRequest("POST", sUrl, strings.NewReader(KeyP.Encode())); err != nil {
		log.Error(err)
		return 0, "", err
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(KeyP.Encode())))
		// req.Header.Set("Cookie", "name=anny")
		if len(opt) > 0 {
			for k, v := range opt[0] {
				req.Header.Set(k, v)
			}
		}
		KeyReq = req
	}

	// 3、发起请求获得结果KeyResponse
	var KeyResponse *http.Response
	if resp, err := (&http.Client{}).Do(KeyReq); err != nil {
		log.Error(err)
		return 0, "", err
	} else {
		defer func() { _ = resp.Body.Close() }()
		KeyResponse = resp
	}

	// 4、请求返回内容KeyBody
	KeyBody := make([]byte, 0)
	if body, err := ioutil.ReadAll(KeyResponse.Body); err != nil {
		log.Error(err)
		return -1, "", err
	} else {
		KeyBody = body
	}

	// 返回数据
	return KeyResponse.StatusCode, string(KeyBody), nil
}

// 接口请求1：Post的json格式参数请求
func PostJson(sUrl string, params interface{}, opt ...map[string]string) (int, string, error) {
	KeyReqBody, _ := json_v1.JsonEncode(params)

	// 2、接口请求配置KeyReq
	var KeyReq *http.Request
	if req, err := http.NewRequest("POST", sUrl, strings.NewReader(KeyReqBody)); err != nil {
		return 0, "", err
	} else {
		req.Header.Set("Content-Type", "application/json")
		// req.Header.Set("Cookie", "name=anny")
		if len(opt) > 0 {
			for k, v := range opt[0] {
				req.Header.Set(k, v)
			}
		}
		KeyReq = req
	}

	// 3、发起请求获得结果KeyResponse
	var KeyResponse *http.Response
	if resp, err := (&http.Client{}).Do(KeyReq); err != nil {
		log.Error(err)
		return 0, "", err
	} else {
		defer func() { _ = resp.Body.Close() }()
		KeyResponse = resp
	}

	// 4、请求返回内容KeyBody
	KeyBody := make([]byte, 0)
	if body, err := ioutil.ReadAll(KeyResponse.Body); err != nil {
		log.Error(err)
		return -1, "", err
	} else {
		KeyBody = body
	}

	// 返回数据
	return KeyResponse.StatusCode, string(KeyBody), nil
}
