package com.coderpad.app;

import java.util.*;
import java.util.concurrent.*;

public class Main {

    static int somethingSlow(int jobId) throws InterruptedException {
        System.out.println("starting job " + jobId);
        Thread.sleep(1000);
        System.out.println("completed job " + jobId);
        return jobId;
    }

    // public static void main(String[] args) throws InterruptedException {

    //     List<Thread> tasks = new ArrayList<>();

    //     final List<String> completedJobIds = new ArrayList<>();
    //     int jobCount = 10;

    //     for (int i = 0; i < jobCount; i++) {
    //         final int index = i;

    //         Thread task = new Thread(()->{
    //           int completedJob = index;

    //           try{
    //             completedJob = somethingSlow(index);
    //           }
    //           catch(InterruptedException ex)
    //           {
    //               Thread.currentThread().interrupt();
    //           }
    //           finally{
    //               completedJobIds.add(String.valueOf(completedJob));
    //           }
    //         });
    //         task.start();
    //         tasks.add(task);
    //     }
        
    //     // for (Thread task : tasks) {
    //     //  task.join();
    //     //}

    //     while(true)
    //     {
    //       int completedTasks = 0;
    //       int runningTasks = 0;

    //       for (Thread task : tasks) {
    //           if(task.isAlive())
    //           {
    //             completedTasks ++;
    //           }
    //           else
    //           {
    //             runningTasks++;
    //           }
    //       }
          
    //       if(completedTasks == jobCount)
    //         break;

    //       System.out.printf("Jobs running: %d, Jobs completed: %d", runningTasks, completedTasks);
    //       Thread.sleep(200);
    //     }

    //     System.out.printf("all jobs done! IDs: %s", String.join(", ",completedJobIds));
    // }

    
    public static void main(String[] args) throws InterruptedException
    {
      int maxConcurrency = 10;
      final List<String> completedJobIds = new ArrayList<>();
      BlockingQueue<Runnable> workQueue = new ArrayBlockingQueue<>(100);

      ExecutorService workerPool = new ThreadPoolExecutor(maxConcurrency, maxConcurrency, 10, TimeUnit.SECONDS, workQueue);

      int jobCount = 10;

      List<Future<?>> jobResults = new ArrayList<>();

      for(int i=0; i < jobCount; i++ )
      {
        final int index = i;

        jobResults.add(workerPool.submit(()->{
          int completedJob = index;

          try{
            completedJob = somethingSlow(index);
          }
          catch(InterruptedException ex)
          {
              Thread.currentThread().interrupt();
          }
          finally{
              completedJobIds.add(String.valueOf(completedJob));
          }
        }));
      }// for
    
      int completedJobCount = 0;

      while(completedJobCount < jobCount)
      {
          completedJobCount = 0;

          for (Future<?> result : jobResults) {
            if (result.isDone())
            {
              completedJobCount++;
            }
          }

          System.out.printf("Jobs Completed: %d, Jobs pending: %d", completedJobCount, jobCount-completedJobCount);
          Thread.sleep(200);
      }

      System.out.printf("\n all jobs done! IDs: %s", String.join(", ",completedJobIds));
      // TODO: Clean shutdown, include interrupt
      workerPool.shutdown();
    }
}