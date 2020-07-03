#include <Arduino.h>
#include <ESP8266WiFi.h>
#include "./config.h"


int ledState = LOW;

unsigned long previousMillis = 0;
const long interval = 1000;

void setup() {
    Serial.begin(115200);
    Serial.println("Booting");

    WiFi.mode(WIFI_STA);
    WiFi.begin(WIFI_SSID, WIFI_PASSPHRASE);

    while (WiFi.waitForConnectResult() != WL_CONNECTED) {
        Serial.println("Connection Failed! Rebooting...");
        delay(5000);
        ESP.restart();
    }

    pinMode(LED_BUILTIN, OUTPUT);

    Serial.println("Ready");
    Serial.print("IP address: ");
    Serial.println(WiFi.localIP());
}

void loop() {
    unsigned long currentMillis = millis();
    if(currentMillis - previousMillis >= interval) {
        previousMillis = currentMillis;
        if (ledState == LOW)
            ledState = HIGH;  // Note that this switches the LED *off*
        else
            ledState = LOW;   // Note that this switches the LED *on*
        digitalWrite(LED_BUILTIN, ledState);
    }
}

