package factory_client_asaas

import "github.com/pericles-luz/go-asaas/pkg/rest_asaas"

func NewClient(configPath string) (*rest_asaas.Rest, error) {
	engine := rest_asaas.NewEngine(map[string]interface{}{"InsecureSkipVerify": true})
	restEntity, err := rest_asaas.NewRest(engine, configPath)
	if err != nil {
		return nil, err
	}
	return restEntity, nil
}
