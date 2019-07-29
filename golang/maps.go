package main

import "fmt"

func main() {

	emails := make(map[string]string)
	
	emails["kibet"] = "kibetleonard8@gmail.com"
    emails["leonard"] = "leonardmagutbiz@gmail.com"

    fmt.Println(emails)
    fmt.Println(len(emails))
    delete(emails,"kibet")
    fmt.Println(emails)

}