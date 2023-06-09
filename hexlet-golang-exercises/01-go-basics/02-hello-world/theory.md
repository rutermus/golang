Основы Go

# Привет, Мир!

Начнем с программы Hello World:

```go
// Определение пакета main
package main
// Импорт пакета fmt
import "fmt"
// Определение функции main
func main() {
  // Вызов функции Print из пакета fmt
  // Отступ 1 таб
  fmt.Print("Hello, World!") // В конце не нужна точка с запятой
}
```

Сначала мы определили пакет `main`:

```go
package main
```

_Пакеты выполняют роль неймспейсов_ (логически обособленные директории с локальным именованием) _и используются для группировки функций_. Все Go-файлы начинаются с объявления пакета, к которому они относятся. Внутри одного пакета может быть множество функций.

После определения пакета мы написали функцию `main()`. В ней описана вся логика программы:

```go
func main() {
  ...
}
```

`main()` — главная функция, которая выполняется при запуске любой Go-программы и является точкой входа в программу. Она не может принимать или возвращать какие-либо аргументы.

Чтобы вывести строку в терминал, мы использовали внешний пакет fmt:

```go
import "fmt"
```

Импорт сторонних пакетов описывается словом `import`. Блок импортов располагается в начале Go-файла сразу после названия пакета.

После импорта мы можем вызывать функции пакета в своем коде.

```go
fmt.Print("Hello, World!");
```

Сторонний пакет — это не объект, а неймспейс. Его используют как префикс к функциям. Обратиться к функции стороннего пакета можно через точку. Функции, как и во многих языках, вызываются через скобки и передачу аргумента внутрь.

Функция `Print()` написана с большой буквы, а `main()` — с маленькой. В Go функция публична, если ее первая буква заглавная. Публичность или экспортируемость означает, что мы можем использовать эту функцию в других пакетах. Если функция начинается с маленькой буквы, то она является приватной и может использоваться только внутри пакета. Функция `fmt.Print()` — публичная, поэтому мы можем вызывать ее в своем коде. А функция `main()` — приватная, она доступна только в рамках нашего пакета `main`.

Рассмотрим аргумент `"Hello, World!"`. Строки практически всегда объявляются в двойных кавычках. Существует еще один способ описать многострочную строку через обратную кавычку `, но он используется в исключительных ситуациях. Мы пока будем использовать только двойные.

Одной из особенностей языка Go является отсутствие точек с запятыми. Чтобы компилятор понял код, необходимы правильные переносы строк и отступы: табами, а не пробелами.
