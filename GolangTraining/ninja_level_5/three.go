package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "White",
		},
		fourWheel: true,
	}
	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Black",
		},
		luxury: true,
	}
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.fourWheel)
	fmt.Println(s.color)
}
