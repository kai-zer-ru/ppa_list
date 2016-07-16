package main

import (
	"fmt"
	"strings"
)

func AddRepoToSourceList(repo_path string) {
	SourceList = append(SourceList, repo_path)
	if (SourceListString == "") {
		SourceListString = fmt.Sprint("%s; " ,repo_path)
	} else {
		SoftListString = fmt.Sprintf("%s; %s", SourceListString, repo_path)
	}
}

func AddRepoToPpaList(repo_path string) {
	PpaList = append(PpaList, repo_path)
	if (PpaListString == "") {
		PpaListString = fmt.Sprint("%s; " ,repo_path)
	} else {
		PpaListString = fmt.Sprintf("%s; %s", PpaListString, repo_path)
	}
}

func AddSoft(soft string) {
	soft_list := strings.Split(soft, " ")
	if (len(soft_list) == 0) {
		return
	}
	if (len(soft_list) == 1) {
		SoftList = append(SoftList, strings.Trim(soft, " "))
	} else {
		for _, soft_name := range soft_list {
			SoftList = append(SoftList, strings.Trim(soft_name, " "))
		}
	}

	if (SoftListString == "") {
		SoftListString = strings.Replace(soft, " ", ";", len(soft_list))
		SoftListString = fmt.Sprint("%s; " ,SoftListString)
	} else {
		soft_string := strings.Replace(soft, " ", ";", len(soft_list))
		SoftListString = fmt.Sprintf("%s %s;", SourceListString, soft_string)
	}
}