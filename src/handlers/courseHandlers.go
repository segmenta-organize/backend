package handlers

import (
	"segmenta/src/controllers"

	"github.com/gin-gonic/gin"
)

// Course Handlers

func GetCourses(c *gin.Context) {
	controllers.GetAllCourses(c)
}

func GetCourseByID(c *gin.Context) {
	controllers.GetCourseByID(c)
}

func CreateCourse(c *gin.Context) {
	controllers.CreateCourse(c)
}

func AutoCreateCourse(c *gin.Context) {
	controllers.AutoCreateCourses(c)
}

func UpdateCourse(c *gin.Context) {
	controllers.UpdateCourse(c)
}

func DeleteCourse(c *gin.Context) {
	controllers.DeleteCourse(c)
}

func CreatePublicCourseFromCourse(c *gin.Context) {
	controllers.CreatePublicCourseFromCourse(c)
}

func UpdatePublicCourseFromCourse(c *gin.Context) {
	controllers.UpdatePublicCourseFromCourse(c)
}

// Chapter Handlers

func GetAllChaptersByCourseID(c *gin.Context) {
	controllers.GetAllChaptersByCourseID(c)
}

func GetOneChapterByID(c *gin.Context) {
	controllers.GetOneChapterByID(c)
}

func CreateChapter(c *gin.Context) {
	controllers.CreateChapter(c)
}

func UpdateChapter(c *gin.Context) {
	controllers.UpdateChapter(c)
}

func DeleteChapter(c *gin.Context) {
	controllers.DeleteChapter(c)
}

// Chapter → Explore Chapter public handlers

func CreatePublicChapterFromChapter(c *gin.Context) {
	controllers.CreatePublicChapterFromChapter(c)
}

func UpdatePublicChapterFromChapter(c *gin.Context) {
	controllers.UpdatePublicChapterFromChapter(c)
}

func DeletePublicChapterFromChapter(c *gin.Context) {
	controllers.DeletePublicChapterFromChapter(c)
}
