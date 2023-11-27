package response

type UserTokenResponse struct {
	Token string `json:"token"`
}

func UserToken(token string) UserTokenResponse {
	return  UserTokenResponse{Token: token}
}
