#!/bin/bash
set -e

clickhouse client -n <<-EOSQL
   CREATE DATABASE IF NOT EXISTS aegis;
   CREATE TABLE IF NOT EXISTS aegis.event_dist (
    "ModelId" UInt64,
    "Date" Date,
    "Time" DateTime,
    "Year" UInt16,
    "Quarter" UInt8,
    "Month" UInt8,
    "DayOfMonth" UInt8,
    "DayOfWeek" UInt8,
    "DeviceId" String,
    "UserId" UInt64,
    "IP" String,
    "ipLocation.Country" String,
    "ipLocation.Province" String,
    "ipLocation.City" String
) ENGINE = MergeTree(Date, (ModelId, Year, Date), 8192);
EOSQL
