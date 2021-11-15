package t

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/glavanan/talanddev/dbconnector"
	"github.com/glavanan/talanddev/model"
	"github.com/glavanan/talanddev/router"
)

var (
	StatusOK      = 200
	StatusBR      = 400
	BodyConvertOk = `{
		"amount": 500,
		"currency": "EUR",
		"convert_to": "ETB",
		"date": "2021-11-12T00:00:00Z"
	}`
	BodyConvertNotOk = `{
		"amount": 500,
		"currency": "ZZZ",
		"convert_to": "ETB",
		"date": "2021-11-12T00:00:00Z"
	}`
)

func TestPostOK(t *testing.T) {
	//Put the dns of a fake DB
	dbconnector.InitDB("host=0.0.0.0 user=postgres password=secretpass dbname=talanddev port=5432")
	ts := httptest.NewServer(router.GetRouter())
	defer ts.Close()
	request, err := http.NewRequest("POST", ts.URL+"/convert", bytes.NewBuffer([]byte(BodyConvertOk)))
	if err != nil {
		t.Errorf("create request fail : " + err.Error())
	}
	resp, err := (&http.Client{}).Do(request)
	if err != nil {
		t.Errorf("request fail : " + err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != StatusOK {
		t.Errorf("Get domain status didn't return StatusOK")
	}
	body, _ := io.ReadAll(resp.Body)
	var historyRep model.History
	err = json.Unmarshal(body, &historyRep)
	if err != nil {
		t.Errorf("Unmarshal fail : " + err.Error())
	}
	if historyRep.Amount != 500 || historyRep.Result != 27073.9775 {
		t.Errorf("Result is not what we expected")
	}
}

func TestPostNotOK(t *testing.T) {
	//Put the dns of a fake DB
	dbconnector.InitDB("host=0.0.0.0 user=postgres password=secretpass dbname=talanddev port=5432")
	ts := httptest.NewServer(router.GetRouter())
	defer ts.Close()
	request, err := http.NewRequest("POST", ts.URL+"/convert", bytes.NewBuffer([]byte(BodyConvertNotOk)))
	if err != nil {
		t.Errorf("create request fail : " + err.Error())
	}
	resp, err := (&http.Client{}).Do(request)
	if err != nil {
		t.Errorf("request fail : " + err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != StatusBR {
		t.Errorf("Get domain status didn't return Status Bad request")
	}
}
