package validcoords

import (
	"regexp"
	"strconv"
	"strings"
)

func IsValidCoordinates(coordinates string) bool {

	coords := strings.Split(strings.Trim(coordinates, " "), ",")

	if len(coords) > 2 {
		return false
	}

	coords[0] = strings.Trim(coords[0], " ")
	coords[1] = strings.Trim(coords[1], " ")

	matched1, _ := regexp.Match(`(^-?\d+\.?\d*$)`, []byte(coords[0]))
	matched2, _ := regexp.Match(`(^-?\d+\.?\d*$)`, []byte(coords[1]))

	if !matched1 || !matched2 {
		return false
	}
	// checks for space between negative sign and the number
	if strings.ContainsAny(coords[0], " ") || strings.ContainsAny(coords[1], " ") {
		return false
	}

	latitude, err := strconv.ParseFloat(coords[0], 64)
	if err != nil {
		return false

	}

	longitude, err := strconv.ParseFloat(coords[1], 64)
	if err != nil {
		return false
	}

	//check to make sure latitude and longitude are within the integer range
	if latitude < -90 || latitude > 90 || longitude < -180 || longitude > 180 {
		return false
	}

	return true
}
