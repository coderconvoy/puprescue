package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/coderconvoy/lazyf"
	"github.com/coderconvoy/msvg"
)

func main() {

	ob := flag.String("o", "", "output-base")
	picfol := flag.String("p", "", "Internal pic root folder")
	flag.Parse()
	if *ob == "" {
		log.Fatal("Needs output base")
	}
	c, err := lazyf.ReadFile("card-data/list.lz")
	if err != nil {
		log.Fatal(err)
	}

	cards := []msvg.Card{}
	for _, v := range c {
		cards = append(cards, NewCard(v, *picfol))
	}
	tot := msvg.Total(cards)
	for i := 0; i < tot; i += 25 {
		bb := msvg.PageA4(25, 5, msvg.CardList(cards, i))
		ioutil.WriteFile(*ob+strconv.Itoa(1+(i/25))+".svg", bb.Bytes(), 0770)
	}

}
