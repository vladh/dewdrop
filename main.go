// © 2022 Vlad-Stefan Harbuz <vlad@vladh.net>
// SPDX-License-Identifier: MIT

package main

import (
	"time"
	"fmt"

	"golang.org/x/exp/io/i2c"

	"github.com/quhar/bme280"
)

func main() {
	d, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, bme280.I2CAddr)
	if err != nil {
		panic(err)
	}

	b := bme280.New(d)
	err = b.Init()
	temp, pres, hum, err := b.EnvData()
	if err != nil {
		panic(err)
	}

	fmt.Printf("date=%s ", time.Now().Format(time.RFC3339))
	fmt.Printf("temp=%f ", temp)
	fmt.Printf("pres=%f ", pres)
	fmt.Printf("hum=%f\n", hum)
}
