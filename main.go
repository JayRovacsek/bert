package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("Clicking is not fun. Let's avoid it.")
	minInterval, maxInterval, err := getUserClickInputs()

	thirdClick := checkIfThirdClickRequired()

	if err != nil {
		fmt.Printf("An error occured: %v", err.Error())
		os.Exit(1)
	}

	if minInterval < 0 || maxInterval < 0 {
		fmt.Printf("Can't do negative times, exiting.")
		os.Exit(0)
	}

	fmt.Printf("Entered min: %v, Maximum: %v\n", minInterval, maxInterval)
	minDuration, maxDuration := parseDurations(minInterval, maxInterval)
	for {
		fmt.Printf("Sleeping between %v and %v\n", minDuration, maxDuration)
		r := minInterval + rand.Float64()*(maxInterval-minInterval)
		sleep := time.Duration(r) * time.Second
		fmt.Printf("Feeling sleepy... Sleeping: %v second(s)\n", r)
		time.Sleep(sleep)
		x, y := robotgo.GetMousePos()
		fmt.Printf("Current mouse position: x:%v y:%v\n", x, y)
		robotgo.MouseClick("left", true)
		if thirdClick {
			r := 0.5 + rand.Float64()*0.5
			sleep := time.Duration(r) * time.Second
			time.Sleep(sleep)
			fmt.Printf("Ended up sleeping for the third clicker: %v seconds", r)
			robotgo.MouseClick("left", true)
		}
	}
}

func getUserClickInputs() (float64, float64, error) {
	var s []float64
	var err error
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter minimum interval: ")
	min, _ := reader.ReadString('\n')
	if x, err := strconv.ParseFloat(min[0:len(min)-1], 64); err == nil {
		s = append(s, x)
		fmt.Print("Enter maximum interval: ")
		max, _ := reader.ReadString('\n')
		if y, err := strconv.ParseFloat(max[0:len(max)-1], 64); err == nil {
			s = append(s, y)
			sort.Float64s(s)
			return s[0], s[1], err
		}
	}
	return 0, 0, err
}

func parseDurations(min float64, max float64) (time.Duration, time.Duration) {
	minDuration := time.Duration(min) * time.Second
	maxDuration := time.Duration(max) * time.Second
	return minDuration, maxDuration
}

func checkIfThirdClickRequired() bool {
	for {
		fmt.Printf("Do you need the third click? [Y/N]")
		reader := bufio.NewReader(os.Stdin)
		required, _ := reader.ReadString('\n')
		input := required[0 : len(required)-1]
		if input == "Y" || input == "y" {
			return true
		} else if input == "N" || input == "n" {
			return false
		}
	}
}
