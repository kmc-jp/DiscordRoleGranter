package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/bwmarrin/discordgo"
)

type Serve struct{}

// Header put html header data
type Header struct {
	UserName string
	Error    string
	Success  string
}

// Footer put html footer data
type Footer struct{}

// Index put html index page data
type Index struct {
	Header Header
	Footer Footer
}

func (s *Serve) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Query().Get("Action") {
	case "AddRole":
		s.addRole(w, r)
	case "List":
		s.memberList(w, r)
	default:
		s.main(w, r)
	}
}

func (s *Serve) main(w http.ResponseWriter, r *http.Request) {
	var index = Index{
		Header: Header{
			UserName: os.Getenv("REMOTE_USER"),
		},
	}

	var t = template.New("index")
	t, err := t.ParseFiles(
		filepath.Join("resources", "index.html"),
		filepath.Join("resources", "header.html"),
		filepath.Join("resources", "footer.html"),
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Parse Error<br>%s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "text/html")

	err = t.ExecuteTemplate(w, "index", index)
	if err != nil {
		fmt.Fprintf(w, "Excute Error!<br>%s", err.Error())
		return
	}
	return
}

func (s *Serve) memberList(w http.ResponseWriter, r *http.Request) {
	type Member struct {
		UserName string
		Nick     string
		ID       string
	}
	type Result struct {
		Status string
		Member []Member
	}
	ses, err := NewDiscord()
	if err != nil {
		SendJSON(w, Result{err.Error(), nil}, http.StatusInternalServerError)
		return
	}

	var members = make([]Member, 0)

	var mems = make([]*discordgo.Member, 0)
	{
		var after string = "0"
	add:
		mem, err := ses.GuildMembers(Settings.Discord.GuildID, after, 1000)
		if err != nil {
			SendJSON(w, Result{err.Error(), nil}, http.StatusInternalServerError)
			return
		}
		mems = append(mems, mem...)
		if len(mem) > 1000 {
			after = mem[999].User.ID
			goto add
		}
	}

	for _, mem := range mems {
		members = append(members,
			Member{
				UserName: mem.User.Username,
				Nick:     mem.Nick,
				ID:       mem.User.ID,
			})
	}

	SendJSON(w, Result{"OK", members}, http.StatusOK)
	return
}

func (s *Serve) addRole(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		Status string
	}

	ses, err := NewDiscord()
	if err != nil {
		SendJSON(w, Result{err.Error()}, http.StatusInternalServerError)
		return
	}

	err = r.ParseForm()
	if err != nil {
		SendJSON(w, Result{err.Error()}, http.StatusBadRequest)
		return
	}

	err = ses.GuildMemberRoleAdd(
		Settings.Discord.GuildID,
		r.FormValue("user_id"),
		Settings.Discord.RoleID,
	)

	if err != nil {
		SendJSON(w, Result{err.Error()}, http.StatusInternalServerError)
		return
	}

	SendJSON(w, Result{"OK"}, http.StatusOK)
	return
}
