package main

import (
	"DesafioTecnico/database"
	"fmt"
	"net/http"

	
)

func main() {

/* 	err := repositorio.AbreSessao()
	if err != nil {
		fmt.Print("erro ao conectar no banco")
	} */
	database.ConnectMongo()
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Ola Mundo")
	})

	http.HandleFunc("/criarmoeda", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/editarmoeda", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/deletarmoeda", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/voteup", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/votedown", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("exibirmoedas", func(w http.ResponseWriter, r *http.Request) {

	})

}
