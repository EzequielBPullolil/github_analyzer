package profileanalyzer

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Repository struct {
	Url, Name string
}

func (e *Repository) CalculateCommits() int {
	cant_commits := 0
	c := colly.NewCollector()
	c.OnHTML("span.Text-sc-17v1xeu-0", func(h *colly.HTMLElement) {
		sCommits, f := strings.CutSuffix(h.Text, "Commits")
		if f {
			t, _ := strings.CutSuffix(sCommits, "Commits")
			final := strings.TrimSpace(t)
			n, _ := strconv.Atoi(final)
			cant_commits = n
		}
	})
	c.Visit(e.Url)
	return cant_commits
}

func (e Repository) HaveReadme() bool {
	var haveReadme = true
	c := colly.NewCollector()
	c.OnHTML("span.Text-sc-17v1xeu-0", func(h *colly.HTMLElement) {
		if h.Attr("class") == "Text-sc-17v1xeu-0 eJOnBv" {

			fmt.Println(h.Text)
		}
	})
	c.Visit(e.Url)
	return haveReadme
}
