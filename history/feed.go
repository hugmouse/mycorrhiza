package history

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/bouncepaw/mycorrhiza/cfg"

	"github.com/gorilla/feeds"
)

func recentChangesFeed(opts FeedOptions) *feeds.Feed {
	feed := &feeds.Feed{
		Title:       "Recent changes",
		Link:        &feeds.Link{Href: cfg.URL},
		Description: "List of 30 recent changes on the wiki",
		Author:      &feeds.Author{Name: "Wikimind", Email: "wikimind@mycorrhiza"},
		Updated:     time.Now(),
	}
	revs := RecentChanges(30)
	groups := opts.grouping.Group(revs)
	for _, grp := range groups {
		item := grp.feedItem(opts)
		feed.Add(&item)
	}
	return feed
}

// RecentChangesRSS creates recent changes feed in RSS format.
func RecentChangesRSS(opts FeedOptions) (string, error) {
	return recentChangesFeed(opts).ToRss()
}

// RecentChangesAtom creates recent changes feed in Atom format.
func RecentChangesAtom(opts FeedOptions) (string, error) {
	return recentChangesFeed(opts).ToAtom()
}

// RecentChangesJSON creates recent changes feed in JSON format.
func RecentChangesJSON(opts FeedOptions) (string, error) {
	return recentChangesFeed(opts).ToJSON()
}

type revisionGroup []Revision

func newRevisionGroup(rev Revision) revisionGroup {
	return revisionGroup([]Revision{rev})
}

func (grp *revisionGroup) addRevision(rev Revision) {
	*grp = append(*grp, rev)
}

func groupRevisionsByMonth(revs []Revision) (res []revisionGroup) {
	var (
		currentYear  int
		currentMonth time.Month
	)
	for _, rev := range revs {
		if rev.Time.Month() != currentMonth || rev.Time.Year() != currentYear {
			currentYear = rev.Time.Year()
			currentMonth = rev.Time.Month()
			res = append(res, newRevisionGroup(rev))
		} else {
			res[len(res)-1].addRevision(rev)
		}
	}
	return res
}

// groupRevisionsByPeriodFromNow groups close-together revisions.
// If two revisions happened within period of each other, they are put in the same group.
func groupRevisionsByPeriod(revs []Revision, period time.Duration) (res []revisionGroup) {
	if len(revs) == 0 {
		return res
	}

	currTime := revs[0].Time
	res = append(res, newRevisionGroup(revs[0]))
	for _, rev := range revs[1:] {
		if currTime.Sub(rev.Time) < period {
			res[len(res)-1].addRevision(rev)
		} else {
			res = append(res, newRevisionGroup(rev))
		}
		currTime = rev.Time
	}
	return res
}

func (grp revisionGroup) feedItem(opts FeedOptions) feeds.Item {
	return feeds.Item{
		Title:       grp.title(opts.groupOrder),
		Author:      grp.author(),
		Id:          grp[len(grp)-1].Hash,
		Description: grp.descriptionForFeed(opts.groupOrder),
		Created:     grp[len(grp)-1].Time, // earliest revision
		Updated:     grp[0].Time,          // latest revision
		Link:        &feeds.Link{Href: cfg.URL + grp[0].bestLink()},
	}
}

func (grp revisionGroup) title(order FeedGroupOrder) string {
	var message string
	switch order {
	case NewToOld:
		message = grp[0].Message
	case OldToNew:
		message = grp[len(grp)-1].Message
	}

	if len(grp) == 1 {
		return message
	} else {
		return fmt.Sprintf("%d edits (%s, ...)", len(grp), message)
	}
}

func (grp revisionGroup) author() *feeds.Author {
	author := grp[0].Username
	for _, rev := range grp[1:] {
		// if they don't all have the same author, return nil
		if rev.Username != author {
			return nil
		}
	}
	return &feeds.Author{Name: author}
}

func (grp revisionGroup) descriptionForFeed(order FeedGroupOrder) string {
	builder := strings.Builder{}
	switch order {
	case NewToOld:
		for _, rev := range grp {
			builder.WriteString(rev.descriptionForFeed())
		}
	case OldToNew:
		for i := len(grp) - 1; i >= 0; i-- {
			builder.WriteString(grp[i].descriptionForFeed())
		}
	}
	return builder.String()
}

type FeedOptions struct {
	grouping   FeedGrouping
	groupOrder FeedGroupOrder
}

func ParseFeedOptions(query url.Values) (FeedOptions, error) {
	grouping, err := parseFeedGrouping(query)
	if err != nil {
		return FeedOptions{}, err
	}
	groupOrder, err := parseFeedGroupOrder(query)
	if err != nil {
		return FeedOptions{}, err
	}
	return FeedOptions{grouping, groupOrder}, nil
}

type FeedGrouping interface {
	Group([]Revision) []revisionGroup
}

func parseFeedGrouping(query url.Values) (FeedGrouping, error) {
	if query.Get("period") == "" {
		return NormalFeedGrouping{}, nil
	} else {
		period, err := time.ParseDuration(query.Get("period"))
		if err != nil {
			return nil, err
		}
		return PeriodFeedGrouping{Period: period}, nil
	}
}

type NormalFeedGrouping struct{}

func (NormalFeedGrouping) Group(revs []Revision) (res []revisionGroup) {
	for _, rev := range revs {
		res = append(res, newRevisionGroup(rev))
	}
	return res
}

type PeriodFeedGrouping struct {
	Period time.Duration
}

func (g PeriodFeedGrouping) Group(revs []Revision) (res []revisionGroup) {
	return groupRevisionsByPeriod(revs, g.Period)
}

type FeedGroupOrder int

const (
	NewToOld FeedGroupOrder = iota
	OldToNew FeedGroupOrder = iota
)

func parseFeedGroupOrder(query url.Values) (FeedGroupOrder, error) {
	switch query.Get("order") {
	case "oldtonew":
		return OldToNew, nil
	case "newtoold":
	case "":
		return NewToOld, nil
	}
	return 0, errors.New("unknown order")
}
