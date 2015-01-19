package main

import "net/http"
import "io/ioutil"
import "regexp"
import "fmt"
import "os"
import "io"

func main() {

    fmt.Println("Executing: ", os.Args)
    if(len(os.Args) >= 4){
    	fmt.Println("Accessing the link: ",os.Args[1])
    	resp, _ := http.Get(os.Args[1])
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		s := string(body)
		r, _ := regexp.Compile(os.Args[2])
		links := r.FindAllStringSubmatch(s,-1)
		for index,element := range links {
			fmt.Println("Link ",index);
			fulllink := os.Args[3]+element[1]
			fmt.Println("Downloading: ",fulllink)
			out, _ := os.Create(element[2])
			defer out.Close()
			resp, _  := http.Get(fulllink)
			defer resp.Body.Close()
			io.Copy(out, resp.Body)
		}	
	}
}