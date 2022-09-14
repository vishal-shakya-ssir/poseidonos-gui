package ibofos

import (
    pos "kouros/pos"
    "kouros/model"
    "kouros/setting"
    "google.golang.org/protobuf/encoding/protojson"
    "kouros/log"
    pb "kouros/api"
    "encoding/json"
    "os"
    "path/filepath"
    "fmt"
    "time"
    "kouros/utils"
)

func CallGetSystemProperty(xrId string, param interface{}, posMngr pos.POSManager) (model.Response, error) {
    result, err1 := posMngr.GetSystemProperty()
    if err1 != nil {
        log.Errorf(commandFailureMsg, GetFuncName(1), err1)
    }
    resByte, err2 := protojson.Marshal(result)
    return HandleResponse(resByte, err2)
}

func CallSetSystemProperty(xrId string, param interface{}, posMngr pos.POSManager) (model.Response, error) {
    var paramStruct pb.SetSystemPropertyRequest_Param
    pByte, err := json.Marshal(param)
    if err != nil {
        log.Errorf(marshalErrMsg, GetFuncName(1), err)
    }
    if err = json.Unmarshal(pByte, &paramStruct); err != nil {
        log.Errorf(unmarshalErrMsg, GetFuncName(1), err)
    }
    result, err1 := posMngr.SetSystemProperty(&paramStruct)
    if err1 != nil {
        log.Errorf(commandFailureMsg, GetFuncName(1), err1)
    }
    resByte, err2 := protojson.Marshal(result)
    return HandleResponse(resByte, err2)
}

func UpdateResponse(response model.Response, res *model.Response, opName string, errorInfoList *[]map[string]interface{}) {
    info := make(map[string]interface{})
    info["id"] = opName
    info["code"] = response.Result.Status.Code
    info["description"] = response.Result.Status.Description
    *errorInfoList = append(*errorInfoList, info)
}

func transport(xrId string, config setting.HostConf, res *model.Response, opName string, errorInfoList *[]map[string]interface{},posMngr pos.POSManager) {
    transportParam := model.CreateTransportRequest_Param{}
    transportParam.TransportType = config.TransportType
    transportParam.BufCacheSize = int32(config.BufCacheSize)
    transportParam.NumSharedBuf = int32(config.NumSharedBuf)
    //create Transport
    response, _ := CallCreateTransport(xrId, transportParam, posMngr)
    UpdateResponse(response, res, opName, errorInfoList)

}

func subSystem(xrId string, config setting.HostConf, res *model.Response, name string, opName string, errorInfoList *[]map[string]interface{},posMngr pos.POSManager) {
    subsystemParam := model.CreateSubsystemRequest_Param{}
    subsystemParam.Nqn = name
    subsystemParam.SerialNumber = config.Serial
    subsystemParam.ModelNumber = config.Model
    subsystemParam.MaxNamespaces = uint32(config.MaxNameSpaces)
    subsystemParam.AllowAnyHost = config.AllowAnyHost
    subsystemParam.AnaReporting = config.AnaReporting

    response, _ := CallCreateSubsystem(xrId, subsystemParam, posMngr)
    //update response
    UpdateResponse(response, res, opName, errorInfoList)

}
func listener(xrId string, config setting.HostConf, res *model.Response, name string, opName string, errorInfoList *[]map[string]interface{},posMngr pos.POSManager) {
    listenerParam := model.AddListenerRequest_Param{}
    listenerParam.Subnqn = name
    listenerParam.TransportType = config.TransportType
    listenerParam.TargetAddress = config.TargetAddress
    listenerParam.TransportServiceId = config.TransportServiceId
    response, _ := CallAddListener(xrId, listenerParam, posMngr)
    //update response
    UpdateResponse(response, res, opName, errorInfoList)

}

func uramDevice(xrId string, config setting.HostConf, res *model.Response, name string, opName string, errorInfoList *[]map[string]interface{},posMngr pos.POSManager) {
    deviceParam := model.CreateDeviceRequest_Param{}
    deviceParam.Name = name
    deviceParam.NumBlocks = uint32(config.NumBlocks)
    deviceParam.BlockSize = uint32(config.BlockSize)
    deviceParam.DevType = config.DevType
    deviceParam.Numa = uint32(config.Numa)
    response, _ := CallCreateDevice(xrId, deviceParam, posMngr)
    //update response
    response.Result.Status.Description = name + ": " + response.Result.Status.Description
    UpdateResponse(response, res, opName, errorInfoList)

}

func RuniBoFOS(xrId string, param interface{},posMngr pos.POSManager) (model.Request, model.Response, error) {
    errorInfo := make(map[string]interface{})
    errorInfoList := []map[string]interface{}{}
    iBoFRequest := model.Request{
        Command: "RUNIBOFOS",
        Rid:     xrId,
    }

    iBoFRequest.Param = param
    res := model.Response{}
    _, pathErr := os.Stat("/usr/local/bin/poseidonos")
    if os.IsNotExist(pathErr) {
        res.Result.Status.Code = 11001
        return iBoFRequest, res, pathErr
    }
    path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
    cmd := fmt.Sprintf("/../script/run_remote_ibofos.sh %s", setting.Config.Server.IBoF.IP)
    err := utils.ExecCmd(path+cmd, false)

    if err != nil {
        res.Result.Status.Code = 11000
    } else {
        res.Result.Status.Code = 0
        res.LastSuccessTime = time.Now().UTC().Unix()
    }

    log.Info("RuniBoFOS result : ", res.Result.Status.Code)
    if res.Result.Status.Code == 0 {
        config := setting.Config.Server.IBoF
        errorCode := 0
        transport(xrId, config, &res, "transport", &errorInfoList, posMngr)
        subSystem(xrId, config, &res, config.Subsystem_1, "subSystem1", &errorInfoList, posMngr)
        subSystem(xrId, config, &res, config.Subsystem_2, "subSystem2", &errorInfoList, posMngr)
        listener(xrId, config, &res, config.Subsystem_1, "addListener1", &errorInfoList, posMngr)
        listener(xrId, config, &res, config.Subsystem_2, "addListener2", &errorInfoList, posMngr)
        uramDevice(xrId, config, &res, config.Uram1, "uram1", &errorInfoList, posMngr)
        uramDevice(xrId, config, &res, config.Uram2, "uram2", &errorInfoList, posMngr)
        for itr := 0; itr < len(errorInfoList); itr++ {
            if errorInfoList[itr]["code"].(int) != 0 {
                errorCode = 1
                break
            }
        }
        errorInfo["errorResponses"] = errorInfoList
        errorInfo["errorCode"] = errorCode
        res.Result.Status.ErrorInfo = errorInfo
    }
    return iBoFRequest, res, err
}

func CallStartPoseidonOS(xrId string, param interface{}, posMngr pos.POSManager) (model.Response, error) {
    //resByte, err := posMngr.StartPoseidonOS()
    _,res, err := RuniBoFOS(xrId,param,posMngr)
    fmt.Println("err>>>",err)
    fmt.Println("respo>>> ",res)
    return res,err
}

func CallStopPoseidonOS(xrId string, param interface{}, posMngr pos.POSManager) (model.Response, error) {
    result, err1 := posMngr.StopPoseidonOS()
    if err1 != nil {
        log.Errorf(commandFailureMsg, GetFuncName(1), err1)
        return model.Response{}, ErrConn
    }
    resByte, err2 := protojson.Marshal(result)
    return HandleResponse(resByte, err2)
}

func CallGetSystemInfo(xrId string, param interface{}, posMngr pos.POSManager) (model.Response, error) {
    result, err1 := posMngr.GetSystemInfo()
    if err1 != nil {
        log.Errorf(commandFailureMsg, GetFuncName(1), err1)
    }
    resByte, err2 := protojson.Marshal(result)
    return HandleResponse(resByte, err2)
}


