package com.example.groid;

import go.rpc.Rpc;
import org.json.*;
import android.view.*;
import android.util.Log;
import android.content.Context;
import android.hardware.*;
import java.util.List;
import java.util.ArrayList;
import java.lang.reflect.Field;

public class RpcHandlerGetSensorsList implements RpcHandlerInterface {
    public JSONObject Handle(Context context, JSONObject payload) {
        JSONObject json = new JSONObject();
        JSONObject jsonFields = new JSONObject();

        SensorManager sensorManager = (SensorManager) context.getSystemService(
            Context.SENSOR_SERVICE
        );

        try {
            for (Field field : Sensor.class.getFields()) {
                if (field.getName().startsWith("TYPE_")) {
                    int id = field.getInt(null);
                    if (sensorManager.getDefaultSensor(id) != null) {
                        jsonFields.put(
                            field.getName(),
                            String.format("%s", id)
                        );
                    }
                }
            }

            json.put("sensors", jsonFields);
        } catch (Exception e) {
            Log.v("", e.toString());
        }
        return json;
    }
}
