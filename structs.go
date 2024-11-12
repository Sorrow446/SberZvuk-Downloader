package main

type Transport struct{}

type WriteCounter struct {
	Total      int64
	TotalStr   string
	Downloaded int64
	Percentage int
	StartTime  int64
}

type Config struct {
	Email            string
	Password         string
	Urls             []string
	Format           int
	FormatStr        string
	OutPath          string
	AlbumTemplate    string
	TrackTemplate    string
	PlaylistTemplate string
	MaxCover         bool
	Lyrics           bool
	SpeedLimit       float64
	ByteLimit        int64
	KeepCover        bool
	FolderForSingle  bool
}

type Args struct {
	Urls             []string `arg:"positional, required"`
	Format           int      `arg:"-f" default:"-1" help:"Download format. 1 = 128 Kbps MP3, 2 = 320 Kbps MP3, 3 = 16/44 FLAC."`
	OutPath          string   `arg:"-o" help:"Where to download to. Path will be made if it doesn't already exist."`
	MaxCover         bool     `arg:"-m" help:"true = max cover size, false = 600x600."`
	Lyrics           bool     `arg:"-l" help:"Get lyrics if available."`
	AlbumTemplate    string   `arg:"-a" help:"Album folder naming template. Vars: album, albumArtist, year."`
	TrackTemplate    string   `arg:"-t" help:"Track filename naming template. Vars: album, albumArtist, artist, genre, title, track, trackPad, trackTotal, year."`
	PlaylistTemplate string   `arg:"-p" help:"Playlist filename naming template. Vars: album, albumArtist, artist, genre, title, track, trackPad, trackTotal, year."`
	SpeedLimit       float64  `arg:"-L" default:"-1" help:"Download speed limit in megabytes. Example: 0.5 = 500 kB/s, 1 = 1 MB/s, 1.5 = 1.5 MB/s. -1 = unlimited."`
}

type ArtistReleases struct {
	GetArtists []struct {
		Typename string `json:"__typename"`
		Releases []struct {
			Typename string `json:"__typename"`
			ID       string `json:"id"`
		} `json:"releases"`
	} `json:"getArtists"`
}

type Auth struct {
	Result struct {
		Token string `json:"token"`
	} `json:"result"`
}

type UserInfo struct {
	Result struct {
		Subscription struct {
			Name       string `json:"name"`
			Expiration int64  `json:"expiration"`
		} `json:"subscription"`
	} `json:"result"`
}

type Release struct {
	ArtistIds []int  `json:"artist_ids"`
	Template  string `json:"template"`
	Title     string `json:"title"`
	Image     struct {
		Src           string `json:"src"`
		Palette       string `json:"palette"`
		PaletteBottom string `json:"palette_bottom"`
	} `json:"image"`
	SearchTitle   string   `json:"search_title"`
	Explicit      bool     `json:"explicit"`
	SearchCredits string   `json:"search_credits"`
	TrackIds      []int    `json:"track_ids"`
	ArtistNames   []string `json:"artist_names"`
	Credits       string   `json:"credits"`
	LabelID       int      `json:"label_id"`
	Availability  int      `json:"availability"`
	HasImage      bool     `json:"has_image"`
	Date          int      `json:"date"`
	Price         int      `json:"price"`
	Type          string   `json:"type"`
	ID            int      `json:"id"`
	GenreIds      []int    `json:"genre_ids"`
}

type Track struct {
	HasFlac        bool        `json:"has_flac"`
	ReleaseID      int         `json:"release_id"`
	Lyrics         interface{} `json:"lyrics"`
	Price          int         `json:"price"`
	SearchCredits  string      `json:"search_credits"`
	Credits        string      `json:"credits"`
	Duration       int         `json:"duration"`
	HighestQuality string      `json:"highest_quality"`
	ID             int         `json:"id"`
	Condition      string      `json:"condition"`
	ArtistIds      []int       `json:"artist_ids"`
	Genres         []string    `json:"genres"`
	Title          string      `json:"title"`
	SearchTitle    string      `json:"search_title"`
	Explicit       bool        `json:"explicit"`
	ReleaseTitle   string      `json:"release_title"`
	Availability   int         `json:"availability"`
	ArtistNames    []string    `json:"artist_names"`
	Template       string      `json:"template"`
	Position       int         `json:"position"`
	Image          struct {
		Src           string `json:"src"`
		Palette       string `json:"palette"`
		PaletteBottom string `json:"palette_bottom"`
	} `json:"image"`
}

type Meta struct {
	Result struct {
		Tracks     map[string]Track    `json:"tracks"`
		Playlists  map[string]Playlist `json:"playlists"`
		RadioWaves struct{}            `json:"radio_waves"`
		Releases   map[string]Release  `json:"releases"`
		Artists    struct{}            `json:"artists"`
		Labels     struct{}            `json:"labels"`
		Users      struct{}            `json:"users"`
	} `json:"result"`
}

type StreamMeta struct {
	Result struct {
		Expire      int64  `json:"expire"`
		ExpireDelta int    `json:"expire_delta"`
		Stream      string `json:"stream"`
	} `json:"result"`
}

type Quality struct {
	Specs     string
	Extension string
	IsFlac    bool
}

type Lyrics struct {
	Result struct {
		Translation interface{} `json:"translation"`
		Type        interface{} `json:"type"`
		Lyrics      interface{} `json:"lyrics"`
	} `json:"result"`
}

type ItemType struct {
	TypeId byte
	ItemId string
}

type Playlist struct {
	ImageUrl    string `json:"image_url"`
	ImageUrlBig string `json:"image_url_big"`
	Title       string `json:"title"`
	TrackIds    []int  `json:"track_ids"`
}
