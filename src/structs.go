package main

import ()

type ZUP struct {
	name    string
	journal Journal
	planner Planner
}

/* Journal Portion */

type Journal struct {
	entries []JournalEntry
	count   uint16
}

type JournalEntry struct {
	id       uint16
	time     string
	date     string
	location string
	content  string
}

/* Planner Portion */

type Planner struct {

}

type TaskList struct {
	id        uint16
	name      string
	tasks     []Task
	time      string
	count     uint16
	countDone uint16
}

type Task struct {
	id   uint16
	name string
	done bool
}
