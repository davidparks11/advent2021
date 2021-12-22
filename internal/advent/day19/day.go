package day19

import 	"github.com/davidparks11/advent2021/internal/math"

type Scanner struct {
	Beacons map[Point]struct{}
}

func (s *Scanner) FindOverlappingBeacons(other *Scanner) (*Scanner, *vector) {
	//loop through all possibly rotataions to find an offset for other scanner
	for _, rotatedBeacons := range other.rotateBeacons() {
		if offset := s.findOffset(rotatedBeacons); offset != nil {
			rotatedBeacons.translate(*offset)    //translate beacon locations to use s as origin
			rotatedBeacons.addBeacons(s.Beacons) //add original scanner's beacons to other's beacons
			return rotatedBeacons, offset
		}
	}

	return nil, nil
}

func (s *Scanner) findOffset(other *Scanner) *vector {
	distances := make(map[vector]int)
	//for each beacon pair, take distance and increment in map
	for sBeacon := range s.Beacons {
		for otherBeacon := range other.Beacons {
			distance := sBeacon.distance(otherBeacon)
			distances[distance]++
		}
	}

	//find where 12 or more beacons share a distance offset
	var offset *vector
	for v, numBeacons := range distances {
		if numBeacons >= 12 {
			offset = &v
			break
		}
	}

	return offset
}

func (s *Scanner) translate(v vector) {
	translated := make(map[Point]struct{})
	for p := range s.Beacons {
		translated[p.translate(v)] = struct{}{}
	}
	s.Beacons = translated
}

func (s *Scanner) addBeacons(m map[Point]struct{}) {
	for p := range m {
		s.Beacons[p] = struct{}{}
	}
}

func (s *Scanner) rotateBeacons() []*Scanner {
	var beaconRotations []*Scanner
	for i, r := range rotations {
		beaconRotations = append(beaconRotations, &Scanner{
			Beacons: make(map[Point]struct{}),
		})
		for p := range s.Beacons {
			p = p.rotate(r)
			beaconRotations[i].Beacons[p] = struct{}{}
		}
	}
	return beaconRotations
}

type Point struct {
	X, Y, Z int
}

func (p Point) dotProduct(other Point) int {
	return p.X*other.X + p.Y*other.Y + p.Z*other.Z
}

func (p Point) distance(other Point) vector {
	return vector{X: p.X - other.X, Y: p.Y - other.Y, Z: p.Z - other.Z}
}

func (p Point) translate(v vector) Point {
	return Point{X: p.X + v.X, Y: p.Y + v.Y, Z: p.Z + v.Z}
}

func (p Point) rotate(m matrix) Point {
	return Point{X: p.dotProduct(m.X), Y: p.dotProduct(m.Y), Z: p.dotProduct(m.Z)}
}

func (p Point) Manhattan(other Point) int {
	return math.Abs(p.X-other.X) + math.Abs(p.Y-other.Y) + math.Abs(p.Z-other.Z)
}

type vector Point

type matrix struct {
	X Point
	Y Point
	Z Point
}

//array of all 24 rotations lazily copied from the internet and manipulated
var rotations = []matrix{
	{Point{0, 0, 1}, Point{0, -1, 0}, Point{1, 0, 0}},
	{Point{-1, 0, 0}, Point{0, 0, -1}, Point{0, -1, 0}},
	{Point{0, 0, 1}, Point{1, 0, 0}, Point{0, 1, 0}},
	{Point{0, 1, 0}, Point{1, 0, 0}, Point{0, 0, -1}},
	{Point{0, -1, 0}, Point{0, 0, 1}, Point{-1, 0, 0}},
	{Point{0, 0, -1}, Point{1, 0, 0}, Point{0, -1, 0}},
	{Point{0, 0, 1}, Point{-1, 0, 0}, Point{0, -1, 0}},
	{Point{0, 0, -1}, Point{-1, 0, 0}, Point{0, 1, 0}},
	{Point{-1, 0, 0}, Point{0, -1, 0}, Point{0, 0, 1}},
	{Point{1, 0, 0}, Point{0, 0, -1}, Point{0, 1, 0}},
	{Point{0, 1, 0}, Point{0, 0, 1}, Point{1, 0, 0}},
	{Point{0, 1, 0}, Point{0, 0, -1}, Point{-1, 0, 0}},
	{Point{0, 0, -1}, Point{0, -1, 0}, Point{-1, 0, 0}},
	{Point{0, -1, 0}, Point{0, 0, -1}, Point{1, 0, 0}},
	{Point{0, 0, -1}, Point{0, 1, 0}, Point{1, 0, 0}},
	{Point{0, 1, 0}, Point{-1, 0, 0}, Point{0, 0, 1}},
	{Point{0, -1, 0}, Point{1, 0, 0}, Point{0, 0, 1}},
	{Point{0, -1, 0}, Point{-1, 0, 0}, Point{0, 0, -1}},
	{Point{1, 0, 0}, Point{0, 0, 1}, Point{0, -1, 0}},
	{Point{-1, 0, 0}, Point{0, 0, 1}, Point{0, 1, 0}},
	{Point{1, 0, 0}, Point{0, 1, 0}, Point{0, 0, 1}},
	{Point{1, 0, 0}, Point{0, -1, 0}, Point{0, 0, -1}},
	{Point{-1, 0, 0}, Point{0, 1, 0}, Point{0, 0, -1}},
	{Point{0, 0, 1}, Point{0, 1, 0}, Point{-1, 0, 0}},
}
