package main

import (
	"encoding/json"
    "fmt"
    "io/ioutil"
)


func main(){
	jsonFile,_ := ioutil.ReadFile("data.json")
	var datas []map[string]interface{}
	json.Unmarshal([]byte(jsonFile) , &datas)

	//success
	allElectronicDevice(datas)
	allPurcasedOn16Jan2020(datas)
	
	//masih salah
	itemsMeetingRoom(datas)
	allFurniture(datas)
	allBrownColor(datas)	
}

// 1. Find items in the Meeting Room.
func itemsMeetingRoom(datas []map[string]interface{}){
	for _, data := range datas{
		tmp := data["placement"]
			fmt.Print(tmp)
	}
}
// 2. Find all electronic devices.
func allElectronicDevice(datas []map[string]interface{}){
	for _, data := range datas{
		if data["type"]=="electronic"{
			fmt.Print(data["inventory_id"])
		}
	}

}
// 3. Find all the furniture.
func allFurniture(datas []map[string]interface{}){
	for _, data := range datas{
		tmp := data["tags"]
			if true{
				fmt.Print(tmp)
			}
	}
}
// 4. Find all items were purchased on 16 Januari 2020.
func allPurcasedOn16Jan2020(datas []map[string]interface{}){
	for _, data := range datas{
		if data["purcased_at"]=="16012020"{
			fmt.Print(data["name"])
		}
	}
}

// 5. Find all items with brown color.
func allBrownColor(datas []map[string]interface{}){
	for _, data := range datas{
		tmp := data["tags"]
			if true{
				fmt.Print(tmp)
			}
	}
}
