# Digital VLF Recording and Analysis System

## Overview
At the moment this microservice serves as the data gateway and device interface for the system. It might be more prudent to split this into two separate services:
* A general "IoT Gateway" style service for running on edge devices (RPi). The gateway will then place all data on a bus, event style.
* A separate device or "Thing" style service for collecting the data and sending it to the gateway.