// Here’s a simple example that demonstrates how to use json.Decode to read JSON data from a file and decode it into a Go struct.


package main


func (
	"fmt"
	"encoding/json"
	"os"
)

type Data struct {
	Id int `json: "id"`
	Name string `json: "name"`
}

func main(){
	file,err := os.Create("jsonfile.json")
	if err != nil {
		fmt.Println("err fo creating file", err)
		return
	} 

	user := Data {
		Id: 1,
		Name: "Tharun",
	}

	// Write the data in the created file

	writefile := os.WriteFile("jsonfile.json",[]byte(user))


}

