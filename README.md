# 🧪 NeoCore PoC – Arquitetura com gRPC e RabbitMQ

Este repositório contém uma **Prova de Conceito (PoC)** com **4 microserviços Go**, divididos entre comunicação **gRPC** e **RabbitMQ**, com o objetivo de comparar **escalabilidade e desempenho** entre os dois modelos.

---

## 🧩 Estrutura do Projeto

```
.
├── micro-api             → Serviço gRPC (Produtor principal)
├── micro-extrator        → Serviço gRPC (Consumidor)
├── micro-queue           → Serviço RabbitMQ (Produtor)
├── micro-queue-consumer  → Serviço RabbitMQ (Consumidor)
└── pkg
    └── proto             → Definições `.proto` e código gerado
```

### Comunicação
| Tipo | Serviço | Descrição |
|------|----------|-----------|
| gRPC | `micro-api` → `micro-extrator` | Comunicação direta via RPC |
| RabbitMQ | `micro-queue` → `micro-queue-consumer` | Comunicação assíncrona via fila |

---

## 🚀 Objetivo

A PoC visa comparar a **escalabilidade**, **latência** e **simplicidade operacional** entre:
- **gRPC**, para chamadas diretas e rápidas entre serviços.
- **RabbitMQ**, para processamento assíncrono e desacoplado.

---

## ⚙️ Requisitos

- **Go 1.22+**
- **Docker** (para subir o RabbitMQ)
- **RabbitMQ** rodando localmente na **porta 5673**

---

## 🐇 Subindo o RabbitMQ

```bash
docker run -d   --name rabbitmq   -p 5673:5672   -p 15673:15672   -e RABBITMQ_DEFAULT_USER=guest   -e RABBITMQ_DEFAULT_PASS=guest   rabbitmq:3-management
```

Interface de administração: [http://localhost:15673](http://localhost:15673)

---

## ▶️ Rodando os Serviços

Cada microserviço pode ser executado de forma independente:

```bash
cd micro-api
go mod tidy
go run cmd/main.go
```

Repita o mesmo processo para os outros diretórios:
- `micro-extrator`
- `micro-queue`
- `micro-queue-consumer`

---

## 🧠 Estrutura de Domínio

Todos os microserviços seguem uma arquitetura inspirada em **DDD (Domain-Driven Design)**, separando responsabilidades em:

- `model/` – Entidades e estruturas de dados  
- `repository/` – Acesso e persistência de dados  
- `service/` – Regras de negócio  
- `worker/` – Consumidores ou workers AMQP (quando aplicável)

---

## 📡 Tecnologias Utilizadas

- Go 1.22+  
- gRPC  
- Protocol Buffers  
- RabbitMQ  
- Docker

---

## 📈 Objetivo Final

Verificar **qual abordagem (gRPC ou RabbitMQ)** se comporta melhor em termos de:
- Escalabilidade
- Latência de comunicação
- Facilidade de manutenção
- Simplicidade operacional

---

🧠 **Autor:** Iago Holek  
📅 **Propósito:** Estudo de arquitetura distribuída em Go  
