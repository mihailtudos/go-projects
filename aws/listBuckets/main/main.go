package main

import (
	"fmt"
	listBuckets "github.com/tudosm/go-projects/aws/list-buckets"
)

func main() {
	buckets, err := listBuckets.ListBuckets()

	if err != nil {
		fmt.Println("Error with listbuckets")
		panic(err)
	}

	for _, bucket := range buckets {
		fmt.Println(*bucket)
	}
}
