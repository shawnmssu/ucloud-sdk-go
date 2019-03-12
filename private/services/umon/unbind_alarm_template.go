//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMon UnbindAlarmTemplate

package umon

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UnbindAlarmTemplateRequest is request schema for UnbindAlarmTemplate action
type UnbindAlarmTemplateRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 资源类型，同DescribeResourceMetric支持的类型，请参考DescribeResourceMetric中的可选资源类型
	ResourceType *string `required:"true"`

	// 资源id
	ResourceId []string `required:"true"`
}

// UnbindAlarmTemplateResponse is response schema for UnbindAlarmTemplate action
type UnbindAlarmTemplateResponse struct {
	response.CommonBase
}

// NewUnbindAlarmTemplateRequest will create request of UnbindAlarmTemplate action.
func (c *UMonClient) NewUnbindAlarmTemplateRequest() *UnbindAlarmTemplateRequest {
	req := &UnbindAlarmTemplateRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// UnbindAlarmTemplate - 解绑告警模板
func (c *UMonClient) UnbindAlarmTemplate(req *UnbindAlarmTemplateRequest) (*UnbindAlarmTemplateResponse, error) {
	var err error
	var res UnbindAlarmTemplateResponse

	err = c.client.InvokeAction("UnbindAlarmTemplate", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
