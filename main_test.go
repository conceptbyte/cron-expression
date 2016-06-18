package cron

import (
	"testing"
)

/**
 * Compile the cron expression
 * @param {testing} t *testing.T Compile Test
 */
func CompileTest(t *testing.T) {
	c := Create("******")

	if c.Compile() != "******" {
		t.Error("Expression did not match")
	}
}

/**
 * Create a daily cron
 * @param {testing} t *testing.T Create Daily
 */
func CreateDailyTest(t *testing.T) {
	c := Create("@daily")

	if c.Compile() != "0 0 * * *" {
		t.Error("Expression did not match")
	}
}
