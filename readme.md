## Aлгоритмы и структуры данных
https://education.vk.company/curriculum/program/lesson/25129/

### Домашнее задание 2:


#### Задача А:
На вход подается не отсортированный массив целых чисел и некоторое целое число «element».
Необходимо написать функцию, которая вернет количество элементов, которые не равны данному числу «element».
Возвращаемое значение после работы функции обозначает количество чисел не равные «element».

##### Пример 1
###### Входные данные:

9
7 8 1 3 11 3 9 4 3
3
###### Выходные данные:

6

#### Задача В:

В школе прошел экзамен по математике. Несколько человек списали решения и были замечены. Этим школьникам поставил 0 баллов. Задача: есть массив с оценками, среди которых есть 0. Необходимо все оценки, равные нулю перенести в конец массива, чтобы все такие школьники оказались в конце списка.

Примечание: первая строка во вводе - число элементов в массиве
##### Пример 1
###### Входные данные:

9
0 33 57 88 60 0 0 80 99
###### Выходные данные:

33 57 88 60 80 99 0 0 0
##### Пример 2
###### Входные данные:

8
0 0 88 0 99 0 100 0 
###### Выходные данные:

88 99 100 0 0 0 0 0

### Домашнее задание 3:


#### Задача А:
Реализовать сортировку вставками (insertion sort)

На вход подаётся массив чисел

На выходе ожидается отсортированный массив чисел



##### Пример 1
###### Входные данные:

5
12 11 13  5  6

###### Выходные данные:

5 6 11 12 13

### Домашнее задание 4:


#### Задача А:
Реализовать нисходящую сортировку слиянием.
 



##### Пример 1
###### Входные данные:

5
12 11 13  5  6

###### Выходные данные:

5 6 11 12 13

#### Задача B:
Сортировка книг

Описание:
Представьте, что у вас есть электронная библиотека, содержащая информацию о книгах. Каждая книга имеет уникальный номер ISBN, название и год издания.

Вам нужно написать программу, которая сортирует все книги по году издания в порядке возрастания. Если две или более книг были изданы в один и тот же год, сортируйте их по названию в алфавитном порядке. Используйте сортировку слиянием для решения этой задачи.


##### Пример 1
###### Входные данные:

3
1234567890 "Война и мир" 1869
9876543210 "1984" 1949
1111111111 "Сказка о рыбаке и рыбке" 1833

###### Выходные данные:

1111111111 "Сказка о рыбаке и рыбке" 1833
1234567890 "Война и мир" 1869
9876543210 "1984" 1949

### Рубежный контроль 1:


#### Задача А:
Дан не отсортированный массив целых чисел. Необходимо перенести в начало массива все четные числа. При этом последовательность четных чисел должна быть сохранена. То есть если 8 стояла после 6, то после переноса в начало, она по-прежнему должна стоять после 6. 



##### Пример 1
###### Входные данные:

7
3 2 4 1 11 8 9

###### Выходные данные:

2 4 8 1 11 3 9

В примере мы видим, что все четные числа перешли вперед сохраняя свое расположение относительно друг друга
 
### Домашнее задание 6:


#### Задача А:
Дана строка s. Строка состоит из английских букв в нижнем регистре.Необходимо из строки удалить все рядом стоящие повторяющиеся буквы. Например, в строке xyyx мы должны удалить yy, а после и оставшиеся xx и того после должна получиться пустая строка. Или же в строке fqffqzzsd после удаления остануться только fsd. Первыми удаляться ff, являющимися третьими и четвертыми символами, затем qq и после уже zz.

##### Пример 1
###### Входные данные:

cdffd

###### Выходные данные:

c
#### Задача В:
Дан массив не отсортированных целых чисел. Написать функцию, которая вернет первое с конца четное число. При написании кода используйте принцип стека. Если массив не содержит четного числа возвращать -1.
##### Пример 1
###### Входные данные:

5
8 9 10 16 9
###### Выходные данные:
16

### Домашнее задание 7:


#### Задача А:
Есть список товаров в магазине отсортированный по возрастанию, необходимо понять, есть ли в этом списке товар с заданной ценой. В случае если цена найдена возвращайте true
 

##### Пример 1
###### Входные данные:

12
100 450 730 800 950 999 1000 3000 3300 8000 9990 10000
999
###### Выходные данные:

true
### Домашнее задание 8:


#### Задача А:
Дан отсортированный по возрастанию массив целых чисел и заданное число.
Заданное число может и не находиться в массиве. Тогда необходимо вернуть предполагаемый индекс, где мог бы находится элемент. Другими словами, найдите правильное место для вставки элемента. Если же число есть, то возвращаем его индекс.


##### Пример 1
###### Входные данные:

5
5 7 9 11 13
6
###### Выходные данные:

1
### Домашнее задание 9:


#### Задача А:
На вход функции подается две строки a и b
Используя хеш таблицу, попытайтесь определить, является ли строка b анаграммой к строке a.

Анаграмма – это слово или фраза, образованные путем перестановки букв другого слова или фразы, как правило, используя все исходные буквы ровно один раз

##### Пример 1
###### Входные данные:

"anagram", "nagaram"
###### Выходные данные:

true
#### Задача B:
Есть строка с многократным повторением разных букв. Надо понять какое максимальное количество повторений встречается в строке. Используя хеш таблицу, записывайте пару буква - количество вхождений в строку.
##### Пример 1
###### Входные данные:

adeuuuuio
###### Выходные данные:

4
### Рубежный контроль 2:


#### Задача А:
Дан отсортированный по возрастанию массив и некоторое целое число. Необходимо найти минимальный подмассив, в котором может содержаться это число методом экспоненциального поиска.


##### Пример 1
###### Входные данные:

11
8 11 12 16 18 21 33 36 48 54 63
16
###### Выходные данные:

2 4

Число 16 находится в подмассиве от индекса 2 до индекса 4
#### Задача B:
Дан массив целых чисел. Необходимо вернуть элемент, который встречается больше половины раз. Если такого элемента нет, возвращайте -1
##### Пример 1
###### Входные данные:

7
7 7 8 8 8 8 9
###### Выходные данные:

8