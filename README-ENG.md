[RUS](https://github.com/larikhide/stepik-go/blob/main/README.md)

## Disclaimer
  
The solutions presented in this repository are solutions to tasks from a free course on stepik.org and are based on my subjective assessment of their optimality. These solutions are provided for reference purposes only, and it is recommended to independently complete the course.

I welcome your comments and suggestions for improving these solutions, particularly in the context of their application in commercial high-load development. Your contributions can greatly enrich and enhance the solutions presented here, taking into account practical experience and modern requirements of professional development.



[3.6.6](https://github.com/larikhide/stepik-go/blob/main/3-6-json/3-6-6-json.go) Read JSON from StdIn (or from a file), perform operations on the data, and return JSON to StdOut  
[3.6.9](https://github.com/larikhide/stepik-go/blob/main/3-6-json/3-6-9-json/3-6-9-json.go) Read JSON from a file or from the internet with unknown structure, find what you need, and calculate  
[3.7.3](https://github.com/larikhide/stepik-go/blob/main/3-7-time/3-7-3-time.go) Accept a date in one format and return it in another  
[3.7.4](https://github.com/larikhide/stepik-go/blob/main/3-7-time/3-7-4-time.go) Accept a date, parse it, and modify it if necessary. Read from StdIn using bufio.NewReader() when you need precise control over the buffer size and reading a specific amount of data at a time.  
[3.7.7](https://github.com/larikhide/stepik-go/blob/main/3-7-time/3-7-7-time.go) Accept a string of dates with a specific delimiter, convert it into an array, and calculate the difference between dates. Read from StdIn using bufio.NewScanner when you need to scan data line by line or based on a specific delimiter.  
[3.7.8](https://github.com/larikhide/stepik-go/blob/main/3-7-time/3-7-8-time.go) Accept duration in a free format and a date in Unix format. Add the duration to the date and output the new date in Unix format.  

[3.8.8](https://github.com/larikhide/stepik-go/blob/main/3-8-parallelism/3-8-8-parallelism.go) Receive data in a channel, modify it, and pass it back to the same channel.  

[3.8.9](https://github.com/larikhide/stepik-go/blob/main/3-8-parallelism/3-8-9-parallelism.go) Receive a string in a channel and send it five times to the same channel with a space added each time.     

[3.8.10](https://github.com/larikhide/stepik-go/blob/main/3-8-parallelism/3-8-9-parallelism.go) There are 2 channels. We receive strings in one channel and send the string to the other channel only if it is different from the previous one.  

[3.9.3]() Call the main function within a goroutine and wait for the completion of its execution.