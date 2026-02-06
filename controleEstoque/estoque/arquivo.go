package estoque

import (
	"encoding/json" // serve para codificar e decodificar dados em formato JSON
	"errors"        // serve para criar e manipular erros
	"os"            // serve para interagir com o sistema operacional (ler e escrever arquivos)
)

// RepositorioArquivo implementa o RepositorioEstoque armazenando produtos em um arquivo JSON
// Cada vez que um produto é adicionado ou atualizado, o arquivo é reescrito com o estado atual do estoque
type RepositorioArquivo struct {
	caminho string
}

// cria um repositório persistido em arquivo
func NovoRepositorioArquivo(caminho string) *RepositorioArquivo {
	return &RepositorioArquivo{
		caminho: caminho,
	}
}

func (r *RepositorioArquivo) Listar() []Produto {
	dados, err := os.ReadFile(r.caminho) // lê o conteúdo do arquivo
	if err != nil { //
	// arquivo ainda não existe -> estoque vazio
	return []Produto{} // retorna uma lista vazia de produtos	
	}

	var produtos []Produto // cria uma variável para armazenar os produtos lidos do arquivo
	json.Unmarshal(dados, &produtos) // decodifica os dados JSON para a variável produtos
	return produtos // retorna a lista de produtos 
}

func (r *RepositorioArquivo) Adicionar(produto Produto) {
	produtos := r.Listar() // lê os produtos atuais do arquivo
	produtos = append(produtos, produto) // adiciona o novo produto à lista
	
	dados, _ := json.MarshalIndent(produtos,"", " ") // codifica a lista de produtos em JSON com indentação
	os.WriteFile(r.caminho, dados, 0644) // os.WriteFile grava os dados no arquivo especificado pelo caminho, 0644 significa as permissões do arquivo
}

func (r *RepositorioArquivo) Atualizar(produto Produto) error {
	produtos := r.Listar() // lê os produtos atuais do arquivo

	for i := range produtos { // percorre a lista de produtos para encontrar o produto com o ID correspondente
		if produtos[i].ID == produto.ID {
			// atualiza o produto na lista
			produtos[i] = produto

			dados, _ := json.MarshalIndent(produtos, "", " ") // codifica a lista atualizada de produtos em JSON com indentação
			os.WriteFile(r.caminho, dados, 0644) // grava os dados no arquivo especificado pelo caminho
			
			return nil // retorna nil se a atualização for bem-sucedida}
		}
	}

	return errors.New("produto não encontrado") // retorna erro se o produto não for encontrado
}
