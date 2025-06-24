
# 🎥 Video Playback Event Simulation

This project simulates user interaction with a video player and emits analytics-style events for testing or prototyping real-world video tracking systems.

## 📋 Overview

The simulator produces an event stream mimicking how a video player behaves, including user actions (like pause, seek), system events (like buffering), and playback metrics (like bitrate and progress).

Each event includes a timestamp and contextually relevant data to allow for downstream metric calculation such as playing time, buffering time, pause durations, errors, etc.



## 📦 Initial Metadata Payload

Sent once per session to identify the user, content, and environment before any playback-related events occur.

| **Event**    | **Value Format**         | **Example Value**                                   | **Notes**                                            |
|--------------|--------------------------|-----------------------------------------------------|------------------------------------------------------|
| `sdk_ver`    | `string`                 | `"1.0.2"`                                           | SDK version emitting the analytics                   |
| `geoip`      | `string (IP address)`    | `"112.215.238.152"`                                 | Derived from client IP or geo-resolver               |
| `uadev`      | `string`                 | `"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Firefox/89.0 Safari/537.36"`                      | Platform, OS, browser, environment, device type      |
| `duuid`      | `UUID string`            | `"c18db03d-d943-412a-a025-462931e944ac"`            | Device/session-level identifier                      |
| `user_id`    | `string`                 | `"user@mail.com"`                                   | Optional user identifier (email, ID, anon token)     |
| `video_id`   | `string` / `int`         | `"95"`                                              | Unique video content identifier                      |
| `video_name` | `string`                 | `"Pacific Rim"`                                     | Human-readable title of the video                    |
| `tags`       | `comma-separated string` | `"Action,Adventure,Sci-Fi,International"`           | Genre or categorization labels                       |


## 🔄 Enriched Event List & Format

Each playback session emits events in response to user or system actions.

### `duration`
- **Value:** `int` (seconds)
- **Example:** `120`
- **Triggered by:** SDK/system on video load
- **Notes:** Sets total video duration for session
- **Required:** ✅ Yes

### `bitrate`
- **Value:** `video_bitrate, audio_bitrate` in Kbps
- **Example:** `3600,256`
- **Triggered by:** load, playing, seek - network change
- **Notes:** Can be emitted multiple times
- **Required:** ✅ Yes

### `load`
- **Value:** `int` (UNIX timestamp)
- **Example:** `1750482971043`
- **Triggered by:** Video player initialization
- **Notes:** First lifecycle event
- **Required:** ✅ Yes

### `play`
- **Value:** `int` (UNIX timestamp)
- **Example:** `1750482971043`
- **Triggered by:** User clicks play
- **Notes:** May occur again after pause
- **Required:** ✅ Yes

### `playing`
- **Value:** `timestamp, position_in_seconds`
- **Example:** `1750482976066, 5.00`
- **Triggered by:** Continuous playback, every 5 seconds
- **Notes:** Resets after pause/seek
- **Required:** ✅ Yes

### `buffer`
- **Value:** `int` (UNIX timestamp)
- **Example:** `1750482976066`
- **Triggered by:** Network stall
- **Notes:**
  Never emit repeated playing events without at least one buffer between.
  Between two playing events, buffer can appear 1–3 times randomly.
  No two buffer events can appear back-to-back without a playing after them.
- **Required:** ⚠️ Conditional

### `pause`
- **Value:** `int` (UNIX timestamp)
- **Example:** `1750482986082`
- **Triggered by:** User pauses playback
- **Notes:** Followed by new play event to resume
- **Required:** ⚠️ Conditional

### `seek`
- **Value:** `timestamp, target_seconds`
- **Example:** `1750482986082, 25.02`
- **Triggered by:** User scrub forward/backward
- **Notes:**
   Breaks playing stream
   May trigger a new bitrate (simulate quality adjustment)
- **Required:** ⚠️ Conditional

### `complete`
- **Value:** `int` (UNIX timestamp)
- **Example:** `1750483106325`
- **Triggered by:** Video reaches end
- **Required:** ⚠️ Conditional


### `unload`
- **Value:** `int` (UNIX timestamp)
- **Example:** `1750483106325`
- **Triggered by:** Player is closed, browser tab exit
- **Notes:** May occur before complete
- **Required:** ✅ Yes

### `error`
- **Value:** `string`
- **Example:** `"video error"`
- **Triggered by:** Playback failure
- **Required:** ⚠️ Optional

---


## 🧾 Output Format

The simulator produces **two types of outputs**:  
1. JSON payloads for REST API submission  
2. Terminal logs

### 1. 📤 JSON Payload

Each batch of events will be formatted as a JSON array:

```json
[
  {
    "tsclient": 1750482976066,
    "sessid": "RFAVKyVwVqLK8WLoqIFFIKSBSIsERvj9",
    "event": "playing",
    "value": "1750482976066,5.00"
  }
]
```

#### 🧾 Field Descriptions

| Field      | Type     | Description                                                                 |
|------------|----------|-----------------------------------------------------------------------------|
| `tsclient` | `int`    | UNIX timestamp (ms) generated by client                                     |
| `sessid`   | `string` | Unique session ID for the simulation run                                    |
| `event`    | `string` | The type of video event (`play`, `pause`, `buffer`, etc.)                   |
| `value`    | `string` | A string payload, often a timestamp or a time range, depending on the event |


### 2. 📺 Terminal Log Output

Each event is printed to terminal in a structured log format:

```
YYYY/MM/DD HH:MM:SS data[index]:: tsclient: <ts>, sessid: <id>, event: <name>, value: <val>
```

#### 🧾 Log Field Mapping

| Field        | Description                                        |
|--------------|----------------------------------------------------|
| `data[N]`    | Index of the conccurrent event in that batch       |
| `tsclient`   | Timestamp in ms                                    |
| `sessid`     | Unique session identifier                          |
| `event`      | Event type name                                    |
| `value`      | Event-specific payload (same as JSON `value`)      |


### 🧵 Concurrency Notes

The simulator is designed to:
- Emit events for **multiple concurrent sessions**
- Assign each session a unique `sessid`
- Maintain **independent timelines** for each playback stream
- Log all concurrent events **chronologically** in both JSON and terminal output

---

# How to Run the Video Playback Simulator

## Basic Usage

Run the simulator with a specific number of concurrent sessions:

```bash
make run concurrent=10
```

## With REST API Integration

To also send events to a REST API endpoint:

```bash
make run concurrent=10 api_url="https://your.api/endpoint" api_key="your-api-key"
```

### Notes

- If `api_url` is provided but `api_key` is missing, the simulator logs an error and skips sending events.
- If `api_key` is provided but `api_url` is missing, it behaves like a local simulation with no network calls.
 
---

## 📦 Use Cases

- Simulate frontend video analytics
- Test backend session reconstruction logic
- Benchmark playback behavior under different conditions
- Generate mock data for dashboards
