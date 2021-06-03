package main

import (

	"cowin/pkg/schedule"
	"time"
	"flag"
	//"os"

)




func main() {
	
	
	
    //pincode:=os.Args[1]
	pincode:=flag.String("p")
	flag.Parse()
	
	for _ = range time.Tick(time.Second * 10){
	
	if

	
schedule.RequestVaccineAvailability(time.Now().Format("02-01-2006"),pincode)==true{

	       break }
	
	}
}
	


	




