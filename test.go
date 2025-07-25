package main

import (
	"fmt"
	"os"
)

SSCS_COMMIT_HISTORY_ENABLED=true
AWS_ACCESS_KEY=AKIA1234567890
token=ghp_1234567890abcdef1234567890abcdef12345678
AWS_ACCESS_KEY=DKIA1234567890
AWS_ACCESS_KEY=ZYIA9834567890
// Convert the secrets above to const
const (
	SSCSCommitHistoryEnabled = true
	AWSAccessKey1        = "AKIA1234567890"
	Token               = "ghp_1234567890abcdef1234567890abcdef12345678"
	AWSAccessKey2		= "DKIA1234567890"
	AWSAccessKey3		= "ZYIA9834567890"
)

func main() {
	data, err := os.ReadFile(".env")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("%q\n", data)

	data2, err := os.ReadFile(".env.example")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("%q\n", data2)
}
