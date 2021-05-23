package main
import (
"fmt"
"io/ioutil"
"log"
"net/http"
"encoding/json"
"strings"
"crypto/tls"
)


var name []string
var locs []string
var image []string

func appHandlerhealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"ok")
}


func appHandler(w http.ResponseWriter, r *http.Request) {
 
resp, err := http.Get("https://rickandmortyapi.com/api/character/?status=alive&species=Human")
	if err != nil {
	fmt.Fprintln(w, err)
	log.Fatal(err)
	}
	instID, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	fmt.Fprintln(w, err)
	log.Fatal(err)
	}
	defer resp.Body.Close()
	//
        rickUrl(w,string(instID))

	parseData := make([]map[string]interface{}, 0, 0)

    	for counter,_ := range name {
		var singleMap = make(map[string]interface{})
		singleMap["name"] = name[counter]
		singleMap["location"] = locs[counter]
		singleMap["image"] = image[counter]
		parseData = append(parseData, singleMap)
    	}

	b, _:= json.Marshal(parseData)
	fmt.Fprintln(w,string(b))
}

func ContainsI(a string, b string) bool {
        return strings.Contains(
            strings.ToLower(a),
            strings.ToLower(b),
        )
    }

func rickUrl(w http.ResponseWriter,data string ){

	var result map[string]interface{}
    	json.Unmarshal([]byte(data), &result)
	
	isChars, ok5 := result["results"]
	_ = isChars	
	if ok5 {
		for _,item:=range isChars.([]interface{}) {
			vdet := item.(map[string]interface{})


			namev := vdet["name"]
			namevstr := fmt.Sprintf("%v", namev)
			imagev := vdet["image"]
			imagevstr := fmt.Sprintf("%v", imagev)
			locationv := vdet["location"]
			md:= locationv.(map[string]interface{})
			locnamev := md["name"]
			locnamevstr := fmt.Sprintf("%v", locnamev)
	
			if ContainsI(locnamevstr,"earth") {
			name = append(name, namevstr)
			locs = append(locs, locnamevstr)
			image = append(image, imagevstr)		
			}
		}
    		
	}


	isInfo, ok1 := result["info"].(map[string]interface{})
	_ = isInfo	
	if ok1 {
		isNext, ok2 := isInfo["next"]
		_ = isNext	
		if ok2 {
    			
			strUrlNext := fmt.Sprintf("%v", isNext)
			
			if strUrlNext != "<nil>" {
				
				resp, err := http.Get(strUrlNext)
				if err != nil {
				fmt.Fprintln(w, err)
				log.Fatal(err)
				}
				instID, err := ioutil.ReadAll(resp.Body)
				if err != nil {
				fmt.Fprintln(w, err)
				log.Fatal(err)
				}
				defer resp.Body.Close()
				//
				rickUrl(w,string(instID))
			}
		}
	}
	
	    

	
}
func main() {
http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
http.HandleFunc("/fetch", appHandler)
http.HandleFunc("/healthcheck", appHandlerhealth)
log.Println("Started, serving on port 8080")
err := http.ListenAndServe(":8080", nil)
if err != nil {
log.Fatal(err.Error())
}

}
