package service

import (
	"fmt"

	"github.com/glavanan/talanddev/model"
)

func Convert(conv model.Convert) error {
	fmt.Printf("%+v", conv)
	return nil
}
