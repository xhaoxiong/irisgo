/**
*@Author: haoxiongxiao
*@Date: 2019/4/3
*@Description: CREATE GO FILE irisgo
 */
package main

import (
	"github.com/spf13/pflag"
	"github.com/xhaoxiong/irisgo/commands/api_template"
	"github.com/xhaoxiong/irisgo/commands/mvc_template"
	"github.com/xhaoxiong/irisgo/utils"
	"log"
	"os"
)

var cliProgramName = pflag.StringP("name", "n", "irisApp", "input -n=$value(your program name)")
var cliApi = pflag.BoolP("api", "a", true, "input --api=$val(true or false) and you will get api template")
var cliMVC = pflag.BoolP("mvc", "m", false, "input --mvc=$val(true or false) and you will get mvc template ")

func main() {

	pflag.Parse()
	currentpath, _ := os.Getwd()
	if !utils.IsExist(currentpath) {
		log.Printf("Application '%s' already exists", currentpath)
		os.Exit(0)
	}

	if *cliMVC {
		mvc_template.CreatedApp(currentpath, *cliProgramName)
		return
	}
	if *cliApi {
		api_template.CreatedApp(currentpath, *cliProgramName)
		return
	}

}
