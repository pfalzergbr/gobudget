package main

func main() {
	apiServer := NewAPIServer(":8080")
	apiServer.Run()
}
