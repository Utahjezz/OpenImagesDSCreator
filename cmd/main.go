package main

import (
	"OpenImagesDSCreator/cmd/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/kniren/gota/dataframe"
)

func getJsonHierarchy() (model.Category, error) {
	var category model.Category
	resp, err := http.Get("https://storage.googleapis.com/openimages/2018_04/bbox_labels_600_hierarchy.json")

	if err != nil {
		fmt.Println("err1")
	}
	defer resp.Body.Close()

	errox := json.NewDecoder(resp.Body).Decode(&category)

	return category, errox
}

func getCsvClassMapping() (dataframe.DataFrame, error) {
	var err error
	resp, err := http.Get("https://storage.googleapis.com/openimages/v5/class-descriptions.csv")

	if err != nil {
		return dataframe.DataFrame{}, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return dataframe.DataFrame{}, err
	}
	ioContent := strings.NewReader(string(respBytes))

	df := dataframe.ReadCSV(ioContent, dataframe.WithDelimiter(','), dataframe.HasHeader(true))
	return df, nil
}

func main() {
	category, err := getJsonHierarchy()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(category)
	classMapping, err := getCsvClassMapping()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(classMapping)
}
