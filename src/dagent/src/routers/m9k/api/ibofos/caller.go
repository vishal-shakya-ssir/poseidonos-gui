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
 *     * Neither the name of Intel Corporation nor the names of its
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
	"pnconnector/src/log"
	"pnconnector/src/routers/m9k/model"
	"bytes"
	"dagent/src/routers/m9k/api"
	"dagent/src/routers/m9k/header"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"fmt"
)

func CalliBoFOS(ctx *gin.Context, f func(string, interface{}) (model.Request, model.Response, error)) {
	req := model.Request{}
	ctx.ShouldBindBodyWith(&req, binding.JSON)
	_, res, err := f(header.XrId(ctx), req.Param)
	fmt.Println("response >>>>>==========:",res)
	api.HttpResponse(ctx, res, err)
}

func CalliBoFOSwithParam(ctx *gin.Context, f func(string, interface{}) (model.Request, model.Response, error), param interface{}) {
	req := model.Request{}
	ctx.ShouldBindBodyWith(&req, binding.JSON)

	if req.Param != nil {
		param = merge(param, req.Param)
	}

	_, res, err := f(header.XrId(ctx), param)
	api.HttpResponse(ctx, res, err)
}

func merge(src interface{}, tar interface{}) interface{} {
	var m map[string]interface{}

	ja, _ := json.Marshal(src)
	json.Unmarshal(ja, &m)

	jb, _ := json.Marshal(tar)
	json.Unmarshal(jb, &m)

	jm, _ := json.Marshal(m)

	var param interface{}

	d := json.NewDecoder(bytes.NewBuffer(jm))
	d.UseNumber()

	if err := d.Decode(&param); err != nil {
		log.Fatal(err)
	}

	return param
}
