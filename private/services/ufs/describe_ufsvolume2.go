//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UFS DescribeUFSVolume2

package ufs

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeUFSVolume2Request is request schema for DescribeUFSVolume2 action
type DescribeUFSVolume2Request struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 文件系统ID
	VolumeId *string `required:"false"`

	// 文件列表起始
	Offset *int `required:"false"`

	// 文件列表长度
	Limit *int `required:"false"`

	// 枚举值，是否拉取性能型文件系统列表，1表示只拉取容量型文件系统列表，默认为0
	OnlyBasic *int `required:"false"`
}

// DescribeUFSVolume2Response is response schema for DescribeUFSVolume2 action
type DescribeUFSVolume2Response struct {
	response.CommonBase

	// 文件系统总数
	TotalCount int

	// 文件系统详细信息列表
	DataSet []UFSVolumeInfo2
}

// NewDescribeUFSVolume2Request will create request of DescribeUFSVolume2 action.
func (c *UFSClient) NewDescribeUFSVolume2Request() *DescribeUFSVolume2Request {
	req := &DescribeUFSVolume2Request{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUFSVolume2 - 获取文件系统列表
func (c *UFSClient) DescribeUFSVolume2(req *DescribeUFSVolume2Request) (*DescribeUFSVolume2Response, error) {
	var err error
	var res DescribeUFSVolume2Response

	err = c.Client.InvokeAction("DescribeUFSVolume2", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}