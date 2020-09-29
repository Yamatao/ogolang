package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	// getting the host current time
	hostTime := time.Now().Round(time.Second)

	fmt.Printf("current time: %+v\n", hostTime)

	// getting ntp time
	url := "2.us.pool.ntp.org"

	ntpTime, err := ntp.Time(url)
	if err != nil {
		log.Fatalf("failed to query time from the '%s': %v", url, err)
	}
	ntpTime = ntpTime.Round(time.Second)

	fmt.Printf("exact time: %+v\n", ntpTime)
}
