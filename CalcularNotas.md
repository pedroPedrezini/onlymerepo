package main

import (
	"fmt"
	"sort"
)

type notasMatematica struct {
	primeiroTri int
	segundoTri  int
	terceiroTri int
}

func soma(n []notasMatematica) {
	for _, v := range n {
		total := v.primeiroTri + v.segundoTri + v.terceiroTri
		falta := 180 - total
		if total > 180 {
			fmt.Println("Aprovado")
		} else {
			fmt.Println("reprovado falta: ", falta)
		}
	}
}

type porNota []notasMatematica

func (p porNota) Len() int           { return len(p) }
func (p porNota) Less(i, j int) bool { return p[i].primeiroTri+p[j].segundoTri+p[i].terceiroTri < 180 }
func (p porNota) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {

	pedro := notasMatematica{
		primeiroTri: 96,
		segundoTri:  0,
		terceiroTri: 82,
	}

	roberto := notasMatematica{
		primeiroTri: 56,
		segundoTri:  110,
		terceiroTri: 66,
	}
	alunos := []notasMatematica{pedro, roberto}
	sort.Sort(porNota(alunos))
	soma(alunos)

}
