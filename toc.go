package gofpdf

import (
	"bytes"
	"fmt"
)

var (
	m0 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	m1 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	m2 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	m3 = []string{"", "M", "MM", "MMM", "I̅V̅",
		"V̅", "V̅I̅", "V̅I̅I̅", "V̅I̅I̅I̅", "I̅X̅"}
	m4 = []string{"", "X̅", "X̅X̅", "X̅X̅X̅", "X̅L̅",
		"L̅", "L̅X̅", "L̅X̅X̅", "L̅X̅X̅X̅", "X̅C̅"}
	m5 = []string{"", "C̅", "C̅C̅", "C̅C̅C̅", "C̅D̅",
		"D̅", "D̅C̅", "D̅C̅C̅", "D̅C̅C̅C̅", "C̅M̅"}
	m6 = []string{"", "M̅", "M̅M̅", "M̅M̅M̅"}
)

type Outlines struct {
	Title string
	Level int
	Page  int
}

func FormatRoman(n int) (string, bool) {
	if n < 1 || n >= 4e6 {
		return "", false
	}
	// this is efficient in Go.  the seven operands are evaluated,
	// then a single allocation is made of the exact size needed for the result.
	return m6[n/1e6] + m5[n%1e6/1e5] + m4[n%1e5/1e4] + m3[n%1e4/1e3] +
			m2[n%1e3/1e2] + m1[n%100/10] + m0[n%10],
		true
}

func (f *Fpdf) GetPages() []*bytes.Buffer {
	return f.pages
}

func (f *Fpdf) GetPage() int {
	return f.page
}

func (f *Fpdf) SetPage(page int) {
	f.page = page
}

func (f *Fpdf) GetOutlines() (lines []Outlines) {
	for _, ol := range f.outlines {
		lines = append(lines, Outlines{Title: ol.text, Level: ol.level, Page: ol.p})
	}

	return
}

func (f *Fpdf) SetPages(pages []*bytes.Buffer) {
	f.pages = pages
	f.page = len(pages) - 1
}

// add page to f head
func (f *Fpdf) AheadPages(ff *Fpdf) {
	var tocs []*bytes.Buffer
	tocs = append(tocs, bytes.NewBufferString(""))

	for k := 1; k <= ff.page; k++ {
		tocs = append(tocs, ff.pages[k])
		//f.page++
		f.pageLinks = append(f.pageLinks, ff.pageLinks[k])
	}

	for i := 1; i <= f.page; i++ {
		tocs = append(tocs, f.pages[i])
	}

	f.pages = tocs
	f.page = len(tocs) - 1
}

// add page to f end
func (f *Fpdf) AppendPages(ff *Fpdf) *Fpdf {

	for k := 1; k <= ff.page; k++ {
		fmt.Println("toc ", k)
		f.pages = append(f.pages, ff.pages[k])
		f.pageLinks = append(f.pageLinks, ff.pageLinks[k])
		f.page++
	}

	return f
}

func (f *Fpdf) EndPage() {
	if f.footerFnc != nil {
		f.inFooter = true
		f.footerFnc()
		f.inFooter = false
	}
	f.endpage()
}
