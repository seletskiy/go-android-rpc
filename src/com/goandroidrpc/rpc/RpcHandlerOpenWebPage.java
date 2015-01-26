package com.goandroidrpc.rpc;

import org.json.JSONException;
import org.json.JSONObject;

import android.app.Activity;
import android.content.Context;
import android.content.Intent;
import android.net.Uri;

public class RpcHandlerOpenWebPage implements RpcHandlerInterface {
    public JSONObject Handle(
        Context context, JSONObject payload
    ) throws JSONException {
        JSONObject json = new JSONObject();
        String url;

        try {
            url = payload.getString("url");
        } catch(JSONException e) {
            json.put("error", e.toString());
            return json;
        }

        Activity activity = (Activity) context;
        try {
            activity.startActivity(new Intent(Intent.ACTION_VIEW).setData(Uri.parse(url)));
        } catch(Exception e) {
            // @TODO
        }
        return json;
    }

    public void destroy() {
        // pass
    }
}
