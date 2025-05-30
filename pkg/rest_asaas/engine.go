package rest_asaas

// Import resty into your code and refer it as `resty`.
import (
	"crypto/tls"
	"errors"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type Engine struct {
	http   *resty.Client
	token  *Token
	config map[string]interface{}
}

func (e *Engine) getHttp() *resty.Client {
	return e.http
}

func (e *Engine) getToken() (*Token, error) {
	if e.token == nil {
		return nil, errors.New("missing authentication token")
	}
	if !e.token.IsValid() {
		e.token = nil
		return nil, errors.New("invalid authentication token")
	}
	return e.token, nil
}

// defines a token to be used in the requests
func (e *Engine) SetToken(token *Token) error {
	if !token.IsValid() {
		return errors.New("token is invalid")
	}
	e.token = token
	return nil
}

// sets variables used in the requests
func (e *Engine) SetConfig(key string, value string) {
	e.config[key] = value
}

// gets variables used in the requests
func (e *Engine) GetConfig(key string) string {
	return e.config[key].(string)
}

// sets variables used in the requests
func (e *Engine) GetConfigData() map[string]interface{} {
	return e.config
}

// posts request to the given link, without token and specific header
func (e *Engine) PostWithHeaderNoAuth(payload map[string]interface{}, link string, header map[string]string) (IResponse, error) {
	resp, err := e.getHttp().R().SetBody(payload).SetHeaders(header).Post(link)
	if err != nil {
		return nil, err
	}
	resp.Time()
	return &Response{
		code: resp.StatusCode(),
		raw:  resp.String(),
	}, nil
}

// gets request to the given link, without token and specific header
func (e *Engine) GetWithHeaderNoAuth(payload map[string]interface{}, link string, header map[string]string) (IResponse, error) {
	data := e.preparePayload(payload)
	resp, err := e.getHttp().R().SetQueryParams(data).SetHeaders(header).Get(link)
	if err != nil {
		return nil, err
	}
	resp.Time()
	return &Response{
		code: resp.StatusCode(),
		raw:  resp.String(),
	}, nil
}

// deletes request to the given link, using the defined token and specific header without authentication
func (e *Engine) DeleteWithHeaderNoAuth(link string, header map[string]string) (IResponse, error) {
	resp, err := e.getHttp().R().SetHeaders(header).Delete(link)
	if err != nil {
		return nil, err
	}
	resp.Time()
	return &Response{
		code: resp.StatusCode(),
		raw:  resp.String(),
	}, nil
}

func (e *Engine) preparePayload(payload map[string]interface{}) map[string]string {
	result := map[string]string{}
	for k, v := range payload {
		switch t := v.(type) {
		case string:
			result[k] = v.(string)
		case bool:
			if v.(bool) {
				result[k] = "true"
				continue
			}
			result[k] = "false"
		default:
			result[k] = fmt.Sprintf("%v", t)
		}
	}
	return result
}

func (e *Engine) NeedAutenticate() bool {
	_, err := e.getToken()
	return err != nil
}

// gets a Rest struct with the given config
// if InsecureSkipVerify is set to true, the client will skip the verification of the server's certificate
func NewEngine(config map[string]interface{}) *Engine {
	client := resty.New()
	if config["InsecureSkipVerify"] != nil && config["InsecureSkipVerify"].(bool) {
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: config["InsecureSkipVerify"].(bool)})
	}
	engine := &Engine{
		http:   client,
		config: config,
		token:  &Token{},
	}
	engine.http.SetHeaders(map[string]string{
		"Content-Type": "application/json",
	})
	engine.getHttp().SetTimeout(1 * time.Minute)
	return engine
}
