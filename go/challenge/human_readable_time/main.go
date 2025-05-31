package main

import "fmt"

func main() {
	fmt.Println(HumanReadableTime1(10000))
	fmt.Println(HumanReadableTime1(10000))
}

func HumanReadableTime1(seconds int) string {
	seconds, minutes, hours := seconds%60, seconds/60%60, seconds/3600
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func HumanReadableTime2(s int) string {
    m, s := s / 60, s % 60
    h, m := m / 60, m % 60
    return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
