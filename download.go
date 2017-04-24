package main

import (
	"./jsontools"
	"fmt"
	"github.com/go-resty/resty"
)

func main() {

	resty.SetOutputDirectory("/tmp57/data")

	pages := jsontools.GetPages("./chhs.json")
	// jsontools.Print(pages)

	for _, p := range pages {
		fmt.Println(p.Filename)
		//fmt.Println(p.Url)

		_, err := resty.R().
			SetOutput(p.Filename + ".json").Get(p.Url + ".json")

		if err != nil {
			fmt.Println(err)
		}

		_, err = resty.R().SetOutput(p.Filename + ".csv").Get(p.Url + ".csv")

		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("done")

}
