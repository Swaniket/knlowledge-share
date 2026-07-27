package types

import "time"

type LearningTrack struct {
	ID                           string
	Title                        string
	Description                  string
	ThumbnailPath                string // COS path
	TrackOwner                   *User
	VideoIds                     []string
	QuizIds                      []string
	RequiredQuizIds              []string
	ContentOrder                 []string // Represents the order of video and quiz in the learning track
	EstimatedDurationToCompleted time.Time
	TotalEnrolled                int64
	AvgRating                    float64
}

type LearningTrackStatus struct {
	ID            string
	TrackId       string
	UserId        string
	StartDate     time.Time
	LastUpdated   time.Time
	CompletedAt   time.Time
	CurrentStatus string // Not started / In Progress / Completed
	CurrentItemID string // Videoid or quiz id
}
