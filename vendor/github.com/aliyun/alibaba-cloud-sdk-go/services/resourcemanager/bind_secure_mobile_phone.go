package resourcemanager

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

// BindSecureMobilePhone invokes the resourcemanager.BindSecureMobilePhone API synchronously
func (client *Client) BindSecureMobilePhone(request *BindSecureMobilePhoneRequest) (response *BindSecureMobilePhoneResponse, err error) {
	response = CreateBindSecureMobilePhoneResponse()
	err = client.DoAction(request, response)
	return
}

// BindSecureMobilePhoneWithChan invokes the resourcemanager.BindSecureMobilePhone API asynchronously
func (client *Client) BindSecureMobilePhoneWithChan(request *BindSecureMobilePhoneRequest) (<-chan *BindSecureMobilePhoneResponse, <-chan error) {
	responseChan := make(chan *BindSecureMobilePhoneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BindSecureMobilePhone(request)
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

// BindSecureMobilePhoneWithCallback invokes the resourcemanager.BindSecureMobilePhone API asynchronously
func (client *Client) BindSecureMobilePhoneWithCallback(request *BindSecureMobilePhoneRequest, callback func(response *BindSecureMobilePhoneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BindSecureMobilePhoneResponse
		var err error
		defer close(result)
		response, err = client.BindSecureMobilePhone(request)
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

// BindSecureMobilePhoneRequest is the request struct for api BindSecureMobilePhone
type BindSecureMobilePhoneRequest struct {
	*requests.RpcRequest
	SecureMobilePhone string `position:"Query" name:"SecureMobilePhone"`
	AccountId         string `position:"Query" name:"AccountId"`
	VerificationCode  string `position:"Query" name:"VerificationCode"`
}

// BindSecureMobilePhoneResponse is the response struct for api BindSecureMobilePhone
type BindSecureMobilePhoneResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBindSecureMobilePhoneRequest creates a request to invoke BindSecureMobilePhone API
func CreateBindSecureMobilePhoneRequest() (request *BindSecureMobilePhoneRequest) {
	request = &BindSecureMobilePhoneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "BindSecureMobilePhone", "", "")
	request.Method = requests.POST
	return
}

// CreateBindSecureMobilePhoneResponse creates a response to parse from BindSecureMobilePhone response
func CreateBindSecureMobilePhoneResponse() (response *BindSecureMobilePhoneResponse) {
	response = &BindSecureMobilePhoneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}