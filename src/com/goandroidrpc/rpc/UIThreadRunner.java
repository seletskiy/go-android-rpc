package com.goandroidrpc.rpc;

import java.util.concurrent.Callable;
import java.util.concurrent.CyclicBarrier;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.FutureTask;
import java.util.concurrent.LinkedBlockingQueue;
import java.util.concurrent.locks.ReentrantLock;

import android.app.Activity;
import android.util.Log;

public class UIThreadRunner {
    final protected Activity mActivity;

    protected ExecutorService mThreadPool;
    protected LinkedBlockingQueue<Task> mTasksQueue;

    public volatile OutsourceExecutor outsourceExecutor;

    UIThreadRunner(Activity activity) {
        Log.v("!!!", String.format("%s", "constructor"));
        mThreadPool = Executors.newFixedThreadPool(2);
        mActivity = activity;
        mTasksQueue = new LinkedBlockingQueue<Task>();

        mThreadPool.submit(new Runnable() {
            public void run() {
                while (true) {
                    Task task;
                    try {
                        task = mTasksQueue.take();
                        try {
                            mTasksQueue.put(task);
                            Log.v("!!!", String.format("%s", task));
                            if (task.future != null) {
                                mActivity.runOnUiThread(task.future);
                                // task.future.isDone() can be false there
                                task.future.get();
                                mTasksQueue.remove(task);
                            } else {
                            }
                        } catch (Exception e) {
                            mTasksQueue.remove(task);
                        }
                    } catch (InterruptedException e) {
                    }
                }
            }
        });
    }

    public Object run(
        Callable<Object> callable
    ) throws Exception {
        Task task = new Task(new FutureTask<Object>(callable));
        mTasksQueue.put(task);
        Object result = task.future.get();
        return result;
    }

    public Object runAndDispatchUI(final Callable<Object> callable) {
        final OutsourceExecutor executor = new OutsourceExecutor();

        FutureTask<Object> task = new FutureTask<Object>(
            new Callable<Object>() {
                @Override
                public Object call() {
                    Object result;
                    try {
                        result = callable.call();
                    } catch (Exception e) {
                        return null;
                    }

                    executor.stop();

                    return result;
                }
            }
        );

        mThreadPool.execute(task);


        executor.runTasks();

        try {
            return task.get();
        } catch (ExecutionException e) {
            // pass
        } catch (InterruptedException e) {
            // pass
        }

        return null;
    }

    public void destroy() {
        mThreadPool.shutdown();
    }

    public class OutsourceExecutor {
        protected ReentrantLock mLock;
        protected CyclicBarrier mBarrier;

        OutsourceExecutor() {
            mLock = new ReentrantLock();
            mBarrier = new CyclicBarrier(2);
        }

        public boolean runTask() {
            try {
                Task task = mTasksQueue.take();
                if (task.future == null) {
                    return false;
                }
                task.future.run();
                return true;
            } catch (InterruptedException e) {
                return false;
            }
        }

        public void stop() {
            try {
                mTasksQueue.put(new Task(null));
            } catch (InterruptedException e) {
            }
        }

        public void runTasks() {
            while (runTask()) {
                // just run tasks
            };
        }
    }

    public class Task {
        public FutureTask<Object> future;

        Task(FutureTask<Object> task) {
            future = task;
        }
    }
}
