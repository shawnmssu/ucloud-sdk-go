// Code is generated by ucloud-model, DO NOT EDIT IT.

package ipsecvpn

/*
VPNTunnelDataSet - DescribeVPNTunnel信息

this model is auto created by ucloud code generater for open api,
you can also see https://docs.ucloud.cn/api for detail.
*/
type VPNTunnelDataSet struct {

	// 创建时间
	CreateTime int

	// IKE参数
	IKEData IKEData

	// IPSec参数
	IPSecData IPSecData

	// 备注
	Remark string

	// 对端网关Id
	RemoteVPNGatewayId string

	// 对端网关名字
	RemoteVPNGatewayName string

	// 用户组
	Tag string

	// 所属VPCId
	VPCId string

	// 所属VOC名字
	VPCName string

	// 所属VPN网关id
	VPNGatewayId string

	// VPN网关名字
	VPNGatewayName string

	// 隧道id
	VPNTunnelId string

	// 隧道名称
	VPNTunnelName string
}