package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeAccessPoints invokes the vpc.DescribeAccessPoints API synchronously
// api document: https://help.aliyun.com/api/vpc/describeaccesspoints.html
func (client *Client) DescribeAccessPoints(request *DescribeAccessPointsRequest) (response *DescribeAccessPointsResponse, err error) {
	response = CreateDescribeAccessPointsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAccessPointsWithChan invokes the vpc.DescribeAccessPoints API asynchronously
// api document: https://help.aliyun.com/api/vpc/describeaccesspoints.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAccessPointsWithChan(request *DescribeAccessPointsRequest) (<-chan *DescribeAccessPointsResponse, <-chan error) {
	responseChan := make(chan *DescribeAccessPointsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAccessPoints(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeAccessPointsWithCallback invokes the vpc.DescribeAccessPoints API asynchronously
// api document: https://help.aliyun.com/api/vpc/describeaccesspoints.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAccessPointsWithCallback(request *DescribeAccessPointsRequest, callback func(response *DescribeAccessPointsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAccessPointsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAccessPoints(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeAccessPointsRequest is the request struct for api DescribeAccessPoints
type DescribeAccessPointsRequest struct {
	*requests.RpcRequest
	Filter               *[]DescribeAccessPointsFilter `position:"Query" name:"Filter"  type:"Repeated"`
	ResourceOwnerId      requests.Integer              `position:"Query" name:"ResourceOwnerId"`
	HostOperator         string                        `position:"Query" name:"HostOperator"`
	ResourceOwnerAccount string                        `position:"Query" name:"ResourceOwnerAccount"`
	Name                 string                        `position:"Query" name:"Name"`
	PageSize             requests.Integer              `position:"Query" name:"PageSize"`
	OwnerId              requests.Integer              `position:"Query" name:"OwnerId"`
	Type                 string                        `position:"Query" name:"Type"`
	PageNumber           requests.Integer              `position:"Query" name:"PageNumber"`
}

// DescribeAccessPointsFilter is a repeated param struct in DescribeAccessPointsRequest
type DescribeAccessPointsFilter struct {
	Value *[]string `name:"Value" type:"Repeated"`
	Key   string    `name:"Key"`
}

// DescribeAccessPointsResponse is the response struct for api DescribeAccessPoints
type DescribeAccessPointsResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	PageNumber     int            `json:"PageNumber" xml:"PageNumber"`
	PageSize       int            `json:"PageSize" xml:"PageSize"`
	TotalCount     int            `json:"TotalCount" xml:"TotalCount"`
	AccessPointSet AccessPointSet `json:"AccessPointSet" xml:"AccessPointSet"`
}

// CreateDescribeAccessPointsRequest creates a request to invoke DescribeAccessPoints API
func CreateDescribeAccessPointsRequest() (request *DescribeAccessPointsRequest) {
	request = &DescribeAccessPointsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeAccessPoints", "vpc", "openAPI")
	return
}

// CreateDescribeAccessPointsResponse creates a response to parse from DescribeAccessPoints response
func CreateDescribeAccessPointsResponse() (response *DescribeAccessPointsResponse) {
	response = &DescribeAccessPointsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
