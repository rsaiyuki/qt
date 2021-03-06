#ifdef __cplusplus
extern "C" {
#endif

int QMaskGenerator_Seed(void* ptr);
void QMaskGenerator_DestroyQMaskGenerator(void* ptr);
void QWebSocket_Abort(void* ptr);
void QWebSocket_ConnectAboutToClose(void* ptr);
void QWebSocket_DisconnectAboutToClose(void* ptr);
void QWebSocket_ConnectBinaryFrameReceived(void* ptr);
void QWebSocket_DisconnectBinaryFrameReceived(void* ptr);
void QWebSocket_ConnectBinaryMessageReceived(void* ptr);
void QWebSocket_DisconnectBinaryMessageReceived(void* ptr);
char* QWebSocket_CloseReason(void* ptr);
void QWebSocket_ConnectConnected(void* ptr);
void QWebSocket_DisconnectConnected(void* ptr);
void QWebSocket_ConnectDisconnected(void* ptr);
void QWebSocket_DisconnectDisconnected(void* ptr);
int QWebSocket_Error(void* ptr);
char* QWebSocket_ErrorString(void* ptr);
int QWebSocket_Flush(void* ptr);
void QWebSocket_IgnoreSslErrors(void* ptr);
int QWebSocket_IsValid(void* ptr);
void* QWebSocket_MaskGenerator(void* ptr);
void QWebSocket_Open(void* ptr, void* url);
char* QWebSocket_Origin(void* ptr);
int QWebSocket_PauseMode(void* ptr);
char* QWebSocket_PeerName(void* ptr);
void QWebSocket_Ping(void* ptr, void* payload);
void QWebSocket_ConnectReadChannelFinished(void* ptr);
void QWebSocket_DisconnectReadChannelFinished(void* ptr);
char* QWebSocket_ResourceName(void* ptr);
void QWebSocket_Resume(void* ptr);
void QWebSocket_SetMaskGenerator(void* ptr, void* maskGenerator);
void QWebSocket_SetPauseMode(void* ptr, int pauseMode);
void QWebSocket_SetProxy(void* ptr, void* networkProxy);
void QWebSocket_SetSslConfiguration(void* ptr, void* sslConfiguration);
void QWebSocket_ConnectStateChanged(void* ptr);
void QWebSocket_DisconnectStateChanged(void* ptr);
void QWebSocket_ConnectTextFrameReceived(void* ptr);
void QWebSocket_DisconnectTextFrameReceived(void* ptr);
void QWebSocket_ConnectTextMessageReceived(void* ptr);
void QWebSocket_DisconnectTextMessageReceived(void* ptr);
void QWebSocket_DestroyQWebSocket(void* ptr);
void* QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator3(void* other);
void* QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator(char* origin);
void* QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator2(void* other);
int QWebSocketCorsAuthenticator_Allowed(void* ptr);
char* QWebSocketCorsAuthenticator_Origin(void* ptr);
void QWebSocketCorsAuthenticator_SetAllowed(void* ptr, int allowed);
void QWebSocketCorsAuthenticator_Swap(void* ptr, void* other);
void QWebSocketCorsAuthenticator_DestroyQWebSocketCorsAuthenticator(void* ptr);
void* QWebSocketServer_NewQWebSocketServer(char* serverName, int secureMode, void* parent);
void QWebSocketServer_ConnectAcceptError(void* ptr);
void QWebSocketServer_DisconnectAcceptError(void* ptr);
void QWebSocketServer_Close(void* ptr);
void QWebSocketServer_ConnectClosed(void* ptr);
void QWebSocketServer_DisconnectClosed(void* ptr);
char* QWebSocketServer_ErrorString(void* ptr);
int QWebSocketServer_HasPendingConnections(void* ptr);
int QWebSocketServer_IsListening(void* ptr);
int QWebSocketServer_MaxPendingConnections(void* ptr);
void QWebSocketServer_ConnectNewConnection(void* ptr);
void QWebSocketServer_DisconnectNewConnection(void* ptr);
void* QWebSocketServer_NextPendingConnection(void* ptr);
void QWebSocketServer_ConnectOriginAuthenticationRequired(void* ptr);
void QWebSocketServer_DisconnectOriginAuthenticationRequired(void* ptr);
void QWebSocketServer_PauseAccepting(void* ptr);
void QWebSocketServer_ResumeAccepting(void* ptr);
int QWebSocketServer_SecureMode(void* ptr);
char* QWebSocketServer_ServerName(void* ptr);
void QWebSocketServer_SetMaxPendingConnections(void* ptr, int numConnections);
void QWebSocketServer_SetProxy(void* ptr, void* networkProxy);
void QWebSocketServer_SetServerName(void* ptr, char* serverName);
int QWebSocketServer_SetSocketDescriptor(void* ptr, int socketDescriptor);
void QWebSocketServer_SetSslConfiguration(void* ptr, void* sslConfiguration);
int QWebSocketServer_SocketDescriptor(void* ptr);
void QWebSocketServer_DestroyQWebSocketServer(void* ptr);

#ifdef __cplusplus
}
#endif