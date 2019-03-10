package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "CHECK")
	fmt.Fprintln(os.Stderr, ".. reading concourse input")

	fmt.Fprintln(os.Stdout, "[{\"num\": \"11\"}]")
	os.Exit(0)
	return

	// // Read input from Concourse
	// storage, err := common.Load(os.Stdin)
	// if err != nil {
	// 	debug.PrintStack()
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }

	// // Bump build number if it's required
	// if storage.Source.DoBump {
	// 	fmt.Fprintln(os.Stderr, ".. bumping the build number")
	// 	err = storage.Bump()
	// 	if err != nil {
	// 		debug.PrintStack()
	// 		fmt.Fprintln(os.Stderr, err)
	// 		os.Exit(1)
	// 	}
	// }

	// // Return saved build number
	// fmt.Fprintln(os.Stderr, ".. returning the build number")
	// buildNumber, err := storage.Get()

	// response := []common.Version{
	// 	{BuildNumber: strconv.Itoa(buildNumber)},
	// }

	// data, err := json.Marshal(response)
	// if err != nil {
	// 	debug.PrintStack()
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }

	// fmt.Println(string(data))
}
