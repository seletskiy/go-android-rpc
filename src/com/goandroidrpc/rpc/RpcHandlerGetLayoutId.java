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

            Activity activity = (Activity) context;
            ViewGroup rootView = (ViewGroup) activity.findViewById(
                R.id.class.getField(layoutName).getInt(null)
            );

            json.put("layout_id", Integer.toString(rootView.getId()));
        } catch (Exception e) {
            // @TODO: proper exception handling
            System.out.println(e);
        }

        return json;
    }
}
