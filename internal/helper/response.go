package helper

import "github.com/labstack/echo/v4"

type Response struct {
	ResponseCode    int         `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	ResponseData    interface{} `json:"responseData"`
}

func ResponseSuccess(c echo.Context, code int, message string, data interface{}) error {
	return c.JSON(code, Response{
		ResponseCode:    code,
		ResponseMessage: message,
		ResponseData:    data,
	})
}

func ResponseError(c echo.Context, code int, message string, data interface{}) error {
	return c.JSON(code, Response{
		ResponseCode:    code,
		ResponseMessage: message,
		ResponseData:    data,
	})
}
