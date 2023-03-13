package main

import (
	"fmt"
	"math/rand"
)

type day struct {
	date        int
	lightPrecip bool
	temperature string
}

type month struct {
	number                  int
	name, earthName, season string
	days                    []day
}

func generateWeather(mo int) (precipitation bool, temperature string) {

	precipRoll := rand.Intn(20) + 1
	var temperatureRoll int
	if mo == 11 || mo < 2 || mo > 4 && mo < 8 {
		temperatureRoll = rand.Intn(20) + 1
	}
	switch mo {
	case 11, 0, 1:
		if precipRoll >= 8 {
			precipitation = true
		}
		//If it's a winter month, it can be cold. January/Abadius is cold more often.
		if temperatureRoll >= 18 || (temperatureRoll >= 16 && mo == 0) {
			temperature = "Mild Cold"
		} else {
			temperature = "Normal"
		}
	case 2, 3, 4, 8, 9, 10:
		if precipRoll >= 15 {
			precipitation = true
		}
		temperature = "Normal"
	case 5, 6, 7:
		if precipRoll == 20 {
			precipitation = true
		}
		//If it's a summer month, it is rarely hot. July/Erastus is warm slightly more often
		if temperatureRoll >= 20 || (temperatureRoll >= 19 && mo == 6) {
			temperature = "Mild Heat"
		} else {
			temperature = "Normal"
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
			year[i].days[j].lightPrecip, year[i].days[j].temperature = generateWeather(i)
		}
	}
	for i := 0; i < len(year); i++ {
		fmt.Println(year[i].number, year[i].name, year[i].earthName, year[i].season)
		for j := 0; j < len(year[i].days); j++ {
			fmt.Println("\t", year[i].days[j])
		}
	}
}
