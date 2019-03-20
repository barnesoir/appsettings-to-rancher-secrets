package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter path to .json file")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	plan, _ := ioutil.ReadFile(text)
	var data map[string]interface{}
	err := json.Unmarshal(plan, &data)

	if err != nil {
		fmt.Println(err)
	}

	getEnvVar("", data)
}

func getEnvVar(prefix string, data map[string]interface{}) {
	for key, val := range data {
		newKey := key

		if len(prefix) > 0 {
			newKey = prefix + "__" + key
		}

		if reflect.TypeOf(val).Kind() == reflect.Map {
			getEnvVar(newKey, val.(map[string]interface{}))
		} else {
			fmt.Println(newKey, "=", val)
		}
	}
}
