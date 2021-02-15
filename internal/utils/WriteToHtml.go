package utils

import (
	"io/ioutil"
	"lagoucoursepull/internal/template"
	"os"
	"strings"
)

// Write2HTML Write2HTML
func Write2HTML(htmlText []byte, htmlDirectory, courseName, title string) error {
	title = strings.Replace(title, "|", "-", -1)
	title = strings.Replace(title, "?", "", -1)
	title = strings.Replace(title, "？", "", -1)
	title = strings.Replace(title, "/", "_", -1)
	title = strings.Replace(title, "\\", "_", -1)
	title = strings.Replace(title, "\"", "", -1)
	title = strings.Replace(title, "'", "", -1)
	dir := htmlDirectory + courseName + "\\"
	err := os.MkdirAll(dir, 0766)
	if err != nil {
		return err
	}
	titleReplacedHTML := strings.Replace(template.LessonTextTemplate, "$LESSON_TITLE", title, 2)
	htmlText = []byte(strings.Replace(titleReplacedHTML, "$SOURCE_HTML_TEXT", string(htmlText), 1))
	err = ioutil.WriteFile(dir+title+".html", htmlText, 0644)
	return err
}
