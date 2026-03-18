package models

import (
	"time"
)

type Course struct {
	CourseID             uint      `gorm:"primaryKey;autoIncrement;column:course_id" json:"course_id"`
	UserID               int       `gorm:"not null;column:user_id" json:"user_id"`
	Title                string    `gorm:"type:varchar(200);not null;column:title" json:"title"`
	Description          *string   `gorm:"type:text;column:description" json:"description,omitempty"`
	Channel              *string   `gorm:"type:varchar(150);column:channel" json:"channel,omitempty"`
	ChannelLink          *string   `gorm:"type:text;column:channel_link" json:"channel_link,omitempty"`
	VideoLink            *string   `gorm:"type:text;column:video_link" json:"video_link,omitempty"`
	ThumbnailImageURL    *string   `gorm:"type:text;column:thumbnail_image_url" json:"thumbnail_image_url,omitempty"`
	Progress             int       `gorm:"default:0;check:progress >= 0 AND progress <= 100;column:progress" json:"progress"`
	SourcePublicCourseID *int      `gorm:"column:source_public_course_id" json:"source_public_course_id,omitempty"`
	SourceVersion        int       `gorm:"not null;column:source_version" json:"source_version"`
	UpdateCheck          bool      `gorm:"not null;default:false;column:update_check" json:"update_check"`
	CreatedAt            time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
}

type Chapter struct {
	ChapterID      uint      `gorm:"primaryKey;autoIncrement;column:chapter_id" json:"chapter_id"`
	CourseID       int       `gorm:"not null;column:course_id" json:"course_id"`
	Title          string    `gorm:"type:varchar(200);not null;column:title" json:"title"`
	VideoTimestamp *string   `gorm:"type:varchar(50);column:video_timestamp" json:"video_timestamp,omitempty"`
	Position       int       `gorm:"not null;check:position > 0;column:position" json:"position"`
	IsCompleted    bool      `gorm:"default:false;column:is_completed" json:"is_completed"`
	CreatedAt      time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
}
