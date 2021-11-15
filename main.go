package main

import "fmt"
import "buffio"
import "os"



func main(){

	input := bufio.NewScanner(os.Stdin)

for input.Scan(){

	fmt.Println("Basic >",input.Text())

}



}
