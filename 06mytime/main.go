package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcoem to My Time in goLang")
	presentTime := time.Now()
	fmt.Println("Present Time is: ", presentTime)
	fmt.Println("Present Time is: ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.September, 9, 23, 12, 0, 0, time.UTC) // year, month, day, hour, minute, second, nanosecond, location
	fmt.Println("Created Date is: ", createdDate.Format("01-02-2006 Monday")) 
}