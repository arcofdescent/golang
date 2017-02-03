// Travelling salesman problem
// ---------------------------
// Given a list of cities and the distances between each pair of cities, what is
// the shortest possible route that visits each city exactly once and returns to
// the origin city?
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	//"sort"
	//"strconv"
	//"strings"
)

type Distance struct {
	City1    string
	City2    string
	Distance float64
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/calcRoute", calcRoute)
	http.Handle("/", r)

	srv := &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: r,
	}

	srv.ListenAndServe()
}

func calcRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got request")
	d := Distance{City1: "Mumbai", City2: "Bangalore", Distance: 45.67}
	json, err := json.Marshal(d)
	if err != nil {
		fmt.Printf("Error while marshalling: %s\n", err)
	}

	s := string(json[:])

	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.WriteString(w, s)
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}

	helper(arr, len(arr))
	return res
}
