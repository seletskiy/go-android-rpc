/*
 * Copyright 2014 The Go Authors. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package com.example.hello;

import go.Go;
import go.rpc.Rpc;
import android.app.Activity;
import android.os.Bundle;
import android.hardware.Sensor;
import android.hardware.SensorManager;
import android.content.Context;
import android.view.View;
import android.view.ViewGroup;
import java.util.*;
import android.util.Log;
import org.json.JSONObject;
import org.json.JSONException;
import org.json.JSONArray;
import android.util.Log;
//import android.hardware.SensorEvent;

/*
 * MainActivity is the entry point for the libhello app.
 *
 * From here, the Go runtime is initialized and a Go function is
 * invoked via gobind language bindings.
 *
 * See example/libhello/README for details.
 */
public class MainActivity extends Activity {
    protected Frontend mFrontend;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        Go.init(getApplicationContext());
        setContentView(R.layout.useless_layout);

        mFrontend = new Frontend();
        Rpc.Link(mFrontend);

        // Pass resources (views) to go (goguibind)
        // @TODO: add multiple layouts support
        // @TODO: become recursive to allow nested ViewGroups
        try {
            JSONObject json = new JSONObject();
            JSONArray jsonViews = new JSONArray();
            ViewGroup rootView = (ViewGroup) findViewById(R.id.useless_layout);
            for(int i = 0; i < rootView.getChildCount(); i++) {
                View childView = rootView.getChildAt(i);
                int resId = childView.getId();
                JSONObject jsonChild = new JSONObject();
                jsonChild.put("id", String.format("%d", resId));
                String resName = childView.getResources().getResourceName(
                    childView.getId());
                Log.v("!!!", String.format("%s", resName));
                jsonChild.put("name", resName);
                jsonViews.put(jsonChild);
            }
            json.put("resources", jsonViews);
            Rpc.Handle(json.toString());
        } catch (JSONException e) {
            System.out.println(e);
            // @TODO: catch exception
        }
    }

    public void someOnClickHandler(View view) {
        // Kabloey
        JSONObject json = new JSONObject();
        try {
            json.put("id", String.format("%d", view.getId()));
            json.put("event", "onClick");
        } catch (JSONException e) {
            // @TODO: catch exception
        }

        Rpc.Handle(json.toString());
    }

    public class Frontend extends Rpc.Frontend.Stub {
        public void Handle(final String payload) {
            // @TODO: refactor that shit
            runOnUiThread(new Runnable() {
                @Override
                public void run() {
                    try {
                        JSONObject json = new JSONObject(payload);
                        View view = findViewById(Integer.parseInt(json.getString("id")));
                        if (json.getString("action").equals("hide")) {
                            view.setVisibility(View.INVISIBLE);
                        }
                        if (json.getString("action").equals("show")) {
                            view.setVisibility(View.VISIBLE);
                        }
                    } catch (JSONException e) {
                        System.out.println(e);
                        // @TODO: catch exception
                    }
                }
            });
        }
    }
}
