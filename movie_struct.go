package main

import "time"

type movie struct {
	Id int
	title string
	description string
	created_at time.Time
	updated_at time.Time
}
