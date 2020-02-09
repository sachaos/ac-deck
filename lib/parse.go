package lib

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

	selector := doc.Find("tbody > tr > td.text-center > a")
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