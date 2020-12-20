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
	monsterMap monsterMap
}

type monsterMap struct {
	imageMap [][2]int
	width    int
	height   int
}

type tile struct {
	images [][][]bool
	edges  [][4]int
}

type connection struct {
	to        int
	direction int
	variation int
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

		images := make([][][]bool, 0)
		edges := make([][4]int, 0)
		for i := 0; i < 8; i++ {
			img, edge := d.getImageVariant(image, i)
			images = append(images, img)
			edges = append(edges, edge)
		}

		d.tiles[num] = &tile{
			images: images,
			edges:  edges,
		}
	}
	d.dimensions = int(math.Sqrt(float64(len(d.tiles))))

	monster := "                  #\n#    ##    ##    ### \n #  #  #  #  #  #   "
	d.monsterMap = monsterMap{
		imageMap: make([][2]int, 0),
	}
	for y, mLines := range strings.Split(monster, "\n") {
		for x, p := range strings.Split(mLines, "") {
			if p == "#" {
				d.monsterMap.imageMap = append(d.monsterMap.imageMap, [2]int{x, y})
				d.monsterMap.width = int(math.Max(float64(d.monsterMap.width), float64(x)))
				d.monsterMap.height = int(math.Max(float64(d.monsterMap.height), float64(y)))
			}
		}
	}

	return nil
}

func (d *Day20) getImageVariant(baseImage [][]bool, variant int) ([][]bool, [4]int) {
	tileDimensions := len(baseImage)

	rotation := variant % 4
	flip := variant / 4

	// rotation
	image := make([][]bool, 0)
	for a := 0; a < tileDimensions; a++ {
		row := make([]bool, 0)
		for b := 0; b < tileDimensions; b++ {
			t := baseImage[a][b]
			switch rotation {
			case 1:
				t = baseImage[b][tileDimensions-1-a]
				break
			case 2:
				t = baseImage[tileDimensions-1-a][tileDimensions-1-b]
				break
			case 3:
				t = baseImage[tileDimensions-1-b][a]
				break
			default:
				t = baseImage[a][b]
				break
			}
			row = append(row, t)
		}
		image = append(image, row)
	}

	// flip
	newImage := make([][]bool, tileDimensions)
	for a := 0; a < tileDimensions; a++ {
		row := make([]bool, tileDimensions)
		for b := 0; b < tileDimensions; b++ {
			switch flip {
			case 1:
				row[b] = image[a][tileDimensions-1-b]

				break
			case 2:
				row[b] = image[tileDimensions-1-a][b]
				break
			default:
				row[b] = image[a][b]
				break
			}
		}
		newImage[a] = row
	}
	image = newImage

	return image, d.getEdges(image)
}

func (d *Day20) getEdges(img [][]bool) [4]int {
	edges := [4]int{}
	count := len(img)
	for p := 0; p < count; p++ {
		if img[0][p] {
			edges[0] += 1 << (count - 1 - p)
		}
		if img[count-1][p] {
			edges[2] += 1 << (count - 1 - p)
		}
		if img[p][count-1] {
			edges[1] += 1 << (count - 1 - p)
		}
		if img[p][0] {
			edges[3] += 1 << (count - 1 - p)
		}
	}
	return edges
}

func (d *Day20) printImage(image [][]bool) {
	for y := 0; y < len(image); y++ {
		for x := 0; x < len(image[y]); x++ {
			if image[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func (d *Day20) isAdjacentTile(edges [4]int, checkTile *tile) (int, int, bool) {
	for variation, cmpEdges := range checkTile.edges {
		for r := 0; r < 4; r++ {
			if edges[r] == cmpEdges[(r+2)%4] {
				return r, variation, true
			}
		}
	}
	return -1, -1, false
}

func (d *Day20) findAdjacentTile(searchTile *tile, variation int) []connection {
	connections := make([]connection, 0)
	edges := searchTile.edges[variation]
	for id, cmpTile := range d.tiles {
		if searchTile == cmpTile {
			continue
		}

		dir, variation, adj := d.isAdjacentTile(edges, cmpTile)
		if adj {
			connections = append(connections, connection{
				to:        id,
				variation: variation,
				direction: dir,
			})
		}
	}
	return connections
}

func (d *Day20) executeA() (int, []int) {
	list := make([]int, 0)
	count := 1
	for id, t := range d.tiles {
		if len(d.findAdjacentTile(t, 0)) == 2 {
			count *= id
			list = append(list, id)
		}
	}
	return count, list
}

func (d *Day20) getFullMap(startTile int) [][]connection {
	passed := make(map[int]struct{})
	baseConn := connection{
		to:        startTile,
		variation: 0,
	}
	fullMap := make([][]connection, d.dimensions)
	for y := 0; y < d.dimensions; y++ {
		rowConn := baseConn
		row := make([]connection, d.dimensions)
		row[0] = baseConn
		for x := 1; x <= d.dimensions; x++ {
			passed[baseConn.to] = struct{}{}

			for _, conn := range d.findAdjacentTile(d.tiles[baseConn.to], baseConn.variation) {
				if _, ok := passed[conn.to]; ok {
					continue
				}
				if conn.direction == 1 || conn.direction == 3 {
					baseConn = conn
					row[x] = conn
				}
			}
		}
		for _, conn := range d.findAdjacentTile(d.tiles[rowConn.to], rowConn.variation) {
			if _, ok := passed[conn.to]; ok {
				continue
			}
			if conn.direction == 2 || conn.direction == 0 {
				baseConn = conn
			}
		}
		fullMap[y] = row
	}
	return fullMap
}

func (d *Day20) findTopLeftTile(corners []int) (int, error) {
	for _, corner := range corners {
		res := d.findAdjacentTile(d.tiles[corner], 0)
		if len(res) != 2 {
			continue
		}
		if res[0].direction == 1 && res[1].direction == 2 || res[0].direction == 2 && res[1].direction == 1 {
			return corner, nil
		}
	}
	return -1, errors.New("top left tile could not be identified")
}

func (d *Day20) removeBorder(ogImage [][]bool) [][]bool {
	count := len(ogImage)
	image := make([][]bool, count-2)
	for y := 1; y < count-1; y++ {
		row := make([]bool, count-2)
		for x := 1; x < count-1; x++ {
			row[x-1] = ogImage[y][x]
		}
		image[y-1] = row
	}
	return image

}

func (d *Day20) buildFullImage(fullMap [][]connection) [][]bool {
	fullImage := make([][]bool, 0)
	for _, partialMap := range fullMap {
		count := 0
		rowMap := make(map[int][]bool)
		for _, conn := range partialMap {
			image := d.removeBorder(d.tiles[conn.to].images[conn.variation])
			for i, row := range image {
				list, ok := rowMap[i]
				if !ok {
					list = make([]bool, 0)
					count = i + 1
				}
				rowMap[i] = append(list, row...)
			}
		}
		for i := 0; i < count; i++ {
			fullImage = append(fullImage, rowMap[i])
		}
	}
	return fullImage
}

func (d *Day20) findMonster(image [][]bool) int {
	monsterCount := 0
	height := len(image)
	width := len(image[0])
	for y := 0; y < height-d.monsterMap.height; y++ {
		for x := 0; x < width-d.monsterMap.width; x++ {
			foundOne := true
			for _, parts := range d.monsterMap.imageMap {
				if !image[y+parts[1]][x+parts[0]] {
					foundOne = false
				}
			}
			if foundOne {
				monsterCount++
			}
		}
	}
	return monsterCount
}

func (d *Day20) countRoughWater(image [][]bool) int {
	count := 0
	for _, row := range image {
		for _, p := range row {
			if p {
				count++
			}
		}
	}
	return count
}

func (d *Day20) executeB(corners []int) int {
	startTile, _ := d.findTopLeftTile(corners)
	fullMap := d.getFullMap(startTile)
	fullImage := d.buildFullImage(fullMap)
	for i := 0; i < 8; i++ {
		img, _ := d.getImageVariant(fullImage, i)
		mCount := d.findMonster(img)
		if mCount > 0 {
			return d.countRoughWater(img) - len(d.monsterMap.imageMap) * mCount
		}
	}
	return -1
}

func (d *Day20) Handle(s string) ([]string, error) {
	err := d.init(s)
	if err != nil {
		return nil, err
	}
	results := make([]string, 0)
	a, corners := d.executeA()
	results = append(results, fmt.Sprintf("%d", a))
	results = append(results, fmt.Sprintf("%d", d.executeB(corners)))
	return results, nil
}
