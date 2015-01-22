package com.goandroidrpc.rpc;

import go.rpc.Rpc;

import org.json.*;

import android.view.*;
import android.content.Context;

import java.lang.reflect.Constructor;

public class RpcHandlerCreateView implements RpcHandlerInterface {
    public JSONObject Handle(
        Context context, JSONObject request
    ) throws JSONException {
        JSONObject result = new JSONObject();

        MainActivity activity = (MainActivity) context;

        try {
            result = createView(
                activity,
                Integer.parseInt(request.getString("id")),
                request.getString("type")
            );
        } catch (JSONException e) {
            result.put("error",
                String.format("error in request: %s", e.getMessage())
            );
        }

        return result;
    }

    protected JSONObject createView(
        MainActivity activity,
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

        activity.orphanViews.put(id, view);

        return result;
    }
}
