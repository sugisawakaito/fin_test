package request

import (
        "github.com/labstack/echo"
)

type UserLoginBody struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func GetUserLoginBody(c echo.Context) (*UserLoginBody, error)  {
	var body UserLoginBody
	if err := c.Bind(&body); err != nil {
		return nil, err
	}
	return &body, nil
}
