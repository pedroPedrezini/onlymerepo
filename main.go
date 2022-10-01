package main

import "fmt"

type macro struct {
	gordura float64

	carbo    float64
	proteina float64
	calorias float64
}

type alimento struct {
	nome             string
	quantidadeGramas float64
	macro
}

func (a alimento) GramaGordura() float64 {
	return a.quantidadeGramas * a.gordura
}

func (a alimento) GramaProteina() float64 {
	return a.quantidadeGramas * a.proteina
}

func (a alimento) GramaCarbo() float64 {
	return a.quantidadeGramas * a.carbo
}

func (a alimento) Calorias() float64 {
	return a.quantidadeGramas * a.calorias
}

func main() {

	frangoCozido := alimento{
		nome:             "frango cozido",
		quantidadeGramas: 100.0,
		macro: macro{
			gordura: 0.04,

			carbo:    0.0,
			proteina: 0.3,
			calorias: 2.0,
		},
	}

	paoDeForma := alimento{
		nome:             "Pão de forma",
		quantidadeGramas: 50.0,
		macro: macro{
			gordura: 0.03,

			carbo:    0.48,
			proteina: 0.08,
			calorias: 2.50,
		},
	}

	alimentosTodos := []alimento{frangoCozido, paoDeForma}

	func(a []alimento) {
		totalGordura := frangoCozido.GramaGordura() + paoDeForma.GramaGordura()
		fmt.Println("total de gorduras em", frangoCozido.quantidadeGramas, "gramas de frango e em", paoDeForma.quantidadeGramas, "gramas de pão de forma:", totalGordura)
		fmt.Println("")
		totalProt := frangoCozido.GramaProteina() + paoDeForma.GramaProteina()
		fmt.Println("total de proteinas em", frangoCozido.quantidadeGramas, "gramas de frango e em", paoDeForma.quantidadeGramas, "gramas de pão de forma:", totalProt)
		fmt.Println("")
		totalCarbo := frangoCozido.GramaCarbo() + paoDeForma.GramaCarbo()
		fmt.Println("total de carboidratos em", frangoCozido.quantidadeGramas, "gramas de frango e em", paoDeForma.quantidadeGramas, "gramas de pão de forma:", totalCarbo)
		fmt.Println("")
		totalCalorias := frangoCozido.Calorias() + paoDeForma.Calorias()
		fmt.Println("total de calorias em", frangoCozido.quantidadeGramas, "gramas de frango e em", paoDeForma.quantidadeGramas, "gramas de pão de forma:", totalCalorias)
		fmt.Println("")
	}(alimentosTodos)

}
