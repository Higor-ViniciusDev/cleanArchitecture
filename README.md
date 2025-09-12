# Clean Architecture (Go)

Repositorio destinado ao desafio FullCycle aplicando princípios de Clean Architecture em um serviço escrito em Go, com GraphQL (gqlgen), migrações de banco e containerização via Docker / Docker Compose.

---

## Índice
1. Visão Geral
2. Principais Tecnologias
3. Arquitetura (Camadas)
4. Estrutura de Pastas
5. Requisitos
6. Configuração de Ambiente (.env)

---

## 1. Visão Geral

O objetivo do projeto é demonstrar a implementação de uma aplicação Go organizada segundo os princípios de Clean Architecture:
- Independência de frameworks
- Separação entre domínios e detalhes de implementação
- Facilidade de manutenção e evolução
- Testabilidade

O projeto utiliza GraphQL (via `gqlgen`) para expor a API, migrações para versionar a estrutura do banco e Docker para empacotamento e orquestração dos serviços.

---

## 2. Principais Tecnologias

| Tecnologia | Uso |
|------------|-----|
| Go (Golang) | Core da aplicação |
| gqlgen | Geração de schema/servidor GraphQL |
| Docker / Docker Compose | Containerização |
| Mysql | Persistência |
| Migrations (pasta `migrations/`) | Evolução de schema |
| Clean Architecture | Estrutura de camadas |

---

## 3. Arquitetura (Camadas)

Estrutura típica utilizada:
- `usecase` : Orquestra lógica de casos de uso.
- `infra` / `repository`: Implementações concretas de persistência, adapters externos.
- `interfaces` : Controllers, GraphQL resolvers, DTOs.
- `configs`: Configuração (carregamento de variáveis / arquivos).
- `migrations`: Scripts de evolução do banco.
- `cmd/SistemaDeOrdem`: Ponto(s) de entrada.

---

## 4. Estrutura de Pastas (resumo)

```
.
├── cmd/
│   └── SistemaDeOrdem/        # (provável entrypoint citado no comando docker)
├── internal/                  # Código interno organizado por contexto/camada
├── api/                       # Definições da camada de interface (resolvers/handlers)
├── configs/                   # Arquivos de configuração
├── migrations/                # Migrações de banco
├── pkg/                       # Pacotes utilitários reutilizáveis
├── gqlgen.yml                 # Config do gqlgen
├── docker-compose.yaml        # Orquestração de serviços
├── Dockerfile                 # Build da aplicação
└── go.mod / go.sum            # Dependências Go
```

---

## 5. Requisitos

Antes de iniciar, garanta que tem instalado:

- Git
- Docker Engine + Docker Compose Plugin  
  Verifique:  
  ```bash
  docker -v
  docker compose version
  ```
- Go (caso deseje rodar sem Docker)  
  Verifique:  
  ```bash
  go version
  ```

Para descobrir o schema, consulte o arquivo `schema.graphqls` (provavelmente gerado/definido em `api` ou conforme indicado no `gqlgen.yml`).

---

## 6. Logs, Rebuild e Troubleshooting

- Reconstruir após mudanças no código:
  ```bash
  docker compose --env-file ./cmd/SistemaDeOrdem/app_config.env up --build -d
  ```

- Seguir logs de um serviço específico:
  ```bash
  docker compose logs -f app
  ```

- Acessar shell no container:
  ```bash
  docker compose exec app sh
  ```

Problemas comuns:
| Sintoma | Possível Causa | Ação |
|--------|----------------|------|
| Conexão recusada ao DB | Serviço ainda subindo | Aguardar / verificar logs |
| Migração falha | Credenciais erradas | Revisar variáveis no .env |
| Porta em uso | Outro processo ocupa 8080 | Alterar APP_PORT e recompôr |

---

## Resumo Rápido (TL;DR)

```bash
git clone https://github.com/Higor-ViniciusDev/cleanArchitecture.git
cd cleanArchitecture
# criar/ajustar ./cmd/SistemaDeOrdem/app_config.env
docker compose --env-file ./cmd/SistemaDeOrdem/app_config.env up --build -d
# acessar: http://localhost:8080/graphql
```

Ajuste quaisquer nomes de serviços, portas e variáveis conforme o conteúdo real do `docker-compose.yaml`.

Por padrão eu separei dois .env o local `.env` e o de produção `app_config.env` dentro de `cmd/SistemaDeOrdem/`. Sinta-se livre para ajustar conforme sua necessidade.
---

Portas Expostas:
- API GraphQL: `http://localhost:8080/graphql` (ajuste conforme `GRAPHQL_SERVER_PORTA`)
- Banco MySQL: `localhost:3306` (ajuste conforme `DB_PORTA`)
- GRPC: `localhost:50051` (ajuste conforme `GRPC_SERVER_PORTA`)
- Web Rest: `http://localhost:8000` (ajuste conforme `WEB_SERVER_PORTA`)
- RabbitMQ: `localhost:5672` (ajuste conforme `RABBITMQ_PORT`)

Marcha no progresso! 🚀