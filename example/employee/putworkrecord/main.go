package main

import (
	"github.com/krrrr38/go-freeehr/freeehr"
	"golang.org/x/oauth2"

	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	clientID := os.Getenv("FREEE_CLIENT_ID")
	if clientID == "" {
		log.Fatal("FREEE_CLIENT_ID env variable not found")
	}
	clientSecret := os.Getenv("FREEE_CLIENT_SECRET")
	if clientSecret == "" {
		log.Fatal("FREEE_CLIENT_SECRET env variable not found")
	}
	employeeID, err := strconv.Atoi(os.Getenv("EMPLOYEE_ID")) // 294712
	if err != nil {
		log.Fatal("EMPLOYEE_ID env variable parse failed", err)
	}

	conf := freeehr.Conf(clientID, clientSecret, "urn:ietf:wg:oauth:2.0:oob")

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v\nEnter code: ", url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("code: %v\n", code)

	token, err := conf.Exchange(oauth2.NoContext, code) // get access token and so on
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("access token: %v\n", token)

	client, _ := freeehr.NewClient(conf.Client(oauth2.NoContext, token))

	clockInAt := time.Now()
	date := freeehr.FreeeDate{clockInAt}

	breakRecordClockInAt := time.Date(clockInAt.Year(), clockInAt.Month(), clockInAt.Day(), 12, 00, 0, 0, time.Local)
	breakRecordClockOutAt := time.Date(clockInAt.Year(), clockInAt.Month(), clockInAt.Day(), 13, 00, 0, 0, time.Local)
	temporaryClockOutAt := time.Date(clockInAt.Year(), clockInAt.Month(), clockInAt.Day(), 18, 30, 0, 0, time.Local)

	// put breakRecords if clockInAt is before than defaut breakRecordClockInAt
	breakRecords := []freeehr.BreakRecord{}
	if clockInAt.Before(breakRecordClockInAt) {
		breakRecords = append(breakRecords, freeehr.BreakRecord{
			ClockInAt:  &freeehr.FreeeDateTime{breakRecordClockInAt},
			ClockOutAt: &freeehr.FreeeDateTime{breakRecordClockOutAt},
		})
	}

	newWorkRecord, _, err := client.Employees.PutWorkRecord(employeeID, date, &freeehr.WorkRecord{
		BreakRecords: breakRecords,
		ClockInAt:    &freeehr.FreeeDateTime{clockInAt},
		ClockOutAt:   &freeehr.FreeeDateTime{temporaryClockOutAt},
	})
	if err != nil {
		fmt.Printf("put error: %v\n", err)
	} else {
		fmt.Printf("put workRecord: %v\n", newWorkRecord)
	}
}
