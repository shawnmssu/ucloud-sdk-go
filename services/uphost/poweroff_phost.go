//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UPHost PoweroffPHost

package uphost

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// PoweroffPHostRequest is request schema for PoweroffPHost action
type PoweroffPHostRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`

	// PHost资源ID
	PHostId *string `required:"true"`
}

// PoweroffPHostResponse is response schema for PoweroffPHost action
type PoweroffPHostResponse struct {
	response.CommonBase

	// PHost 的资源ID
	PHostId string
}

// NewPoweroffPHostRequest will create request of PoweroffPHost action.
func (c *UPHostClient) NewPoweroffPHostRequest() *PoweroffPHostRequest {
	req := &PoweroffPHostRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// PoweroffPHost - 断电物理云主机
func (c *UPHostClient) PoweroffPHost(req *PoweroffPHostRequest) (*PoweroffPHostResponse, error) {
	var err error
	var res PoweroffPHostResponse

	err = c.client.InvokeAction("PoweroffPHost", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
