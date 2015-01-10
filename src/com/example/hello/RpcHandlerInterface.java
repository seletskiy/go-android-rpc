package com.example.hello;

import android.app.Activity;
import org.json.JSONObject;;

public interface RpcHandlerInterface {
    JSONObject Handle(Activity activity, JSONObject request);
}
