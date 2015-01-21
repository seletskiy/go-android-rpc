package com.goandroidrpc.rpc;

import go.rpc.Rpc;

import org.json.*;

import android.view.*;
import android.util.Log;
import android.content.Context;
import android.app.Activity;

import java.util.List;
import java.util.ArrayList;
import java.lang.reflect.Constructor;
import java.lang.reflect.Method;

import java.util.concurrent.FutureTask;
import java.util.concurrent.Callable;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.CancellationException;
import java.lang.InterruptedException;

import java.util.Map;
import java.util.HashMap;


public class RpcHandlerCallViewMethod implements RpcHandlerInterface {
    protected Map<Integer, View> mOrphanViews;

    RpcHandlerCallViewMethod() {
        mOrphanViews = new HashMap<Integer, View>();
    }

    public JSONObject Handle(
        Context context, JSONObject request
    ) throws JSONException {
        JSONObject result = new JSONObject();

        try {
            String methodName = request.getString("viewMethod");
            String id = request.getString("id");
            if (methodName.equals("new")) {
                result = createView(
                    (Activity) context,
                    Integer.parseInt(id),
                    request.getString("type")
                );
            } else if (methodName.equals("attach")) {
                result = attachView(
                    (Activity) context,
                    Integer.parseInt(id),
                    Integer.parseInt(request.getString("viewGroupId"))
                );
            } else {
                JSONArray methodArgs = request.getJSONArray("args");

                View view;
                if (mOrphanViews.containsKey(Integer.parseInt(id))) {
                    try {
                        view = mOrphanViews.get(Integer.parseInt(id));
                    } catch(Exception e) {
                        result.put("error",
                            String.format(
                                "view with ID '%d' is not dynamically created",
                                id
                            )
                        );
                        return result;
                    }
                } else {
                    view = ((Activity) context).findViewById(
                        Integer.parseInt(id));
                }

                result = callMethod(
                    (Activity) context,
                    view,
                    methodName, request.getString("type"),
                    methodArgs
                );
            }
        } catch (JSONException e) {
            result.put("error",
                String.format("error in request: %s", e.getMessage())
            );
        }

        return result;
    }

    protected JSONObject createView(
        Activity activity,
        Integer id, String viewType
    ) throws JSONException {
        JSONObject result = new JSONObject();

        Class viewClass;

        try {
            viewClass = Class.forName(viewType);
        } catch(Exception e) {
            result.put(
                "error",
                String.format("class not found '%s'", viewType)
            );
            return result;
        }

        Constructor[] constructors = viewClass.getConstructors();

        View view;
        try {
            // @TODO: actually, find exact constructor.
            view = (View) constructors[0].newInstance(activity);
        } catch(Exception e) {
            // @TODO: properly handle exception
            result.put(
                "error",
                String.format("%s", e)
            );
            return result;
        }

        view.setId(id);

        mOrphanViews.put(id, view);

        return result;
    }

    protected JSONObject attachView(
        Activity activity,
        Integer id,
        Integer targetViewId
    ) throws JSONException {
        JSONObject result = new JSONObject();

        targetViewId = R.id.useless_layout;
        final ViewGroup targetView = (ViewGroup) activity.findViewById(
            targetViewId
        );

        final View orphanView;
        try {
            orphanView = mOrphanViews.get(id);
        } catch(Exception e) {
            result.put("error",
                String.format(
                    "view with ID '%d' is not either not exist or already attached",
                    id
                )
            );
            return result;
        }


        activity.runOnUiThread(new Runnable(){
            public void run() {
                try {
                    targetView.addView(orphanView);
                } catch(Exception e) {
                    // @TODO
                    Log.v("!!!", e.toString());
                }
            }
        });

        return result;
    }

    protected JSONObject callMethod(
        Activity activity,
        final View view, String methodName, String viewType,
        JSONArray methodArgs
    ) {
        JSONObject result = new JSONObject();

        Log.v("!!!", String.format("%s %s %s", view, methodName, viewType));
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

        for (int i = 0; i < methodArgs.length(); i++) {
            try {
                if (methodArgs.get(i) instanceof Integer) {
                    // because get(i) will return Integer class, not int,
                    // which is not acceptable for methods, that wants int.
                    requestedParams.add(methodArgs.getInt(i));
                    requestedParamTypes.add(int.class);
                } else if (methodArgs.get(i) instanceof Double) {
                    // downcast to float
                    requestedParamTypes.add(float.class);
                    requestedParams.add((float) methodArgs.getDouble(i));
                } else {
                    requestedParams.add(methodArgs.get(i));
                    requestedParamTypes.add(methodArgs.get(i).getClass());
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

        try {
            UIThreadCaller uiThreadCaller = new UIThreadCaller(activity);
            Object callerResult = uiThreadCaller.call(
                viewObject,
                methodToCall,
                requestedParams
            );
            Typer typer = new Typer();
            String strResult = "";
            if (typer.isSimpleType(callerResult)) {
                strResult = callerResult.toString();
            }
            result.put("result", strResult);
        } catch (Exception e) {
            try {
                result.put("error", e.toString());
            } catch (JSONException je) {
                Log.v("!!! UIThreadCaller failed to put error to JSON", je.toString());
            }
        }

        return result;
    }

    public class UIThreadCaller {
        protected Activity mActivity;

        UIThreadCaller(Activity activity) {
            mActivity = activity;
        }

        public Object call(
            final Object mViewObject,
            final Method mMethodToCall,
            final List<Object> mRequestedParams
        )
        throws
            InterruptedException,
            ExecutionException,
            CancellationException
        {
            FutureTask<Object> futureResult = new FutureTask<Object>(
                new Callable<Object> () {
                    @Override
                    public Object call() throws Exception {
                        Object result = mMethodToCall.invoke(
                            mViewObject,
                            mRequestedParams.toArray()
                        );
                        return result;
                    }
                }
            );
            mActivity.runOnUiThread(futureResult);
            Object result = futureResult.get();
            return result;
        }
    }

    public class Typer {
        public boolean isSimpleType(Object var) {
            if (
                var != null && (
                    var instanceof Integer ||
                    var instanceof String ||
                    var instanceof Boolean
                )
            ) {
                return true;
            }
            return false;
        }
    }
}
