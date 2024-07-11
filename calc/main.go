package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/carleesmarty/CalcLib"
)

type Handler struct {
	calculator CalcLib.Calculator
	out        io.Writer
}

func NewHandler(calculator CalcLib.Calculator, out io.Writer) *Handler {
	return &Handler{out: out, calculator: calculator}
}

func (h *Handler) Handle(args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("usage: calc <operator> <value> <value>")
	}

	num1, err1 := strconv.Atoi(args[1])
	num2, err2 := strconv.Atoi(args[2])

	if err1 != nil {
		return fmt.Errorf("invalid arguments: %s", os.Args[1])
	}
	if err2 != nil {
		return fmt.Errorf("invalid arguments: %s", os.Args[2])
	}

	result := h.calculator.Calculate(num1, num2)

	byteResult := []byte(strconv.Itoa(result))
	_, err := h.out.Write(byteResult)
	if err != nil {
		return fmt.Errorf("could not write to output: %s", err)
	}
	return err
}

func main() {

	//firstInt, err1 := strconv.Atoi(os.Args[1])
	//secondInt, err2 := strconv.Atoi(os.Args[2])
	//fmt.Println(os.Args[0], os.Args[1])
	//if err1 != nil || err2 != nil {
	//	panic("Not a valid int.")
	//}
	//fmt.Println(CalcLib.Addition{}.Calculate(firstInt, secondInt))

	args := []string{"add", "10", "67"}
	calculator := CalcLib.Addition{}

	handler := NewHandler(calculator, os.Stdout)
	err := handler.Handle(args)
	if err != nil {
		log.Printf("Error handling calculation: %s", err)
		os.Exit(1)
	}

}
