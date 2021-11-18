package utilities

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"log"
)

func SendReq(words []string) {
	testWord := "rabbit"
	response, err := soup.Get("https://google.com/search?q=" + testWord + "&tbm=isch")
	doc := soup.HTMLParse(response)
	links := doc.FindAll("img")
	for _, link := range links {
		//fmt.Println(link.Text(), "| Link :", link.Attrs()["src"])
		fmt.Println(link.Attrs()["src"])
	}
	//fmt.Println(links)

	if err != nil {
		log.Fatalf("error is : %v", err)
	}
}
