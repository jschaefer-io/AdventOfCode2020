package solutions

import (
	"errors"
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

func (d *Day20) executeA() (int, []int) {
	corners := make([]int, 0)
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
			corners = append(corners, id)
		}
	}
	return res, corners
}

func (d *Day20) resolveVariant(id int) (int, error) {
	for _, tile := range d.tiles {
		for _, conn := range tile.connections {
			if conn != nil && conn.tile == id {
				return conn.variation, nil
			}
		}
	}
	return 0, errors.New("could not resolve variant")
}

func (d *Day20) getNextTile(con connection, a, b int) []connection {
	visited := make(map[int]struct{})
	list := []connection{
		con,
	}
	cId := con.tile
	for x := 0; x < d.dimensions-1; x++ {
		visited[cId] = struct{}{}
		cTile := d.tiles[cId]

		next := cTile.connections[a]
		if next == nil {
			next = cTile.connections[b]
		}
		if _, ok := visited[next.tile]; ok {
			next = cTile.connections[b]
		}
		if next == nil {
			break
		}
		cId = next.tile
		list = append(list, *next)
	}
	return list
}

func (d *Day20) executeB(corners []int) int {
	variant, err := d.resolveVariant(corners[0])
	if err != nil {
		return -1
	}
	initConn := connection{
		tile:      corners[0],
		variation: variant,
	}
	rows := d.getNextTile(initConn, 2, 0)
	for _, rowConn := range rows {
		cols := d.getNextTile(rowConn, 1, 3)
		fmt.Println(cols)
	}

	//fmt.Println(corner)
	return 1
}

func (d *Day20) Handle(s string) ([]string, error) {
	err := d.init(s)
	if err != nil {
		return nil, err
	}
	d.sortTiles()

	results := make([]string, 0)
	a, corners := d.executeA()
	results = append(results, fmt.Sprintf("%d", a))
	results = append(results, fmt.Sprintf("%d", d.executeB(corners)))
	return results, nil
}
