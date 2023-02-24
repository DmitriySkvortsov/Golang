# func signature

Есть фигуры:
```
type figures int

const (
    square figures = iota // квадрат
    circle // круг
    triangle // равносторонний треугольник
) 
```

## Напишите функцию с такой сигнатурой:
```
func area(figures) (func(float64) float64, bool) 
```
## Функция должна возвращать:
- функцию для вычисления площади фигуры;
- true, если фигура известна;
- false, если фигура неизвестна.

## Нужно, чтобы её можно было применять так:
```
ar, ok := area(myFigure)
if !ok {
    fmt.Println("Ошибка")
    return
}
myArea := ar(x) 
```
