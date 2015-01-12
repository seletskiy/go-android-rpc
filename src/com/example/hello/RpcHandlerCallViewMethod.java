package com.example.hello;

import go.rpc.Rpc;
import org.json.*;
import android.view.*;
import android.util.Log;
import android.content.Context;
import android.hardware.*;
import android.app.Activity;
import java.util.List;
import java.util.ArrayList;
import java.lang.reflect.Field;
import android.widget.*;
import java.lang.reflect.Method;
import java.lang.reflect.Type;

public class RpcHandlerCallViewMethod implements RpcHandlerInterface {
    public JSONObject Handle(Context context, JSONObject payload) {
        JSONObject result = new JSONObject();

        Activity activity = (Activity) context;
        String id;
        String viewType;
        String methodName;
        JSONArray jsonArgs;
        try {
            methodName = payload.getString("viewMethod");
            id = payload.getString("id");
            viewType = String.format(
                "android.widget.%s",
                payload.getString("type")
            );

            jsonArgs = payload.getJSONArray("args");
        } catch(Exception e) {
            // @TODO
            Log.v("!!!", e.toString());
            return result;
        }

        final View view = activity.findViewById(Integer.parseInt(id));
        final Object viewObject;

        Method[] allMethods;
        try {
            viewObject = Class.forName(viewType).cast(view);
            allMethods = Class.forName(viewType).getDeclaredMethods();
        } catch(Exception e) {
            // @TODO
            Log.v("!!!", e.toString());
            return result;
        }

        final List<Object> args = new ArrayList<Object>();

        for (int i = 0; i < jsonArgs.length(); i++) {
            try {
                args.add(jsonArgs.get(i));
            } catch(Exception e) {
                Log.v("!!!", e.toString());
            }
        }

        Method targetMethod = null;
        for (Method method : allMethods) {
            if (!method.getName().equals(methodName)) {
                continue;
            }

            Class[] paramTypes = method.getParameterTypes();

            if (paramTypes.length != args.size()) {
                continue;
            }

            boolean signatureMatched = true;
            int argIndex = 0;
            for (Class paramType : paramTypes) {
                if (!paramType.isAssignableFrom(args.get(argIndex).getClass())) {
                    signatureMatched = false;
                    break;
                }

                argIndex++;
            }

            if (signatureMatched) {
                targetMethod = method;
                break;
            }
        }

        //Log.v("!!!", String.format("%s", targetMethod));

        if (targetMethod == null) {
            // @TODO
            return new JSONObject();
        }


        final Method methodToCall = targetMethod;

        activity.runOnUiThread(new Runnable(){
            public void run() {
                try {
                    methodToCall.invoke(viewObject, args.toArray());
                } catch(Exception e) {
                    // @TODO
                    Log.v("!!!", e.toString());
                }

                view.invalidate();
            }
        });

        return result;
    }
}
