package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`        //フィールドタグkey:"value"形式
	Color  bool `json:"color,omitempty"` //omitemptyはフィールドがゼロ値の場合は出力しない、という指定
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

//Goのデータ構造をJSONへ変換することをマーシャリングという
func main() {
	//公開されているフィールドだけ出力される
	// data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", "    ") //読みやすい形

	if err != nil {
		log.Fatalf("JSON marchaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	//アンマーシャリング、タイトル情報だけ取得してその他はいらない
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("%s\n", titles)
}
