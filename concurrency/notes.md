- Goroutines are managed by Goruntime , Gorcheduler

- G-M-P model 

    - G = goroutine
    - M = Machine / OS Thread
    - P = Processor / scheduling context , run queues

        - When a new Goroutine is created G is called
        - It's own growable stack , program counter , stack pointer, status,links
        - G is not an OS thread
    

- Syscalls on M(Thread), FileIO, NetworkIO these kinds of calls are generally IO calls they blcok the M
- Whenever M is blocked due to IO operations, go uses netpoller

- Once these are unblocked , they are scheduled on Global Run Queue
- Gs on global queue are picked by P and then executed on corrosponding M(Thread)


- How Async Await are different than Go Routines
- What are stackless and stackfull routines/threads
- What would happen when the all M are blocked --> Go Creates new Ms , once unblocked it would delete additionally created M
- How many goroutines , go can handle --> in thousands

