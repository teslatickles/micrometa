/*
Controlling LED Brightness Via PWM
LED lights up gradually,and then goes out gradually,repeatedly
********************************************************************
*/

const int ledPin = 9;

void setup() {
  pinMode(ledPin, OUTPUT);
}

void loop() {
  // increase the brightness of led
  for (int a = 0; a <= 255; a++) {
    analogWrite(ledPin, a);
    delay(8);
  }
  // decrease the brightness of led
  for (int a = 255; a >= 0; a--) {
    analogWrite(ledPin, a);
    delay(8);
  }

  delay(800);
}
