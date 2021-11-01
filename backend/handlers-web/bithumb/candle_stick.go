package bithumb

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"toTheMoon/backend/cerror"
)

type ReqCandleStick struct {
	Ticker string `json:"ticker" example:"BTC"`
	ChartInterval string `json:"chart_interval" example:"ex) 1m, 1H ,24H"`
}


type Result struct {
	Time        int     `json:"time"`
	MarketPrice float64 `json:"market_price"`
}

type CandleStickResp struct {
	Status string `json:"status"`
	Data   []interface{}
}


// @tags bithumb
// @Summary 빗썸 캔들스틱 API
// @Description 빗썸 캔들스틱 API
// @Param body body ReqCandleStick true "티커, 차트간격 기준"
// @Accept json
// @Produce json
// @Success 200 {string} Hello World
// @Router /api/web/bithumb/candle_stick [post]
func CandleStick(c *gin.Context) {


	var body ReqCandleStick
	if err := c.ShouldBind(&body); err != nil {
		panic(cerror.BadRequestWithMsg("binding error"))
	}

	ticker := body.Ticker
	chartInterval := body.ChartInterval


	reqURL := fmt.Sprintf("https://api.bithumb.com/public/candlestick/%s_KRW/%s", ticker, chartInterval)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
	var candleStick CandleStickResp
	if err := json.Unmarshal(bytes, &candleStick); err != nil {
		fmt.Println(err)
		panic(err)
	}

	FetchNum := 5



	//result := make([]Result,FetchNum)
	fmt.Println(candleStick.Status)
	for i := len(candleStick.Data); i > len(candleStick.Data)-FetchNum; i-- {

		fmt.Println(candleStick.Data[i-1])
	}


	c.JSON(http.StatusOK, candleStick)
}
