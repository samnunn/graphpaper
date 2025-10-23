package main

import "codeberg.org/go-pdf/fpdf"

func main() {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	const (
		w      = 210
		h      = 297
		vscale = 0.9967 // correction for my brother printer
	)

	pdf.SetLineWidth(0.01)
	pdf.SetDrawColor(0, 0, 0)

	for y := range h + 5 {
		pdf.Line(-1, float64(y)*vscale, w+1, float64(y)*vscale)
	}

	for x := range w + 5 {
		pdf.Line(float64(x), -1, float64(x), h+1)
	}

	err := pdf.OutputFileAndClose("/tmp/graphpaper.pdf")
	if err != nil {
		panic(err)
	}
	println("/tmp/graphpaper.pdf")
}
