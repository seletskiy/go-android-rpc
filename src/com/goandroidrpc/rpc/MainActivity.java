package com.goandroidrpc.rpc;

import go.Go;
import go.rpc.Rpc;
import android.app.Activity;
import android.os.Bundle;
import android.content.Context;
import java.util.*;
import org.json.*;
import android.util.Log;
import android.view.*;
import com.goandroidrpc.rpc.UIThreadRunner;

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
    protected void onPause() {
        super.onResume();
        rpcBackend.destroy();
    }

    @Override
    protected void onResume() {
        super.onResume();
        uiThreadRunner = new UIThreadRunner(this);
        rpcBackend = new RpcBackendCaller(uiThreadRunner);
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

                Log.v("!!! src/com/goandroidrpc/rpc/MainActivity.java:72", String.format("%s", json));

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
