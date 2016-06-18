package cron

import (
	"time"
)

type Cron struct {
	expression string
}

const (
	MINUTE  = 0
	HOUR    = 1
	DAY     = 2
	MONTH   = 3
	WEEKDAY = 4
	YEAR    = 5
)

/**
 * Named constructor for the cron parser.
 */
func Create(expression string) *Cron {
	switch expression {
	case "@daily":
		expression = "0 0 * * *"
	}
	return &Cron{expression}
}

/**
 * Return the compiled expression
 */
func (cron *Cron) Compile() string {
	return cron.expression
}

/**
 * Validate the given cron expression
 */
func (cron *Cron) Validate() bool {
	//
}

/**
 * Get the next running dates
 */
func (cron *Cron) Next(count int) time.Time {
	//
}

/**
 * Get the previous running dates
 */
func (cron *Cron) Previous(count int) time.Time {
	//
}

/**
 * Check if the cron is currently due.
 */
func (cron *Cron) IsDue() bool {
	//
}
