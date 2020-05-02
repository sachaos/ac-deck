package atcoder

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseLangVersion(t *testing.T) {
	{
		file, err := os.Open("testdata/submit_old.html")
		if err != nil {
			t.Fatal(err)
		}

		version, err := ParseLangVersion(file)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, LangOld, version)
	}

	{
		file, err := os.Open("testdata/submit_new.html")
		if err != nil {
			t.Fatal(err)
		}

		version, err := ParseLangVersion(file)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, LangNew, version)
	}
}

func TestParseTaskPage(t *testing.T) {
	{
		file, err := os.Open("testdata/abc153_a.html")
		if err != nil {
			t.Fatal(err)
		}

		task, err := ParseTaskPage(file)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, &Task{
			Name: "A - Serval vs Monster",
			Examples: []*Example{
				{
					In:  "10 4",
					Exp: "3",
				},
				{
					In:  "1 10000",
					Exp: "1",
				},
				{
					In:  "10000 1",
					Exp: "10000",
				},
			},
		}, task)
	}

	{
		file, err := os.Open("testdata/abc153_b.html")
		if err != nil {
			t.Fatal(err)
		}

		task, err := ParseTaskPage(file)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, &Task{
			Name: "B - Common Raccoon vs Monster",
			Examples: []*Example{
				{
					In:  "10 3\n4 5 6",
					Exp: "Yes",
				},
				{
					In:  "20 3\n4 5 6",
					Exp: "No",
				},
				{
					In:  "210 5\n31 41 59 26 53",
					Exp: "Yes",
				},
				{
					In:  "211 5\n31 41 59 26 53",
					Exp: "No",
				},
			},
		}, task)
	}
}

func TestParseTasksPage(t *testing.T) {
	file, err := os.Open("testdata/abc153_tasks.html")
	if err != nil {
		t.Fatal(err)
	}

	tasksPaths, err := ParseTasksPage(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, []string{
		"/contests/abc153/tasks/abc153_a",
		"/contests/abc153/tasks/abc153_b",
		"/contests/abc153/tasks/abc153_c",
		"/contests/abc153/tasks/abc153_d",
		"/contests/abc153/tasks/abc153_e",
		"/contests/abc153/tasks/abc153_f",
	}, tasksPaths)
}

func TestParseSubmissions(t *testing.T) {
	t.Run("Parse Waiting Response", func(t *testing.T) {
		file, err := os.Open("testdata/submissions_waiting.html")
		if err != nil {
			t.Fatal(err)
		}

		statuses, err := ParseSubmissions(file)
		if err != nil {
			t.Error(err)
		}
		status := statuses[0]

		assert.Equal(t, &Status{
			SubmissionDate: "2020-03-28 17:02:29+0900",
			Problem:        "D - String Equivalence",
			Language:       "Python3 (3.4.3)",
			Point:          "0",
			CodeLength:     "250 Byte",
			Result:         "WJ",
			ElapsedTime:    "",
			Memory:         "",
		}, status)
	})

	t.Run("Parse Finished Response", func(t *testing.T) {
		file, err := os.Open("testdata/submissions_waiting.html")
		if err != nil {
			t.Fatal(err)
		}

		statuses, err := ParseSubmissions(file)
		if err != nil {
			t.Error(err)
		}
		status := statuses[1]

		assert.Equal(t, &Status{
			SubmissionDate: "2020-03-28 16:55:10+0900",
			Problem:        "D - String Equivalence",
			Language:       "Python3 (3.4.3)",
			Point:          "400",
			CodeLength:     "250 Byte",
			Result:         "AC",
			ElapsedTime:    "114 ms",
			Memory:         "4340 KB",
		}, status)
	})

	t.Run("Parse Processing Response", func(t *testing.T) {
		file, err := os.Open("testdata/submissions_processing.html")
		if err != nil {
			t.Fatal(err)
		}

		statuses, err := ParseSubmissions(file)
		if err != nil {
			t.Error(err)
		}
		status := statuses[0]

		assert.Equal(t, &Status{
			SubmissionDate: "2020-03-28 17:06:34+0900",
			Problem:        "D - String Equivalence",
			Language:       "Python3 (3.4.3)",
			Point:          "0",
			CodeLength:     "250 Byte",
			Result:         "2/10",
			ElapsedTime:    "",
			Memory:         "",
		}, status)
	})
}
