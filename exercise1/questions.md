Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?

Concurrency is about dealing with multiple tasks at once by interleaving their execution. Tasks make progress by switching between them, but only one may execute at any given moment (on a single core).

Parallelism is about executing multiple tasks simultaneously, requiring multiple processors or cores where tasks literally run at the same time.


What is the difference between a *race condition* and a *data race*? 

A race condition is when your program's outcome depends on timing - like if two threads are racing to do something and whoever wins changes the result. It's about the order things happen in.

A data race is more specific - it's when multiple threads access the same variable at the same time without any protection (like a mutex), and at least one of them is writing to it. This corrupts your data.

 
*Very* roughly - what does a *scheduler* do, and how does it do it?
A scheduler decides which thread gets to run on the CPU and when. It maintains a queue of runnable threads and uses a timer interrupt to periodically stop the current thread (preemption), save its state, pick the next thread from the queue, and restore that thread's state so it can continue running.

### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?

Threads solve problems where you need to do multiple things concurrently.

Responsiveness: Keep the UI interactive while doing background work. Without threads, your application freezes while waiting for operations to complete.

Performance on multicore systems: Threads let you use all available CPU cores. A single-threaded program only uses one core, wasting the others.

I/O efficiency: While one thread waits for slow disk or network operations, other threads can do useful work instead of the entire program sitting idle.

Natural problem modeling: Some problems are inherently concurrent. A web server handling multiple client connections, a game with separate rendering and physics systems, or a download manager fetching multiple files all map naturally to multiple threads.

Real-time requirements: Audio playback or sensor data collection need dedicated threads with guaranteed timing, separate from other application logic.

Threads let you overlap waiting time with computation, utilize multiple cores, and structure programs around concurrent activities rather than forcing everything into a sequential flow.


Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?

Fibers (or green threads/coroutines) are like threads, but they're managed by your program or language runtime instead of the operating system.
Regular threads are scheduled by the OS kernel. Fibers are scheduled by your program itself in user space.
Why use fibers:
They're much cheaper. Creating an OS thread takes milliseconds and several megabytes of memory. A fiber takes microseconds and a few kilobytes. You can have millions of fibers where you could only have thousands of threads.
Context switching is faster. Switching between fibers happens in user space without expensive system calls. This makes them 10-100x faster to switch than OS threads.
They're great for I/O-heavy programs. A web server handling 100,000 connections can give each one its own fiber without running out of resources. Doing this with OS threads would crash your system.
The code stays simple. You can write normal-looking blocking code without callbacks or complex async patterns.
The downside: Fibers don't give you true parallelism on their own since they're cooperatively scheduled within one or a few OS threads. If one fiber does blocking work, it can hold up others. Languages like Go solve this by multiplexing many fibers onto multiple OS threads automatically.


Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?

Both.
Easier: For naturally concurrent problems like web servers or responsive UIs, threads make the code cleaner. You write straightforward blocking code instead of managing complex callback chains.
Harder: You get new bugs that don't exist in sequential code - race conditions, deadlocks, data races. These are hard to reproduce and debug because they depend on timing.
You have to think about synchronization constantly. Which data is shared? What needs locks? What order should locks be acquired? Get it wrong and your program hangs or corrupts data.
Testing is harder because bugs might only appear under specific timing or load conditions.

What do you think is best - *shared variables* or *message passing*?

As a programmer, I much prefer working with message passing. It leads to clearer code, which brings fewer and less complicated bugs in development, as well as much better maintainability.
