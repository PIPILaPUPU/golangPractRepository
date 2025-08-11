package newrequest

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func req() {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/time", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
