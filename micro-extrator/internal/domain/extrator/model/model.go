package model

import (
	"fmt"
)

type Exame struct {
	NomeProcedimento string
	Funcao           string
	QtdeApresentada  int
	QtdePaga         int
	Custo            float64
	Filme            bool
	Honorario        float64
	Glosa            float64
	CodGlosa         string
	Total            float64
	Executante       string
	LocalAtendimento string
	PrestadorArquivo string
	Acomodacao       string
	HoraRealizacao   string
	Acrescimo        float64
	PercentualVia    float64
}

type ExameUseCase struct{}

func NewExameUseCase() *ExameUseCase {
	return &ExameUseCase{}
}

// Process simula processamento de CSV e retorna exames
func (u *ExameUseCase) Process(filePath string) ([]Exame, error) {
	fmt.Println("Processando arquivo:", filePath)

	exames := []Exame{
		{
			NomeProcedimento: "Exame de sangue",
			Funcao:           "Rotina",
			QtdeApresentada:  1,
			QtdePaga:         1,
			Custo:            50.0,
			Total:            50.0,
			Executante:       "Dr. Fulano",
		},
	}

	return exames, nil
}
