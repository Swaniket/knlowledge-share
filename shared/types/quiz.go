package types

import "time"

type Quiz struct {
	ID                      string
	Title                   string
	QuizOwner               *User
	PassThreshold           float64
	EstimatedTimeToComplete time.Time
	Questions               []*QuizQuestions
}

type QuizQuestions struct {
	ID           string
	QuestionText string
	Options      []*Options
	Answer       *Options
}

type Options struct {
	ID    string
	Value string
}

type QuizLibrary struct {
	ID           string
	LibraryOwner *User
	QuizIds      []string
}
