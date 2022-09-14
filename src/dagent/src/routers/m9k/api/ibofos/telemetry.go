package ibofos

import (
    pos "kouros/pos"
    "kouros/model"
    "google.golang.org/protobuf/encoding/protojson"
    "kouros/log"
)

func CallStartTelemetry(xrId string, param interface{}, posMngr pos.POSManager) (model.Response, error) {
    result, err1 := posMngr.StartTelemetry()
    if err1 != nil {
        log.Errorf(commandFailureMsg, GetFuncName(1), err1)
    }
    resByte, err2 := protojson.Marshal(result)
    return HandleResponse(resByte, err2)
}

func CallStopTelemetry(xrId string, param interface{}, posMngr pos.POSManager) (model.Response, error) {
    result, err1 := posMngr.StopTelemetry()
    if err1 != nil {
        log.Errorf(commandFailureMsg, GetFuncName(1), err1)
    }
    resByte, err2 := protojson.Marshal(result)
    return HandleResponse(resByte, err2)
}

