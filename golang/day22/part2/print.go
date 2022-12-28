package part2

import "fmt"

func (board *Board) PrintOrientation() {
	for i := range board.Face {
		fmt.Printf("Face %d is %d rotated %d times = [ ", i, board.Face[i].Location, board.Face[i].Orientation)
		for j := range board.Face[i].Edge {
			if board.Face[i].Edge[j].Newface == nil {
				fmt.Printf("-:-- ")
			} else {
				fmt.Printf("%d:%2d ", board.Face[i].Edge[j].Newface.Id, board.Face[i].Edge[j].Newface.Location)
			}
		}
		fmt.Printf("]\n")
	}
}

func (s State) Print() string {
	return fmt.Sprintf("%d:[%d %d %s]",
		s.Face.Id, s.x, s.y, EncodeDirection(s.z),
	)
}
