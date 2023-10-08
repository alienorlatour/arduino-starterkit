# Love-o-meter

## The circuit

### Data sheets

- [TMP36](https://www.analog.com/media/en/technical-documentation/data-sheets/TMP35_36_37.pdf)

## Changeblog

This is pretty straightforward. What slowed me down was

1. the fact that pin A0 is called `AMC0`
2. finding out how to read the temperature from there

Apparently there is a bug in AVR that prevents me from using the standard output as debugger.
