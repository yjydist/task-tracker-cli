package main

import (
	"task-tracker/cmd"
	"task-tracker/util"
)

func init(){
	util.Load()
	util.Save()
}

func main() {
	cmd.RootCmd.Execute()

}
