Interviewer:
https://www.linkedin.com/in/gavinyue/
https://gavinyue.com/

# Technical Screening

## Instructions 1

Here, we have a simple for loop inside main. It iterates 10 times. Each iteration, it calls somethingSlow(). somethingSlow() represents external work that is outside of our control. We have a 1-second sleep statement to represent that.

Change this such that the entire program completes in about 1 second, with the following constraints:

You may NOT modify somethingSlow()
You may NOT use external packages or frameworks (JDK only)
You may NOT use java.util.concurrent.Executors, ExecutorService, or CompletableFuture

## Instructions 2

somethingSlow() now returns the result of its work. Collect the results from each job and, back in main outside the for loop, print them.

## Instructions 3

Each job is now long-lived. From the shell, show me how to view the number of jobs currently running.

## Instructions 5

One problem we may face when introducing concurrency is that too much may have a negative impact. First, we will add more jobs.

Update your code such that only 4 jobs run at any one time, with these constraints:

All jobs must be queued up front.
Jobs must start as soon as there is capacity. There should be no batch-queueing of new jobs after earlier jobs terminate.
If your solution uses a long-lived mechanism from which jobs execute, assign a unique identifier to that mechanism. Annotate each job's result with that identifier.

Design your solution such that the work function (somethingSlow) is configurable — not hard-coded inside the worker.

Evaluate each job's result in main as soon as it is available. If the result code's name is STATUS_SUCCESS, write the result to stdout. Otherwise, write to stderr. Include all data from the result structure.