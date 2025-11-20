package models

import ()

/* Zup data representation */

type Zup struct {
	name    string
	journal Journal
}

/* Journal Portion */

type Journal struct {
	entries []JournalEntry
	count   int
}

type JournalEntry struct {
	id       int
	time     string
	date     string
	location string
	content  string
}

