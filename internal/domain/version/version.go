package version

import "time"

type Version struct {
	ID            int       `json:"id"`
	NumberVersion string    `json:"number_version"`
	Changelog     string    `json:"changelog"`
	CreatedAt     time.Time `json:"created_at"`
}

//dtos
