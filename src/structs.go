package main

import ()

type Journal struct {
	id uint16
	entries []JournalEntry
	count uint16
}

type JournalEntry struct {
	id uint16
	time string
	date string
	location string
	content string
}

type TaskList struct {
	id uint16
	tasks []Task
	count uint16
	countDone uint16
}

type Task struct {
	id uint16
	name string
	done bool
}
