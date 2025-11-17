package main

import (
	"time"
)

func (j *Journal) newEntry(content string, location string) {
	now := time.Now()
	j.count += 1
	j.entries := append(j.entries, JournalEntry{
		id: j.count,
		time: now.Format("15:04:02"),
		date: now.Format(),
		location: location,
		content: content,
	})
}
