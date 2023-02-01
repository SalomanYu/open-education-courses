package api

import (
	"github.com/SalomanYu/open-education-courses/src/models"

	"github.com/tidwall/gjson"
)

func GetCourse(url string) (course models.Course) {
	json := getJson(url)
	course.Title = gjson.Get(json, "title").String()
	course.Description = gjson.Get(json, "description").String()
	course.StartedAt = gjson.Get(json, "started_at").String()
	course.FinishedAt = gjson.Get(json, "finished_at").String()
	course.Skills = gjson.Get(json, "results").String()
	course.Image = gjson.Get(json, "image").String()
	course.Requirements = gjson.Get(json, "requirements").String()
	course.Url = gjson.Get(json, "external_url").String()
	course.LecturesCount = gjson.Get(json, "lectures").Int()
	course.DurationInWeek = gjson.Get(json, "duration.value").Int()
	course.HasCertificate = gjson.Get(json, "cert").Bool()
	parseTeachers(&course, json)
	return
}


func parseTeachers(course *models.Course, json string){
	data := gjson.Get(json, "teachers").Array()
	for _, i := range data {
		course.TeachersName = append(course.TeachersName, i.Get("display_name").String())
		course.TeachersDescriptions = append(course.TeachersDescriptions, i.Get("description").String())
		course.TeachersImages = append(course.TeachersImages, i.Get("image").String())
	}
}