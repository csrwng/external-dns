package ehpc

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

// ListJobTemplates invokes the ehpc.ListJobTemplates API synchronously
// api document: https://help.aliyun.com/api/ehpc/listjobtemplates.html
func (client *Client) ListJobTemplates(request *ListJobTemplatesRequest) (response *ListJobTemplatesResponse, err error) {
	response = CreateListJobTemplatesResponse()
	err = client.DoAction(request, response)
	return
}

// ListJobTemplatesWithChan invokes the ehpc.ListJobTemplates API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listjobtemplates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListJobTemplatesWithChan(request *ListJobTemplatesRequest) (<-chan *ListJobTemplatesResponse, <-chan error) {
	responseChan := make(chan *ListJobTemplatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListJobTemplates(request)
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

// ListJobTemplatesWithCallback invokes the ehpc.ListJobTemplates API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listjobtemplates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListJobTemplatesWithCallback(request *ListJobTemplatesRequest, callback func(response *ListJobTemplatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListJobTemplatesResponse
		var err error
		defer close(result)
		response, err = client.ListJobTemplates(request)
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

// ListJobTemplatesRequest is the request struct for api ListJobTemplates
type ListJobTemplatesRequest struct {
	*requests.RpcRequest
	Name       string           `position:"Query" name:"Name"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListJobTemplatesResponse is the response struct for api ListJobTemplates
type ListJobTemplatesResponse struct {
	*responses.BaseResponse
	RequestId  string    `json:"RequestId" xml:"RequestId"`
	TotalCount int       `json:"TotalCount" xml:"TotalCount"`
	PageNumber int       `json:"PageNumber" xml:"PageNumber"`
	PageSize   int       `json:"PageSize" xml:"PageSize"`
	Templates  Templates `json:"Templates" xml:"Templates"`
}

// CreateListJobTemplatesRequest creates a request to invoke ListJobTemplates API
func CreateListJobTemplatesRequest() (request *ListJobTemplatesRequest) {
	request = &ListJobTemplatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "ListJobTemplates", "ehs", "openAPI")
	return
}

// CreateListJobTemplatesResponse creates a response to parse from ListJobTemplates response
func CreateListJobTemplatesResponse() (response *ListJobTemplatesResponse) {
	response = &ListJobTemplatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
