# Go Мікросервіси з REST API та gRPC

## Опис

Цей проект демонструє розробку мікросервісів на Go з використанням REST API та gRPC для комунікації, конкурентних
можливостей Go, тестування без зовнішніх бібліотек, деплойменту за допомогою Docker та Kubernetes, а також застосування
принципів TDD.

## Структура Проекту

````
├── api
│   ├── order.proto
│   └── user.proto
├── services
│   ├── order
│   │   ├── main.go
│   │   ├── order.go
│   │   ├── order_test.go
│   │   └── Dockerfile
│   └── user
│       ├── main.go
│       ├── user.go
│       ├── user_test.go
│       └── Dockerfile
├── k8s
│   ├── order-deployment.yaml
│   ├── user-deployment.yaml
│   └── service.yaml
├── docker-compose.yml
└── README.md
````

## Інструкція по Запуску

# Вимоги

* Go 1.16 або новіше
* Docker
* Docker Compose
* Kubernetes (локальний кластер або Minikube)

# Кроки для Запуску

1. Клонування Репозиторію

```bash 
git clone https://github.com/yourusername/go-microservices.git
cd go-microservices
```

2. Запуск за Допомогою Docker Compose

```bash 
docker-compose up --build
```

3. Деплоймент на Kubernetes

```bash
kubectl apply -f k8s/
```

4. Тестування Сервісів

```bash
   go test ./services/...
```

## Тестування

Для запуску тестів виконайте:

```bash
go test ./services/...
