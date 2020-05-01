// Level: medium
//
// Source: https://www.youtube.com/watch?v=3Q_oYDQ2whs
//
// There are 2 persons, each has his own calendar for the day and available work hours.
// Find: all available slots in calendar for these 2 persons. Not just for meeting, all free time for scheduling.
// Like scheduling assistant in Outlook.
//
// Sample input:
// Person 1
// Calendar [[9:00, 10:30], [12:00, 13:00], [16:00, 18:00]]
// Work Hours [9:00, 20:00]
//
// Person 2
// Calendar [[10:00, 11:30], [12:30, 14:30], [14:30, 15:00], [16:00, 17:00]]
// Work Hours [10:00, 18:30]
//
// Meeting: 30 мin
//
// Sample output:
// [[11:30, 12:00], [15:00, 16:00], [18:00, 18:30]]
//
package main

import "fmt"

func main() {
	fmt.Println(findFreeSlots(
		[][]string{
			{"9:00", "10:30"},
			{"12:00", "13:00"},
			{"16:00", "18:00"},
		},
		[][]string{
			{"10:00", "11:30"},
			{"12:30", "14:30"},
			{"14:30", "15:00"},
			{"16:00", "17:00"},
		},
		"30",
	))
}

func findFreeSlots(p1Cal, p2Cal [][]string, meetingDuration string) [][]string {
	return [][]string{
		{"10:00", "11:30"},
		{"12:30", "14:30"},
		{"14:30", "15:00"},
		{"16:00", "17:00"},
	}
}
