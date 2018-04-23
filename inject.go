package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"strings"
)

func main() {
	// inject ssm params
	paths := os.Getenv("AWS_ENV_PATH")
	if paths == "" {
		log.Println("missing AWS_ENV_PATH, aborting.")
		return
	}
	environ := append(*ExportVariables(paths), os.Environ()...)
	// passthrough execution
	binary, lookErr := exec.LookPath(os.Args[1])
	 if lookErr != nil {
		 binary, lookErr = exec.LookPath("sh")
		 if lookErr != nil {
				 panic(lookErr)
		 }
	 }
	execErr := syscall.Exec(binary, os.Args[1:], environ)
  if execErr != nil {
      panic(execErr)
  }
}

func ExportVariables(paths string) *[]string {
	// init client
	session := session.Must(session.NewSession())
	client := ssm.New(session)
	ret := []string{}
	// parse and load each path
	for _, path := range strings.Split(paths , ";") {
		input := &ssm.GetParametersByPathInput{
			Path:           &path,
			WithDecryption: aws.Bool(true),
		}
		// load params per path
		nextToken := ""
		for {
			if nextToken != "" {
				input.SetNextToken(nextToken)
			}
			output, err := client.GetParametersByPath(input)
			if err != nil {
				panic(err)
			}
			for _, param := range output.Parameters {
				key := strings.Trim((*param.Name)[len(path):], "/")
				ret = append(ret, fmt.Sprintf("%s=%s", key, *param.Value))
			}
			if output.NextToken == nil {
				break;
			}
			nextToken = *output.NextToken
		}
	}
	return &ret
}
