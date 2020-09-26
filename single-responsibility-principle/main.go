package main

import (
	"log"
	"encoding/json"
	"fmt"
	"math"
)

type circle struct{ radius float64 }

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) name() string {
	return "circle"
}

type square struct{ length float64 }

func (s square) area() float64 {
	return s.length * s.length
}
func (s square) name() string {
	return "square"
}

type shape interface{
	area() float64
	name() string
}

type outputter struct{}

func(o outputter)Text(s shape) string{
	return fmt.Sprintf("area of the %s: %f",s.name(),s.area())
}
func(o outputter)JSON(s shape) string{
	res:=struct{
		Name string `json:"name"`
		Area float64 `json:"area"`
	}{
		Name: s.name(),
		Area: s.area(),
	}
	bs,err:=json.Marshal(res)
	if err!=nil{
		log.Fatal(err)
	}
	return string(bs)
}

func main() {
	c := circle{radius: 5}
	s := square{length: 6}

	o:=outputter{}
	oText:=o.Text(c)
	oJson:=o.JSON(s)
	fmt.Println(oText)
	fmt.Println(oJson)
}
