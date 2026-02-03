package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestUploads(t *testing.T) {
	url := "http://127.0.0.1:8080/upload"

	payload := strings.NewReader("-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"file\"\r\n\r\n[object Object]\r\n-----011000010111000001101001--\r\n\r\n")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NzAwOTc3MDMsImlhdCI6MTc3MDA5NDEwMywidXNlcklkIjoiMSJ9.tDUJzh9egGEIu-an9jmK0uvDTik0IMKcelxx27wI4Wo")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("User-Agent", "PostmanRuntime-ApipostRuntime/1.1.0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("content-type", "multipart/form-data; boundary=---011000010111000001101001")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
