package models

import (
	"time"
)

type ExploreCourse struct {
	ExploreCourseCourseID uint      `gorm:"primaryKey;autoIncrement;column:explore_course_id" json:"explore_course_id"`
	CreatorID             int       `gorm:"not null;column:creator_id" json:"creator_id"`
	Title                 string    `gorm:"type:varchar(200);not null;column:title" json:"title"`
	Description           *string   `gorm:"type:text;column:description" json:"description,omitempty"`
	Channel               *string   `gorm:"type:varchar(150);column:channel" json:"channel,omitempty"`
	ChannelLink           *string   `gorm:"type:text;column:channel_link" json:"channel_link,omitempty"`
	VideoLink             *string   `gorm:"type:text;column:video_link" json:"video_link,omitempty"`
	ThumbnailImageURL     *string   `gorm:"type:text;column:thumbnail_image_url" json:"thumbnail_image_url,omitempty"`
	Category              *string   `gorm:"type:varchar(100);column:category" json:"category,omitempty"`
	Version               *int      `gorm:"not null;column:version" json:"version"`
	CreatedAt             time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	UpdatedAt             time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
}

type ExploreChapter struct {
	ExploreChapterID uint      `gorm:"primaryKey;autoIncrement;column:explore_chapter_id" json:"explore_chapter_id"`
	ExploreCourseID  uint      `gorm:"not null;column:explore_course_id" json:"explore_course_id"`
	Title            string    `gorm:"type:varchar(200);not null;column:title" json:"title"`
	Description      *string   `gorm:"type:text;column:description" json:"description,omitempty"`
	Order            int       `gorm:"not null;column:order" json:"order"`
	CreatedAt        time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
}
