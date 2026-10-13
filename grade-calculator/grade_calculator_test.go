package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 55, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 40, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam", 70, Exam)
	gradeCalculator.AddGrade("essay", 70, Essay)

	if actual := gradeCalculator.GetFinalGrade(); actual != "C" {
		t.Errorf("Expected C, got %s", actual)
	}
}

func TestGetGradeD(t *testing.T) {
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam", 60, Exam)
	gradeCalculator.AddGrade("essay", 60, Essay)

	if actual := gradeCalculator.GetFinalGrade(); actual != "D" {
		t.Errorf("Expected D, got %s", actual)
	}
}

func TestEmptyGrades(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	if actual := gradeCalculator.GetFinalGrade(); actual != "F" {
		t.Errorf("Expected F for empty grades, got %s", actual)
	}
}

func TestGradeTypeString(t *testing.T) {
	tests := []struct {
		gradeType GradeType
		expected  string
	}{
		{Assignment, "assignment"},
		{Exam, "exam"},
		{Essay, "essay"},
	}

	for _, test := range tests {
		if actual := test.gradeType.String(); actual != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, actual)
		}
	}
}
