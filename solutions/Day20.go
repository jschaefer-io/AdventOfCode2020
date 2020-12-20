package solutions

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Day20 struct {
	tiles      map[int]*tile
	dimensions int
}

type tile struct {
	image       [][]bool
	signatures  [][4]int
	connections [4]*connection
}

type connection struct {
	tile      int
	variation int
}

func newTile(image [][]bool) *tile {
	dimension := len(image)
	baseSignature := [4]int{}
	revSignature := [4]int{}
	offset := dimension - 1
	for x := 0; x < dimension; x++ {
		if image[0][x] {
			baseSignature[0] += 1 << (offset - x)
			revSignature[0] += 1 << x
		}
		if image[x][offset] {
			baseSignature[1] += 1 << (offset - x)
			revSignature[1] += 1 << x
		}
		if image[offset][x] {
			baseSignature[2] += 1 << (offset - x)
			revSignature[2] += 1 << x
		}
		if image[x][0] {
			baseSignature[3] += 1 << (offset - x)
			revSignature[3] += 1 << x
		}
	}

	variations := make([][4]int, 0)
	for rotation := 0; rotation < 4; rotation++ {
		baseVariation := [4]int{}
		revVariation := [4]int{}
		for i := 0; i < 4; i++ {
			index := (rotation + i) % 4
			baseVariation[i] = baseSignature[index]
			revVariation[i] = revSignature[index]
		}
		variations = append(variations, baseVariation)
		variations = append(variations, revVariation)
	}

	return &tile{
		image:      image,
		signatures: variations,
	}
}

func (d *Day20) init(s string) error {
	d.tiles = make(map[int]*tile)
	splitG := regexp.MustCompile("((|\\r)\\n){2}")
	tileNum := regexp.MustCompile("(\\d+)")
	for _, tileString := range splitG.Split(strings.TrimSpace(s), -1) {
		parts := strings.Split(tileString, "\n")
		num, err := strconv.Atoi(tileNum.FindStringSubmatch(parts[0])[0])
		if err != nil {
			return err
		}
		image := make([][]bool, 0)
		for y := 1; y < len(parts); y++ {
			row := make([]bool, 0)
			for _, x := range strings.Split(strings.TrimSpace(parts[y]), "") {
				row = append(row, x == "#")
			}
			image = append(image, row)
		}
		d.tiles[num] = newTile(image)
	}
	d.dimensions = int(math.Sqrt(float64(len(d.tiles))))
	return nil
}

func (d *Day20) findAdjacentTiles(tileId int, variationIndex int) []*connection {
	currentTile := d.tiles[tileId]
	v := currentTile.signatures[variationIndex]

	tiles := make([]*connection, 0)
	for directionIndex := 0; directionIndex < 4; directionIndex++ {
		for id, tile := range d.tiles {
			if id == tileId || currentTile.connections[directionIndex] != nil {
				continue
			}
			for variant, signature := range tile.signatures {
				if signature[(directionIndex+2)%4] == v[directionIndex] {
					conn := connection{
						tile:      id,
						variation: variant,
					}

					tiles = append(tiles, &conn)
					currentTile.connections[directionIndex] = &conn
				}
			}
		}
	}
	return tiles
}

func (d *Day20) sortTiles() {
	var id int
	for id = range d.tiles {
		break
	}
	queue := make([]connection, 0)
	queue = append(queue, connection{
		tile:      id,
		variation: 0,
	})

	done := make(map[int]struct{})
	for i := 0; i < len(queue); i++ {
		if _, ok := done[queue[i].tile]; ok {
			continue
		}
		connections := d.findAdjacentTiles(queue[i].tile, queue[i].variation)
		for _, connection := range connections {
			queue = append(queue, *connection)
			done[queue[i].tile] = struct{}{}
		}
	}
}

func (d *Day20) executeA() int {
	res := 1
	for id, tile := range d.tiles {
		c := 0
		for _, connection := range tile.connections {
			if connection != nil {
				c++
			}
		}
		if c == 2 {
			res *= id
		}
	}
	return res
}

func (d *Day20) executeB() int {
	//fmt.Println("test")
	return 1
}

func (d *Day20) Handle(s string) ([]string, error) {
	err := d.init(s)
	if err != nil {
		return nil, err
	}
	d.sortTiles()

	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	results = append(results, fmt.Sprintf("%d", d.executeB()))
	return results, nil
}
