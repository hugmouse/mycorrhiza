// Code generated by qtc from "history.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/history.qtpl:1
package views

//line views/history.qtpl:1
import "fmt"

//line views/history.qtpl:2
import "net/http"

//line views/history.qtpl:3
import "time"

//line views/history.qtpl:5
import "github.com/bouncepaw/mycorrhiza/cfg"

//line views/history.qtpl:6
import "github.com/bouncepaw/mycorrhiza/util"

//line views/history.qtpl:7
import "github.com/bouncepaw/mycorrhiza/user"

//line views/history.qtpl:8
import "github.com/bouncepaw/mycorrhiza/hyphae"

//line views/history.qtpl:9
import "github.com/bouncepaw/mycorrhiza/history"

//line views/history.qtpl:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/history.qtpl:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/history.qtpl:12
func StreamPrimitiveDiffHTML(qw422016 *qt422016.Writer, rq *http.Request, h *hyphae.Hypha, u *user.User, hash string) {
//line views/history.qtpl:12
	qw422016.N().S(`
`)
//line views/history.qtpl:14
	text, err := history.PrimitiveDiffAtRevision(h.TextPartPath(), hash)
	if err != nil {
		text = err.Error()
	}

//line views/history.qtpl:18
	qw422016.N().S(`
`)
//line views/history.qtpl:19
	StreamNavHTML(qw422016, rq, h.Name, "history")
//line views/history.qtpl:19
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<article>
		<h1>Diff `)
//line views/history.qtpl:23
	qw422016.E().S(util.BeautifulName(h.Name))
//line views/history.qtpl:23
	qw422016.N().S(` at `)
//line views/history.qtpl:23
	qw422016.E().S(hash)
//line views/history.qtpl:23
	qw422016.N().S(`</h1>
		<pre class="codeblock"><code>`)
//line views/history.qtpl:24
	qw422016.E().S(text)
//line views/history.qtpl:24
	qw422016.N().S(`</code></pre>
	</article>
</main>
</div>
`)
//line views/history.qtpl:28
}

//line views/history.qtpl:28
func WritePrimitiveDiffHTML(qq422016 qtio422016.Writer, rq *http.Request, h *hyphae.Hypha, u *user.User, hash string) {
//line views/history.qtpl:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/history.qtpl:28
	StreamPrimitiveDiffHTML(qw422016, rq, h, u, hash)
//line views/history.qtpl:28
	qt422016.ReleaseWriter(qw422016)
//line views/history.qtpl:28
}

//line views/history.qtpl:28
func PrimitiveDiffHTML(rq *http.Request, h *hyphae.Hypha, u *user.User, hash string) string {
//line views/history.qtpl:28
	qb422016 := qt422016.AcquireByteBuffer()
//line views/history.qtpl:28
	WritePrimitiveDiffHTML(qb422016, rq, h, u, hash)
//line views/history.qtpl:28
	qs422016 := string(qb422016.B)
//line views/history.qtpl:28
	qt422016.ReleaseByteBuffer(qb422016)
//line views/history.qtpl:28
	return qs422016
//line views/history.qtpl:28
}

//line views/history.qtpl:30
func StreamRecentChangesHTML(qw422016 *qt422016.Writer, n int) {
//line views/history.qtpl:30
	qw422016.N().S(`
<div class="layout">
<main class="main-width recent-changes">
	<h1>Recent Changes</h1>

	<nav class="recent-changes__count">
		See 
	`)
//line views/history.qtpl:37
	for i, m := range []int{20, 50, 100} {
//line views/history.qtpl:37
		qw422016.N().S(`
	`)
//line views/history.qtpl:38
		if i > 0 {
//line views/history.qtpl:38
			qw422016.N().S(`
		<span aria-hidden="true">|</span>
	`)
//line views/history.qtpl:40
		}
//line views/history.qtpl:40
		qw422016.N().S(`
	`)
//line views/history.qtpl:41
		if m == n {
//line views/history.qtpl:41
			qw422016.N().S(`
		<b>`)
//line views/history.qtpl:42
			qw422016.N().D(m)
//line views/history.qtpl:42
			qw422016.N().S(`</b>
	`)
//line views/history.qtpl:43
		} else {
//line views/history.qtpl:43
			qw422016.N().S(`
		<a href="/recent-changes/`)
//line views/history.qtpl:44
			qw422016.N().D(m)
//line views/history.qtpl:44
			qw422016.N().S(`">`)
//line views/history.qtpl:44
			qw422016.N().D(m)
//line views/history.qtpl:44
			qw422016.N().S(`</a>
	`)
//line views/history.qtpl:45
		}
//line views/history.qtpl:45
		qw422016.N().S(`
	`)
//line views/history.qtpl:46
	}
//line views/history.qtpl:46
	qw422016.N().S(`
		recent changes
	</nav>

	<p><img class="icon" width="20" height="20" src="/static/icon/feed.svg">Subscribe via <a href="/recent-changes-rss">RSS</a>, <a href="/recent-changes-atom">Atom</a> or <a href="/recent-changes-json">JSON feed</a>.</p>

	`)
//line views/history.qtpl:57
	qw422016.N().S(`

	`)
//line views/history.qtpl:60
	changes := history.RecentChanges(n)
	var year, day int
	var month time.Month

//line views/history.qtpl:63
	qw422016.N().S(`
	<section class="recent-changes__list" role="feed">
	`)
//line views/history.qtpl:65
	if len(changes) == 0 {
//line views/history.qtpl:65
		qw422016.N().S(`
		<p>Could not find any recent changes.</p>
	`)
//line views/history.qtpl:67
	} else {
//line views/history.qtpl:67
		qw422016.N().S(`
		`)
//line views/history.qtpl:68
		for i, entry := range changes {
//line views/history.qtpl:68
			qw422016.N().S(`

		`)
//line views/history.qtpl:70
			y, m, d := entry.Time.UTC().Date()

//line views/history.qtpl:70
			qw422016.N().S(`
		`)
//line views/history.qtpl:71
			if d != day || m != month || y != year {
//line views/history.qtpl:71
				qw422016.N().S(`
		<h2 class="recent-changes__heading">
			`)
//line views/history.qtpl:73
				qw422016.E().S(fmt.Sprintf("%04d-%02d-%02d", y, m, d))
//line views/history.qtpl:73
				qw422016.N().S(`
		</h2>
		`)
//line views/history.qtpl:75
				year, month, day = y, m, d

//line views/history.qtpl:75
				qw422016.N().S(`
		`)
//line views/history.qtpl:76
			}
//line views/history.qtpl:76
			qw422016.N().S(`

		<div class="recent-changes__entry" role="article"
		    aria-setsize="`)
//line views/history.qtpl:79
			qw422016.N().D(n)
//line views/history.qtpl:79
			qw422016.N().S(`" aria-posinset="`)
//line views/history.qtpl:79
			qw422016.N().D(i)
//line views/history.qtpl:79
			qw422016.N().S(`">
			 `)
//line views/history.qtpl:80
			qw422016.N().S(recentChangesEntry(entry))
//line views/history.qtpl:80
			qw422016.N().S(`
		</div>

		`)
//line views/history.qtpl:83
		}
//line views/history.qtpl:83
		qw422016.N().S(`
	`)
//line views/history.qtpl:84
	}
//line views/history.qtpl:84
	qw422016.N().S(`
	`)
//line views/history.qtpl:85
	qw422016.N().S(helpTopicBadgeHTML("en", "recent_changes"))
//line views/history.qtpl:85
	qw422016.N().S(`
	</section>
</main>
</div>
`)
//line views/history.qtpl:89
}

//line views/history.qtpl:89
func WriteRecentChangesHTML(qq422016 qtio422016.Writer, n int) {
//line views/history.qtpl:89
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/history.qtpl:89
	StreamRecentChangesHTML(qw422016, n)
//line views/history.qtpl:89
	qt422016.ReleaseWriter(qw422016)
//line views/history.qtpl:89
}

//line views/history.qtpl:89
func RecentChangesHTML(n int) string {
//line views/history.qtpl:89
	qb422016 := qt422016.AcquireByteBuffer()
//line views/history.qtpl:89
	WriteRecentChangesHTML(qb422016, n)
//line views/history.qtpl:89
	qs422016 := string(qb422016.B)
//line views/history.qtpl:89
	qt422016.ReleaseByteBuffer(qb422016)
//line views/history.qtpl:89
	return qs422016
//line views/history.qtpl:89
}

//line views/history.qtpl:91
func streamrecentChangesEntry(qw422016 *qt422016.Writer, rev history.Revision) {
//line views/history.qtpl:91
	qw422016.N().S(`
<div>
	<time class="recent-changes__entry__time">
		`)
//line views/history.qtpl:94
	qw422016.E().S(rev.Time.UTC().Format("15:04 UTC"))
//line views/history.qtpl:94
	qw422016.N().S(`
	</time>
	<span class="recent-changes__entry__message">`)
//line views/history.qtpl:96
	qw422016.E().S(rev.Hash)
//line views/history.qtpl:96
	qw422016.N().S(`</span>

	`)
//line views/history.qtpl:98
	if rev.Username != "anon" {
//line views/history.qtpl:98
		qw422016.N().S(`
	<span class="recent-changes__entry__author">
		&mdash; <a href="/hypha/`)
//line views/history.qtpl:100
		qw422016.E().S(cfg.UserHypha)
//line views/history.qtpl:100
		qw422016.N().S(`/`)
//line views/history.qtpl:100
		qw422016.E().S(rev.Username)
//line views/history.qtpl:100
		qw422016.N().S(`" rel="author">`)
//line views/history.qtpl:100
		qw422016.E().S(rev.Username)
//line views/history.qtpl:100
		qw422016.N().S(`</a>
	</span>
	`)
//line views/history.qtpl:102
	}
//line views/history.qtpl:102
	qw422016.N().S(`
</div>
<div>
	<span class="recent-changes__entry__links">
		`)
//line views/history.qtpl:106
	qw422016.N().S(rev.HyphaeLinksHTML())
//line views/history.qtpl:106
	qw422016.N().S(`
	</span>
	<span class="recent-changes__entry__message">
		`)
//line views/history.qtpl:109
	qw422016.E().S(rev.Message)
//line views/history.qtpl:109
	qw422016.N().S(`
	</span>
</div>
`)
//line views/history.qtpl:112
}

//line views/history.qtpl:112
func writerecentChangesEntry(qq422016 qtio422016.Writer, rev history.Revision) {
//line views/history.qtpl:112
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/history.qtpl:112
	streamrecentChangesEntry(qw422016, rev)
//line views/history.qtpl:112
	qt422016.ReleaseWriter(qw422016)
//line views/history.qtpl:112
}

//line views/history.qtpl:112
func recentChangesEntry(rev history.Revision) string {
//line views/history.qtpl:112
	qb422016 := qt422016.AcquireByteBuffer()
//line views/history.qtpl:112
	writerecentChangesEntry(qb422016, rev)
//line views/history.qtpl:112
	qs422016 := string(qb422016.B)
//line views/history.qtpl:112
	qt422016.ReleaseByteBuffer(qb422016)
//line views/history.qtpl:112
	return qs422016
//line views/history.qtpl:112
}

//line views/history.qtpl:114
func StreamHistoryHTML(qw422016 *qt422016.Writer, rq *http.Request, hyphaName, list string) {
//line views/history.qtpl:114
	qw422016.N().S(`
`)
//line views/history.qtpl:115
	StreamNavHTML(qw422016, rq, hyphaName, "history")
//line views/history.qtpl:115
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<article class="history">
		<h1>History of `)
//line views/history.qtpl:119
	qw422016.E().S(util.BeautifulName(hyphaName))
//line views/history.qtpl:119
	qw422016.N().S(`</h1>
		`)
//line views/history.qtpl:120
	qw422016.N().S(list)
//line views/history.qtpl:120
	qw422016.N().S(`
	</article>
</main>
</div>
`)
//line views/history.qtpl:124
}

//line views/history.qtpl:124
func WriteHistoryHTML(qq422016 qtio422016.Writer, rq *http.Request, hyphaName, list string) {
//line views/history.qtpl:124
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/history.qtpl:124
	StreamHistoryHTML(qw422016, rq, hyphaName, list)
//line views/history.qtpl:124
	qt422016.ReleaseWriter(qw422016)
//line views/history.qtpl:124
}

//line views/history.qtpl:124
func HistoryHTML(rq *http.Request, hyphaName, list string) string {
//line views/history.qtpl:124
	qb422016 := qt422016.AcquireByteBuffer()
//line views/history.qtpl:124
	WriteHistoryHTML(qb422016, rq, hyphaName, list)
//line views/history.qtpl:124
	qs422016 := string(qb422016.B)
//line views/history.qtpl:124
	qt422016.ReleaseByteBuffer(qb422016)
//line views/history.qtpl:124
	return qs422016
//line views/history.qtpl:124
}
