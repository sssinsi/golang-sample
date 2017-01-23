package media

import (
	"io"
	"time"
)

type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}

type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

type Audio interface {
	// Stream() (io.ReadCloser, error)
	// RunningTime() time.Duration
	// Format() string //ex: MP3,WAV
	Streamer()
}

type Video interface {
	// Stream() (io.ReadCloser, error)
	// RunningTime() time.Duration
	// Format() string //ex:MP4, WMV
	Streamer()
	Resolution() (x, y int)
}

type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string //ex:MP4, WMV
}
