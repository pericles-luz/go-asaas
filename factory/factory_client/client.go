package factory_client

import "github.com/pericles-luz/go-asaas/rest"

func NewClient(configPath string) (*rest.Rest, error) {
	engine := rest.NewEngine(map[string]interface{}{"InsecureSkipVerify": true})
	restEntity, err := rest.NewRest(engine, configPath)
	if err != nil {
		return nil, err
	}
	return restEntity, nil
}
