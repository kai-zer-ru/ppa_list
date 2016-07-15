package main

import (
	"fmt"
	"strings"
	"os"
)

var (
	template = "template"
	LANG_SELECT_TEXT = map[string]string {
		RUSSIAN_LANG: "Сменить язык",
		ENGLISH_LANG : "Change language",
	}
	TITLES = map[string]string {
		RUSSIAN_LANG: "Добавление репозитория",
		ENGLISH_LANG : "Add repository",
	}
)

func Welcome(lang string) string {
	WelcomeText := ""
	switch lang {
	case RUSSIAN_LANG:
		WelcomeText = "Добро пожаловать в PPA list manager версии "
	case ENGLISH_LANG:
		WelcomeText = "Welcome to PPA liast mamger v."
	default:
		WelcomeText = "Добро пожаловать в PPA list manager версии "
	}
	return WelcomeText
}

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
	Page = strings.Replace(Page,"PAGE_TITLE", title,1)
	Page = strings.Replace(Page,"CONTENT", "%s",1)
	Page = strings.Replace(Page,"LANG_SELECT_TEXT", LANG_SELECT_TEXT[lang],1)
	return Page
}

func PrintHTML(content, title, lang string) string {
	template := LoadTemplate(lang, title)
	HTML := fmt.Sprintf(template,content)
	return HTML
}


func GetMainPage(lang string) string {
	return PrintHTML(fmt.Sprintf(LoadPage("main", lang),  PpaListString, SoftListString, strings.Join(PpaList, "; sudo add-apt-repository"), strings.Join(SoftList, " ")), fmt.Sprintf("%s %s!", Welcome(lang), Version), lang)
}

func GetAddRepoPage(lang string) string {
	return PrintHTML(LoadPage("add_repo", lang), TITLES[lang], lang)
}