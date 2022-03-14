#include <Servo.h>

class ServoControl {
    Servo servo;

    bool running = false;
    int initialAngle;
    int currentPos;
    int targetPos;
    int increment;
    int updateInterval;
    unsigned long lastUpdate;
    unsigned long currentMillis;

    public:
    ServoControl() {
        increment = 1;
        initialAngle = 90;
    }

    void setUpdateInterval(int interval) {
        updateInterval = interval;
    }

    void setInitialAngle(int angle) {
        initialAngle = angle;
    }

    void attach(int pin) {
        servo.attach(pin);
        // TODO: Add delay?
        servo.write(initialAngle);
        currentPos = initialAngle;
    }

    void detach() {
        servo.detach();
    }

    void moveTo(int angle) {
        targetPos = angle;
        running = true;
        increment = abs(increment);

        if (currentPos > targetPos) {
            increment = -increment;
        }
    }

    bool isRunning() {
        return running;
    }

    void update() {
        if (!running) {
            return;
        }

        currentMillis = millis();

        if ((currentMillis - lastUpdate) > updateInterval) {
            servo.write(currentPos);
            lastUpdate = currentMillis;
            currentPos += increment;

            // Serial.printf("Current: %d, Target: %d\n", currentPos, targetPos);

            if (currentPos == targetPos) {
                // TODO: Fire a callback?
                running = false;
            }
        }
    }
};

