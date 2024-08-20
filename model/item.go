package model

import "kevinsheeran/walrus-backend/constant/enums"

type Item struct {
	Title   string
	Name    string
	Type    enums.ItermType
	Value   string
	Options []string
}
