package main

import "fmt"

type SkinColors int32

const (
	BLACK = iota
	WHITE
	YELLOW
	RED
	BLUE
)

var skinColors = map[SkinColors]string{
	BLACK:  "brown",
	WHITE:  "cream",
	YELLOW: "asian",
}

func (sc SkinColors) returnerOfSkinColor() string {
	return skinColors[sc]
}

func main() {
	fmt.Print(skinColors[BLACK])
}
