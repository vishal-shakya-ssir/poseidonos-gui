package ibofos

import (
	"a-module/src/routers/m9k/model"
)

func ListArrayDevice(xrId string, param interface{}) (model.Request, model.Response, error) {
	return arraySender(xrId, param, "LISTARRAYDEVICE")
}

func LoadArray(xrId string, param interface{}) (model.Request, model.Response, error) {
	return arraySender(xrId, param, "LOADARRAY")
}

func CreateArray(xrId string, param interface{}) (model.Request, model.Response, error) {
	return arraySender(xrId, param, "CREATEARRAY")
}

func DeleteArray(xrId string, param interface{}) (model.Request, model.Response, error) {
	return arraySender(xrId, param, "DELETEARRAY")
}

func ArrayInfo(xrId string, param interface{}) (model.Request, model.Response, error) {
	return arraySender(xrId, param, "ARRAYINFO")
}

func arraySender(xrId string, param interface{}, command string) (model.Request, model.Response, error) {
	return Requester{xrId, param, model.ArrayParam{}}.Send(command)
}
