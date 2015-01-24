package com.goandroidrpc.rpc;

import go.rpc.Rpc;

import java.util.concurrent.Callable;
import java.util.concurrent.RejectedExecutionException;

import android.os.Looper;

public class RpcBackendCaller {
    protected UIThreadRunner mUiThreadRunner;

    RpcBackendCaller(UIThreadRunner uiThreadRunner) {
        mUiThreadRunner = uiThreadRunner;
    }

    public Object call(final String payload) {
        if (Looper.myLooper().getThread() != Looper.getMainLooper().getThread()) {
            return Rpc.CallBackend(payload);
        }


        Object result;
        try {
            result = mUiThreadRunner.runAndDispatchUI(new Callable<Object>() {
                public Object call() {
                    return Rpc.CallBackend(payload);
                }
            });
        } catch(RejectedExecutionException e) {
            return null;
        }

        return result;
    }

    public void destroy() {
        mUiThreadRunner.destroy();
    }
}
