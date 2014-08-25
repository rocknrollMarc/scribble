package main

import(
	"os"
	"fmt"
	//"net/http"
	//"github.com/zenazn/goji"
	//"github.com/zenazn/goji/web"
)

var environmentVariables []string

func init() {
	environmentVariables = os.Environ()
}

func main() {

}
