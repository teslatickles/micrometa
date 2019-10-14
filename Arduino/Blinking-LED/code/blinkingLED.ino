//Script for making an LED blink
//We can use the built-in LED in pin 13 (or plug one in).

int ledPin = 13; 
int blink = 1000
//Here we setup the pin modes.
void setup()
{
    //LED will be used as output pin.
    pinMode(ledPin, OUTPUT);
}

//Here we setup the main program
void loop()
{
    //We send a HIGH voltage signal to turn the LED on
    digitalWrite(ledPin, HIGH);
    //We wait for "blink" milliseconds
    delay(blink);
    //We send a LOW voltage signal to turn the LED off
    digitalWrite(ledPin, LOW);
    //We wait for "blink" milliseconds
    delay(blink);
    //Repeat
}
