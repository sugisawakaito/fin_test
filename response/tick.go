package response

import (
	"fin_test/model"
)

type TickInfoResponse struct {
	Open int `json:"open"`
	Close int `json:"close"`
	High int `json:"high"`
	Low int `json:"low"`
}


func GetTickInfo(tick *model.Tick) TickInfoResponse {
	res := TickInfoResponse{
		Open: tick.Open,
		Close: tick.Close,
		High: tick.High,
		Low: tick.Low,
	}

	return res
}
