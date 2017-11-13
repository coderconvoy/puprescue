package main

import (
	"fmt"
	"strings"

	svg "github.com/ajstarks/svgo"
	"github.com/coderconvoy/lazyf"
)

type Card struct {
	Name     string
	Num      int
	Pic      string
	BasePics []string
	Hero     bool
	pfolder  string
}

func NewCard(dt lazyf.LZ, pfolder string) Card {
	res := Card{
		Name:     dt.Name,
		Num:      dt.PIntD(1, "ex0"),
		Pic:      dt.PStringD(strings.ToLower(dt.Name)+".svg", "Pic", "pic"),
		BasePics: dt.PStringAr("Base", "base"),
		Hero:     dt.PBoolD(false, "Hero", "hero"),
		pfolder:  pfolder,
	}
	if len(res.BasePics) != 0 {
		fmt.Println(res.BasePics)
	}
	return res
}

func (c Card) Count() int {
	return c.Num
}

func (c Card) Svg(cw, ch int, g *svg.SVG) {
	fcol := "#ffbbbb"
	if c.Hero {
		fcol = "#bbbbff"
	}
	g.Rect(0, 0, cw, ch, fmt.Sprintf("stroke:black;stroke-width:%d;fill:%s;", cw/20, fcol))

	g.Text(cw/2, (ch*3)/20, c.Name, fmt.Sprintf("text-anchor:middle;font-size:%d;font-weight:bold;", cw/10))

	g.Image(cw/10, cw/10, cw*8/10, ch*8/10, c.Pic)

	blen := len(c.BasePics)
	for i, v := range c.BasePics {
		g.Image(cw/2+cw/16+(i*cw)/4-(blen*cw)/8, (ch*8)/10, cw/8, ch/8, v)
	}
}
