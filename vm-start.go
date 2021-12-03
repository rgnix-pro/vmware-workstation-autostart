package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// structured json
// VSIs struct which contains array of vsi's
type VSIs struct {
	VSIs []Vsi `json:"vs"`
}

// Vsi struct which contains a Name, location and bools for enabling
type Vsi struct {
	Location string `json:"location"`
	Name     string `json:"name"`
	Start    bool   `json:"start"`
	Gui      bool   `json:"gui"`
}

// err check
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// int vars
	VMstart := ""
	VMgui := ""

	// unstructured data
	jsonFile, err := os.Open("config.json")
	check(err)

	defer jsonFile.Close()

	byteVal, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteVal), &result)

	vmrun := fmt.Sprint(result["vmrun"])
	
	startJSON, err := os.Open(fmt.Sprint(result["vm-start"]))
	check(err)


	defer startJSON.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(startJSON)

	// initialize our vsi array
	var vsi VSIs

	// jsonFile's content into 'vsi' which was defined above in structs
	json.Unmarshal(byteValue, &vsi)

	// structured json
	// print results
	fmt.Println("")
	for i := 0; i < len(vsi.VSIs); i++ {
		fmt.Println("VS Name: " + vsi.VSIs[i].Name)
		fmt.Println("VS Location: " + vsi.VSIs[i].Location)
		fmt.Printf("VS Start: %v\n", vsi.VSIs[i].Start)
		fmt.Printf("VS Gui: %v\n", vsi.VSIs[i].Gui)

		if vsi.VSIs[i].Start {
			VMstart = "start"
		}

		if vsi.VSIs[i].Gui {
			VMgui = "gui"
		} else {
			VMgui = "nogui"
		}

		// start the vsi's
		out, err := exec.Command(vmrun, VMstart, vsi.VSIs[i].Location+vsi.VSIs[i].Name+"\\"+vsi.VSIs[i].Name+".vmx", VMgui).Output()
		check(err)
		
		if err == nil {
			fmt.Println("Status: Starting...")
			fmt.Printf("%s\n", out)
		}

	}

}
