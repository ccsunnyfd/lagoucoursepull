package main

import (
	"encoding/json"
	"lagoucoursepull/internal/model"
	"lagoucoursepull/internal/utils"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/jessevdk/go-flags"
)

var opt option

func init() {
	_, err := flags.Parse(&opt)
	if err != nil {
		log.Fatalf("command line flags parse error: %v", err)
	}
}

func main() {
	courseID := opt.CourseID

	// import config
	newConf := utils.GetConf()
	var (
		requestCourseListParams = "courseId=" + strconv.Itoa(courseID) // query params
		cookie                  = newConf.GetCookie()                  // cookie header field
		auth                    = newConf.GetAuth()                    // cookie header authorization
		htmlDir                 = newConf.GetHTMLDir()                 // place to store lesson text files

		requestCourseListURL    = "https://gate.lagou.com/v1/neirong/kaiwu/getCourseLessons"
		requestCourseListMethod = "GET"

		requestLessonTextURL         = "https://gate.lagou.com/v1/neirong/kaiwu/getCourseLessonDetail"
		requestLessonTextMethod      = "GET"
		requestLessonTextParamPrefix = "lessonId="
	)

	// authorization header field
	var clr *model.LagouRequest = model.NewLagouRequestWithParams(requestCourseListMethod, requestCourseListURL, requestCourseListParams)

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
				defer wg.Done()
				time.Sleep(time.Duration(randSource.Intn(500)) * time.Millisecond)
				body, err = lhc.Do(ltr, cookie, auth)
				if err != nil {
					log.Fatal(err)
				}
				var tcr model.TextContentResponse
				json.Unmarshal(body, &tcr)
				htmlText := []byte(tcr.Content.TextContent)

				// Write to html
				err = utils.Write2HTML(htmlText, htmlDir, courseNameComplete, title)
				if err != nil {
					log.Fatal(err)
				}
			}(ltr, title)
		}
	}

	wg.Wait()

}
