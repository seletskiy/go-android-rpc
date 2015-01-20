package com.goandroidrpc.rpc;

import go.android.Rpc;
import org.json.*;
import android.util.Log;
import android.content.Context;
import android.hardware.*;
import java.util.List;
import java.util.ArrayList;

public class RpcHandlerSubscribeToSensorValues implements RpcHandlerInterface {
    protected List<Listener> mListeners;

    RpcHandlerSubscribeToSensorValues() {
        mListeners = new ArrayList<Listener>();
    }

    public JSONObject Handle(Context context, JSONObject payload) {
        JSONObject json = new JSONObject();
        try {
            int sensorId = Integer.parseInt(payload.getString("sensor_id"));

            SensorManager sensorManager = (SensorManager) context.getSystemService(
                Context.SENSOR_SERVICE
            );

            Listener listener = new Listener(sensorId);

            sensorManager.registerListener(listener,
                sensorManager.getDefaultSensor(sensorId),
                SensorManager.SENSOR_DELAY_NORMAL
            );

            mListeners.add(listener);
        } catch (Exception e) {
            Log.v("!!!", e.toString());
        }
        return json;
    }


    public class Listener implements SensorEventListener {
        protected int mSensorId;

        Listener(int id) {
            mSensorId = id;
        }

        @Override
        public void onAccuracyChanged(Sensor sensor, int value) {}

        @Override
        public void onSensorChanged(SensorEvent event) {
            JSONObject json = new JSONObject();
            JSONArray jsonValues = new JSONArray();
            JSONObject jsonData = new JSONObject();

            try {
                json.put("event", "sensorChange");

                jsonData.put("sensor_id", String.format("%s", mSensorId));
                for (float value : event.values) {
                    jsonValues.put(value);
                }

                jsonData.put("values", jsonValues);

                json.put("data", jsonData);
            } catch (Exception e) {
                // @TODO
            }

            Rpc.CallBackend(json.toString());
        }
    }
}
