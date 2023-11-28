package handler

import (
	        "github.com/labstack/echo"
		"fin_test/request"
		"fin_test/service"
		"fin_test/response"

)

func GetTickInfo(c echo.Context) error {
	params, err :=  request.GetTickInfoParams(c)
	if err != nil {
		c.JSON(400, nil)
	}


	tickInfo, err := service.GetTickInfo(params.Code, params.Year, params.Month, params.Day, params.Hour)


	c.JSON(200, response.GetTickInfo(tickInfo))
	return nil
}
