package main

import (
	"encoding/json"
	"fmt"
	"lagoucoursepull/internal/model"
	"lagoucoursepull/internal/utils"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	var requestCourseListParams string = "courseId=594" // query params
	var cookie string = ``                              // cookie header field
	var auth string = ""                                // authorization header field

	var htmlDirectory string = "C:\\拉勾教育\\"

	var requestCourseListURL string = "https://gate.lagou.com/v1/neirong/kaiwu/getCourseLessons"
	var requestCourseListMethod string = "GET"
	var clr *model.LagouRequest = model.NewLagouRequestWithParams(requestCourseListMethod, requestCourseListURL, requestCourseListParams)

	var requestLessonTextURL string = "https://gate.lagou.com/v1/neirong/kaiwu/getCourseLessonDetail"
	var requestLessonTextMethod string = "GET"
	var requestLessonTextParamPrefix string = "lessonId="

	lhc := utils.NewLagouHTTPClient()

	// Pull courseList from lagou edu
	body, err := lhc.Do(clr, cookie, auth)
	if err != nil {
		log.Fatal(err)
	}
	var ccr model.CourseContentResponse
	json.Unmarshal(body, &ccr)

	courseSectionList := ccr.Content.CourseSectionList
	courseName := ccr.Content.CourseName

	var wg sync.WaitGroup
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, courseSection := range courseSectionList {
		for _, courseLesson := range courseSection.CourseLessons {
			courseNameComplete := strconv.Itoa(courseLesson.CourseID) + "-" + courseName
			var lessonID int = courseLesson.ID
			var title string = courseLesson.Theme
			// Pull lessonText from lagou edu
			params := requestLessonTextParamPrefix + strconv.Itoa(lessonID)
			var ltr *model.LagouRequest = model.NewLagouRequestWithParams(requestLessonTextMethod, requestLessonTextURL, params)
			wg.Add(1)
			go func(*model.LagouRequest, string) {
				time.Sleep(time.Duration(randSource.Intn(500)) * time.Millisecond)
				body, err = lhc.Do(ltr, cookie, auth)
				if err != nil {
					log.Fatal(err)
				}
				if strings.Contains(title, "事务处理与恢复（下）") {
					fmt.Println(string(body))
				}
				var tcr model.TextContentResponse
				json.Unmarshal(body, &tcr)
				htmlText := []byte(tcr.Content.TextContent)

				// Write to html
				err = utils.WriteToHTML(htmlText, htmlDirectory, courseNameComplete, title)
				if err != nil {
					log.Fatal(err)
				}
				wg.Done()
			}(ltr, title)
		}
	}

	wg.Wait()

}
