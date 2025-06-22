
# 🎥 Video Playback Event Simulation

This project simulates user interaction with a video player and emits analytics-style events for testing or prototyping real-world video tracking systems.

---

## 📋 Overview

The simulator produces an event stream mimicking how a video player behaves, including user actions (like pause, seek), system events (like buffering), and playback metrics (like bitrate and progress).

Each event includes a timestamp and contextually relevant data to allow for downstream metric calculation such as playing time, buffering time, pause durations, errors, etc.

---

## 📦 Initial Metadata Payload

Sent once per session to identify the user, content, and environment before any playback-related events occur.

| **Event**    | **Value Format**         | **Example Value**                                   | **Notes**                                            |
|--------------|--------------------------|-----------------------------------------------------|------------------------------------------------------|
| `sdk_ver`    | `string`                 | `"1.0.2"`                                           | SDK version emitting the analytics                   |
| `geoip`      | `string` (IP address)    | `"112.215.238.152"`                                 | Derived from client IP or geo-resolver               |
| `uadev`      | `string`                 | `"iOS,16.2,Edge,Browser,iPad"`                      | Platform, OS, browser, environment, device type      |
| `duuid`      | `UUID string`            | `"c18db03d-d943-412a-a025-462931e944ac"`            | Device/session-level identifier                      |
| `user_id`    | `string`                 | `"user@mail.com"`                                   | Optional user identifier (email, ID, anon token)     |
| `video_id`   | `string` / `int`         | `"95"`                                              | Unique video content identifier                      |
| `video_name` | `string`                 | `"Pacific Rim"`                                     | Human-readable title of the video                    |
| `tags`       | `comma-separated string` | `"Action,Adventure,Sci-Fi,International"`           | Genre or categorization labels                       |

---

## 🔄 Enriched Event List & Format

Each playback session emits events in response to user or system actions.

### `duration`
- **Value:** `int` (seconds)
- **Example:** `120`
- **Triggered by:** SDK/system on video load
- **Notes:** Sets total duration for session
- **Required:** ✅ Yes

---

### `bitrate`
- **Value:** `[video_bitrate, audio_bitrate]` in bps
- **Example:** `[200000, 64000]`
- **Triggered by:** load, playing, network change
- **Notes:** Can be emitted multiple times
- **Required:** ✅ Yes

---

### `load`
- **Value:** `int` (UNIX timestamp)
- **Example:** `1750482971043`
- **Triggered by:** Video player initialization
- **Notes:** First lifecycle event
- **Required:** ✅ Yes

---

### `play`
- **Value:** `int` (UNIX timestamp)
- **Example:** `1750482971043`
- **Triggered by:** User clicks play
- **Notes:** May occur multiple times after pause
- **Required:** ✅ Yes

---

### `playing`
- **Value:** `[timestamp, position_in_seconds]`
- **Example:** `[1750482976066, 5.00]`
- **Triggered by:** Continuous playback, every 5 seconds
- **Notes:** Resets after pause/seek
- **Required:** ✅ Yes

---

### `buffer`
- **Value:** `int` (UNIX timestamp)
- **Example:** `1750482976066`
- **Triggered by:** Network stall
- **Notes:** May interrupt playing
- **Required:** ⚠️ Conditional

---

### `pause`
- **Value:** `int` (UNIX timestamp)
- **Example:** `1750482986082`
- **Triggered by:** User pauses playback
- **Notes:** Followed by new play event to resume
- **Required:** ⚠️ Conditional

---

### `seek`
- **Value:** `[timestamp, target_seconds]`
- **Example:** `[1750482986082, 25.02]`
- **Triggered by:** User scrub forward/backward
- **Notes:** Breaks playing stream
- **Required:** ⚠️ Conditional

---

### `complete`
- **Value:** `int` (UNIX timestamp)
- **Example:** `1750483106325`
- **Triggered by:** Video reaches end
- **Required:** ⚠️ Conditional

---

### `unload`
- **Value:** `int` (UNIX timestamp)
- **Example:** `1750483106325`
- **Triggered by:** Player is closed, browser tab exit
- **Notes:** May occur before complete
- **Required:** ✅ Yes

---

### `error`
- **Value:** `string`
- **Example:** `"video error"`
- **Triggered by:** Playback failure
- **Required:** ⚠️ Optional

---

## 🧠 Notes

- All timestamps are **UNIX time in milliseconds**
- `playing` occurs every 5 seconds during uninterrupted playback
- `pause`, `seek`, and `buffer` reset the `playing` cycle
- Events are **ordered chronologically**, starting from metadata → `load` → `play` → playback events → `unload` or `complete`

---

## 🧾 Output Format

The simulator produces **two types of outputs**:  
1. JSON payloads suitable for REST API submission  
2. Human-readable terminal logs

---

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

---

### 2. 📺 Terminal Log Output

Each event is printed to terminal in a structured log format:

```
2025/06/21 12:16:16 data[0]:: tsclient: 1750482976066, sessid: RFAVKyVwVqLK8WLoqIFFIKSBSIsERvj9, group: video_measure, label: playing, value: 1750482976066,5.00
```

#### 🧾 Log Field Mapping

| Field        | Description                                        |
|--------------|----------------------------------------------------|
| `data[N]`    | Index of the event in that batch                   |
| `tsclient`   | Timestamp in ms                                    |
| `sessid`     | Unique session identifier                          |
| `group`      | Fixed value: `video_measure`                       |
| `label`      | Event type name (same as `event`)                  |
| `value`      | Event-specific payload (same as JSON `value`)      |

---

### 🧵 Concurrency Notes

The simulator is designed to:
- Emit events for **multiple concurrent sessions**
- Assign each session a unique `sessid`
- Maintain **independent timelines** for each playback stream
- Log all concurrent events **chronologically** in both JSON and terminal output

---

## 📦 Use Cases

- Simulate frontend video analytics
- Test backend session reconstruction logic
- Benchmark playback behavior under different conditions
- Generate mock data for dashboards
