package cloud

// contains functionality for retrieving data from cloud buckets

import (
	"os"
	"io"
	"context"
	"strings"
	"log"
	//"github.com/gin-gonic/gin"
	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

const bucketName = "raw-digested-data"

const projectName = "crypto-dashboard-469901"

func ReadFromBucket(client *storage.Client, coin_name string, ctx context.Context) []string {

	// pulls data from the cloud bucket for that coin name

	bkt := client.Bucket(bucketName)

	// list all objects in the bucket
	query := &storage.Query{Prefix: coin_name + "_digested_data"}

	var names []string
	it := bkt.Objects(ctx, query)

	for {

		//log.Println(os.Getwd())

		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// check if the object is a directory
		if !strings.HasSuffix(attrs.Name, ".json") {
			continue
		}

		// for each object, download the file and write it to a local folder
		rc, err := bkt.Object(attrs.Name).NewReader(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer rc.Close()
		body, err := io.ReadAll(rc)
		if err != nil {
			log.Fatal(err)
		}

		// write the body to a file in local storage
		file, err := os.Create("./data/" + attrs.Name)
		if err != nil {
			log.Printf("Error creating file: %v", err)
			return nil
		}
		defer file.Close()

		_, err = file.Write(body)
		if err != nil {
			log.Printf("Error writing to file: %v", err)
			return nil
		}

		names = append(names, attrs.Name)
	}

	return names

}
