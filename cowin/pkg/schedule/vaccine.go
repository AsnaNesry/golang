package schedule

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	
	

)

type Center struct {
	CentreId int `json:"center_id"`
}

type VaccineCenter struct {
	Centers []Center `json:"centers"`
}

func RequestVaccineAvailability(date string,pincode string) bool{

        var status bool
        status =false
	client := &http.Client{}
	
	req, err := http.NewRequest("GET", "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByPin",nil)
	
	
	
	
	
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")
	
	q:=req.URL.Query()
	q.Set("pincode",pincode)
	q.Set("date",date)
	req.URL.RawQuery=q.Encode()
	
	
	//fmt.Println(req.URL.String())
	
	

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)

	}
	
	
	var jsonData = []byte(string(body))
	
	var center VaccineCenter
	
	err = json.Unmarshal(jsonData, &center)
	
	if err != nil {
		log.Println(err)
	}
	
	
	//fmt.Println(center.Centers)
	
	if len(center.Centers)>0{
	
	fmt.Println("Available")
	
	log.Println(string(body))
	
	status=true
	
	}else{
	
	fmt.Println("not available")
	
	}

return status
}	

       
     
      
