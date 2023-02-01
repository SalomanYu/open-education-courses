package main

import (
	"github.com/SalomanYu/open-education-courses/src/api"
	// "api-open-data/src/database"
	// "api-open-data/src/models"
	"github.com/SalomanYu/open-education-courses/src/mongo"
	"fmt"
	// "sync"
)


func main() {
	links := generatePageListLinks(990)
	for _, item := range links {
		coursesLinks := api.GetCoursesList(item)
		for _, item := range coursesLinks {
			course := api.GetCourse(item)
			mongo.AddCourse(&course)
		}
	}
}

// func main() {
// 	var wg sync.WaitGroup
// 	pages := generatePageListLinks(990)
// 	wg.Add(len(pages))
// 	for _, item := range pages {
// 		go savePage(item, &wg)
// 	}
// 	wg.Wait()
// }


// func savePage(url string, wg *sync.WaitGroup) {
// 	var courses []models.Course
// 	pageLinks := api.GetCoursesList(url)
// 	for _, item := range pageLinks {
// 		courses = append(courses, api.GetCourse(item))
// 	}
// 	database.AddMultipleCourses(courses)
// 	fmt.Println("Added courses from page:", url)
// 	wg.Done()
// }


func generatePageListLinks(count int) (links []string) {
	for i:=1; i<count; i++ {
		links = append(links, fmt.Sprintf("https://courses.openedu.ru/api/courses/v1/courses/?page=%d", i))
	}
	return
}