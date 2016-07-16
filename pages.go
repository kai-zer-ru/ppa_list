package main

import (
	"fmt"
	"strings"
	"os"
)

var (
	template = "template"
	OTHER_LANG = map[string]string {
		RUSSIAN_LANG: fmt.Sprintf("<a href=\"?lang=%s\"><img id=\"lang\" src=\"/static/images/en.ico\"></a>",ENGLISH_LANG),
		ENGLISH_LANG: fmt.Sprintf("<a href=\"?lang=%s\"><img id=\"lang\" src=\"/static/images/ru.ico\"></a>",RUSSIAN_LANG),
	}

	PAGES = map[string]map[string]string {
		"main": map[string]string {
			RUSSIAN_LANG: "Главная",
			ENGLISH_LANG: "Main",
		},
		"add_repo": map[string]string {
			RUSSIAN_LANG: "Добавить репозиторий",
			ENGLISH_LANG: "Add repository",
		},
		"repo_list": map[string]string {
			RUSSIAN_LANG: "Ваши репозитории",
			ENGLISH_LANG: "Your repositories",
		},
		"about": map[string]string {
			RUSSIAN_LANG: "О сервисе",
			ENGLISH_LANG: "About service",
		},
		"contacts": map[string]string {
			RUSSIAN_LANG: "Контакты",
			ENGLISH_LANG: "Contacts",
		},
	}

	TITLES = map[string]map[string]string {
		"main": map[string]string{
			RUSSIAN_LANG: "Добро пожаловать в PPA list manager версии ",
			ENGLISH_LANG : "Welcome to PPA liast mamger v.",
		},
		"add_repo": map[string]string{
			RUSSIAN_LANG: "Добавление репозитория",
			ENGLISH_LANG : "Add repository",
		},
		"repo_list": map[string]string{
			RUSSIAN_LANG: "Ваши репозитории",
			ENGLISH_LANG : "Your repositories",
		},
		"about": map[string]string{
			RUSSIAN_LANG: "О сервисе",
			ENGLISH_LANG : "About service",
		},
		"contacts": map[string]string{
			RUSSIAN_LANG: "Контакты",
			ENGLISH_LANG : "Contacts",
		},
	}
	ERRORS = map[int]map[string]string {
		ERROR_REPO_NOT_ADDED : map[string]string {
			RUSSIAN_LANG: "Репозиторий не добавлен",
			ENGLISH_LANG : "Repository not added",
		},
	}
)

func LoadPage(name, lang string) string {
	pages_path := "/opt/ppalist/pages"
	PageFile, err := os.Open(fmt.Sprintf("%s/%s/%s.html", pages_path, lang, name))
	if err != nil {
		fmt.Println("Erro while loading Soft list: ", err.Error())
		return ""
	}
	defer PageFile.Close()

	stat, err := PageFile.Stat()
	if err != nil {
		fmt.Println("Erro while get Soft list size: ", err.Error())
		return ""
	}
	bs := make([]byte, stat.Size())
	_, err = PageFile.Read(bs)
	if err != nil {
		return ""
	}

	Page := string(bs)
	return Page
}

func LoadTemplate(lang, title string) string {
	template_path := fmt.Sprintf("/opt/ppalist/pages/templates/%s.html",template)
	PageFile, err := os.Open(template_path)
	if err != nil {
		fmt.Println("Erro while loading Soft list: ", err.Error())
		return ""
	}
	defer PageFile.Close()

	stat, err := PageFile.Stat()
	if err != nil {
		fmt.Println("Erro while get Soft list size: ", err.Error())
		return ""
	}
	bs := make([]byte, stat.Size())
	_, err = PageFile.Read(bs)
	if err != nil {
		return ""
	}

	Page := string(bs)
	Page = strings.Replace(Page,"/PAGE_TITLE/", title,1)
	Page = strings.Replace(Page,"/CONTENT/", "%s",1)
	Page = strings.Replace(Page,"/LANG_Q/", lang,4)
	Page = strings.Replace(Page,"/LANG/", OTHER_LANG[lang],1)
	Page = strings.Replace(Page, "/HOME/", PAGES["main"][lang], 1)
	Page = strings.Replace(Page, "/ADD_REPO/", PAGES["add_repo"][lang], 1)
	Page = strings.Replace(Page, "/REPO_LIST/", PAGES["repo_list"][lang], 1)
	Page = strings.Replace(Page, "/ABOUT/", PAGES["about"][lang], 1)
	Page = strings.Replace(Page, "/CONTACTS/", PAGES["contacts"][lang], 1)
	return Page
}

func PrintHTML(content, title, lang string, to_replace map[string]string) string {
	template := LoadTemplate(lang, title)
	HTML := fmt.Sprintf(template, content)
	if (len(to_replace) > 0) {
		for text, replaced := range to_replace {
			HTML = strings.Replace(HTML, text, replaced,1)
		}
	}
	return HTML
}

func GetMainPage(lang string) string {
	to_replace := map[string]string {}
	content := LoadPage("main", lang)
	return PrintHTML(content, fmt.Sprintf("%s %s!", TITLES["main"][lang], Version), lang, to_replace)
}

func GetReposPage(lang string) string {
	fmt.Println(PpaListString)
	fmt.Println(PpaList)
	to_replace := map[string]string {}
	content := fmt.Sprintf(LoadPage("repo_list", lang),  strings.Join(strings.Split(PpaListString,";"),"</br>"), SoftListString, strings.Join(SourceList, "</br>"),strings.Join(PpaList, "</br> sudo add-apt-repository "), strings.Join(SoftList, " "))
	return PrintHTML(content, TITLES["repo_list"][lang], lang, to_replace)
}

func GetAddRepoPage(lang, error_msg string) string {
	to_replace := map[string]string {
		"ERROR": error_msg,
	}
	content := LoadPage("add_repo", lang)
	return PrintHTML(content, TITLES["add_repo"][lang], lang, to_replace)
}

func GetContactsPage(lang string) string {
	to_replace := map[string]string {}
	content := LoadPage("contacts", lang)
	return PrintHTML(content, TITLES["about"][lang], lang, to_replace)
}