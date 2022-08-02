package model

type BankAndCoin struct {
	TotalChange             float64 `param:"totalChange" json:"totalChange"  form:"totalChange"`
	OneThousandBahtNote     int     `param:"1000" json:"1000"  form:"1000"`
	FiveHundredBahtNote     int     `param:"500" json:"500"  form:"500"`
	HundredBahtNote         int     `param:"100" json:"100"  form:"100"`
	FiftyBahtNote           int     `param:"50" json:"50"  form:"50"`
	TwentyBahtNote          int     `param:"20" json:"20"  form:"20"`
	TenBahtCoin             int     `param:"10" json:"10"  form:"10"`
	FiveBahtCoin            int     `param:"5" json:"5"  form:"5"`
	OneBahtCoin             int     `param:"1" json:"1"  form:"1"`
	TwentyFiveStangBahtCoin int     `param:"0.25" json:"0.25"  form:"0.25"`
}

type PaymentAndPrice struct {
	Payment float64 `param:"payment" json:"payment"  form:"payment"`
	Price   float64 `param:"price" json:"price"  form:"payment"`
}

func (b *BankAndCoin) SetChangeList(changeList map[float64]int) {
	b.OneThousandBahtNote = changeList[1000]
	b.FiveHundredBahtNote = changeList[500]
	b.HundredBahtNote = changeList[100]
	b.FiftyBahtNote = changeList[50]
	b.TwentyBahtNote = changeList[20]
	b.TenBahtCoin = changeList[10]
	b.FiveBahtCoin = changeList[5]
	b.OneBahtCoin = changeList[1]
	b.TwentyFiveStangBahtCoin = changeList[0.25]
}

func (b *BankAndCoin) SetTotalChange(totalChange float64) {
	b.TotalChange = totalChange
}
