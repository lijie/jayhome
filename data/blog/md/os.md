## Notes of Modern Operating Systems

### PROCESS

#### THREAD

##### Why need thread

1. In many applications, multipile activities are going on at once.
2. Threads are light weight than processes.
3. Threads yield no performance gain when all of them are cpu bound, but when there is substantial computing and also substantial I/O, having threads allow these activites to overlap, thus speeding up the application.
4. Threads are usefull on systems with multiple cpus, where real parallelism are possible.

The key is "improve performance".

#### The classical thread model

#### POSIX Thread



