package com.goandroidrpc.rpc;

import java.util.concurrent.BrokenBarrierException;
import java.util.concurrent.Callable;
import java.util.concurrent.CyclicBarrier;
import java.util.concurrent.FutureTask;
import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.ReentrantLock;

import android.app.Activity;
import android.os.Looper;
import android.util.Log;

public class UIThreadRunner {
    protected Activity mActivity;
    protected ReentrantLock mExecutionLock;

    public volatile OutsourceExecutor outsourceExecutor;

    UIThreadRunner(Activity activity) {
        mActivity = activity;
        mExecutionLock = new ReentrantLock();
    }

    public Object run(
        Callable<Object> callable
    ) throws Exception {
        mExecutionLock.lock();
        Log.v("!!!", String.format("%s", Thread.currentThread()));

        Object result;
        try {
            if (outsourceExecutor != null) {
                Log.v("!!!", String.format("%s", "outsource executor"));
                result = outsourceExecutor.submit(callable).get();
            } else {
                Log.v("!!!", String.format("%s", "current executor"));
                FutureTask<Object> task = new FutureTask<Object>(callable);
                mActivity.runOnUiThread(task);
                result = task.get();
            }
        } finally {
            mExecutionLock.unlock();
        }

        return result;
    }

    public OutsourceExecutor outsourceExecutor() {
        mExecutionLock.lock();
        outsourceExecutor = new OutsourceExecutor();
        mExecutionLock.unlock();
        return outsourceExecutor;
    }

    public void leaveOutsourcing() {
        mExecutionLock.lock();
        outsourceExecutor.stop();
        outsourceExecutor = null;
        mExecutionLock.unlock();
    }

    public class OutsourceExecutor {
        protected ReentrantLock mLock;
        protected CyclicBarrier mBarrier;

        protected FutureTask<Object> mTask;

        OutsourceExecutor() {
            mTask = null;
            mLock = new ReentrantLock();
            mBarrier = new CyclicBarrier(2);
        }

        public FutureTask<Object> submit(
            Callable<Object> callback
        ) throws InterruptedException, BrokenBarrierException {
            if (callback == null) {
                mTask = null;
            } else {
                mTask = new FutureTask<Object>(callback);
            }
            mBarrier.await();
            return mTask;
        }

        public boolean runTask() {
            try {
                mBarrier.await();
                if (mTask == null) {
                    return false;
                }
                mTask.run();
                return true;
            } catch (InterruptedException e) {
                return false;
            } catch (BrokenBarrierException e) {
                return false;
            }
        }

        protected void stop() {
            try {
                submit(null);
            } catch (InterruptedException e) {
            } catch (BrokenBarrierException e) {
            }
        }

        public void outsource() {
            while (runTask()) {
                // just run tasks
            };
        }
    }
}
