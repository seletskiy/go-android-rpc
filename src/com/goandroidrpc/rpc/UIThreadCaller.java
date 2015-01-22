package com.goandroidrpc.rpc;

import android.app.Activity;
import android.os.Looper;
import android.util.Log;

import java.util.List;
import java.lang.reflect.Method;
import java.util.concurrent.FutureTask;
import java.util.concurrent.Callable;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.CancellationException;
import java.lang.InterruptedException;

public class UIThreadCaller {
    protected Activity mActivity;

    UIThreadCaller(Activity activity) {
        mActivity = activity;
    }

    public Object call(
        Callable<Object> callable
        //final Object mViewObject,
        //final Method mMethodToCall,
        //final List<Object> mRequestedParams
    ) throws Exception {
        FutureTask<Object> futureResult = new FutureTask<Object>(
            callable
        );
        mActivity.runOnUiThread(futureResult);
        Object result = futureResult.get();
        return result;
    }
}
