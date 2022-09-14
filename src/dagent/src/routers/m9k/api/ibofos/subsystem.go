package ibofos

import (
    pos "kouros/pos"
    "kouros/model"
    pb "kouros/api"
    "encoding/json"
    "kouros/log"
    "google.golang.org/protobuf/encoding/protojson"
)

func CallAddListener(xrId string, param interface{}, posMngr pos.POSManager) (model.Response, error) {
    var paramStruct pb.AddListenerRequest_Param
    pByte, err := json.Marshal(param)
    if err != nil {
        log.Errorf(marshalErrMsg, GetFuncName(1), err)
    }
    if err = json.Unmarshal(pByte, &paramStruct); err != nil {
        log.Errorf(unmarshalErrMsg, GetFuncName(1), err)
    }
    result, err1 := posMngr.AddListener(&paramStruct)
    if err1 != nil {
        log.Errorf(commandFailureMsg, GetFuncName(1), err1)
    }
    resByte, err2 := protojson.Marshal(result)
    return HandleResponse(resByte, err2)

}

func CallCreateSubsystem(xrId string, param interface{}, posMngr pos.POSManager) (model.Response, error) {
    var paramStruct pb.CreateSubsystemRequest_Param
    pByte, err := json.Marshal(param)
    if err != nil {
        log.Errorf(marshalErrMsg, GetFuncName(1), err)
    }
    if err = json.Unmarshal(pByte, &paramStruct); err != nil {
        log.Errorf(unmarshalErrMsg, GetFuncName(1), err)
    }
    result, err1 := posMngr.CreateSubsystem(&paramStruct)
    if err1 != nil {
        log.Errorf(commandFailureMsg, GetFuncName(1), err1)
    }
    resByte, err2 := protojson.Marshal(result)
    return HandleResponse(resByte, err2)

}

func CallCreateTransport(xrId string, param interface{}, posMngr pos.POSManager) (model.Response, error) {
    var paramStruct pb.CreateTransportRequest_Param
    pByte, err := json.Marshal(param)
    if err != nil {
        log.Errorf(marshalErrMsg, GetFuncName(1), err)
    }
    if err = json.Unmarshal(pByte, &paramStruct); err != nil {
        log.Errorf(unmarshalErrMsg, GetFuncName(1), err)
    }
    result, err1 := posMngr.CreateTransport(&paramStruct)
    if err1 != nil {
        log.Errorf(commandFailureMsg, GetFuncName(1), err1)
    }
    resByte, err2 := protojson.Marshal(result)
    return HandleResponse(resByte, err2)

}

func CallDeleteSubsystem(xrId string, param interface{}, posMngr pos.POSManager) (model.Response, error) {
    var paramStruct pb.DeleteSubsystemRequest_Param
    pByte, err := json.Marshal(param)
    if err != nil {
        log.Errorf(marshalErrMsg, GetFuncName(1), err)
    }
    if err = json.Unmarshal(pByte, &paramStruct); err != nil {
        log.Errorf(unmarshalErrMsg, GetFuncName(1), err)
    }
    result, err1 := posMngr.DeleteSubsystem(&paramStruct)
    if err1 != nil {
        log.Errorf(commandFailureMsg, GetFuncName(1), err1)
    }
    resByte, err2 := protojson.Marshal(result)
    return HandleResponse(resByte, err2)

}

func CallSubsystemInfo(xrId string, param interface{}, posMngr pos.POSManager) (model.Response, error) {
    var paramStruct pb.SubsystemInfoRequest_Param
    pByte, err := json.Marshal(param)
    if err != nil {
        log.Errorf(marshalErrMsg, GetFuncName(1), err)
    }
    if err = json.Unmarshal(pByte, &paramStruct); err != nil {
        log.Errorf(unmarshalErrMsg, GetFuncName(1), err)
    }
    result, err1 := posMngr.SubsystemInfo(&paramStruct)
    if err1 != nil {
        log.Errorf(commandFailureMsg, GetFuncName(1), err1)
    }
    resByte, err2 := protojson.Marshal(result)
    return HandleResponse(resByte, err2)
}

func CallListSubsystem(xrId string, param interface{}, posMngr pos.POSManager) (model.Response, error) {
    result, err1 := posMngr.ListSubsystem()
    if err1 != nil {
        log.Errorf(commandFailureMsg, GetFuncName(1), err1)
    }
    resByte, err2 := protojson.Marshal(result)
    return HandleResponse(resByte, err2)
}

