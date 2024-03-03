package main

import (
    "fmt"
    "strconv"
)

func main() {
    var (
        cpf string
        cpf_digitos [11]int
        res bool
    )

    fmt.Print("Digite o CPF à validar: ")
    fmt.Scanln(&cpf)

    if len(cpf) != 11 {
        fmt.Println("Tamanho do CPF inválido!")
        return
    }

    for i := 0; i < 11; i++ {
        digit, _ := strconv.Atoi(string(cpf[i]))
        cpf_digitos[i] = digit
    }

    res = ValidaCPF(cpf_digitos)

    if res == true {
        fmt.Println("✔️ CPF Válido!")
    } else {
        fmt.Println("❌ CPF Inválido!")
    }
}

func ValidaCPF(num [11]int) bool {
    var i, j, soma, resto, dv1, dv2 int

    // Digito verificador numero #1
    soma = 0
    j = 10

    for i = 0; i < 9; i++ {
        soma += num[i] * j
        j -= 1
    }

    resto = soma % 11
    if resto < 2 {
        dv1  = 0
    } else {
        dv1 = 11 - resto
    }

    // Digito verificador numero #2
    soma = 0
    j = 11

    for i = 0; i < 10; i++ {
        soma += num[i] * j
        j -= 1
    }

    resto = soma % 11
    if resto < 2 {
        dv2 = 0
    } else {
        dv2 = 11 - resto
    }

    if num[9] == dv1 && num[10] == dv2 {
        return true
    } else {
        return false
    }
}
