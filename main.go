package main

func main() {
	a := App{}
	a.Initialize("root", "", "gorest")

	a.Run(":8080")
}
