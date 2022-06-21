package interwiki

import (
	"embed"
	"github.com/bouncepaw/mycorrhiza/viewutil"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

var (
	//go:embed *html
	fs             embed.FS
	ruTranslation  = ``
	chainInterwiki viewutil.Chain
	chainNameTaken viewutil.Chain
)

func InitHandlers(rtr *mux.Router) {
	chainInterwiki = viewutil.CopyEnRuWith(fs, "view_interwiki.html", ruTranslation)
	chainNameTaken = viewutil.CopyEnRuWith(fs, "view_name_taken.html", ruTranslation)
	rtr.HandleFunc("/interwiki", handlerInterwiki)
	rtr.HandleFunc("/interwiki/add-entry", handlerAddEntry).Methods(http.MethodPost)
}

func handlerAddEntry(w http.ResponseWriter, rq *http.Request) {
	wiki := Wiki{
		Name:           rq.PostFormValue("name"),
		Aliases:        strings.Split(rq.PostFormValue("aliases"), ","),
		URL:            rq.PostFormValue("url"),
		LinkHrefFormat: rq.PostFormValue("link-href-format"),
		ImgSrcFormat:   rq.PostFormValue("img-src-format"),
		Engine:         WikiEngine(rq.PostFormValue("engine")),
	}
	wiki.canonize()
	if err := addEntry(&wiki); err != nil {
		viewNameTaken(viewutil.MetaFrom(w, rq), &wiki, err.Error())
		return
	}
	saveInterwikiJson()
	http.Redirect(w, rq, "/interwiki", http.StatusSeeOther)
}

type nameTakenData struct {
	*viewutil.BaseData
	*Wiki
	TakenName string
}

func viewNameTaken(meta viewutil.Meta, wiki *Wiki, takenName string) {
	viewutil.ExecutePage(meta, chainNameTaken, nameTakenData{
		BaseData:  &viewutil.BaseData{},
		Wiki:      wiki,
		TakenName: takenName,
	})
}

func handlerInterwiki(w http.ResponseWriter, rq *http.Request) {
	viewInterwiki(viewutil.MetaFrom(w, rq))
}

type interwikiData struct {
	*viewutil.BaseData
	Entries []*Wiki
	CanEdit bool
	Error   string
}

func viewInterwiki(meta viewutil.Meta) {
	viewutil.ExecutePage(meta, chainInterwiki, interwikiData{
		BaseData: &viewutil.BaseData{},
		Entries:  listOfEntries,
		CanEdit:  meta.U.Group == "admin",
		Error:    "",
	})
}
