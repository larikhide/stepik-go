[ENG](https://github.com/larikhide/stepik-go/blob/main/3-8-parallelism/README-ENG.md)  

# 3-8-8  

Напишите функцию которая принимает канал и число N типа int. Необходимо вернуть значение N+1 в канал. 
Функция должна называться task().

# 3-8-9  

Напишите функцию которая принимает канал и строку. Необходимо отправить эту же строку в заданный канал 5 раз, добавив к ней пробел.

Функция должна называться task2().  

# 3-8-10  

Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет значения на следующий этап конвейера только если оно отличается от того, что пришло ранее.

Ваша функция должна принимать два канала - inputStream и outputStream, в первый вы будете получать строки, во второй вы должны отправлять значения без повторов. В итоге в outputStream должны остаться значения, которые не повторяются подряд. Не забудьте закрыть канал ;)

Функция должна называться removeDuplicates()
