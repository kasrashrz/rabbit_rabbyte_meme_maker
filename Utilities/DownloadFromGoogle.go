package utilities

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"log"
)

func SendReq(words []string) {

	response, err := soup.Get("https://google.com/search?q=hi&tbm=isch")

	doc := soup.HTMLParse(response)
	links := doc.FindAll("img")
	for _, link := range links {
		fmt.Println(link.Text(), "| Link :", link.Attrs()["src"])
	}

	if err != nil {
		log.Fatalf("error is : %v", err)
	}

}
