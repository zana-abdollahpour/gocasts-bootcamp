package main

import "fmt"

type Clock struct {
	hours   int
	minutes int
}

func NewClock(h, m int) Clock {
	minutes := m % 60

	if minutes < 0 {
		minutes += 60
		h--
	}

	hours := (h + (m / 60)) % 24

	if hours < 0 {
		hours += 24
	}

	return Clock{
		hours,
		minutes,
	}
}

func (c Clock) Add(m int) Clock {
	return NewClock(c.hours, (c.minutes + m))
}

func (c Clock) Subtract(m int) Clock {
	return NewClock(c.hours, (c.minutes - m))
}

func (c Clock) String() string {
	return fmt.Sprintf("%.2d:%.2d", c.hours, c.minutes)
}

func main() {
	fmt.Println(NewClock(4, 44))
}
