[ENG](https://github.com/larikhide/stepik-go/blob/style/README-ENG.md)

## Disclaimer
Решения, представленные в данном репозитории, являются решениями задач из бесплатного курса на [stepik.org](https://stepik.org/course/54403/syllabus) и представлены в соответствии с моим субъективным мнением об их оптимальности. Эти решения предоставлены исключительно для ознакомления, и рекомендуется пройти сам курс самостоятельно.

Я буду рад вашим комментариям и предложениям по улучшению данных решений, особенно в контексте их применения в коммерческой хайлоад разработке. Ваш вклад может значительно обогатить и доработать представленные здесь решения, учитывая практический опыт и современные требования профессиональной разработки.

  
[3.6.6](https://github.com/larikhide/stepik-go/blob/main/3-6-json/3-6-6-json.go) Прочитать JSON из StdIn (или из файла) сделать что-то с данными, вернуть JSON в StdOut  
[3.6.9](https://github.com/larikhide/stepik-go/blob/main/3-6-json/3-6-9-json/3-6-9-json.go) Прочитать JSON из файла или из интернета с неивестной структурой, найти там то что тебе нужно и посчитать  
[3.7.3](https://github.com/larikhide/stepik-go/blob/main/3-7-time/3-7-3-time.go) принять дату в одном формате, вернуть в другом  
[3.7.4](https://github.com/larikhide/stepik-go/blob/main/3-7-time/3-7-4-time.go) принять дату, распарсить ее и изменить, если нужно. Чтение из StdIn с помощью bufio.NewReader() когда вам нужен точный контроль над размером буфера и чтением определенного количества данных за раз.   
[3.7.7](https://github.com/larikhide/stepik-go/blob/main/3-7-time/3-7-7-time.go) принять строку из дат со специфическим разделителем, превратить это в массив, посчитать разницу между датами. Чтение из StdIn с помощью bufio.NewScanner когда вам нужно сканировать данные построчно или по определенному разделителю.  
[3.7.8](https://github.com/larikhide/stepik-go/blob/main/3-7-time/3-7-8-time.go) принять продолжительность в свободном формате, дату в Unix формате. Добавить к дате продолжительность и выдать новую дату в Unix  

[3.8.8](https://github.com/larikhide/stepik-go/blob/main/3-8-parallelism/3-8-8-parallelism.go) принять в канал данные, изменить их и передать в тот же канал

[3.8.9](https://github.com/larikhide/stepik-go/blob/main/3-8-parallelism/3-8-9-parallelism.go) принять в канал строку и отправить ее пять раз в этот же канал с добавлением пробела  

[3.8.10](https://github.com/larikhide/stepik-go/blob/main/3-8-parallelism/3-8-9-parallelism.go) есть 2 канала. принимаем строки в один, отправляем эту строку в другой, если она не совпадает с предыдущей  