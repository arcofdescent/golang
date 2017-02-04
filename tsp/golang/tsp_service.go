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
	"io/ioutil"
	"math"
	"net/http"
	//"sort"
	//"strconv"
	//"strings"
)

type Distance struct {
	Id1      string
	Id2      string
	Distance float64
}

type Point struct {
	Id string
	X  float64
	Y  float64
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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error while reading body: %s\n", err)
	}
	//fmt.Printf("body: %s\n", body)

	points := make([]Point, 0)
	err = json.Unmarshal(body, &points)
	if err != nil {
		fmt.Printf("Error while unmarshalling: %s\n", err)
	}
	//fmt.Printf("struct: %#v", d)

	distances := calcDistanceBetweenPoints(points)
	fmt.Printf("distances: %#v\n", distances)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.WriteString(w, "OK")
}

func calcDistanceBetweenPoints(points []Point) []Distance {

	distances := make([]Distance, 0)

	comb(len(points), 2, func(c []int) {
		//fmt.Printf("c: %#v\n", c)
		id1 := points[c[0]].Id
		x1 := points[c[0]].X
		y1 := points[c[0]].Y
		id2 := points[c[1]].Id
		x2 := points[c[1]].X
		y2 := points[c[1]].Y

		dst := math.Sqrt(((y2 - y1) * (y2 - y1)) + ((x2 - x1) * (x2 - x1)))
		distance := Distance{Id1: id1, Id2: id2, Distance: dst}
		fmt.Printf("dst: %#v\n", distance)
		distances = append(distances, distance)
	})

	return distances
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

func comb(n, m int, emit func([]int)) {
	s := make([]int, m)
	last := m - 1
	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < n; j++ {
			s[i] = j
			if i == last {
				emit(s)
			} else {
				rc(i+1, j+1)
			}
		}
		return
	}
	rc(0, 0)
}
