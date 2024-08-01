/*

some Agent
Copyright tg:@cybittheir, 2024

*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	appname string
	version string
	build   string
	silent  bool = false // display all messages
)

const (
	def_period  = 60 // default circle period
	err_timeout = 10 // timeout if connections error
)

type netMACIP struct {
	ip  string
	mac string
}

func Greeting() {

	fmt.Println(appname, version, "build", build)
	fmt.Println("Simple agent for checking links")
	fmt.Println("https://github.com/cybittheir/sAgent")

}

func sendQuery(url, token, urlQuery string) (bool, error) {
	// Make HTTP GET request
	response, err := http.Get(url + "?UID=" + token + urlQuery)
	//	resp, err := http.Post(url + "?UID=" + token,urlQuery)
	if err == nil {
		if !silent {
			io.WriteString(os.Stdout, fmt.Sprintf("Sending info to server: \n%s", urlQuery))
		}
		// Copy data from the response to standard output
		_, warning := io.Copy(os.Stdout, response.Body)
		if warning != nil {
			return false, warning
		}
	} else {
		return false, err
	}
	defer response.Body.Close()
	return true, nil

}

func readConfJSON() ([]byte, error) {

	// Open our jsonFile

	var byteValue []byte
	jsonFile, err := os.Open("conf.json")
	// if we os.Open returns an error then handle it

	if err == nil {

		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		byteValue, warning := io.ReadAll(jsonFile)

		return byteValue, warning
	} else {
		log.Println("Cannot open conf.json")
		log.Println(err)
		os.Exit(1)
	}
	return byteValue, err

}

func initArgs() {
	// reading starting parameters

	if len(os.Args) != 1 {

		arg := os.Args[1]

		if arg == "-h" || arg == "-help" {

			fmt.Println("")
			fmt.Println("Use conf.json file configuration:")
			fmt.Println("connect:")
			fmt.Println("   period: 60 // seconds, 60 sec minimum")
			fmt.Println("   url: https://[url]")
			fmt.Println("   token: [token] // ?UID=token")
			fmt.Println("   pin: 000000 // ?pin=[pin]")
			fmt.Println("   batch: batch.bat // context with tasklist.exe /FO CSV /NH | find '%1'")
			fmt.Println("   path: C:\\[PATH]\\ //path to batch-file")
			fmt.Println("check:")
			fmt.Println("   process: //tests applications are running. =9 if OK, =1 if failed")
			fmt.Println("      app1: app1.exe")
			fmt.Println("      app2: app2.exe")
			fmt.Println("   device: //tests hosts are reachable. =failed if NOT")
			fmt.Println("      dev1:")
			fmt.Println("         ip: 192.168.0.1")
			fmt.Println("         port: 80")
			fmt.Println("      dev2:")
			fmt.Println("         ip: 192.168.0.2")
			fmt.Println("         port: 22")
			fmt.Println("")
			fmt.Println("Use -s (OR -silent) option for hiding all messages except errors")
			fmt.Println("Use Ctrl+C for exit")
			fmt.Println("")

			os.Exit(0)
		} else if arg == "-s" || arg == "-silent" {
			silent = true // display only main messages & errors
		}
	}
}

func printMSGln(message string) {
	if !silent {
		fmt.Println(message)
	}
}

func printMSG(message string) {
	if !silent {
		fmt.Print(message)
	}
}

func main() {

	Greeting()

	initArgs() // read commandline parameters

	byteValue, err := readConfJSON()

	if err == nil {
		printMSGln("Successfully opened config file")
	} else {
		log.Println(err)
		os.Exit(1)
	}

	var confResult map[string]map[string]string
	var checkProc map[string]map[string]map[string]string
	var checkConn map[string]map[string]map[string]map[string]string

	json.Unmarshal([]byte(byteValue), &confResult)
	json.Unmarshal([]byte(byteValue), &checkProc)
	json.Unmarshal([]byte(byteValue), &checkConn)

	//	var emptyConfig bool
	//	var timePeriod int

	// checking config parameters
	f, err := os.Open("c:/tc")
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		fmt.Println(v.Name(), v.IsDir())
	}
}
