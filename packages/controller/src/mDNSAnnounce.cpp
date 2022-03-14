#include <ESP8266mDNS.h>

class mDNSAnnounce {
    private:
        String name;

    public:
        mDNSAnnounce(String name) {
            this->name = name;
        }

        void setup() {
            if (MDNS.begin(this->name)) {
                Serial.println("mDNS started");
                Serial.print("I am: ");
                Serial.println(this->name);

                // announce a bogus service so the server can find us
                MDNS.addService("MILTON", "tcp", 2354);
            }
            else {
                Serial.println("mDNS announce failed!");
            }
        }

        void update() {
            MDNS.update();
        }
};

