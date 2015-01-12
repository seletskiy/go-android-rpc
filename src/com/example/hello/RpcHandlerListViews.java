package com.example.hello;

import go.rpc.Rpc;
import org.json.*;
import android.view.*;
import android.util.Log;
import android.app.Activity;
import android.content.Context;

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
                JSONObject jsonChild = new JSONObject();
                jsonChild.put("id", String.format("%d", resId));
                String resName = childView.getResources().getResourceEntryName(
                    childView.getId()
                );
                //jsonChild.put("name", resName);
                jsonViews.put(resName, jsonChild);
            }
            json.put("resources", jsonViews);
        } catch (Exception e) {
            // @TODO: proper exception handling
            System.out.println(e);
        }

        return json;
    }
}
