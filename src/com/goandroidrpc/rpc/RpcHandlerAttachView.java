package com.goandroidrpc.rpc;

import go.rpc.Rpc;

import org.json.*;

import android.view.*;
import android.util.Log;
import android.content.Context;

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

        activity.runOnUiThread(new Runnable(){
            public void run() {
                try {
                    targetView.addView(orphanView);
                } catch(Exception e) {
                    // @TODO
                    Log.v("!!!", e.toString());
                }
            }
        });

        return result;
    }
}
