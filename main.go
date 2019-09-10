package main

import (
	"fmt"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"crypto/md5"
	"encoding/hex"
)

func hash_file_md5(filePath string) (string, error) {
	var returnMD5String string

	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}

	defer file.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	hashInBytes := hash.Sum(nil)[:16]

	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil
}

func main() {
	fmt.Println("Yo whats up3")
	// Read our config options from the config.json file

	// Read file
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Print(err)
	}
	// Define data structure
	type config_options struct {
		INPUT string
		OUTPUT string
	}

	var obj config_options

	// unmarshall it
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("input dir: %s\n", obj.INPUT)
	fmt.Printf("output dir: %s\n", obj.OUTPUT)

	// For each file in our INPUT Dir we want to get the following info
	// full path + name to file
	// size
	// md5 hash

	type FILEDATA struct {
		SIZE int64
		MD5 string
	}

	filemap := make(map[string]FILEDATA)


	err1 := filepath.Walk(obj.INPUT, func(path string, info os.FileInfo, err1 error) error {
		if err1 != nil {
			return err1
		}
		// info.Size() is file size in bytes, file name doesn't change this
		fmt.Println(path, info.Size())
		// get the md5 hash of the file
		hash, err3 := hash_file_md5(path)
		if err3 == nil {
			filemap[path] = FILEDATA{info.Size(), hash}
		}
		if err3 != nil {
			fmt.Println("Hash error: %s", err3)
		}
		return nil
	})
	if err1 != nil {
		log.Println(err1)
	}

	fmt.Printf("%+v\n", filemap)
}