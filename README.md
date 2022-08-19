## WhatsApp Status Trimmer (WAST) - Go

Simple Go Program to trim/split a video into 30s videos for WhatsApp Status. 
This project is golang port of [whatsapp-status-trimmer](https://github.com/tfkhdyt/whatsapp-status-trimmer)

## Built With

- [Go](https://go.dev/)
- [ffmpeg-go](https://github.com/u2takey/ffmpeg-go)

## Prerequisites

- [Go](https://go.dev/)
- [FFmpeg](https://ffmpeg.org/)

## Installation

```bash
git clone https://github.com/tfkhdyt/wast-go.git
cd wast-go
go install .
```

## Build 

```bash
cd wast-go
go build .
```

## Usage

```bash
wast-go <start timestamp> <end timestamp> <input file>
wast-go <options>
```

## Example

```bash
wast-go 00:00:23 00:01:40 weightless.mp4
```

The output would be 3 videos:

1. 00:00:23 - 00:00:53
2. 00:00:53 - 00:01:23
3. 00:01:23 - 00:01:40

## Options

```bash
-h, --help      Show help menu
-v, --version   Show app version
```

## Differences from the Node.js version

- Rewritten in Go
- Process is running asynchronously (thanks to goroutine)
- Fix file/folder name with space(s) issue
- Less bloated (bye bye node_modules)

## Contact

<p align=center>
  <a href="https://facebook.com/tfkhdyt142"><img height="30" src="https://upload.wikimedia.org/wikipedia/commons/5/51/Facebook_f_logo_%282019%29.svg"></a>&nbsp;
  <a href="https://twitter.com/tfkhdyt142"><img height="28" src="https://upload.wikimedia.org/wikipedia/commons/4/4f/Twitter-logo.svg"></a>&nbsp;
  <a href="https://instagram.com/_tfkhdyt_"><img height="30" src="https://upload.wikimedia.org/wikipedia/commons/e/e7/Instagram_logo_2016.svg"></a>&nbsp;
  <a href="https://youtube.com/tfkhdyt"><img height="30" src="https://upload.wikimedia.org/wikipedia/commons/a/a0/YouTube_social_red_circle_%282017%29.svg"></a>&nbsp;
  <a href="https://t.me/tfkhdyt"><img height="30" src="https://upload.wikimedia.org/wikipedia/commons/8/83/Telegram_2019_Logo.svg"></a>&nbsp;
  <a href="https://www.linkedin.com/mwlite/in/taufik-hidayat-6793aa200"><img height="30" src="https://upload.wikimedia.org/wikipedia/commons/8/81/LinkedIn_icon.svg"></a>
  <a href="https://pddikti.kemdikbud.go.id/data_mahasiswa/QUUyNzdEMjktNDk0Ri00RTlDLUE4NzgtNkUwRDBDRjIxOUNB"><img height="30" src="https://i.postimg.cc/YSB2c3DG/1619598282440.png"></a>
  <a href="https://tfkhdyt.my.id/"><img height="31" src="https://www.svgrepo.com/show/295345/internet.svg"></a>
</p>
