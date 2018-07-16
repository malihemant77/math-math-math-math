package main

import (
	"math/rand"
	"strings"
	"fmt"
	"strconv"
)

func drill15m() Drill {
	modes := []int{1, 2, 3, 4, 5}

	gen := func(mode int) []int {
		friend := 10 - mode

		var nums []int
		var n = randint(5, 10)

		nums = append(nums, mode)
		nums = append(nums, mode)
		nums = append(nums, friend)
		nums = append(nums, friend)

		for i := 0; i < n-4; i++ {
			r := rand.Float32()
			num := 0
			if r < 0.25 {
				num = mode
			} else if r < 0.5 {
				num = friend
			} else {
				num = randint(1, 10)
			}

			nums = append(nums, num)
		}

		rand.Shuffle(len(nums), func(i, j int) {
			nums[i], nums[j] = nums[j], nums[i]
		})

		return nums
	}

	var sheets []Sheet
	for i := 0; i < 5; i++ {
		mode := modes[i]

		var questions []Question
		for j := 0; j < 15; j++ {
			nums := gen(mode)
			var numsString []string
			for _, num := range nums {
				numsString = append(numsString, fmt.Sprintf("%d", num))
			}

			sum := 0
			for _, num := range nums {
				sum += num
			}

			text := strings.Join(numsString, "  ") + "  =  _____ (   ) [" + strconv.Itoa(checksum(sum)) + "]"

			questions = append(questions, Question{Text: text})
		}

		sheet := Sheet{PageNumber: i + 1, Questions: questions}
		sheet.Intro = fmt.Sprintf("Write the sum of all numbers to the left of the equal sign. (Hint: there are many %ds and %ds). Write the checksum in the parentheses.", mode, 10-mode)

		sheets = append(sheets, sheet)
	}

	return Drill{Name: "15m", Sheets: sheets, ColumnCount: 1, MarginBottom: "2em"}
}
