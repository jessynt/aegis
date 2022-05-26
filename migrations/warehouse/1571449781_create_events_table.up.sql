CREATE TABLE events (
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
) ENGINE = MergeTree(Date, (ModelId, Year, Date), 8192);