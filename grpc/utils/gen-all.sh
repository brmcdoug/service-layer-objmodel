#!/bin/bash
./gen-c-headers.sh
cd ../python
gen-bindings.sh
cd ../cpp
gen-bindings.sh
