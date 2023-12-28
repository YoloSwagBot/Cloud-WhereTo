package main

// 1
import (
 "fmt"
 "time"

 "github.com/go-co-op/gocron"
)


func runCronJobs() {
 // 3
 s := gocron.NewScheduler(time.UTC)

 // 4
 s.Every(12).Hours().Do(func() {
  fmt.Println("Hello, there world.")
 })

 // 5
 s.StartBlocking()
}

// 6
func main() {
 runCronJobs()
}