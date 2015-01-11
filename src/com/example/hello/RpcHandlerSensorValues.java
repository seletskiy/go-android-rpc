package com.example.hello;

import go.rpc.Rpc;
import org.json.*;
import android.view.*;
import android.util.Log;
import android.content.Context;
import android.hardware.*;

public class RpcHandlerSensorValues implements RpcHandlerInterface {

    public JSONObject Handle(Context context, JSONObject payload) {
        JSONObject json = new JSONObject();
		try {
			int sensorId = Integer.parseInt(payload.getString("sensor_id"));

			SensorManager sensorManager = (SensorManager) context.getSystemService(Context.SENSOR_SERVICE);

			MyListener listener = new MyListener();

			sensorManager.registerListener(listener,
					sensorManager.getDefaultSensor(sensorId),
					SensorManager.SENSOR_DELAY_NORMAL);

			float[] values = listener.getValues();

			//java.lang.NullPointerException
			json.put("ax_value", Float.toString(values[0]));
			Log.v("ax_value", Float.toString(values[0]));
		} catch (Exception e) {
			Log.v("", e.toString());
		}
        return json;
    }


	public class MyListener implements SensorEventListener {
		float[] values;

		@Override
		public void onAccuracyChanged(Sensor sensor, int value) {}

		@Override
		public void onSensorChanged(SensorEvent event) {
			values = event.values;
		}

		public float[] getValues() {
			return values;
		}
	}
}
