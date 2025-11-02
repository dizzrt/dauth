#!/bin/bash

cd api && buf generate && cd ..
cd cmd && wire gen && cd ..
