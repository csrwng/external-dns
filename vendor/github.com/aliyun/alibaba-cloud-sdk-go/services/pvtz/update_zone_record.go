package pvtz

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

// UpdateZoneRecord invokes the pvtz.UpdateZoneRecord API synchronously
// api document: https://help.aliyun.com/api/pvtz/updatezonerecord.html
func (client *Client) UpdateZoneRecord(request *UpdateZoneRecordRequest) (response *UpdateZoneRecordResponse, err error) {
	response = CreateUpdateZoneRecordResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateZoneRecordWithChan invokes the pvtz.UpdateZoneRecord API asynchronously
// api document: https://help.aliyun.com/api/pvtz/updatezonerecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateZoneRecordWithChan(request *UpdateZoneRecordRequest) (<-chan *UpdateZoneRecordResponse, <-chan error) {
	responseChan := make(chan *UpdateZoneRecordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateZoneRecord(request)
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

// UpdateZoneRecordWithCallback invokes the pvtz.UpdateZoneRecord API asynchronously
// api document: https://help.aliyun.com/api/pvtz/updatezonerecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateZoneRecordWithCallback(request *UpdateZoneRecordRequest, callback func(response *UpdateZoneRecordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateZoneRecordResponse
		var err error
		defer close(result)
		response, err = client.UpdateZoneRecord(request)
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

// UpdateZoneRecordRequest is the request struct for api UpdateZoneRecord
type UpdateZoneRecordRequest struct {
	*requests.RpcRequest
	Rr           string           `position:"Query" name:"Rr"`
	Lang         string           `position:"Query" name:"Lang"`
	RecordId     requests.Integer `position:"Query" name:"RecordId"`
	Type         string           `position:"Query" name:"Type"`
	Ttl          requests.Integer `position:"Query" name:"Ttl"`
	Priority     requests.Integer `position:"Query" name:"Priority"`
	Value        string           `position:"Query" name:"Value"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
}

// UpdateZoneRecordResponse is the response struct for api UpdateZoneRecord
type UpdateZoneRecordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	RecordId  int    `json:"RecordId" xml:"RecordId"`
}

// CreateUpdateZoneRecordRequest creates a request to invoke UpdateZoneRecord API
func CreateUpdateZoneRecordRequest() (request *UpdateZoneRecordRequest) {
	request = &UpdateZoneRecordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("pvtz", "2018-01-01", "UpdateZoneRecord", "pvtz", "openAPI")
	return
}

// CreateUpdateZoneRecordResponse creates a response to parse from UpdateZoneRecord response
func CreateUpdateZoneRecordResponse() (response *UpdateZoneRecordResponse) {
	response = &UpdateZoneRecordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
