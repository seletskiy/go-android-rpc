package com.example.hello;

import android.content.Context;
import org.json.JSONObject;;

public interface RpcHandlerInterface {
    JSONObject Handle(Context context, JSONObject request);
}
