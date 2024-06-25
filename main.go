package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        fmt.Println("API key not set")
        return
    }

    // Trim any leading or trailing whitespace from the API key
    apiKey = strings.TrimSpace(apiKey)
    fmt.Println("Using API key:", apiKey)


    url := "https://api.openai.com/v1/chat/completions"

    requestBody, err := json.Marshal(map[string]interface{}{
        "model": "gpt-3.5-turbo",
        "messages": []map[string]string{
            // {
            //     "role": "system",
            //     "content": "You are a helpful assistant.",
            // },
            // {
            //     "role": "user",
            //     "content": "Who won the world series in 2020?",
            // },
            // {
            //     "role": "assistant",
            //     "content": "The Los Angeles Dodgers won the World Series in 2020.",
            // },
            {
                "role": "user",
                "content": "Where was the last UCL final played?",
            },
        },
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
    if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
        message := choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
        fmt.Println("Generated Message:", message)
    } else {
        fmt.Println("No message generated")
    }
}
