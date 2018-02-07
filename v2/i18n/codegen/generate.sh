#!/bin/sh
go build && ./codegen -cout ../pluralspecs_gen.go -tout ../pluralspecs_gen_test.go && \
    gofmt -w=true ../pluralspecs_gen.go && \
    gofmt -w=true ../pluralspecs_gen_test.go && \
    rm codegen
