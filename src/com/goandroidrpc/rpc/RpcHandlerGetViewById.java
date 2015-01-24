package com.goandroidrpc.rpc;

import org.json.JSONException;
import org.json.JSONObject;

import android.app.Activity;
import android.content.Context;
import android.view.View;

public class RpcHandlerGetViewById implements RpcHandlerInterface {
    public JSONObject Handle(
        Context context, JSONObject payload
    ) throws JSONException {
        // Pass resources (views) to go (goguibind)
        // @TODO: become recursive to allow nested ViewGroups
        JSONObject json = new JSONObject();
        String name;

        try {
            name = payload.getString("id");
        } catch(JSONException e) {
            json.put("error", e.toString());
            return json;
        }

        int id;
        try {
            id = R.id.class.getField(name).getInt(null);
        } catch(IllegalAccessException e) {
            json.put("error", e.toString());
            return json;
        } catch(NoSuchFieldException e) {
            // in case of string-encode numeric
            id = Integer.parseInt(name);
        }


        MainActivity activity = (MainActivity) context;
        try {
            View view;
            if (activity.orphanViews.containsKey(id)) {
                view = activity.orphanViews.get(id);
            } else {
                view = ((Activity) context).findViewById(id);
            }

            if (view == null) {
                json.put("error", "not found");
                return json;
            }

            json.put("id", Integer.toString(view.getId()));
            json.put("type", view.getClass().getName());
        } catch (Exception e) {
            json.put("error", e.toString());
        }

        return json;
    }

    public void destroy() {
        // pass
    }
}
