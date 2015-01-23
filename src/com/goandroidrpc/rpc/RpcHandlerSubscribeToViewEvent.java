package com.goandroidrpc.rpc;

import go.rpc.Rpc;

import org.json.*;

import android.app.Activity;
import android.view.*;
import android.os.Looper;
import android.util.Log;
import android.content.Context;

import java.util.List;
import java.util.ArrayList;
import java.util.concurrent.Callable;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.FutureTask;
import java.lang.Class;
import java.lang.Character;
import java.lang.reflect.Method;

import android.view.View.*;

public class RpcHandlerSubscribeToViewEvent implements RpcHandlerInterface {
    public JSONObject Handle(Context context, JSONObject payload) {
        JSONObject result = new JSONObject();

        String id;
        String eventName;
        String viewType;

        try {
            id = payload.getString("id");
            viewType = payload.getString("type");
        } catch(Exception e) {
            // @TODO
            Log.v("!!!", e.toString());
            return result;
        }

        try {
            eventName = payload.getString("event");
            String eventNameUcFirst = Character
                .toString(eventName.charAt(0))
                .toUpperCase()+eventName.substring(1);

            MainActivity activity = (MainActivity) context;
            ListenerFactory listenerFactory = new ListenerFactory(activity);

            Object listener;
            Method[] factoryMethods;
            Method[] viewMethods;

            final View view;
            if (activity.orphanViews.containsKey(Integer.parseInt(id))) {
                view = activity.orphanViews.get(Integer.parseInt(id));
            } else {
                view = ((Activity) context).findViewById(Integer.parseInt(id));
            }

            final Object viewObject;

            try {
                factoryMethods = listenerFactory.getClass().getDeclaredMethods();
                viewObject = Class.forName(viewType).cast(view);
                //viewMethods = Class.forName(viewType).getDeclaredMethods();
                viewMethods = View.class.getDeclaredMethods();
            } catch(Exception e) {
                // @TODO
                Log.v("!!!", e.toString());
                return result;
            }

            for (Method factoryMethod : factoryMethods) {
                if (!factoryMethod.getName().equals(eventName)) {
                    continue;
                }

                for (Method viewMethod : viewMethods) {
                    if (!viewMethod.getName().equals("set"+eventNameUcFirst+"Listener")) {
                        continue;
                    }

                    final List<Object> args = new ArrayList<Object>();
                    listener = factoryMethod.invoke(listenerFactory, args.toArray());
                    args.add(listener);
                    viewMethod.invoke(viewObject, args.toArray());
                }
            }
        } catch (Exception e) {
            Log.v("!!!", e.toString());
        }

        return result;
    }

    public class ListenerFactory {
        protected MainActivity mActivity;

        ListenerFactory(MainActivity activity) {
            mActivity = activity;
        }

        public OnClickListener onClick() {
            return new OnClickListener() {
                public void onClick(View v) {
                    JSONObject json = new JSONObject();
                    JSONObject jsonData = new JSONObject();

                    try {
                        json.put("event", "click");
                        jsonData.put("viewId", String.format("%d", v.getId()));
                        json.put("data", jsonData);
                    } catch (Exception e) {
                        // @TODO
                    }

                    mActivity.rpcBackend.call(json.toString());
                }
            };
        }

        public OnTouchListener onTouch() {
            return new OnTouchListener() {
                public boolean onTouch(View v, MotionEvent event) {
                    JSONObject json = new JSONObject();
                    JSONObject jsonData = new JSONObject();

                    try {
                        json.put("event", "touch");
                        jsonData.put("viewId", String.format("%d", v.getId()));
                        json.put("data", jsonData);
                    } catch (Exception e) {
                        // @TODO
                    }

                    Object result = mActivity.rpcBackend.call(json.toString());

                    return (Boolean)result;
                }
            };
        }
    }
}
