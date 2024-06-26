#include <Arduino.h>
#include <ESP8266WiFi.h>
#include <ArduinoOTA.h>

#include "./debug.h"
#include "./config.h"

#include "./MiltonControl.cpp"

MiltonControl Milton;

int ledState = LOW;

unsigned long previousMillis = 0;
const long interval = 1000;

// sanity checks
#if OTA_ENABLE && !defined(OTA_PASSWORD)
    #error "Please uncomment and set `OTA_PASSWORD` in config.h"
#endif

void setup() {
    // use `Serial` for debugging `setup` and `RemoteDebug` for everything else
    Serial.begin(115200);
    Serial.println("\n\nBooting");

    WiFi.mode(WIFI_STA);
    WiFi.begin(WIFI_SSID, WIFI_PASSPHRASE);

    while (WiFi.waitForConnectResult() != WL_CONNECTED) {
        Serial.println("Connection Failed! Rebooting...");
        delay(5000);
        ESP.restart();
    }

    pinMode(LED_BUILTIN, OUTPUT);

#if OTA_ENABLE
    // Port defaults to 8266
    ArduinoOTA.setPort(OTA_PORT);

    // Hostname defaults to esp8266-[ChipID]
    ArduinoOTA.setHostname(OTA_HOSTNAME);
    ArduinoOTA.setPassword(OTA_PASSWORD);

    ArduinoOTA.onStart([]() {
        String type;
        if (ArduinoOTA.getCommand() == U_FLASH)
            type = "sketch";
        else // U_SPIFFS
            type = "filesystem";

        // NOTE: if updating SPIFFS this would be the place to unmount SPIFFS using SPIFFS.end()
        Serial.println("Start updating " + type);
    });
    ArduinoOTA.onEnd([]() {
        Serial.println("\nEnd");
    });
    ArduinoOTA.onProgress([](unsigned int progress, unsigned int total) {
        Serial.printf("Progress: %u%%\r", (progress / (total / 100)));
    });
    ArduinoOTA.onError([](ota_error_t error) {
        Serial.printf("Error[%u]: ", error);
        if (error == OTA_AUTH_ERROR) Serial.println("Auth Failed");
        else if (error == OTA_BEGIN_ERROR) Serial.println("Begin Failed");
        else if (error == OTA_CONNECT_ERROR) Serial.println("Connect Failed");
        else if (error == OTA_RECEIVE_ERROR) Serial.println("Receive Failed");
        else if (error == OTA_END_ERROR) Serial.println("End Failed");
    });
    ArduinoOTA.begin();
#endif

    Serial.println("Ready");
    Serial.print("IP address: ");
    Serial.println(WiFi.localIP());

    Milton.setup();
}

void loop() {
    #if OTA_ENABLE
        ArduinoOTA.handle();
    #endif

    unsigned long currentMillis = millis();
    if(currentMillis - previousMillis >= interval) {
        previousMillis = currentMillis;
        if (ledState == LOW)
            ledState = HIGH;  // Note that this switches the LED *off*
        else
            ledState = LOW;   // Note that this switches the LED *on*
        digitalWrite(LED_BUILTIN, ledState);

        debug("Toggling the LED");
    }

    Milton.handle();
}

