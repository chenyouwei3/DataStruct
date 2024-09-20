package main

func main() {
}

type DequqTEST struct {
	data []int
}

func (d *DequqTEST) PushFront(value int) {
	d.data = append([]int{value}, d.data...)
}
