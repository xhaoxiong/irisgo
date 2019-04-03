/**
*@Author: haoxiongxiao
*@Date: 2019/4/3
*@Description: CREATE GO FILE irisgo
 */
package main

import (
	"github.com/spf13/pflag"
	"irisgo/commands/new"
	"irisgo/utils"
	"log"
	"os"
)

func main() {

	pflag.Parse()
	currentpath, _ := os.Getwd()
	if !utils.IsExist(currentpath) {
		log.Printf("Application '%s' already exists", currentpath)
		os.Exit(0)
	}
	for v := range pflag.Args() {
		if pflag.Args()[v] == "new" {
			if pflag.Args()[v+1] != "" {
				appName := pflag.Args()[v+1]
				commands.CreatedApp(currentpath, appName)
			}
		}
	}

}
