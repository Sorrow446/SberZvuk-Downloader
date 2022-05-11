# Zvuk-Downloader
[Zvuk (Звук)](https://dereferer.me/?https://zvuk.com/) downloader written in Go.
![](https://i.imgur.com/N706D0w.png)
[Windows, Linux, macOS and Android binaries](https://github.com/Sorrow446/SberZvuk-Downloader/releases)

# Setup
Input credentials into config file.
Configure any other options if needed.
|Option|Info|
| --- | --- |
|email|Email address.
|password|Password.
|format|Download format. 1 = 128 Kbps MP3, 2 = 320 Kbps MP3, 3 = 16/24-bit FLAC.
|outPath|Where to download to. Path will be made if it doesn't already exist.
|albumTemplate|Album folder naming template. Vars: album, albumArtist, year.
|trackTemplate|Track filename naming template. Vars: album, albumArtist, artist, genre, title, track, trackPad, trackTotal, year.
|maxCover|true = max cover size, false = 600x600.
|lyrics|Get lyrics if available.
|speedLimit|Download speed limit in megabytes. Example: 0.5 = 500 kB/s, 1 = 1 MB/s, 1.5 = 1.5 MB/s. -1 = unlimited.
|keepCover|true = don't delete covers from album folders.
|token|Use token to auth instead of credentials. Will only be used if not empty. Get from https://zvuk.com/api/v2/tiny/profile.

# Usage
Args take priority over the same config file options.

Download two albums with config file format:   
`zvuk_dl_x64.exe https://zvuk.com/release/14607525 https://zvuk.com/release/14024820`

Download a single album and from two text files in format 2 with lyrics:   
`zvuk_dl_x64.exe https://zvuk.com/release/14607525 G:\1.txt G:\2.txt -f 2 -l`

```
 _____         _      ____                _           _
|__   |_ _ _ _| |_   |    \ ___ _ _ _ ___| |___ ___ _| |___ ___
|   __| | | | | '_|  |  |  | . | | | |   | | . | .'| . | -_|  _|
|_____|\_/|___|_,_|  |____/|___|_____|_|_|_|___|__,|___|___|_|

Usage: zvuk_dl_x64.exe [--format FORMAT] [--outpath OUTPATH] [--maxcover] [--lyrics] [--albumtemplate ALBUMTEMPLATE] [--tracktemplate TRACKTEMPLATE] [--speedlimit SPEEDLIMIT] URLS [URLS ...]

Positional arguments:
  URLS

Options:
  --format FORMAT, -f FORMAT
                         Download format. 1 = 128 Kbps MP3, 2 = 320 Kbps MP3, 3 = 16/44 FLAC. [default: -1]
  --outpath OUTPATH, -o OUTPATH
                         Where to download to. Path will be made if it doesn't already exist.
  --maxcover, -m         true = max cover size, false = 600x600.
  --lyrics, -l           Get lyrics if available.
  --albumtemplate ALBUMTEMPLATE, -a ALBUMTEMPLATE
                         Album folder naming template. Vars: album, albumArtist, year.
  --tracktemplate TRACKTEMPLATE, -t TRACKTEMPLATE
                         Track filename naming template. Vars: album, albumArtist, artist, genre, title, track, trackPad, trackTotal, year.
  --speedlimit SPEEDLIMIT, -L SPEEDLIMIT
                         Download speed limit in megabytes. Example: 0.5 = 500 kB/s, 1 = 1 MB/s, 1.5 = 1.5 MB/s. -1 = unlimited. [default: -1]
  --help, -h             display this help and exit
```
 
# Disclaimer
- I will not be responsible for how you use Zvuk Downloader.    
- Zvuk brand and name is the registered trademark of its respective owner.    
- Zvuk Downloader has no partnership, sponsorship or endorsement with Zvuk.
