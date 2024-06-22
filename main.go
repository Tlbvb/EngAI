
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        fmt.Println("API key not set",apiKey)
        return
    }

    url := "https://api.openai.com/v1/audio/translations"

    requestBody, err := json.Marshal(map[string]interface{}{
        "model": "whisper-1",
		"prompt": "Once upon a time, in a land far, far away...",
		"max_tokens": 50,
	})
    if err != nil {
        fmt.Println("Error marshaling request body:", err)
        return
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    var result map[string]interface{}
    if err := json.Unmarshal(body, &result); err != nil {
        fmt.Println("Error unmarshaling response:", err)
        return
    }
	fmt.Println(result)
    //fmt.Println("Generated Text:", result["choices"].([]interface{})[0].(map[string]interface{})["text"])
}
