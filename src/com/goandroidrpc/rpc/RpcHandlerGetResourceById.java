package com.goandroidrpc.rpc;

import org.json.*;
import android.content.Context;

public class RpcHandlerGetResourceById implements RpcHandlerInterface {
    public JSONObject Handle(Context context, JSONObject payload) {
        JSONObject json = new JSONObject();
        try {
            String resourceName = payload.getString("resource");

            int resourceId = context.getResources()
                .getIdentifier(resourceName, "id", context.getPackageName());

            json.put("resource_id", Integer.toString(resourceId));
        } catch (Exception e) {
            // @TODO: proper exception handling
            System.out.println(e);
        }

        return json;
    }

    public void destroy() {
        // pass
    }
}
