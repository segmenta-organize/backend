package handlers

import (
	"segmenta/src/controllers"

	"github.com/gin-gonic/gin"
)

// Explore Course Handlers

func GetAllExploreCourses(c *gin.Context) {
	controllers.GetAllExploreCourses(c)
}

func GetExploreCourseByID(c *gin.Context) {
	controllers.GetExploreCourseByID(c)
}

func SearchCourses(c *gin.Context) {
	controllers.SearchCourses(c)
}

func GetAllCoursesByCategoryForExplore(c *gin.Context) {
	controllers.GetAllCoursesByCategoryForExplore(c)
}

func EnrollInCourse(c *gin.Context) {
	controllers.EnrollInCourse(c)
}

func EditPublicCourse(c *gin.Context) {
	controllers.EditPublicCourse(c)
}

func DeletePublicCourse(c *gin.Context) {
	controllers.DeletePublicCourse(c)
}

// Explore Course Chapter Handlers

func GetAllExploreChapterByExploreCourseID(c *gin.Context) {
	controllers.GetAllExploreChapterByExploreCourseID(c)
}

func GetOneExploreChapterByCourseID(c *gin.Context) {
	controllers.GetOneExploreChapterByCourseID(c)
}

func CreateExploreChapter(c *gin.Context) {
	controllers.CreateExploreChapter(c)
}

func UpdateExploreChapter(c *gin.Context) {
	controllers.UpdateExploreChapter(c)
}

func DeleteExploreChapter(c *gin.Context) {
	controllers.DeleteExploreChapter(c)
}
