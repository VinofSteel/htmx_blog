# Templ Blog

Essa aplicação é uma demonstração para a utilização da library [Templ](https://templ.guide/) como forma de gerar uma aplicação fullstack, utilizando Go como o motor da aplicação no backend e JavaScript vanilla no front. A ideia foi conceptualizada pra a execução de um blog com a utilização da um componente de editor de texto, o [Quill.js](https://quilljs.com/) e criar um banco de dados em [PostgreSQL](https://www.postgresql.org/) relativamente complexo para a realização disso. No desenvolvimento, também optei pela utilização do [TypeScript](https://www.typescriptlang.org/) para facilitar o desenvolvimento e adicionar segurança ao scripting feito no frontend.

Nessa aplicação, é possível fazer todo o processo de autenticação esperado de um usuário (CRUD de usuário e lógica de Autenticação) e um CRUD de artigos de um blog que conta com diversas opções de customização. Como dito no parágrafo acima, essa aplicação é apenas uma demonstração de um sistema de edição de texto que será utilizado em outro site, mas é fundamental para a familiarização com as bibliotecas escolhidas.

## Tecnologias utilizadas:
- [Go](https://go.dev/) v1.22.3
  - [Fiber](https://gofiber.io/) v2.52.5
  - [Air](https://github.com/air-verse/air) v1.52.3
  - [sqlc](https://docs.sqlc.dev/en/latest/) v1.26.0
  - [Goose](https://github.com/pressly/goose) v3.21.1
  - [Validator](https://pkg.go.dev/github.com/go-playground/validator/v10) v10.22.0
  - [GoDotEnv](https://pkg.go.dev/github.com/joho/godotenv) v1.5.1
  - [Templ](https://templ.guide/) v0.2.747
- [PostgreSQL](https://www.postgresql.org/) v16.3
- [Node](https://nodejs.org/) v20.17.0
  - [TypeScript](https://www.typescriptlang.org/) v5.5.4
  - [ESBuild](https://esbuild.github.io/) v0.23.1
  - [Quill](https://quilljs.com/) v2.0.2
  - [Axios](https://axios-http.com/) v1.7.4
- [Docker](https://www.docker.com/)

## Rodando a API:
Essa é uma API que foi construída para ser rodada utilizando o Docker, uma poderosa ferramenta que permite a utilização de imagens de sistemas operacionais inteiros para a execução de código de forma consistente. Especificamente, para rodar essa API, utilizamos o [Docker Compose](https://docs.docker.com/compose/), que permite a utilização de múltiplos containers em um único serviço. Para rodar nosso código, é só seguir os seguintes passos:
1. [Instale e configure](https://docs.docker.com/guides/getting-started/) o Docker.
2. [Instale e configure](https://docs.docker.com/compose/install/) o Docker Compose.
3. Crie um arquivo `.env` na raiz deste repositório e preencha as variáveis conforme o arquivo `.env.example`.
   1. Note que é extremamente importante preencher a variável `ENV` corretamente com o valor `development`, pois preencher ela com o valor `production` gera o executável para deployment sem utilizar o Docker Compose e não permite o hot-reloading da aplicação.
4. Abra um terminal na raiz deste repositório.
5. Execute o comando `docker-compose up --build` 
   1. Note que este comando pode variar dependendo da sua instalação ou distribuição do Linux, utilizando Debian no WSL 2, o meu é `docker compose up --build`.
   2. Também é possível executar isso usando o [Docker Desktop](https://www.docker.com/products/docker-desktop/), mas visto como eu não o utilizo pessoalmente, não deixei isso nessa documentação.
6. Caso você tenha configurado seu ambiente corretamente, você deverá ter 2 containers rodando, um do banco de dados e um da API em si (rode o comando `docker container ls -a` para verificar). Em `development`, o container da API utiliza a ferramenta `Air` e graças as configurações do Docker Compose, o hot-reloading funciona. Em `production`, apenas é gerado um executável único que então é executado pela API, logo o hot-reloading não está disponível. É só acessar as rotas da API no seu browser favorito.
7. Caso você esteja em ambiente de desenvolvimento, as migrations do banco de dados não são rodadas automaticamente. No diretório raiz do repositório, rode o comando `make m-up` e elas serão rodadas no seu banco de dados (novamente, não se esqueça de preencher o seu `.env` antes de fazer isso, ou nada vai funcionar). Caso você utilize Windows e não tenha o `Make` instalado para rodar comandos do `Makefile`, eu recomendo a instalação [nesse link](https://gnuwin32.sourceforge.net/packages/make.htm) já que ele é essencial para o desenvolvimento em Go. Se ainda assim você você preferir não instalar, terá que montar a sua string de conexão do PostgreSQL manualmente e rodar o comando `goose -dir sql/schema postgres (PG_CONN_STRING) up` no seu terminal.
8. Quando tiver terminado a utilização, apenas dê um `ctrl + c` no terminal e execute o comando `docker-compose down` caso deseje deletar os containers. Como a aplicação possui um volume, os dados não serão perdidos no seu banco a não ser que você delete o volume também.
