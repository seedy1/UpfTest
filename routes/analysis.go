package routes

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/seedy1/UpfTest/helpers"
	pkg "github.com/seedy1/UpfTest/packages"
)

// handleSSE handles the /analysis endpoint
func HandleSSE(writer http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		log.Println("handleSSE - Method not allowed, only GET is allowed on this endpoint.", http.StatusMethodNotAllowed)
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// get duration from request
	durationStr := req.URL.Query().Get("duration")
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		log.Println("handleSSE - Invalid duration", http.StatusBadRequest)
		http.Error(writer, "Invalid duration", http.StatusBadRequest)
		return
	}

	// get dimension from request which can be: likes comments, favorites, retweets
	dimensionStr := req.URL.Query().Get("dimension")
	if dimensionStr != "likes" && dimensionStr != "comments" && dimensionStr != "favorites" && dimensionStr != "retweets" {
		log.Println("handleSSE - Invalid dimension", http.StatusBadRequest)
		http.Error(writer, "Please enter a known dimension", http.StatusBadRequest)
		return
	}

	// read SSE server for given duration and returns an array of the data
	rawDataArray, err := readSSEServer(duration)
	if err != nil {
		log.Println("handleSSE - Error reading SSE server")
		http.Error(writer, "Error reading SSE server. Try again", http.StatusInternalServerError)
		return
	}

	dictionaryEnSv := make(map[int]pkg.Data)
	for i, dataString := range rawDataArray {
		var data pkg.Data
		if err := json.Unmarshal([]byte(dataString[6:]), &data); err != nil {
			log.Printf("Error parsing data: %v\n", err)
			continue
		} else {
			dictionaryEnSv[i] = data
		}
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	aggregateResult := pkg.AggResult(dictionaryEnSv, dimensionStr)
	aggregateResultJson, err := json.Marshal(aggregateResult)
	writer.Write(aggregateResultJson)
}

// reads SSE server for given duration and returns the data as a slice of strings
func readSSEServer(duration time.Duration) ([]string, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", helpers.SSEURL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Accept", "text/event-stream")
	ctx, cancel := context.WithTimeout(request.Context(), duration)
	defer cancel()
	request = request.WithContext(ctx)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Println("readSSEServer - Invalid status code:", response.StatusCode)
		return nil, errors.New("Invalid status code" + string(response.StatusCode))
	}

	// read the body
	var eventsStream []string
	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		eventData := scanner.Text()
		if eventData == "" {
			continue
		}
		eventsStream = append(eventsStream, eventData)
	}
	return eventsStream, nil
}
