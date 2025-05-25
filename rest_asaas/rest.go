package rest_asaas

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/pericles-luz/go-asaas/model"
	"github.com/pericles-luz/go-base/pkg/utils"
)

var (
	ErrMissingEngine       = errors.New("missing rest engine")
	ErrSubscriptionFailed  = errors.New("subscription failed")
	ErrPlanCreationFailed  = errors.New("plan creation failed")
	ErrPlanNotFound        = errors.New("plan not found")
	ErrCardCreationFailed  = errors.New("card creation failed")
	ErrCardRetrievalFailed = errors.New("card retrieval failed")

	ErrAuthenticationRequired   = errors.New("authentication required")
	ErrMissingAutenticationData = errors.New("missing authentication data")
	ErrSubscriptionIDIsRequired = errors.New("subscription id is required")
	ErrCustomerCreationFailed   = errors.New("customer creation failed")
	ErrCustomerNotFound         = errors.New("customer not found")
)

type IResponse interface {
	GetCode() int
	GetRaw() string
}

type IToken interface {
	SetValidity(validity string) error
	SetKey(key string)
	IsValid() bool
	GetValidity() string
	GetKey() string
}

type IEngine interface {
	SetToken(token *Token) error
	NeedAutenticate() bool
	PostWithHeaderNoAuth(payload map[string]interface{}, link string, header map[string]string) (IResponse, error)
	GetWithHeaderNoAuth(payload map[string]interface{}, link string, header map[string]string) (IResponse, error)
	DeleteWithHeaderNoAuth(link string, header map[string]string) (IResponse, error)
}

type Rest struct {
	engine   IEngine
	baseLink string

	credential *model.Credential
}

func NewRest(engine IEngine, credentials string) (*Rest, error) {
	if engine == nil {
		return nil, ErrMissingEngine
	}
	configJSON, err := os.ReadFile(credentials)
	if err != nil {
		return nil, err
	}
	credential := model.NewCredential()
	if err := credential.Unmarshal(configJSON); err != nil {
		return nil, err
	}
	return &Rest{
		engine:     engine,
		credential: credential,
		baseLink:   credential.Link,
	}, nil
}

func (r *Rest) SetBaseLink(baseLink string) {
	r.baseLink = baseLink
}

func (r *Rest) getLink(link string) string {
	return r.baseLink + link
}

func (r *Rest) SetToken(token *Token) error {
	return r.engine.SetToken(token)
}

func (r *Rest) showResponse(result IResponse) {
	if result == nil {
		fmt.Println("Result is nil")
		return
	}
	fmt.Println("Code: ", result.GetCode())
	fmt.Println("Raw: ", result.GetRaw())
}

func (r *Rest) Authenticate() error {
	if r.engine == nil {
		return ErrMissingEngine
	}
	if !r.engine.NeedAutenticate() {
		return nil
	}
	if r.credential == nil {
		return ErrMissingAutenticationData
	}
	if err := r.credential.Validate(); err != nil {
		return err
	}
	token := NewToken(r.credential.AccessToken, 60)
	return r.SetToken(token)
}

func (r *Rest) CreateCustomer(customer *model.Customer) (*model.Customer, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}
	if err := customer.Validate(); err != nil {
		return nil, err
	}
	if r.engine.NeedAutenticate() {
		return nil, ErrAuthenticationRequired
	}
	fmt.Println("Customer: ", string(utils.MapInterfaceToBytes(customer.ToMap())))
	fmt.Println("Link: ", r.getLink("/v3/customers"))
	result, err := r.engine.PostWithHeaderNoAuth(customer.ToMap(), r.getLink("/v3/customers"), map[string]string{
		"access_token": r.credential.AccessToken,
		"accept":       "application/json",
		"content-type": "application/json",
	})
	if err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	if result.GetCode() != http.StatusOK {
		r.showResponse(result)
		errResponse, err := NewErrorResponse([]byte(result.GetRaw()))
		if err != nil {
			return nil, ErrCustomerCreationFailed
		}
		if errResponse.HasErrors() {
			fmt.Println("Error: ", errResponse.String())
			return nil, errResponse.Return()
		}
		return nil, ErrCustomerCreationFailed
	}
	customerResponse := model.NewCustomer()
	if err := customerResponse.Unmarshal([]byte(result.GetRaw())); err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	return customerResponse, nil
}

func (r *Rest) GetCustomer(customerID string) (*model.Customer, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}
	if customerID == "" {
		return nil, model.ErrCustomerIDIsRequired
	}
	result, err := r.engine.GetWithHeaderNoAuth(nil, r.getLink("/v3/customers/"+customerID), map[string]string{
		"access_token": r.credential.AccessToken,
		"accept":       "application/json",
	})
	if err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	if result.GetCode() != http.StatusOK {
		r.showResponse(result)
		errResponse, err := NewErrorResponse([]byte(result.GetRaw()))
		if err != nil {
			return nil, ErrCustomerCreationFailed
		}
		if errResponse.HasErrors() {
			fmt.Println("Error: ", errResponse.String())
			return nil, errResponse.Return()
		}
		return nil, ErrCustomerNotFound
	}
	customer := model.NewCustomer()
	if err := customer.Unmarshal([]byte(result.GetRaw())); err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	return customer, nil
}

func (r *Rest) ListCustomers(filter map[string]interface{}) (*model.CustomerList, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}
	result, err := r.engine.GetWithHeaderNoAuth(filter, r.getLink("/v3/customers"), map[string]string{
		"access_token": r.credential.AccessToken,
		"accept":       "application/json",
	})
	if err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	if result.GetCode() != http.StatusOK {
		r.showResponse(result)
		errResponse, err := NewErrorResponse([]byte(result.GetRaw()))
		if err != nil {
			return nil, ErrCustomerCreationFailed
		}
		if errResponse.HasErrors() {
			fmt.Println("Error: ", errResponse.String())
			return nil, errResponse.Return()
		}
		return nil, ErrCustomerNotFound
	}
	customers := model.NewCustomerList()
	if err := customers.Unmarshal([]byte(result.GetRaw())); err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	return customers, nil
}

func (r *Rest) Subscribe(subscription *model.Subscription) (*model.Subscription, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}
	if err := subscription.Validate(); err != nil {
		return nil, err
	}
	if r.engine.NeedAutenticate() {
		return nil, ErrAuthenticationRequired
	}
	fmt.Println("Subscription: ", string(utils.MapInterfaceToBytes(subscription.ToMap())))
	fmt.Println("Link: ", r.getLink("/v3/subscriptions"))
	result, err := r.engine.PostWithHeaderNoAuth(subscription.ToMap(), r.getLink("/v3/subscriptions"), map[string]string{
		"access_token": r.credential.AccessToken,
		"accept":       "application/json",
		"content-type": "application/json",
	})
	if err != nil {
		return nil, err
	}
	if result.GetCode() != http.StatusOK {
		r.showResponse(result)
		errResponse, err := NewErrorResponse([]byte(result.GetRaw()))
		if err != nil {
			return nil, ErrCustomerCreationFailed
		}
		if errResponse.HasErrors() {
			fmt.Println("Error: ", errResponse.String())
			return nil, errResponse.Return()
		}
		return nil, ErrSubscriptionFailed
	}
	subscriptionResponse := model.NewSubscription()
	if err := subscriptionResponse.Unmarshal([]byte(result.GetRaw())); err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	return subscriptionResponse, nil
}

func (r *Rest) GetSubscription(subscriptionID string) (*model.Subscription, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}
	if subscriptionID == "" {
		return nil, ErrSubscriptionIDIsRequired
	}
	result, err := r.engine.GetWithHeaderNoAuth(nil, r.getLink("/v3/subscriptions/"+subscriptionID), map[string]string{
		"access_token": r.credential.AccessToken,
		"accept":       "application/json",
	})
	if err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	if result.GetCode() != http.StatusOK {
		r.showResponse(result)
		errResponse, err := NewErrorResponse([]byte(result.GetRaw()))
		if err != nil {
			return nil, ErrCustomerCreationFailed
		}
		if errResponse.HasErrors() {
			fmt.Println("Error: ", errResponse.String())
			return nil, errResponse.Return()
		}
		return nil, ErrSubscriptionFailed
	}
	subscription := model.NewSubscription()
	if err := subscription.Unmarshal([]byte(result.GetRaw())); err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return nil, err
	}
	return subscription, nil
}

func (r *Rest) Unsubscribe(subscriptionID string) error {
	if err := r.Authenticate(); err != nil {
		return err
	}
	result, err := r.engine.DeleteWithHeaderNoAuth(r.getLink("/v3/subscriptions/"+subscriptionID), map[string]string{
		"access_token": r.credential.AccessToken,
		"accept":       "application/json",
	})
	if err != nil {
		r.showResponse(result)
		fmt.Println("Error: ", err)
		return err
	}
	if result.GetCode() != http.StatusOK {
		r.showResponse(result)
		return ErrSubscriptionFailed
	}
	return nil
}
