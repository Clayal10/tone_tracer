# Tone Tracer

A guitar tone visualizer built with an arduino.

## The plan

<img src="./docs/flow.drawio.svg" width="500" height="400">

## Tech Stack

### Firmware

The arduino requires use of its C++ wrapper and IDE, so that will be used. The arduino will be responsible for:

- Capturing the analog input.
    - Guitar notes can range from ~80hz to 1khz, so a sample rate of ~15,000 should be adequate.
- Encoding it into some bare bones format.
- Sending it to a computer for real data manipulation.

### Software

Go will be the primary language. For visualization, [go-gl](https://github.com/go-gl/gl) will probably be the best solution.
