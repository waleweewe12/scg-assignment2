package main

import (
	"assignment2/model"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
)

//fake DB
var bankAndCoin = map[float64]int{
	1000: 10,
	500:  20,
	100:  15,
	50:   20,
	20:   30,
	10:   20,
	5:    20,
	1:    20,
	0.25: 50,
}

var bankAndCoinSeq = []float64{1000, 500, 100, 50, 20, 10, 5, 1, 0.25}

func calculate(payment float64, productPrice float64) (float64, map[float64]int) {
	change := payment - productPrice
	changeList := map[float64]int{}
	restore := make([]int, 0)
	if change < 0 {
		return change, changeList
	}
	for _, v := range bankAndCoinSeq {
		res := int(math.Floor(change / v))
		if bankAndCoin[v]-res < 0 {
			res = bankAndCoin[v]
			bankAndCoin[v] = 0
		} else {
			bankAndCoin[v] -= res
		}
		change -= v * float64(res)
		restore = append(restore, res)
		changeList[v] = res
	}
	//restore when bank and coin not enougth
	if change != 0 {
		for i, v := range bankAndCoinSeq {
			bankAndCoin[v] += restore[i]
		}
	}
	return change, changeList
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) (err error) {
		return c.String(http.StatusOK, "health check !")
	})
	e.POST("/change", func(c echo.Context) (err error) {
		pp := new(model.PaymentAndPrice)
		if err = c.Bind(pp); err != nil {
			return
		}
		change, changeList := calculate(pp.Payment, pp.Price)
		if change != 0 {
			e := new(model.Error)
			e.SetTotalChange(pp.Payment - pp.Price)
			e.SetError(change)
			return c.JSON(http.StatusBadRequest, e)
		}
		b := new(model.BankAndCoin)
		b.SetTotalChange(pp.Payment - pp.Price)
		b.SetChangeList(changeList)
		return c.JSON(http.StatusOK, b)
	})
	e.Logger.Fatal(e.Start(":9000"))
}
