[RUS](https://github.com/larikhide/stepik-go/blob/main/3-8-parallelism/README.md)

# 3-8-8  

Write a function that takes a channel and an integer N as input. The function should return the value N+1 into the channel. The function should be named task().  

# 3-8-9  

Write a function that takes a channel and a string as parameters. The function should send the same string to the specified channel 5 times, appending a space to it each time.

The function should be named task2().  

# 3-8-10   

Write a pipeline element (function) that remembers the previous value and sends values to the next stage of the pipeline only if they differ from the previously received value.

Your function should take two channels as parameters - inputStream and outputStream. You will receive strings in the inputStream, and you should send values without duplicates to the outputStream. In the end, the outputStream should only contain values that do not repeat consecutively. Don't forget to close the channel ;)

The function should be named removeDuplicates().