package main

import (
	"fmt"
	"os"
	"strconv"

	"codeberg.org/go-pdf/fpdf"
)

func main() {
	// take size from command line argument, default to 1mm
	size := 1
	if len(os.Args) > 1 {
		chosen_size, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		size = chosen_size
	}

	// init fpdf
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	const (
		w      = 210
		h      = 297
		vscale = 0.9967 // correction for my brother printer
	)

	pdf.SetLineWidth(0.01)
	pdf.SetDrawColor(0, 0, 0)

	// draw grid
	for y := 0; y < h+5; y += size {
		pdf.Line(-1, float64(y)*vscale, w+1, float64(y)*vscale)
	}

	for x := 0; x < w+5; x += size {
		pdf.Line(float64(x), -1, float64(x), h+1)
	}

	// output
	err := pdf.OutputFileAndClose("./graphpaper.pdf")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d mm -> ./graphpaper.pdf\n", size)
}
