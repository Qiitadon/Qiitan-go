#!/bin/sh

rm -f ./fibonacci.qtn

qiitan -compile ./fibonacci.qiitan || {
    echo >&2 "Failed to compile qiitan."
    exit 1
}

perf stat -r5 qiitan fibonacci.qtn
