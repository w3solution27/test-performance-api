# Documentação Técnica

## **1. Resumo**
Este projeto é uma API desenvolvida em GoLang com foco em boas práticas de programação e arquitetura. Ele implementa um conjunto de rotas RESTful para gerenciar itens (CRUD) e foi projetado para servir de base para testes de performance usando o K6.

---

## **2. Tecnologias Utilizadas**
- **Linguagem:** GoLang (versão 1.20 ou superior)
- **Framework:** net/http nativo do Go
- **Testes de Performance:** K6
- **Estrutura Modular:**
  - **Controllers**: Gerencia a lógica de negócio.
  - **Routes**: Configuração de endpoints.
  - **Middleware**: Logger para registrar requisições.
  - **Services**: Contém a lógica da aplicação.
  - **Models**: Define as entidades e tipos de dados.

---

## **3. Estrutura do Projeto**
```plaintext
/project
├── main.go                 # Ponto de entrada da aplicação
├── go.mod                  # Gerenciador de dependências Go
├── routes/
│   └── routes.go           # Definição de rotas
├── middleware/
│   └── logger.go           # Middleware de logging
├── controllers/
│   └── item_controller.go  # Controladores
├── services/
│   └── item_service.go     # Lógica de negócio
├── models/
│   └── item.go             # Definição do modelo
├── k6/
│   └── test.js             # Script de teste de performance
```
---

## **4. Configuração do Ambiente**

### **4.1 Pré-requisitos**
Antes de começar, certifique-se de que os seguintes softwares estejam instalados:

GoLang (versão 1.20 ou superior)
K6 (para testes de performance)

### **4.2 Clonando o Repositório**
Clone o repositório do projeto:

```plaintext
git clone https://github.com/w3solution27/test-performance-api.git
cd project
Configure os módulos Go:

go mod tidy
5. Execução do Projeto
5.1 Inicializando o Servidor
Execute o seguinte comando para iniciar o servidor:

go run main.go
A aplicação estará disponível em: http://localhost:8080
```

## **6. Endpoints Disponíveis**

Base URL: http://localhost:8080/api/v1

GET	/items	Lista todos os itens
POST	/items	Cria um novo item

## **7. Testes de Performance**

### **7.1 Script de Teste**
O script para testes de performance está localizado em k6/test.js e foi configurado para simular requisições GET ao endpoint /items.

Exemplo de Script

```plaintext
import http from 'k6/http';
import { check, sleep } from 'k6';

export default function () {
  const res = http.get('http://localhost:8080/api/v1/items');
  check(res, { 'status is 200': (r) => r.status === 200 });
  sleep(1);
}
```

### **7.2 Execução do Teste**
Execute o script de teste com o comando:

```plaintext
k6 run k6/test.js
```

## **8. Boas Práticas Implementadas**
- Organização Modular: Código dividido em camadas (controllers, services, routes, middleware, models) para facilitar manutenção e expansão.

- Middleware Reutilizável: Logger implementado como middleware para ser aplicado a múltiplas rotas.

- Separação de Responsabilidades: Cada componente da aplicação possui funções específicas, promovendo clareza e coesão.

- Testes de Performance: O uso de K6 ajuda a validar a robustez e escalabilidade da aplicação.
