package service

import (
	"encoding/json"
	"fmt"
)

type Results struct{
	WrapperType string `json:"wrapperType"`
	Kind string `json:"kind"`
	TrackId int64 `json:"trackId"`
	ArtistName string `json:"artistName"`
	TrackName string `json:"trackName"`
	TrackPrice float64 `json:"trackPrice"`
	TrackRenalPrice float64 `json:"trackRentalPrice"`
	Country string `json:"country"`
	Currency string `json:"currency"`
	PrimaryGenreName string `json:"primaryGenreName"`
	HasITunesExtras bool `json:"hasITunesExtras"`
	LongDescription string `json:"longDescription"`
}

type apiJSON struct {
	ResultCount int64 `json:"resultCount"`
	Results []Results `json:"results"`
}

// "/apijson"路由解析jsonData
// TODO 接收传入的JSON
func Unmarshalapijson(jsonData []byte){

	//jsonData := []byte(`
	//{
	//	"resultCount": 1,
	//	"results": [{
	//		"wrapperType": "track",
	//		"kind": "feature-movie",
	//		"trackId": 1493208771,
	//		"artistName": "Sam Mendes",
	//		"trackName": "1917",
	//		"trackCensoredName": "1917",
	//		"trackViewUrl": "https://itunes.apple.com/tw/movie/1917/id1493208771?uo=4",
	//		"previewUrl": "https://video-ssl.itunes.apple.com/itunes-assets/Video123/v4/cb/3e/f8/cb3ef81f-5ee7-7d10-fb56-0b52e3b69e7e/mzvf_14336057711372979641.640x356.h264lc.U.p.m4v",
	//		"artworkUrl30": "https://is3-ssl.mzstatic.com/image/thumb/Video123/v4/d9/02/8c/d9028c2e-156e-2001-d3f0-4b1218df1f25/source/30x30bb.jpg",
	//		"artworkUrl60": "https://is3-ssl.mzstatic.com/image/thumb/Video123/v4/d9/02/8c/d9028c2e-156e-2001-d3f0-4b1218df1f25/source/60x60bb.jpg",
	//		"artworkUrl100": "https://is3-ssl.mzstatic.com/image/thumb/Video123/v4/d9/02/8c/d9028c2e-156e-2001-d3f0-4b1218df1f25/source/100x100bb.jpg",
	//		"collectionPrice": 90.00,
	//		"trackPrice": 90.00,
	//		"trackRentalPrice": 80,
	//		"collectionHdPrice": 90,
	//		"trackHdPrice": 90,
	//		"trackHdRentalPrice": 80,
	//		"releaseDate": "2020-09-20T07:00:00Z",
	//		"collectionExplicitness": "notExplicit",
	//		"trackExplicitness": "notExplicit",
	//		"trackTimeMillis": 7138623,
	//		"country": "TWN",
	//		"currency": "TWD",
	//		"primaryGenreName": "劇情片",
	//		"contentAdvisoryRating": "輔15級",
	//		"longDescription": "《007：空降危機》、《007：惡魔四伏》、《美國心玫瑰情》奧斯卡金獎導演山姆·曼德斯，用他的獨特視角執導這部第一次世界大戰史詩鉅片。在一戰情勢最緊繃之際，兩名年輕的英國士兵史考菲（《神奇大隊長》喬治·麥凱飾）和布雷克（《冰與火之歌：權力遊戲》迪恩查爾斯·查普曼 飾），奉命執行一項看似不可能的任務。他們必須跟時間賽跑，越過敵區，傳達一個口信，讓布雷克的親兄弟在內的上千名士兵，逃過一場致命攻擊行動。",
	//		"hasITunesExtras": true
	//	}]
	//}`)
	var apiJson apiJSON
	err := json.Unmarshal(jsonData, &apiJson)
	if err != nil {
		fmt.Printf("json unmarshal failed, err%v\n", err.Error())
	}
	fmt.Print(apiJson.ResultCount, apiJson.Results)
}
