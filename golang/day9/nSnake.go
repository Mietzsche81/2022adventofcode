type Snake struct {
	length int
	link   []Link
}

type Link struct {
	dimension int
	head      []int
	tail      []int
}

func NewSnake(length int, dimension int) Snake {
	NewLink := func(dimension int) Link {
		state := make([]int, dimension)
		return Link{
			dimension: dimension,
			head:      state,
			tail:      state,
		}
	}
	link := make([]Link, length-1)
	for i := 0; i < length-1; i++ {
		link[i] = NewLink(dimension)
	}
	return Snake{
		length: length,
		link:   link,
	}
}
