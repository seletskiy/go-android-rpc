package com.goandroidrpc.rpc;

import org.json.*;

import android.content.Context;

import android.media.*;

public class RpcHandlerControlMusicPlayback implements RpcHandlerInterface {
    public JSONObject Handle(
        Context context, JSONObject request
    ) throws JSONException {
        JSONObject result = new JSONObject();

        try {
            Integer resourceId = request.getInt("resource_id");
            String action = request.getString("action");
            Boolean loop = request.getBoolean("loop");

            MediaPlayer mediaPlayer = MediaPlayer.create(context, resourceId);

            mediaPlayer.setLooping(loop);

            if (action.equals("start")) {
                mediaPlayer.start();
            }

            if (action.equals("pause")) {
                mediaPlayer.pause();
            }

            if (action.equals("stop")) {
                mediaPlayer.stop();
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
