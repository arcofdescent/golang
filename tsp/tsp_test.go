package tsp

import "testing"
import "fmt"
import "reflect"
import "time"

func TestCalcDistanceBetweenPoints(t *testing.T) {
	fmt.Println("TestCalcDistanceBetweenPoints")

	points := []Point{
		{Id: "P1", X: 164.51666259765625, Y: 304},
		{Id: "P2", X: 206.51666259765625, Y: 156},
		{Id: "P3", X: 334.51666259765625, Y: 180},
	}

	expected := []Distance{
		{Id1: "P1", Id2: "P2", Distance: 153.8440769090575},
		{Id1: "P1", Id2: "P3", Distance: 210.41863035387337},
		{Id1: "P2", Id2: "P3", Distance: 130.23056476879765},
	}

	got := calcDistanceBetweenPoints(points)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("calcDistanceBetweenPoints()")
	}
}

func TestCalcShortestRoute(t *testing.T) {
	fmt.Println("TestCalcShortestRoute")
	startTime := time.Now()

	points := []Point{
		{Id: "P1", X: 164.51666259765625, Y: 304},
		{Id: "P2", X: 206.51666259765625, Y: 156},
		{Id: "P3", X: 334.51666259765625, Y: 180},
	}

	distances := calcDistanceBetweenPoints(points)

	res := calcShortestRoute(len(points), distances)
	duration := time.Since(startTime)
	res.Duration = fmt.Sprintf("%v", duration)
	fmt.Printf("res: %#v\n", res)
	fmt.Printf("duration: %v\n", duration)
}
