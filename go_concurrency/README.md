# Task 3 - Explain Go concurrency code

1. **How the highlighted constructs work?**
    ```go
    cnp := make(chan func(), 10)
        for i := 0; i < 4; i++ {
            go func() {
                for f := range cnp {
                    f()
                }
            }()
        }
        cnp <- func() {
            fmt.Println("HERE1")
        }
    ```
    
    First, a buffered channel is created with a capacity of 10 (of type `func()`). Channels are used to communicate between goroutines.

    Then, 4 goroutines are created using the `go` keyword. Each goroutine listens on the channel `cnp` and executes the function received on the channel.

    A function literal is sent to the channel `cnp` using the `<-` operator. The function literal prints `HERE1` to the console.


2. **Use-cases of what these constructs could be used for?**

    Some of the use-cases are:
    - Job scheduling systems and task queues
    - Implementing a pool of workers that execute tasks concurrently.

3. **What is the significance of the for loop with 4 iterations?**

    The for loop with 4 iterations starts 4 worker goroutines, which will pull functions from the cnp channel and execute them. They help in parallel execution of tasks.

4. **What is the significance of make(chan func(), 10)?**

    The `make(chan func(), 10)` creates a buffered channel of functions with a capacity of 10. This means that the channel can hold up to 10 functions before blocking the sender.

    `make` takes in the type and optionally the size. Here the 1st parameter is of `func()` type and the 2nd parameter is the buffer size of 10.

5. **Why is “HERE1” not getting printed?**

    "HERE1" is not getting printed because the main function exits before any of the worker goroutines have a chance to pull the function from the channel and execute it.

    Since, the main function exits immediately after sending the function to the channel, the worker goroutines don't get a chance to execute the function and are terminated. 

    We can use synchronization mechanisms like `sync.WaitGroup` or even `time.Sleep` to wait for the worker goroutines to finish executing before exiting the main function.