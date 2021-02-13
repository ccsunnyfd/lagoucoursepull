package model

// CourseLesson CourseLesson
type CourseLesson struct {
	ID        int
	CourseID  int
	SectionID int
	Theme     string
}

// CourseSections CourseSections
type CourseSections struct {
	CourseLessons []CourseLesson
}

// CourseContent CourseContent
type CourseContent struct {
	CourseName        string
	CourseSectionList []CourseSections
}

// CourseContentResponse CourseContentResponse
type CourseContentResponse struct {
	Content CourseContent
}
