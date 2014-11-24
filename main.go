package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry-incubator/receptor"
	"github.com/pivotal-cf-experimental/veritas/say"
)

func main() {
	if len(os.Args) <= 2 {
		PrintUsageAndExit()
	}

	receptorAddr := os.Getenv("RECEPTOR")
	if receptorAddr == "" {
		fmt.Println("No RECEPTOR set")
		PrintUsageAndExit()
	}

	client := receptor.NewClient(receptorAddr)

	if os.Args[1] == "destroy-task" {
		for _, taskGuid := range os.Args[2:] {
			say.Println(0, "Destroying Task: %s", say.Red(taskGuid))
			client.CancelTask(taskGuid)
			client.DeleteTask(taskGuid)
		}
		os.Exit(0)
	}

	if os.Args[1] == "destroy-lrp" {
		for _, processGuid := range os.Args[2:] {
			say.Println(0, "Destroying LRP: %s", say.Red(processGuid))
			client.DeleteDesiredLRP(processGuid)
		}
		os.Exit(0)
	}

	PrintUsageAndExit()
}

func PrintUsageAndExit() {
	fmt.Println(`Usage:
worf destroy-task TASK_GUID TASK_GUID ...
worf destroy-lrp PROCESS_GUID PROCESS_GUID ...

Set the receptor address with the RECEPTOR environment:
    export RECEPTOR=http://username:password@receptor.ketchup.cf-app.com

The address for a local Diego Edge box can be set via: 
    export RECEPTOR=http://receptor.192.168.11.11.xip.io
`)
	os.Exit(1)
}

func ExitIfError(err error) {
	if err != nil {
		fmt.Printf("Got an unexpected error:\n%s\n", err.Error())
		os.Exit(1)
	}
}
