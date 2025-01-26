package main

import (
	"task-tracker/cmd"
	"task-tracker/util"
)

func init() {
	util.Load()
}

func main() {
	defer util.Save()
	cmd.RootCmd.Execute()

}
