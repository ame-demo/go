package main

import (
	"encoding/json"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ame "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ame/v20190916"
)

func main()  {
	credential := common.NewCredential(
		"",
		"",
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 5
	//cpf.HttpProfile.Endpoint = "cvm.ap-guangzhou.tencentcloudapi.com"
	cpf.SignMethod = "HmacSHA1"

	client, _ := ame.NewClient(credential, "ap-guangzhou", cpf)
	// 获取分类
	request := ame.NewDescribeStationsRequest()
	request.Limit = common.Uint64Ptr(10)
	request.Offset = common.Uint64Ptr(0)

	// get response structure
	response, err := client.DescribeStations(request)
	// API errors
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// unexpected errors
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(response.Response)
	fmt.Printf("%s\n", b)

	// 获取分类音乐列表
	request1 := ame.NewDescribeItemsRequest()
	request1.Limit = common.Uint64Ptr(10)
	request1.Offset = common.Uint64Ptr(0)
	request1.CategoryId = response.Response.Stations[0].CategoryID
	response1, err := client.DescribeItems(request1)
	b, _ = json.Marshal(response1.Response)
	fmt.Printf("%s\n", b)

	// 获取歌曲信息
	request2 := ame.NewDescribeMusicRequest()
	request2.ItemId = response1.Response.Items[0].ItemID
	// 请使用C端用户唯一标识
	request2.IdentityId = common.StringPtr("1234")
	request2.SubItemType = common.StringPtr("MP3-128K-FTD")
	response2, err := client.DescribeMusic(request2)
	b, _ = json.Marshal(response2.Response)
	fmt.Printf("%s\n", b)

	// 获取歌词信息
	request3 := ame.NewDescribeLyricRequest()
	request3.ItemId = response1.Response.Items[0].ItemID
	response3, err := client.DescribeLyric(request3)
	b, _ = json.Marshal(response3.Response)
	fmt.Printf("%s\n", b)
}
