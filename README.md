# parkingLot
Solution for basic parking lot problem in Golang.

**Author: Amogh Mishra**

## Package

Package list in this application:

#### parking
```parking``` is main package, for initialize a parking lot. In each parking lot there is information of slot capacity
 and a set of slots that will be occupied by the car.


#### slot
```slot``` is a cleared area that is intended for parking car, with identity serial index number.


#### car
```car``` is basic object.  ```car```fill a ```slot``` and reduce parking lot slot capacity.

#### commands
```commands``` is a module that contains a set of commands. which serves to control the *lifecycle* of the
 above 3 objects.

## Setup application

Run bash ```setup``` inside ```bin``` directory
```sh
$ cd bin
$ ./setup
```
That command has automatically run unit testing, unit testing commands contained in the file  ```bin/run_functional_tests```

This application fully controlled by command. Run bash ```parking_lot``` inside bin directory with 2 options:

* The inputs commands are expected and taken from the file specified
```sh
$ ./parking_lot [input_filepath]
```
* Or start the program in interactive mode.
```sh
$ ./parking_lot
```
**Command list**

* ```>> create_parking_lot [capacity]```
Initialization of parking lot with parameters of slot capacity. This command must be run first to initialize the parking lot.

* ```>> park [car_registration_number] [car_color]```
The parking car in available slot with identity of registration number and color.


* ```>> leave [slot_number]```
Car leaves the parking lot so that the slot can be occupied by the another car will park.

* ```>> status```
For print parking area status in table format.
```Slot No.    Registration No     Colour```

* ```>> registration_numbers_for_cars_with_colour [car_color]```
For print all registration car with specific color who was parking.

* ```>> slot_numbers_for_cars_with_colour [car_color]```
For print all slot number with specific car color who was parking.

* ```>> slot_number_for_registration_number [car_registration_number]```
For print slot number with specific car registration color who was parking.
