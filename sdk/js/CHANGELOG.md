# Change Log

All notable changes to this project will be documented in this file in
accordance with
[![keepachangelog 1.0.0](https://img.shields.io/badge/keepachangelog-1.0.0-brightgreen.svg)](http://keepachangelog.com/en/1.0.0/)

## \[Unreleased]

## 0.0.0-alpha.20

### Removed

- `node/api/client#pkg_content_list` & `node/api/client#pkg_content_get` methods; use `node/api/client#data_get`

## 0.0.0-alpha.19

### Added

- reconnect on websocket close

## 0.0.0-alpha.18

### Added

- ack to event_stream_get call to enable websocket back pressure

