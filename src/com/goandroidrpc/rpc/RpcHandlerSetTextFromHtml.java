package com.goandroidrpc.rpc;

import java.util.concurrent.Callable;

import org.json.JSONException;
import org.json.JSONObject;

import android.app.Activity;
import android.content.Context;
import android.text.Html;
import android.widget.TextView;

public class RpcHandlerSetTextFromHtml implements RpcHandlerInterface {
    public JSONObject Handle(
        Context context, JSONObject request
    ) throws JSONException {
        JSONObject result = new JSONObject();

        MainActivity activity = (MainActivity) context;

        String id;
        String html;
        try {
            html = request.getString("html");
            id = request.getString("id");

            TextView view;
            if (activity.orphanViews.containsKey(Integer.parseInt(id))) {
                view = (TextView) activity.orphanViews.get(Integer.parseInt(id));
            } else {
                view = (TextView) ((Activity) context).findViewById(Integer.parseInt(id));
            }


            final TextView viewToCall = view;
            final String htmlToSet = html;
            try {
                activity.uiThreadRunner.run(
                    new Callable<Object> () {
                        @Override
                        public Object call() throws Exception {
                            viewToCall.setText(Html.fromHtml(htmlToSet));
                            return null;
                        }
                    }
                );
            } catch(Exception e) {
                result.put("error",
                    String.format("error in request: %s", e.getMessage())
                );
            }
        } catch (JSONException e) {
            result.put("error",
                String.format("error in request: %s", e.getMessage())
            );
        }

        return result;
    }

    public void destroy() {
        // pass
    }
}
