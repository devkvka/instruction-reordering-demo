# Instruction reordering

This is a small demonstration that modern computers do not always execute
instructions linearly (in the order they appear in source code).

## The experiment

Two threads (in this case goroutines) run the following instructions concurrently.

| Thread 1 | Thread 2 |
|----------|----------|
| x = 1    | y = 1    |
| a = y    | b = x    |

If the instructions were performed strictly sequentially we would expect to always 
see a, b or both end up with the value 1, because both threads always set the other
variable before the other thread reads it. However if we repeat this until we observe 
that ```a == 0 && b == 0```, it always happens eventually.

This outcome is only possible if both threads had their write performed after the read.
The reordering can be done by the compiler, CPU, or even both which makes this reproducible
in any language that supports starting either real or green threads.

## Output example

```
% go build .
% ./instruction_reordering
Found (a == 0 && b == 0) after 336 iterations
```
