#ifdef __cplusplus
extern "C" {
#endif

int QAccelerometer_AccelerationMode(void* ptr);
void* QAccelerometer_Reading(void* ptr);
void* QAccelerometer_NewQAccelerometer(void* parent);
void QAccelerometer_ConnectAccelerationModeChanged(void* ptr);
void QAccelerometer_DisconnectAccelerationModeChanged(void* ptr);
void QAccelerometer_SetAccelerationMode(void* ptr, int accelerationMode);
void QAccelerometer_DestroyQAccelerometer(void* ptr);
int QAccelerometerFilter_Filter(void* ptr, void* reading);
double QAccelerometerReading_X(void* ptr);
double QAccelerometerReading_Y(void* ptr);
double QAccelerometerReading_Z(void* ptr);
void QAccelerometerReading_SetX(void* ptr, double x);
void QAccelerometerReading_SetY(void* ptr, double y);
void QAccelerometerReading_SetZ(void* ptr, double z);
void* QAltimeter_Reading(void* ptr);
void* QAltimeter_NewQAltimeter(void* parent);
void QAltimeter_DestroyQAltimeter(void* ptr);
int QAltimeterFilter_Filter(void* ptr, void* reading);
double QAltimeterReading_Altitude(void* ptr);
void QAltimeterReading_SetAltitude(void* ptr, double altitude);
int QAmbientLightFilter_Filter(void* ptr, void* reading);
int QAmbientLightReading_LightLevel(void* ptr);
void QAmbientLightReading_SetLightLevel(void* ptr, int lightLevel);
void* QAmbientLightSensor_Reading(void* ptr);
void* QAmbientLightSensor_NewQAmbientLightSensor(void* parent);
void QAmbientLightSensor_DestroyQAmbientLightSensor(void* ptr);
int QAmbientTemperatureFilter_Filter(void* ptr, void* reading);
double QAmbientTemperatureReading_Temperature(void* ptr);
void QAmbientTemperatureReading_SetTemperature(void* ptr, double temperature);
void* QAmbientTemperatureSensor_Reading(void* ptr);
void* QAmbientTemperatureSensor_NewQAmbientTemperatureSensor(void* parent);
void QAmbientTemperatureSensor_DestroyQAmbientTemperatureSensor(void* ptr);
void* QCompass_Reading(void* ptr);
void* QCompass_NewQCompass(void* parent);
void QCompass_DestroyQCompass(void* ptr);
int QCompassFilter_Filter(void* ptr, void* reading);
double QCompassReading_Azimuth(void* ptr);
double QCompassReading_CalibrationLevel(void* ptr);
void QCompassReading_SetAzimuth(void* ptr, double azimuth);
void QCompassReading_SetCalibrationLevel(void* ptr, double calibrationLevel);
int QDistanceFilter_Filter(void* ptr, void* reading);
double QDistanceReading_Distance(void* ptr);
void QDistanceReading_SetDistance(void* ptr, double distance);
void* QDistanceSensor_Reading(void* ptr);
void* QDistanceSensor_NewQDistanceSensor(void* parent);
void QDistanceSensor_DestroyQDistanceSensor(void* ptr);
void* QGyroscope_Reading(void* ptr);
void* QGyroscope_NewQGyroscope(void* parent);
void QGyroscope_DestroyQGyroscope(void* ptr);
int QGyroscopeFilter_Filter(void* ptr, void* reading);
double QGyroscopeReading_X(void* ptr);
double QGyroscopeReading_Y(void* ptr);
double QGyroscopeReading_Z(void* ptr);
void QGyroscopeReading_SetX(void* ptr, double x);
void QGyroscopeReading_SetY(void* ptr, double y);
void QGyroscopeReading_SetZ(void* ptr, double z);
int QHolsterFilter_Filter(void* ptr, void* reading);
int QHolsterReading_Holstered(void* ptr);
void QHolsterReading_SetHolstered(void* ptr, int holstered);
void* QHolsterSensor_Reading(void* ptr);
void* QHolsterSensor_NewQHolsterSensor(void* parent);
void QHolsterSensor_DestroyQHolsterSensor(void* ptr);
int QIRProximityFilter_Filter(void* ptr, void* reading);
double QIRProximityReading_Reflectance(void* ptr);
void QIRProximityReading_SetReflectance(void* ptr, double reflectance);
void* QIRProximitySensor_Reading(void* ptr);
void* QIRProximitySensor_NewQIRProximitySensor(void* parent);
void QIRProximitySensor_DestroyQIRProximitySensor(void* ptr);
int QLightFilter_Filter(void* ptr, void* reading);
double QLightReading_Lux(void* ptr);
void QLightReading_SetLux(void* ptr, double lux);
double QLightSensor_FieldOfView(void* ptr);
void* QLightSensor_Reading(void* ptr);
void* QLightSensor_NewQLightSensor(void* parent);
void QLightSensor_SetFieldOfView(void* ptr, double fieldOfView);
void QLightSensor_DestroyQLightSensor(void* ptr);
void* QMagnetometer_Reading(void* ptr);
int QMagnetometer_ReturnGeoValues(void* ptr);
void QMagnetometer_SetReturnGeoValues(void* ptr, int returnGeoValues);
void* QMagnetometer_NewQMagnetometer(void* parent);
void QMagnetometer_ConnectReturnGeoValuesChanged(void* ptr);
void QMagnetometer_DisconnectReturnGeoValuesChanged(void* ptr);
void QMagnetometer_DestroyQMagnetometer(void* ptr);
int QMagnetometerFilter_Filter(void* ptr, void* reading);
double QMagnetometerReading_CalibrationLevel(void* ptr);
double QMagnetometerReading_X(void* ptr);
double QMagnetometerReading_Y(void* ptr);
double QMagnetometerReading_Z(void* ptr);
void QMagnetometerReading_SetCalibrationLevel(void* ptr, double calibrationLevel);
void QMagnetometerReading_SetX(void* ptr, double x);
void QMagnetometerReading_SetY(void* ptr, double y);
void QMagnetometerReading_SetZ(void* ptr, double z);
int QOrientationFilter_Filter(void* ptr, void* reading);
int QOrientationReading_Orientation(void* ptr);
void QOrientationReading_SetOrientation(void* ptr, int orientation);
void* QOrientationSensor_Reading(void* ptr);
void* QOrientationSensor_NewQOrientationSensor(void* parent);
void QOrientationSensor_DestroyQOrientationSensor(void* ptr);
int QPressureFilter_Filter(void* ptr, void* reading);
double QPressureReading_Pressure(void* ptr);
double QPressureReading_Temperature(void* ptr);
void QPressureReading_SetPressure(void* ptr, double pressure);
void QPressureReading_SetTemperature(void* ptr, double temperature);
void* QPressureSensor_Reading(void* ptr);
void* QPressureSensor_NewQPressureSensor(void* parent);
void QPressureSensor_DestroyQPressureSensor(void* ptr);
int QProximityFilter_Filter(void* ptr, void* reading);
int QProximityReading_Close(void* ptr);
void QProximityReading_SetClose(void* ptr, int close);
void* QProximitySensor_Reading(void* ptr);
void* QProximitySensor_NewQProximitySensor(void* parent);
void QProximitySensor_DestroyQProximitySensor(void* ptr);
int QRotationFilter_Filter(void* ptr, void* reading);
double QRotationReading_X(void* ptr);
double QRotationReading_Y(void* ptr);
double QRotationReading_Z(void* ptr);
void QRotationReading_SetFromEuler(void* ptr, double x, double y, double z);
int QRotationSensor_HasZ(void* ptr);
void* QRotationSensor_Reading(void* ptr);
void* QRotationSensor_NewQRotationSensor(void* parent);
void QRotationSensor_ConnectHasZChanged(void* ptr);
void QRotationSensor_DisconnectHasZChanged(void* ptr);
void QRotationSensor_SetHasZ(void* ptr, int hasZ);
void QRotationSensor_DestroyQRotationSensor(void* ptr);
int QSensor_AxesOrientationMode(void* ptr);
int QSensor_BufferSize(void* ptr);
int QSensor_CurrentOrientation(void* ptr);
int QSensor_DataRate(void* ptr);
char* QSensor_Description(void* ptr);
int QSensor_EfficientBufferSize(void* ptr);
int QSensor_Error(void* ptr);
void* QSensor_Identifier(void* ptr);
int QSensor_IsActive(void* ptr);
int QSensor_IsAlwaysOn(void* ptr);
int QSensor_IsBusy(void* ptr);
int QSensor_IsConnectedToBackend(void* ptr);
int QSensor_MaxBufferSize(void* ptr);
int QSensor_OutputRange(void* ptr);
void* QSensor_Reading(void* ptr);
void QSensor_SetActive(void* ptr, int active);
void QSensor_SetAlwaysOn(void* ptr, int alwaysOn);
void QSensor_SetAxesOrientationMode(void* ptr, int axesOrientationMode);
void QSensor_SetBufferSize(void* ptr, int bufferSize);
void QSensor_SetDataRate(void* ptr, int rate);
void QSensor_SetIdentifier(void* ptr, void* identifier);
void QSensor_SetOutputRange(void* ptr, int index);
void QSensor_SetUserOrientation(void* ptr, int userOrientation);
int QSensor_SkipDuplicates(void* ptr);
void* QSensor_Type(void* ptr);
int QSensor_UserOrientation(void* ptr);
void* QSensor_NewQSensor(void* ty, void* parent);
void QSensor_ConnectActiveChanged(void* ptr);
void QSensor_DisconnectActiveChanged(void* ptr);
void QSensor_AddFilter(void* ptr, void* filter);
void QSensor_ConnectAlwaysOnChanged(void* ptr);
void QSensor_DisconnectAlwaysOnChanged(void* ptr);
void QSensor_ConnectAvailableSensorsChanged(void* ptr);
void QSensor_DisconnectAvailableSensorsChanged(void* ptr);
void QSensor_ConnectAxesOrientationModeChanged(void* ptr);
void QSensor_DisconnectAxesOrientationModeChanged(void* ptr);
void QSensor_ConnectBufferSizeChanged(void* ptr);
void QSensor_DisconnectBufferSizeChanged(void* ptr);
void QSensor_ConnectBusyChanged(void* ptr);
void QSensor_DisconnectBusyChanged(void* ptr);
int QSensor_ConnectToBackend(void* ptr);
void QSensor_ConnectCurrentOrientationChanged(void* ptr);
void QSensor_DisconnectCurrentOrientationChanged(void* ptr);
void QSensor_ConnectDataRateChanged(void* ptr);
void QSensor_DisconnectDataRateChanged(void* ptr);
void* QSensor_QSensor_DefaultSensorForType(void* ty);
void QSensor_ConnectEfficientBufferSizeChanged(void* ptr);
void QSensor_DisconnectEfficientBufferSizeChanged(void* ptr);
int QSensor_IsFeatureSupported(void* ptr, int feature);
void QSensor_ConnectMaxBufferSizeChanged(void* ptr);
void QSensor_DisconnectMaxBufferSizeChanged(void* ptr);
void QSensor_ConnectReadingChanged(void* ptr);
void QSensor_DisconnectReadingChanged(void* ptr);
void QSensor_RemoveFilter(void* ptr, void* filter);
void QSensor_ConnectSensorError(void* ptr);
void QSensor_DisconnectSensorError(void* ptr);
void QSensor_SetCurrentOrientation(void* ptr, int currentOrientation);
void QSensor_SetEfficientBufferSize(void* ptr, int efficientBufferSize);
void QSensor_SetMaxBufferSize(void* ptr, int maxBufferSize);
void QSensor_SetSkipDuplicates(void* ptr, int skipDuplicates);
void QSensor_ConnectSkipDuplicatesChanged(void* ptr);
void QSensor_DisconnectSkipDuplicatesChanged(void* ptr);
int QSensor_Start(void* ptr);
void QSensor_Stop(void* ptr);
void QSensor_ConnectUserOrientationChanged(void* ptr);
void QSensor_DisconnectUserOrientationChanged(void* ptr);
void QSensor_DestroyQSensor(void* ptr);
void QSensorBackend_AddDataRate(void* ptr, double min, double max);
int QSensorBackend_IsFeatureSupported(void* ptr, int feature);
void QSensorBackend_SensorBusy(void* ptr);
void QSensorBackend_SensorError(void* ptr, int error);
void QSensorBackend_AddOutputRange(void* ptr, double min, double max, double accuracy);
void QSensorBackend_NewReadingAvailable(void* ptr);
void* QSensorBackend_Reading(void* ptr);
void* QSensorBackend_Sensor(void* ptr);
void QSensorBackend_SensorStopped(void* ptr);
void QSensorBackend_SetDataRates(void* ptr, void* otherSensor);
void QSensorBackend_SetDescription(void* ptr, char* description);
void QSensorBackend_Start(void* ptr);
void QSensorBackend_Stop(void* ptr);
void* QSensorBackendFactory_CreateBackend(void* ptr, void* sensor);
void QSensorChangesInterface_SensorsChanged(void* ptr);
int QSensorFilter_Filter(void* ptr, void* reading);
void* QSensorGesture_NewQSensorGesture(char* ids, void* parent);
char* QSensorGesture_GestureSignals(void* ptr);
char* QSensorGesture_InvalidIds(void* ptr);
int QSensorGesture_IsActive(void* ptr);
void QSensorGesture_StartDetection(void* ptr);
void QSensorGesture_StopDetection(void* ptr);
char* QSensorGesture_ValidIds(void* ptr);
void QSensorGesture_DestroyQSensorGesture(void* ptr);
void* QSensorGestureManager_NewQSensorGestureManager(void* parent);
char* QSensorGestureManager_GestureIds(void* ptr);
void QSensorGestureManager_ConnectNewSensorGestureAvailable(void* ptr);
void QSensorGestureManager_DisconnectNewSensorGestureAvailable(void* ptr);
char* QSensorGestureManager_RecognizerSignals(void* ptr, char* gestureId);
int QSensorGestureManager_RegisterSensorGestureRecognizer(void* ptr, void* recognizer);
void* QSensorGestureManager_QSensorGestureManager_SensorGestureRecognizer(char* id);
void QSensorGestureManager_DestroyQSensorGestureManager(void* ptr);
char* QSensorGesturePluginInterface_Name(void* ptr);
char* QSensorGesturePluginInterface_SupportedIds(void* ptr);
void QSensorGesturePluginInterface_DestroyQSensorGesturePluginInterface(void* ptr);
void QSensorGestureRecognizer_CreateBackend(void* ptr);
char* QSensorGestureRecognizer_GestureSignals(void* ptr);
char* QSensorGestureRecognizer_Id(void* ptr);
int QSensorGestureRecognizer_IsActive(void* ptr);
void QSensorGestureRecognizer_StartBackend(void* ptr);
void QSensorGestureRecognizer_StopBackend(void* ptr);
void QSensorGestureRecognizer_DestroyQSensorGestureRecognizer(void* ptr);
void* QSensorManager_QSensorManager_CreateBackend(void* sensor);
int QSensorManager_QSensorManager_IsBackendRegistered(void* ty, void* identifier);
void QSensorManager_QSensorManager_RegisterBackend(void* ty, void* identifier, void* factory);
void QSensorManager_QSensorManager_SetDefaultBackend(void* ty, void* identifier);
void QSensorManager_QSensorManager_UnregisterBackend(void* ty, void* identifier);
void QSensorPluginInterface_RegisterSensors(void* ptr);
void* QSensorReading_Value(void* ptr, int index);
int QSensorReading_ValueCount(void* ptr);
int QTapFilter_Filter(void* ptr, void* reading);
int QTapReading_IsDoubleTap(void* ptr);
int QTapReading_TapDirection(void* ptr);
void QTapReading_SetDoubleTap(void* ptr, int doubleTap);
void QTapReading_SetTapDirection(void* ptr, int tapDirection);
void* QTapSensor_Reading(void* ptr);
int QTapSensor_ReturnDoubleTapEvents(void* ptr);
void QTapSensor_SetReturnDoubleTapEvents(void* ptr, int returnDoubleTapEvents);
void* QTapSensor_NewQTapSensor(void* parent);
void QTapSensor_ConnectReturnDoubleTapEventsChanged(void* ptr);
void QTapSensor_DisconnectReturnDoubleTapEventsChanged(void* ptr);
void QTapSensor_DestroyQTapSensor(void* ptr);
int QTiltFilter_Filter(void* ptr, void* reading);
double QTiltReading_XRotation(void* ptr);
double QTiltReading_YRotation(void* ptr);
void QTiltReading_SetXRotation(void* ptr, double x);
void QTiltReading_SetYRotation(void* ptr, double y);
void* QTiltSensor_NewQTiltSensor(void* parent);
void* QTiltSensor_Reading(void* ptr);
void QTiltSensor_DestroyQTiltSensor(void* ptr);
void QTiltSensor_Calibrate(void* ptr);

#ifdef __cplusplus
}
#endif