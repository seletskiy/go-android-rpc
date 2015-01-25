package com.goandroidrpc.rpc;

import java.util.concurrent.Callable;

import org.json.JSONException;
import org.json.JSONObject;

import android.content.Context;
import android.util.Log;
import android.view.View;
import android.view.ViewGroup;

public class RpcHandlerRemoveView implements RpcHandlerInterface {
    public JSONObject Handle(
        Context context, JSONObject request
    ) throws JSONException {
        JSONObject result = new JSONObject();

        MainActivity activity = (MainActivity) context;

        try {
            result = removeView(
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

    protected JSONObject removeView(
        final MainActivity activity,
        final Integer id,
        Integer viewGroupId
    ) throws JSONException {
        JSONObject result = new JSONObject();

        final ViewGroup viewGroup = (ViewGroup) activity.findViewById(viewGroupId);
        if (viewGroup == null) {
            result.put("error", "view group is not found");
            return result;
        }

        final View viewToRemove = (View) viewGroup.findViewById(id);
        if (viewToRemove == null) {
            result.put("error", "view to remove is not found");
            return result;
        }

        final ViewGroup parentView = (ViewGroup) viewToRemove.getParent();

        try {
            activity.uiThreadRunner.run(
                new Callable<Object> () {
                    @Override
                    public Object call() throws Exception {
                        try {
                            parentView.removeView(viewToRemove);
                        } catch(Exception e) {
                            // @TODO
                            Log.v("!!! src/com/goandroidrpc/rpc/RpcHandlerRemoveView.java:55", String.format("%s", e));
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
