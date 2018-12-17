package rds

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

// DescribeDBInstanceUser invokes the rds.DescribeDBInstanceUser API synchronously
// api document: https://help.aliyun.com/api/rds/describedbinstanceuser.html
func (client *Client) DescribeDBInstanceUser(request *DescribeDBInstanceUserRequest) (response *DescribeDBInstanceUserResponse, err error) {
	response = CreateDescribeDBInstanceUserResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstanceUserWithChan invokes the rds.DescribeDBInstanceUser API asynchronously
// api document: https://help.aliyun.com/api/rds/describedbinstanceuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceUserWithChan(request *DescribeDBInstanceUserRequest) (<-chan *DescribeDBInstanceUserResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceUser(request)
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

// DescribeDBInstanceUserWithCallback invokes the rds.DescribeDBInstanceUser API asynchronously
// api document: https://help.aliyun.com/api/rds/describedbinstanceuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceUserWithCallback(request *DescribeDBInstanceUserRequest, callback func(response *DescribeDBInstanceUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceUserResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceUser(request)
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

// DescribeDBInstanceUserRequest is the request struct for api DescribeDBInstanceUser
type DescribeDBInstanceUserRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ConnectionString     string           `position:"Query" name:"ConnectionString"`
}

// DescribeDBInstanceUserResponse is the response struct for api DescribeDBInstanceUser
type DescribeDBInstanceUserResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DBInstanceName string `json:"DBInstanceName" xml:"DBInstanceName"`
	InternalDBFlag string `json:"InternalDBFlag" xml:"InternalDBFlag"`
}

// CreateDescribeDBInstanceUserRequest creates a request to invoke DescribeDBInstanceUser API
func CreateDescribeDBInstanceUserRequest() (request *DescribeDBInstanceUserRequest) {
	request = &DescribeDBInstanceUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceUser", "rds", "openAPI")
	return
}

// CreateDescribeDBInstanceUserResponse creates a response to parse from DescribeDBInstanceUser response
func CreateDescribeDBInstanceUserResponse() (response *DescribeDBInstanceUserResponse) {
	response = &DescribeDBInstanceUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
