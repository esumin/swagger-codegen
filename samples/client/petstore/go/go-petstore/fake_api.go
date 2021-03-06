/* 
 * Swagger Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * OpenAPI spec version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package petstore

import (
	"net/url"
	"time"
	"encoding/json"
)

type FakeApi struct {
	Configuration *Configuration
}

func NewFakeApi() *FakeApi {
	configuration := NewConfiguration()
	return &FakeApi{
		Configuration: configuration,
	}
}

func NewFakeApiWithBasePath(basePath string) *FakeApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &FakeApi{
		Configuration: configuration,
	}
}

/**
 * To test \&quot;client\&quot; model
 *
 * @param body client model
 * @return *Client
 */
func (a FakeApi) TestClientModel(body Client) (*Client, *APIResponse, error) {

	var localVarHttpMethod = "Patch"
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/fake"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body

	var successPayload = new(Client)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(localVarHttpResponse.RawResponse), err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(localVarHttpResponse.RawResponse), err
}

/**
 * Fake endpoint for testing various parameters 假端點 偽のエンドポイント 가짜 엔드 포인트 
 * Fake endpoint for testing various parameters 假端點 偽のエンドポイント 가짜 엔드 포인트 
 *
 * @param number None
 * @param double None
 * @param patternWithoutDelimiter None
 * @param byte_ None
 * @param integer None
 * @param int32_ None
 * @param int64_ None
 * @param float None
 * @param string_ None
 * @param binary None
 * @param date None
 * @param dateTime None
 * @param password None
 * @return void
 */
func (a FakeApi) TestEndpointParameters(number float32, double float64, patternWithoutDelimiter string, byte_ string, integer int32, int32_ int32, int64_ int64, float float32, string_ string, binary string, date time.Time, dateTime time.Time, password string) (*APIResponse, error) {

	var localVarHttpMethod = "Post"
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/fake"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(http_basic_test)' required
	// http basic authentication required
	if a.Configuration.UserName != "" || a.Configuration.Password != ""{
		localVarHeaderParams["Authorization"] =  "Basic " + a.Configuration.GetBasicAuthEncodedString()
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/xml; charset=utf-8", "application/json; charset=utf-8",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/xml; charset=utf-8",
"application/json; charset=utf-8",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarFormParams["integer"] = a.Configuration.APIClient.ParameterToString(integer, "")
	localVarFormParams["int32_"] = a.Configuration.APIClient.ParameterToString(int32_, "")
	localVarFormParams["int64_"] = a.Configuration.APIClient.ParameterToString(int64_, "")
	localVarFormParams["number"] = a.Configuration.APIClient.ParameterToString(number, "")
	localVarFormParams["float"] = a.Configuration.APIClient.ParameterToString(float, "")
	localVarFormParams["double"] = a.Configuration.APIClient.ParameterToString(double, "")
	localVarFormParams["string_"] = a.Configuration.APIClient.ParameterToString(string_, "")
	localVarFormParams["patternWithoutDelimiter"] = a.Configuration.APIClient.ParameterToString(patternWithoutDelimiter, "")
	localVarFormParams["byte_"] = a.Configuration.APIClient.ParameterToString(byte_, "")
	localVarFormParams["binary"] = a.Configuration.APIClient.ParameterToString(binary, "")
	localVarFormParams["date"] = a.Configuration.APIClient.ParameterToString(date, "")
	localVarFormParams["dateTime"] = a.Configuration.APIClient.ParameterToString(dateTime, "")
	localVarFormParams["password"] = a.Configuration.APIClient.ParameterToString(password, "")


	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return NewAPIResponse(localVarHttpResponse.RawResponse), err
	}

	return NewAPIResponse(localVarHttpResponse.RawResponse), err
}

/**
 * To test enum parameters
 *
 * @param enumFormStringArray Form parameter enum test (string array)
 * @param enumFormString Form parameter enum test (string)
 * @param enumHeaderStringArray Header parameter enum test (string array)
 * @param enumHeaderString Header parameter enum test (string)
 * @param enumQueryStringArray Query parameter enum test (string array)
 * @param enumQueryString Query parameter enum test (string)
 * @param enumQueryInteger Query parameter enum test (double)
 * @param enumQueryDouble Query parameter enum test (double)
 * @return void
 */
func (a FakeApi) TestEnumParameters(enumFormStringArray []string, enumFormString string, enumHeaderStringArray []string, enumHeaderString string, enumQueryStringArray []string, enumQueryString string, enumQueryInteger float32, enumQueryDouble float64) (*APIResponse, error) {

	var localVarHttpMethod = "Get"
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/fake"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	var collectionFormat = "csv"
	if collectionFormat == "multi" {
		for _, value := range enumQueryStringArray {
			localVarQueryParams.Add("enum_query_string_array", value)
		}
	} else {
		localVarQueryParams.Add("enum_query_string_array", a.Configuration.APIClient.ParameterToString(enumQueryStringArray, collectionFormat))
	}
			localVarQueryParams.Add("enum_query_string", a.Configuration.APIClient.ParameterToString(enumQueryString, ""))
			localVarQueryParams.Add("enum_query_integer", a.Configuration.APIClient.ParameterToString(enumQueryInteger, ""))
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	// header params "enum_header_string_array"
	localVarHeaderParams["enum_header_string_array"] = a.Configuration.APIClient.ParameterToString(enumHeaderStringArray, "")
	// header params "enum_header_string"
	localVarHeaderParams["enum_header_string"] = a.Configuration.APIClient.ParameterToString(enumHeaderString, "")

	localVarFormParams["enumFormStringArray"] = a.Configuration.APIClient.ParameterToString(enumFormStringArray, "")
	localVarFormParams["enumFormString"] = a.Configuration.APIClient.ParameterToString(enumFormString, "")
	localVarFormParams["enumQueryDouble"] = a.Configuration.APIClient.ParameterToString(enumQueryDouble, "")


	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return NewAPIResponse(localVarHttpResponse.RawResponse), err
	}

	return NewAPIResponse(localVarHttpResponse.RawResponse), err
}

