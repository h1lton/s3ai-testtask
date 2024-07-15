# Тестовое задание

## Задание

Создать REST API на Golang, которое имитирует работу банкомата. Приложение должно поддерживать операции по созданию аккаунтов, пополнению баланса, снятию средств и проверке баланса. Все операции должны логироваться в консоль. Для обработки операций использовать горутины.


## Требования

### Интерфейсы:

Определить интерфейс `BankAccount` с методами:
- `Deposit(amount float64) error` - пополнение баланса.
- `Withdraw(amount float64) error` - снятие средств.
- `GetBalance() float64` - получение текущего баланса.

Реализовать этот интерфейс в структуре `Account`.


### Горутины:

Для каждой операции (пополнение, снятие, проверка баланса) должна создаваться отдельная горутина.
Использовать каналы для передачи данных между горутинами.

Основные операции:
- Создание нового аккаунта.
- Пополнение баланса.
- Снятие средств.
- Проверка баланса.


### Логирование:

Все операции должны логироваться в консоль с указанием времени выполнения операции и идентификатора аккаунта.


### Потокобезопасность:

Обеспечить потокобезопасность операций с балансом аккаунта.

### REST API:

Реализовать следующие эндпоинты:
- `POST /accounts` - создание нового аккаунта.
- `POST /accounts/{id}/deposit` - пополнение баланса.
- `POST /accounts/{id}/withdraw` - снятие средств.
- `GET /accounts/{id}/balance` - проверка баланса.
