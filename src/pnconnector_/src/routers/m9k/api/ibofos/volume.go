/*
 *   BSD LICENSE
 *   Copyright (c) 2021 Samsung Electronics Corporation
 *   All rights reserved.
 *
 *   Redistribution and use in source and binary forms, with or without
 *   modification, are permitted provided that the following conditions
 *   are met:
 *
 *     * Redistributions of source code must retain the above copyright
 *       notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above copyright
 *       notice, this list of conditions and the following disclaimer in
 *       the documentation and/or other materials provided with the
 *       distribution.
 *     * Neither the name of Samsung Electronics Corporation nor the names of its
 *       contributors may be used to endorse or promote products derived
 *       from this software without specific prior written permission.
 *
 *   THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 *   "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 *   LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 *   A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 *   OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 *   SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 *   LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 *   DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 *   THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 *   (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 *   OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */
package ibofos

import (
	"pnconnector/src/errors"
	"pnconnector/src/influxdb"
	"kouros/model"
)

func CreateVolume(xrId string, param interface{}) (model.Request, model.Response, error) {
	req, res, err := Requester{xrId, param, model.VolumeParam{}}.Send("CREATEVOLUME")

	if err != nil {
		return req, res, err
	}

	// Calling a go routine to mark the creation time of a volume in influxdb
	// The creation time is used in querying the influxdb for volume level data
	// This is currently a workaround for the issue that POS is crating duplicate volume Ids
	// The function is made as a go routine so that it does not cause much sde effects to the Create Volume API
	go influxdb.CreateVolume(ListVolume, param, res)

	return req, res, err
}

func UpdateVolume(xrId string, param interface{}) (model.Request, model.Response, error) {
	return Requester{xrId, param, model.VolumeParam{}}.Send("UPDATEVOLUMEQOS")
}

func MountVolume(xrId string, param interface{}) (model.Request, model.Response, error) {
	return Requester{xrId, param, model.VolumeParam{}}.Send("MOUNTVOLUME")
}

func CallSubSystemAuto(xrId string, res *model.Response, param map[string]interface{}, opName string, errorInfoList *[]map[string]interface{}) {
	subsystemParam := model.SubSystemParam{}
	subsystemParam.SUBNQN, _ = param["subnqn"].(string)
	_, response, _ := CreateSubSystemAuto(xrId, subsystemParam)
	//update response
	UpdateResponse(response, res, opName, errorInfoList)

}
func CallAddlistener(xrId string, res *model.Response, param map[string]interface{}, opName string, errorInfoList *[]map[string]interface{}) {
	subsystemParam := model.SubSystemParam{}
	subsystemParam.SUBNQN, _ = param["subnqn"].(string)
	subsystemParam.TRANSPORTTYPE, _ = param["transport_type"].(string)
	subsystemParam.TARGETADDRESS, _ = param["target_address"].(string)
	subsystemParam.TRANSPORTSERVICEID, _ = param["transport_service_id"].(string)
	_, response, _ := AddListener(xrId, subsystemParam)
	//update response
	UpdateResponse(response, res, opName, errorInfoList)

}
func CallMountVolume(xrId string, res *model.Response, param map[string]interface{}, opName string, errorInfoList *[]map[string]interface{}) {
	volumeParam := model.VolumeParam{}
	volumeParam.Name, _ = param["name"].(string)
	volumeParam.Array, _ = param["array"].(string)
	volumeParam.SubNQN, _ = param["subnqn"].(string)
	_, response, _ := MountVolume(xrId, volumeParam)
	//update response
	UpdateResponse(response, res, opName, errorInfoList)
}
func MountVolumeWithSubSystem(xrId string, param interface{}) (model.Request, model.Response, error) {
	errorInfo := make(map[string]interface{})
	errorInfoList := []map[string]interface{}{}
	iBoFRequest := model.Request{
		Command: "RUNIBOFOS",
		Rid:     xrId,
	}
	iBoFRequest.Param = param
	res := model.Response{}
	reqMap := param.(map[string]interface{})
	errorCode := 0
	CallSubSystemAuto(xrId, &res, reqMap, "SubSystemAuto", &errorInfoList)
	CallAddlistener(xrId, &res, reqMap, "AddListener", &errorInfoList)
	CallMountVolume(xrId, &res, reqMap, "MountVolume", &errorInfoList)
	for itr := 0; itr < len(errorInfoList); itr++ {
		if errorInfoList[itr]["code"].(int) != 0 {
			errorCode = 1
			break
		}
	}
	errorInfo["errorResponses"] = errorInfoList
	errorInfo["errorCode"] = errorCode
	res.Result.Status.ErrorInfo = errorInfo
	return iBoFRequest, res, nil
	//return Requester{xrId, param, model.VolumeParam{}}.Send("MOUNTVOLUME")
}

func UnmountVolume(xrId string, param interface{}) (model.Request, model.Response, error) {
	return Requester{xrId, param, model.VolumeParam{}}.Send("UNMOUNTVOLUME")
}

func DeleteVolume(xrId string, param interface{}) (model.Request, model.Response, error) {
	req, res, err := Requester{xrId, param, model.VolumeParam{}}.Send("DELETEVOLUME")

	if err != nil {
		return req, res, err
	}

	err = influxdb.DeleteVolume()

	if err != nil {
		err = errors.New("Request Success, but Influx Error : " + err.Error())
	}
	return req, res, err
}

func ListVolume(xrId string, param interface{}) (model.Request, model.Response, error) {
	return volumeSender(xrId, param, "LISTVOLUME")
}

func VolumeInfo(xrId string, param interface{}) (model.Request, model.Response, error) {
	return volumeSender(xrId, param, "VOLUMEINFO")
}

func UpdateVolumeQoS(xrId string, param interface{}) (model.Request, model.Response, error) {
	return volumeSender(xrId, param, "UPDATEVOLUMEQOS")
}

func RenameVolume(xrId string, param interface{}) (model.Request, model.Response, error) {
	return volumeSender(xrId, param, "RENAMEVOLUME")
}

func ResizeVolume(xrId string, param interface{}) (model.Request, model.Response, error) {
	return volumeSender(xrId, param, "RESIZEVOLUME")
}

func GetMaxVolumeCount(xrId string, param interface{}) (model.Request, model.Response, error) {
	return volumeSender(xrId, param, "GETMAXVOLUMECOUNT")
}

func GetHostNQN(xrId string, param interface{}) (model.Request, model.Response, error) {
	return volumeSender(xrId, param, "GETHOSTNQN")
}

func volumeSender(xrId string, param interface{}, command string) (model.Request, model.Response, error) {
	return Requester{xrId, param, model.VolumeParam{}}.Send(command)
}
