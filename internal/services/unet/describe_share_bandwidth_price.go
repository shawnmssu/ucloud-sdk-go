//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UNet DescribeShareBandwidthPrice

package unet

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeShareBandwidthPriceRequest is request schema for DescribeShareBandwidthPrice action
type DescribeShareBandwidthPriceRequest struct {
	request.CommonBase

	// 付费方式, 预付费:Year 按年,Month 按月,Dynamic 按需;后付费:Postpay(按月)
	ChargeType *string `required:"true"`

	// 共享带宽值(预付费)/共享带宽峰值(后付费)
	ShareBandwidth *int `required:"true"`

	// 购买时长(预付费)
	Quantity *int `required:"false"`

	// 共享带宽保底值(后付费)
	ShareBandwidthGuarantee *int `required:"false"`
}

// DescribeShareBandwidthPriceResponse is response schema for DescribeShareBandwidthPrice action
type DescribeShareBandwidthPriceResponse struct {
	response.CommonBase

	// 共享带宽总价
	TotalPrice int
}

// NewDescribeShareBandwidthPriceRequest will create request of DescribeShareBandwidthPrice action.
func (c *UNetClient) NewDescribeShareBandwidthPriceRequest() *DescribeShareBandwidthPriceRequest {
	req := &DescribeShareBandwidthPriceRequest{}
	c.client.SetupRequest(req)
	return req
}

// DescribeShareBandwidthPrice - 获取共享带宽价格
func (c *UNetClient) DescribeShareBandwidthPrice(req *DescribeShareBandwidthPriceRequest) (*DescribeShareBandwidthPriceResponse, error) {
	var err error
	var res DescribeShareBandwidthPriceResponse

	err = c.client.InvokeAction("DescribeShareBandwidthPrice", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
