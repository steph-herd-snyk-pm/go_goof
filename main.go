package main

import (
	"bufio"
	"encoding/xml" // this dependency has vulnerability CVE-2020-29509 for the used Go version 1.15.4 (not found by Snyk)
	"fmt"
	"github.com/gogo/protobuf/plugin/unmarshal" // this has a dependency on github.com/gogo/protobuf v1.2.1, which has vulnerability CVE-2021-3121 (not found by Snyk)
	"github.com/cucumber/godog/colors"
	"github.com/gin-gonic/gin"   // this dependency has vulnerability CVE-2020-28483 (found by Snyk)
	"golang.org/x/text/language" // this dependency has vulnerability CVE-2020-28851 for the used Go version 1.15.4 (not found by Snyk)
	"log"
	"os"
)

var opts = godog.Options{Output: colors.Colored(os.Stdout)}

func main() {
	_, _, err := language.ParseAcceptLanguage("ES-v-00-u-000-00")
	fmt.Println("Error:", err)

	var data interface{}
	if err := xml.NewDecoder(bufio.NewReader(os.Stdin)).Decode(&data); err != nil {
		panic(err)
	}
	log.Println(data)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"client_ip": c.ClientIP(),
		})
	})
}
