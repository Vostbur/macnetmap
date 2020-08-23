package main

import (
	"io/ioutil"
	"path"
	"regexp"
)

func readFile(filename string) (string, error) {
	var data []byte
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func parseShow(show string, reExp string) (map[string]string, error) {
	re, err := regexp.Compile(reExp)
	if err != nil {
		return nil, err
	}
	result := re.FindAllStringSubmatch(show, -1)
	resMap := make(map[string]string)
	for _, group := range result {
		resMap[group[1]] = group[2]
	}
	return resMap, nil
}

func parseHostname(raw string, reExp string) (string, error) {
	re, err := regexp.Compile(reExp)
	if err != nil {
		return "", err
	}
	match := re.FindStringSubmatch(raw)[1]
	return match, nil
}

func readDir(dir string) ([]string, error) {
	fileObjects, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var files []string
	for _, f := range fileObjects {
		if !f.IsDir() {
			files = append(files, path.Join(dir, f.Name()))
		}
	}
	return files, nil
}
