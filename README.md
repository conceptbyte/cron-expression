# Cron Expressions

This package is designed to parse existing cron expressions as well as create new expressions using a human friendly fluent API.

It can calculate if a given expression is due for running at the current time and also determine past and future run schedules.

This package is inspired by the amazing PHP cron expression package by @mtdowling.

## Anatomy of a cron expression
A cron expression is a string that denotes the schedule in which certain commands should be executed. Given below is the basic structure of a cron.

    *    *    *    *    *    *
    -    -    -    -    -    -
    |    |    |    |    |    |
    |    |    |    |    |    + year [optional]
    |    |    |    |    +----- day of week (0 - 7) (Sunday=0 or 7)
    |    |    |    +---------- month (1 - 12)
    |    |    +--------------- day of month (1 - 31)
    |    +-------------------- hour (0 - 23)
    +------------------------- min (0 - 59)

## Package API
//TODO: WIP