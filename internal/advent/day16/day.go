package day16

import (
	"fmt"
	"strconv"
)

//packet types
const (
	typeSum         = iota
	typeProduct     = iota
	typeMin         = iota
	typeMax         = iota
	typeLiteral     = iota
	typeGreaterThan = iota
	typeLessThan    = iota
	typeEqualTo     = iota
)

const (
	lengthType15 = 0
)

type Bits struct {
	bits        []bool
	index       int
	sumVersions bool
}

func (b *Bits) Decode() int {
	return b.unWrap()
}

func (b *Bits) unWrap() int {
	version, packetType := b.versionType()
	result := 0
	if packetType == typeLiteral {
		if b.sumVersions {
			b.parseLiteral()
			result = version
		} else {
			result = b.parseLiteral()
		}
	} else {
		subPackets := b.subPackets()
		if b.sumVersions {
			result = b.sum(subPackets) + version
		} else {
			switch packetType {
			case typeSum:
				result = b.sum(subPackets)
			case typeProduct:
				result = b.product(subPackets)
			case typeMin:
				result = b.min(subPackets)
			case typeMax:
				result = b.max(subPackets)
			case typeGreaterThan:
				result = b.greaterThan(subPackets)
			case typeLessThan:
				result = b.lessThan(subPackets)
			case typeEqualTo:
				result = b.equals(subPackets)
			}
		}
	}

	return result
}

func (b *Bits) subPackets() []int {
	var packets []int
	if b.parseChunk(1) == lengthType15 {
		bitLength := b.parseChunk(15)
		start := b.index
		for b.index < start+bitLength {
			packets = append(packets, b.unWrap())
		}
	} else {
		numPackets := b.parseChunk(11)
		for i := 0; i < numPackets; i++ {
			packets = append(packets, b.unWrap())
		}
	}
	return packets
}

func (b *Bits) sum(packets []int) (sum int) {
	if len(packets) == 0 {
		panic("sum of zero packets not allowed")
	}
	for _, v := range packets {
		sum += v
	}
	return
}

func (b *Bits) product(packets []int) int {
	product := 1
	if len(packets) == 0 {
		panic("product of zero packets not allowed")
	}
	for _, v := range packets {
		product *= v
	}
	return product
}

func (b *Bits) min(packets []int) int {
	if len(packets) == 0 {
		panic("min of zero packets not allowed")
	}
	min := packets[0]
	for _, v := range packets {
		if v < min {
			min = v
		}
	}
	return min
}

func (b *Bits) max(packets []int) int {
	if len(packets) == 0 {
		panic("max of zero packets not allowed")
	}
	max := packets[0]
	for _, v := range packets {
		if v > max {
			max = v
		}
	}
	return max
}

func (b *Bits) greaterThan(packets []int) int {
	if len(packets) != 2 {
		panic(fmt.Sprintf("greater than must have exactly 2 packets, found %d", len(packets)))
	}
	return boolToInt(packets[0] > packets[1])
}

func (b *Bits) lessThan(packets []int) int {
	if len(packets) != 2 {
		panic(fmt.Sprintf("less than must have exactly 2 packets, found %d", len(packets)))
	}
	return boolToInt(packets[0] < packets[1])
}

func (b *Bits) equals(packets []int) int {
	if len(packets) != 2 {
		panic(fmt.Sprintf("equal to must have exactly 2 packets, found %d", len(packets)))
	}
	return boolToInt(packets[0] == packets[1])
}

func (b *Bits) parseLiteral() int {
	var literal int
	var lastChunk bool
	for !lastChunk {
		lastChunk = b.parseChunk(1) == 0
		literal = literal<<4 + b.parseChunk(4)
	}
	return literal
}

func (b *Bits) versionType() (int, int) {
	version := b.parseChunk(3)
	packetType := b.parseChunk(3)

	return version, packetType
}

//parseChunk reads a set of bits from Bits.index to target (excluding
//target) and that set of bits parse as an int.
func (b *Bits) parseChunk(target int) int {
	var result int
	for i := b.index; i < len(b.bits) && i < b.index+target; i++ {
		result = result<<1 + boolToInt(b.bits[i])
	}

	b.index += target
	return result
}

//small helper to get a 1 or 0 out of a bool. There most likely a better way
//to do this.
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

//String returns a clearer string representation of Bits for debugging for
//printing.
func (b Bits) String() string {
	str := fmt.Sprintf("index: %d\t[", b.index)
	for i, v := range b.bits {
		var numVal int
		if v {
			numVal = 1
		}
		if i == 0 {
			str += fmt.Sprintf("%d", numVal)
		} else {
			str += fmt.Sprintf(", %d", numVal)
		}
	}
	return str + "]"
}

func ParseInput(input []string, sumVersions bool) *Bits {
	b := Bits{sumVersions: sumVersions}
	for i := 0; i < len(input[0]); i++ {
		val, err := strconv.ParseUint(input[0][i:i+1], 16, 4)
		if err != nil {
			panic(err.Error())
		}
		for n := 3; n >= 0; n-- {
			b.bits = append(b.bits, val>>n&1 == 1)
		}
	}
	return &b
}
