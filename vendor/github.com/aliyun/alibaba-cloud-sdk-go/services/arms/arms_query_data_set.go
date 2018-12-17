package arms

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

// ARMSQueryDataSet invokes the arms.ARMSQueryDataSet API synchronously
// api document: https://help.aliyun.com/api/arms/armsquerydataset.html
func (client *Client) ARMSQueryDataSet(request *ARMSQueryDataSetRequest) (response *ARMSQueryDataSetResponse, err error) {
	response = CreateARMSQueryDataSetResponse()
	err = client.DoAction(request, response)
	return
}

// ARMSQueryDataSetWithChan invokes the arms.ARMSQueryDataSet API asynchronously
// api document: https://help.aliyun.com/api/arms/armsquerydataset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ARMSQueryDataSetWithChan(request *ARMSQueryDataSetRequest) (<-chan *ARMSQueryDataSetResponse, <-chan error) {
	responseChan := make(chan *ARMSQueryDataSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ARMSQueryDataSet(request)
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

// ARMSQueryDataSetWithCallback invokes the arms.ARMSQueryDataSet API asynchronously
// api document: https://help.aliyun.com/api/arms/armsquerydataset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ARMSQueryDataSetWithCallback(request *ARMSQueryDataSetRequest, callback func(response *ARMSQueryDataSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ARMSQueryDataSetResponse
		var err error
		defer close(result)
		response, err = client.ARMSQueryDataSet(request)
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

// ARMSQueryDataSetRequest is the request struct for api ARMSQueryDataSet
type ARMSQueryDataSetRequest struct {
	*requests.RpcRequest
	DatasetId     requests.Integer                `position:"Query" name:"DatasetId"`
	IntervalInSec requests.Integer                `position:"Query" name:"IntervalInSec"`
	DateStr       requests.Integer                `position:"Query" name:"DateStr"`
	MinTime       requests.Integer                `position:"Query" name:"MinTime"`
	MaxTime       requests.Integer                `position:"Query" name:"MaxTime"`
	Dimensions    *[]ARMSQueryDataSetDimensions   `position:"Query" name:"Dimensions"  type:"Repeated"`
	Measures      *[]string                       `position:"Query" name:"Measures"  type:"Repeated"`
	IsDrillDown   requests.Boolean                `position:"Query" name:"IsDrillDown"`
	OrderByKey    string                          `position:"Query" name:"OrderByKey"`
	Limit         requests.Integer                `position:"Query" name:"Limit"`
	ReduceTail    requests.Boolean                `position:"Query" name:"ReduceTail"`
	HungryMode    requests.Boolean                `position:"Query" name:"HungryMode"`
	RequiredDims  *[]ARMSQueryDataSetRequiredDims `position:"Query" name:"RequiredDims"  type:"Repeated"`
	OptionalDims  *[]ARMSQueryDataSetOptionalDims `position:"Query" name:"OptionalDims"  type:"Repeated"`
}

// ARMSQueryDataSetDimensions is a repeated param struct in ARMSQueryDataSetRequest
type ARMSQueryDataSetDimensions struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
	Type  string `name:"Type"`
}

// ARMSQueryDataSetRequiredDims is a repeated param struct in ARMSQueryDataSetRequest
type ARMSQueryDataSetRequiredDims struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
	Type  string `name:"Type"`
}

// ARMSQueryDataSetOptionalDims is a repeated param struct in ARMSQueryDataSetRequest
type ARMSQueryDataSetOptionalDims struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
	Type  string `name:"Type"`
}

// ARMSQueryDataSetResponse is the response struct for api ARMSQueryDataSet
type ARMSQueryDataSetResponse struct {
	*responses.BaseResponse
	Data string `json:"Data" xml:"Data"`
}

// CreateARMSQueryDataSetRequest creates a request to invoke ARMSQueryDataSet API
func CreateARMSQueryDataSetRequest() (request *ARMSQueryDataSetRequest) {
	request = &ARMSQueryDataSetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2016-11-25", "ARMSQueryDataSet", "", "")
	return
}

// CreateARMSQueryDataSetResponse creates a response to parse from ARMSQueryDataSet response
func CreateARMSQueryDataSetResponse() (response *ARMSQueryDataSetResponse) {
	response = &ARMSQueryDataSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
