#include <ESP8266WiFi.h>
#include <PubSubClient.h>
#include "./debug.h"

class MQTT {
    private:
        PubSubClient mqtt;
        WiFiClient espClient;

        uint32_t lastCheck = 0;
        uint32_t currentMillis = 0;
        uint32_t checkInterval = 2000;
        String clientID;
        String user;
        String pass;

        // stores the callback that's executed when a successful connection to
        // the server has been made
        void (*connectedCallback)(PubSubClient &mqtt);

    public:
        MQTT(
            IPAddress server,
            uint16_t port,
            String user,
            String pass,
            String clientID
        ) : mqtt{ espClient } {
            this->clientID = clientID;
            this->user = user;
            this->pass = pass;

            mqtt.setServer(server, port);
        }

        void setup() {
            // noop
        }

        void handle() {
            if (!mqtt.connected()) {
                reconnect();
            }

            mqtt.loop();
        }

        void reconnect() {
            currentMillis = millis();

            // stop if it's too early to attempt to reconnect again
            if ((currentMillis - lastCheck) < checkInterval) {
                return;
            }

            debug("Attempting MQTT connection... ");

            if (mqtt.connect(this->clientID.c_str(), this->user.c_str(), pass.c_str())) {
                debug("connected");

                if (connectedCallback) {
                    connectedCallback(mqtt);
                }
            }
            else {
                debug("\nError: MQTT connection failed with state: %s", mqtt.state());
                debug("Attempting again in %n ms", checkInterval);
            }

            lastCheck = currentMillis;
        }

        void setIntervalCheck(int interval) {
            checkInterval = interval;
        }

        template<typename Functor>
        void setMessageCallback(Functor cb) {
            mqtt.setCallback(cb);
        }

        template<typename Functor>
        void setConnectedCallback(Functor cb) {
            connectedCallback = cb;
        }
};

