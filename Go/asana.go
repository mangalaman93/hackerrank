package main

import (
  "fmt"
  "time"
)

const (
  cDateForm = "02/01/2006"
)

var (
  dayToNum = map[string]int{
    "Monday": 1,
    "Tuesday": 2,
    "Wednesday": 3,
    "Thursday": 4,
    "Friday": 5,
    "Saturday": 6,
    "Sunday": 0,
  }
)

func getDate(baseTime time.Time, day string) string {
  daysToAdd := (dayToNum[day] - int(baseTime.Weekday()) + 7) % 7
  newTime := baseTime.AddDate(0, 0, daysToAdd)
  return newTime.Format(cDateForm)
}

func rotateSlice(days []string, day int) []string {
  index := 0
  for i, d := range days {
    if dayToNum[d] >= day {
      index = i
      break
    }
  }

  return append(days[index:], days[:index]...)
}

func recurringTask(firstDate string, k int, daysOfTheWeek []string, n int) []string {
  firstTime, _ := time.Parse(cDateForm, firstDate)
  day := int(firstTime.Weekday())
  daysOfTheWeek = rotateSlice(daysOfTheWeek, day)

  i := 0
  result := make([]string, n)
  for i < n {
    for _, day := range daysOfTheWeek {
      if i == n {
        break
      }

      result[i] = getDate(firstTime, day)
      i++
    }

    firstTime = firstTime.AddDate(0, 0, k*7)
  }

  return result
}

func main() {
  // ["01/01/2015", "05/01/2015", "15/01/2015", "19/01/2015"]
  fmt.Println(recurringTask("01/01/2015", 2, []string{"Monday", "Thursday"}, 4))
}
