package main

import (
	"encoding/json"
	"fmt"
	// "github.com/tommitoan/bazica"
	"github.com/phawitb/bazi-go"
	"time"
)

func main() {
	// Calculate current ba-zi chart
	// loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	// now := time.Now()
	// gender := 0 // 0 = female & 1 = male
	
	// /* Example: 
	// Time to calculate: 1990-12-31 6:30 - Timezone: HoChiMinh / Vietnam - Gender: Male
	// loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	loc, _ := time.LoadLocation("Asia/Bangkok")
	// now := time.Date(1990, time.Month(12), 31, 6, 30, 0, 0, loc)
	// now := time.Date(1991, time.Month(1), 1, 12, 0, 0, 0, loc)
	now := time.Date(2032, time.Month(4), 22, 2, 30, 0, 0, loc)
	gender := 0
	// */

	
	

	chart, err := bazica.GetBaziChart(now, loc, gender)
	if err != nil {
		fmt.Println(err)
	}
	jsonData, _ := json.Marshal(chart)
	fmt.Println(string(jsonData))
}
