package atcoder

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"strings"
)

func ParseTasksPage(r io.Reader) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	selector := doc.Find("tbody > tr > td:first-child > a")
	paths := make([]string, selector.Length())
	selector.Each(func(i int, selection *goquery.Selection) {
		attr, _ := selection.Attr("href")
		paths[i] = attr
	})

	return paths, nil
}

func ParseTaskPage(r io.Reader) (*Task, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	selector := doc.Find("span.lang-ja .part > section > h3")
	if selector.Length() == 0 {
		return nil, fmt.Errorf("examples not found")
	}

	var inputSelections []*goquery.Selection
	var outputSelections []*goquery.Selection
	selector.Each(func(_ int, selection *goquery.Selection) {
		if strings.HasPrefix(selection.Text(), "入力例") {
			inputSelections = append(inputSelections, selection.Parent().Find("pre"))
		}
		if strings.HasPrefix(selection.Text(), "出力例") {
			outputSelections = append(outputSelections, selection.Parent().Find("pre"))
		}
	})

	if len(inputSelections) != len(outputSelections) {
		return nil, fmt.Errorf("input & output count mismatch")
	}

	examples := make([]*Example, len(inputSelections))
	for i := range inputSelections {
		examples[i] = &Example{
			In:  strings.TrimSpace(inputSelections[i].Text()),
			Exp: strings.TrimSpace(outputSelections[i].Text()),
		}
	}

	name := doc.Find("title").Text()

	return &Task{
		Name: name,
		Examples: examples,
	}, nil
}

func ParseCSRFToken(r io.Reader) (string, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return "", err
	}

	tokenEl := doc.Find("input[name=csrf_token]")
	token, exists := tokenEl.Attr("value")
	if !exists {
		return "", fmt.Errorf("value not found")
	}

	return token, nil
}

func ParseLangVersion(r io.Reader) (string, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return "", err
	}

	var result string
	doc.Find("option").Each(func(_ int, selection *goquery.Selection) {
		val, _ := selection.Attr("value")
		if val == "3001" {
			result = LangOld
		}

		if val == "4001" {
			result = LangNew
		}
	})

	if result == "" {
		return "", fmt.Errorf("cannot detect lang")
	}

	return result, nil
}

func ParseSubmissions(r io.Reader) ([]*Status, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	statuses := []*Status{}
	selector := doc.Find(".panel-submission > .table-responsive > table > tbody > tr")
	selector.Each(func(_ int, selection *goquery.Selection) {
		timeEl := selection.Find("td:nth-child(1) > time")
		problemEl := selection.Find("td:nth-child(2) > a")
		langEl := selection.Find("td:nth-child(4)")
		pointEl := selection.Find("td:nth-child(5)")
		codeLengthEl := selection.Find("td:nth-child(6)")
		resultEl := selection.Find("td:nth-child(7)")

		_, exists := resultEl.Attr("colspan")

		var elapsedTime string
		var memory string
		if !exists {
			elapsedTime = selection.Find("td:nth-child(8)").Text()
			memory = selection.Find("td:nth-child(9)").Text()
		}
		result := strings.TrimSpace(resultEl.Text())

		statuses = append(statuses, &Status{
			SubmissionDate: timeEl.Text(),
			Problem:        problemEl.Text(),
			Language:       langEl.Text(),
			Point:          pointEl.Text(),
			CodeLength:     codeLengthEl.Text(),
			Result:         result,
			ElapsedTime:    elapsedTime,
			Memory:         memory,
		})
	})

	return statuses, nil
}