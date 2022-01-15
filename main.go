package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

	"github.com/alexflint/go-arg"
	"github.com/bogem/id3v2"
	"github.com/dustin/go-humanize"
	"github.com/go-flac/flacpicture"
	"github.com/go-flac/flacvorbis"
	"github.com/go-flac/go-flac"
)

const apiBase = "https://sber-zvuk.com/"

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
		"User-Agent", "OpenPlay|4.9.1|Android|7.1.2|samsung SM-N976N",
	)
	req.Header.Add(
		"Referer", "https://sber-zvuk.com/",
	)
	return http.DefaultTransport.RoundTrip(req)
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Downloaded += uint64(n)
	percentage := float64(wc.Downloaded) / float64(wc.Total) * float64(100)
	wc.Percentage = int(percentage)
	fmt.Printf("\r%d%%, %s/%s ", wc.Percentage, humanize.Bytes(wc.Downloaded), wc.TotalStr)
	return n, nil
}

func initErr(errText string, err error) {
	errString := fmt.Sprintf("%s\n%s", errText, err)
	panic(errString)
}

func getScriptDir() (string, error) {
	var (
		ok    bool
		err   error
		fname string
	)
	if filepath.IsAbs(os.Args[0]) {
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
	scriptDir := filepath.Dir(fname)
	return scriptDir, nil
}

func readTxtFile(path string) ([]string, error) {
	var lines []string
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
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
	for _, url := range urls {
		if strings.HasSuffix(url, ".txt") && !contains(txtPaths, url) {
			txtLines, err := readTxtFile(url)
			if err != nil {
				return nil, err
			}
			for _, txtLine := range txtLines {
				if !contains(processed, txtLine) {
					processed = append(processed, txtLine)
				}
			}
			txtPaths = append(txtPaths, url)
		} else {
			if !contains(processed, url) {
				processed = append(processed, url)
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
		cfg.OutPath = "SberZvuk downloads"
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
	err := os.MkdirAll(path, 0755)
	return err
}

func checkUrl(url string) string {
	const regexString = `^https://sber-zvuk.com/release/(\d+)$`
	regex := regexp.MustCompile(regexString)
	match := regex.FindStringSubmatch(url)
	if match == nil {
		return ""
	}
	return match[1]
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

func getMeta(albumId, token string) (*Meta, error) {
	req, err := http.NewRequest(http.MethodGet, apiBase+"api/tiny/releases", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-auth-token", token)
	query := url.Values{}
	query.Set("ids", albumId)
	query.Set("include", "track,")
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

func sanitize(filename string) string {
	regex := regexp.MustCompile(`[\/:*?"><|]`)
	sanitized := regex.ReplaceAllString(filename, "_")
	return sanitized
}

func parseTrackIds(trackIds []int) string {
	var parsed string
	for _, trackId := range trackIds {
		parsed += strconv.Itoa(trackId) + ","
	}
	return parsed[:len(parsed)-1]
}

func getStreamMeta(trackId, formatStr, token string) (*StreamMeta, error) {
	req, err := http.NewRequest(http.MethodGet, apiBase+"api/tiny/track/stream", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-auth-token", token)
	query := url.Values{}
	query.Set("id", trackId)
	query.Set("quality", formatStr)
	req.URL.RawQuery = query.Encode()
	do, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer do.Body.Close()
	if do.StatusCode != http.StatusOK {
		return nil, errors.New(do.Status)
	}
	var obj StreamMeta
	err = json.NewDecoder(do.Body).Decode(&obj)
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
		"/streamfl?": {"16-bit / 44.1 kHz FLAC", ".flac", true},
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

func parseTrackMeta(meta *Track, albMeta map[string]string, trackNum, trackTotal int) map[string]string {
	albMeta["artist"] = strings.Join(meta.ArtistNames, ", ")
	albMeta["genre"] = strings.Join(meta.Genres, ", ")
	albMeta["title"] = meta.Title
	albMeta["track"] = strconv.Itoa(trackNum)
	albMeta["trackPad"] = fmt.Sprintf("%02d", trackNum)
	albMeta["trackTotal"] = strconv.Itoa(trackTotal)
	return albMeta
}

func parseTemplate(templateText string, tags map[string]string) string {
	var buffer bytes.Buffer
	for {
		err := template.Must(template.New("").Parse(templateText)).Execute(&buffer, tags)
		if err == nil {
			break
		}
		fmt.Println("Failed to parse template. Default will be used instead.")
		templateText = "{{.trackPad}}. {{.title}}"
		buffer.Reset()
	}
	return buffer.String()
}

func downloadTrack(trackPath, url string) error {
	f, err := os.Create(trackPath)
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
	totalBytes := uint64(do.ContentLength)
	counter := &WriteCounter{Total: totalBytes, TotalStr: humanize.Bytes(totalBytes)}
	_, err = io.Copy(f, io.TeeReader(do.Body, counter))
	fmt.Println("")
	return err
}

func downloadCover(url, path string, maxCover bool) error {
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
	err = f.Save(decTrackPath)
	return err
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
	err = tag.Save()
	return err
}

func writeTags(decTrackPath, coverPath string, isFlac bool, tags map[string]string) error {
	var (
		err     error
		imgData []byte
	)
	if coverPath != "" {
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
	defer f.Close()
	f.Write([]byte(lyrics))
	return nil
}

func main() {
	fmt.Println(`
 _____ _           _____         _      ____                _           _         
|   __| |_ ___ ___|__   |_ _ _ _| |_   |    \ ___ _ _ _ ___| |___ ___ _| |___ ___ 
|__   | . | -_|  _|   __| | | | | '_|  |  |  | . | | | |   | | . | .'| . | -_|  _|
|_____|___|___|_| |_____|\_/|___|_,_|  |____/|___|_____|_|_|_|___|__,|___|___|_|
`)
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
		initErr("Failed to parse config file.", err)
	}
	err = makeDirs(cfg.OutPath)
	if err != nil {
		initErr("Failed to make output folder.", err)
	}
	token, err := auth(cfg.Email, cfg.Password)
	if err != nil {
		initErr("Failed to auth.", err)
	}
	userInfo, err := getUserInfo(token)
	if err != nil {
		initErr("Failed to get user info.", err)
	}
	if reflect.ValueOf(userInfo.Result.Subscription).IsZero() {
		panic("Subscription required.")
	}
	fmt.Println(
		"Signed in successfully - " + userInfo.Result.Subscription.Name + "\n",
	)
	albumTotal := len(cfg.Urls)
	for albumNum, url := range cfg.Urls {
		fmt.Printf("Album %d of %d:\n", albumNum+1, albumTotal)
		albumId := checkUrl(url)
		if albumId == "" {
			fmt.Println("Invalid URL:", url)
			continue
		}
		meta, err := getMeta(albumId, token)
		if err != nil {
			fmt.Printf("Failed to fetch album metadata.\n%s", err)
			continue
		}
		albumRelease := meta.Result.Releases[albumId]
		parsedAlbMeta := parseAlbumMeta(&albumRelease)
		albumFolder := parsedAlbMeta["albumArtist"] + " - " + parsedAlbMeta["album"]
		fmt.Println(albumFolder)
		if len(albumFolder) > 120 {
			fmt.Println("Album folder was chopped as it exceeds 120 characters.")
			albumFolder = albumFolder[:120]
		}
		albumPath := filepath.Join(cfg.OutPath, sanitize(albumFolder))
		err = makeDirs(albumPath)
		if err != nil {
			fmt.Println("Failed to make album folder.\n", err)
			continue
		}
		trackIds := albumRelease.TrackIds
		coverPath := filepath.Join(albumPath, "cover.jpg")
		err = downloadCover(albumRelease.Image.Src, coverPath, cfg.MaxCover)
		if err != nil {
			fmt.Println("Failed to get cover.\n", err)
			coverPath = ""
		}
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
				fmt.Println("Failed to get track stream meta.\n", err)
				continue
			}
			streamUrl := streamMeta.Result.Stream
			quality := queryRetQuality(streamUrl)
			if quality == nil {
				fmt.Println("The API returned an unsupported format.")
				continue
			}
			trackFname := parseTemplate(cfg.TrackTemplate, parsedMeta)
			sanTrackFname := sanitize(trackFname)
			trackPath := filepath.Join(albumPath, sanTrackFname+quality.Extension)
			exists, err := fileExists(trackPath)
			if err != nil {
				fmt.Printf("Failed to check if track already exists locally.\n%s", err)
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
			err = downloadTrack(trackPath, streamUrl)
			if err != nil {
				fmt.Printf("Failed to download track.\n%s", err)
				continue
			}
			err = writeTags(trackPath, coverPath, quality.IsFlac, parsedMeta)
			if err != nil {
				fmt.Printf("Failed to write tags.\n%s", err)
				continue
			}
			if cfg.Lyrics {
				lyrics, err := getLyrics(trackIdStr, token)
				if err != nil {
					fmt.Printf("Failed to get lyrics.\n%s", err)
					break
				}
				if lyrics == "" {
					break
				}
				lyricsPath := filepath.Join(albumPath, sanTrackFname+".lrc")
				err = writeLyrics(lyrics, lyricsPath)
				if err != nil {
					fmt.Printf("Failed to write lyrics.\n%s", err)
					break
				}
				fmt.Println("Wrote lyrics.")
			}
		}
	}
}
