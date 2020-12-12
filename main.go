/**
 *
 * Copyright 2020 Victor Shinya
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
package main

import (
	"github.com/IBM/go-sdk-core/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

var (
	vpc *vpcv1.VpcV1
	err error
)

// Main function is called when OpenWhisk call the action
func Main(params map[string]interface{}) map[string]interface{} {
	if vpc == nil {
		apiKey := params["apikey"].(string)
		vpc, err = vpcv1.NewVpcV1(&vpcv1.VpcV1Options{Authenticator: &core.IamAuthenticator{ApiKey: apiKey}})
	}
	if err == nil {
		instancesID := params["instance_id"].([]string)
		actionType := params["type"].(string)
		for _, instance := range instancesID {
			_, _, err = vpc.CreateInstanceAction(&vpcv1.CreateInstanceActionOptions{
				InstanceID: &instance,
				Type:       &actionType,
			})
		}
	}
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	result := make(map[string]interface{})
	result["status"] = "Done"
	result["error"] = errMsg
	return result
}
