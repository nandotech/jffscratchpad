package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"strconv"
)

func main() {
	background := flag.String("bg", "000000", "background color")
	foreground := flag.String("fg", "000000", "foreground color")
	flag.Parse()

	fg, err := parseColor(*foreground)
	if err != nil {
		log.Fatal(err)
	}
	bg, err := parseColor(*background)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("background is %s\n", *background)

	draw(fg, bg)

}

func parseColor(s string) (color.Color, error) {
	v, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		log.Fatalf("background is not a color: %v", err)
	}
	b := uint8(v % 256)
	g := uint8(v / 256 % 256)
	r := uint8((v / (256 * 256)) % 256)

	return color.RGBA{R: r, G: g, B: b, A: 255}, nil
}

func printColor(c color.Color) string {
	r, g, b, a := c.RGBA()
	r, g, b, a = r/256, g/256, b/256, a/256
	return fmt.Sprintf("rgba(%v,%v,%v,%v)", r, g, b, a)
}

func draw(fg, bg color.Color) {
	//r,g,b,a := c.RGBA()
	fmt.Printf("drawing with foreground %s and background %s\n", printColor(fg), printColor(bg))
}
