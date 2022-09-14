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

package m9k

import (
	"dagent/src/routers/m9k/api/dagent"
	"dagent/src/routers/m9k/api/ibofos"
	"dagent/src/routers/m9k/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"os"
	"path/filepath"
	//amoduleIBoFOS "pnconnector/src/routers/m9k/api/ibofos"
	"kouros/model"
	"kouros/utils"
	"reflect"
	"strings"
    "kouros"
    pos "kouros/pos"
    /*"fmt"
    "encoding/json"
    "google.golang.org/protobuf/encoding/protojson"*/
)

func Route(router *gin.Engine) {
	uri := router.Group("/api")
    posMngr,_ := kouros.NewPOSManager(pos.GRPC)
    posMngr.Init("dagent","127.0.0.1:50055")
    /*var mRes model.Response
    temp,_ := posMngr.ListDevices()
    //res1,_ := protojson.Marshal(temp)
    //fmt.Println(string(res1))
    if mErr := json.Unmarshal(temp, &mRes); mErr != nil {
        panic(mErr)
    }
    fmt.Println("out ",mRes,err)
    */
	// Doc Static
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dir = strings.ReplaceAll(dir, "/bin", "/doc")
	uri.StaticFS("/dagent/v1/doc", http.Dir(dir))

	//uri.Use(middleware.CheckBasicAuth())
	//uri.Use(middleware.CheckAPIActivate())
	uri.Use(middleware.CheckHeader)
	uri.Use(middleware.CheckBody)
	uri.Use(middleware.ResponseHeader)
	// D-Agent
	dagentPath := uri.Group("/dagent/v1")
	{
		dagentPath.GET("/heartbeat", func(ctx *gin.Context) {
			dagent.CallDagent(ctx, dagent.HeartBeat)
		})
		dagentPath.GET("/version", func(ctx *gin.Context) {
			dagent.CallDagent(ctx, dagent.Version)
		})
		dagentPath.DELETE("/dagent", func(ctx *gin.Context) {
			dagent.CallDagent(ctx, dagent.KillDAgent)
		})
		dagentPath.DELETE("/ibofos", func(ctx *gin.Context) {
			dagent.CallDagent(ctx, dagent.ForceKillIbof)
		})
	}

	// iBoFOSPath
	iBoFOSPath := uri.Group("/ibofos/v1")
	iBoFOSPath.Use(middleware.PostHandler)

	// System
	{
		iBoFOSPath.POST("/system", func(ctx *gin.Context) {
            ibofos.CalliBoFOS_new(ctx, ibofos.CallStartPoseidonOS, posMngr)
		})
		iBoFOSPath.DELETE("/system", func(ctx *gin.Context) {
            ibofos.CalliBoFOS_new(ctx, ibofos.CallStopPoseidonOS, posMngr)
		})
		iBoFOSPath.GET("/system", func(ctx *gin.Context) {
            ibofos.CalliBoFOS_new(ctx, ibofos.CallGetSystemInfo, posMngr)
		})
		iBoFOSPath.POST("/system/property", func(ctx *gin.Context) {
            ibofos.CalliBoFOS_new(ctx, ibofos.CallSetSystemProperty, posMngr)
		})
		iBoFOSPath.GET("/system/property", func(ctx *gin.Context) {
            ibofos.CalliBoFOS_new(ctx, ibofos.CallGetSystemProperty, posMngr)
		})
	}

	// Device
	{

		iBoFOSPath.GET("/devices", func(ctx *gin.Context) {
            ibofos.CalliBoFOS_new(ctx, ibofos.CallListDevices, posMngr)
		})
		iBoFOSPath.POST("/device", func(ctx *gin.Context) {
            //ibofos.CalliBoFOSwithParam_new(ctx, ibofos.CallCreateDevice, param, posMngr)
            ibofos.CalliBoFOS_new(ctx, ibofos.CallCreateDevice, posMngr)

		})
		iBoFOSPath.GET("/devices/:deviceName/scan", func(ctx *gin.Context) {
			deviceName := ctx.Param("deviceName")
			if deviceName == "all" {
                ibofos.CalliBoFOS_new(ctx, ibofos.CallScanDevice, posMngr)
			} else {
				// 404 return
			}
		})
		iBoFOSPath.GET("/devices/:deviceName/smart", func(ctx *gin.Context) {
			deviceName := ctx.Param("deviceName")
			param := model.DeviceParam{Name: deviceName}
            ibofos.CalliBoFOSwithParam_new(ctx, ibofos.CallGetDeviceSmartLog, param, posMngr)
		})
	}

	// Array
	{
		iBoFOSPath.POST("/array", func(ctx *gin.Context) {
			if validateNumOfDevice(ctx) {
				param := model.ArrayParam{}
				param.Name = ctx.Param("arrayName")
                ibofos.CalliBoFOSwithParam_new(ctx, ibofos.CallCreateArray, param, posMngr)

			}
		})
		iBoFOSPath.POST("/autoarray", func(ctx *gin.Context) {
			if validateNumOfDevice(ctx) {
				param := model.AutocreateArrayParam{}
                ibofos.CalliBoFOSwithParam_new(ctx, ibofos.CallAutoCreateArray, param, posMngr)
			}
		})
		iBoFOSPath.GET("/array/:arrayName", func(ctx *gin.Context) {
			param := model.ArrayParam{}
			param.Name = ctx.Param("arrayName")
            ibofos.CalliBoFOSwithParam_new(ctx, ibofos.CallArrayInfo, param, posMngr)
		})
		iBoFOSPath.GET("/arrays", func(ctx *gin.Context) {
			param := model.ArrayParam{}
            ibofos.CalliBoFOSwithParam_new(ctx, ibofos.CallListArray, param, posMngr)
		})
		iBoFOSPath.POST("/array/:arrayName/mount", func(ctx *gin.Context) {
			arrayName := ctx.Param("arrayName")
			param := model.ArrayParam{Name: arrayName}
            ibofos.CalliBoFOSwithParam_new(ctx, ibofos.CallMountArray, param, posMngr)
		})
		iBoFOSPath.DELETE("/array/:arrayName/mount", func(ctx *gin.Context) {
			arrayName := ctx.Param("arrayName")
			param := model.ArrayParam{Name: arrayName}
            ibofos.CalliBoFOSwithParam_new(ctx, ibofos.CallUnmountArray, param, posMngr)
		})
		iBoFOSPath.GET("/arrays/reset", func(ctx *gin.Context) {
			param := model.ArrayParam{}
            ibofos.CalliBoFOSwithParam_new(ctx, ibofos.CallResetMBR, param, posMngr)
		})

		iBoFOSPath.DELETE("/array/:arrayName", func(ctx *gin.Context) {
			param := model.ArrayParam{}
			param.Name = ctx.Param("arrayName")
            ibofos.CalliBoFOSwithParam_new(ctx, ibofos.CallDeleteArray, param, posMngr)
		})
		iBoFOSPath.POST("/array/:arrayName/devices", func(ctx *gin.Context) {
			param := model.ArrayParam{}
			param.Array = ctx.Param("arrayName")
            ibofos.CalliBoFOSwithParam_new(ctx, ibofos.CallAddDevice, param, posMngr)
		})
        iBoFOSPath.POST("/array/:arrayName/replace", func(ctx *gin.Context) {
            param := model.ArrayParam{}
            param.Array = ctx.Param("arrayName")
            ibofos.CalliBoFOSwithParam_new(ctx, ibofos.CallReplaceArrayDevice, param, posMngr)
        })
		iBoFOSPath.DELETE("/array/:arrayName/devices/:deviceName", func(ctx *gin.Context) {
			param := model.ArrayParam{}
			param.Array = ctx.Param("arrayName")
			param.Spare = []model.Device{{DeviceName: ctx.Param("deviceName")}}
            ibofos.CalliBoFOSwithParam_new(ctx, ibofos.CallRemoveDevice, param, posMngr)

		})
	}
	//Subsystem
	{
		iBoFOSPath.POST("/transport", func(ctx *gin.Context) {
			param := model.SubSystemParam{}
            ibofos.CalliBoFOSwithParam_new(ctx, ibofos.CallCreateTransport, param, posMngr)
		})

		iBoFOSPath.POST("/listener", func(ctx *gin.Context) {
			param := model.SubSystemParam{}
            ibofos.CalliBoFOSwithParam_new(ctx, ibofos.CallAddListener, param, posMngr)
		})
		iBoFOSPath.GET("/subsystem", func(ctx *gin.Context) {
            ibofos.CalliBoFOS_new(ctx, ibofos.CallListSubsystem, posMngr)
		})
		iBoFOSPath.POST("/subsystem", func(ctx *gin.Context) {
            ibofos.CalliBoFOS_new(ctx, ibofos.CallCreateSubsystem, posMngr)
		})
        iBoFOSPath.POST("/subsysteminfo", func(ctx *gin.Context) {
            ibofos.CalliBoFOS_new(ctx, ibofos.CallSubsystemInfo, posMngr)
        })
		iBoFOSPath.DELETE("/subsystem", func(ctx *gin.Context) {
			param := model.SubSystemParam{}
            ibofos.CalliBoFOSwithParam_new(ctx, ibofos.CallDeleteSubsystem, param, posMngr)
		})

	}
    /*
	// Volume
	{
		iBoFOSPath.POST("/volumes", func(ctx *gin.Context) {
			if multiVolRes, ok := dagent.IsMultiVolume(ctx); ok {
				req := model.Request{}
				ctx.ShouldBindBodyWith(&req, binding.JSON)
				reqMap := req.Param.(map[string]interface{})
				res := model.Response{}
				if reflect.TypeOf(reqMap["namesuffix"]).Kind() == reflect.String || reqMap["namesuffix"].(float64) < 0 {
					res.Result.Status, _ = util.GetStatusInfo(11060)
					ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, &res)
					return

				}
				if reflect.TypeOf(reqMap["size"]).Kind() == reflect.String || reqMap["size"].(float64) <= 0 {
					res.Result.Status, _ = util.GetStatusInfo(2033)
					ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, &res)
					return

				}
				dagent.ImplementAsyncMultiVolume(ctx, amoduleIBoFOS.CreateVolume, &multiVolRes, dagent.CREATE_VOLUME)
			} else {
				ibofos.CalliBoFOS(ctx, amoduleIBoFOS.CreateVolume)
			}
		})
		iBoFOSPath.GET("/volumelist/:arrayName", func(ctx *gin.Context) {
			arrayName := ctx.Param("arrayName")
			param := model.VolumeParam{Array: arrayName}
			ibofos.CalliBoFOSwithParam(ctx, amoduleIBoFOS.ListVolume, param)
		})
		iBoFOSPath.GET("/array/:arrayName/volume/:volumeName", func(ctx *gin.Context) {
			arrayName := ctx.Param("arrayName")
			volumeName := ctx.Param("volumeName")
			param := model.VolumeParam{Array: arrayName, Name: volumeName}
			ibofos.CalliBoFOSwithParam(ctx, amoduleIBoFOS.VolumeInfo, param)
		})
		iBoFOSPath.PATCH("/volumes/:volumeName", func(ctx *gin.Context) {
			volumeName := ctx.Param("volumeName")
			param := model.VolumeParam{Name: volumeName}
			ibofos.CalliBoFOSwithParam(ctx, amoduleIBoFOS.RenameVolume, param)
		})
		iBoFOSPath.GET("/volumes/maxcount", func(ctx *gin.Context) {
			ibofos.CalliBoFOS(ctx, amoduleIBoFOS.GetMaxVolumeCount)
		})
		iBoFOSPath.DELETE("/volumes/:volumeName", func(ctx *gin.Context) {
			volumeName := ctx.Param("volumeName")
			param := model.VolumeParam{Name: volumeName}
			ibofos.CalliBoFOSwithParam(ctx, amoduleIBoFOS.DeleteVolume, param)
		})
		iBoFOSPath.POST("/volumes/:volumeName/mount", func(ctx *gin.Context) {
			if multiVolRes, ok := dagent.IsMultiVolume(ctx); ok {
				dagent.ImplementAsyncMultiVolume(ctx, amoduleIBoFOS.MountVolume, &multiVolRes, dagent.MOUNT_VOLUME)
			} else {
				volumeName := ctx.Param("volumeName")
				param := model.VolumeParam{Name: volumeName}
				ibofos.CalliBoFOSwithParam(ctx, amoduleIBoFOS.MountVolume, param)
			}
		})
		iBoFOSPath.POST("/volumes/:volumeName/mount/subsystem", func(ctx *gin.Context) {
			volumeName := ctx.Param("volumeName")
			param := model.VolumeParam{Name: volumeName}
			ibofos.CalliBoFOSwithParam(ctx, amoduleIBoFOS.MountVolumeWithSubSystem, param)
		})
		iBoFOSPath.DELETE("/volumes/:volumeName/mount", func(ctx *gin.Context) {
			volumeName := ctx.Param("volumeName")
			param := model.VolumeParam{Name: volumeName}
			ibofos.CalliBoFOSwithParam(ctx, amoduleIBoFOS.UnmountVolume, param)
		})
		iBoFOSPath.PATCH("/volumes/:volumeName/qos", func(ctx *gin.Context) {
			volumeName := ctx.Param("volumeName")
			param := model.VolumeParam{Name: volumeName}
			ibofos.CalliBoFOSwithParam(ctx, amoduleIBoFOS.UpdateVolumeQoS, param)
		})
	}
	//QOS
	iBoFOSPath.POST("/qos", func(ctx *gin.Context) {
		ibofos.CalliBoFOS(ctx, amoduleIBoFOS.QOSCreateVolumePolicies)
	})
	iBoFOSPath.POST("/qos/reset", func(ctx *gin.Context) {
		ibofos.CalliBoFOS(ctx, amoduleIBoFOS.QOSResetVolumePolicies)
	})
	iBoFOSPath.POST("/qos/policies", func(ctx *gin.Context) {
		ibofos.CalliBoFOS(ctx, amoduleIBoFOS.QOSListPolicies)
	})*/

	//Telemetry
	iBoFOSPath.POST("/telemetry", func(ctx *gin.Context) {
        ibofos.CalliBoFOS_new(ctx, ibofos.CallStartTelemetry, posMngr)
	})
	iBoFOSPath.DELETE("/telemetry", func(ctx *gin.Context) {
        ibofos.CalliBoFOS_new(ctx, ibofos.CallStopTelemetry, posMngr)
	})

	// Logger Commands
	iBoFOSPath.POST("/logger/filter", func(ctx *gin.Context) {
        ibofos.CalliBoFOS_new(ctx, ibofos.CallApplyLogFilter, posMngr)
	})
	iBoFOSPath.GET("/logger/info", func(ctx *gin.Context) {
        ibofos.CalliBoFOS_new(ctx, ibofos.CallLoggerInfo, posMngr)
	})
	iBoFOSPath.POST("/logger/level", func(ctx *gin.Context) {
        level := ctx.Param("level")
        param := model.LoggerParam{Level: level}
        ibofos.CalliBoFOSwithParam_new(ctx, ibofos.CallSetLogLevel, param, posMngr)
	})
	iBoFOSPath.GET("/logger/level", func(ctx *gin.Context) {
        ibofos.CalliBoFOS_new(ctx, ibofos.CallGetLogLevel, posMngr)
	})
    iBoFOSPath.POST("/logger/preference", func(ctx *gin.Context) {
        ibofos.CalliBoFOS_new(ctx, ibofos.CallSetLogPreference, posMngr)
    })

	// Developer Commands
	iBoFOSPath.POST("/devel/event-wrr/reset", func(ctx *gin.Context) {
        ibofos.CalliBoFOS_new(ctx, ibofos.CallResetEventWRRPolicy, posMngr)
	})
	iBoFOSPath.POST("/devel/event-wrr/update", func(ctx *gin.Context) {
        ibofos.CalliBoFOS_new(ctx, ibofos.CallUpdateEventWRRPolicy, posMngr)

	})
	iBoFOSPath.DELETE("/devel/:arrayName/rebuild", func(ctx *gin.Context) {
		arrayName := ctx.Param("arrayName")
		param := model.ArrayParam{Name: arrayName}
        ibofos.CalliBoFOSwithParam_new(ctx, ibofos.CallStopRebuilding, param, posMngr)
	})

}
func validateNumOfDevice(ctx *gin.Context) bool {
	req := model.Request{}
	var numOfDevice int
	ctx.ShouldBindBodyWith(&req, binding.JSON)
	reqMap := req.Param.(map[string]interface{})
	res := model.Response{}
	if _, ok := reqMap["data"]; ok {
		numOfDevice = len(reqMap["data"].([]interface{}))
	} else {
		numOfDevice = int(reqMap["numData"].(float64))
	}
	if strings.ToLower(reqMap["raidtype"].(string)) == "raid10" && numOfDevice%2 != 0 {
		res.Result.Status, _ = util.GetStatusInfo(2517)
		ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, &res)
		return false
	}
	return true
}
