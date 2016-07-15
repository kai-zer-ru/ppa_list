package main

import (
	"flag"
	"fmt"
	"time"
	"net/http"
	"strings"
	"net"
	"os"
)

var (

	MainConfig      = Config {}
	PpaList         = make([]string,0)
	PpaListString   string
	SoftList        = make([]string,0)
	SoftListString   string
	Version         = "0.01"

)

func parseFlags() bool {
	var info 			= flag.Bool("info", false, "")
	var hangup 			= flag.Bool("hangup", false, "")
	var interrupt		= flag.Bool("interrupt", false, "")
	flag.Parse()
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


func ReadPpaList() error {
	PpaListPath := MainConfig.GetConfString("PpaListPath", "/opt/ppalist/ppalist")
	PpaListFile, err := os.Open(PpaListPath)
	if err != nil {
		fmt.Println("Erro while loading PPA list: ", err.Error())
		return err
	}
	defer PpaListFile.Close()

	stat, err := PpaListFile.Stat()
	if err != nil {
		fmt.Println("Erro while get PPA list size: ", err.Error())
		return err
	}
	bs := make([]byte, stat.Size())
	_, err = PpaListFile.Read(bs)
	if err != nil {
		return err
	}

	PpaListString = string(bs)
	PpaList = strings.Split(PpaListString, ";")
	return nil
}

func ReadSoftList() error {
	SoftListPath := MainConfig.GetConfString("SoftListPath", "/opt/ppalist/softlist")
	SoftListFile, err := os.Open(SoftListPath)
	if err != nil {
		fmt.Println("Erro while loading Soft list: ", err.Error())
		return err
	}
	defer SoftListFile.Close()

	stat, err := SoftListFile.Stat()
	if err != nil {
		fmt.Println("Erro while get Soft list size: ", err.Error())
		return err
	}
	bs := make([]byte, stat.Size())
	_, err = SoftListFile.Read(bs)
	if err != nil {
		return err
	}

	SoftListString = string(bs)
	SoftList = strings.Split(SoftListString, " ")
	return nil
}

func main() {
	resultParse := parseFlags()
	if (resultParse == true) {
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
	http.HandleFunc("/", application)
	http.HandleFunc("/add_repo", add_repo)
	is_unix := MainConfig.GetConfBool("UseUnix",false)
	if is_unix ==  false{
		address := MainConfig.GetConfString("RunPath","localhost:3333")
		fmt.Println(address)
		panic(http.ListenAndServe(address,nil))
	} else {
		unix_sock := MainConfig.GetConfString("RunPath","/var/run/ppa-list.sock")
		os.Remove(unix_sock)
		l,err := net.Listen("unix",unix_sock)
		os.Chmod(unix_sock, 0777)
		if err != nil {
			panic(fmt.Sprintf("Listen socket %s error: %s", unix_sock,err))
		}
		panic(http.Serve(l,nil))
	}
}

func application(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Start Application")
	content_type := MainConfig.GetConfString("ContentType","text/html")
	w.Header().Set("Content-Type",content_type)

	lang := req.FormValue("lang")
	if (lang == "") {
		lang = MainConfig.GetConfString("DefaultLanguage", RUSSIAN_LANG)
	}
	PageBody := GetMainPage(lang)
	fmt.Fprintf(w,"%s",PageBody)
	fmt.Println("End Application")
}

func add_repo(w http.ResponseWriter, req *http.Request) {
	content_type := MainConfig.GetConfString("ContentType","text/html")
	w.Header().Set("Content-Type",content_type)

	lang := req.FormValue("lang")
	if (lang == "") {
		lang = MainConfig.GetConfString("DefaultLanguage", RUSSIAN_LANG)
	}
	PageBody := GetAddRepoPage(lang)
	fmt.Fprintf(w,"%s",PageBody)
}
