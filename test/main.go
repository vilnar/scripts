package main

import (
	"log"
	"regexp"
)

func CheckLink(link string) bool {
	re := regexp.MustCompile(`https?://.+`)
	match := re.MatchString(link)
	return match
}

func main() {
	log.Println("start clean lnk")

	link := "https://t"

	log.Printf("is match %+v\n", CheckLink(link))
}
