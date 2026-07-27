package types

import "time"

type Video struct {
	ID                  string
	Title               string
	Description         string
	Duration            time.Time
	VideoRequiredInPath bool
	VideoOwner          *User
	HasResources        bool
	Files               *VideoFiles
	Track               *LearningTrack
	Views               int64
	Likes               int64
}

type VideoFiles struct {
	Video         *COSFile
	Thumbnail     *COSFile
	ResourceLinks []string
	ResourceFiles []*COSFile // Links in COS
}

type COSFile struct {
	COSFileAbsolutePath string
}

type VideoStatus struct {
	ID                 string
	VideoId            string
	Status             string // Not started / In Progress / Completed
	LastViewedDuration time.Time
	TrackStatusId      string // LearningTrackStatus.ID
}

type VideoLikes struct {
	UserId  string
	VideoId string
	LikedAt time.Time
}

type VideoLibrary struct {
	ID           string
	LibraryOwner *User
	VideoIds     []*string
}
