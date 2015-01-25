package com.goandroidrpc.rpc;

import java.util.concurrent.Callable;

import org.json.JSONObject;

import android.app.Activity;
import android.content.Context;
import android.util.Log;

public class RpcHandlerChangeLayout implements RpcHandlerInterface {
    public JSONObject Handle(Context context, JSONObject payload) {
        JSONObject result = new JSONObject();
        try {
            final String layoutName = payload.getString("layout");

            Activity activity = (Activity) context;

            final MainActivity mainActivity = ((MainActivity) activity);
            mainActivity.uiThreadRunner.run(new Callable<Object>() {
                public Object call() {
                    try {
                        mainActivity.setContentView(R.layout.class.getField(layoutName).getInt(null));
                    } catch(NoSuchFieldException e) {
                    } catch(IllegalAccessException e) {
                        // @TODO
                    }
                    return null;
                }
            });
        } catch (Exception e) {
            // @TODO: proper exception handling
            Log.v("!!! EXCEPT", String.format("%s", e));
        }

        return result;
    }

    public void destroy() {
        // pass
    }
}
