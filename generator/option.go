package generator

import "time"

// Option defines log generator options
type Option struct {
	Format    string
	Output    string
	Type      string
	Number    int
	Bytes     int
	Sleep     time.Duration
	Delay     time.Duration
	SplitBy   int
	Overwrite bool
	Forever   bool
}
