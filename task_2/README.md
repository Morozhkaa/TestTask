## Backend test task (GO)

- Сделать клиента для получения курсов ([page](https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1)).
- Добавить возможность получать курс для определенной криптовалюты.
- Курсы обновляем не чаще чем раз в 10 минут.

### Usage
1. Запуск с помощью команды:  `go run main.go` - вывод курсов всех криптовалют.

2. Чтобы вывести курсы определенных криптовалют, можно передать их названия в качестве аргументов. Например: 

`go run main.go "Polymesh" "Bitcoin Gold" "Baby Doge Coin"`

Вывод:
```
Bitcoin Gold (btg) - 16.56000$
Baby Doge Coin (babydoge) - 0.00000$
Polymesh (polyx) - 0.12504$
```