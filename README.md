# Milton
An M to N (**M**il**toN**) irrigation system.  
Where M is multiple pump controllers and N is multiple plants per pump.

## Status
Very much in alpha stage.

## Rationale
My amazing wife likes to create a small jungle of colorful plants on our balcony during the spring and summer months. This is great and all but all those plants require constant watering (sometimes two times a day), which is tedious and finding someone to care for them when we're on vacations is also not easy.

Searching for existing irrigation solutions online didn't yield any results based on our criteria:
1. The system should be able to water multiple plants
2. The system should be able to supply each plant with a different amount of
   water
3. The system should be remotely configurable

Most of the smart irrigation systems needed one water pump per flower pot which is expensive and not very practical, especially if you have around 30 flower pots on your balcony. Other systems ran one tube across all the plants and supplied each flower pot with the same amount of water. There just wasn't anything like what we needed.

Thus the idea for Milton was born.

## The setup

The idea is that Milton is divided into two components:
1. The server - where the user configures the watering schedule for each individual flower pot
2. The water pump controller - which receives commands from the server, like which flower pot to irrigate and how much water to use

The server is a pretty straightforward CRUD application which sends commands via MQTT to the pump controller based on a cron schedule. The interesting (and most difficult part) is how the pump controller, controlling only one pump, can dispense water to multiple flower pots independently.

Each flower pot gets one end of a small tube. The other end is connected to **a coupler** that chooses which flower pot currently receives water.

## The coupler
The coupler is the crux of the project. It's the hardest thing to get right and probably the reason why I couldn't find similar irrigation systems in the DIY spaces on the web.

The tubes from each flower pot will be arranged in a pattern and one tube from the water pump that also has a water meter connected to the controller will go from tube to tube and dispense the needed amount of water and then move to the next pot.

The easiest approach is to use a [corexy](https://corexy.com/) system with rods and bushings, driven by two stepper motors, similar to how some 3d printers work. However this approach is bulky and expensive and sort of an overkill for a balcony irrigation system.

The two ideas I'm currently exploring are a revolver-type coupler, where the tubes are organized in a circle, just like bullets in a revolver. The other is based on the design of the [plot clock](https://www.youtube.com/watch?v=iOLFP90DneY) project. The latter is my current approach, however due to a data loss incident I lost a lot of progress on the microcontroller code with the inverse kinematics formulas.
