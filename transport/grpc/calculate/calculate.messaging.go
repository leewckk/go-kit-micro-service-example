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

// create time : 2022-08-11 17:04:33.05022532 +0800 CST m=+0.010543209

package calculate

import (
	context "context"
	fmt "fmt"
	grpc "github.com/go-kit/kit/transport/grpc"
	grpc1 "google.golang.org/grpc"
	calculate1 "micro-service-sample/endpoint/calculate"
	calculate "micro-service-sample/pb/golang/pkg/calculate"
	reflect "reflect"
)

type MessagingServiceTransport struct {
	__GetMessage__ grpc.Handler
}

func RegisterMessagingProc(s *grpc1.Server, opts ...grpc.ServerOption) {

	calculate.RegisterMessagingServer(s, NewMessagingTransport(opts...))
}

func NewMessagingTransport(opts ...grpc.ServerOption) *MessagingServiceTransport {

	var t MessagingServiceTransport
	t.__GetMessage__ = grpc.NewServer(
		calculate1.MakeMessagingGetMessageEndpoint(),
		DecodeMessagingGetMessageRequest,
		EncodeMessagingGetMessageResponse,
		opts...,
	)
	return &t
}

///
func (this *MessagingServiceTransport) GetMessage(ctx context.Context, request *calculate.GetMessageRequest) (*calculate.GetMessageResponse, error) {
	_, resp, err := this.__GetMessage__.ServeGRPC(ctx, request)
	if nil != err {
		return nil, err
	}
	return resp.(*calculate.GetMessageResponse), nil
}

///

//// THIS IS FOR SERVICE DEOCDE REQUEST
//// PB.REQUEST -> SERVICE.REQUEST
func DecodeMessagingGetMessageRequest(ctx context.Context, request interface{}) (interface{}, error) {
	if req, ok := request.(*calculate.GetMessageRequest); ok {
		r := calculate1.GetMessageRequest{}
		/// TODO
		r.MessageId = req.MessageId
		r.UserId = req.UserId
		return &r, nil
	}
	return nil, fmt.Errorf("Error convert interface : %v -> *%v ", reflect.TypeOf(request), "calculate.GetMessageRequest")
}

////

//// THIS IS FOR SERVICE ENCODE RESPONSE
//// SERVICE.RESPONSE -> PB.RESPONSE
func EncodeMessagingGetMessageResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	if response, ok := resp.(*calculate1.GetMessageResponse); ok {
		r := calculate.GetMessageResponse{}
		/// TODO
		r.MessageId = response.MessageId
		r.UserId = response.UserId
		r.Text = response.Text
		return &r, nil
	}
	return nil, fmt.Errorf("Error convert interface : %v -> *%v", reflect.TypeOf(resp), "calculate1.GetMessageResponse")
}

////
