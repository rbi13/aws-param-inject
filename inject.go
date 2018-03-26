package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"log"
	"os"
	"strings"
)

func main() {
	paths := os.Getenv("AWS_ENV_PATH")
	if  paths == "" {
		log.Println("missing AWS_ENV_PATH, aborting.")
		return
	}

	log.Println(paths)
	for _, path := range strings.Split(paths ,";") {
		ExportVariables(path, "")
	}
}

func CreateClient() *ssm.SSM {
	session := session.Must(session.NewSession())
	return ssm.New(session)
}

func ExportVariables(path string, nextToken string) {
	client := CreateClient()

	input := &ssm.GetParametersByPathInput{
		Path:           &path,
		WithDecryption: aws.Bool(true),
	}

	if nextToken != "" {
		input.SetNextToken(nextToken)
	}

	output, err := client.GetParametersByPath(input)

	if err != nil {
		log.Panic(err)
	}

	// for _, element := range output.Parameters {
	// 	PrintExportParameter(path, element)
	// }

	if output.NextToken != nil {
		ExportVariables(path, *output.NextToken)
	}
}

func PrintExportParameter(path string, parameter *ssm.Parameter) {
	name := *parameter.Name
	value := *parameter.Value

	env := strings.Trim(name[len(path):], "/")
	value = strings.Replace(value, "\n", "\\n", -1)

	fmt.Printf("export %s=$'%s'\n", env, value)
}
