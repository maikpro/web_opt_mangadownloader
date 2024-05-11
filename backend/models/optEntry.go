package models

/* {
"id":1307,
"name":"Patt",
"number":1113,
"category_id":3,
"arc_id":41,
"specials_id":0,
"lang":"ger",
"pages":15,
"is_available":true,
"date":"26.04.2024",
"href":"https://onepiece-tube.com/manga/kapitel/1113/1"
},
*/

type OPTEntry struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Number      int    `json:"number"`
	Pages       int    `json:"pages"`
	Date        string `json:"date"`
	IsAvailable bool   `json:"is_available"`
}
