package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	colors "github.com/daviddengcn/go-colortext"
)

const (
	// No change of color
	none = colors.Color(iota)
	black
	red
	green
	yellow
	blue
	magenta
	cyan
	white
)

type project struct {
	contributors                    int
	margin, profit, expenses, total float64
}

type contributor struct {
	hours int
	price float64
}

func main() {
	project := new(project)
	project.askInput()
	project.total = project.expenses * project.margin
	project.profit = project.total - project.expenses

	project.output()
}

func (p *project) askInput() {

	colors.Foreground(blue, true)

	fmt.Print("Número de colaboradores: ")
	_, err := fmt.Scanln(&p.contributors)

	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= p.contributors; i++ {

		c := new(contributor)

		question1 := fmt.Sprintf("Remuneração/hora colaborador %d: ", i)
		question2 := fmt.Sprintf("Número de horas colaborador %d: ", i)

		if p.contributors == 1 {
			question1 = "Remuneração/hora do colaborador: "
			question2 = "Número de horas do colaborador: "
		}

		fmt.Print(question1)
		_, err := fmt.Scanln(&c.price)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(question2)
		_, err = fmt.Scanln(&c.hours)
		if err != nil {
			log.Fatal(err)
		}

		p.expenses += float64(c.hours) * c.price
	}

	fmt.Print("Margem K: ")
	var m float64
	_, err = fmt.Scanln(&m)
	p.margin = float64(m)/100 + 1

	if err != nil {
		log.Fatal(err)
	}
	colors.ResetColor()
}

func (p *project) output() {

	colors.Foreground(red, true)
	fmt.Printf("\nCustos: %.2f\n", p.expenses)
	colors.ResetColor()

	colors.Foreground(green, true)
	fmt.Printf("Margem: %.2f\n", p.profit)
	colors.ResetColor()

	colors.Foreground(white, true)
	fmt.Printf("Total: %.2f\n\n", p.total)

	colors.ResetColor()
	colors.Foreground(blue, true)
	fmt.Print("Prima enter para sair...")
	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
}
