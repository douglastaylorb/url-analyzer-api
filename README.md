# URL Analyzer API - Backend

Este é o componente de backend do projeto URL Analyzer. Ele é responsável por analisar URLs e gerar informações detalhadas, como o IP do servidor, métodos HTTP permitidos e links encontrados na página. Além disso, suporta a geração de relatórios em PDF.

---

## Funcionalidades

- **Análise de URLs:** Extração de informações como endereço IP, servidores e os links disponíveis na página analisada.
- **Geração de Relatórios:** Criação de relatórios formatados em PDF contendo os resultados da análise.
- **Rotas Estruturadas:** Organização e abstração de rotas, controladores e serviços para maior clareza do código.

---

## Tecnologias Utilizadas

- **Linguagem:** Go (Golang)
- **Framework:** Gin - Web Framework para gerenciamento de rotas
- **Banco de Dados:** SQLite ou outro compatível, com abstração no pacote `database/sql`
- **PDF:** `gofpdf` - Para geração de arquivos PDF
- **Docker Compose:** Facilita execução e preparo do ambiente

---

## Estrutura do Projeto

Segue a estrutura de diretórios do projeto:
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

---

## Configuração e Execução

### Pré-Requisitos

- [Go](https://golang.org/doc/install) >= 1.17

### Configuração Local

1. **Baixe e Instale o Repositório**

   Clone o repositório e instale as dependências:

   ```
   git clone https://github.com/seu-usuario/url-analyzer-api.git
   cd url-analyzer-api
   go mod tidy

2. **Exemplo de Requisição:**

  // json
  {
    "url": "https://example.com"
  }


Exemplo de Resposta:

json
{
  "ip": "127.0.0.1",
  "serverInfo": "nginx",
  "performanceMetrics": "120ms",
  "allowedMethods": "GET, POST",
  "hrefs": [
    "https://example.com/page1",
    "https://example.com/page2"
  ],
  "contentType": {
    "text/html": "90%",
    "image/png": "10%"
  }
}
GET /api/download-pdf
Gera e retorna um relatório no formato PDF para uma URL.

Parâmetros da URL:

url: A URL a ser analisada (Ex: https://example.com)
Resposta: Um arquivo PDF gerado para download.