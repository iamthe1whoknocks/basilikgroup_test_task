Тестовое задание
Сделать сервис на golang, который будет выполнять арифметические операции над числами.
У http сервиса должно быть 4 api метода (add, sub, mul, div), каждый из которых выполняет соответствующую арифметическую операцию (сложения, вычитания, умножения и деления).
Каждый API метод на вход получает 2 get параметра - числа, с которыми надо выполнить операцию.
Пример запроса
http://localhost/api/add?a=1&b=15
Пример ответа
{
    "Success":true,
    "ErrCode":"",
    "Value":16
}
Обернуть сервис в docker контейнер, таким образом, чтобы сборка сервиса выполнялась непосредственно внутри контейнера.
Покрыть решение авто-тестами, которые проверяют, что решение хоть как-то работает
