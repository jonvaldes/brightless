# brightless
Simple Linux tool to modify monitor brightness below what hardware is capable of (in the style of f.lux)

To compile this tool you'll need to have Go installed (https://golang.org)

To build, download source code and execute:

    go build

Instructions:

    Usage: brightless [DELTA]
    Modify the monitor brightness, adding DELTA to it.
    This program allows brightness to be between 0.10 and 1.00.
    
    Examples
    - To increase brightness by 1/10th
      brightless 0.1
      
    - To dim monitor by 1/10th
      brightless -0.1
      
    - To set full brightness
      brightless 1
      
    - To set full dimming
      brightless -1

    
