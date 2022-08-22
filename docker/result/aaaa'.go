// package main
// import (
// 	"fmt"
// 	"os/exec"
// )

// func command() {

// 	cmd := exec.Command("ls", "-a")
// 	cmd.Dir = "/code"
	
// 	output, err := cmd.Output()

// 	// output : [46 10 46 46 10 99 111 109 109 97 110 100 46 103 111 10 102 111 114 46 103 111 10 112 114 105 110 116 46 103 111 10 115 101 114 118 101 114 46 103 111 10]
// 	// err : <nil>
// 	// fmt.Println(output)
// 	// fmt.Println(err)

// 	// 아래 코드 안됨
// 	// fmt.Println(stirng([46 10 46 46 10 99 111 109 109 97 110 100 46 103 111 10 102 111 114 46 103 111 10 112 114 105 110 116 46 103 111 10 115 101 114 118 101 114 46 103 111 10]))

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(string(output))
// 	}

// }
