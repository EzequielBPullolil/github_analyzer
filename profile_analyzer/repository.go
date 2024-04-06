package profileanalyzer

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Repository struct {
	Url, Name, ShorUrl    string
	have_readme, is_empty bool
	cant_commits          int
}

func (r *Repository) GetData() {
	c := colly.NewCollector()

	r.haveReadme(c)
	r.calculateCommits(c)
	c.Visit(r.Url)
}
func (e Repository) HaveReadme() bool { return e.have_readme }
func (e Repository) CantCommits() int { return e.cant_commits }
func (e Repository) IsEmpty() bool    { return e.cant_commits < 5 }

func (e *Repository) calculateCommits(c *colly.Collector) {
	c.OnHTML("span.Text-sc-17v1xeu-0", func(h *colly.HTMLElement) {
		sCommits, f := strings.CutSuffix(h.Text, "Commits")
		if f {
			t, _ := strings.CutSuffix(sCommits, "Commits")
			final := strings.TrimSpace(t)
			n, _ := strconv.Atoi(final)
			e.cant_commits = n
		}
	})
}
func (e *Repository) haveReadme(c *colly.Collector) {
	c.OnHTML("div.react-directory-truncate > a[title]", func(h *colly.HTMLElement) {
		if h.Attr("title") == "README.md" || h.Attr("title") == "README.MD" {
			e.have_readme = true
		}
	})
}
