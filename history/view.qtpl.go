// Code generated by qtc from "view.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line history/view.qtpl:1
package history

//line history/view.qtpl:1
import "fmt"

//line history/view.qtpl:2
import "github.com/bouncepaw/mycorrhiza/cfg"

// HyphaeLinksHTML returns a comma-separated list of hyphae that were affected by this revision as HTML string.

//line history/view.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line history/view.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line history/view.qtpl:5
func (rev Revision) StreamHyphaeLinksHTML(qw422016 *qt422016.Writer) {
//line history/view.qtpl:5
	qw422016.N().S(`
`)
//line history/view.qtpl:7
	for i, hyphaName := range rev.hyphaeAffected() {
//line history/view.qtpl:8
		if i > 0 {
//line history/view.qtpl:8
			qw422016.N().S(`<span aria-hidden="true">, </span>`)
//line history/view.qtpl:10
		}
//line history/view.qtpl:10
		qw422016.N().S(`<a href="/primitive-diff/`)
//line history/view.qtpl:11
		qw422016.E().S(rev.Hash)
//line history/view.qtpl:11
		qw422016.N().S(`/`)
//line history/view.qtpl:11
		qw422016.E().S(hyphaName)
//line history/view.qtpl:11
		qw422016.N().S(`">`)
//line history/view.qtpl:11
		qw422016.E().S(hyphaName)
//line history/view.qtpl:11
		qw422016.N().S(`</a>`)
//line history/view.qtpl:12
	}
//line history/view.qtpl:13
	qw422016.N().S(`
`)
//line history/view.qtpl:14
}

//line history/view.qtpl:14
func (rev Revision) WriteHyphaeLinksHTML(qq422016 qtio422016.Writer) {
//line history/view.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line history/view.qtpl:14
	rev.StreamHyphaeLinksHTML(qw422016)
//line history/view.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line history/view.qtpl:14
}

//line history/view.qtpl:14
func (rev Revision) HyphaeLinksHTML() string {
//line history/view.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
//line history/view.qtpl:14
	rev.WriteHyphaeLinksHTML(qb422016)
//line history/view.qtpl:14
	qs422016 := string(qb422016.B)
//line history/view.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
//line history/view.qtpl:14
	return qs422016
//line history/view.qtpl:14
}

// descriptionForFeed generates a good enough HTML contents for a web feed.

//line history/view.qtpl:17
func (rev *Revision) streamdescriptionForFeed(qw422016 *qt422016.Writer) {
//line history/view.qtpl:17
	qw422016.N().S(`
<p><b>`)
//line history/view.qtpl:18
	qw422016.E().S(rev.Message)
//line history/view.qtpl:18
	qw422016.N().S(`</b> (by `)
//line history/view.qtpl:18
	qw422016.E().S(rev.Username)
//line history/view.qtpl:18
	qw422016.N().S(` at `)
//line history/view.qtpl:18
	qw422016.E().S(rev.TimeString())
//line history/view.qtpl:18
	qw422016.N().S(`)</p>
<p>Hyphae affected: `)
//line history/view.qtpl:19
	rev.StreamHyphaeLinksHTML(qw422016)
//line history/view.qtpl:19
	qw422016.N().S(`</p>
<pre><code>`)
//line history/view.qtpl:20
	qw422016.E().S(rev.textDiff())
//line history/view.qtpl:20
	qw422016.N().S(`</code></pre>
`)
//line history/view.qtpl:21
}

//line history/view.qtpl:21
func (rev *Revision) writedescriptionForFeed(qq422016 qtio422016.Writer) {
//line history/view.qtpl:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line history/view.qtpl:21
	rev.streamdescriptionForFeed(qw422016)
//line history/view.qtpl:21
	qt422016.ReleaseWriter(qw422016)
//line history/view.qtpl:21
}

//line history/view.qtpl:21
func (rev *Revision) descriptionForFeed() string {
//line history/view.qtpl:21
	qb422016 := qt422016.AcquireByteBuffer()
//line history/view.qtpl:21
	rev.writedescriptionForFeed(qb422016)
//line history/view.qtpl:21
	qs422016 := string(qb422016.B)
//line history/view.qtpl:21
	qt422016.ReleaseByteBuffer(qb422016)
//line history/view.qtpl:21
	return qs422016
//line history/view.qtpl:21
}

// WithRevisions returns an html representation of `revs` that is meant to be inserted in a history page.

//line history/view.qtpl:24
func StreamWithRevisions(qw422016 *qt422016.Writer, hyphaName string, revs []Revision) {
//line history/view.qtpl:24
	qw422016.N().S(`
`)
//line history/view.qtpl:25
	for _, grp := range groupRevisionsByMonth(revs) {
//line history/view.qtpl:25
		qw422016.N().S(`
	`)
//line history/view.qtpl:27
		currentYear := grp[0].Time.Year()
		currentMonth := grp[0].Time.Month()
		sectionId := fmt.Sprintf("%d-%d", currentYear, currentMonth)

//line history/view.qtpl:30
		qw422016.N().S(`
<section class="history__month">
	<a href="#`)
//line history/view.qtpl:32
		qw422016.E().S(sectionId)
//line history/view.qtpl:32
		qw422016.N().S(`" class="history__month-anchor">
		<h2 id="`)
//line history/view.qtpl:33
		qw422016.E().S(sectionId)
//line history/view.qtpl:33
		qw422016.N().S(`" class="history__month-title">`)
//line history/view.qtpl:33
		qw422016.N().D(currentYear)
//line history/view.qtpl:33
		qw422016.N().S(` `)
//line history/view.qtpl:33
		qw422016.E().S(currentMonth.String())
//line history/view.qtpl:33
		qw422016.N().S(`</h2>
	</a>
	<ul class="history__entries">
        `)
//line history/view.qtpl:36
		for _, rev := range grp {
//line history/view.qtpl:36
			qw422016.N().S(`
            `)
//line history/view.qtpl:37
			rev.streamasHistoryEntry(qw422016, hyphaName)
//line history/view.qtpl:37
			qw422016.N().S(`
        `)
//line history/view.qtpl:38
		}
//line history/view.qtpl:38
		qw422016.N().S(`
	</ul>
</section>
`)
//line history/view.qtpl:41
	}
//line history/view.qtpl:41
	qw422016.N().S(`
`)
//line history/view.qtpl:42
}

//line history/view.qtpl:42
func WriteWithRevisions(qq422016 qtio422016.Writer, hyphaName string, revs []Revision) {
//line history/view.qtpl:42
	qw422016 := qt422016.AcquireWriter(qq422016)
//line history/view.qtpl:42
	StreamWithRevisions(qw422016, hyphaName, revs)
//line history/view.qtpl:42
	qt422016.ReleaseWriter(qw422016)
//line history/view.qtpl:42
}

//line history/view.qtpl:42
func WithRevisions(hyphaName string, revs []Revision) string {
//line history/view.qtpl:42
	qb422016 := qt422016.AcquireByteBuffer()
//line history/view.qtpl:42
	WriteWithRevisions(qb422016, hyphaName, revs)
//line history/view.qtpl:42
	qs422016 := string(qb422016.B)
//line history/view.qtpl:42
	qt422016.ReleaseByteBuffer(qb422016)
//line history/view.qtpl:42
	return qs422016
//line history/view.qtpl:42
}

//line history/view.qtpl:44
func (rev *Revision) streamasHistoryEntry(qw422016 *qt422016.Writer, hyphaName string) {
//line history/view.qtpl:44
	qw422016.N().S(`
<li class="history__entry">
	<a class="history-entry" href="/rev/`)
//line history/view.qtpl:46
	qw422016.E().S(rev.Hash)
//line history/view.qtpl:46
	qw422016.N().S(`/`)
//line history/view.qtpl:46
	qw422016.E().S(hyphaName)
//line history/view.qtpl:46
	qw422016.N().S(`">
        <time class="history-entry__time">`)
//line history/view.qtpl:47
	qw422016.E().S(rev.timeToDisplay())
//line history/view.qtpl:47
	qw422016.N().S(`</time>
    </a>
	<span class="history-entry__hash"><a href="/primitive-diff/`)
//line history/view.qtpl:49
	qw422016.E().S(rev.Hash)
//line history/view.qtpl:49
	qw422016.N().S(`/`)
//line history/view.qtpl:49
	qw422016.E().S(hyphaName)
//line history/view.qtpl:49
	qw422016.N().S(`">`)
//line history/view.qtpl:49
	qw422016.E().S(rev.Hash)
//line history/view.qtpl:49
	qw422016.N().S(`</a></span>
	<span class="history-entry__msg">`)
//line history/view.qtpl:50
	qw422016.E().S(rev.Message)
//line history/view.qtpl:50
	qw422016.N().S(`</span>
	`)
//line history/view.qtpl:51
	if rev.Username != "anon" {
//line history/view.qtpl:51
		qw422016.N().S(`
        <span class="history-entry__author">by <a href="/hypha/`)
//line history/view.qtpl:52
		qw422016.E().S(cfg.UserHypha)
//line history/view.qtpl:52
		qw422016.N().S(`/`)
//line history/view.qtpl:52
		qw422016.E().S(rev.Username)
//line history/view.qtpl:52
		qw422016.N().S(`" rel="author">`)
//line history/view.qtpl:52
		qw422016.E().S(rev.Username)
//line history/view.qtpl:52
		qw422016.N().S(`</a></span>
    `)
//line history/view.qtpl:53
	}
//line history/view.qtpl:53
	qw422016.N().S(`
</li>
`)
//line history/view.qtpl:55
}

//line history/view.qtpl:55
func (rev *Revision) writeasHistoryEntry(qq422016 qtio422016.Writer, hyphaName string) {
//line history/view.qtpl:55
	qw422016 := qt422016.AcquireWriter(qq422016)
//line history/view.qtpl:55
	rev.streamasHistoryEntry(qw422016, hyphaName)
//line history/view.qtpl:55
	qt422016.ReleaseWriter(qw422016)
//line history/view.qtpl:55
}

//line history/view.qtpl:55
func (rev *Revision) asHistoryEntry(hyphaName string) string {
//line history/view.qtpl:55
	qb422016 := qt422016.AcquireByteBuffer()
//line history/view.qtpl:55
	rev.writeasHistoryEntry(qb422016, hyphaName)
//line history/view.qtpl:55
	qs422016 := string(qb422016.B)
//line history/view.qtpl:55
	qt422016.ReleaseByteBuffer(qb422016)
//line history/view.qtpl:55
	return qs422016
//line history/view.qtpl:55
}
