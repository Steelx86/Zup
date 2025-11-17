package main

import ()

type JournalEntry struct {
	id uint
	time string
	day string
	location string
	content string
}

type TaskList struct {
	id uint
	tasks []Task
	numberDone uint16
	numberAvailable uint16
}

type Task struct {
	id uint
	name string
	done bool
}
