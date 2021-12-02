package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// VSI struct which contains array of vsi's
type VSIs struct {
	VSIs []Vsi `json:"vs"`
}

// VM struct which contains a Name a location and bools for enabling
type Vsi struct {
	Location string `json:"location"`
	Name     string `json:"name"`
	Start    bool   `json:"start"`
	Gui      bool   `json:"gui"`
}

func main() {

	// variables
	VMstart := ""
	VMgui := ""
	vmrun := "C:\\Program Files (x86)\\VMware\\VMware Workstation\\vmrun.exe"

	// Open our jsonFile
	jsonFile, err := os.Open("D:\\Virtual Machines\\AutoStart\\start-vsi.json")

	// if os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that to parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// initialize our vsi array
	var vsi VSIs

	// jsonFile's content into 'vsi' which was defined above
	json.Unmarshal(byteValue, &vsi)

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
		if err != nil {
			fmt.Print("Status: Error! ", err)
		} else {
			fmt.Println("Status: Started...")
			fmt.Printf("%s\n", out)
		}

	}

}
