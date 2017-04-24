package main

import (
	"./jsontools"
	//"github.com/go-resty/resty"
)

func main() {

	pages := jsontools.GetPages("./chhs.json")
    jsontools.Print(pages)


/*
	resty.SetOutputDirectory("/tmp57/data")

	_, err := resty.R().
		SetOutput("medreview.json").
		Get("https://chhs.data.ca.gov/api/views/pch7-48qc/rows.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("done")
*/
}
