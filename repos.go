package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

func AddRepoToSourceList(repo_path string) {
	repo_path = strings.Trim(repo_path, " ")
	if _, ok := SourceMap[repo_path]; ok {
		return
	}
	if (len(SourceList) == 1 && SourceList[0] == "") {
		SourceListString = fmt.Sprintf("%s; " ,repo_path)
	} else {
		SourceListString = fmt.Sprintf("%s; %s", SourceListString, repo_path)
	}
	if (len(SourceList) == 1 && SourceList[0] == "") {
		SourceList = []string{repo_path}
	} else {
		SourceList = append(SourceList, repo_path)
	}
	SourceMap[repo_path] = 1
	WriteFile(MainConfig.GetConfString("SourceListPath", "/opt/ppalist/sourcelist"), SourceListString)
}

func AddRepoToPpaList(repo_path string) {
	if _, ok := PpaMap[repo_path]; ok {
		return
	}
	repo_path = strings.Trim(repo_path, " ")
	if (len(PpaList) == 1 && PpaList[0] == "") {
		PpaListString = fmt.Sprintf("%s;" ,repo_path)
	} else {
		PpaListString = fmt.Sprintf("%s;%s", PpaListString, repo_path)
	}
	if (len(PpaList) == 1 && PpaList[0] == "") {
		PpaList = []string {repo_path}
	} else {
		PpaList = append(PpaList, repo_path)
	}
	PpaMap[repo_path] = 1
	WriteFile(MainConfig.GetConfString("PpaListPath", "/opt/ppalist/ppalist"), PpaListString)
}

func AddSoft(soft string) {
	fmt.Println("len(SoftList) = ", len(SoftList))
	soft_list := strings.Split(soft, " ")
	if (len(soft_list) == 0) {
		return
	}

	soft = strings.Trim(soft, " ")
	if (len(SoftList) == 1 && SoftList[0] == "") {
		SoftListString = strings.Replace(soft, " ", ";", len(soft_list))
		SoftListString = fmt.Sprintf("%s;" ,SoftListString)
	} else {
		soft_string := strings.Replace(soft, " ", "; ", len(soft_list))
		if _, ok := SoftMap[soft]; !ok {
			SoftListString = fmt.Sprintf("%s%s;", SourceListString, soft_string)
		}

	}

	if (len(soft_list) == 1) {
		if (len(SoftList) == 1 && SoftList[0] == "") {
			SoftList = []string{soft}
		} else {
			if _, ok := SoftMap[soft]; !ok {
				SoftList = append(SoftList, soft)
			}
		}
	} else {
		for _, soft_name := range soft_list {
			if (len(SoftList) == 1 && SoftList[0] == "") {
				SoftList = []string{soft_name}
				SoftMap[soft_name] = 1
			} else {
				if _, ok := SoftMap[soft_name]; !ok {
					SoftList = append(SoftList, soft_name)
					SoftMap[soft_name] = 1
				}
			}
		}
	}
	WriteFile(MainConfig.GetConfString("SoftListPath", "/opt/ppalist/softlist"), SoftListString)
}

func WriteFile(filename, data string) {
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}

func ReadPpaList() error {
	PpaListPath := fmt.Sprintf("%s%s", MainDir, ppalist_filename)
	PpaListFile, err := os.Open(PpaListPath)
	if err != nil {
		fmt.Println("Erro while loading PPA list: ", err.Error())
		fmt.Println("Create New PPA list file")
		os.Create(PpaListPath)
		err2 := ReadPpaList()
		return err2
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

	PpaListString = strings.Replace(string(bs), " ", "", 99999)
	fmt.Println("PpaListString = ", PpaListString)
	PpaList = strings.Split(PpaListString, ";")
	if (len(PpaList) == 1 && PpaList[0] == "") {
		return nil
	}
	for _, ppa_name := range PpaList {
		PpaMap[ppa_name] = 1
	}
	return nil
}

func ReadSourceList() error {
	SourceListPath := fmt.Sprintf("%s%s", MainDir, sourcelistfilename)
	SourceListFile, err := os.Open(SourceListPath)
	if err != nil {
		fmt.Println("Erro while loading Source list: ", err.Error())
		fmt.Println("Create New Source list file")
		os.Create(SourceListPath)
		err2 := ReadSourceList()
		return err2
	}
	defer SourceListFile.Close()

	stat, err := SourceListFile.Stat()
	if err != nil {
		fmt.Println("Erro while get PPA list size: ", err.Error())
		return err
	}
	bs := make([]byte, stat.Size())
	_, err = SourceListFile.Read(bs)
	if err != nil {
		return err
	}

	SourceListString = strings.Replace(string(bs), " ", "", 99999)
	fmt.Println("SourceListString = ", SourceListString)
	SourceList = strings.Split(SourceListString, ";")
	if (len(SourceList) == 1 && SourceList[0] == "") {
		return nil
	}
	for _, source_name := range SourceList {
		SourceMap[source_name] = 1
	}
	return nil
}

func ReadSoftList() error {
	SoftListPath := fmt.Sprintf("%s%s", MainDir, softlistfilename)
	SoftListFile, err := os.Open(SoftListPath)
	if err != nil {
		fmt.Println("Erro while loading Soft list: ", err.Error())
		fmt.Println("Create New Soft list file")
		os.Create(SoftListPath)
		err2 := ReadSoftList()
		return err2
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

	SoftListString = strings.Replace(string(bs), " ", "", 99999)
	fmt.Println("SoftListString = ", SoftListString)
	SoftList = strings.Split(SoftListString, ";")
	if (len(SoftList) == 1 && SoftList[0] == "") {
		return nil
	}
	for _, soft_name := range SoftList {
		SoftMap[soft_name] = 1
	}
	return nil
}
