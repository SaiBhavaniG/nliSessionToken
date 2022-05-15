package main

import "fmt"

func main() {
	fmt.Println("Hi")
}
/*
api is called - client do request server(backend) 
IN Return Backend(server) will respond

1. https://localhost:8080/api/v1/token?deviceid=12 : GET (Request Method)
2. Request: Query: deviceid=12
		    Bddy: {json} //SCHEMA : POST, PDATE, PUT, DELETE, GET
3. Response: {JSON} //SCHEMA
*/