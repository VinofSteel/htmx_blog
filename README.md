# Templ Blog

Essa aplicação é uma aplicação de demonstração para a utilização da library [Templ](https://templ.guide/) como forma de gerar uma API que sirva HTML para um frontend coeso, utilizando Go como o motor da aplicação. A ideia foi conceptualizada pra a execução de um blog com a utilização da um componente de editor de texto, o [Quill.js](https://quilljs.com/) e criar um banco de dados em [PostgreSQL](https://www.postgresql.org/) relativamente complexo para a realização disso.

Nessa aplicação, é possível fazer todo o processo de autentiação esperado de um usuário (CRUD de usuário e lógica de Autenticação) e um CRUD de artigos de um blog que conta com diversas opções de customização.

## Tecnologias utilizadas:
- [Go](https://go.dev/) v1.22.3
  - [Fiber](https://gofiber.io/) v2.52.5
  - [Air](https://github.com/air-verse/air)
  - [sqlc](https://docs.sqlc.dev/en/latest/) v1.26.0
  - [Goose](https://github.com/pressly/goose) v3.21.1
  - [Validator](https://pkg.go.dev/github.com/go-playground/validator/v10) v10.22.0
  - [GoDotEnv](https://pkg.go.dev/github.com/joho/godotenv) v1.5.1
  - [Templ](https://templ.guide/) v0.2.747
- [PostgreSQL](https://www.postgresql.org/) v16.3
- [HTMX](https://htmx.org/) v2.0.0
- [Docker](https://www.docker.com/)

## Rodando a API:
Essa é uma API que foi construída para ser rodada utilizando o Docker, uma poderosa ferramenta que permite a utilização de imagens de sistemas operacionais inteiros para a execução de código de forma consistente. Especificamente, para rodar essa API, utilizamos o [Docker Compose](https://docs.docker.com/compose/), que permite a utilização de múltiplos containers em um único serviço. Para rodar nosso código, é só seguir os seguintes passos:
1. [Instale e configure](https://docs.docker.com/guides/getting-started/) o Docker.
2. [Instale e configure](https://docs.docker.com/compose/install/) o Docker Compose.
3. Crie um arquivo `.env` na raiz deste repositório e preencha as variáveis conforme o arquivo `.env.example`.
   1. Note que é extremamente importante preencher a variável `ENV` corretamente com o valor `development`, pois preencher ela com o valor `production` gera o executável para deployment e não permite o hot-reloading da aplicação.
4. Abra um terminal na raiz deste repositório.
5. Execute o comando `docker-compose up --build` 
   1. Note que este comando pode variar dependendo da sua instalação ou distribuição do Linux, utilizando Debian no WSL 2, o meu é `docker compose up --build`.
   2. Também é possível executar isso usando o [Docker Desktop](https://www.docker.com/products/docker-desktop/), mas visto como eu não o utilizo pessoalmente, não deixei isso nessa documentação.
6. Pronto! Caso você tenha configurado seu ambiente corretamente, você deverá ter 2 containers rodando, um do banco de dados e um da API em si. Em `development`, o container da API utiliza a ferramenta `Air` e graças as configurações do Docker Compose, o hot-reloading funciona. Em `production`, apenas é gerado um executável único que então é executado pela API, logo o hot-reloading não está disponível. É só acessar as rotas da API no seu browser favorito.
7. Quando tiver terminado a utilização, apenas dê um `ctrl + C` no terminal e execute o comando `docker-compose down` caso deseje deletar os containers.
   1. Executar o comando `docker-compose down` irá **DELETAR** os containers, o que significa que você perderá **QUALQUER MODIFICAÇÃO FEITA NO BANCO DE DADOS**. Se você planeja utilizar os dados além da execução desse momento, NÃO execute este comando.
   2. Note que, assim como o `docker-compose up` esse comando pode variar dependendo da sua instalação ou distribuição do Linux, utilizando Debian no WSL 2, o meu é `docker compose down`.
