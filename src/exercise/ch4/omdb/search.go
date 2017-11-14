package omdb

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"strings"
	"os"
	"io/ioutil"
	"io"
	"bytes"
)

func SearchMovie(key string) (*SearchResult, error) {
	searchUrl := OmdbApiURL + key
	resp, err := http.Get(searchUrl)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search failed: %s", resp.Status)
	}

	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

func SavePoster(imgUrl string) (int64, error) {

	if imgUrl == "N/A" {
		fmt.Println(imgUrl + " is not url")
		return 0, nil
	}

	path := strings.Split(imgUrl, "/")
	var name string
	if len(path) > 1 {
		name = path[len(path)-1]
	}
	fmt.Printf("name = %s\n", name)
	out, err := os.Create(name)
	defer out.Close()
	resp, err := http.Get(imgUrl)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal("Download image error!")
		return 0, nil
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Download image error!")
		return 0, nil
	}

	pix, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Save image error!")
		return 0, nil
	}
	return io.Copy(out, bytes.NewBuffer(pix))

}
