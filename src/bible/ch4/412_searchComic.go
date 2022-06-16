package ch4

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type XKCDresponse struct {
	Month string
	Num int64
	Link string
	Year string
	News string
	SafeTitle string
	Transcript string
	Alt string
	Img string
	Title string
	Day string
}

func searchComic(comicID int)(*XKCDresponse, error){
	//q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(fmt.Sprintf("https://xkcd.com/%d/info.0.json", comicID))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result XKCDresponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	//fmt.Println(result)
	resp.Body.Close()
	return &result, nil
}


func Practice412(){
	//The popular web comic xkcd has a JSON interface. For example,
	//a request to http://xkcd.com/571/info.0.json produces a detailed description of comic 571,
	//one of many favorites. Download each URL (once!) and build an offline index.
	//Write a tool xkcd that, using this index,
	//prints the URL and transcript of each comic that matches a search term provided on the command line.
	//TODO: 可以添加一个函数，以title作为key进行检索。
	for i := 571; i <= 600; i++ {
		result, err := searchComic(i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v \n", result)
		comicTitle := result.Title
		dataString, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			log.Fatalf("JSON MarshalIndent failed: %s", err)
		}
		content := []byte(dataString)
		err = ioutil.WriteFile(
				fmt.Sprintf("/Users/qiujiayu/GolandProjects/macsgolang/utils/xkcdComic/%v.json", comicTitle),
				content,
				0644,
			)
		if err != nil {
			fmt.Println(">>>>>>>>>>>>>>>>>>> error start >>>>>>>>>>>>>>>>>")
			fmt.Println(err)
			fmt.Println(">>>>>>>>>>>>>>>>>>>  error end  >>>>>>>>>>>>>>>>>")
		}
	}
}
