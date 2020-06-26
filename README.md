# drone-upx

[![Build Status](https://cloud.drone.io/api/badges/cnbattle/drone-upx/status.svg)](https://cloud.drone.io/cnbattle/drone-upx)
[![LINK](https://img.shields.io/badge/link-Github-%23FF4D5B.svg)](https://github.com/cnbattle/drone-upx) 
[![Go Report Card](https://goreportcard.com/badge/github.com/cnbattle/drone-upx)](https://goreportcard.com/report/github.com/cnbattle/drone-upx)
[![Docker Pulls](https://img.shields.io/docker/pulls/cnbattle/drone-upx.svg)](https://img.shields.io/docker/pulls/cnbattle/drone-upx.svg)
[![GoDoc](https://godoc.org/github.com/cnbattle/drone-upx?status.svg)](https://godoc.org/github.com/cnbattle/drone-upx)
[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE)

## Use
```
  - name : upx
    image: cnbattle/drone-upx
    settings:
      level: 9 //default 5
      save_file: ./executable_upx_file
      original_file: ./executable_original_file
```