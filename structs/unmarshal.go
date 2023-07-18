package Structs

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func FuncResponseArtists() ([]artists, int) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return []artists{}, http.StatusInternalServerError
	}

	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []artists{}, http.StatusInternalServerError
	}
	var responseArtists []artists
	err = json.Unmarshal(responseData, &responseArtists)
	if err != nil {
		return []artists{}, http.StatusInternalServerError
	}
	return responseArtists, http.StatusOK
}

func FuncResponseOneArtist(myid string) (artists, int) {
	responseArt, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + myid)
	if err != nil {
		return artists{}, http.StatusInternalServerError
	}
	defer responseArt.Body.Close()
	responseArtData, err := ioutil.ReadAll(responseArt.Body)
	if err != nil {
		return artists{}, http.StatusInternalServerError
	}
	var responseOneArist artists
	err = json.Unmarshal(responseArtData, &responseOneArist)
	if err != nil {
		return artists{}, http.StatusInternalServerError
	}
	return responseOneArist, http.StatusOK
}

func FuncResponseRelation(myid string) (relation, int) {
	responseRel, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + myid)
	if err != nil {
		return relation{}, http.StatusInternalServerError
	}
	defer responseRel.Body.Close()
	responseRelData, err := ioutil.ReadAll(responseRel.Body)
	if err != nil {
		return relation{}, http.StatusInternalServerError
	}
	var responseRelation relation

	err = json.Unmarshal(responseRelData, &responseRelation)
	if err != nil {
		return relation{}, http.StatusInternalServerError
	}
	return responseRelation, http.StatusOK
}
