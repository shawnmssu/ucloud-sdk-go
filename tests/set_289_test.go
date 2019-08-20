// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet289(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("HostName", "auto_host_test3")
	ctx.SetVar("Password", "Z3VhbmxpeXVhbm1pbWExMjM=")
	ctx.SetVar("ChargeType", "Month")
	ctx.SetVar("CreateCPU", 1)
	ctx.SetVar("CreateMem", 1024)
	ctx.SetVar("ImageId", "#{u_get_image_resource($Region,$Zone)}")
	ctx.SetVar("BootSize", 20)
	ctx.SetVar("BootType", "CLOUD_SSD")
	ctx.SetVar("DiskSize", 20)
	ctx.SetVar("DiskType", "CLOUD_NORMAL")
	ctx.SetVar("BootBackup", "NONE")
	ctx.SetVar("DiskBackup", "NONE")
	ctx.SetVar("UHostType", "N2")
	ctx.SetVar("UDiskType", "DataDisk")
	ctx.SetVar("Size", 1)
	ctx.SetVar("UDataArkMode", "No")
	ctx.SetVar("UDiskName", "auto_udisk_noArk")

	testSet289DescribeUDiskPrice00(&ctx)
	testSet289CheckUDiskAllowance01(&ctx)
	testSet289CreateUDisk02(&ctx)
	testSet289DescribeUDisk03(&ctx)
	testSet289CreateUDiskSnapshot04(&ctx)
	testSet289DescribeUDisk05(&ctx)
	testSet289DescribeSnapshot06(&ctx)
	testSet289CreateUDiskSnapshot07(&ctx)
	testSet289DescribeSnapshot08(&ctx)
	testSet289CreateUDiskSnapshot09(&ctx)
	testSet289DescribeUDiskSnapshot10(&ctx)
	testSet289CreateUDiskSnapshot11(&ctx)
	testSet289CloneUDiskSnapshot12(&ctx)
	testSet289CloneUDiskSnapshot13(&ctx)
	testSet289RestoreUDisk14(&ctx)
	testSet289DescribeUDisk15(&ctx)
	testSet289DeleteSnapshot16(&ctx)
	testSet289DescribeSnapshot17(&ctx)
	testSet289DeleteSnapshot18(&ctx)
	testSet289DescribeSnapshot19(&ctx)
	testSet289DeleteUDiskSnapshot20(&ctx)
	testSet289DescribeSnapshot21(&ctx)
	testSet289DescribeUDiskSnapshot22(&ctx)
	testSet289DeleteUDisk23(&ctx)
	testSet289DescribeUDisk24(&ctx)
	testSet289DeleteUDisk25(&ctx)
	testSet289DescribeUDisk26(&ctx)
	testSet289DeleteUDisk27(&ctx)
}

func testSet289DescribeUDiskPrice00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewDescribeUDiskPriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UDataArkMode", ctx.GetVar("UDataArkMode")))

	ctx.NoError(utest.SetReqValue(req, "Size", ctx.GetVar("Size")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Quantity", 1))

	ctx.NoError(utest.SetReqValue(req, "DiskType", ctx.GetVar("UDiskType")))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DescribeUDiskPrice(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet289CheckUDiskAllowance01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := pudiskClient.NewCheckUDiskAllowanceRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "Size", ctx.GetVar("Size")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Count", 1))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return pudiskClient.CheckUDiskAllowance(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "CheckUDiskAllowanceResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet289CreateUDisk02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(2) * time.Second)

	req := udiskClient.NewCreateUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UDataArkMode", ctx.GetVar("UDataArkMode")))

	ctx.NoError(utest.SetReqValue(req, "Size", ctx.GetVar("Size")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Quantity", 0))

	ctx.NoError(utest.SetReqValue(req, "Name", ctx.GetVar("UDiskName")))

	ctx.NoError(utest.SetReqValue(req, "DiskType", ctx.GetVar("UDiskType")))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.CreateUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "CreateUDiskResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

	ctx.Vars["udisk_noArk_Id"] = ctx.Must(utest.GetValue(resp, "UDiskId.0"))
}

func testSet289DescribeUDisk03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := udiskClient.NewDescribeUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_noArk_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DescribeUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet289CreateUDiskSnapshot04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewCreateUDiskSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_noArk_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Name", "snapshot_01"))

	ctx.NoError(utest.SetReqValue(req, "Comment", "comment_01"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.CreateUDiskSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "CreateUDiskSnapshotResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

	ctx.Vars["snapshot1_Id"] = ctx.Must(utest.GetValue(resp, "SnapshotId.0"))
}

func testSet289DescribeUDisk05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(120) * time.Second)

	req := udiskClient.NewDescribeUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_noArk_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DescribeUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("DataSet.0.SnapshotLimit", 3, "str_eq"),
			ctx.NewValidator("DataSet.0.SnapshotCount", 1, "str_eq"),
			ctx.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet289DescribeSnapshot06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(120) * time.Second)

	req := puhostClient.NewDescribeSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "SnapshotIds", ctx.GetVar("snapshot1_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return puhostClient.DescribeSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("UHostSnapshotSet.0.DiskId", ctx.GetVar("udisk_noArk_Id"), "str_eq"),
			ctx.NewValidator("UHostSnapshotSet.0.State", "Normal", "str_eq"),
		},
		MaxRetries:    50,
		RetryInterval: 3 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet289CreateUDiskSnapshot07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewCreateUDiskSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_noArk_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Name", "snapshot_02"))

	ctx.NoError(utest.SetReqValue(req, "Comment", "comment_01"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.CreateUDiskSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

	ctx.Vars["snapshot2_Id"] = ctx.Must(utest.GetValue(resp, "SnapshotId.0"))
}

func testSet289DescribeSnapshot08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(120) * time.Second)

	req := puhostClient.NewDescribeSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "SnapshotIds", ctx.GetVar("snapshot2_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return puhostClient.DescribeSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("UHostSnapshotSet.0.DiskId", ctx.GetVar("udisk_noArk_Id"), "str_eq"),
			ctx.NewValidator("UHostSnapshotSet.0.State", "Normal", "str_eq"),
		},
		MaxRetries:    30,
		RetryInterval: 5 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet289CreateUDiskSnapshot09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewCreateUDiskSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_noArk_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Name", "snapshot_03"))

	ctx.NoError(utest.SetReqValue(req, "Comment", "comment_03"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.CreateUDiskSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "CreateUDiskSnapshotResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

	ctx.Vars["snapshot3_Id"] = ctx.Must(utest.GetValue(resp, "SnapshotId.0"))
}

func testSet289DescribeUDiskSnapshot10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(120) * time.Second)

	req := udiskClient.NewDescribeUDiskSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "SnapshotId", ctx.GetVar("snapshot3_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DescribeUDiskSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "DescribeUDiskSnapshotResponse", "str_eq"),
			ctx.NewValidator("DataSet.0.Status", "Normal", "str_eq"),
			ctx.NewValidator("DataSet.0.UDiskId", ctx.GetVar("udisk_noArk_Id"), "str_eq"),
		},
		MaxRetries:    30,
		RetryInterval: 5 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet289CreateUDiskSnapshot11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := udiskClient.NewCreateUDiskSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_noArk_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Name", "snapshot_04"))

	ctx.NoError(utest.SetReqValue(req, "Comment", "comment_04"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.CreateUDiskSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 16999, "gt"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet289CloneUDiskSnapshot12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewCloneUDiskSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UDataArkMode", "Yes"))

	ctx.NoError(utest.SetReqValue(req, "SourceId", ctx.GetVar("snapshot1_Id")))

	ctx.NoError(utest.SetReqValue(req, "Size", ctx.GetVar("Size")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Quantity", 0))

	ctx.NoError(utest.SetReqValue(req, "Name", "clone_snap1_Ark"))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.CloneUDiskSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "CloneUDiskSnapshotResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 2 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

	ctx.Vars["clone_snap1_Ark"] = ctx.Must(utest.GetValue(resp, "UDiskId.0"))
}

func testSet289CloneUDiskSnapshot13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := udiskClient.NewCloneUDiskSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UDataArkMode", "No"))

	ctx.NoError(utest.SetReqValue(req, "SourceId", ctx.GetVar("snapshot1_Id")))

	ctx.NoError(utest.SetReqValue(req, "Size", ctx.GetVar("Size")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Quantity", 0))

	ctx.NoError(utest.SetReqValue(req, "Name", "clone_snap1_noArk"))

	ctx.NoError(utest.SetReqValue(req, "ChargeType", "Month"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.CloneUDiskSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "CloneUDiskSnapshotResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

	ctx.Vars["clone_snap1_noArk"] = ctx.Must(utest.GetValue(resp, "UDiskId.0"))
}

func testSet289RestoreUDisk14(ctx *utest.TestContext) {
	time.Sleep(time.Duration(30) * time.Second)

	req := udiskClient.NewRestoreUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_noArk_Id")))

	ctx.NoError(utest.SetReqValue(req, "SnapshotId", ctx.GetVar("snapshot1_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.RestoreUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "RestoreUDiskResponse", "str_eq"),
		},
		MaxRetries:    10,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet289DescribeUDisk15(ctx *utest.TestContext) {
	time.Sleep(time.Duration(50) * time.Second)

	req := udiskClient.NewDescribeUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_noArk_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DescribeUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		},
		MaxRetries:    30,
		RetryInterval: 3 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet289DeleteSnapshot16(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := puhostClient.NewDeleteSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "SnapshotId", ctx.GetVar("snapshot1_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return puhostClient.DeleteSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("SnapshotId", ctx.GetVar("snapshot1_Id"), "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet289DescribeSnapshot17(ctx *utest.TestContext) {
	time.Sleep(time.Duration(1) * time.Second)

	req := puhostClient.NewDescribeSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "SnapshotIds", ctx.GetVar("snapshot1_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return puhostClient.DescribeSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("TotalCount", 0, "str_eq"),
			ctx.NewValidator("PerDiskQuota", 3, "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet289DeleteSnapshot18(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := puhostClient.NewDeleteSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "SnapshotId", ctx.GetVar("snapshot2_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return puhostClient.DeleteSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("SnapshotId", ctx.GetVar("snapshot2_Id"), "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet289DescribeSnapshot19(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := puhostClient.NewDescribeSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "SnapshotIds", ctx.GetVar("snapshot2_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return puhostClient.DescribeSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("TotalCount", 0, "str_eq"),
			ctx.NewValidator("PerDiskQuota", 3, "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet289DeleteUDiskSnapshot20(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewDeleteUDiskSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "SnapshotId", ctx.GetVar("snapshot3_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DeleteUDiskSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "DeleteUDiskSnapshotResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet289DescribeSnapshot21(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := puhostClient.NewDescribeSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "SnapshotIds", ctx.GetVar("snapshot3_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return puhostClient.DescribeSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("TotalCount", 0, "str_eq"),
			ctx.NewValidator("PerDiskQuota", 3, "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet289DescribeUDiskSnapshot22(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewDescribeUDiskSnapshotRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_noArk_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DescribeUDiskSnapshot(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "DescribeUDiskSnapshotResponse", "str_eq"),
			ctx.NewValidator("TotalCount", 0, "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 3 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet289DeleteUDisk23(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewDeleteUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("udisk_noArk_Id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DeleteUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet289DescribeUDisk24(ctx *utest.TestContext) {
	time.Sleep(time.Duration(90) * time.Second)

	req := udiskClient.NewDescribeUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("clone_snap1_Ark")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DescribeUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		},
		MaxRetries:    30,
		RetryInterval: 3 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet289DeleteUDisk25(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewDeleteUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("clone_snap1_Ark")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DeleteUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet289DescribeUDisk26(ctx *utest.TestContext) {
	time.Sleep(time.Duration(20) * time.Second)

	req := udiskClient.NewDescribeUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("clone_snap1_noArk")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DescribeUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("DataSet.0.Status", "Available", "str_eq"),
		},
		MaxRetries:    30,
		RetryInterval: 3 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Fatal(err)
		ctx.T.FailNow()

	}

}

func testSet289DeleteUDisk27(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := udiskClient.NewDeleteUDiskRequest()

	ctx.NoError(utest.SetReqValue(req, "Zone", ctx.GetVar("Zone")))

	ctx.NoError(utest.SetReqValue(req, "UDiskId", ctx.GetVar("clone_snap1_noArk")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return udiskClient.DeleteUDisk(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
			ctx.NewValidator("Action", "DeleteUDiskResponse", "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}
