// Code generated by qtc from "readers.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/readers.qtpl:1
package views

//line views/readers.qtpl:1
import "net/http"

//line views/readers.qtpl:2
import "strings"

//line views/readers.qtpl:3
import "path"

//line views/readers.qtpl:4
import "os"

//line views/readers.qtpl:6
import "github.com/bouncepaw/mycorrhiza/cfg"

//line views/readers.qtpl:7
import "github.com/bouncepaw/mycorrhiza/hyphae"

//line views/readers.qtpl:8
import "github.com/bouncepaw/mycorrhiza/mimetype"

//line views/readers.qtpl:9
import "github.com/bouncepaw/mycorrhiza/tree"

//line views/readers.qtpl:10
import "github.com/bouncepaw/mycorrhiza/user"

//line views/readers.qtpl:11
import "github.com/bouncepaw/mycorrhiza/util"

//line views/readers.qtpl:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/readers.qtpl:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/readers.qtpl:13
func StreamAttachmentMenuHTML(qw422016 *qt422016.Writer, rq *http.Request, h *hyphae.Hypha, u *user.User) {
//line views/readers.qtpl:13
	qw422016.N().S(`
<div class="layout">
<main class="main-width attachment-tab">
	<h1>Attachment of `)
//line views/readers.qtpl:16
	qw422016.N().S(beautifulLink(h.Name))
//line views/readers.qtpl:16
	qw422016.N().S(`</h1>
	`)
//line views/readers.qtpl:17
	if h.BinaryPath == "" {
//line views/readers.qtpl:17
		qw422016.N().S(`
	<p class="explanation">This hypha has no attachment, you can upload it here. <a href="/help/en/attachment" class="shy-link">What are attachments?</a></p>
	`)
//line views/readers.qtpl:19
	} else {
//line views/readers.qtpl:19
		qw422016.N().S(`
	<p class="explanation">You can manage the hypha's attachment on this page. <a href="/help/en/attachment" class="shy-link">What are attachments?</a></p>
	`)
//line views/readers.qtpl:21
	}
//line views/readers.qtpl:21
	qw422016.N().S(`

	<section class="amnt-grid">

	`)
//line views/readers.qtpl:25
	if h.BinaryPath != "" {
//line views/readers.qtpl:25
		qw422016.N().S(`
		`)
//line views/readers.qtpl:27
		mime := mimetype.FromExtension(path.Ext(h.BinaryPath))
		fileinfo, err := os.Stat(h.BinaryPath)

//line views/readers.qtpl:28
		qw422016.N().S(`
		`)
//line views/readers.qtpl:29
		if err == nil {
//line views/readers.qtpl:29
			qw422016.N().S(`
		<fieldset class="amnt-menu-block">
			<legend class="modal__title modal__title_small">Stat</legend>
			<p class="modal__confirmation-msg"><b>File size:</b> `)
//line views/readers.qtpl:32
			qw422016.N().DL(fileinfo.Size())
//line views/readers.qtpl:32
			qw422016.N().S(` bytes</p>
			<p><b>MIME type:</b> `)
//line views/readers.qtpl:33
			qw422016.E().S(mime)
//line views/readers.qtpl:33
			qw422016.N().S(`</p>
		</fieldset>
		`)
//line views/readers.qtpl:35
		}
//line views/readers.qtpl:35
		qw422016.N().S(`

		`)
//line views/readers.qtpl:37
		if strings.HasPrefix(mime, "image/") {
//line views/readers.qtpl:37
			qw422016.N().S(`
		<fieldset class="amnt-menu-block">
			<legend class="modal__title modal__title_small">Include</legend>
			<p class="modal__confirmation-msg">This attachment is an image. To include it n a hypha, use a syntax like this:</p>
			<pre class="codebleck"><code>img { `)
//line views/readers.qtpl:41
			qw422016.E().S(h.Name)
//line views/readers.qtpl:41
			qw422016.N().S(` }</code></pre>
		</fieldset>
		`)
//line views/readers.qtpl:43
		}
//line views/readers.qtpl:43
		qw422016.N().S(`
	`)
//line views/readers.qtpl:44
	}
//line views/readers.qtpl:44
	qw422016.N().S(`

	`)
//line views/readers.qtpl:46
	if u.CanProceed("upload-binary") {
//line views/readers.qtpl:46
		qw422016.N().S(`
	<form action="/upload-binary/`)
//line views/readers.qtpl:47
		qw422016.E().S(h.Name)
//line views/readers.qtpl:47
		qw422016.N().S(`"
			method="post" enctype="multipart/form-data"
			class="upload-binary modal amnt-menu-block">
		<fieldset class="modal__fieldset">
			<legend class="modal__title modal__title_small">Attach</legend>
			<p class="modal__confirmation-msg">You can upload a new attachment. Please do not upload too big pictures unless you need to because may not want to wait for big pictures to load.</p>
			<label for="upload-binary__input"></label>
			<input type="file" id="upload-binary__input" name="binary">

			<input type="submit" class="btn stick-to-bottom" value="Upload">
		</fieldset>
	</form>
	`)
//line views/readers.qtpl:59
	}
//line views/readers.qtpl:59
	qw422016.N().S(`

	`)
//line views/readers.qtpl:61
	if h.BinaryPath != "" && u.CanProceed("unattach-confirm") {
//line views/readers.qtpl:61
		qw422016.N().S(`
	<form action="/unattach-confirm/`)
//line views/readers.qtpl:62
		qw422016.E().S(h.Name)
//line views/readers.qtpl:62
		qw422016.N().S(`" method="post" class="modal amnt-menu-block">
		<fieldset class="modal__fieldset">
			<legend class="modal__title modal__title_small">Unattach</legend>
			<p class="modal__confirmation-msg">Please note that you don't have to unattach before uploading a new attachment.</p>
			<input type="submit" class="btn" value="Unattach">
		</fieldset>
	</form>
	`)
//line views/readers.qtpl:69
	}
//line views/readers.qtpl:69
	qw422016.N().S(`

	</section>
</main>
</div>
`)
//line views/readers.qtpl:74
}

//line views/readers.qtpl:74
func WriteAttachmentMenuHTML(qq422016 qtio422016.Writer, rq *http.Request, h *hyphae.Hypha, u *user.User) {
//line views/readers.qtpl:74
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/readers.qtpl:74
	StreamAttachmentMenuHTML(qw422016, rq, h, u)
//line views/readers.qtpl:74
	qt422016.ReleaseWriter(qw422016)
//line views/readers.qtpl:74
}

//line views/readers.qtpl:74
func AttachmentMenuHTML(rq *http.Request, h *hyphae.Hypha, u *user.User) string {
//line views/readers.qtpl:74
	qb422016 := qt422016.AcquireByteBuffer()
//line views/readers.qtpl:74
	WriteAttachmentMenuHTML(qb422016, rq, h, u)
//line views/readers.qtpl:74
	qs422016 := string(qb422016.B)
//line views/readers.qtpl:74
	qt422016.ReleaseByteBuffer(qb422016)
//line views/readers.qtpl:74
	return qs422016
//line views/readers.qtpl:74
}

// If `contents` == "", a helpful message is shown instead.

//line views/readers.qtpl:77
func StreamHyphaHTML(qw422016 *qt422016.Writer, rq *http.Request, h *hyphae.Hypha, contents string) {
//line views/readers.qtpl:77
	qw422016.N().S(`
`)
//line views/readers.qtpl:79
	siblings, subhyphae, prevHyphaName, nextHyphaName := tree.Tree(h.Name)
	u := user.FromRequest(rq)

//line views/readers.qtpl:81
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<article id="hypha">
		<div class="jump-btn">
			<a class="jump-btn__link" href="#hypha-bottom">↓</a>
		</div>
		<div class="btn edit-btn">
			<a class="edit-btn__link" href="/edit/`)
//line views/readers.qtpl:89
	qw422016.E().S(h.Name)
//line views/readers.qtpl:89
	qw422016.N().S(`">Edit text</a>
		</div>
		`)
//line views/readers.qtpl:91
	qw422016.N().S(NaviTitleHTML(h))
//line views/readers.qtpl:91
	qw422016.N().S(`
		`)
//line views/readers.qtpl:92
	if h.Exists {
//line views/readers.qtpl:92
		qw422016.N().S(`
			`)
//line views/readers.qtpl:93
		qw422016.N().S(contents)
//line views/readers.qtpl:93
		qw422016.N().S(`
		`)
//line views/readers.qtpl:94
	} else {
//line views/readers.qtpl:94
		qw422016.N().S(`
		    `)
//line views/readers.qtpl:95
		streamnonExistentHyphaNotice(qw422016, h, u)
//line views/readers.qtpl:95
		qw422016.N().S(`
		`)
//line views/readers.qtpl:96
	}
//line views/readers.qtpl:96
	qw422016.N().S(`
	</article>
	<section class="prevnext">
		`)
//line views/readers.qtpl:99
	if prevHyphaName != "" {
//line views/readers.qtpl:99
		qw422016.N().S(`
		<a class="prevnext__el prevnext__prev" href="/hypha/`)
//line views/readers.qtpl:100
		qw422016.E().S(prevHyphaName)
//line views/readers.qtpl:100
		qw422016.N().S(`" rel="prev">← `)
//line views/readers.qtpl:100
		qw422016.E().S(util.BeautifulName(path.Base(prevHyphaName)))
//line views/readers.qtpl:100
		qw422016.N().S(`</a>
		`)
//line views/readers.qtpl:101
	}
//line views/readers.qtpl:101
	qw422016.N().S(`
		`)
//line views/readers.qtpl:102
	if nextHyphaName != "" {
//line views/readers.qtpl:102
		qw422016.N().S(`
		<a class="prevnext__el prevnext__next" href="/hypha/`)
//line views/readers.qtpl:103
		qw422016.E().S(nextHyphaName)
//line views/readers.qtpl:103
		qw422016.N().S(`" rel="next">`)
//line views/readers.qtpl:103
		qw422016.E().S(util.BeautifulName(path.Base(nextHyphaName)))
//line views/readers.qtpl:103
		qw422016.N().S(` →</a>
		`)
//line views/readers.qtpl:104
	}
//line views/readers.qtpl:104
	qw422016.N().S(`
	</section>
`)
//line views/readers.qtpl:106
	StreamSubhyphaeHTML(qw422016, subhyphae)
//line views/readers.qtpl:106
	qw422016.N().S(`
	<section id="hypha-bottom">
		<div class="jump-btn">
    		<a class="jump-btn__link" href="#hypha">↑</a>
   		</div>
   		`)
//line views/readers.qtpl:111
	streamhyphaInfo(qw422016, rq, h)
//line views/readers.qtpl:111
	qw422016.N().S(`
	</section>
</main>
`)
//line views/readers.qtpl:114
	streamsiblingHyphaeHTML(qw422016, siblings)
//line views/readers.qtpl:114
	qw422016.N().S(`
</div>
`)
//line views/readers.qtpl:116
	streamviewScripts(qw422016)
//line views/readers.qtpl:116
	qw422016.N().S(`
`)
//line views/readers.qtpl:117
}

//line views/readers.qtpl:117
func WriteHyphaHTML(qq422016 qtio422016.Writer, rq *http.Request, h *hyphae.Hypha, contents string) {
//line views/readers.qtpl:117
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/readers.qtpl:117
	StreamHyphaHTML(qw422016, rq, h, contents)
//line views/readers.qtpl:117
	qt422016.ReleaseWriter(qw422016)
//line views/readers.qtpl:117
}

//line views/readers.qtpl:117
func HyphaHTML(rq *http.Request, h *hyphae.Hypha, contents string) string {
//line views/readers.qtpl:117
	qb422016 := qt422016.AcquireByteBuffer()
//line views/readers.qtpl:117
	WriteHyphaHTML(qb422016, rq, h, contents)
//line views/readers.qtpl:117
	qs422016 := string(qb422016.B)
//line views/readers.qtpl:117
	qt422016.ReleaseByteBuffer(qb422016)
//line views/readers.qtpl:117
	return qs422016
//line views/readers.qtpl:117
}

//line views/readers.qtpl:119
func StreamRevisionHTML(qw422016 *qt422016.Writer, rq *http.Request, h *hyphae.Hypha, contents, revHash string) {
//line views/readers.qtpl:119
	qw422016.N().S(`
`)
//line views/readers.qtpl:121
	siblings, subhyphae, _, _ := tree.Tree(h.Name)

//line views/readers.qtpl:122
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<article>
		<p>Please note that viewing attachments of hyphae is not supported in history for now.</p>
		`)
//line views/readers.qtpl:127
	qw422016.N().S(NaviTitleHTML(h))
//line views/readers.qtpl:127
	qw422016.N().S(`
		`)
//line views/readers.qtpl:128
	qw422016.N().S(contents)
//line views/readers.qtpl:128
	qw422016.N().S(`
	</article>
`)
//line views/readers.qtpl:130
	StreamSubhyphaeHTML(qw422016, subhyphae)
//line views/readers.qtpl:130
	qw422016.N().S(`
</main>
`)
//line views/readers.qtpl:132
	streamsiblingHyphaeHTML(qw422016, siblings)
//line views/readers.qtpl:132
	qw422016.N().S(`
</div>
`)
//line views/readers.qtpl:134
	streamviewScripts(qw422016)
//line views/readers.qtpl:134
	qw422016.N().S(`
`)
//line views/readers.qtpl:135
}

//line views/readers.qtpl:135
func WriteRevisionHTML(qq422016 qtio422016.Writer, rq *http.Request, h *hyphae.Hypha, contents, revHash string) {
//line views/readers.qtpl:135
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/readers.qtpl:135
	StreamRevisionHTML(qw422016, rq, h, contents, revHash)
//line views/readers.qtpl:135
	qt422016.ReleaseWriter(qw422016)
//line views/readers.qtpl:135
}

//line views/readers.qtpl:135
func RevisionHTML(rq *http.Request, h *hyphae.Hypha, contents, revHash string) string {
//line views/readers.qtpl:135
	qb422016 := qt422016.AcquireByteBuffer()
//line views/readers.qtpl:135
	WriteRevisionHTML(qb422016, rq, h, contents, revHash)
//line views/readers.qtpl:135
	qs422016 := string(qb422016.B)
//line views/readers.qtpl:135
	qt422016.ReleaseByteBuffer(qb422016)
//line views/readers.qtpl:135
	return qs422016
//line views/readers.qtpl:135
}

//line views/readers.qtpl:137
func streamviewScripts(qw422016 *qt422016.Writer) {
//line views/readers.qtpl:137
	qw422016.N().S(`
`)
//line views/readers.qtpl:138
	for _, scriptPath := range cfg.ViewScripts {
//line views/readers.qtpl:138
		qw422016.N().S(`
<script src="`)
//line views/readers.qtpl:139
		qw422016.E().S(scriptPath)
//line views/readers.qtpl:139
		qw422016.N().S(`"></script>
`)
//line views/readers.qtpl:140
	}
//line views/readers.qtpl:140
	qw422016.N().S(`
`)
//line views/readers.qtpl:141
}

//line views/readers.qtpl:141
func writeviewScripts(qq422016 qtio422016.Writer) {
//line views/readers.qtpl:141
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/readers.qtpl:141
	streamviewScripts(qw422016)
//line views/readers.qtpl:141
	qt422016.ReleaseWriter(qw422016)
//line views/readers.qtpl:141
}

//line views/readers.qtpl:141
func viewScripts() string {
//line views/readers.qtpl:141
	qb422016 := qt422016.AcquireByteBuffer()
//line views/readers.qtpl:141
	writeviewScripts(qb422016)
//line views/readers.qtpl:141
	qs422016 := string(qb422016.B)
//line views/readers.qtpl:141
	qt422016.ReleaseByteBuffer(qb422016)
//line views/readers.qtpl:141
	return qs422016
//line views/readers.qtpl:141
}
