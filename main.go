package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/bogem/id3v2"
	"github.com/dustin/go-humanize"
	"github.com/go-flac/flacpicture"
	"github.com/go-flac/flacvorbis"
	"github.com/go-flac/go-flac"
)

const (
	megabyte            = 1000000
	apiBase             = "https://zvuk.com/"
	albumRegexString    = `^https://zvuk.com/release/(\d+)$`
	playlistRegexString = `^https://zvuk.com/playlist/(\d+)$`
	tokRegexString      = `^[\da-zA-Z]{32}$`
	userAgent           = "OpenPlay|4.10.2|Android|7.1.2|Asus ASUS_Z01QD"
	trackTemplate       = "{{.trackPad}}. {{.title}}"
	albumTemplate       = "{{.albumArtist}} - {{.album}}"
)

var (
	jar, _     = cookiejar.New(nil)
	client     = &http.Client{Jar: jar, Transport: &Transport{}}
	qualityMap = map[int]string{
		1: "mid",
		2: "high",
		3: "flac",
	}
)

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add(
		"User-Agent", userAgent,
	)
	req.Header.Add(
		"Referer", apiBase,
	)
	return http.DefaultTransport.RoundTrip(req)
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	var speed int64 = 0
	n := len(p)
	wc.Downloaded += int64(n)
	percentage := float64(wc.Downloaded) / float64(wc.Total) * float64(100)
	wc.Percentage = int(percentage)
	toDivideBy := time.Now().UnixMilli() - wc.StartTime
	if toDivideBy != 0 {
		speed = int64(wc.Downloaded) / toDivideBy * 1000
	}
	fmt.Printf("\r%d%% @ %s/s, %s/%s ", wc.Percentage, humanize.Bytes(uint64(speed)),
		humanize.Bytes(uint64(wc.Downloaded)), wc.TotalStr)
	return n, nil
}

func handleErr(errText string, err error, _panic bool) {
	errString := errText + "\n" + err.Error()
	if _panic {
		panic(errString)
	}
	fmt.Println(errString)
}

func wasRunFromSrc() bool {
	buildPath := filepath.Join(os.TempDir(), "go-build")
	return strings.HasPrefix(os.Args[0], buildPath)
}

func getScriptDir() (string, error) {
	var (
		ok    bool
		err   error
		fname string
	)
	runFromSrc := wasRunFromSrc()
	if runFromSrc {
		_, fname, _, ok = runtime.Caller(0)
		if !ok {
			return "", errors.New("Failed to get script filename.")
		}
	} else {
		fname, err = os.Executable()
		if err != nil {
			return "", err
		}
	}
	return filepath.Dir(fname), nil
}

func readTxtFile(path string) ([]string, error) {
	var lines []string
	f, err := os.OpenFile(path, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

func contains(lines []string, value string) bool {
	for _, line := range lines {
		if strings.EqualFold(line, value) {
			return true
		}
	}
	return false
}

func processUrls(urls []string) ([]string, error) {
	var (
		processed []string
		txtPaths  []string
	)
	for _, _url := range urls {
		if strings.HasSuffix(_url, ".txt") && !contains(txtPaths, _url) {
			txtLines, err := readTxtFile(_url)
			if err != nil {
				return nil, err
			}
			for _, txtLine := range txtLines {
				if !contains(processed, txtLine) {
					processed = append(processed, txtLine)
				}
			}
			txtPaths = append(txtPaths, _url)
		} else {
			if !contains(processed, _url) {
				processed = append(processed, _url)
			}
		}
	}
	return processed, nil
}

func parseCfg() (*Config, error) {
	cfg, err := readConfig()
	if err != nil {
		return nil, err
	}
	args := parseArgs()
	if args.SpeedLimit != -1 {
		cfg.SpeedLimit = args.SpeedLimit
	}
	if cfg.SpeedLimit != -1 && cfg.SpeedLimit <= 0 {
		return nil, errors.New("Invalid speed limit.")
	}
	cfg.ByteLimit = int64(megabyte * cfg.SpeedLimit)
	if cfg.SpeedLimit != -1 {
		fmt.Printf("Download speed limiting is active, limit: %s/s.\n",
			humanize.Bytes(uint64(cfg.ByteLimit)))
	}
	if args.Format != -1 {
		cfg.Format = args.Format
	}
	if !(cfg.Format >= 1 && cfg.Format <= 3) {
		return nil, errors.New("Format must be between 1 and 3.")
	}
	cfg.FormatStr = qualityMap[cfg.Format]
	if args.OutPath != "" {
		cfg.OutPath = args.OutPath
	}
	if args.MaxCover {
		cfg.MaxCover = args.MaxCover
	}
	if args.Lyrics {
		cfg.Lyrics = args.Lyrics
	}
	if args.TrackTemplate != "" {
		cfg.TrackTemplate = args.TrackTemplate
	}
	if cfg.OutPath == "" {
		cfg.OutPath = "Zvuk downloads"
	}
	if cfg.TrackTemplate == "" {
		cfg.TrackTemplate = "{{.trackPad}}. {{.title}}"
	}
	cfg.Urls, err = processUrls(args.Urls)
	if err != nil {
		errString := fmt.Sprintf("Failed to process URLs.\n%s", err)
		return nil, errors.New(errString)
	}
	return cfg, nil
}

func readConfig() (*Config, error) {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}
	var obj Config
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func parseArgs() *Args {
	var args Args
	arg.MustParse(&args)
	return &args
}

func makeDirs(path string) error {
	return os.MkdirAll(path, 0755)
}

func checkToken(token string) bool {
	regex := regexp.MustCompile(tokRegexString)
	return regex.MatchString(token)
}

func auth(email, pwd string) (string, error) {
	data := url.Values{}
	data.Set("email", email)
	data.Set("password", pwd)
	req, err := http.NewRequest(http.MethodPost, apiBase+"api/tiny/login/email", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	do, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer do.Body.Close()
	if do.StatusCode != http.StatusOK {
		return "", errors.New(do.Status)
	}
	var obj Auth
	err = json.NewDecoder(do.Body).Decode(&obj)
	if err != nil {
		return "", err
	}
	return obj.Result.Token, nil
}

func getToken(token, email, password string) (string, error) {
	var err error
	if token != "" {
		if !checkToken(token) {
			return "", errors.New("Invalid token.")
		}
	} else {
		token, err = auth(email, password)
		if err != nil {
			return "", err
		}
	}
	return token, nil
}

func getUserInfo(token string) (*UserInfo, error) {
	req, err := http.NewRequest(http.MethodGet, apiBase+"api/v2/tiny/profile", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-auth-token", token)
	do, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer do.Body.Close()
	if do.StatusCode != http.StatusOK {
		return nil, errors.New(do.Status)
	}
	var obj UserInfo
	err = json.NewDecoder(do.Body).Decode(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func checkUrl(url string) ItemType {
	matchAlbum := regexp.MustCompile(albumRegexString).FindStringSubmatch(url)
	if matchAlbum == nil {
		matchPlaylist := regexp.MustCompile(playlistRegexString).FindStringSubmatch(url)
		if matchPlaylist == nil {
			return ItemType{0, ""}
		} else {
			return ItemType{2, matchPlaylist[1]}
		}
	}
	return ItemType{1, matchAlbum[1]}
}

func getMeta(itemType ItemType, token string) (*Meta, error) {

	var api string
	switch itemType.TypeId {
	case 1:
		api = "api/tiny/releases"
	case 2:
		api = "api/tiny/playlists"
	case 3:
		//TODO single track
	}
	req, err := http.NewRequest(http.MethodGet, apiBase+api, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-auth-token", token)
	query := url.Values{}
	query.Set("ids", itemType.ItemId)
	query.Set("include", "track")
	req.URL.RawQuery = query.Encode()
	do, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer do.Body.Close()
	if do.StatusCode != http.StatusOK {
		return nil, errors.New(do.Status)
	}
	var obj Meta
	err = json.NewDecoder(do.Body).Decode(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func sanitize(filename string, isFolder bool) string {
	var regexStr string
	if isFolder {
		regexStr = `[:*?"><|]`
	} else {
		regexStr = `[\/:*?"><|]`
	}
	return regexp.MustCompile(regexStr).ReplaceAllString(filename, "_")
}

func parseTrackIds(trackIds []int) string {
	var parsed string
	for _, trackId := range trackIds {
		parsed += strconv.Itoa(trackId) + ","
	}
	return parsed[:len(parsed)-1]
}

func getStreamMeta(trackId, formatStr, token string) (*StreamMeta, error) {
	var do *http.Response
	req, err := http.NewRequest(http.MethodGet, apiBase+"api/tiny/track/stream", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-auth-token", token)
	query := url.Values{}
	query.Set("id", trackId)
	query.Set("quality", formatStr)
	req.URL.RawQuery = query.Encode()
	for i := 0; i < 5; i++ {
		do, err = client.Do(req)
		if err != nil {
			return nil, err
		}
		if do.StatusCode == http.StatusTeapot && i != 4 {
			do.Body.Close()
			fmt.Printf("Got a HTTP 418, %d attempt(s) remaining.\n", 4-i)
			time.Sleep(time.Second * 3)
			continue
		}
		if do.StatusCode != http.StatusOK {
			do.Body.Close()
			return nil, errors.New(do.Status)
		}
		break
	}
	var obj StreamMeta
	err = json.NewDecoder(do.Body).Decode(&obj)
	do.Body.Close()
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func queryQualities(track *Track, cfg *Config) {
	avail := true
	format := cfg.Format
	highestQual := track.HighestQuality
	if format == 3 && highestQual != "flac" {
		avail = false
	} else if format == 2 && highestQual == "med" {
		avail = false
	}
	if !avail {
		fmt.Println("Track unavailable in your chosen quality.")
		cfg.FormatStr = highestQual
	}
}

func queryRetQuality(streamUrl string) *Quality {
	qualityMap := map[string]Quality{
		"/stream?":   {"128 Kbps MP3", ".mp3", false},
		"/streamhq?": {"320 Kbps MP3", ".mp3", false},
		"/streamfl?": {"FLAC", ".flac", true},
	}
	for k, v := range qualityMap {
		if strings.Contains(streamUrl, k) {
			return &v
		}
	}
	return nil
}

func fileExists(path string) (bool, error) {
	f, err := os.Stat(path)
	if err == nil {
		return !f.IsDir(), nil
	} else if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func parseAlbumMeta(meta *Release) map[string]string {
	parsedMeta := map[string]string{
		"album":       meta.Title,
		"albumArtist": strings.Join(meta.ArtistNames, ", "),
		"year":        strconv.Itoa(meta.Date)[:4],
	}
	return parsedMeta
}

func parsePlaylistMeta(meta *Playlist) map[string]string {
	parsedMeta := map[string]string{
		"album": meta.Title,
	}
	return parsedMeta
}

func parseTrackMeta(meta *Track, albMeta map[string]string, trackNum, trackTotal int) map[string]string {
	albMeta["artist"] = strings.Join(meta.ArtistNames, ", ")
	albMeta["genre"] = strings.Join(meta.Genres, ", ")
	albMeta["title"] = meta.Title
	albMeta["track"] = strconv.Itoa(trackNum)
	albMeta["trackPad"] = fmt.Sprintf("%02d", trackNum)
	albMeta["trackTotal"] = strconv.Itoa(trackTotal)
	return albMeta
}

func parseTemplate(templateText, defTemplate string, tags map[string]string) string {
	var buffer bytes.Buffer
	for {
		err := template.Must(template.New("").Parse(templateText)).Execute(&buffer, tags)
		if err == nil {
			break
		}
		fmt.Println("Failed to parse template. Default will be used instead.")
		templateText = defTemplate
		buffer.Reset()
	}
	return html.UnescapeString(buffer.String())
}

func downloadTrack(trackPath, url string, byteLimit int64) error {
	f, err := os.OpenFile(trackPath, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Range", "bytes=0-")
	do, err := client.Do(req)
	if err != nil {
		return err
	}
	defer do.Body.Close()
	if do.StatusCode != http.StatusOK && do.StatusCode != http.StatusPartialContent {
		return errors.New(do.Status)
	}
	totalBytes := do.ContentLength
	counter := &WriteCounter{Total: totalBytes, TotalStr: humanize.Bytes(uint64(totalBytes)),
		StartTime: time.Now().UnixMilli()}
	if byteLimit == -1000000 {
		_, err = io.Copy(f, io.TeeReader(do.Body, counter))
	} else {
		for range time.Tick(time.Second * 1) {
			_, err = io.CopyN(f, io.TeeReader(do.Body, counter), byteLimit)
			if errors.Is(err, io.EOF) {
				err = nil
				break
			}
			if err != nil {
				break
			}
		}
	}
	fmt.Println("")
	return err
}

func downloadPlaylistCover(playlist Playlist, path string, maxCover bool) error {
	var url string
	if maxCover {
		url = apiBase + playlist.ImageUrlBig
	} else {
		url = apiBase + playlist.ImageUrl
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	req, err := client.Get(url)
	if err != nil {
		return err
	}
	defer req.Body.Close()
	if req.StatusCode != http.StatusOK {
		return errors.New(req.Status)
	}
	_, err = io.Copy(f, req.Body)
	return err
}

func downloadAlbumCover(url, path string, maxCover bool) error {
	rep := ""
	if !maxCover {
		rep = "&size=600x600"
	}
	url = strings.Replace(url, "&size={size}", rep, 1)
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	req, err := client.Get(url)
	if err != nil {
		return err
	}
	defer req.Body.Close()
	if req.StatusCode != http.StatusOK {
		return errors.New(req.Status)
	}
	_, err = io.Copy(f, req.Body)
	return err
}

func writeFlacTags(decTrackPath string, tags map[string]string, imgData []byte) error {
	f, err := flac.ParseFile(decTrackPath)
	if err != nil {
		return err
	}
	tag := flacvorbis.New()
	for k, v := range tags {
		tag.Add(strings.ToUpper(k), v)
	}
	tagMeta := tag.Marshal()
	f.Meta = append(f.Meta, &tagMeta)
	if imgData != nil {
		picture, err := flacpicture.NewFromImageData(
			flacpicture.PictureTypeFrontCover, "", imgData, "image/jpeg",
		)
		if err != nil {
			return err
		}
		pictureMeta := picture.Marshal()
		f.Meta = append(f.Meta, &pictureMeta)
	}
	return f.Save(decTrackPath)
}

func writeMp3Tags(decTrackPath string, tags map[string]string, imgData []byte) error {
	tags["track"] += "/" + tags["trackTotal"]
	resolve := map[string]string{
		"album":       "TALB",
		"artist":      "TPE1",
		"albumArtist": "TPE2",
		"genre":       "TCON",
		"title":       "TIT2",
		"track":       "TRCK",
		"year":        "TYER",
	}
	tag, err := id3v2.Open(decTrackPath, id3v2.Options{Parse: true})
	if err != nil {
		return err
	}
	defer tag.Close()
	for k, v := range tags {
		resolved, ok := resolve[k]
		if ok {
			tag.AddTextFrame(resolved, tag.DefaultEncoding(), v)
		}
	}
	if imgData != nil {
		imgFrame := id3v2.PictureFrame{
			Encoding:    id3v2.EncodingUTF8,
			MimeType:    "image/jpeg",
			PictureType: id3v2.PTFrontCover,
			Picture:     imgData,
		}
		tag.AddAttachedPicture(imgFrame)
	}
	return tag.Save()
}

func writeTags(decTrackPath, coverPath string, isFlac bool, tags map[string]string, writeCover bool) error {
	var (
		err     error
		imgData []byte
	)
	if coverPath != "" && writeCover {
		imgData, err = ioutil.ReadFile(coverPath)
		if err != nil {
			return err
		}
	}
	delete(tags, "trackPad")
	if isFlac {
		err = writeFlacTags(decTrackPath, tags, imgData)
	} else {
		err = writeMp3Tags(decTrackPath, tags, imgData)
	}
	return err
}

func getLyrics(trackId, token string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, apiBase+"api/tiny/musixmatch/lyrics", nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("x-auth-token", token)
	query := url.Values{}
	query.Set("track_id", trackId)
	req.URL.RawQuery = query.Encode()
	do, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer do.Body.Close()
	if do.StatusCode != http.StatusOK {
		return "", errors.New(do.Status)
	}
	var obj Lyrics
	err = json.NewDecoder(do.Body).Decode(&obj)
	if err != nil {
		return "", err
	}
	if obj.Result.Lyrics == nil {
		return "", nil
	}
	return obj.Result.Lyrics.(string), nil
}

func writeLyrics(lyrics, path string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	_, err = f.WriteString(lyrics)
	f.Close()
	return err
}

func downloadTracks(trackIds []int, meta *Meta, cfg *Config, parsedAlbMeta map[string]string, token string, albumPath string, coverPath string, writeCover bool) {
	trackTotal := len(trackIds)
	for trackNum, trackId := range trackIds {
		trackNum++
		trackIdStr := strconv.Itoa(trackId)
		track := meta.Result.Tracks[trackIdStr]
		parsedMeta := parseTrackMeta(&track, parsedAlbMeta, trackNum, trackTotal)
		if cfg.Format != 1 {
			queryQualities(&track, cfg)
		}
		streamMeta, err := getStreamMeta(trackIdStr, cfg.FormatStr, token)
		if err != nil {
			handleErr("Failed to get track stream metadata.", err, false)
			continue
		}
		streamUrl := streamMeta.Result.Stream
		quality := queryRetQuality(streamUrl)
		if quality == nil {
			fmt.Println("The API returned an unsupported format.")
			continue
		}
		trackFname := parseTemplate(cfg.TrackTemplate, trackTemplate, parsedMeta)
		sanTrackFname := sanitize(trackFname, false)
		trackPath := filepath.Join(albumPath, sanTrackFname+quality.Extension)
		exists, err := fileExists(trackPath)
		if err != nil {
			handleErr("Failed to check if track already exists locally.", err, false)
			continue
		}
		if exists {
			fmt.Println("Track already exists locally.")
			continue
		}
		fmt.Printf(
			"Downloading track %d of %d: %s - %s\n", trackNum, trackTotal, parsedMeta["title"],
			quality.Specs,
		)
		err = downloadTrack(trackPath, streamUrl, cfg.ByteLimit)
		if err != nil {
			handleErr("Failed to download track.", err, false)
			continue
		}
		err = writeTags(trackPath, coverPath, quality.IsFlac, parsedMeta, writeCover)
		if err != nil {
			handleErr("Failed to write tags.", err, false)
			continue
		}
		if cfg.Lyrics {
			lyrics, err := getLyrics(trackIdStr, token)
			if err != nil {
				handleErr("Failed to get lyrics.", err, false)
				continue
			}
			if lyrics == "" {
				continue
			}
			lyricsPath := filepath.Join(albumPath, sanTrackFname+".lrc")
			err = writeLyrics(lyrics, lyricsPath)
			if err != nil {
				handleErr("Failed to write lyrics.", err, false)
				continue
			}
			fmt.Println("Wrote lyrics.")
		}
	}
	if coverPath != "" && !cfg.KeepCover {
		err := os.Remove(coverPath)
		if err != nil {
			handleErr("Failed to delete cover.", err, false)
		}
	}
}

func init() {
	fmt.Println(`
 _____         _      ____                _           _         
|__   |_ _ _ _| |_   |    \ ___ _ _ _ ___| |___ ___ _| |___ ___ 
|   __| | | | | '_|  |  |  | . | | | |   | | . | .'| . | -_|  _|
|_____|\_/|___|_,_|  |____/|___|_____|_|_|_|___|__,|___|___|_|
   `)
}

func main() {
	scriptDir, err := getScriptDir()
	if err != nil {
		panic(err)
	}
	err = os.Chdir(scriptDir)
	if err != nil {
		panic(err)
	}
	cfg, err := parseCfg()
	if err != nil {
		handleErr("Failed to parse config file.", err, true)
	}
	err = makeDirs(cfg.OutPath)
	if err != nil {
		handleErr("Failed to make output path.", err, true)
	}
	token, err := getToken(cfg.Token, cfg.Email, cfg.Password)
	if err != nil {
		handleErr("Failed to auth.", err, true)
	}
	userInfo, err := getUserInfo(token)
	if err != nil {
		handleErr("Failed to get user info.", err, true)
	}
	if reflect.ValueOf(userInfo.Result.Subscription).IsZero() {
		panic("Subscription required.")
	}
	fmt.Println(
		"Signed in successfully - " + userInfo.Result.Subscription.Name + "\n",
	)
	itemTotal := len(cfg.Urls)
	for itemNum, url := range cfg.Urls {
		fmt.Printf("Album/Playlist %d of %d:\n", itemNum+1, itemTotal)
		itemType := checkUrl(url)
		switch itemType.TypeId {
		case 0:
			// not supported yet type
			fmt.Println("Invalid URL:", url)
			continue

		case 1:
			//album
			meta, err := getMeta(itemType, token)
			if err != nil {
				handleErr("Failed to get album metadata.", err, false)
				continue
			}
			albumRelease := meta.Result.Releases[itemType.ItemId]
			parsedMeta := parseAlbumMeta(&albumRelease)
			albumFolder := parseTemplate(cfg.AlbumTemplate, albumTemplate, parsedMeta)
			fmt.Println(parsedMeta["albumArtist"] + " - " + parsedMeta["album"])
			if len(albumFolder) > 120 {
				fmt.Println("Album folder was chopped as it exceeds 120 characters.")
				albumFolder = albumFolder[:120]
			}
			sanAlbumFolder := sanitize(albumFolder, true)
			path := filepath.Join(cfg.OutPath, strings.TrimSuffix(sanAlbumFolder, "."))
			err = makeDirs(path)
			if err != nil {
				handleErr("Failed to make album folder.", err, false)
				continue
			}
			coverPath := filepath.Join(path, "cover.jpg")
			err = downloadAlbumCover(albumRelease.Image.Src, coverPath, cfg.MaxCover)
			if err != nil {
				handleErr("Failed to get cover.", err, false)
				coverPath = ""
			}
			trackIds := albumRelease.TrackIds
			downloadTracks(trackIds, meta, cfg, parsedMeta, token, path, coverPath, true)

		case 2:
			//playlist
			meta, err := getMeta(itemType, token)
			if err != nil {
				handleErr("Failed to get playlist metadata.", err, false)
				continue
			}
			playlist := meta.Result.Playlists[itemType.ItemId]
			playlistFolder := playlist.Title
			if len(playlistFolder) > 120 {
				fmt.Println("Playlist folder was chopped as it exceeds 120 characters.")
				playlistFolder = playlistFolder[:120]
			}
			sanPlaylistFolder := sanitize(playlistFolder, true)
			path := filepath.Join(cfg.OutPath, strings.TrimSuffix(sanPlaylistFolder, "."))
			err = makeDirs(path)
			if err != nil {
				handleErr("Failed to make playlist folder.", err, false)
				continue
			}
			coverPath := filepath.Join(path, "cover.jpg")
			err = downloadPlaylistCover(playlist, coverPath, cfg.MaxCover)
			if err != nil {
				handleErr("Failed to get cover.", err, false)
				coverPath = ""
			}

			parsedMeta := parsePlaylistMeta(&playlist)
			downloadTracks(playlist.TrackIds, meta, cfg, parsedMeta, token, path, coverPath, false)
		}
	}
}
