package v3

import (
	"fmt"
	"golang.org/x/text/currency"
	"io"
	"net/http"
	"os"
	"testing"
)

var key string

func init() {
	var file, err0 = os.Open("key")
	if err0 != nil {
		panic(err0)
	}

	var data, err1 = io.ReadAll(file)
	if err1 != nil {
		panic(err1)
	}

	key = string(data)
}

func Test(t *testing.T) {
	var httpClient = new(http.Client)
	var client, err0 = NewClient(key, httpClient)
	if err0 != nil {
		t.Fatal(err0)
	}

	var resp0, err1 = client.Status()
	if err1 != nil {
		t.Fatal(err1)
	}

	fmt.Printf("%+v", *resp0)
	fmt.Println("")

	var resp1, err2 = client.Latest(LatestRequest{
		From: currency.EUR.String(),
		To: []string{
			currency.AUD.String(),
			currency.JPY.String(),
		},
	})
	if err2 != nil {
		t.Fatal(err2)
	}

	fmt.Printf("%+v", *resp1)
	fmt.Println("")
}
