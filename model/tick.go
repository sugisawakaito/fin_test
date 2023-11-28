package model

type Tick struct {
	Open int
	High int
	Low int
	Close int
}

func TickNomarize(tickValues []int) *Tick {
	open := 0
	close := 0
       	low := tickValues[0]
	high := 0
	for i, t := range tickValues {
		if i == 0 {
			open = t
		}
		if i == len(tickValues) -1 {
			close = t
		}
		if low > t {
			low = t
		}
		if  high < t {
			high = t
		}
	}
	tick := Tick{
		Open: open,
		High: high,
		Low: low,
		Close: close,
	}

	return &tick
}
