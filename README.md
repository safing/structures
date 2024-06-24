# Go Structures

A small collection of useful data structures and helpers.

## container

A []byte slice on steriods that helps to reduce reallocations.

## dsd

DSD stands for dynamically structured data. It has an identifier for the format used, so file and wire encoding can be simply switched.
This makes it easier / more efficient to store different data types in a k/v data storage.

## varint

This is just a convenience wrapper around `encoding/binary`, because we use varints a lot.
