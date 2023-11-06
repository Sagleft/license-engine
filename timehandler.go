package license

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"time"

	swissknife "github.com/Sagleft/swiss-knife"
)

const (
	defaultTimeZone = "Europe/Moscow"
)

type timeHandler struct{}

type worldTimeAPIResponse struct {
	Time string `json:"utc_datetime"`
}

type timeAPIResponse struct {
	Time string `json:"dateTime"`
}

type worldClockAPIResponse struct {
	Time string `json:"currentDateTime"`
}

type geonamesAPIResponse struct {
	Time string `json:"time"`
}

func newTimeHandler() timeHandler {
	return timeHandler{}
}

type timeParserFunc func() (*time.Time, error)

func (h *timeHandler) getCurrentTime() (*time.Time, error) {
	handlers := []timeParserFunc{
		h.parseTimeFromWorldAPI,
		h.parseTimeFromTimeAPI,
		h.parseTimeFromWorldClockAPI,
		h.parseTimeFromGeoNamesAPI,
	}

	for i, handler := range handlers {
		timeResult, err := handler()
		if err != nil {
			log.Println("time handler #" + strconv.Itoa(i) + ": " + err.Error())
			continue
		}
		log.Println("use time result from handler #" + strconv.Itoa(i))
		return timeResult, nil
	}

	return nil, errors.New("all time API servers are offline?")
}

func (h *timeHandler) parseTimeFromWorldAPI() (*time.Time, error) {
	// API GET
	apiURL := "http://worldtimeapi.org/api/timezone/" + defaultTimeZone
	responseBytes, err := swissknife.HttpGET(apiURL)
	if err != nil {
		return nil, err
	}

	timeResult := worldTimeAPIResponse{}
	err = json.Unmarshal(responseBytes, &timeResult)
	if err != nil {
		return nil, errors.New("failed to unmarshal api response json: " + err.Error())
	}

	timeParsed, err := time.Parse(time.RFC3339Nano, timeResult.Time)
	if err != nil {
		return nil, errors.New("failed to parse time (api): " + err.Error())
	}

	return &timeParsed, nil
}

func (h *timeHandler) parseTimeFromTimeAPI() (*time.Time, error) {
	// API GET
	apiURL := "https://www.timeapi.io/api/Time/current/zone?timeZone=" + defaultTimeZone
	responseBytes, err := swissknife.HttpGET(apiURL)
	if err != nil {
		return nil, err
	}

	timeResult := timeAPIResponse{}
	err = json.Unmarshal(responseBytes, &timeResult)
	if err != nil {
		return nil, errors.New("failed to unmarshal api response json: " + err.Error())
	}

	// 2021-12-31T02:13:28.272809
	timeLayout := "2006-01-02T15:04:05.999999999"
	timeParsed, err := time.Parse(timeLayout, timeResult.Time)
	if err != nil {
		return nil, errors.New("failed to parse time (api): " + err.Error())
	}

	return &timeParsed, nil
}

func (h *timeHandler) parseTimeFromWorldClockAPI() (*time.Time, error) {
	// API GET
	apiURL := "http://worldclockapi.com/api/json/utc/now"
	responseBytes, err := swissknife.HttpGET(apiURL)
	if err != nil {
		return nil, err
	}

	timeResult := worldClockAPIResponse{}
	err = json.Unmarshal(responseBytes, &timeResult)
	if err != nil {
		return nil, errors.New("failed to unmarshal api response json: " + err.Error())
	}

	timeLayout := "2006-01-02T15:04Z"
	timeParsed, err := time.Parse(timeLayout, timeResult.Time)
	if err != nil {
		return nil, errors.New("failed to parse time (api): " + err.Error())
	}

	return &timeParsed, nil
}

func (h *timeHandler) parseTimeFromGeoNamesAPI() (*time.Time, error) {
	// API GET
	apiURL := "http://api.geonames.org/timezoneJSON?formatted=true&lat=55.753220&lng=37.622513&username=demo&style=full"
	responseBytes, err := swissknife.HttpGET(apiURL)
	if err != nil {
		return nil, err
	}

	timeResult := geonamesAPIResponse{}
	err = json.Unmarshal(responseBytes, &timeResult)
	if err != nil {
		return nil, errors.New("failed to unmarshal api response json: " + err.Error())
	}

	if timeResult.Time == "" {
		return nil, errors.New("failed to get timestamp from service, empty time found. please try again later")
	}

	timeParsed, err := time.Parse("2006-01-02 15:04", timeResult.Time)
	if err != nil {
		return nil, errors.New("failed to parse time (api): " + err.Error())
	}

	return &timeParsed, nil
}
