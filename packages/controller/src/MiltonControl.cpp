#include "./mDNSAnnounce.cpp"
#include "./config.h"

class MiltonControl {
    private:
        mDNSAnnounce *mdns;

    public:
        MiltonControl() {
            this->mdns = new mDNSAnnounce(MDNS_NAME);
        }

        void setup() {
            this->mdns->setup();
        }

        void handle() {
            this->mdns->update();
        }
};

