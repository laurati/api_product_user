package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/laurati/api_product_user/internal/entity"
)

func GetAddressService(cep string) (*entity.Address, error) {

	urlViaCep := fmt.Sprintf(`http://viacep.com.br/ws/%s/json`, cep)

	req, err := http.Get(urlViaCep)
	if err != nil {
		log.Printf("request err %e", err)
		return nil, err
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("response err %e", err)
		return nil, err
	}
	var addressInput entity.AddressInput
	err = json.Unmarshal(res, &addressInput)
	if err != nil {
		log.Printf("err parsing response %e", err)
		return nil, err
	}

	return &entity.Address{
		Cep:        addressInput.Cep,
		Logradouro: addressInput.Logradouro,
		Bairro:     addressInput.Bairro,
		Localidade: addressInput.Localidade,
	}, nil
}
