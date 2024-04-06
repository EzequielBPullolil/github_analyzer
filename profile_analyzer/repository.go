package profileanalyzer

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Repository struct {
	Url, Name, ShorUrl string
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
	c.OnHTML("span[data-content]", func(h *colly.HTMLElement) {
		if h.Attr("data-content") == "README" {
			haveReadme = false
		}
	})
	c.Visit(e.Url)
	return haveReadme
}
