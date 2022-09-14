package ibofos
import (
    "kouros/model"
    "runtime"
    "strings"
    "reflect"
    "encoding/json"
    "fmt"
    "kouros/log"
    "errors"
)
var (
    errLocked    = errors.New("Locked out buddy")
    ErrSending   = errors.New("Sending error")
    ErrReceiving = errors.New("Receiving error")
    ErrJson      = errors.New("Json error")
    ErrRes       = errors.New("Response error")
    ErrConn      = errors.New("Could not connect to the CLI server. Is PoseidonOS running?")
    ErrJsonType  = errors.New("Json Type Validation Error")
    unmarshalErrMsg   = "Could not Unmarshal json response in "
    marshalErrMsg     = "Failed to marshal the param structure in "
    commandFailureMsg = "Received error from grpc server in "

)


func GetFuncName(level int) string {
    pc, _, _, ok := runtime.Caller(level)
    if !ok {
        return ""
    }

    fn := runtime.FuncForPC(pc)
    if fn == nil {
        return ""
    }
    //fmt.Println("fun name ", fn.Name())
    nameArr := strings.Split(fn.Name(), ".")
    name := nameArr[len(nameArr)-1]
    return name
}

func HandleResponse(resByte []byte, err error) (model.Response, error) {
    var res model.Response
    if err != nil {
        log.Errorf(marshalErrMsg, GetFuncName(2), err)
        return model.Response{}, err
    }
    fmt.Println("response ",string(resByte),reflect.TypeOf(resByte))
    if err1 := json.Unmarshal(resByte, &res); err1 != nil {
        log.Errorf(unmarshalErrMsg, GetFuncName(2), err1)
        return model.Response{}, err1
    }
    fmt.Println("trace 2", GetFuncName(2))
    //fmt.Println("response ", res)
    fmt.Println("resByte string", string(resByte))
    return res, nil
}

