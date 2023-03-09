package main

import (
	"fmt"
	"math/rand"
)

type day struct {
	date        int
	lightPrecip bool
}

type month struct {
	days                    []day
	name, earthName, season string
	number                  int
}

func generateWeather(mo int) (precipitation bool) {

	precipRoll := rand.Intn(20) + 1
	fmt.Println("Precipitation roll is:", precipRoll)
	switch mo {
	case 11, 0, 1:
		if precipRoll >= 8 {
			precipitation = true
		}
	case 2, 3, 4, 8, 9, 10:
		if precipRoll >= 15 {
			precipitation = true
		}
	case 5, 6, 7:
		if precipRoll == 20 {
			precipitation = true
		}
	}
	return
}

func main() {
	year := make([]month, 12)
	monthNames := [12]string{"Abadius", "Calistril", "Pharast", "Gohzran", "Desnus", "Sarenith",
		"Erastus", "Arodus", "Rova", "Lamashan", "Neth", "Kuthona"}
	monthEarthNames := [12]string{"January", "February", "March", "April", "May", "June", "July",
		"August", "September", "October", "November", "December"}
	daysInMonth := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	for i := 0; i < len(year); i++ {
		year[i].name = monthNames[i]
		year[i].earthName = monthEarthNames[i]
		switch i {
		case 11, 0, 1:
			year[i].season = "Winter"
		case 2, 3, 4:
			year[i].season = "Spring"
		case 5, 6, 7:
			year[i].season = "Summer"
		case 8, 9, 10:
			year[i].season = "Autumn"
		}
		year[i].number = i + 1
		year[i].days = make([]day, daysInMonth[i])
		for j := 0; j < daysInMonth[i]; j++ {
			year[i].days[j].date = j + 1
			year[i].days[j].lightPrecip = generateWeather(i)
		}
	}
	for i := 0; i < len(year); i++ {
		fmt.Println(year[i])
	}
}
