package main

import "fmt"

type Circle struct {
	On    bool
	Ammo  int
	Power int
}

func (c *Circle) Shoot() bool {
	if c.Ammo > 0 && c.On != false {
		c.Ammo -= 1
		return true
	} else {
		return false
	}
}
func (c *Circle) RideBike() bool {
	if c.Power > 0 && c.On != false {
		c.Power -= 1
		return true
	} else {
		return false
	}
}
func main() {

	testStruct := new(Circle)
	fmt.Print(testStruct)
}
