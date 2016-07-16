package main

import (
	"flag"
	"fmt"
	"time"
	"net/http"
	"strings"
	"strconv"
)

var (

	MainConfig          = Config {}
	PpaList             = make([]string,0)
	PpaMap              = map[string]int {}
	PpaListString       string
	SourceList          = make([]string,0)
	SourceMap              = map[string]int {}
	SourceListString    string
	SoftList            = make([]string,0)
	SoftMap              = map[string]int {}
	SoftListString      string
	Version             = "0.01"
	chttp               = http.NewServeMux()
)

func parseFlags() bool {
	var info 			= flag.Bool("info", false, "")
	var hangup 			= flag.Bool("hangup", false, "")
	var interrupt		= flag.Bool("interrupt", false, "")
	var conf            = flag.String("conf", "/opt/ppalist/main.cfg", "")
	flag.Parse()
	if *conf != "" {
		MainConfig.Configuration_filename = *conf
	}
	if *info {
		a := GetPid()
		if len(a) > 0 {
			fmt.Println("Started:")
			for _, p := range a {
				fmt.Println(fmt.Sprintf("%d\t%s\n", p.Pid, p.Path))
			}
		} else {
			fmt.Println(fmt.Sprintf("Started not found.\n"))
		}
		return true
	}

	if *interrupt {
		// Берем список PID запущенных процессов
		a := GetPid()
		if len(a) > 0 {
			// Отправляем HUP для каждого процесса kill -1 {PID}
			fmt.Println("Send HUP...")
			for _, p := range a {
				fmt.Println(fmt.Sprintf("%d", p.Pid))
				// Отправляем сигнал на перезапуск
				ExecCmd("kill", []string{"-2", fmt.Sprintf("%d", p.Pid)})
			}

			for {
				<-time.After(1 * time.Second)
				a := GetPid()
				if len(a) == 0 {
					fmt.Println("All process stoped.\n")
					return true
				} else {
					fmt.Println("Kill in process...\nWorked count = %v\n", len(a))
				}
			}
		} else {
			fmt.Println("Started not found.\n")
		}
		return true
	}
	if *hangup {
		// Берем список PID запущенных процессов
		a := GetPid()
		if len(a) > 0 {
			// Отправляем HUP для каждого процесса kill -1 {PID}
			fmt.Println("Send HUP...")
			for _, p := range a {
				fmt.Println(fmt.Sprintf("%d", p.Pid))
				// Отправляем сигнал на перезапуск
				ExecCmd("kill", []string{"-1", fmt.Sprintf("%d", p.Pid)})
			}

			for {
				<-time.After(1 * time.Second)
				a := GetPid()
				if len(a) == 0 {
					fmt.Println("All process stoped.\n")
					break
				} else {
					fmt.Println("Kill in process... Worked count = %v\n", len(a))
				}
			}
		} else {
			fmt.Println("Started not found.\n")
		}
		return true
	}
	return false
}



func main() {
	MainConfig.Configuration_filename = "/opt/ppalist/main.cfg"
	resultParse := parseFlags()
	if (resultParse == true) {
		return
	}
	MainConfig.Configuration = map[string]string{}

	if (!MainConfig.ReadConfiguration()) {
		fmt.Println("Read Configuration error")
		return
	}

	err := ReadPpaList()
	if (err != nil) {
		fmt.Println(err.Error())
		return
	}
	err = ReadSoftList()
	if (err != nil) {
		fmt.Println(err.Error())
		return
	}
	err = ReadSourceList()
	if (err != nil) {
		fmt.Println(err.Error())
		return
	}
	http.HandleFunc("/", application)
	http.HandleFunc("/add_repo", add_repo)
	http.HandleFunc("/add_new_repo", add_new_repo)
	http.HandleFunc("/repo_list", repo_list)
	http.HandleFunc("/contacts", contacts)
	chttp.Handle("/", http.FileServer(http.Dir(MainConfig.GetConfString("StaticDir","/opt/ppalist/"))))
	address := MainConfig.GetConfString("RunPath","localhost:3333")
	fmt.Println(address)
	panic(http.ListenAndServe(address,nil))
}

func application(w http.ResponseWriter, req *http.Request) {
	if (strings.Contains(req.URL.Path, ".")) {
		chttp.ServeHTTP(w, req)
		return
	}
	content_type := MainConfig.GetConfString("ContentType","text/html")
	w.Header().Set("Content-Type",content_type)

	lang := req.FormValue("lang")
	if (lang == "") {
		lang = MainConfig.GetConfString("DefaultLanguage", RUSSIAN_LANG)
	}
	PageBody := GetMainPage(lang)
	fmt.Fprintf(w,"%s",PageBody)
}

func contacts(w http.ResponseWriter, req *http.Request) {
	content_type := MainConfig.GetConfString("ContentType","text/html")
	w.Header().Set("Content-Type",content_type)

	lang := req.FormValue("lang")
	if (lang == "") {
		lang = MainConfig.GetConfString("DefaultLanguage", RUSSIAN_LANG)
	}
	PageBody := GetContactsPage(lang)
	fmt.Fprintf(w,"%s",PageBody)
}

func repo_list(w http.ResponseWriter, req *http.Request) {
	content_type := MainConfig.GetConfString("ContentType","text/html")
	w.Header().Set("Content-Type",content_type)

	lang := req.FormValue("lang")
	if (lang == "") {
		lang = MainConfig.GetConfString("DefaultLanguage", RUSSIAN_LANG)
	}
	PageBody := GetReposPage(lang)
	fmt.Fprintf(w,"%s",PageBody)
}

func add_repo(w http.ResponseWriter, req *http.Request) {
	content_type := MainConfig.GetConfString("ContentType","text/html")
	w.Header().Set("Content-Type",content_type)

	lang := req.FormValue("lang")
	if (lang == "") {
		lang = MainConfig.GetConfString("DefaultLanguage", RUSSIAN_LANG)
	}
	error_code_str := req.FormValue("error")
	PageBody := ""
	if (error_code_str != "") {
		error_code,_ := strconv.Atoi(error_code_str)
		PageBody = GetAddRepoPage(lang, ERRORS[error_code][lang])
		fmt.Fprintf(w,"%s",PageBody)
		return
	}
	PageBody = GetAddRepoPage(lang, "")
	fmt.Fprintf(w,"%s",PageBody)
}

func add_new_repo(w http.ResponseWriter, req *http.Request) {
	content_type := MainConfig.GetConfString("ContentType","text/html")
	w.Header().Set("Content-Type",content_type)

	lang := req.FormValue("lang")
	if (lang == "") {
		lang = MainConfig.GetConfString("DefaultLanguage", RUSSIAN_LANG)
	}
	repo_path := req.FormValue("repo_path")
	soft := req.FormValue("soft")

	PageBody := ""
	if (soft == "" || repo_path == "") {
		http.Redirect(w, req, fmt.Sprintf("/add_repo?error=%v", ERROR_REPO_NOT_ADDED), 301)
		return
	}

	if (repo_path[0:3] == "deb") {
		// adding to sourcelist
		AddRepoToSourceList(repo_path)
		AddSoft(soft)
	} else if (repo_path[0:3] == "ppa") {
		// adding to ppalist
		AddRepoToPpaList(repo_path)
		AddSoft(soft)
	}

	PageBody = GetAddRepoPage(lang, REPO_IS_ADDED[lang])
	fmt.Fprintf(w,"%s",PageBody)
}
