package request

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func MakeRequest(filepath, bucketName string, headers map[string]string) (resp *http.Response, err error) {

	client := &http.Client{}
	contents, _ := ioutil.ReadFile(filepath)
	reader := bytes.NewReader(contents)

	request, err := http.NewRequest("PUT", "http://"+bucketName+".s3.amazonaws.com/"+filepath, reader)

	if err != nil {
		log.Fatal(err)
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	resp, err = client.Do(request)

	defer resp.Body.Close()

	return
}
