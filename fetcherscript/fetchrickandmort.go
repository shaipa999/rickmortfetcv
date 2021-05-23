package main
import (
"fmt"
"io/ioutil"
"log"
"net/http"
"encoding/json"
"strings"
)



func rickUrlln(data string ){

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
			fmt.Printf("%v,%v,%v \n",namevstr,locnamevstr,imagevstr)	
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
				fmt.Println( err)
				log.Fatal(err)
				}
				instID, err := ioutil.ReadAll(resp.Body)
				if err != nil {
				fmt.Println( err)
				log.Fatal(err)
				}
				defer resp.Body.Close()
				//
				rickUrlln(string(instID))
			}
		}
	}
	
	    

	
}

func ContainsI(a string, b string) bool {
        return strings.Contains(
            strings.ToLower(a),
            strings.ToLower(b),
        )
    }



func main() {
	resp, err := http.Get("https://rickandmortyapi.com/api/character/?status=alive&species=Human")
	if err != nil {
	fmt.Println(err)
	log.Fatal(err)
	}
	instID, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	fmt.Println(err)
	log.Fatal(err)
	}
	defer resp.Body.Close()
	//
        rickUrlln(string(instID))
}
