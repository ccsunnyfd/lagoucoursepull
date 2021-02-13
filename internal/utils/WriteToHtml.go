package utils

import (
	"io/ioutil"
	"os"
	"strings"
)

// WriteToHTML WriteToHTML
func WriteToHTML(htmlText []byte, htmlDirectory, courseName, title string) error {
	title = strings.Replace(title, "|", "-", -1)
	title = strings.Replace(title, "?", "", -1)
	title = strings.Replace(title, "？", "", -1)
	title = strings.Replace(title, "/", "_", -1)
	title = strings.Replace(title, "\\", "_", -1)
	dir := htmlDirectory + courseName + "\\"
	err := os.MkdirAll(dir, 0766)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dir+title+".html", htmlText, 0644)
	return err
}
