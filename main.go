package main

import (
	"fmt"
	"time"
	"strconv"
)

type Procesos interface {
	Mostrar()
	Buscar() uint64
}

type ListaProcesos struct {
	Contenido []Procesos
}

type Proceso struct {
	Id, Incremento uint64
	Bandera bool
}

func (lp *ListaProcesos) Mostrar() {
	for i := 0; i < len(lp.Contenido); i = i + 1{
		go lp.Contenido[i].Mostrar()
	}
}

func (lp *ListaProcesos) Buscar(id uint64) int {
	var i int
	for ubicacion, _ := range lp.Contenido {
		if id == lp.Contenido[ubicacion].Buscar() {
			i = ubicacion
			break
		}
	}
	return i
}

func (lp *ListaProcesos) Eliminar(id uint64) {
	eliminar := lp.Buscar(id)
	
	if (eliminar == len(lp.Contenido) - 1) {
		lp.Contenido = append(lp.Contenido[:eliminar])
	} else {
		lp.Contenido = append(lp.Contenido[:eliminar], lp.Contenido[eliminar+1:]...)
	}
}

func (proceso *Proceso) Start() {
	proceso.Incremento = 0
	proceso.Bandera = false

	for {
		proceso.Incremento = proceso.Incremento + 1
		time.Sleep(time.Millisecond * 500)
	}
}

func (proceso *Proceso) Buscar() uint64 {
	return proceso.Id
}

func (proceso *Proceso) Mostrar() {
	proceso.Bandera = !proceso.Bandera

	for {
		fmt.Println("id " + strconv.FormatUint(proceso.Id, 10) + ": " + strconv.FormatUint(proceso.Incremento, 10))
		time.Sleep(time.Millisecond * 500)

		if proceso.Bandera {
			return
		}
	}
}

func main(){
	lista := new(ListaProcesos)
	var contador uint64 = 0
	stop := false
	var opc int

	for !stop {
		fmt.Println("1) Agregar Proceso")
		fmt.Println("2) Mostrar Procesos")
		fmt.Println("3) Eliminar Proceso")
		fmt.Println("0) Salir")
		fmt.Scan(&opc)
		switch(opc) {
			case 1:
				proceso := new(Proceso)
				proceso.Id = contador
				contador++
				go proceso.Start()
				lista.Contenido = append(lista.Contenido, proceso)
				break
			case 2:
				lista.Mostrar()
				fmt.Println()
				break
			case 3:
				var input uint64
				fmt.Println("Ingrese Id a eliminar: ")
				fmt.Scan(&input)
				lista.Eliminar(input)
				break
			case 0:
				stop = true
		}
	}
}