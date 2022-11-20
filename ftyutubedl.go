//package main
package ftyutubedl 

import (
	"fmt"
	//"log"
	"os/exec"
	"strings"
)

var testString string = "https://www.youtube.com/watch?v=zCIpWFYDJ8s"

func YouTubeDl(u string) {

	fmt.Printf("Downloading from Youtube: \n%s\n", u)
	command := exec.Command("yt-dlp", "--output", "%(id)s.%(ext)s", u)
	command.Dir = "/usr/local/Cellar/nginx/1.23.1/html"
	i, err := command.Output()
	if strings.Contains(string(i), "has already been downloaded") {
		fmt.Println("")
		fmt.Printf("Error: %s already downloaded\n",u)
	}else {
	fmt.Printf("\nDownload successful. \n")
	fmt.Printf("%s\n",u)
	}
	if err != nil {
		return
	}

}
func main() {
/*	cmd := exec.Command("ls", "-lah")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out--------?:\n%s\n", string(out))
	command := exec.Command("yt-dlp", "--output", "%(id)s.%(ext)s", testString)
	command.Dir = "/usr/local/Cellar/nginx/1.23.1/html"
	i, err := command.Output()
	fmt.Printf("OUTPUT---->\n%s\n", string(i))
	if strings.Contains(string(i), "has already been downloaded") {
		fmt.Println("Already downloaded")
	}
	if err != nil {
		return
	}
*/
	YouTubeDl(testString)
}
