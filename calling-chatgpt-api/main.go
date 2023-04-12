package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Preencha com o seu API key
	apiKey := "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"

	// Dados da mensagem para o ChatGPT
	input := "Canal do Youtube HunCoding"

	// Constrói a requisição HTTP POST
	requestBody, err := json.Marshal(map[string]interface{}{
		"model":       "text-davinci-003",
		"prompt":      input,
		"max_tokens":  4000,
		"temperature": 1.0,
	})
	if err != nil {
		fmt.Println("Erro ao construir corpo da requisição: ", err.Error())
		return
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Erro ao criar requisição: ", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Executa a requisição
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao executar requisição: ", err.Error())
		return
	}
	defer resp.Body.Close()

	// Lê a resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta: ", err.Error())
		return
	}

	// Exibe a resposta em formato JSON
	fmt.Println(string(body))
}
