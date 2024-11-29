package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "VitaoCLI"}

	var nome, email, senha string

	var cmd = &cobra.Command{
		Use:   "Create",
		Short: "Crie um novo usuário",
		Run: func(cmd *cobra.Command, args []string) {
			if nome == "" {
				fmt.Println("Nome não pode estar vazio.s")
			}

			if email == "" {
				fmt.Println("Email não pode estar vazio.")
			}

			if len(senha) < 6 {
				fmt.Println("A senha deve ter pelo menos 6 caracteres.s")
			}

			fmt.Printf("Nome: %s\nEmail: %s\nSenha %s\n", nome, email, senha)
		},
	}

	cmd.Flags().StringVarP(&nome, "nome", "n", "", "Nome do usuário")
	cmd.Flags().StringVarP(&email, "email", "e", "", "Email do usuário")
	cmd.Flags().StringVarP(&senha, "senha", "s", "", "Senha do usuário")

	rootCmd.AddCommand(cmd)
	rootCmd.Execute()

}
