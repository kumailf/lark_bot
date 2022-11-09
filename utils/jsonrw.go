package utils

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type PersonInfo struct {
// 	Name    string
// 	age     int32
// 	Sex     bool
// 	Hobbies []string
// }

// func readJsonFile(jsondata interface{}, filepath) interface{} {

// 	filePtr, err := os.Open("person_info.json")
// 	if err != nil {
// 		fmt.Println("Open file failed [Err:%s]", err.Error())
// 	}
// 	defer filePtr.Close()

// 	var person []PersonInfo

// 	// 创建json解码器
// 	decoder := json.NewDecoder(filePtr)
// 	err = decoder.Decode(&person)
// 	if err != nil {
// 		fmt.Println("Decoder failed", err.Error())

// 	} else {
// 		fmt.Println("Decoder success")
// 		fmt.Println(person)
// 	}
// 	return person
// }

// func writeJsonFile() {
// 	personInfo := []PersonInfo{{"David", 30, true, []string{"跑步", "读书", "看电影"}}, {"Lee", 27, false, []string{"工作", "读书", "看电影"}}}
// 	old := readFile()
// 	personInfo = append(personInfo, old...)

// 	// 创建文件
// 	filePtr, err := os.Create("person_info.json")
// 	if err != nil {
// 		fmt.Println("Create file failed", err.Error())
// 		return
// 	}
// 	defer filePtr.Close()

// 	// 创建Json编码器
// 	encoder := json.NewEncoder(filePtr)

// 	err = encoder.Encode(personInfo)
// 	if err != nil {
// 		fmt.Println("Encoder failed", err.Error())

// 	} else {
// 		fmt.Println("Encoder success")
// 	}
// }

// func appendJsonFile() {

// }
