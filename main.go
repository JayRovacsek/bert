package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
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
		fmt.Println(fmt.Sprintf("Clicked at: %v, %v", x, y))
		if thirdClick {
			fmt.Println(fmt.Sprintf("Third click: %v", thirdClick))
			r := 0.5 + rand.Float64()*0.5
			sleep := time.Duration(r) * time.Second
			time.Sleep(sleep)
			fmt.Println(fmt.Sprintf("Ended up sleeping for the third clicker: %v seconds", r))
			robotgo.MouseClick("left", true)
			fmt.Println(fmt.Sprintf("Clicked at: %v, %v", x, y))
		}
	}
}

func getUserClickInputs() (float64, float64, error) {
	var s []float64
	var err error
	var temp string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter min value: ")
	temp, _ = reader.ReadString('\n')
	min, _ := strconv.ParseFloat(strings.TrimSpace(temp), 64)

	fmt.Print("Enter max value: ")
	temp, _ = reader.ReadString('\n')
	max, _ := strconv.ParseFloat(strings.TrimSpace(temp), 64)

	s = append(s, min, max)

	sort.Float64s(s)

	return s[0], s[1], err
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
		input := strings.TrimSpace(required)

		if input == "Y" || input == "y" {
			return true
		}

		if input == "N" || input == "n" {
			return false
		}
	}
}
