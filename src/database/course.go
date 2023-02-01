package database

import (
	"github.com/SalomanYu/open-education-courses/src/models"
	"database/sql"
	"fmt"
	"strings"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "edwica"
	password = "123"
	dbname = "courses"
)

func TryConnect() {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", conn)
	checkErr(err)
	defer db.Close()
	err = db.Ping()
	checkErr(err)
}

func AddCourse(course models.Course) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", conn)
	checkErr(err)
	defer db.Close()
	
	smt := `INSERT INTO open_education (url, title, started_at, finished_at, img, skills, description, requirements, duration_in_week, lectures_count, has_certificate, teachers_name, teachers_description, teachers_image) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
	tx, _ := db.Begin()
	_, err = db.Exec(smt, course.Url, course.Title, course.StartedAt, course.FinishedAt, course.Image, course.Skills, course.Description, course.Requirements, course.DurationInWeek, course.LecturesCount, course.HasCertificate, pq.Array(course.TeachersName), pq.Array(course.TeachersDescriptions), pq.Array(course.TeachersImages))
	if err != nil {
		// tx.Rollback()
		panic(err)
	} else {
		tx.Commit()
		fmt.Println("added course:", course.Title)
	}
}

func AddMultipleCourses(data []models.Course) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", conn)
	checkErr(err)
	defer db.Close()
	valueStrings := []string{}
	valueArgs := []interface{}{}
	valueInsertCount := 1
	for _, course := range data {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", valueInsertCount, valueInsertCount+1, valueInsertCount+2, valueInsertCount+3, valueInsertCount+4, valueInsertCount+5, valueInsertCount+6, valueInsertCount+7, valueInsertCount+8, valueInsertCount+9, valueInsertCount+10, valueInsertCount+11, valueInsertCount+12, valueInsertCount+13))
		valueArgs = append(valueArgs, course.Url)
		valueArgs = append(valueArgs, course.Title)
		valueArgs = append(valueArgs, course.StartedAt)
		valueArgs = append(valueArgs, course.FinishedAt)
		valueArgs = append(valueArgs, course.Image)
		valueArgs = append(valueArgs, course.Skills)
		valueArgs = append(valueArgs, course.Description)
		valueArgs = append(valueArgs, course.Requirements)
		valueArgs = append(valueArgs, course.DurationInWeek)
		valueArgs = append(valueArgs, course.LecturesCount)
		valueArgs = append(valueArgs, course.HasCertificate)
		valueArgs = append(valueArgs, pq.Array(course.TeachersName))
		valueArgs = append(valueArgs, pq.Array(course.TeachersDescriptions))
		valueArgs = append(valueArgs, pq.Array(course.TeachersImages))
		valueInsertCount += 14
	}
	if len(valueArgs) == 0 {
		return
	}
	
	smt := `INSERT INTO open_education (url, title, started_at, finished_at, img, skills, description, requirements, duration_in_week, lectures_count, has_certificate, teachers_name, teachers_description, teachers_image) VALUES %s`
	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))
	tx, _ := db.Begin()
	_, err = db.Exec(smt, valueArgs...)
	if err != nil {
		fmt.Println(valueArgs...)
		panic(err)
	} else {
		fmt.Println("added courses:", len(data))
		tx.Commit()
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}