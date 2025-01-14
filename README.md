# Инструкция пользования

## 1. Соберите проект 

`go build -o [name] main.go`

## 2. Можете запускать проект

### 2.1 С помощью сборки

`./[name] command`

### 2.2 С помощью *main.go*

`go run main.go command`

## 3. Комманды

`add "description"` - Добавить задачу в список <br>

`list` - Посмотреть весь список задач <br>

`list todo/in progress/done` - Посмотреть задачи с указанным статусом <br>

`update id "new description"` - Обновить описание задачи по *ID* <br>

`mark-in-progress id` - Изменить статус на in progress задаче с указанным *ID* <br>

`mark-done id` - Изменить статус на done задаче с указанным *ID* <br>

`delete id` - Удалить задачи из списка <br>