//This is designed to be a script/a program that starts at start up
//docker start name
//https://gobyexample.com/spawning-processes
//go run .
//go build .
//go build docker.go main.go
//The formating might be weird. I'm programming in VIM
package main

import ("fmt"
        "os/exec"
)
type dockerInfo struct {
        Base string
        start string
        containerName string
        //sudo string //Update sudo
}

func dockerLaunch() {
        mongo := dockerInfo{"docker","start", "nodeTest"}
//      var complete = mongo.Base + mongo.space + mongo.containerName
        fmt.Println("Launched: " + mongo.containerName)
        var launch = exec.Command(mongo.Base, mongo.start, mongo.containerName)
        //var launch = exec.Command("docker", "start", "nodeTest")
        var launchOutput, err = launch.Output() //Gets the output from exec.command
        fmt.Println(string(launchOutput))
        if err != nil { //Nil == NULL
                panic(err) //err is a keyword in Go panic is another built in function to handle unexpected error. The panic function will print the error
        }
}