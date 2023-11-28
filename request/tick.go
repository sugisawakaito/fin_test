package request

import (
        "github.com/labstack/echo"
)

type TickInfoParams struct {
	 Code string `query:"code"`
	Year int `query:"year"`
	Month int `query:"month"`
	Day int `query:"day"`
	Hour int `query:"hour"`
}

func GetTickInfoParams(c echo.Context) (*TickInfoParams, error) {
	var params TickInfoParams
	if err := c.Bind(&params); err != nil {
		return nil, err
	}
	return &params, nil
}
