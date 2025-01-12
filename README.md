# URL Analyzer Backend

Este é o componente de backend do projeto URL Analyzer. Ele é responsável por analisar URLs e gerar informações detalhadas, como o IP do servidor, métodos HTTP permitidos e links encontrados na página. Além disso, suporta a geração de relatórios em PDF.

### Funcionalidades

- Análise de URLs: Extração de informações como endereço IP, servidores, e os links disponíveis na página analisada.
- Geração de Relatórios: Criação de relatórios formatados em PDF contendo os resultados da análise.
- Rotas Estruturadas: Organização e abstração de rotas, controladores e serviços para maior clareza do código.

### Tecnologias Utilizadas

- Linguagem: Go (Golang)
- Framework: Gin - Web Framework para gerenciamento de rotas
- Banco de Dados: SQLite ou outro compatível, com abstração no pacote database/sql
- PDF: gofpdf - Para geração de arquivos PDF
- Docker Compose - Facilita execução e preparo do ambiente

### Configuração e Execução

1. **Pré-requisitos**

   - Go instalado (versão >=1.16)
   - GOPATH configurado corretamente
  
2. **Instalando Dependências**
   
   Execute o seguinte comando para instalar as bibliotecas necessárias:

   ```
   go mod tidy
3. **Rodando o Servidor**

  Clone o reposítorio e instale as dependências: 
  
  ```
  git clone https://github.com/seu-usuario/url-analyzer-api.git
  cd url-analyzer-api
  go mod tidy
4. **Endpoints Disponíveis**
  - POST /api/analyze: Analisar uma URL fornecida no formato JSON.
  - GET /api/download-pdf?url={your_url}: Gera e baixa um relatório PDF para a URL fornecida.

  Estrutura do Projeto

  URL-ANALYZER-API
  ├── assets/                       # Fontes para PDF
  │   ├── Roboto-Bold.ttf           # Fonte Roboto (Negrito)
  │   └── Roboto-Regular.ttf        # Fonte Roboto (Regular)
  ├── controllers/                  # Diretório dedicado aos controladores
  │   └── urlController.go          # Lógica relacionada a rotas de URLs
  ├── database/                     # Configuração de database e conexão
  │   ├── db.go                     # Inicialização e manipulação CRUD
  │   └── models/
  │       └── url.go                # Estrutura do modelo de URL (abstração dos dados)
  ├── routes/                       # Declaração das rotas HTTP
  │   └── routes.go                 # Agrupamento de rotas e middlewares
  ├── services/                     # Contém serviços gerais reutilizáveis
  │   └── urlService.go             # Lógica principal da análise de URLs e PDFs
  ├── .env                          # Variáveis de ambiente
  ├── docker-compose.yml            # Definições Docker para rodar a aplicação
  ├── go.mod                        # Lista de dependências do projeto
  ├── go.sum                        # Checksum das dependências
  ├── main.go                       # Ponto de entrada principal
  └── README.md                     # Documentação

  Certifique-se de ter todas as fontes necessárias no diretório fonts.

  Contribuição
  Sinta-se à vontade para abrir issues e sugestões para melhoria do projeto.