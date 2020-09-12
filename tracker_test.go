package sfapi

import (
	"testing"
)

func TestMileStoneTimeUpdate(t *testing.T) {
	testCases := []struct {
		DueDate  string
		Expected string
	}{
		{
			DueDate:  "2006-01-02 15:04:05",
			Expected: "2006-01-02 15:04:05",
		},
		{
			DueDate:  "",
			Expected: "0001-01-01 00:00:00",
		},
		{
			DueDate:  "invalid date",
			Expected: "0001-01-01 00:00:00",
		},
		{
			DueDate:  "01-02-2006 15:04:05",
			Expected: "0001-01-01 00:00:00",
		},
	}

	for _, testCase := range testCases {
		m := Milestone{
			DueDate: testCase.DueDate,
		}

		formatted := m.DueTime().Format("2006-01-02 15:04:05")
		if formatted != testCase.Expected {
			t.Errorf("%q expected: %q, got %q", testCase.DueDate, testCase.Expected, formatted)
		}
	}
}

func TestTicketTimeUpdat(t *testing.T) {
	testCases := []struct {
		CreatedDate         string
		ModDate             string
		ExpectedCreatedTime string
		ExpectedModTime     string
	}{
		{
			CreatedDate:         "2006-01-02 15:04:05",
			ModDate:             "2006-01-02 15:04:05",
			ExpectedModTime:     "2006-01-02 15:04:05",
			ExpectedCreatedTime: "2006-01-02 15:04:05",
		},
		{
			CreatedDate:         "",
			ModDate:             "",
			ExpectedModTime:     "0001-01-01 00:00:00",
			ExpectedCreatedTime: "0001-01-01 00:00:00",
		},
		{
			CreatedDate:         "invalid date",
			ModDate:             "invalid date",
			ExpectedModTime:     "0001-01-01 00:00:00",
			ExpectedCreatedTime: "0001-01-01 00:00:00",
		},
		{
			CreatedDate:         "01-02-2006 15:04:05",
			ModDate:             "01-02-2006 15:04:05",
			ExpectedModTime:     "0001-01-01 00:00:00",
			ExpectedCreatedTime: "0001-01-01 00:00:00",
		},
	}

	for _, testCase := range testCases {
		m := Ticket{
			CreatedDate: testCase.CreatedDate,
			ModDate:     testCase.ModDate,
		}

		formatted := m.CreatedTime().Format("2006-01-02 15:04:05")
		if formatted != testCase.ExpectedCreatedTime {
			t.Errorf("created expected: %q, got %q", testCase.ExpectedCreatedTime, formatted)
		}
		formatted = m.ModTime().Format("2006-01-02 15:04:05")
		if formatted != testCase.ExpectedModTime {
			t.Errorf("mod expected: %q, got %q", testCase.ExpectedModTime, formatted)
		}
	}
}
