package com.goandroidrpc.rpc;

import go.rpc.Rpc;

import java.util.concurrent.Callable;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.FutureTask;

import org.json.JSONException;
import org.json.JSONObject;

import android.os.Looper;
import android.util.Log;

public class RpcBackendCaller {
    protected ExecutorService mPool;
    protected UIThreadRunner mUiThreadRunner;

    RpcBackendCaller(UIThreadRunner uiThreadRunner) {
        mPool = Executors.newFixedThreadPool(2);
        mUiThreadRunner = uiThreadRunner;
    }

    public Object call(final String payload) {
        if (Looper.myLooper() != Looper.getMainLooper()) {
            return Rpc.CallBackend(payload);
        }

        Log.v("!!!", String.format("%s", "init outsourcer"));
        final UIThreadRunner.OutsourceExecutor outsourcing = mUiThreadRunner.outsourceExecutor();
        Log.v("!!!", String.format("%s", "after init outsourcer"));

        FutureTask<Object> task = new FutureTask<Object>(
            new Callable<Object>() {
                @Override
                public Object call() {
                    Log.v("!!!", String.format("%s", "before callbackend"));
                    Log.v("!!!", String.format("%s", payload));
                    String reply = Rpc.CallBackend(payload);
                    Log.v("!!!", String.format("%s", "before stop"));
                    mUiThreadRunner.leaveOutsourcing();
                    try {
                        return new JSONObject(reply).get("result");
                    } catch (JSONException e) {
                        return null;
                    }
                }
            }
        );

        mPool.execute(task);

        Log.v("!!!", String.format("%s", "before outsource"));
        outsourcing.outsource();
        Log.v("!!!", String.format("%s", "after outsource"));

        try {
            return task.get();
        } catch (ExecutionException e) {
            // pass
        } catch (InterruptedException e) {
            // pass
        }

        return null;
    }
}
