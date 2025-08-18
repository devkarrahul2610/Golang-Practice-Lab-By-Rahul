# Golang-Practice-Lab-By-Rahul

In Golang, go routine is light-weight,user-space thread managed by go runtime. It is extremly cheap to create starting with just
2KB of stack trace, and scales dynamically. You start it using the go keyword.
ex. go doTask()
behind the scenes what happed here? Ans--> Go doesn't create a new OS thread for each goroutine. Instead it uses its own GMP schedular
to multiplex thousands of goroutines over a small number of OS threads.
In Go, concurrancy means managing multiple task at once (Interleaved Execution) while parrelelisum means execute multiple task
simulteniously on multiple CPU cores.
By default Go is concurrent,But with multiple cores (GOMAXPROS), it can be parellel too.
runtime.GOMAXPROS(runtime.NumCPU) enables true parellelisum

Channels are typed conduits through which goroutines communicate. They enforce synchronization and make data sharing safe without
explicit locks.