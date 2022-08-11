// This code is genetated by protoc-gen-gokit-micro.

// The MIT License (MIT)

// Copyright (c) 2022 leewckk@126.com

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// PLUGIN: protoc-gen-gokit-micro
//		VERSION : 0.0.1-1-g3674669
//		GIT TAG : 0.0.1-1-g3674669
//		GIT HASH : 3674669d9e78d07a85347b1a1fa579727193f466
//		BUILDER VERSION : go version go1.16.5 linux/amd64
//		BUILD TIME: : 2022-08-11 17:04:24

// create time : 2022-08-11 17:04:33.050783319 +0800 CST m=+0.011101212

package calculate

import (
	fmt "fmt"
	gin "github.com/gin-gonic/gin"
	http1 "github.com/go-kit/kit/transport/http"
	http3 "github.com/leewckk/go-kit-micro-service/middlewares/transport/http"
	http2 "github.com/leewckk/go-kit-micro-service/router/http"
	calculate "micro-service-sample/endpoint/calculate"
	http "net/http"
)

type MessagingServiceTransport struct {
	options []http1.ServerOption
}

func RegisterMessagingProc(r *http2.Router, opts ...http3.ServerOption) {

	objs := NewMessagingTransport(opts...)
	for _, obj := range objs {
		r.Register(obj)
	}

}

func NewMessagingTransport(opts ...http3.ServerOption) []*http3.RouterObject {

	objects := make([]*http3.RouterObject, 0, 0)
	objects = append(objects,
		http3.NewTransportObject("GET",
			"/v1/messages/:messageId",
			http3.NewServer(calculate.MakeMessagingGetMessageEndpoint(),
				DecodeMessagingGetMessageRequestGET,
				EncodeMessagingGetMessageResponse,
				opts...)))

	objects = append(objects,
		http3.NewTransportObject("GET",
			"/v1/users/:userId/messages/:messageId",
			http3.NewServer(calculate.MakeMessagingGetMessageEndpoint(),
				DecodeMessagingGetMessageRequestGET,
				EncodeMessagingGetMessageResponse,
				opts...)))

	return objects
}

//// THIS IS FOR SERVICE DEOCDE REQUEST
//// CLIENT.REQUEST -> SERVICE.REQUEST
func DecodeMessagingGetMessageRequestGET(c *gin.Context) (interface{}, error) {
	var request calculate.GetMessageRequest

	if err := c.ShouldBindUri(&request); err == nil {
		return &request, nil
	} else {
		return nil, fmt.Errorf("Error decode request , uri: %v, err: %v", c.Request.URL, err)
	}
}

//// THIS IS FOR SERVICE DEOCDE REQUEST
//// CLIENT.REQUEST -> SERVICE.REQUEST
func DecodeMessagingGetMessageRequestJSON(c *gin.Context) (interface{}, error) {
	var request calculate.GetMessageRequest

	if err := c.ShouldBindJSON(&request); err == nil {
		return &request, nil
	} else {
		return nil, fmt.Errorf("Error decode request , uri: %v, err: %v", c.Request.URL, err)
	}
}

////

//// THIS IS FOR SERVICE ENCODE RESPONSE
//// SERVICE.RESPONSE -> CLIENT.RESPONSE
func EncodeMessagingGetMessageResponse(c *gin.Context, resp interface{}) error {
	c.JSON(http.StatusOK, resp)
	return nil
}

////
