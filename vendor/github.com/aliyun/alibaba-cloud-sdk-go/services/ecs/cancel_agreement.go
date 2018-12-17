package ecs

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

// CancelAgreement invokes the ecs.CancelAgreement API synchronously
// api document: https://help.aliyun.com/api/ecs/cancelagreement.html
func (client *Client) CancelAgreement(request *CancelAgreementRequest) (response *CancelAgreementResponse, err error) {
	response = CreateCancelAgreementResponse()
	err = client.DoAction(request, response)
	return
}

// CancelAgreementWithChan invokes the ecs.CancelAgreement API asynchronously
// api document: https://help.aliyun.com/api/ecs/cancelagreement.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelAgreementWithChan(request *CancelAgreementRequest) (<-chan *CancelAgreementResponse, <-chan error) {
	responseChan := make(chan *CancelAgreementResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelAgreement(request)
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

// CancelAgreementWithCallback invokes the ecs.CancelAgreement API asynchronously
// api document: https://help.aliyun.com/api/ecs/cancelagreement.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelAgreementWithCallback(request *CancelAgreementRequest, callback func(response *CancelAgreementResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelAgreementResponse
		var err error
		defer close(result)
		response, err = client.CancelAgreement(request)
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

// CancelAgreementRequest is the request struct for api CancelAgreement
type CancelAgreementRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AgreementType        string           `position:"Query" name:"AgreementType"`
}

// CancelAgreementResponse is the response struct for api CancelAgreement
type CancelAgreementResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCancelAgreementRequest creates a request to invoke CancelAgreement API
func CreateCancelAgreementRequest() (request *CancelAgreementRequest) {
	request = &CancelAgreementRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CancelAgreement", "ecs", "openAPI")
	return
}

// CreateCancelAgreementResponse creates a response to parse from CancelAgreement response
func CreateCancelAgreementResponse() (response *CancelAgreementResponse) {
	response = &CancelAgreementResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
