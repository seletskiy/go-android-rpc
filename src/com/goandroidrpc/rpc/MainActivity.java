package com.goandroidrpc.rpc;

import go.Go;
import go.rpc.Rpc;

import java.util.HashMap;
import java.util.Map;

import org.json.JSONObject;

import android.app.Activity;
import android.content.Context;
import android.media.AudioFormat;
import android.media.AudioManager;
import android.media.AudioTrack;
import android.os.Bundle;
import android.util.Log;
import android.view.View;


/*
 * MainActivity is the entry point for RPC endpoint of go-android-rpc.
 *
 * From here, the Go runtime is initialized and a Go function is
 * invoked via gobind language bindings.
 */
public class MainActivity extends Activity {
    public Map<Integer, View> orphanViews;

    public RpcBackendCaller rpcBackend;
    public RpcFrontend rpcFrontend;

    public UIThreadRunner uiThreadRunner;

    public MainActivity() {
        orphanViews = new HashMap<Integer, View>();
        uiThreadRunner = new UIThreadRunner(this);
        rpcBackend = new RpcBackendCaller(uiThreadRunner);
        playSound();
    }

    private void playSound() {
        Log.v("!!!", "Sound started");
        int mBufferSize = AudioTrack.getMinBufferSize(44100,
                AudioFormat.CHANNEL_OUT_MONO,    
                AudioFormat.ENCODING_PCM_8BIT);

        AudioTrack mAudioTrack = new
            AudioTrack(AudioManager.STREAM_MUSIC, 44100,
                    AudioFormat.CHANNEL_OUT_MONO,    
                    AudioFormat.ENCODING_PCM_8BIT,
                    mBufferSize,
                    AudioTrack.MODE_STREAM);

        double[] mSound = new double[5*44100];

        short[] mBuffer = new short[5*44100];

        for (int i = 0; i < mSound.length; i++) {
                mSound[i] = Math.sin((2.0*Math.PI * 440.0/44100.0*(double)i));
                mBuffer[i] = (short) (mSound[i]*Short.MAX_VALUE);
            }

        mAudioTrack.setStereoVolume(1.0f, 1.0f);
        mAudioTrack.play();

        mAudioTrack.write(mBuffer, 0, mSound.length);
        mAudioTrack.stop();
        mAudioTrack.release();
        Log.v("!!!", "Sound stopped");
    }

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        Go.init(getApplicationContext());
        setContentView(R.layout.main_layout);

        rpcFrontend = new RpcFrontend(this);
        Rpc.StartBackend(rpcFrontend);
    }

    @Override
    protected void onDestroy() {
        super.onDestroy();
        rpcBackend.destroy();
        // no need for now
        //rpcFrontend.destroy();
    }

    @Override
    protected void onStop() {
        super.onStop();
        rpcBackend.stop();
    }

    public class RpcFrontend extends Rpc.Frontend.Stub {
        protected Context mContext;
        protected Map<String, RpcHandlerInterface> mHandlers;

        RpcFrontend(Context context) {
            mContext = context;
            mHandlers = new HashMap<String, RpcHandlerInterface>();
        }

        public String CallFrontend(final String payload) {
            try {
                JSONObject json = new JSONObject(payload);

                String handlerName = String.format(
                    "%s.RpcHandler%s",
                    this.getClass().getPackage().getName(),
                    json.getString("method")
                );

                // @TODO: handle errors

                RpcHandlerInterface handler = mHandlers.get(handlerName);
                if (handler == null) {
                    handler = (RpcHandlerInterface) Class.forName(
                        handlerName
                    ).newInstance();

                    mHandlers.put(handlerName, handler);
                }

                return handler.Handle(mContext, json).toString();
            } catch (Exception e) {
                // @TODO: proper exception handling
                Log.v("!!!", e.toString());
            }

            // @TODO: properly return json error
            return "error";
        }

        public void destroy() {
            for (Map.Entry<String, RpcHandlerInterface> handler : mHandlers.entrySet()) {
                handler.getValue().destroy();
            }
        }
    }
}
