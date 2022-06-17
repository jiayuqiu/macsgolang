package ch4

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type XKCDresponse struct {
	Month      string
	Num        int64
	Link       string
	Year       string
	News       string
	SafeTitle  string
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

func searchComic(comicID int) (*XKCDresponse, error) {
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
	resp.Body.Close()
	return &result, nil
}

func Practice412() {
	//The popular web comic xkcd has a JSON interface. For example,
	//a request to http://xkcd.com/571/info.0.json produces a detailed description of comic 571,
	//one of many favorites. Download each URL (once!) and build an offline index.
	//Write a tool xkcd that, using this index,
	//prints the URL and transcript of each comic that matches a search term provided on the command line.
	//TODO: 可以添加一个函数，以title作为key进行检索。
	// for i := 571; i <= 575; i++ {
	// 	result, err := searchComic(i)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	comicTitle := result.Title
	// 	dataString, err := json.MarshalIndent(result, "", "    ")
	// 	if err != nil {
	// 		log.Fatalf("JSON MarshalIndent failed: %s", err)
	// 	}
	// 	content := []byte(dataString)
	// 	err = ioutil.WriteFile(
	// 		fmt.Sprintf("/Users/qiujiayu/GolandProjects/macsgolang/utils/xkcdComic/%v.json", comicTitle),
	// 		content,
	// 		0644,
	// 	)
	// 	if err != nil {
	// 		fmt.Println(">>>>>>>>>>>>>>>>>>> error start >>>>>>>>>>>>>>>>>")
	// 		fmt.Println(err)
	// 		fmt.Println(">>>>>>>>>>>>>>>>>>>  error end  >>>>>>>>>>>>>>>>>")
	// 	}
	// }

	// 读取文件夹内的所有内容，并生成hash表
	// key文件title, value为json对应struct
	folderPath := "/Users/qiujiayu/GolandProjects/macsgolang/utils/xkcdComic/"
	files, _ := ioutil.ReadDir(folderPath)
	localComicInfo := make(map[string]*XKCDresponse)
	for _, f := range files {
		// fmt.Println(f.Name())
		// 读取json文件，并将文件中的内容转换为struct
		fmt.Println(folderPath + f.Name())
		strBytes, err := ioutil.ReadFile(folderPath + f.Name())
		// fmt.Printf("%v %T", strBytes, strBytes)

		if err != nil {
			fmt.Println("读取json文件失败", err)
			fmt.Println(f.Name())
			return
		}

		// var res XKCDresponse
		res := &XKCDresponse{} // 此处初始化需要用指针

		err = json.Unmarshal(strBytes, res)
		if err != nil {
			fmt.Println("解析数据失败", err)
			return
		}
		localComicInfo[f.Name()] = res
	}

	fmt.Println(localComicInfo["Addiction.json"])
	fmt.Println(localComicInfo["Addiction.json"].Title)
}
