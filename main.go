package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type IGameRecord struct {
	Name          string
	Platform      string
	Medal         string
	Complete_time int
	Genre         string
	//completed_dlcs []string
}

func main() {
	in_filename := "./assets/beaten_games.csv"
	out_filename := "./assets/beaten_games.json"

	data, error := ioutil.ReadFile(in_filename)

	if error != nil {
		fmt.Println("Error reading file. ", error.Error())
		return
	}

	r := csv.NewReader(strings.NewReader(string(data)))

	games := []IGameRecord{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		name := record[0]
		platform := record[1]
		medal := record[2]
		complete_time, err := strconv.Atoi(record[3])
		genre := record[4]
		var dlcs_str string = record[5]

		if dlcs_str != "" {
			fmt.Println(dlcs_str)
		}

		game_record := IGameRecord{
			Name:          name,
			Platform:      platform,
			Medal:         medal,
			Complete_time: complete_time,
			Genre:         genre,
		}

		//json_abc, err := json.Marshal(game_record)
		//log.Println("json: ", string(json_abc))

		games = append(games, game_record)
	}

	json_bytes, json_error := json.MarshalIndent(games, "", "   ")

	if json_error != nil {
		log.Fatal("Couldn't convert to json")
		return
	}
	write_result := ioutil.WriteFile(out_filename, json_bytes, 0644)

	if write_result == nil {
		fmt.Println("Successfuly wrote json!")
	} else {
		log.Fatal("Couldn't write json. ", write_result.Error())
	}
}
