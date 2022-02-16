# SberZvuk-Downloader
[SberZvuk (СберЗвук)](https://dereferer.me/?https://sber-zvuk.com/) downloader written in Go.
![](https://i.imgur.com/wx90jGV.png)
[Windows, Linux, macOS and Android binaries](https://github.com/Sorrow446/SberZvuk-Downloader/releases)

# Setup
Input credentials into config file.
Configure any other options if needed.
|Option|Info|
| --- | --- |
|email|Email address.
|password|Password.
|format|Download format. 1 = 128 Kbps MP3, 2 = 320 Kbps MP3, 3 = 16/44 FLAC.
|outPath|Where to download to. Path will be made if it doesn't already exist.
|trackTemplate|Track filename naming template. Vars: album, albumArtist, artist, genre, title, track, trackPad, trackTotal, year.
|maxCover|true = max cover size, false = 600x600.
|lyrics|Get lyrics if available.

# Usage
Args take priority over the same config file options.

Download two albums with config file format:   
`sberzvuk_dl_x64.exe https://sber-zvuk.com/release/14607525 https://sber-zvuk.com/release/14024820`

Download a single album and from two text files in format 2 with lyrics:   
`sberzvuk_dl_x64.exe https://sber-zvuk.com/release/14607525 G:\1.txt G:\2.txt -f 2 -l`

```
 _____ _           _____         _      ____                _           _
|   __| |_ ___ ___|__   |_ _ _ _| |_   |    \ ___ _ _ _ ___| |___ ___ _| |___ ___
|__   | . | -_|  _|   __| | | | | '_|  |  |  | . | | | |   | | . | .'| . | -_|  _|
|_____|___|___|_| |_____|\_/|___|_,_|  |____/|___|_____|_|_|_|___|__,|___|___|_|

Usage: sberzvuk_dl_x64.exe [--format FORMAT] [--outpath OUTPATH] [--maxcover] [--lyrics] [--tracktemplate TRACKTEMPLATE] URLS [URLS ...]

Positional arguments:
  URLS

Options:
  --format FORMAT, -f FORMAT
                         Download format. 1 = 128 Kbps MP3, 2 = 320 Kbps MP3, 3 = 16/24-bit FLAC. [default: -1]
  --outpath OUTPATH, -o OUTPATH
                         Where to download to. Path will be made if it doesn't already exist.
  --maxcover, -m         true = max cover size, false = 600x600.
  --lyrics, -l           Get lyrics if available.
  --tracktemplate TRACKTEMPLATE, -f TRACKTEMPLATE
                         Track filename naming template. Vars: album, albumArtist, artist, genre, title, track, trackPad, trackTotal, year.
  --help, -h             display this help and exit
```
 
# Disclaimer
- I will not be responsible for how you use SberZvuk Downloader.    
- SberZvuk brand and name is the registered trademark of its respective owner.    
- SberZvuk Downloader has no partnership, sponsorship or endorsement with SberZvuk.
