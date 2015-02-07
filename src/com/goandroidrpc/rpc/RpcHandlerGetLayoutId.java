package com.goandroidrpc.rpc;

import go.rpc.Rpc;
import org.json.*;
import android.view.*;
import android.util.Log;
import android.app.Activity;
import android.content.Context;

public class RpcHandlerGetLayoutId implements RpcHandlerInterface {
    public JSONObject Handle(Context context, JSONObject payload) {
        JSONObject json = new JSONObject();
        try {
            String layoutName = payload.getString("layout");
            json.put("layout_id",
                Integer.toString(
                    R.layout.class.getField(layoutName).getInt(null)
                )
            );
        } catch (Exception e) {
            Log.v("!!! src/com/goandroidrpc/rpc/RpcHandlerGetLayoutById.java:21", String.format("%s", e));
        }

        return json;
    }

    public void destroy() {
        // pass
    }
}
