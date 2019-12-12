// Code is generated by ucloud-model, DO NOT EDIT IT.

package driver

import (
	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

type serviceClient interface {
	AddResponseHandler(ucloud.ResponseHandler) error
}

type serviceFactory func(interface{}, interface{}) serviceClient

// newServiceClient will build a service Client
func newServiceClient(product string, cfg interface{}, cred interface{}) serviceClient {
	for k, v := range serviceFactoryMap {
		if k == product {
			c := v(cfg, cred)
			return c
		} else if len(product) == 0 {
			c := ucloud.NewClient(cfg.(*ucloud.Config), cred.(*auth.Credential))
			return c
		}
	}

	return nil
}

var serviceFactoryMap = map[string]serviceFactory{
	"UHost": func(cfg interface{}, cred interface{}) serviceClient {
		return uhost.NewClient(cfg.(*ucloud.Config), cred.(*auth.Credential))
	},
}
