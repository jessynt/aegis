CREATE DATABASE aegis on cluster '{cluster}';

CREATE TABLE aegis.event_replica on cluster '{cluster}' (
    "ModelId" UInt64,
    "Date" Date DEFAULT toDate(Time),
    "Time" DateTime,
    "Year" UInt16,
    "Quarter" UInt8,
    "Month" UInt8,
    "DayOfMonth" UInt8,
    "DayOfWeek" UInt8,
    "DeviceId" String,
    "UserId" Int64,
    "IP" String,
    "Country" String,
    "Province" String,
    "City" String
) ENGINE = ReplicatedMergeTree('/clickhouse/tables/{cluster}-{shard}/event_replica', '{replica}')
PARTITION BY Month
ORDER BY (ModelId, Year, Date)
SETTINGS index_granularity = 8192;

CREATE TABLE aegis.event_dist ON CLUSTER '{cluster}' AS aegis.event_replica ENGINE = Distributed('{cluster}', aegis, event_replica, rand())