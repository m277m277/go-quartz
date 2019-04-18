package quartz

import (
	"fmt"
	"strconv"
	"time"
)

func indexes(search []string, target []string) []int {
	searchIndexes := make([]int, 0, len(search))
	for _, a := range search {
		searchIndexes = append(searchIndexes, intVal(target, a))
	}
	return searchIndexes
}

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func fillRange(from int, to int) ([]int, error) {
	if to < from {
		return nil, cronError("fillRange")
	}
	len := (to - from) + 1
	arr := make([]int, len)
	for i, j := from, 0; i <= to; i, j = i+1, j+1 {
		arr[j] = i
	}
	return arr, nil
}

func fillStep(from int, step int, max int) ([]int, error) {
	if max < from {
		return nil, cronError("fillStep")
	}
	len := ((max - from) / step) + 1
	arr := make([]int, len)
	for i, j := from, 0; i <= max; i, j = i+step, j+1 {
		arr[j] = i
	}
	return arr, nil
}

func normalize(field string, tr []string) int {
	i, err := strconv.Atoi(field)
	if err == nil {
		return i
	}
	return intVal(tr, field)
}

func inScope(i int, min int, max int) bool {
	if i >= min && i <= max {
		return true
	}
	return false
}

func cronError(cause string) error {
	return fmt.Errorf("Invalid cron expression: %s", cause)
}

// Align single digit values for time.UnixDate format
func alignDigit(next int, prefix string) string {
	if next < 10 {
		return prefix + strconv.Itoa(next)
	}
	return strconv.Itoa(next)
}

func step(prev int, next int, max int) int {
	diff := next - prev
	if diff < 0 {
		return diff + max
	}
	return diff
}

func intVal(target []string, search string) int {
	for i, v := range target {
		if v == search {
			return i
		}
	}
	return -1 //TODO: return error
}

// Unsafe strconv.Atoi
func atoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func maxDays(month int, year int) int {
	if month == 2 && isLeapYear(year) {
		return 29
	}
	return daysInMonth[month]
}

func isLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	} else if year%100 != 0 {
		return true
	} else if year%400 != 0 {
		return false
	}
	return true
}

func dayOfTheWeek(y int, m int, d int) string {
	t := []int{0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4}
	if m < 3 {
		y--
	}
	return days[((y + y/4 - y/100 + y/400 + t[m-1] + d) % 7)]
}

func NowNano() int64 {
	return time.Now().UTC().UnixNano()
}

func isOutdated(_time int64) bool {
	return _time < NowNano()-(time.Second*30).Nanoseconds()
}
