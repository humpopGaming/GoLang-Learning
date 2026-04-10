package main

import "fmt"

func isBetween(score, low, high int) bool {
	return score >= low && score <= high
}

func letterGrade(score int) string {
	switch {
	case isBetween(score, 90, 100):
		return "A"
	case isBetween(score, 80, 89):
		return "B"
	case isBetween(score, 70, 79):
		return "C"
	case isBetween(score, 60, 69):
		return "D"
	case isBetween(score, 0, 59):
		return "F"
	}
	return ""
}

func gradeDescription(grade string) string {
	switch grade {
	case "A":
		return "Excellent"
	case "B":
		return "Good"
	case "C":
		return "Average"
	case "D":
		return "Below Average"
	case "F":
		return "Failing"
	}
	return "Invalid grade"
}

func main() {
	scores := []int{95, 82, 74, 65, 43}
	for i := 0; i < len(scores); i++ {
		score := scores[i]
		lg := letterGrade(score)
		gd := gradeDescription(lg)

		fmt.Printf("Score: %v → Grade: %v (%v)\n", score, lg, gd)
	}
}
