# Функции, структуры, указатели

## func STDIN STDOUT

Напишите функцию f(), которая будет принимать строку text и выводить ее (печатать на экране).

### Решение
#### 01_func_in_out


## func minimum num

Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.
Вводится четыре числа.
Необходимо вернуть из функции наименьшее из 4-х данных чисел

### Решение
#### 02_min_num


## func vote

Напишите "функцию голосования", возвращающую то значение (0 или 1), которое среди значений ее аргументов x, y, z встречается чаще.
Вводится 3 числа - x, y и z (x, y и z равны 0 или 1).
Необходимо вернуть значение функции от x, y и z.

### Решение
#### 03_func_vote


## func fibonachi

Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1. Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ... Напишите функцию, которая по указанному натуральному n возвращает φn.
Вводится одно число n.
Необходимо вывести  значение φn.

### Решение
#### 04_func_fibo


## sumInt

Напишите функцию sumInt, принимающую переменное количество аргументов типа int, 
и возвращающую количество полученных функцией аргументов и их сумму. 
```
Пример вызова вашей функции:
a, b := sumInt(1, 0)
fmt.Println(a, b)
```
Результат: 2, 1

### Решение
#### 05_func_sumInt


## link

Напишите функцию, которая умножает значения на которые ссылаются два указателя и выводит 
получившееся произведение в консоль. Ввод данных уже написан за вас.

### Решение
#### 06_mult_links


## some revers

Поменяйте местами значения переменных на которые ссылаются указатели. 
После этого переменные нужно вывести.

### Решение
#### 07_revers_int


## shoot or not

Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, 
int, int соответственно. У этой структуры должны быть методы: Shoot и RideBike, которые не 
принимают аргументов, но возвращают значение bool.

Если значение On == false, то оба метода вернут false.

Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод 
возвращает true), если его нет, то метод вернет false. Метод RideBike работает также, 
но только зависит от свойства Power.

Чтобы проверить, что вы все сделали правильно, вы должны создать указатель на экземпляр этой
структуры с именем testStruct в функции main, в дальнейшем программа проверит результат.

### Решение
#### 08_shoot_ride


## wrong line

На вход подается строка. Нужно определить, является ли она правильной или нет. Правильная строка начинается с заглавной буквы и заканчивается точкой. Если строка правильная - вывести Right иначе - вывести Wrong

Sample Input:
- Быть или не быть.
Sample Output:
- Right

### Решение
#### 09_be_or_not


## Palindrom

На вход подается строка. Нужно определить, является ли она палиндромом. 
Если строка является палиндромом - вывести Палиндром иначе - вывести Нет. 
Все входные строчки в нижнем регистре.

### Решение
#### 10_palindrom


## In string

Даются две строки X и S. Нужно найти и вывести первое вхождение подстроки S в строке X. 
Если подстроки S нет в строке X - вывести -1

### Решение
#### 11_in_string


## delete salt

На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные 
символы (считая с нуля)

Sample Input:
- ihgewlqlkot
Sample Output:
- hello

### Решение
#### 12_salt


## delete symb

Дается строка. Нужно удалить все символы, которые встречаются более одного раза 
и вывести получившуюся строку

Sample Input:
- zaabcbd
Sample Output:
- zcd

### Решение
#### 13_del_symb


## Password

Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные 
требования. Длина пароля должна быть не менее 5 символов, он должен содержать только 
арабские цифры и буквы латинского алфавита. На вход подается строка-пароль. 
Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

Sample Input:
- fdsghdfgjsdDD1
Sample Output:
- Ok


### Решение
#### 14


## divide

Вы должны вызвать функцию divide которая делит два числа, но она возвращает не только 
результат деления, но и информацию об ошибке. В случае какой-либо ошибки вы должны 
вывести "ошибка", если ошибки нет - результат функции. 
Функция divide(a int, b int) (int, error) получает на вход два числа которые нужно 
поделить и возвращает результат (int) и ошибку (error).

### Решение
#### 15_divide


## hypotenuse

На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы

### Решение
#### 16_hypotenuse

##

Дана строка, содержащая только английские буквы (большие и маленькие). Добавить символ 
‘*’ (звездочка) между буквами (перед первой буквой и после последней символ ‘*’
добавлять не нужно).

Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков.
Вывести строку, которая получится после добавления символов '*'.

Sample Input:
- LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO
Sample Output:
- L*I*t*B*e*o*F*L*c*S*G*B*O*F*Q*x*M*H*o*I*u*D*D*W*c*q*c*V*g*k*c*R*o*A*e*o*c*X*O


### Решение
#### 17_enter_salt

## maximum

Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.
Вводится строка ненулевой длины. Известно также, что длина строки не превышает 
1000 знаков и строка содержит только арабские цифры.

Выведите максимальную цифру, которая встречается во введенной строке.

### Решение
#### 18_max_in_num


## squares

На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа 
и вывести получившееся число. 

Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1. 
Единица в квадрате - 1. В итоге получаем 811181

Sample Input:
- 9119
Sample Output:
- 811181

### Решение
#### 19_squares


## fluctuations

Требуется вычислить период колебаний (t) математического маятника 
(мы округлили некоторые значения для удобства проверки), 
для этого нужно найти циклическую частоту колебания пружинного маятника (w), 
в формуле w встречается масса которую также нужно найти, все нужные формулы приведены ниже:


Напишите три функции, каждая из которых будет выполнять конкретную формулу. 
Название функций обязательно должны соответствовать букве формулы: T(), W() и M(). 
Для того чтобы найти t - необходимо сначала найти w, и т.д. 
Так что используйте результат функции W() в формуле функции T() - то-есть 
вызывайте функцию W() в T(). Аналогично и с W(), M(). 
```
t = w/6 , 

w = sqrt{k/m}, 

m=p*v
```

### Решение
#### 

