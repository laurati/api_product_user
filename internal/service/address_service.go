package service

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/laurati/api_product_user/internal/entity"
)

func GetAddressService(cep string) *entity.Address {

	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json")
	if err != nil {
		log.Printf("request err %e", err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("response err %e", err)
	}
	var addressInput entity.AddressInput
	err = json.Unmarshal(res, &addressInput)
	if err != nil {
		log.Printf("err parsing response %e", err)
	}

	return &entity.Address{
		Cep:        addressInput.Cep,
		Logradouro: addressInput.Logradouro,
		Bairro:     addressInput.Bairro,
		Localidade: addressInput.Localidade,
	}
}
