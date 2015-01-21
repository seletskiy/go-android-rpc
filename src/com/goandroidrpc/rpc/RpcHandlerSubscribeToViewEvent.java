package com.goandroidrpc.rpc;

import go.rpc.Rpc;

import org.json.*;
import android.app.Activity;

import android.view.*;
import android.util.Log;
import android.content.Context;

import java.util.List;
import java.util.ArrayList;
import java.lang.Class;
import java.lang.Character;
import java.lang.reflect.Method;

import android.view.View.*;

public class RpcHandlerSubscribeToViewEvent implements RpcHandlerInterface {
    public JSONObject Handle(Context context, JSONObject payload) {
        JSONObject result = new JSONObject();

        Activity activity = (Activity) context;
        String viewId;
        String eventName;
        String viewType;

        try {
            viewId = payload.getString("id");
            viewType = String.format(
                "android.widget.%s",
                payload.getString("type")
            );
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

            ListenerFactory listenerFactory = new ListenerFactory();
            Object listener;
            Method[] factoryMethods;
            Method[] viewMethods;

            final View view = activity.findViewById(Integer.parseInt(viewId));
            Log.v("!!!", String.format("%s", view));
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
        public OnClickListener onClick() {
            return new OnClickListener() {
                public void onClick(View v) {
                    JSONObject json = new JSONObject();
                    JSONObject jsonData = new JSONObject();

                    try {
                        json.put("event", "click");

                        jsonData.put("view_id", String.format("%d", v.getId()));

                        json.put("data", jsonData);
                    } catch (Exception e) {
                        // @TODO
                    }

                    Rpc.CallBackend(json.toString());
                }
            };
        }

        public OnTouchListener onTouch() {
            return new OnTouchListener() {
                public boolean onTouch(View v, MotionEvent event) {
                    JSONObject json = new JSONObject();
                    JSONObject jsonData = new JSONObject();
                    Log.v("!!!", "got here");

                    try {
                        json.put("event", "touch");

                        jsonData.put("view_id", String.format("%d", v.getId()));

                        json.put("data", jsonData);
                    } catch (Exception e) {
                        // @TODO
                    }

                    Rpc.CallBackend(json.toString());

                    return true;
                }
            };
        }
    }
}
