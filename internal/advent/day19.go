package advent

import (
	"strings"

	. "github.com/davidparks11/advent2021/internal/advent/day19"
)

type beaconScanner struct {
	dailyProblem
}

func NewBeaconScanner() Problem {
	return &beaconScanner{
		dailyProblem{
			day: 19,
		},
	}
}

func (s *beaconScanner) Solve() interface{} {
	input := s.GetInputLines()
	var results []int
	results = append(results, s.sharedBeaconCount(input))
	results = append(results, s.maxScannerDistance(input))
	return results
}

func (b *beaconScanner) sharedBeaconCount(input []string) int {
	scanners := b.ParseInput(input)
	current := scanners[0]
	for i := 1; i < len(scanners); i++ {
		if newScanner, _ := current.FindOverlappingBeacons(scanners[i]); newScanner != nil {
			current = newScanner
		} else {
			scanners = append(scanners, scanners[i])
		}
	}
	return len(current.Beacons)
}

func (b *beaconScanner) maxScannerDistance(input []string) int {
	scanners := b.ParseInput(input)
	current := scanners[0]
	offsets := []Point{{}} //populate initial offsets with zero vector for first scanner
	for i := 1; i < len(scanners); i++ {
		if newScanner, distance := current.FindOverlappingBeacons(scanners[i]); newScanner != nil {
			current = newScanner
			offsets = append(offsets, Point(*distance))
		} else {
			scanners = append(scanners, scanners[i])
		}
	}

	maxDistance := 0
	for i := 0; i < len(offsets)-1; i++ {
		for j := i; j < len(offsets); j++ {
			distance := offsets[i].Manhattan(offsets[j])
			if distance > maxDistance {
				maxDistance = distance
			}

		}
	}
	return maxDistance
}

func (b *beaconScanner) ParseInput(input []string) []*Scanner {
	var scanners []*Scanner
	for _, line := range input {
		switch {
		case strings.Contains(line, "scanner"):
			scanners = append(scanners, &Scanner{Beacons: make(map[Point]struct{})})
		case line == "":
			continue
		default:
			nums := CommaSplitInts(line)
			scanners[len(scanners)-1].Beacons[Point{X: nums[0], Y: nums[1], Z: nums[2]}] = struct{}{}
		}
	}

	return scanners
}
