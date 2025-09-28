package esepunittests

import "math"

/*
type GradeCalculator struct {
	assignments []Grade
	exams       []Grade
	essays      []Grade
}
*/
type GradeCalculator struct {
	grades []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

const (
	assignment_weight = 0.5
	exam_weight       = 0.35
	essay_weight      = 0.15
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		grades: make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{Name: name, Grade: grade, Type: gradeType})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignment_average := computeAverage(gc.grades, Assignment)
	exam_average := computeAverage(gc.grades, Exam)
	essay_average := computeAverage(gc.grades, Essay)

	weighted_grade := float64(assignment_average)*assignment_weight + float64(exam_average)*exam_weight + float64(essay_average)*essay_weight

	return int(math.Round(weighted_grade))
}

func computeAverage(grades []Grade, gradeType GradeType) int {

	var sum, count int

	for _, g := range grades {
		if g.Type == gradeType {
			sum = sum + g.Grade
			count = count + 1
		}
	}

	if count == 0 {
		return 0
	}

	return sum / count
}
