# Chapter1
Concurrency is an interesting word because it means different thigs to differents people in
our filed. In addition to "concurrenty" you may have heard the words, "asynchronous", "parallel"
or "threaded" bandied about. some people take these words to mean the same thing and other people
very specifically delineate between each of those words. If we're to spend an entire book's
worth of time discussing concyrrency, it would be benefical to first spend some time discussing
what we mean when we say "concurrencyconcurrency"




After reading a lot of information about concurrency in this book, i found some interesing
thing:

## Why Is Concurrency Hard?
Concurrenct code is notoriously difficult to get right. It usually takes a few iterations
to get it working as expected, and even then it's not uncommon for bugs to exist in code
for years before some change in timing (/heavier disk utilization, more users logged into
  the system, etc.) causes a previously undiscovered buf to rear its head. Indeed, for this
  very book, I've gotten as many eyes aspossible on the code to try and mitigate this.


## Atomicity

When something is considered atomic, or to have the property of atomicity,
this means that whithin the context that it is operating, it is indivisible,
or uninterruptible.

So what does that reall mean, and why is this important to know when working with
concurrent code?

The first thint that's very important is the word "context". Something may be
atomic in one conetext but not another. Operations that are atomic in the context
of the operating system; operations that are atomic within the context fo the 
opreating system may not be atomic within the context of your machine; and opearations
that are atomic within the context of your machine may not be atomic within the
context of your application. In other words, the atomicity of an operation can change
dependinf on the currently deifned scope. This fact can work both for and against you!


## Deadlocs, Livelocksm and Starvation
The previous sections have all benn about discussing program
correctness in that if these issues are managed correctly,
your program will never give an incorrect answer. Unfornunately,
even if you successfully habdle these classes of issues, there 
is another class of issues to contend with: deadklocks, livelocks,
and starvation.

### Deadlock
A deadlocked program is one in which all concurrent processes are
waiting on one another. In thi state, the program will never
recover without outside intervention.

If that sounds grim, it's because it is! The Go runtime attempts
to do its part and will detect some deadlocks (all goroutines 
must be blocked, or "asleep"), but this doesn't do much to help 
you prvend deadlocks