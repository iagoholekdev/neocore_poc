package model

import "fmt"

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

func ProcessExameFile(path string) error {
	// Aqui você faria parsing do CSV/XLS e retornaria erro se algo falhar
	fmt.Println("Processando arquivo:", path)
	return nil
}
