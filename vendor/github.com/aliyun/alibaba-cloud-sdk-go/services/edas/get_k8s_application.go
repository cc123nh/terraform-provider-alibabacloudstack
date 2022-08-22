package edas

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

// GetK8sApplication invokes the edas.GetK8sApplication API synchronously
func (client *Client) GetK8sApplication(request *GetK8sApplicationRequest) (response *GetK8sApplicationResponse, err error) {
	response = CreateGetK8sApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// GetK8sApplicationWithChan invokes the edas.GetK8sApplication API asynchronously
func (client *Client) GetK8sApplicationWithChan(request *GetK8sApplicationRequest) (<-chan *GetK8sApplicationResponse, <-chan error) {
	responseChan := make(chan *GetK8sApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetK8sApplication(request)
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

// GetK8sApplicationWithCallback invokes the edas.GetK8sApplication API asynchronously
func (client *Client) GetK8sApplicationWithCallback(request *GetK8sApplicationRequest, callback func(response *GetK8sApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetK8sApplicationResponse
		var err error
		defer close(result)
		response, err = client.GetK8sApplication(request)
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

// GetK8sApplicationRequest is the request struct for api GetK8sApplication
type GetK8sApplicationRequest struct {
	*requests.RoaRequest
	AppId string `position:"Query" name:"AppId"`
	From  string `position:"Query" name:"From"`
	RegionId string  `position:"Query" name:"RegionId"`
}

// GetK8sApplicationResponse is the response struct for api GetK8sApplication
type GetK8sApplicationResponse struct {
	*responses.BaseResponse
	Code       int        `json:"Code" xml:"Code"`
	Message    string     `json:"Message" xml:"Message"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Applcation Applcation `json:"Applcation" xml:"Applcation"`
}

// CreateGetK8sApplicationRequest creates a request to invoke GetK8sApplication API
func CreateGetK8sApplicationRequest() (request *GetK8sApplicationRequest) {
	request = &GetK8sApplicationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "GetK8sApplication", "/roa/pop/v5/changeorder/co_application", "Edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetK8sApplicationResponse creates a response to parse from GetK8sApplication response
func CreateGetK8sApplicationResponse() (response *GetK8sApplicationResponse) {
	response = &GetK8sApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}