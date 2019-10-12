**Controlling Raspberry Pi (3) with Golang**

- This is a simple example using Go to control an LED connected to Raspberry Pi's GPIO.

 *This example is intended to be built (using build command below) on a Desktop/Laptop,
then transferred wirelessly to Raspberry Pi using scp (secured copy) in the terminal.*  

**What it does, exactly**

- This simple Go application allows one to press the UP arrow key or DOWN arrow key on keyboard
to control turning an LED On or Off, respectively.

**Step-by-Step**

   - The GPIO pin that is designated in this example is pin **18**, which is physical pin **12**.
    
*Note: Of course, you can change this, just follow an online guide to GPIO for Raspi.*

   **DON"T FORGET YOU NEED A RESISTOR TO LIMIT THE CURRENT GOING THROUGH THE RESISTOR!**

   - Clone this repo: `git clone [location of this repo]` in terminal/command line.

   - Build an executable binary file using this command:

    ```env GOOS=linux GOARCH=arm GOARM=5 build```

   - this means using linux operating system, on raspi 3, which means ARM architecture and 5 designates for Raspi **3**

   - scp the binary file to target Raspberry Pi 3 (make sure Raspi is booted up and has internet connection):

  ``scp rpio pi@[pi's ip address here]:/home/pi/[name you want executable to be on pi]``

    You will be prompted to enter the password for your Raspi.

    *Note: you can find your pi's ip by typing `ifconfig` in terminal*

   - Now, we need to run the executable on the Raspi:

    *You can either ssh into the pi from another computer or boot up the Raspberry Pi, then double-click the executable file found at /home/pi/ and select
    `Run from terminal`.*

   - Now, try pressing Up arrow key and Down arrow key to see the LED turn On and Off, respectively!

    Thanks for checking this out,

    - Hunter Hartline (teslatickles)