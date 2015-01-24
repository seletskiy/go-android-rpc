package com.goandroidrpc.rpc;

import org.json.JSONObject;

import android.app.Activity;
import android.content.Context;
import android.content.res.Resources.NotFoundException;
import android.util.Log;
import android.view.View;
import android.view.ViewGroup;

public class RpcHandlerListViews implements RpcHandlerInterface {
    public JSONObject Handle(Context context, JSONObject payload) {
        // Pass resources (views) to go (goguibind)
        // @TODO: become recursive to allow nested ViewGroups
        JSONObject json = new JSONObject();
        try {
            String layoutName = payload.getString("layout");

            JSONObject jsonViews = new JSONObject();
            Activity activity = (Activity) context;
            ViewGroup rootView = (ViewGroup) activity.findViewById(
                R.id.class.getField(layoutName).getInt(null)
            );

            for(int i = 0; i < rootView.getChildCount(); i++) {
                View childView = rootView.getChildAt(i);
                int resId = childView.getId();
                Log.v("!!!", String.format("%s", resId));
                JSONObject jsonChild = new JSONObject();
                jsonChild.put("id", String.format("%d", resId));
                jsonChild.put("type", childView.getClass().getName());

                String resName;
                try {
                    resName = childView.getResources().getResourceEntryName(
                        childView.getId()
                    );
                } catch(NotFoundException e) {
                    resName = String.format("%s", resId);
                }

                jsonViews.put(resName, jsonChild);
            }

            json.put("resources", jsonViews);
        } catch (Exception e) {
            // @TODO: proper exception handling
            Log.v("!!!", String.format("%s", e));
        }

        return json;
    }

    public void destroy() {
        // pass
    }
}
