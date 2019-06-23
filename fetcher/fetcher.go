package fetcher

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/mlbright/forecast/v2"
)

type Query struct {
	Latitude  string
	Longitude string
}

type Result struct {
	Answer      string
	Description string
}

func GetWeather(q Query) *forecast.Forecast {
	err := godotenv.Load()

	key := os.Getenv("FORECAST_API_KEY")

	f, err := forecast.Get(key, q.Latitude, q.Longitude, "now", forecast.UK, forecast.English)
	if err != nil {
		log.Println(err)
	}

	return f
}

func SetResult(f *forecast.Forecast) Result {
	lines := []string{}
	var description string
	trigger := 16.0
	goodConditions := []string{
		"partly-cloudy-day", // YES
		"clear-day",         // YES
	}

	temp := f.Currently.ApparentTemperature

	if temp >= trigger {
		lines = []string{
			"Hell yeah",
			"Of course",
			"Get the legs out",
			"Totes",
			"Flat out",
			"No Doubt",
			"It bloody well is",
		}

		t := strconv.Itoa(int(temp))

		description = "It's " + forecastIconToWord(f.Currently.Icon) + " " + t + " Degrees"
	}

	warmer_hours := []forecast.DataPoint{}

	for _, h := range f.Hourly.Data {
		if h.ApparentTemperature > trigger && contains(goodConditions, h.Icon) {
			warmer_hours = append(warmer_hours, h)
		}
	}

	return Result{
		Answer:      random(lines),
		Description: description,
	}
}

func forecastIconToWord(icon string) string {
	words := map[string][]string{
		"clear-day": []string{
			"nice",
			"sunny",
			"roasting",
			"swealtering",
			"schwealterin'",
			"dynamite",
			"unreal",
			"amazing",
			"lovely",
			"hot",
			"warm",
			"great",
		},
		"clear-night": []string{
			"starry",
			"unreal",
			"lovely",
		},
		"rain": []string{
			"wet",
			"soggy",
			"damp",
			"moist",
			"cold",
			"coul",
			"baltic",
			"freezing",
			"chilly",
		},
		"snow": []string{
			"snowy",
			"artic",
			"baltic",
			"freezing",
			"freezin'",
			"chilly",
		},
		"sleet": []string{
			"wet",
			"wet and snowy",
			"freezin'",
			"freezing",
			"damp",
			"moist",
			"cold",
			"coul",
			"baltic",
			"chilly",
		},
		"wind": []string{
			"windy",
			"chilly",
			"cold",
			"coul",
			"baltic",
			"blustery",
		},
		"fog": []string{
			"foggy",
		},
		"cloudy": []string{
			"grey",
			"shite",
			"cloudy",
			"ballix",
		},
		"partly-cloudy-day": []string{
			"cloudy",
			"not too bad",
		},
		"partly-cloudy-night": []string{
			"starry",
		},
	}

	word := words[icon][rand.Intn(len(words[icon]))]
	var prefix_word []string

	vowels := []string{"a", "e", "i", "o", "u"}
	if contains(vowels, string(word[0])) {
		prefix_word = []string{"an", word}
	} else {
		prefix_word = []string{"a", word}
	}

	return strings.Join(prefix_word, " ")
}

func contains(arr []string, inc string) bool {
	for _, a := range arr {
		if a == inc {
			return true
		}
	}

	return false
}

func random(arr []string) string {
	rand.Seed(time.Now().Unix())
	return arr[rand.Intn(len(arr))]
}
