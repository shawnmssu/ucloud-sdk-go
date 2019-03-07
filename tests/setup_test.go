package tests

import (
	"os"
	"testing"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"

	"github.com/ucloud/ucloud-sdk-go/services/pathx"
	"github.com/ucloud/ucloud-sdk-go/services/uaccount"
	"github.com/ucloud/ucloud-sdk-go/services/ubill"
	"github.com/ucloud/ucloud-sdk-go/services/udb"
	"github.com/ucloud/ucloud-sdk-go/services/udisk"
	"github.com/ucloud/ucloud-sdk-go/services/udpn"
	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/services/ulb"
	"github.com/ucloud/ucloud-sdk-go/services/umem"
	"github.com/ucloud/ucloud-sdk-go/services/unet"
	"github.com/ucloud/ucloud-sdk-go/services/uphost"
	"github.com/ucloud/ucloud-sdk-go/services/vpc"

	pudisk "github.com/ucloud/ucloud-sdk-go/private/services/udisk"
	puhost "github.com/ucloud/ucloud-sdk-go/private/services/uhost"
	pulb "github.com/ucloud/ucloud-sdk-go/private/services/ulb"
	pumem "github.com/ucloud/ucloud-sdk-go/private/services/umem"
	pumon "github.com/ucloud/ucloud-sdk-go/private/services/umon"

	iubill "github.com/ucloud/ucloud-sdk-go/internal/services/ubill"
	iudataark "github.com/ucloud/ucloud-sdk-go/internal/services/udataark"
	iudb "github.com/ucloud/ucloud-sdk-go/internal/services/udb"
	iudisk "github.com/ucloud/ucloud-sdk-go/internal/services/udisk"
	iudpn "github.com/ucloud/ucloud-sdk-go/internal/services/udpn"
	iuhost "github.com/ucloud/ucloud-sdk-go/internal/services/uhost"
	iulb "github.com/ucloud/ucloud-sdk-go/internal/services/ulb"
	iumem "github.com/ucloud/ucloud-sdk-go/internal/services/umem"
	iumon "github.com/ucloud/ucloud-sdk-go/internal/services/umon"
	iunet "github.com/ucloud/ucloud-sdk-go/internal/services/unet"
	iuphost "github.com/ucloud/ucloud-sdk-go/internal/services/uphost"
)

var config *ucloud.Config
var client *ucloud.Client
var uhostClient *uhost.UHostClient
var unetClient *unet.UNetClient
var ulbClient *ulb.ULBClient
var vpcClient *vpc.VPCClient
var uaccountClient *uaccount.UAccountClient
var pathxClient *pathx.PathXClient
var udiskClient *udisk.UDiskClient
var udbClient *udb.UDBClient
var umemClient *umem.UMemClient
var udpnClient *udpn.UDPNClient
var ubillClient *ubill.UBillClient
var uphostClient *uphost.UPHostClient

var puhostClient *puhost.UHostClient
var pudiskClient *pudisk.UDiskClient
var pumemClient *pumem.UMemClient
var pumonClient *pumon.UMonClient
var pulbClient *pulb.ULBClient

var iuhostClient *iuhost.UHostClient
var iunetClient *iunet.UNetClient
var iulbClient *iulb.ULBClient
var iudiskClient *iudisk.UDiskClient
var iumonClient *iumon.UMonClient
var iudataarkClient *iudataark.UDataArkClient
var iudbClient *iudb.UDBClient
var iumemClient *iumem.UMemClient
var iudpnClient *iudpn.UDPNClient
var iubillClient *iubill.UBillClient
var iuphostClient *iuphost.UPHostClient

func TestMain(m *testing.M) {
	testSetup()
	exitCode := m.Run()
	testTeardown()
	os.Exit(exitCode)
}

func testSetup() {
	cfg := ucloud.NewConfig()
	cfg.MaxRetries = 1
	cfg.LogLevel = log.DebugLevel
	cfg.Region = os.Getenv("UCLOUD_REGION")
	cfg.ProjectId = os.Getenv("UCLOUD_PROJECT_ID")

	config = &cfg

	credential := auth.NewCredential()
	credential.PrivateKey = os.Getenv("UCLOUD_PRIVATE_KEY")
	credential.PublicKey = os.Getenv("UCLOUD_PUBLIC_KEY")

	cfg.SetActionLevel("DescribeImage", log.WarnLevel)
	cfg.SetActionLevel("GetRegion", log.WarnLevel)
	cfg.SetActionLevel("DescribeUDBParamGroup", log.WarnLevel)

	client = ucloud.NewClient(&cfg, &credential)
	uhostClient = uhost.NewClient(&cfg, &credential)
	unetClient = unet.NewClient(&cfg, &credential)
	ulbClient = ulb.NewClient(&cfg, &credential)
	vpcClient = vpc.NewClient(&cfg, &credential)
	uaccountClient = uaccount.NewClient(&cfg, &credential)
	pathxClient = pathx.NewClient(&cfg, &credential)
	udiskClient = udisk.NewClient(&cfg, &credential)
	udbClient = udb.NewClient(&cfg, &credential)
	umemClient = umem.NewClient(&cfg, &credential)
	udpnClient = udpn.NewClient(&cfg, &credential)
	ubillClient = ubill.NewClient(&cfg, &credential)
	uphostClient = uphost.NewClient(&cfg, &credential)

	iudataarkClient = iudataark.NewClient(&cfg, &credential)
	iudiskClient = iudisk.NewClient(&cfg, &credential)
	iuhostClient = iuhost.NewClient(&cfg, &credential)
	iulbClient = iulb.NewClient(&cfg, &credential)
	iumonClient = iumon.NewClient(&cfg, &credential)
	iunetClient = iunet.NewClient(&cfg, &credential)
	iudbClient = iudb.NewClient(&cfg, &credential)
	iumemClient = iumem.NewClient(&cfg, &credential)
	iudpnClient = iudpn.NewClient(&cfg, &credential)
	iubillClient = iubill.NewClient(&cfg, &credential)
	iuphostClient = iuphost.NewClient(&cfg, &credential)

	pudiskClient = pudisk.NewClient(&cfg, &credential)
	puhostClient = puhost.NewClient(&cfg, &credential)
	pumemClient = pumem.NewClient(&cfg, &credential)
	pumonClient = pumon.NewClient(&cfg, &credential)
	pulbClient = pulb.NewClient(&cfg, &credential)

	log.Info("setup test fixtures ...")
}

func testTeardown() {}
