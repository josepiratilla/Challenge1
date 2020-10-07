package drum

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

const trackLenght = 16

type track struct {
	number  int
	name    string
	pattern []bool
}

type pattern struct {
	version string
	tempo   float32
	tracks  []track
}

//DecodeFile main function for the challenge
func DecodeFile(file string) (string, error) {

	data, _ := ioutil.ReadFile(file)
	p := new(pattern)
	pointer := 6 //We skip the SPLICE prefix
	pendingData := int(binary.BigEndian.Uint64(data[pointer : pointer+8]))
	pointer += 8
	p.version = readString(data[pointer : pointer+32])
	pointer += 32
	p.tempo = readFloat32(data[pointer : pointer+4])
	pointer += 4
	for i := 0; pointer < pendingData; i++ {
		newTrack := new(track)
		p.tracks = append(p.tracks, *newTrack)
		p.tracks[i].number = int(data[pointer])
		pointer++
		nameLength := int(binary.BigEndian.Uint32(data[pointer : pointer+4]))
		pointer += 4
		p.tracks[i].name = string(data[pointer : pointer+nameLength])
		pointer += nameLength
		p.tracks[i].pattern = readBools(data[pointer : pointer+trackLenght])
		pointer += trackLenght
	}

	return p.String(), nil
}

func readString(b []byte) string {
	var j int
	for i := range b {
		if b[i] == 0 {
			j = i
			break
		}
	}
	return string(b[:j])
}

func readFloat32(b []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(b))
}

func readBools(b []byte) []bool {
	out := make([]bool, len(b))
	for i := range b {
		if b[i] == 0 {
			out[i] = false
		} else {
			out[i] = true
		}
	}
	return out
}

func (p *pattern) String() string {
	var b strings.Builder

	b.Grow(500) //To avoid reallocation

	b.WriteString(fmt.Sprintf("Saved with HW Version: %s\n", p.version))
	if p.tempo == float32(int(p.tempo)) {
		b.WriteString(fmt.Sprintf("Tempo: %.0f\n", p.tempo))
	} else {
		b.WriteString(fmt.Sprintf("Tempo: %.1f\n", p.tempo))
	}
	for i := range p.tracks {
		b.WriteString(p.tracks[i].String())
	}

	return b.String()
}

func (t *track) String() string {
	var b strings.Builder

	b.Grow(22)
	b.WriteString(fmt.Sprintf("(%d) %s	", t.number, t.name))
	b.WriteString("|")
	for i := range t.pattern {
		if t.pattern[i] {
			b.WriteString("x")
		} else {
			b.WriteString("-")
		}
		if i%4 == 3 {
			b.WriteString("|")
		}
	}
	b.WriteString("\n")
	return b.String()

}

//
// [00-05]  6 bytes: 6 characters -> "SPLICE"
// [06-0D]  8 bytes: UINT64 (big endian) -> Lenght of the rest of the file (excluding the first 14 bytes)
// [0E-2D] 32 bytes: String with the "Saved with HW Version Value", filling with nulls.
// [2E-31]  4 bytes: FLOAT32 (little endian) -> Tempo
// From here we start a track
//    1 bytes: UINT: The track number
//    4 bytes: INT (Big endian) The length of the next string
//	  N bytes: String of char ending at "01" The name of the track
//	 16 bytes: [16]BOOL, represented each one in a byte ("01" true, "00" false)
