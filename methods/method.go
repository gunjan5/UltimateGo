//Declare a struct that represents a baseball player. Include name, atBats and hits.
//Declare a method that calculates a players batting average. The formula is Hits / AtBats.
//Declare a slice of this type and initialize the slice with several players.
//Iterate over the slice displaying the players name and batting average.

package main

import (
	"fmt"
)

type player struct {
	name         string
	atBats, hits int
}

func (p *player) average() float64 {

	return float64(p.hits) / float64(p.atBats)

}

func main() {
	players := []player{
		{"bill", 10, 6},
		{"billy", 12, 7},
		{"bilster", 9, 4},
	}

	for _, v := range players {
		fmt.Println(v.name, v.average())
	}

}
