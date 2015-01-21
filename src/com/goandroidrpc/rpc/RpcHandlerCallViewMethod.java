package com.goandroidrpc.rpc;

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
import java.util.concurrent.Future;
import java.util.concurrent.FutureTask;
import java.util.concurrent.Callable;

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

            // @TODO: move package name to go code generator
            String packageName = "android.widget";
            if (payload.getString("type").equals("View")) {
                packageName = "android.view";
            }

            viewType = String.format(
                "%s.%s",
                packageName,
                payload.getString("type")
            );

            if (!payload.isNull("args")) {
                jsonArgs = payload.getJSONArray("args");
            } else {
                jsonArgs = new JSONArray();
            }
        } catch(Exception e) {
            // @TODO
            Log.v("!!! CVM init", e.toString());
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
            Log.v("!!! CVM methods", e.toString());
            return result;
        }

        final List<Object> requestedParams = new ArrayList<Object>();
        final List<Class> requestedParamTypes = new ArrayList<Class>();

        for (int i = 0; i < jsonArgs.length(); i++) {
            try {
                if (jsonArgs.get(i) instanceof Integer) {
                    // because get(i) will return Integer class, not int,
                    // which is not acceptable for methods, that wants int.
                    requestedParams.add(jsonArgs.getInt(i));
                    requestedParamTypes.add(int.class);
                } else if (jsonArgs.get(i) instanceof Double) {
                    // downcast to float
                    requestedParamTypes.add(float.class);
                    requestedParams.add((float) jsonArgs.getDouble(i));
                } else {
                    requestedParams.add(jsonArgs.get(i));
                    requestedParamTypes.add(jsonArgs.get(i).getClass());
                }
            } catch(Exception e) {
                Log.v("!!! CVM fill", e.toString());
            }
        }

        Method targetMethod = null;
        for (Method method : allMethods) {
            if (!method.getName().equals(methodName)) {
                continue;
            }

            Class[] paramTypes = method.getParameterTypes();

            if (paramTypes.length != requestedParams.size()) {
                continue;
            }

            boolean signatureMatched = true;
            int argIndex = 0;
            for (Class paramType : paramTypes) {
                if (!paramType.isAssignableFrom(requestedParamTypes.get(argIndex))) {
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

        if (targetMethod == null) {
            // @TODO
            return new JSONObject();
        }

        final Method methodToCall = targetMethod;

        FutureTask<String> futureResult = new FutureTask<String>(
            new Callable<String> () {
                @Override
                public String call() throws Exception {
                    Object result = methodToCall.invoke(viewObject, requestedParams.toArray());
                    if (
                            result != null && (
                                result instanceof Integer ||
                                result instanceof String ||
                                result instanceof Boolean
                            )
                    ) {
                        return result.toString();
                    } else {
                        return "";
                    }
                }
            }
        );
        activity.runOnUiThread(futureResult);
        try {
            result.put("result", futureResult.get());
        } catch (Exception e) {
            try {
                result.put("error", e.toString());
            } catch (JSONException je) {
                Log.v("!!! CVM put result", je.toString());
            }
        }

        return result;
    }
}
