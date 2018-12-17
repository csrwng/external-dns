package nas

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

// ModifyAccessGroup invokes the nas.ModifyAccessGroup API synchronously
// api document: https://help.aliyun.com/api/nas/modifyaccessgroup.html
func (client *Client) ModifyAccessGroup(request *ModifyAccessGroupRequest) (response *ModifyAccessGroupResponse, err error) {
	response = CreateModifyAccessGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyAccessGroupWithChan invokes the nas.ModifyAccessGroup API asynchronously
// api document: https://help.aliyun.com/api/nas/modifyaccessgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAccessGroupWithChan(request *ModifyAccessGroupRequest) (<-chan *ModifyAccessGroupResponse, <-chan error) {
	responseChan := make(chan *ModifyAccessGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyAccessGroup(request)
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

// ModifyAccessGroupWithCallback invokes the nas.ModifyAccessGroup API asynchronously
// api document: https://help.aliyun.com/api/nas/modifyaccessgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAccessGroupWithCallback(request *ModifyAccessGroupRequest, callback func(response *ModifyAccessGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyAccessGroupResponse
		var err error
		defer close(result)
		response, err = client.ModifyAccessGroup(request)
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

// ModifyAccessGroupRequest is the request struct for api ModifyAccessGroup
type ModifyAccessGroupRequest struct {
	*requests.RpcRequest
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	Description     string `position:"Query" name:"Description"`
}

// ModifyAccessGroupResponse is the response struct for api ModifyAccessGroup
type ModifyAccessGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyAccessGroupRequest creates a request to invoke ModifyAccessGroup API
func CreateModifyAccessGroupRequest() (request *ModifyAccessGroupRequest) {
	request = &ModifyAccessGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "ModifyAccessGroup", "nas", "openAPI")
	return
}

// CreateModifyAccessGroupResponse creates a response to parse from ModifyAccessGroup response
func CreateModifyAccessGroupResponse() (response *ModifyAccessGroupResponse) {
	response = &ModifyAccessGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
