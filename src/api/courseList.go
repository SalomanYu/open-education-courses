package api

import (
	"github.com/tidwall/gjson"
	"fmt"
)


func GetCoursesList(url string) (links []string) {
	json := getJson(url)
	data := gjson.Get(json, "results").Array()
	for _, i := range data {
		org := i.Get("org")
		number := i.Get("number")
		links = append(links, fmt.Sprintf("https://openedu.ru/api/courses/export/%s/%s?format=json", org, number))
	}
	return
}


