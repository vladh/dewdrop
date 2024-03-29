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
	d, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, 0x76)
	if err != nil {
		panic(err)
	}

	b := bme280.New(d)
	err = b.Init()
	temp, pres, hum, err := b.EnvData()
	if err != nil {
		panic(err)
	}

	fmt.Printf("date=%s:00%s ",
		time.Now().Format("2006-01-02T15:04"),
		time.Now().Format("Z07:00"))
	fmt.Printf("temp=%f ", temp)
	fmt.Printf("pres=%f ", pres)
	fmt.Printf("hum=%f\n", hum)
}
