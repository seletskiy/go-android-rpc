package com.goandroidrpc.rpc;

public class JsonConverter {
    public String Convert(Object var) {
        if (!isSimpleType(var)) {
            return "";
        }

        return var.toString();
    }

    public boolean isSimpleType(Object var) {
        if (
            var != null && (
                var instanceof Integer ||
                var instanceof String ||
                var instanceof Boolean
            )
        ) {
            return true;
        }
        return false;
    }
}
