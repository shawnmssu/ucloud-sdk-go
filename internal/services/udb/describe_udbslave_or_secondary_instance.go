//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB DescribeUDBSlaveOrSecondaryInstance

package udb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeUDBSlaveOrSecondaryInstanceRequest is request schema for DescribeUDBSlaveOrSecondaryInstance action
type DescribeUDBSlaveOrSecondaryInstanceRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`

	// 主DB实例，master或者primary节点
	DBId *string `required:"true"`

	// DB种类，分为SQL和NOSQL,取值分别为mysql、mongo
	ClassType *string `required:"true"`
}

// DescribeUDBSlaveOrSecondaryInstanceResponse is response schema for DescribeUDBSlaveOrSecondaryInstance action
type DescribeUDBSlaveOrSecondaryInstanceResponse struct {
	response.CommonBase

	// DB实例信息列表 UDBSlaveOrSecondarySet
	DataSet []UDBSlaveOrSecondarySet
}

// NewDescribeUDBSlaveOrSecondaryInstanceRequest will create request of DescribeUDBSlaveOrSecondaryInstance action.
func (c *UDBClient) NewDescribeUDBSlaveOrSecondaryInstanceRequest() *DescribeUDBSlaveOrSecondaryInstanceRequest {
	req := &DescribeUDBSlaveOrSecondaryInstanceRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUDBSlaveOrSecondaryInstance - 获取udb实例从库信息，指定ClassType（仅作为内部api使用）
func (c *UDBClient) DescribeUDBSlaveOrSecondaryInstance(req *DescribeUDBSlaveOrSecondaryInstanceRequest) (*DescribeUDBSlaveOrSecondaryInstanceResponse, error) {
	var err error
	var res DescribeUDBSlaveOrSecondaryInstanceResponse

	err = c.client.InvokeAction("DescribeUDBSlaveOrSecondaryInstance", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
