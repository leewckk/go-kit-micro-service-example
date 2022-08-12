package service

import (
	"context"
	"encoding/json"
	"micro-service-sample/endpoint/calculate"
	"net/http"

	"github.com/sirupsen/logrus"
)

func init() {

	// func ImplementCalculateAdd(impl func() CalculateAdd) {
	// 	__CalculateAddProc__ = impl()
	// }
	calculate.ImplementCalculateAdd(func() calculate.CalculateAdd {
		return func(ctx context.Context, request *calculate.AddRequest) (*calculate.CalculateResult, error) {
			resp := calculate.CalculateResult{
				Status: http.StatusOK,
				Msg:    "ok",
				Result: request.Element0 + request.Element1,
			}
			js, _ := json.MarshalIndent(request, "", "  ")
			logrus.Info(string(js))
			return &resp, nil
		}
	})

}
