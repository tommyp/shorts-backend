package fetcher

import (
	"log"
	"math/rand"
	"os"
	"sort"
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

	warmerHours := []forecast.DataPoint{}

	for _, h := range f.Hourly.Data {
		if h.ApparentTemperature > trigger && contains(goodConditions, h.Icon) {
			warmerHours = append(warmerHours, h)
		}
	}

	warmerHours = sortHours(warmerHours)

	t := strconv.Itoa(int(temp))

	if temp >= trigger && contains(goodConditions, f.Currently.Icon) {
		lines = []string{
			"Hell yeah",
			"Of course",
			"Get the legs out",
			"Totes",
			"Flat out",
			"No Doubt",
			"It bloody well is",
		}

		description = "It's " + forecastIconToWord(f.Currently.Icon) + " " + t + " Degrees"
	} else if len(warmerHours) >= 1 {
		// Not warm now but a warmer hour later
		lines = []string{
			"Not now, but it'll be warmer later",
			"Give it a chance",
			"Houl yer horses",
			"Relax yer kacks",
			"Don't worry",
			"Not yet",
		}

		description = "It's a " + forecastIconToWord(f.Currently.Icon) + " " + t + " degrees right now, but it'll be " + forecastIconToWord(f.Currently.Icon) + " " + strconv.Itoa(int(warmerHours[0].ApparentTemperature)) + " degrees later"
	} else {
		lines = []string{
			"No way",
			"Hell no",
			"Are you not wise?",
			"Jeans flat out",
			"Fraid not",
			"Way on",
			"Away on",
			"Fuck away off",
			"Are you having a giraffe?",
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
	var prefixWord []string

	vowels := []string{"a", "e", "i", "o", "u"}
	if contains(vowels, string(word[0])) {
		prefixWord = []string{"an", word}
	} else {
		prefixWord = []string{"a", word}
	}

	return strings.Join(prefixWord, " ")
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

func sortHours(hours []forecast.DataPoint) []forecast.DataPoint {
	sort.SliceStable(hours, func(i, j int) bool {
		return hours[i].ApparentTemperature > hours[j].ApparentTemperature
	})

	return hours
}
