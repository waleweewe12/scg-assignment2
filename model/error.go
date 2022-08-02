package model

type Error struct {
	TotalChange float64
	Message     string
}

func (e *Error) SetError(change float64) {
	if change > 0 {
		e.Message = "Bank and coin not enougth"
	} else {
		e.Message = "payment less than price"
	}
}

func (e *Error) SetTotalChange(totalChange float64) {
	e.TotalChange = totalChange
}
