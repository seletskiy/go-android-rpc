package com.goandroidrpc.rpc;

import java.lang.reflect.Constructor;

import org.json.JSONException;
import org.json.JSONObject;

import android.content.Context;
import android.util.Log;
import android.view.View;
import android.view.LayoutInflater;

public class RpcHandlerInflateView implements RpcHandlerInterface {
    public JSONObject Handle(
        Context context, JSONObject request
    ) throws JSONException {
        JSONObject result = new JSONObject();

        MainActivity activity = (MainActivity) context;

        try {
            LayoutInflater inflater = (LayoutInflater)activity.getSystemService(
                MainActivity.LAYOUT_INFLATER_SERVICE
            );

            View view = inflater.inflate(
                Integer.parseInt(request.getString("res")),
                null
            );
            Integer id = Integer.parseInt(request.getString("id"));

            view.setId(id);
            activity.orphanViews.put(id, view);
        } catch (JSONException e) {
            result.put("error",
                String.format("error in request: %s", e.getMessage())
            );
        }

        return result;
    }

    public void destroy() {
        // pass
    }
}
