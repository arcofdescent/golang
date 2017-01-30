// Travelling salesman problem
// ---------------------------
// Given a list of cities and the distances between each pair of cities, what is
// the shortest possible route that visits each city exactly once and returns to
// the origin city?
package main

import (
	"fmt"
	"sort"
	"strings"
)

var cities = [...]string{"Mumbai", "Delhi", "Bangalore"}

// set the distances
var distance = map[string]float32{
	"del_mum": 56,
	"ban_mum": 26,
	"ban_del": 96,
}

func main() {
	fmt.Println("start")

	d := getDistance("Delhi", "Bangalore")
	fmt.Println(d)
}

func getDistance(city1 string, city2 string) float32 {

	// get first 3 letters
	city1 = strings.ToLower(city1[:3])
	city2 = strings.ToLower(city2[:3])

	// create a slice for sorting
	cities := []string{city1, city2}
	sort.Strings(cities)

	// key into the distance map
	dist_key := strings.Join(cities, "_")
	return distance[dist_key]
}
