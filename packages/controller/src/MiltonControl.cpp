#include <PubSubClient.h>
#include "./mDNSAnnounce.cpp"
#include "./MQTT.cpp"
#include "./config.h"

class MiltonControl {
    private:
        mDNSAnnounce *mdns;
        MQTT *mqtt;

    public:
        MiltonControl() {
            this->mdns = new mDNSAnnounce(MDNS_NAME);

            // connect to the MQTT server and use the provided MDNS_NAME as a
            // unique identifier
            this->mqtt = new MQTT(MQTT_SERVER, MQTT_PORT, MQTT_USER, MQTT_PASS, MDNS_NAME);

            Serial.println("Adding callback");
            this->mqtt->setConnectedCallback([&](PubSubClient &mqtt) {
                mqtt.publish("esp/test", "HELLO FROM MILTON");
                mqtt.subscribe("esp/test");
            });

            this->mqtt->setMessageCallback([&] (
                char* topic,
                byte* payload,
                unsigned int length
            ) {
                this->handleMQTTMessage(topic, payload, length);
            });
        }

        void setup() {
            this->mdns->setup();
            this->mqtt->setup();
        }

        void handle() {
            this->mdns->update();
            this->mqtt->handle();
        }

        void handleMQTTMessage(char* topic, byte* payload, long length) {
            debug("Message arrived in topic: %s", topic);

            // if (topic == 'commands') {
            //     if (payload == 'water_plant') {
            //
            //     }
            //     else if (payload == 'move_to_pipe') {
            //
            //     }
            // }

            debug("Message: %s", payload);
            // for (int i = 0; i < length; i++) {
            //     debug((char)payload[i]);
            // }
            debug("-----------------------");
        }

        // void handleMQTTMessage(char* topic, byte* payload, long length) {
        //     Serial.print("Message arrived in topic: ");
        //     Serial.println(topic);
        //
        //     Serial.print("Message:");
        //     for (int i = 0; i < length; i++) {
        //         Serial.print((char)payload[i]);
        //     }
        //
        //     Serial.println();
        //     Serial.println("-----------------------");
        // }
};

