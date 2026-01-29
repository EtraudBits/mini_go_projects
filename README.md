# Mini Projetos em Go 🚀

Repositório de mini projetos desenvolvidos em **Go (Golang)** para fins de estudo e aprimoramento de habilidades em programação.

## 📚 Sobre

Esta coleção contém diversos mini projetos que exploram diferentes aspectos da programação, com foco em:

- **Lógica de Programação** - Resolução de problemas e algoritmos
- **Estruturas de Dados** - Implementação e manipulação de dados
- **Organização de Código** - Clean Architecture e boas práticas
- **Padrões de Design** - Design patterns e princípios SOLID
- **Testes Unitários** - TDD e garantia de qualidade
- **CLI Development** - Aplicações de linha de comando
- **Validação e Tratamento de Erros** - Error handling robusto

## 🎯 Objetivos

- 📖 Aprender e praticar conceitos fundamentais de Go
- 🏗️ Desenvolver habilidades de arquitetura de software
- 🧪 Aplicar testes unitários e TDD
- 📝 Criar código limpo, legível e bem documentado
- 🔧 Explorar as bibliotecas padrão do Go
- 💡 Resolver problemas práticos do dia a dia

## 🗂️ Estrutura dos Projetos

Cada mini projeto neste repositório segue uma estrutura organizada e profissional:

```
mini_go_projects/
├── projeto1/
│   ├── README.md              # Documentação específica do projeto
│   ├── go.mod                 # Módulo Go
│   ├── cmd/                   # Código executável
│   │   └── main.go
│   └── internal/              # Lógica de negócio
│       └── package/
│           ├── *.go           # Implementação
│           └── *_test.go      # Testes
├── projeto2/
│   └── ...
└── README.md                  # Este arquivo
```

## 📋 Padrões Adotados

Todos os projetos seguem boas práticas de desenvolvimento:

### ✅ Organização

- Separação clara entre entrada/saída e lógica de negócio
- Uso do diretório `internal/` para código não exportável
- Módulos Go independentes para cada projeto

### ✅ Qualidade de Código

- Código comentado e documentado
- Funções pequenas e com responsabilidade única
- Nomenclatura clara e descritiva
- Tratamento adequado de erros

### ✅ Testes

- Testes unitários para cada funcionalidade
- Table-driven tests quando apropriado
- Cobertura de casos de sucesso e erro
- Testes de edge cases

### ✅ Documentação

- README.md detalhado em cada projeto
- Comentários explicativos no código
- Instruções de uso e exemplos
- Descrição da lógica implementada

## 🛠️ Tecnologias e Conceitos

### Linguagem

- **Go 1.22+** - Linguagem principal

### Práticas de Engenharia

- Clean Architecture
- Separation of Concerns
- SOLID Principles
- Design Patterns
- Error Handling Patterns

### Ferramentas

- Go Modules - Gerenciamento de dependências
- Go Testing - Framework de testes nativo
- Git & GitHub - Controle de versão

## 🚀 Como Usar

### Pré-requisitos

- Go 1.22 ou superior instalado
- Git para clonar o repositório

### Clonando o Repositório

```bash
git clone https://github.com/EtraudBits/mini_go_projects.git
cd mini_go_projects
```

### Executando um Projeto

```bash
# Navegue até o diretório do projeto
cd nome_do_projeto/cmd

# Execute o projeto
go run .
```

### Rodando os Testes

```bash
# Testes de um projeto específico
cd nome_do_projeto
go test ./...

# Testes com cobertura
go test ./... -cover

# Testes detalhados
go test ./... -v
```

## 📖 Aprendizado Contínuo

Este repositório é um **portfólio vivo** de aprendizado, onde:

- 🔄 Novos projetos são adicionados regularmente
- 📈 Projetos existentes são refatorados conforme evolução do conhecimento
- 🎓 Cada projeto explora um conceito ou técnica diferente
- 💬 Código é documentado para facilitar revisão futura

## 🎓 Tópicos Abordados

Os projetos neste repositório cobrem uma variedade de tópicos:

- **Fundamentos**: Tipos de dados, estruturas de controle, funções
- **Estruturas de Dados**: Arrays, slices, maps, structs
- **Programação Orientada a Objetos**: Métodos, interfaces, composição
- **Concorrência**: Goroutines, channels (projetos futuros)
- **Entrada/Saída**: Manipulação de arquivos, stdin/stdout
- **Validação**: Expressões regulares, validadores customizados
- **Algoritmos**: Ordenação, busca, manipulação de strings
- **Padrões**: Factory, Strategy, Repository (conforme aplicável)

## 🌟 Diferenciais

- 📝 **Documentação Completa**: Cada projeto possui README detalhado
- 🧪 **Testes Abrangentes**: Cobertura de casos normais e extremos
- 🏗️ **Arquitetura Limpa**: Código organizado e modular
- 💬 **Código Comentado**: Explicações em português para facilitar compreensão
- 🎯 **Foco Educacional**: Prioridade em aprendizado e boas práticas

## 📊 Status dos Projetos

Cada projeto possui seu próprio ciclo de desenvolvimento:

- ✅ **Completo**: Totalmente implementado e testado
- 🚧 **Em Desenvolvimento**: Implementação em andamento
- 📝 **Planejado**: Próximos projetos a serem desenvolvidos

## 🤝 Contribuições

Este é um repositório pessoal de estudos, mas sugestões e feedbacks são sempre bem-vindos!

## 📄 Licença

Este projeto é de código aberto para fins educacionais.

## 👨‍💻 Autor

Desenvolvido por **Duarte** como parte do portfólio de aprendizado em Go.

---

**Nota**: Este repositório está em constante evolução. Novos projetos e melhorias são adicionados regularmente conforme o progresso nos estudos de Go e engenharia de software.

🔗 **GitHub**: [EtraudBits/mini_go_projects](https://github.com/EtraudBits/mini_go_projects)
