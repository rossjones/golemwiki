package main

import (
	//"fmt"
	"github.com/flosch/pongo"
	"net/http"
	"path"
)

var templates = make(map[string]*pongo.Template)

func init_templates() {
	root := config.Server["templates"]
	templates["index.html"] = pongo.Must(pongo.FromFile(path.Join(root, "index.html"), nil))
	templates["about.html"] = pongo.Must(pongo.FromFile(path.Join(root, "about.html"), nil))
	templates["history.html"] = pongo.Must(pongo.FromFile(path.Join(root, "history.html"), nil))
	templates["add/page.html"] = pongo.Must(pongo.FromFile(path.Join(root, "add/page.html"), nil))
	templates["add/group.html"] = pongo.Must(pongo.FromFile(path.Join(root, "add/group.html"), nil))
}

func renderTemplate(tmpl string, ctx pongo.Context, w http.ResponseWriter) {
	//err := templates[tmpl].ExecuteRW(w, &ctx)
	t, err := pongo.FromFile(path.Join(config.Server["templates"], tmpl), nil)
	//err := templates[tmpl].ExecuteRW(w, &ctx)

	// Add some general items to context from the config
	for k, v := range config.Wiki {
		ctx["_"+k] = v
	}

	ctx["_groups"] = get_all_groups()
	err = t.ExecuteRW(w, &ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homepage_view(w http.ResponseWriter, r *http.Request) {
	renderTemplate("index.html", pongo.Context{"active": "home"}, w)
}

func about_view(w http.ResponseWriter, r *http.Request) {
	renderTemplate("about.html", pongo.Context{"active": "about"}, w)
}

func history_view(w http.ResponseWriter, r *http.Request) {
	renderTemplate("history.html", pongo.Context{"active": "history"}, w)
}

func admin_view(w http.ResponseWriter, r *http.Request) {
	renderTemplate("admin.html", pongo.Context{"active": "admin"}, w)
}

func add_page_view(w http.ResponseWriter, r *http.Request) {
	renderTemplate("add/page.html", pongo.Context{"active": "add_page"}, w)
}

func add_group_view(w http.ResponseWriter, r *http.Request) {
	renderTemplate("add/group.html", pongo.Context{"active": "add_group"}, w)
}
