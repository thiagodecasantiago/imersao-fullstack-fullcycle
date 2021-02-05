# Desafio 02 - Consumo e publicação de mensagens no Kafka

> #Informações do desafio
>
>Nesse desafio você deverá criar uma simples aplicação utilizando Golang que seja capaz de publicar uma mensagem em um tópico do Apache Kafka e também consumir mensagens desse tópico.
>
>Para deixar mais claro a separação das responsabilidades, crie uma pasta chamada producer e outra consumer. Em cada uma delas crie um arquivo main.go que será responsável por produzir e consumir as mensagens respectivamente

## Setup

 - Clone repo:
 
  ```git clone https://github.com/thiagodecasantiago/imersao-fullstack-fullcycle.git```
 - Change to directory 'desafio02':
 
  ```mv desafio02```
 - Create containers:
 
 ```docker-compose up -d```
 - Enter desafio02_app_1 container:
 
 ```docker exec -it desafio02_app_1 bash```

## Usage

### Producer

- Publish message YOUR MESSAGE:

```go run producer/main.go "YOUR MESSAGE"```

### Consumer

- Run consumer:

```go run consumer/main.go```

- Quit pressing `CTRL + C`
