/**
*@Author: haoxiongxiao
*@Date: 2019/4/3
*@Description: CREATE GO FILE irisgo
 */
package main

import (
	"github.com/spf13/pflag"
	"github.com/xhaoxiong/irisgo/commands/new"
	"github.com/xhaoxiong/irisgo/utils"
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
	appName := ""

	if len(pflag.Args()) == 0 {
		appName = "irisApp"
		commands.CreatedApp(currentpath, appName)
	}else{
		for v := range pflag.Args() {
			if v == 0 {
				if pflag.Args()[v] == "new" {
					if pflag.Args()[v+1] != "" {
						appName = pflag.Args()[v+1]
						commands.CreatedApp(currentpath, appName)
					}
				}
			}
		}
	}



}
