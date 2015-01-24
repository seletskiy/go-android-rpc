package com.goandroidrpc.rpc;

import java.util.concurrent.Callable;

import org.json.JSONException;
import org.json.JSONObject;

import android.content.Context;
import android.util.Log;
import android.view.View;
import android.view.ViewGroup;

public class RpcHandlerAttachView implements RpcHandlerInterface {
    public JSONObject Handle(
        Context context, JSONObject request
    ) throws JSONException {
        JSONObject result = new JSONObject();

        MainActivity activity = (MainActivity) context;

        try {
            result = attachView(
                activity,
                Integer.parseInt(request.getString("id")),
                Integer.parseInt(request.getString("viewGroupId"))
            );
        } catch (JSONException e) {
            result.put("error",
                String.format("error in request: %s", e.getMessage())
            );
        }

        return result;
    }

    protected JSONObject attachView(
        MainActivity activity,
        Integer id,
        Integer targetViewId
    ) throws JSONException {
        JSONObject result = new JSONObject();

        final ViewGroup targetView = (ViewGroup) activity.findViewById(
            targetViewId
        );

        final View orphanView;
        try {
            orphanView = activity.orphanViews.get(id);
        } catch(Exception e) {
            result.put("error",
                String.format(
                    "view with ID '%d' is not either not exist or already attached",
                    id
                )
            );
            return result;
        }

        activity.orphanViews.remove(id);

        try {
            activity.uiThreadRunner.run(
                new Callable<Object> () {
                    @Override
                    public Object call() throws Exception {
                        try {
                            targetView.addView(orphanView);
                        } catch(Exception e) {
                            // @TODO
                            Log.v("!!!", e.toString());
                        }

                        return null;
                    }
                }
            );
        } catch (Exception e) {
            result.put("error", e.toString());
        }

        return result;
    }

    public void destroy() {
        // pass
    }
}
