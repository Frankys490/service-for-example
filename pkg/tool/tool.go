package tool

import (
	"errors"
	"net/http"
	"resenje.org/logging"
	"service-for-example/internal/model"
	"service-for-example/pkg/custom_errors"
)

func Response(err error, httpRes any) (res model.Response) {
	if err == nil {
		res.Data = httpRes
		res.Code = http.StatusOK
		res.Message = "Successful"
		return res
	}

	var e *custom_errors.ErrHttp

	logging.Info(err.Error())

	if errors.As(err, &e) {
		var finalErrMessage string

		for err != nil {
			finalErrMessage = err.Error()
			err = errors.Unwrap(err)
		}

		res.Code = e.Code
		res.Message = finalErrMessage

		return
	}

	res.Code = custom_errors.ErrInternal.Code
	res.Message = custom_errors.ErrInternal.Error()

	return
}
