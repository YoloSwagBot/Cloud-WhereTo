package main

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func midnightPSTJob() {
	// This function will be executed at midnight PST
	fmt.Println("Cron job executed at NOW PST:", time.Now())
}

func main() {
	// Create a new cron instance
	c := cron.New()

	// Schedule the job to run every day at midnight PST
	// "0 0 * * *" means midnight every day
	// "America/Los_Angeles" is the time zone for PST
	_, err := c.AddFunc("47 6 * * *", midnightPSTJob).SetTimeZone(time.FixedZone("America/Los_Angeles", -8*60*60))
	if err != nil {
		log.Fatal("Error scheduling cron job:", err)
	}

	// Start the cron scheduler
	c.Start()

	// Run the program indefinitely
	select {}
}