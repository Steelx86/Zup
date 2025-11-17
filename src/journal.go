package main

import (
	"time"
	"os"
	// "crypto/aes"
)

func createJournal() Journal {
	return Journal{
		entries: nil,
		count: 0,
	}
}

func loadJournal(file *os.File) Journal {
	return Journal {}
}

func writeJournal(file *os.File) string {
	return ""
}

func (j *Journal) newEntry(location string, content string) {
	now := time.Now()
	j.count += 1
	j.entries = append(j.entries, JournalEntry{
		id:       j.count,
		time:     now.Format("15:04:05"),
		date:     now.Format("2025-01-02"),
		location: location,
		content:  content,
	})
}


