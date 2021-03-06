//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem ModifyURedisGroupPassword

package umem

import (
	"encoding/base64"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// ModifyURedisGroupPasswordRequest is request schema for ModifyURedisGroupPassword action
type ModifyURedisGroupPasswordRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`

	// 组的ID
	GroupId *string `required:"true"`

	// 新密码字符串，如要取消密码，此值为空字符串
	Password *string `required:"true"`

	//
	ResourceType *string `required:"false"`
}

// ModifyURedisGroupPasswordResponse is response schema for ModifyURedisGroupPassword action
type ModifyURedisGroupPasswordResponse struct {
	response.CommonBase
}

// NewModifyURedisGroupPasswordRequest will create request of ModifyURedisGroupPassword action.
func (c *UMemClient) NewModifyURedisGroupPasswordRequest() *ModifyURedisGroupPasswordRequest {
	req := &ModifyURedisGroupPasswordRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// ModifyURedisGroupPassword - 修改主备密码/重置密码
func (c *UMemClient) ModifyURedisGroupPassword(req *ModifyURedisGroupPasswordRequest) (*ModifyURedisGroupPasswordResponse, error) {
	var err error
	var res ModifyURedisGroupPasswordResponse
	req.Password = ucloud.String(base64.StdEncoding.EncodeToString([]byte(ucloud.StringValue(req.Password))))

	err = c.client.InvokeAction("ModifyURedisGroupPassword", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
