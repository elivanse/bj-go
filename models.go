package models

import ()

type org struct {
	id_org   string `json:"id_org"`
	desc_org string `json:"desc_org"`
}

type bolxanio struct {
	boletin_id_boletin int `json:"boletin_id_boletin"`
	anio               int `json:"anio"`
	numero             int `json:"numero"`
}

type digesto struct {
	boletin_id_boletin int `json:"boletin_id_boletin"`
	anio               int `anio:"anio"`
	id_org             int `id_org:"id_org"`
}

type User struct {
	Name          string    `json:"name"`
	Password      string    `json:"password"`
	PreferredFish []string  `json:"preferredFish"`
	CreatedAt     time.Time `json:"createdAt"`
}
