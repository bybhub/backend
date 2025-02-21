## BYB BACKEND

### Entrypoint
```bash
go run cmd/server/main.go
```

## Começando no projeto
1. Garanta que você tem permissão para realizar o clone do repositório, vá até o diretório que deseja ter o projeto e execute o comando:
```bash
git clone git@github.com:bybhub/backend.git byb-backend
```

2. Certifique-se de ter o Go instalado e funcionando - [Download GO](https://go.dev/dl/).

3. Instale os pacotes e libs com o comando:
```bash
go mod tidy
```
4. Com tudo instalado você já pode começar a rodar o projeto com o comando:
```bash
go run cmd/server/main.go
```
## Importante
### Tenha o Docker instalado
O Docker será uma ferramenta fundamental para nosso desenvolvimento. Ele permite que tenhamos um container para nossa aplicação, proporcionando um ambiente controlado, único e replicável. Com isso, garantimos que o código seja executado da mesma forma em todas as máquinas locais e no servidor, minimizando problemas de compatibilidade e facilitando a colaboração entre os membros da equipe.

## Atualizações e Entregas
### PR's e Git Flow
Para garantir uma gestão eficiente do nosso código, manteremos o fluxo de pull requests dividido seguindo a metodologia Git Flow. Isso nos permitirá organizar melhor o desenvolvimento, facilitando a colaboração e a revisão do código. O Git Flow é uma abordagem popular que ajuda a gerenciar versões e a implementar novas funcionalidades de maneira estruturada. Para mais informações sobre como utilizar o Git Flow, consulte este artigo: [Gitflow](https://www.alura.com.br/artigos/git-flow-o-que-e-como-quando-utilizar?srsltid=AfmBOoryp2QatJYqUKRo5Y_Uwf-xIGkdwsy2N8E7k3sPLgR1tdfprtJb)

Em caso de dúvida, consulte o PR aberto na branch feature/example.