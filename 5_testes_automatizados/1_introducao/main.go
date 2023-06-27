package enderecos

import "strings"

// Verifica se um endereço tem um tipo valido e o retorna
func TiposDeEndereco(endereco string) string{
	tiposValidos := []string{"Rua", "Avenida", "Estrada", "Rodovia"}
	enderecoEmLetraMaiuscula := strings.ToUpper(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMaiuscula, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}
	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}
	return "Tipo Invalido"
}