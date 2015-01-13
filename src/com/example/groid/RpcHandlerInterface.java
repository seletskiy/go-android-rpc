package com.example.groid;

import android.content.Context;
import org.json.JSONObject;

public interface RpcHandlerInterface {
    JSONObject Handle(Context context, JSONObject request);
}
