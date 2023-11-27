package handler

import (
	"fmt"
	"fin_test/request"
	"fin_test/service"
	"fin_test/response"
        "github.com/labstack/echo"
)

func UserLogin(c echo.Context) error {
	body, err := request.GetUserLoginBody(c)
	if err != nil {
		c.JSON(400, nil)
	}

	encriptedUserInfo := service.GetEncriptedUserInfo(body.UserName, body.Password)
	c.JSON(200, response.UserToken(encriptedUserInfo))
	return nil
}

func GetFlag(c echo.Context) error {
	body, err := request.GetFlagBody(c)
	if err != nil {
		c.JSON(400, nil)
	}

	fmt.Println("##############################")
	fmt.Println("flag is: %s", body.Flag)
	fmt.Println("##############################")

	c.JSON(200, nil)
	return nil
}
