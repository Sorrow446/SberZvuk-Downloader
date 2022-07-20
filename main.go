package main

import (
	"bufio"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-flac/flacpicture"
	"github.com/machinebox/graphql"
	_ "github.com/mattn/go-sqlite3"
	"html"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
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
	"unicode"

	"github.com/alexflint/go-arg"
	"github.com/bogem/id3v2"
	"github.com/dustin/go-humanize"
	"github.com/go-flac/flacvorbis"
	"github.com/go-flac/go-flac"
)

const (
	megabyte              = 1000000
	apiBase               = "https://zvuk.com/"
	qraphql               = "api/v1/graphql"
	albumRegexString      = `^https://zvuk.com/release/(\d+)$`
	playlistRegexString   = `^https://zvuk.com/playlist/(\d+)$`
	artistRegexString     = `^https://zvuk.com/artist/(\d+)$`
	trackTemplateAlbum    = "{{.trackPad}}-{{.title}}"
	trackTemplatePlaylist = "{{.artist}} - {{.title}}"
	albumTemplate         = "{{.albumArtist}} - {{.album}}"
	authHeader            = "x-auth-token"
)

var userAgents = []string{
	"OpenPlay|4.9.4|Android|7.1|HTC One X10",
	"OpenPlay|4.10.1|Android|7.1.2|Sony Xperia Z5",
	"OpenPlay|4.10.2|Android|7.1|Sony Xperia XZ",
	"OpenPlay|4.10.3|Android|7.1.2|Asus ASUS_Z01QD",
	"OpenPlay|4.11.2|Android|8|Nexus 6P",
	"OpenPlay|4.11.4|Android|8.1|Samsung Galaxy S6",
	"OpenPlay|4.11.5|Android|9|Samsung Galaxy S7",
	"OpenPlay|4.12.3|Android|10|Samsung Galaxy S8",
	"OpenPlay|4.13|Android|11|Samsung Galaxy S9",
	"OpenPlay|4.14|Android|12|Google Pixel 4 XL",
}

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
		"User-Agent", userAgents[rand.Int()%len(userAgents)],
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
		speed = wc.Downloaded / toDivideBy * 1000
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

func randomPause(minPause, duration int) {
	time.Sleep(time.Duration(minPause+rand.Intn(duration)) * time.Second)
}

func wasRunFromSrc() bool {
	buildPath := filepath.Join(os.TempDir(), "go-build")
	return strings.HasPrefix(os.Args[0], buildPath)
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
		cfg.TrackTemplate = trackTemplateAlbum
	}
	if args.PlaylistTemplate != "" {
		cfg.TrackTemplate = trackTemplatePlaylist
	}
	if args.AlbumTemplate != "" {
		cfg.AlbumTemplate = albumTemplate
	}
	if cfg.OutPath == "" {
		cfg.OutPath = "Zvuk downloads"
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

func getToken(email, pwd string) (string, error) {
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
	req.Header.Add(authHeader, token)
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
	matchArtist := regexp.MustCompile(artistRegexString).FindStringSubmatch(url)
	if matchArtist == nil {
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
	return ItemType{3, matchArtist[1]}
}

func getMeta(apiUrl string, itemId string, token string) (*Meta, error) {
	req, err := http.NewRequest(http.MethodGet, apiBase+apiUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(authHeader, token)
	query := url.Values{}
	query.Set("ids", itemId)
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

func getStreamMeta(trackId, formatStr, token string) (*StreamMeta, error) {
	var do *http.Response
	req, err := http.NewRequest(http.MethodGet, apiBase+"api/tiny/track/stream", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(authHeader, token)
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
			randomPause(3, 7)
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
	highestQuality := track.HighestQuality
	if format == 3 && highestQuality != "flac" {
		avail = false
	} else if format == 2 && highestQuality == "med" {
		avail = false
	}
	if !avail {
		fmt.Println("Track unavailable in your chosen quality.")
		cfg.FormatStr = highestQuality
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

func uniqueStrings(stringSlices ...[]string) []string {
	uniqueMap := map[string]bool{}

	for _, intSlice := range stringSlices {
		for _, number := range intSlice {
			uniqueMap[number] = true
		}
	}

	result := make([]string, 0, len(uniqueMap))

	for key := range uniqueMap {
		result = append(result, key)
	}

	return result
}

func getArtistAlbumId(artistId string, limit int, offset int, token string) []string {
	graphqlClient := graphql.NewClient(apiBase + qraphql)
	graphqlRequest := graphql.NewRequest(`
				query getArtistReleases($id: ID!, $limit: Int!, $offset: Int!) { getArtists(ids: [$id]) { __typename releases(limit: $limit, offset: $offset) { __typename ...ReleaseGqlFragment } } } fragment ReleaseGqlFragment on Release { id }
			`)
	graphqlRequest.Var("id", artistId)
	graphqlRequest.Var("limit", limit)
	graphqlRequest.Var("offset", offset)
	graphqlRequest.Header.Add(authHeader, token)
	var graphqlResponse interface{}
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
	}
	jsonString, _ := json.Marshal(graphqlResponse)
	var obj ArtistReleases
	json.Unmarshal(jsonString, &obj)

	var albumsIds = make([]string, 0)
	if len(obj.GetArtists) > 0 {
		for _, element := range obj.GetArtists[0].Releases {
			if element.ID != "" {
				albumsIds = append(albumsIds, element.ID)
			}
		}
	}
	return albumsIds
}

func parseAlbumMeta(meta *Release) map[string]string {
	parsedMeta := map[string]string{
		"album":       sanitize(meta.Title, false),
		"albumArtist": strings.Join(meta.ArtistNames, ", "),
		"year":        strconv.Itoa(meta.Date)[:4],
		"tracks":      strconv.Itoa(len(meta.TrackIds)),
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

func parseTrackTemplate(isPlaylist bool, cfg *Config, trackTotal int, tags map[string]string) string {
	var buffer bytes.Buffer
	var templateText string
	var defTemplate string

	if isPlaylist {
		templateText = cfg.PlaylistTemplate
		defTemplate = trackTemplatePlaylist
	} else {
		templateText = cfg.TrackTemplate
		defTemplate = trackTemplateAlbum
		if cfg.FolderForSingle == false && trackTotal == 1 {
			templateText = cfg.PlaylistTemplate
			defTemplate = trackTemplatePlaylist
		}
	}

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

func parseAlbumTemplate(cfg *Config, tags map[string]string) string {
	var buffer bytes.Buffer
	templateText := cfg.AlbumTemplate
	tracks, _ := strconv.Atoi(tags["tracks"])
	if cfg.FolderForSingle == false && tracks == 1 {
		templateText = "{{.albumArtist}}"
	}

	for {
		err := template.Must(template.New("").Parse(templateText)).Execute(&buffer, tags)
		if err == nil {
			break
		}
		fmt.Println("Failed to parse template. Default will be used instead.")
		templateText = albumTemplate
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

func downloadRelease(itemId string, token string, cfg *Config) {
	meta, err := getMeta("api/tiny/releases", itemId, token)
	if err != nil || reflect.ValueOf(meta.Result.Releases).IsZero() || len(meta.Result.Releases) == 0 {
		handleErr("Failed to get album metadata.", err, false)
		return
	}

	albumRelease := meta.Result.Releases[itemId]
	parsedMeta := parseAlbumMeta(&albumRelease)
	albumFolder := parseAlbumTemplate(cfg, parsedMeta)
	fmt.Println(parsedMeta["albumArtist"] + " - " + parsedMeta["album"])
	if len(albumFolder) > 120 {
		fmt.Println("Album folder was chopped as it exceeds 120 characters.")
		albumFolder = albumFolder[:120]
	}
	sanAlbumFolder := sanitize(albumFolder, true)
	path := filepath.Join(cfg.OutPath, strings.TrimRightFunc(sanAlbumFolder, func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsNumber(r) }))
	err = makeDirs(path)
	if err != nil {
		handleErr("Failed to make album folder.", err, false)
		return
	}
	coverPath := filepath.Join(path, "cover.jpg")
	err = downloadAlbumCover(albumRelease.Image.Src, coverPath, cfg.MaxCover)
	if err != nil {
		handleErr("Failed to get cover.", err, false)
		coverPath = ""
	}

	downloadTracks(albumRelease.TrackIds, meta, cfg, parsedMeta, token, path, coverPath, true, false)
}

func downloadPlaylist(itemId string, token string, cfg *Config) {
	meta, err := getMeta("api/tiny/playlists", itemId, token)
	if err != nil || reflect.ValueOf(meta.Result.Playlists).IsZero() || len(meta.Result.Playlists) == 0 {
		handleErr("Failed to get playlist metadata.", err, false)
		return
	}

	playlist := meta.Result.Playlists[itemId]
	playlistFolder := playlist.Title
	if len(playlistFolder) > 120 {
		fmt.Println("Playlist folder was chopped as it exceeds 120 characters.")
		playlistFolder = playlistFolder[:120]
	}
	sanPlaylistFolder := sanitize(playlistFolder, false)
	path := filepath.Join(cfg.OutPath, strings.TrimSuffix(sanPlaylistFolder, "."))
	err = makeDirs(path)
	if err != nil {
		handleErr("Failed to make playlist folder.", err, false)
	}
	coverPath := filepath.Join(path, "cover.jpg")
	err = downloadPlaylistCover(playlist, coverPath, cfg.MaxCover)
	if err != nil {
		handleErr("Failed to get cover.", err, false)
		coverPath = ""
	}

	parsedMeta := parsePlaylistMeta(&playlist)
	downloadTracks(playlist.TrackIds, meta, cfg, parsedMeta, token, path, coverPath, false, true)
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

func extractFLACComment(fileName string) (*flacvorbis.MetaDataBlockVorbisComment, int) {
	f, err := flac.ParseFile(fileName)
	if err != nil {
		panic(err)
	}

	var cmt *flacvorbis.MetaDataBlockVorbisComment
	var cmtIdx int
	for idx, meta := range f.Meta {
		if meta.Type == flac.VorbisComment {
			cmt, err = flacvorbis.ParseFromMetaDataBlock(*meta)
			cmtIdx = idx
			if err != nil {
				panic(err)
			}
		}
	}
	return cmt, cmtIdx
}

func writeFlacTags(decTrackPath string, tags map[string]string, imgData []byte) error {
	f, err := flac.ParseFile(decTrackPath)
	if err != nil {
		return err
	}
	tag, idx := extractFLACComment(decTrackPath)
	if tag == nil && idx > 0 {
		tag = flacvorbis.New()
	}
	for k, v := range tags {
		tag.Add(strings.ToUpper(k), v)
	}
	tagMeta := tag.Marshal()
	if idx > 0 {
		f.Meta[idx] = &tagMeta
	} else {
		f.Meta = append(f.Meta, &tagMeta)
	}
	if imgData != nil {
		picture, err := flacpicture.NewFromImageData(
			flacpicture.PictureTypeFrontCover, "", imgData, "image/jpeg",
		)
		if err != nil {
			handleErr("Tag picture error", err, false)
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
		tags["DATE"] = tags["year"]
		tags["PERFORMER"] = tags["albumArtist"]
		tags["TRACKNUMBER"] = tags["track"]
		delete(tags, "trackTotal")
		/*delete(tags, "year")
		delete(tags, "albumArtist")
		delete(tags, "track")*/
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
	req.Header.Add(authHeader, token)
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

func downloadTracks(trackIds []int, meta *Meta, cfg *Config, parsedAlbMeta map[string]string, token string, albumPath string, coverPath string, writeCover bool, isPlaylist bool) {
	trackTotal := len(trackIds)
	for trackNum, trackId := range trackIds {
		randomPause(1, 5)
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

		var trackName string
		trackName = parseTrackTemplate(isPlaylist, cfg, trackTotal, parsedMeta)
		sanTrackName := sanitize(trackName, false)
		trackPath := filepath.Join(albumPath, sanTrackName+quality.Extension)
		exists, err := fileExists(trackPath)
		if err != nil {
			handleErr("Failed to check if track already exists locally.", err, false)
			continue
		}
		if exists {
			fmt.Println("Track already exists locally.")
			continue
		}
		fmt.Printf("Downloading track %d of %d: %s - %s\n", trackNum, trackTotal, parsedMeta["title"], quality.Specs)
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
			lyricsPath := filepath.Join(albumPath, sanTrackName+".lrc")
			err = writeLyrics(lyrics, lyricsPath)
			if err != nil {
				handleErr("Failed to write lyrics.", err, false)
				continue
			}
			fmt.Println("Wrote lyrics.")
		}
		if cfg.FolderForSingle == false && trackTotal == 1 {
			err := os.Remove(coverPath)
			if err != nil {
				handleErr("Failed to delete cover.", err, false)
			}
		}
	}
	if coverPath != "" && !cfg.KeepCover {
		err := os.Remove(coverPath)
		if err != nil {
			handleErr("Failed to delete cover.", err, false)
		}
	}
}

func createTokenTable(dbFile string) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `create table authToken (id integer not null primary key, token text);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}
	defer db.Close()
}

func insertTokenDb(dbFile string, tokenValue string) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("delete from authToken")
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into authToken(id, token) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(0, tokenValue)
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func getTokenFromDb(dbFile string) (error, string) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return err, ""
	}

	stmt, err := db.Prepare("select token from authToken limit 1")
	if err != nil {
		return err, ""
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow().Scan(&name)
	if err != nil {
		return err, ""
	}
	return nil, name
}

func createAndFillDbWithToken(cfg *Config, needTable bool) string {
	if needTable {
		createTokenTable("./sqlite.db")
	}

	token, err := getToken(cfg.Email, cfg.Password)
	if err != nil {
		handleErr("Failed to auth.", err, true)
	}
	insertTokenDb("./sqlite.db", token)
	return token
}

func tokenDb(dbFile string, cfg *Config) (string, *UserInfo) {
	dbExist, _ := fileExists(dbFile)
	if dbExist {
		err, token := getTokenFromDb(dbFile)
		if err != nil {
			os.Remove(dbFile)
			return createAndFillDbWithToken(cfg, true), nil
		} else {
			userInfo, err := getUserInfo(token)
			if err != nil {
				handleErr("Failed to get user info, re-login..", err, true)
				return createAndFillDbWithToken(cfg, false), nil
			}
			return token, userInfo
		}
	} else {
		return createAndFillDbWithToken(cfg, true), nil
	}
}

func init() {
	fmt.Println(`
 _____         _      ____                _           _         
|__   |_ _ _ _| |_   |    \ ___ _ _ _ ___| |___ ___ _| |___ ___ 
|   __| | | | | '_|  |  |  | . | | | |   | | . | .'| . | -_|  _|
|_____|\_/|___|_,_|  |____/|___|_____|_|_|_|___|__,|___|___|_|*
   `)
}

func main() {
	rand.Seed(time.Now().UnixNano())
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

	token, userInfo := tokenDb("./sqlite.db", cfg)
	if userInfo == nil {
		us, err := getUserInfo(token)
		if err != nil {
			handleErr("Failed to get user info.", err, true)
		} else {
			userInfo = us
		}
	}
	if reflect.ValueOf(userInfo.Result.Subscription).IsZero() {
		panic("Subscription required.")
	}
	fmt.Println("Signed in successfully - " + userInfo.Result.Subscription.Name)
	t := time.UnixMilli(userInfo.Result.Subscription.Expiration)
	fmt.Println("Subscription valid until: " + t.Format(time.RFC822))

	err = makeDirs(cfg.OutPath)
	if err != nil {
		handleErr("Failed to make output path.", err, true)
	}

	var downloadItems []ItemType
	for _, url := range cfg.Urls {
		itemType := checkUrl(url)
		downloadItems = append(downloadItems, itemType)
	}

	var releaseIds []string
	itemTotal := len(downloadItems)
	for itemNum, itemType := range downloadItems {
		if itemType.TypeId == 3 {
			//artist
			fmt.Printf("Get all artist releases...")
			firstFifty := getArtistAlbumId(itemType.ItemId, 50, 0, token)
			lastFifty := getArtistAlbumId(itemType.ItemId, 50, 49, token)
			releaseIds = uniqueStrings(firstFifty, lastFifty)

		} else {
			switch itemType.TypeId {
			case 1:
				fmt.Printf("Album %d of %d:\n", itemNum+1, itemTotal)
				downloadRelease(itemType.ItemId, token, cfg)
			case 2:
				fmt.Printf("Playlist %d of %d:\n", itemNum+1, itemTotal)
				downloadPlaylist(itemType.ItemId, token, cfg)
			}
		}
	}

	if len(releaseIds) > 0 {
		itemTotal := len(releaseIds)
		for itemNum, releaseId := range releaseIds {
			fmt.Printf("Album %d of %d:\n", itemNum+1, itemTotal)
			downloadRelease(releaseId, token, cfg)
			randomPause(3, 7)
		}
	}
}
