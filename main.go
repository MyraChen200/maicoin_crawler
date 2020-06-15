package main

import "fmt"
import "net/http"
import "io/ioutil"
// import "encoding/json"
import "github.com/tidwall/gjson"


func main() {
	currency_api_endpoint := "https://max-api.maicoin.com/api/v2/summary"
	req, err := http.Get(currency_api_endpoint)

	if err != nil {
		fmt.Println("error %s",err)
	}
	defer req.Body.Close()
	
	body, err := ioutil.ReadAll(req.Body)
	bodyString := string(body)
	currency := gjson.Get(bodyString, "tickers.ethtwd.buy")
        xxx := gjson.Get(bodyString, "tickers.ethtwd")
	fmt.Println(xxx)
        fmt.Println("目前價格: ", currency)
}
