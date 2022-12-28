package part2

import (
	"fmt"
	"log"
)

func (board *Board) createMetaboard(data []string) {
	board.Meta = make([][]int, len(data)/50)
	id := 0
	for i := range board.Meta {
		board.Meta[i] = make([]int, len(data[i])/50)
		for j := range board.Meta[i] {
			if data[50*i][50*j] != ' ' {
				board.Meta[i][j] = id
				for x := 0; x < 50; x++ {
					board.Face[id].value[x] = data[50*i+x][50*j : 50*(j+1)]
				}
				board.Face[id].Id = id
				id++
			} else {
				board.Meta[i][j] = -1
			}
		}
	}
}

var locationNeighbors = map[int][4]int{
	1:  {2, -3, -2, 3},
	-1: {-2, -3, 2, 3},
	2:  {-1, -3, 1, 3},
	-2: {1, -3, -1, 3},
	3:  {2, 1, -2, -1},
	-3: {2, -1, -2, 1},
}

/*
 *	 1: front
 *	-1: back
 *	 2: right
 *	-2: left
 *	 3: top
 *	-3: bottom
 */
func (board *Board) findOrientation() {
	// Verify metaboard
	if len(board.Meta) == 0 {
		log.Fatal("Must create metaboard before finding edges")
	}
	for f := range board.Face {
		board.Face[f].Orientation = -1
	}
	// 1st pass: Find first cube with multiple neighbors, assign as front
	assigned := make([]int, 0, 6)
	for f := range board.Face {
		// count the edges
		edges := make(map[int]*Face)
		for d := range directions {
			if board.Face[f].Edge[d].Newface != nil {
				// found a neighbor, count the edge
				edges[d] = board.Face[f].Edge[d].Newface
			}
		}
		if len(edges) < 2 {
			// If singly connected, don't use as front.
			continue
		} else {
			// Assign as front with no rotated orientation
			board.Face[f].Location = 1
			board.Face[f].Orientation = 0
			assigned = append(assigned, 1)
			// Assign the edges accordingly
			for dir, face := range edges {
				iFace := 0
				switch dir {
				case 0:
					iFace = 2
				case 1:
					iFace = -3
				case 2:
					iFace = -2
				case 3:
					iFace = 3
				}
				// Assign as unrotated because directly attached
				face.Location = iFace
				face.Orientation = 0
				fmt.Printf("Face %d neighbor %d is %d so orient %d\n", f, dir, face.Id, iFace)
				assigned = append(assigned, iFace)
			}
			break
		}
	}

	// 2nd pass: induction until all orientations found
	fmt.Println(assigned)
	for len(assigned) < 6 {
		for f := range board.Face {
			face := &board.Face[f]
			if face.Location == 0 {
				// Not assigned, cannot use to induct
				continue
			}
			neighborOrientation := locationNeighbors[face.Location]
			if face.Orientation < 0 {
				// Induce orientation
				face.Orientation = 0
				for i := range directions {
					orienter := face.Edge[i].Newface
					if orienter == nil {
						// not a face
						continue
					} else if orienter.Location == 0 {
						// unknown orientation, can't use as orienter
						continue
					} else {
						// Have an orienter, use to determine my orientation
						for orienter.Location != neighborOrientation[toD(i+face.Orientation)] {
							face.Orientation++
						}
						break
					}
				}
			}
			fmt.Printf("%d:%d searching for %v\n", face.Id, face.Location, neighborOrientation)
			for d := range directions {
				other := face.Edge[d].Newface
				if other == nil {
					// no face in this direction, nothing to do
					continue
				} else if other.Location != 0 {
					// face is already assigned, skip
					continue
				} else {
					// assign neighbor based on location & orientation
					fmt.Printf("Face %d neighbor %d is %d so induct %d\n", f, d, other.Id, neighborOrientation[toD(d+face.Orientation)])
					other.Location = neighborOrientation[toD(d+face.Orientation)]
					assigned = append(assigned, other.Location)
					fmt.Println(assigned)
				}
			}
		}
	}
}

func in[T comparable](l []T, x T) bool {
	for i := range l {
		if l[i] == x {
			return true
		}
	}
	return false
}

func (board *Board) findNeighbors() {
	// 1st pass: Find immediate neighbors, which must be connected directly
	for i := range board.Meta {
		for j := range board.Meta[i] {
			if board.Meta[i][j] == -1 {
				// no face at x,y
				continue
			}
			face := board.Meta[i][j]
			for d, direction := range directions {
				x, y := i+direction[0], j+direction[1]
				if x < 0 || x >= len(board.Meta) {
					// row out of bounds
					continue
				} else if y < 0 || y >= len(board.Meta[i]) {
					// column out of bounds
					continue
				} else if board.Meta[x][y] == -1 {
					// no face at x,y
					continue
				} else {
					// found a neighbor, save the edge
					neighbor := board.Meta[x][y]
					board.Face[face].Edge[d].Newface = &(board.Face[neighbor])
					board.Face[face].Edge[d].Newedge = opposeD(d)
					fmt.Printf("Face %d direct neighbor %d is %d\n", face, d, neighbor)
				}
			}
		}
	}

	// 2nd pass: Find corner neighbors
	for face := range board.Face {
		for iNeighbor := range board.Face[face].Edge {
			neighbor := board.Face[face].Edge[iNeighbor].Newface
			if neighbor == nil {
				// No neighbor here, cannot find neighbor of neighbor yet
				continue
			}
			// get forward neighbor of neighbor
			iNext := toD(iNeighbor + 1)
			next := board.Face[face].Edge[iNext].Newface
			// If forward neighbor exits
			if next != nil {
				// If neighbor is already assigned
				if neighbor.Edge[iNext].Newface != nil || next.Edge[iNeighbor].Newface != nil {
					// Don't overwrite, will overfold
					continue
				}
				// link forwards
				neighbor.Edge[iNext].Newface = next
				neighbor.Edge[iNext].Newedge = iNeighbor
				fmt.Printf("- Face %d corner neighbor %d is %d\n", neighbor.Id, iNext, next.Id)
				// link backwards
				next.Edge[iNeighbor].Newface = neighbor
				next.Edge[iNeighbor].Newedge = iNext
				fmt.Printf("  Face %d corner neighbor %d is %d\n", next.Id, iNeighbor, neighbor.Id)
			}
		}
	}

}

func (board *Board) FindAllEdges() {
	// Verify metaboard
	if len(board.Meta) == 0 {
		log.Fatal("Must create metaboard before finding edges")
	}

	board.findNeighbors()
	board.findOrientation()

	// Zip together by orientation

	/*
		// Track the total number of neighbors found for loop condition
		found := 0
		// 2nd pass: Find 2nd degree neighbors, align by folding:
		for face := range board.face {
			for iNeighbor := range board.face[face].edge {
				neighbor := board.face[face].edge[iNeighbor].newface
				if neighbor == nil {
					// No neighbor here, cannot find neighbor of neighbor yet
					continue
				}
				// get forward neighbor of neighbor
				iNext := toD(iNeighbor + 1)
				next := board.face[face].edge[iNext].newface
				// If forward neighbor is assigned
				if next != nil {
					// link forwards
					neighbor.edge[iNext].newface = next
					neighbor.edge[iNext].newz = opposeD(iNeighbor)
					found++
					// link backwards
					next.edge[iNeighbor].newface = neighbor
					next.edge[iNeighbor].newz = opposeD(iNext)
					found++
				}
			}
		}

		// 3rd pass: horizontal induction
		for face := range board.face {

			// In plane
			// Out of plane
		}

		// 4th pass: vertical induction
	*/
}

func toD(i int) (d int) {
	d = i % 4
	if d < 0 {
		d = 4 - d
	}
	return
}

func opposeD(i int) (d int) {
	return toD(i + 2)
}
