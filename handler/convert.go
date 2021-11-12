package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glavanan/talanddev/model"
	"github.com/glavanan/talanddev/service"
)

func Convert(c *gin.Context) {
	var convert model.Convert
	fmt.Printf("In Handler")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	err = json.Unmarshal(body, &convert)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	service.Convert(convert)
}
