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
	Email         string
	Password      string
	Urls          []string
	Format        int
	FormatStr     string
	OutPath       string
	AlbumTemplate string
	TrackTemplate string
	MaxCover      bool
	Lyrics        bool
	SpeedLimit    float64
	ByteLimit     int64
	KeepCover     bool
	Token         string
}

type Args struct {
	Urls          []string `arg:"positional, required"`
	Format        int      `arg:"-f" default:"-1" help:"Download format. 1 = 128 Kbps MP3, 2 = 320 Kbps MP3, 3 = 16/44 FLAC."`
	OutPath       string   `arg:"-o" help:"Where to download to. Path will be made if it doesn't already exist."`
	MaxCover      bool     `arg:"-m" help:"true = max cover size, false = 600x600."`
	Lyrics        bool     `arg:"-l" help:"Get lyrics if available."`
	AlbumTemplate string   `arg:"-a" help:"Album folder naming template. Vars: album, albumArtist, year."`
	TrackTemplate string   `arg:"-f" help:"Track filename naming template. Vars: album, albumArtist, artist, genre, title, track, trackPad, trackTotal, year."`
	SpeedLimit    float64  `arg:"-L" default:"-1" help:"Download speed limit in megabytes. Example: 0.5 = 500 kB/s, 1 = 1 MB/s, 1.5 = 1.5 MB/s. -1 = unlimited."`
}

type Auth struct {
	Result struct {
		Token string `json:"token"`
	} `json:"result"`
}

type UserInfo struct {
	Result struct {
		Profile struct {
			Username      string `json:"username"`
			AllowExplicit bool   `json:"allow_explicit"`
			Image         struct {
				H             int         `json:"h"`
				Palette       interface{} `json:"palette"`
				PaletteBottom interface{} `json:"palette_bottom"`
				W             int         `json:"w"`
				Src           interface{} `json:"src"`
			} `json:"image"`
			IsActive     bool        `json:"is_active"`
			Phone        interface{} `json:"phone"`
			Birthday     interface{} `json:"birthday"`
			ID           int         `json:"id"`
			IsEditor     bool        `json:"is_editor"`
			HideGender   bool        `json:"hide_gender"`
			Name         string      `json:"name"`
			IsAnonymous  bool        `json:"is_anonymous"`
			Created      int64       `json:"created"`
			Gender       string      `json:"gender"`
			Registered   int64       `json:"registered"`
			HideBirthday bool        `json:"hide_birthday"`
			Token        string      `json:"token"`
			IsRegistered bool        `json:"is_registered"`
			Email        string      `json:"email"`
		} `json:"profile"`
		Subscription struct {
			Status            string      `json:"status"`
			PartnerTitle      string      `json:"partner_title"`
			PlanID            int         `json:"plan_id"`
			Name              string      `json:"name"`
			Title             string      `json:"title"`
			IsTrial           bool        `json:"is_trial"`
			Price             float64     `json:"price"`
			IsRecurrent       bool        `json:"is_recurrent"`
			DetailsInfo       interface{} `json:"details_info"`
			Start             int64       `json:"start"`
			ServicesAvailable []string    `json:"services_available"`
			Expiration        int64       `json:"expiration"`
			Logo              struct {
				Source string `json:"source"`
				Image  string `json:"image"`
			} `json:"logo"`
			Duration       int    `json:"duration"`
			Partner        string `json:"partner"`
			App            string `json:"app"`
			ID             int    `json:"id"`
			PaymentDetails struct {
				SberAcquiringOrderID      string      `json:"sber_acquiring_order_id"`
				PayWithSbrfSpasiboBonuses interface{} `json:"pay_with_sbrf_spasibo_bonuses"`
				ConnectingToSbrfSpasibo   interface{} `json:"connecting_to_sbrf_spasibo"`
			} `json:"payment_details"`
			PlanPrice float64 `json:"plan_price"`
		} `json:"subscription"`
		AllSubscriptions []struct {
			Status            string      `json:"status"`
			PartnerTitle      string      `json:"partner_title"`
			PlanID            int         `json:"plan_id"`
			Name              string      `json:"name"`
			Title             string      `json:"title"`
			IsTrial           bool        `json:"is_trial"`
			Price             float64     `json:"price"`
			IsRecurrent       bool        `json:"is_recurrent"`
			DetailsInfo       interface{} `json:"details_info"`
			Start             int64       `json:"start"`
			ServicesAvailable []string    `json:"services_available"`
			Expiration        int64       `json:"expiration"`
			Logo              struct {
				Source string `json:"source"`
				Image  string `json:"image"`
			} `json:"logo"`
			Duration       int    `json:"duration"`
			Partner        string `json:"partner"`
			App            string `json:"app"`
			ID             int    `json:"id"`
			PaymentDetails struct {
				SberAcquiringOrderID      string      `json:"sber_acquiring_order_id"`
				PayWithSbrfSpasiboBonuses interface{} `json:"pay_with_sbrf_spasibo_bonuses"`
				ConnectingToSbrfSpasibo   interface{} `json:"connecting_to_sbrf_spasibo"`
			} `json:"payment_details"`
			PlanPrice float64 `json:"plan_price"`
		} `json:"all_subscriptions"`
		Settings struct {
			EditingGeo  bool `json:"editing_geo"`
			VkImport    bool `json:"vk_import"`
			LoginScreen struct {
			} `json:"login_screen"`
			AdsInDviews                bool     `json:"ads_in_dviews"`
			SearchResultsOrder         []string `json:"search_results_order"`
			FullscreenAdCount          int      `json:"fullscreen_ad_count"`
			PlayerID                   int      `json:"player_id"`
			ClikstreamBgUpdateDuration int      `json:"clikstream_bg_update_duration"`
			Tele2HeaderEnrichment      bool     `json:"tele2_header_enrichment"`
			AppCountForPopUp           int      `json:"app_count_for_pop_up"`
			Bundles                    struct {
				SberID struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"sber_id"`
				AboutPremiumX2Premium struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"about premium x2 premium"`
				GenresLight struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"genres_light"`
				SberPrimeSubPromo struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"SberPrime_Sub_Promo"`
				WhatSNew struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"what's_new"`
				YotaLogo struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"Yota_logo"`
				NoInternetAndroidXhdpi struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"no_internet_android_xhdpi"`
				NewCollection struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"new_collection"`
				Sbermobile struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"Sbermobile"`
				EmptyStates struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"empty states"`
				SberprimePush21152021 struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"sberprime_push_21.15.2021"`
				LavaSpherePng struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"lava_sphere.png"`
				BeelineLogo struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"beeline_logo"`
				Tele2RuBundleXxhdpi struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"tele2_ru_bundle_xxhdpi"`
				MegafonPlus struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"megafon_plus"`
				AppleGenericBundle struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"apple_generic_bundle"`
				Hifi struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"hifi"`
				LoginPhoneNewMail struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"LoginPhone_new&Mail"`
				AppleGeneric25Bundle struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"apple_generic25_bundle"`
				SberPrimePlusZSubLogo struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"SberPrime&Plus&Z_sub_logo"`
				SberbankInsurance struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"sberbank-insurance"`
				ZvukPrimeAndroidTest struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"Zvuk+prime_android_test"`
				Headphones struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"headphones"`
				SberzvukEmptyAvatarNewPng struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"sberzvuk_empty_avatar_new.png"`
				SberprimeSberprimePlusSberprimeZ struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"Sberprime_SberprimePlus_SberprimeZ"`
				XhdpiOneYear struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"xhdpi_one_year"`
				RecommendationGrid struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"recommendation_grid"`
				ZvukPlus struct {
					Enabled bool     `json:"enabled"`
					Target  []string `json:"target"`
					Archive string   `json:"archive"`
				} `json:"zvuk-plus"`
			} `json:"bundles"`
			ActionKitPages struct {
				CollectionPlaylistEmptyStateHistory struct {
					Style    string `json:"style"`
					Messages []struct {
						Text    string `json:"text"`
						Image   string `json:"image"`
						TitleEn string `json:"title_en"`
						TextEn  string `json:"text_en"`
						Title   string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL  string `json:"url"`
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"collection_playlist_empty_state_history"`
				AssistantPromo struct {
					Comment  string `json:"comment"`
					Messages []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style string `json:"style"`
						Title string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"assistant_promo"`
				AppleGenericSubscriptionOffer struct {
					Messages []struct {
						TextAz     string `json:"text_az"`
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleAz    string `json:"title_az"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleEn string `json:"title_en"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style string `json:"style"`
						Title string `json:"title"`
					} `json:"actions"`
					Source           string      `json:"source"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"apple_generic_subscription_offer"`
				Videopodcast1122021 struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL             string `json:"url"`
							Title           string `json:"title"`
							ShowsMiniPlayer bool   `json:"shows_mini_player"`
							Name            string `json:"name"`
							InWebkit        bool   `json:"in_webkit"`
						} `json:"action"`
						Style   string `json:"style"`
						Title   string `json:"title"`
						TitleUk string `json:"title_uk,omitempty"`
						TitleEn string `json:"title_en,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"videopodcast_1.12.2021"`
				CollectionArtistEmptyStateDownloaded struct {
					Style    string `json:"style"`
					Messages []struct {
						Text   string `json:"text"`
						Image  string `json:"image"`
						TextEn string `json:"text_en"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Query string `json:"query"`
							Type  string `json:"type"`
							Name  string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"collection_artist_empty_state_downloaded"`
				Videopodcast24122021Newyeargift struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL             string `json:"url"`
							Title           string `json:"title"`
							ShowsMiniPlayer bool   `json:"shows_mini_player"`
							Name            string `json:"name"`
							InWebkit        bool   `json:"in_webkit"`
						} `json:"action"`
						Style   string `json:"style"`
						Title   string `json:"title"`
						TitleUk string `json:"title_uk,omitempty"`
						TitleEn string `json:"title_en,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"videopodcast_24.12.2021_newyeargift"`
				WavesEmptyArtist struct {
					Style    string `json:"style"`
					Messages []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"waves_empty_artist"`
				SberprimeLogin struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						TextAz     string `json:"text_az"`
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						TitleAz    string `json:"title_az"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk string `json:"title_uk"`
						TextUk  string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Style           string `json:"style"`
						ShouldColorIcon bool   `json:"should_color_icon,omitempty"`
						Title           string `json:"title"`
						TitleAz         string `json:"title_az"`
						TitleEn         string `json:"title_en"`
						TextColor       string `json:"text_color,omitempty"`
						BgColor         string `json:"bg_color,omitempty"`
						Action          struct {
							Register     bool   `json:"register"`
							Name         string `json:"name"`
							LoginTrigger string `json:"login_trigger"`
						} `json:"action"`
						TitleUk string `json:"title_uk"`
						Icon    string `json:"icon,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"sberprime-login"`
				AgreementFromSberButton struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						TextAz     string `json:"text_az"`
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						TitleAz    string `json:"title_az"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk string `json:"title_uk"`
						TextUk  string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Type       string `json:"type"`
						Checkboxes []struct {
							Checked  bool   `json:"checked"`
							TextAz   string `json:"text_az"`
							Text     string `json:"text"`
							Required bool   `json:"required"`
							TextEn   string `json:"text_en"`
							TextUk   string `json:"text_uk"`
							ID       string `json:"id"`
							ErrorAz  string `json:"error_az,omitempty"`
							ErrorUk  string `json:"error_uk,omitempty"`
							Error    string `json:"error,omitempty"`
							ErrorEn  string `json:"error_en,omitempty"`
						} `json:"checkboxes,omitempty"`
						RequiresAccept bool `json:"requires_accept,omitempty"`
						Action         struct {
							Name    string `json:"name"`
							Success struct {
								Fail struct {
									ID   string `json:"id"`
									Name string `json:"name"`
								} `json:"fail"`
								Name    string `json:"name"`
								Success struct {
									ID   string `json:"id"`
									Name string `json:"name"`
								} `json:"success"`
							} `json:"success"`
							Title        string `json:"title"`
							TitleEn      string `json:"title_en"`
							TitleAz      string `json:"title_az"`
							LoginTrigger string `json:"login_trigger"`
							TitleUk      string `json:"title_uk"`
						} `json:"action,omitempty"`
						Style   string `json:"style,omitempty"`
						Title   string `json:"title,omitempty"`
						TitleEn string `json:"title_en,omitempty"`
						TitleAz string `json:"title_az,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"agreement_from_sber_button"`
				DcollectionEmptyStateGridheader struct {
					Comment  string `json:"comment"`
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
							Image string `json:"image"`
						} `json:"background"`
						BrandedBackground struct {
							Gradient      bool   `json:"gradient"`
							Width         int    `json:"width"`
							Height        int    `json:"height"`
							PaletteBottom string `json:"palette_bottom"`
						} `json:"branded_background"`
						TitleUk string `json:"title_uk"`
						TextUk  string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Style     string `json:"style"`
						Title     string `json:"title"`
						TitleEn   string `json:"title_en"`
						TextColor string `json:"text_color"`
						BgColor   string `json:"bg_color"`
						Action    struct {
							Fail struct {
								Name string `json:"name"`
							} `json:"fail"`
							Page    string `json:"page"`
							Success struct {
								Name string `json:"name"`
							} `json:"success"`
							ID string `json:"id"`
						} `json:"action"`
						TitleUk string `json:"title_uk"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"dcollection_empty_state_gridheader"`
				DpodcastsEmptyStateGridheader struct {
					Style    string `json:"style"`
					Messages []struct {
						Text              string `json:"text"`
						BrandedBackground struct {
							Gradient      bool   `json:"gradient"`
							Width         int    `json:"width"`
							Height        int    `json:"height"`
							PaletteBottom string `json:"palette_bottom"`
						} `json:"branded_background"`
						Background struct {
							Color string `json:"color"`
							Image string `json:"image"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL  string `json:"url"`
							Name string `json:"name"`
						} `json:"action"`
						BgColor   string `json:"bg_color"`
						Style     string `json:"style"`
						TextColor string `json:"text_color"`
						Title     string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"dpodcasts_empty_state_gridheader"`
				CollectionAbooksEmptyStateDownloaded struct {
					Style    string `json:"style"`
					Messages []struct {
						Text   string `json:"text"`
						Image  string `json:"image"`
						TextEn string `json:"text_en"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL  string `json:"url"`
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"collection_abooks_empty_state_downloaded"`
				CollectionPlaylistEmptyStateDownloaded struct {
					Style    string `json:"style"`
					Messages []struct {
						Text   string `json:"text"`
						Image  string `json:"image"`
						TextEn string `json:"text_en"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL  string `json:"url"`
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"collection_playlist_empty_state_downloaded"`
				MarketingOneYearSubscription struct {
					Style    string `json:"style"`
					Messages []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name    string `json:"name"`
							Success struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"success"`
							SubscriptionType    string `json:"subscription_type"`
							AndroidSubscription string `json:"android_subscription"`
							Fail                struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"fail"`
							Subscription string `json:"subscription"`
						} `json:"action"`
						Style string `json:"style"`
						Title string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"marketing_one_year_subscription"`
				TestRadio struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Detail     string `json:"detail"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL             string `json:"url"`
							Title           string `json:"title"`
							ShowsMiniPlayer bool   `json:"shows_mini_player"`
							Name            string `json:"name"`
							InWebkit        bool   `json:"in_webkit"`
						} `json:"action"`
						Style   string `json:"style"`
						Title   string `json:"title"`
						TitleUk string `json:"title_uk,omitempty"`
						TitleEn string `json:"title_en,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"test_radio"`
				OfflineModeCollection struct {
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk string `json:"title_uk"`
						TextUk  string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						TitleUk string `json:"title_uk"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"offline_mode_collection"`
				CollectionPlaylistEmptyState struct {
					Style    string `json:"style"`
					Messages []struct {
						Text   string `json:"text"`
						Image  string `json:"image"`
						TextEn string `json:"text_en"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL  string `json:"url"`
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"collection_playlist_empty_state"`
				CollectionAlbumEmptyStateDownloaded struct {
					Style    string `json:"style"`
					Messages []struct {
						Text   string `json:"text"`
						Image  string `json:"image"`
						TextEn string `json:"text_en"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL  string `json:"url"`
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"collection_album_empty_state_downloaded"`
				WavesStartArtist struct {
					Style    string `json:"style"`
					Messages []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TextEn string `json:"text_en"`
						Title  string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						TitleEn string `json:"title_en"`
						Style   string `json:"style"`
						Rounded bool   `json:"rounded,omitempty"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"waves_start_artist"`
				CollectionAllTracksEmptyStateDownloaded struct {
					Style    string `json:"style"`
					Messages []struct {
						Text   string `json:"text"`
						Image  string `json:"image"`
						TextEn string `json:"text_en"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL  string `json:"url"`
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"collection_all_tracks_empty_state_downloaded"`
				SubscriptionExpired struct {
					Style                string `json:"style"`
					CloseGestureDisabled bool   `json:"close_gesture_disabled"`
					Messages             []struct {
						TextAz     string `json:"text_az"`
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						TitleAz    string `json:"title_az"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk string `json:"title_uk"`
						TextUk  string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"subscription_expired"`
				CollectionPlaylistEmptyStateFavourite struct {
					Style    string `json:"style"`
					Messages []struct {
						Text    string `json:"text"`
						Image   string `json:"image"`
						TitleEn string `json:"title_en"`
						TextEn  string `json:"text_en"`
						Title   string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Page string `json:"page"`
							ID   string `json:"id"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"collection_playlist_empty_state_favourite"`
				CollectionProfileEmptyState struct {
					Style    string `json:"style"`
					Messages []struct {
						Text   string `json:"text"`
						Image  string `json:"image"`
						TextEn string `json:"text_en"`
					} `json:"messages"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"collection_profile_empty_state"`
				AgreementBitch struct {
					Style    string `json:"style"`
					Messages []struct {
						TextAz     string `json:"text_az"`
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						TitleAz    string `json:"title_az"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk string `json:"title_uk"`
						TextUk  string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Style   string `json:"style"`
						Title   string `json:"title"`
						TitleAz string `json:"title_az"`
						TitleEn string `json:"title_en"`
						Action  struct {
							Name                  string `json:"name"`
							DismissLoginActionkit int    `json:"dismiss_login_actionkit"`
						} `json:"action"`
						TitleUk string `json:"title_uk"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"agreement_bitch"`
				NoInternet struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk string `json:"title_uk"`
						TextUk  string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style string `json:"style"`
						Title string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"no_internet"`
				TestVideo struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Detail     string `json:"detail"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL             string `json:"url"`
							Title           string `json:"title"`
							ShowsMiniPlayer bool   `json:"shows_mini_player"`
							Name            string `json:"name"`
							InWebkit        bool   `json:"in_webkit"`
						} `json:"action"`
						Style   string `json:"style"`
						Title   string `json:"title"`
						TitleUk string `json:"title_uk,omitempty"`
						TitleEn string `json:"title_en,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"test_video"`
				NotEnoughStorage struct {
					Comment  string `json:"comment"`
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style     string `json:"style"`
						TextColor string `json:"text_color,omitempty"`
						TitleEn   string `json:"title_en"`
						Title     string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"not_enough_storage"`
				PlusSubscriptionChangeYear struct {
					Style    string `json:"style"`
					Messages []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Fail struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"fail"`
							AndroidSubscription string `json:"android_subscription"`
							Name                string `json:"name"`
							Success             struct {
								ID   string `json:"id"`
								Name string `json:"name"`
							} `json:"success"`
							Subscription string `json:"subscription"`
						} `json:"action"`
						Style string `json:"style"`
						Title string `json:"title"`
					} `json:"actions"`
					Source           string      `json:"source"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"plus_subscription_change_year"`
				AppleGenericAnnualSubscriptionFailure struct {
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
							ID   string `json:"id"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_annual_subscription_failure"`
				FirstStartShufflePaywall struct {
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						Detail     string `json:"detail"`
						TextEn     string `json:"text_en"`
						DetailUk   string `json:"detail_uk"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk  string `json:"title_uk"`
						DetailEn string `json:"detail_en"`
						TextUk   string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
							ID   string `json:"id"`
						} `json:"action"`
						TitleUk string `json:"title_uk"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string      `json:"source"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"first_start_shuffle_paywall"`
				AppleGenericUpsellAnnualAgain struct {
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name    string `json:"name"`
							Success struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"success"`
							SubscriptionType    string `json:"subscription_type"`
							AndroidSubscription string `json:"android_subscription"`
							Fail                struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"fail"`
							Subscription string `json:"subscription"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_upsell_annual_again"`
				Videolink23122021GAYAZOVSBROTHERS struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL             string `json:"url"`
							Title           string `json:"title"`
							ShowsMiniPlayer bool   `json:"shows_mini_player"`
							Name            string `json:"name"`
							InWebkit        bool   `json:"in_webkit"`
						} `json:"action"`
						Style   string `json:"style"`
						Title   string `json:"title"`
						TitleUk string `json:"title_uk,omitempty"`
						TitleEn string `json:"title_en,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"videolink_23.12.2021_GAYAZOVS-BROTHERS"`
				WavesFirstStart struct {
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						TitleEn string `json:"title_en"`
						Style   string `json:"style"`
						Rounded bool   `json:"rounded,omitempty"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"waves_first_start"`
				AppleGenericUpsellAnnualCancel struct {
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Style     string `json:"style"`
						Title     string `json:"title"`
						TitleEn   string `json:"title_en"`
						TextColor string `json:"text_color,omitempty"`
						BgColor   string `json:"bg_color,omitempty"`
						Action    struct {
							Name    string `json:"name"`
							Success struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"success"`
							SubscriptionType    string `json:"subscription_type"`
							AndroidSubscription string `json:"android_subscription"`
							Fail                struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"fail"`
							Subscription string `json:"subscription"`
						} `json:"action"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_upsell_annual_cancel"`
				CollectionAlbumEmptyState struct {
					Style    string `json:"style"`
					Messages []struct {
						Text   string `json:"text"`
						Image  string `json:"image"`
						TextEn string `json:"text_en"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL  string `json:"url"`
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"collection_album_empty_state"`
				PlusSubscriptionSuccess struct {
					Style    string `json:"style"`
					Messages []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style string `json:"style"`
						Title string `json:"title"`
					} `json:"actions"`
					Source           string      `json:"source"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"plus_subscription_success"`
				DartistsEmptyStateGridheader struct {
					Comment  string `json:"comment"`
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
							Image string `json:"image"`
						} `json:"background"`
						BrandedBackground struct {
							Gradient      bool   `json:"gradient"`
							Width         int    `json:"width"`
							Height        int    `json:"height"`
							PaletteBottom string `json:"palette_bottom"`
						} `json:"branded_background"`
					} `json:"messages"`
					Actions []struct {
						Style     string `json:"style"`
						Title     string `json:"title"`
						TitleEn   string `json:"title_en"`
						TextColor string `json:"text_color"`
						BgColor   string `json:"bg_color"`
						Action    struct {
							Query string `json:"query"`
							Name  string `json:"name"`
						} `json:"action"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"dartists_empty_state_gridheader"`
				ArtistsEmptyStateGridheader struct {
					Comment  string `json:"comment"`
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
							Image string `json:"image"`
						} `json:"background"`
						BrandedBackground struct {
							Gradient      bool   `json:"gradient"`
							Width         int    `json:"width"`
							Height        int    `json:"height"`
							PaletteBottom string `json:"palette_bottom"`
						} `json:"branded_background"`
					} `json:"messages"`
					Actions []struct {
						Style     string `json:"style"`
						Title     string `json:"title"`
						TitleEn   string `json:"title_en"`
						TextColor string `json:"text_color"`
						BgColor   string `json:"bg_color"`
						Action    struct {
							Query string `json:"query"`
							Name  string `json:"name"`
						} `json:"action"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"artists_empty_state_gridheader"`
				ExplicitBlock struct {
					Style    string `json:"style"`
					Messages []struct {
						Text       string `json:"text"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style string `json:"style"`
						Title string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"explicit_block"`
				StoriesError struct {
					Style    string `json:"style"`
					Messages []struct {
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style string `json:"style"`
						Title string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"stories_error"`
				PodcastsEmptyStateGridheader struct {
					Style    string `json:"style"`
					Messages []struct {
						Text              string `json:"text"`
						BrandedBackground struct {
							Gradient      bool   `json:"gradient"`
							Width         int    `json:"width"`
							Height        int    `json:"height"`
							PaletteBottom string `json:"palette_bottom"`
						} `json:"branded_background"`
						Background struct {
							Color string `json:"color"`
							Image string `json:"image"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL  string `json:"url"`
							Name string `json:"name"`
						} `json:"action"`
						BgColor   string `json:"bg_color"`
						Style     string `json:"style"`
						TextColor string `json:"text_color"`
						Title     string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"podcasts_empty_state_gridheader"`
				AppleGenericMultitasking struct {
					Style                string `json:"style"`
					CloseGestureDisabled bool   `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Detail     string `json:"detail"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Agreement string `json:"agreement"`
					Actions   []struct {
						Action struct {
							Fail struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"fail"`
							AndroidSubscription string `json:"android_subscription"`
							Name                string `json:"name"`
							Success             struct {
								ID   string `json:"id"`
								Name string `json:"name"`
							} `json:"success"`
							Subscription string `json:"subscription"`
						} `json:"action"`
						Style           string `json:"style"`
						BackgroundImage string `json:"background_image,omitempty"`
						Title           string `json:"title"`
						TextColor       string `json:"text_color,omitempty"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_multitasking"`
				AppleGenericPayViaCard struct {
					Style    string `json:"style"`
					Messages []struct {
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL     string `json:"url"`
							Success struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"success"`
							Name string `json:"name"`
							Auth bool   `json:"auth"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_pay_via_card"`
				ProfileNoPlaylistsCurrentUserEmptyState struct {
					Style    string `json:"style"`
					Messages []struct {
						Text   string `json:"text"`
						Image  string `json:"image"`
						TextEn string `json:"text_en"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"profile_no_playlists_current_user_empty_state"`
				Videolink22122021MoominTroll struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL             string `json:"url"`
							Title           string `json:"title"`
							ShowsMiniPlayer bool   `json:"shows_mini_player"`
							Name            string `json:"name"`
							InWebkit        bool   `json:"in_webkit"`
						} `json:"action"`
						Style   string `json:"style"`
						Title   string `json:"title"`
						TitleUk string `json:"title_uk,omitempty"`
						TitleEn string `json:"title_en,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"videolink_22.12.2021_moomin-troll"`
				AppleGenericPushProhibited struct {
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk string `json:"title_uk"`
						TextUk  string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						TitleUk string `json:"title_uk"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
						TitleAz string `json:"title_az,omitempty"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_push_prohibited"`
				WavesStartTrack struct {
					Style    string `json:"style"`
					Messages []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TextEn string `json:"text_en"`
						Title  string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						TitleEn string `json:"title_en"`
						Style   string `json:"style"`
						Rounded bool   `json:"rounded,omitempty"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"waves_start_track"`
				CollectionArtistEmptyState struct {
					Style    string `json:"style"`
					Messages []struct {
						Text   string `json:"text"`
						Image  string `json:"image"`
						TextEn string `json:"text_en"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Query string `json:"query"`
							Type  string `json:"type"`
							Name  string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"collection_artist_empty_state"`
				ProfileEmptyStateGridheader struct {
					Style    string `json:"style"`
					Messages []struct {
						Title string `json:"title"`
					} `json:"messages"`
					Actions          []interface{} `json:"actions"`
					Source           string        `json:"source"`
					PinActions       bool          `json:"pin_actions"`
					AllowCloseButton bool          `json:"allow_close_button"`
					Type             string        `json:"type"`
				} `json:"profile_empty_state_gridheader"`
				TestPromocodeWebview struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL      string `json:"url"`
							Name     string `json:"name"`
							Auth     bool   `json:"auth"`
							InWebkit bool   `json:"in_webkit"`
						} `json:"action"`
						Style   string `json:"style"`
						Title   string `json:"title"`
						TitleUk string `json:"title_uk,omitempty"`
						TitleEn string `json:"title_en,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"test_promocode_webview"`
				AppleGenericUpsellAnnualExplanation struct {
					Comment  string `json:"comment"`
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Style     string `json:"style"`
						Title     string `json:"title"`
						TitleEn   string `json:"title_en"`
						TextColor string `json:"text_color,omitempty"`
						BgColor   string `json:"bg_color,omitempty"`
						Action    struct {
							Name    string `json:"name"`
							Success struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"success"`
							SubscriptionType    string `json:"subscription_type"`
							AndroidSubscription string `json:"android_subscription"`
							Fail                struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"fail"`
							Subscription string `json:"subscription"`
						} `json:"action"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_upsell_annual_explanation"`
				DwavesEmptyStateGridheader struct {
					Comment  string `json:"comment"`
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
							Image string `json:"image"`
						} `json:"background"`
						BrandedBackground struct {
							Gradient      bool   `json:"gradient"`
							Width         int    `json:"width"`
							Height        int    `json:"height"`
							PaletteBottom string `json:"palette_bottom"`
						} `json:"branded_background"`
					} `json:"messages"`
					Actions []struct {
						Style     string `json:"style"`
						Title     string `json:"title"`
						TitleEn   string `json:"title_en"`
						TextColor string `json:"text_color"`
						BgColor   string `json:"bg_color"`
						Action    struct {
							URL  string `json:"url"`
							Fail struct {
								Name string `json:"name"`
							} `json:"fail"`
							Name    string `json:"name"`
							Success struct {
								Name string `json:"name"`
							} `json:"success"`
						} `json:"action"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"dwaves_empty_state_gridheader"`
				RegisteredCachePaywall struct {
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						Detail     string `json:"detail"`
						TextEn     string `json:"text_en"`
						DetailUk   string `json:"detail_uk"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk  string `json:"title_uk"`
						DetailEn string `json:"detail_en"`
						TextUk   string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
							ID   string `json:"id"`
						} `json:"action"`
						TitleUk string `json:"title_uk"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"registered_cache_paywall"`
				AttachSberError struct {
					Style    string `json:"style"`
					Messages []struct {
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Type string `json:"type"`
							Name string `json:"name"`
						} `json:"action"`
						Style string `json:"style"`
						Title string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"attach_sber_error"`
				AlbumsEmptyStateGridheader struct {
					Comment  string `json:"comment"`
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
							Image string `json:"image"`
						} `json:"background"`
						BrandedBackground struct {
							Gradient      bool   `json:"gradient"`
							Width         int    `json:"width"`
							Height        int    `json:"height"`
							PaletteBottom string `json:"palette_bottom"`
						} `json:"branded_background"`
					} `json:"messages"`
					Actions []struct {
						Style     string `json:"style"`
						Title     string `json:"title"`
						TitleEn   string `json:"title_en"`
						TextColor string `json:"text_color"`
						BgColor   string `json:"bg_color"`
						Action    struct {
							URL  string `json:"url"`
							Fail struct {
								Name string `json:"name"`
							} `json:"fail"`
							Name    string `json:"name"`
							Success struct {
								Name string `json:"name"`
							} `json:"success"`
						} `json:"action"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"albums_empty_state_gridheader"`
				AttachAccountSber struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Style     string `json:"style"`
						Title     string `json:"title"`
						TitleEn   string `json:"title_en,omitempty"`
						TextColor string `json:"text_color"`
						BgColor   string `json:"bg_color"`
						Action    struct {
							Fail struct {
								ID   string `json:"id"`
								Name string `json:"name"`
							} `json:"fail"`
							Type string `json:"type"`
							Name string `json:"name"`
						} `json:"action"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"attach_account_sber"`
				ZeroTestAkit struct {
					Comment              string `json:"comment"`
					Style                string `json:"style"`
					CloseGestureDisabled bool   `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text,omitempty"`
						Image      string `json:"image"`
						Detail     string `json:"detail,omitempty"`
						Background struct {
							Color string `json:"color"`
							Image string `json:"image"`
						} `json:"background"`
						Title   string `json:"title"`
						Bullets []struct {
							Text string `json:"text"`
						} `json:"bullets,omitempty"`
					} `json:"messages"`
					Agreement string `json:"agreement"`
					Actions   []struct {
						Style    string `json:"style"`
						Subtitle string `json:"subtitle,omitempty"`
						Title    string `json:"title"`
						Detail   string `json:"detail,omitempty"`
						Action   struct {
							Fail struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"fail"`
							AndroidSubscription string `json:"android_subscription"`
							Name                string `json:"name"`
							Success             struct {
								ID   string `json:"id"`
								Name string `json:"name"`
							} `json:"success"`
							Subscription string `json:"subscription"`
						} `json:"action"`
						Icon            string `json:"icon"`
						ShouldColorIcon bool   `json:"should_color_icon,omitempty"`
						TitleEn         string `json:"title_en,omitempty"`
						TextColor       string `json:"text_color,omitempty"`
						BgColor         string `json:"bg_color,omitempty"`
					} `json:"actions"`
					MultipleBannersTitle    string `json:"multiple_banners_title"`
					Source                  string `json:"source"`
					PinActions              bool   `json:"pin_actions"`
					MultipleBannersSubtitle string `json:"multiple_banners_subtitle"`
					AllowCloseButton        bool   `json:"allow_close_button"`
					Type                    string `json:"type"`
				} `json:"0test-akit"`
				OfflineModeDownload struct {
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk string `json:"title_uk"`
						TextUk  string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						TitleUk string `json:"title_uk"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"offline_mode_download"`
				AppleGenericSubscriptionSuccess struct {
					Comment  string `json:"comment"`
					Messages []struct {
						TextAz     string `json:"text_az"`
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						TitleAz    string `json:"title_az"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk string `json:"title_uk"`
						TextUk  string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"apple_generic_subscription_success"`
				BeelineKgSubscriptionFailure struct {
					Messages []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style string `json:"style"`
						Title string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"beeline_kg_subscription_failure"`
				Videopodcast24192021 struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL             string `json:"url"`
							Title           string `json:"title"`
							ShowsMiniPlayer bool   `json:"shows_mini_player"`
							Name            string `json:"name"`
							InWebkit        bool   `json:"in_webkit"`
						} `json:"action"`
						Style   string `json:"style"`
						Title   string `json:"title"`
						TitleUk string `json:"title_uk,omitempty"`
						TitleEn string `json:"title_en,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"videopodcast_24.19.2021"`
				CollectionAbooksEmptyState struct {
					Style    string `json:"style"`
					Messages []struct {
						Text   string `json:"text"`
						Image  string `json:"image"`
						TextEn string `json:"text_en"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL  string `json:"url"`
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"collection_abooks_empty_state"`
				TrackPremiumOnly struct {
					Comment  string `json:"comment"`
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						Detail     string `json:"detail"`
						TextEn     string `json:"text_en"`
						DetailUk   string `json:"detail_uk"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk  string `json:"title_uk"`
						DetailEn string `json:"detail_en"`
						TextUk   string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Fail struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"fail"`
							AndroidSubscription string `json:"android_subscription"`
							Name                string `json:"name"`
							Success             struct {
								ID   string `json:"id"`
								Name string `json:"name"`
							} `json:"success"`
							Subscription string `json:"subscription"`
						} `json:"action"`
						TitleUk string `json:"title_uk"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"track_premium_only"`
				AppleGenericSubscriptionSuccessSberprimeSub struct {
					Comment              string `json:"comment"`
					CloseGestureDisabled bool   `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Detail     string `json:"detail"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						BgColor   string `json:"bg_color"`
						Style     string `json:"style"`
						TextColor string `json:"text_color"`
						Title     string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"apple_generic_subscription_success_sberprime_sub"`
				Videopodcast3112021 struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL             string `json:"url"`
							Title           string `json:"title"`
							ShowsMiniPlayer bool   `json:"shows_mini_player"`
							Name            string `json:"name"`
							InWebkit        bool   `json:"in_webkit"`
						} `json:"action"`
						Style   string `json:"style"`
						Title   string `json:"title"`
						TitleUk string `json:"title_uk,omitempty"`
						TitleEn string `json:"title_en,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"videopodcast_3.11.2021"`
				CollectionAllTracksEmptyState struct {
					Style    string `json:"style"`
					Messages []struct {
						Text   string `json:"text"`
						Image  string `json:"image"`
						TextEn string `json:"text_en"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL  string `json:"url"`
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"collection_all_tracks_empty_state"`
				GiftCodeWebview struct {
					Style    string `json:"style"`
					Messages []struct {
						WebContentURL string `json:"web_content_url"`
					} `json:"messages"`
					Actions          []interface{} `json:"actions"`
					Source           string        `json:"source"`
					PinActions       bool          `json:"pin_actions"`
					AllowCloseButton bool          `json:"allow_close_button"`
					Type             string        `json:"type"`
				} `json:"gift_code_webview"`
				AnnualSubscription struct {
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name    string `json:"name"`
							Success struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"success"`
							SubscriptionType    string `json:"subscription_type"`
							AndroidSubscription string `json:"android_subscription"`
							Fail                struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"fail"`
							Subscription string `json:"subscription"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"annual_subscription"`
				PermissionMicrophone struct {
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Agreement string `json:"agreement"`
					Actions   []struct {
						Action struct {
							Name    string `json:"name"`
							Success struct {
								Name string `json:"name"`
							} `json:"success"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"permission-microphone"`
				AppleGenericSubscriptionFailure struct {
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
							ID   string `json:"id"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_subscription_failure"`
				AppleGenericSubscriptionFailureSberprimeSub struct {
					Comment              string `json:"comment"`
					CloseGestureDisabled bool   `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Detail     string `json:"detail"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
							ID   string `json:"id"`
						} `json:"action"`
						BgColor   string `json:"bg_color"`
						Style     string `json:"style"`
						TextColor string `json:"text_color"`
						Title     string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"apple_generic_subscription_failure_sberprime_sub"`
				CollectionPlaylistEmptyStateFavouriteDownloaded struct {
					Style    string `json:"style"`
					Messages []struct {
						Text    string `json:"text"`
						Image   string `json:"image"`
						TitleEn string `json:"title_en"`
						TextEn  string `json:"text_en"`
						Title   string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Page string `json:"page"`
							ID   string `json:"id"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"collection_playlist_empty_state_favourite_downloaded"`
				RestoreSubscriptionSuccess struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style string `json:"style"`
						Title string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"restore_subscription_success"`
				Videopodcast17112021 struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL             string `json:"url"`
							Title           string `json:"title"`
							ShowsMiniPlayer bool   `json:"shows_mini_player"`
							Name            string `json:"name"`
							InWebkit        bool   `json:"in_webkit"`
						} `json:"action"`
						Style   string `json:"style"`
						Title   string `json:"title"`
						TitleUk string `json:"title_uk,omitempty"`
						TitleEn string `json:"title_en,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"videopodcast_17.11.2021"`
				DisableOfflineMode struct {
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"disable_offline_mode"`
				OfflineModeRadio struct {
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk string `json:"title_uk"`
						TextUk  string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						TitleUk string `json:"title_uk"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"offline_mode_radio"`
				WavesEmptyTrack struct {
					Style    string `json:"style"`
					Messages []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"waves_empty_track"`
				LoginSuccess struct {
					Messages []struct {
						TextAz     string `json:"text_az"`
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleAz    string `json:"title_az"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleEn string `json:"title_en"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style string `json:"style"`
						Title string `json:"title"`
					} `json:"actions"`
					Source           string      `json:"source"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"login_success"`
				FavouriteTracksEmptyStateGridheader struct {
					Comment  string `json:"comment"`
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
							Image string `json:"image"`
						} `json:"background"`
						BrandedBackground struct {
							Gradient      bool   `json:"gradient"`
							Width         int    `json:"width"`
							Height        int    `json:"height"`
							PaletteBottom string `json:"palette_bottom"`
						} `json:"branded_background"`
					} `json:"messages"`
					Actions []struct {
						Style     string `json:"style"`
						Title     string `json:"title"`
						TitleEn   string `json:"title_en"`
						TextColor string `json:"text_color"`
						BgColor   string `json:"bg_color"`
						Action    struct {
							Fail struct {
								Name string `json:"name"`
							} `json:"fail"`
							Page    string `json:"page"`
							Success struct {
								Name string `json:"name"`
							} `json:"success"`
							ID string `json:"id"`
						} `json:"action"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"favourite_tracks_empty_state_gridheader"`
				DfavouriteTracksEmptyStateGridheader struct {
					Comment  string `json:"comment"`
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
							Image string `json:"image"`
						} `json:"background"`
						BrandedBackground struct {
							Gradient      bool   `json:"gradient"`
							Width         int    `json:"width"`
							Height        int    `json:"height"`
							PaletteBottom string `json:"palette_bottom"`
						} `json:"branded_background"`
					} `json:"messages"`
					Actions []struct {
						Style     string `json:"style"`
						Title     string `json:"title"`
						TitleEn   string `json:"title_en"`
						TextColor string `json:"text_color"`
						BgColor   string `json:"bg_color"`
						Action    struct {
							Fail struct {
								Name string `json:"name"`
							} `json:"fail"`
							Page    string `json:"page"`
							Success struct {
								Name string `json:"name"`
							} `json:"success"`
							ID string `json:"id"`
						} `json:"action"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"dfavourite_tracks_empty_state_gridheader"`
				OfflineModeLyrics struct {
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk string `json:"title_uk"`
						TextUk  string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						TitleUk string `json:"title_uk"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"offline_mode_lyrics"`
				TestBeelineKGKit struct {
					Messages []struct {
						Image   string `json:"image"`
						Bullets []struct {
							Text  string `json:"text"`
							Image string `json:"image"`
						} `json:"bullets"`
						Detail     string `json:"detail"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Fail struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"fail"`
							AndroidSubscription string `json:"android_subscription"`
							Name                string `json:"name"`
							Success             struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"success"`
							Subscription string `json:"subscription"`
						} `json:"action"`
						Style string `json:"style"`
						Title string `json:"title"`
					} `json:"actions"`
					PinActions bool   `json:"pin_actions"`
					Type       string `json:"type"`
				} `json:"test_beeline_KG_kit"`
				CollectionPodcastEmptyState struct {
					Style    string `json:"style"`
					Messages []struct {
						Text   string `json:"text"`
						Image  string `json:"image"`
						TextEn string `json:"text_en"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL  string `json:"url"`
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"collection_podcast_empty_state"`
				TrackUnavailable struct {
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk string `json:"title_uk"`
						TextUk  string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
						} `json:"action"`
						TitleUk string `json:"title_uk"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"track_unavailable"`
				PremiumOfflineFirstTime struct {
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style string `json:"style"`
						Title string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"premium_offline_first_time"`
				WavesStartRelease struct {
					Style    string `json:"style"`
					Messages []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TextEn string `json:"text_en"`
						Title  string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						TitleEn string `json:"title_en"`
						Style   string `json:"style"`
						Rounded bool   `json:"rounded,omitempty"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"waves_start_release"`
				Videolink24122021EtoyamoguInclusion struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL             string `json:"url"`
							Title           string `json:"title"`
							ShowsMiniPlayer bool   `json:"shows_mini_player"`
							Name            string `json:"name"`
							InWebkit        bool   `json:"in_webkit"`
						} `json:"action"`
						Style   string `json:"style"`
						Title   string `json:"title"`
						TitleUk string `json:"title_uk,omitempty"`
						TitleEn string `json:"title_en,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"videolink_24.12.2021_etoyamogu-inclusion"`
				AudiobooksEmptyStateGridheader struct {
					Style    string `json:"style"`
					Messages []struct {
						Text              string `json:"text"`
						BrandedBackground struct {
							Gradient      bool   `json:"gradient"`
							Width         int    `json:"width"`
							Height        int    `json:"height"`
							PaletteBottom string `json:"palette_bottom"`
						} `json:"branded_background"`
						Background struct {
							Color string `json:"color"`
							Image string `json:"image"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL  string `json:"url"`
							Name string `json:"name"`
						} `json:"action"`
						BgColor   string `json:"bg_color"`
						Style     string `json:"style"`
						TextColor string `json:"text_color"`
						Title     string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"audiobooks_empty_state_gridheader"`
				DalbumsEmptyStateGridheader struct {
					Comment  string `json:"comment"`
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
							Image string `json:"image"`
						} `json:"background"`
						BrandedBackground struct {
							Gradient      bool   `json:"gradient"`
							Width         int    `json:"width"`
							Height        int    `json:"height"`
							PaletteBottom string `json:"palette_bottom"`
						} `json:"branded_background"`
					} `json:"messages"`
					Actions []struct {
						Style     string `json:"style"`
						Title     string `json:"title"`
						TitleEn   string `json:"title_en"`
						TextColor string `json:"text_color"`
						BgColor   string `json:"bg_color"`
						Action    struct {
							URL  string `json:"url"`
							Fail struct {
								Name string `json:"name"`
							} `json:"fail"`
							Name    string `json:"name"`
							Success struct {
								Name string `json:"name"`
							} `json:"success"`
						} `json:"action"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"dalbums_empty_state_gridheader"`
				AppleGenericUpsellTryInProfile struct {
					Style    string `json:"style"`
					Messages []struct {
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_upsell_try_in_profile"`
				TriggerAction10182021Premium struct {
					Comment  string `json:"comment"`
					Messages []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
							ID   string `json:"id"`
						} `json:"action"`
						Style   string `json:"style"`
						Title   string `json:"title"`
						TitleUk string `json:"title_uk,omitempty"`
						TitleEn string `json:"title_en,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"trigger_action_10.18.2021_premium"`
				DaudiobooksEmptyStateGridheader struct {
					Style    string `json:"style"`
					Messages []struct {
						Text              string `json:"text"`
						BrandedBackground struct {
							Gradient      bool   `json:"gradient"`
							Width         int    `json:"width"`
							Height        int    `json:"height"`
							PaletteBottom string `json:"palette_bottom"`
						} `json:"branded_background"`
						Background struct {
							Color string `json:"color"`
							Image string `json:"image"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL  string `json:"url"`
							Name string `json:"name"`
						} `json:"action"`
						BgColor   string `json:"bg_color"`
						Style     string `json:"style"`
						TextColor string `json:"text_color"`
						Title     string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"daudiobooks_empty_state_gridheader"`
				OfflineModeStreaming struct {
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk string `json:"title_uk"`
						TextUk  string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						TitleUk string `json:"title_uk"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"offline_mode_streaming"`
				Videolink22122021DmitrienkoHabib struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL             string `json:"url"`
							Title           string `json:"title"`
							ShowsMiniPlayer bool   `json:"shows_mini_player"`
							Name            string `json:"name"`
							InWebkit        bool   `json:"in_webkit"`
						} `json:"action"`
						Style   string `json:"style"`
						Title   string `json:"title"`
						TitleUk string `json:"title_uk,omitempty"`
						TitleEn string `json:"title_en,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"videolink_22.12.2021_dmitrienko-habib"`
				LotsOfSubscriptions struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Detail     string `json:"detail"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Section string `json:"section"`
							Name    string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						Title   string `json:"title"`
						TitleUk string `json:"title_uk,omitempty"`
						TitleEn string `json:"title_en,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"lots_of_subscriptions"`
				AirplaneMode struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk string `json:"title_uk"`
						TextUk  string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Style           string `json:"style"`
						ShouldColorIcon bool   `json:"should_color_icon,omitempty"`
						Title           string `json:"title"`
						TitleEn         string `json:"title_en,omitempty"`
						TextColor       string `json:"text_color,omitempty"`
						BgColor         string `json:"bg_color,omitempty"`
						Action          struct {
							Name string `json:"name"`
							ID   string `json:"id"`
						} `json:"action"`
						TitleUk string `json:"title_uk,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"airplane_mode"`
				SbermobilePush15122021 struct {
					Comment              string `json:"comment"`
					CloseGestureDisabled bool   `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL      string `json:"url"`
							Name     string `json:"name"`
							InWebkit bool   `json:"in_webkit"`
						} `json:"action"`
						Style string `json:"style"`
						Title string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"Sbermobile_Push_15.12.2021"`
				SberInsuranceSuccess struct {
					Comment              string `json:"comment"`
					CloseGestureDisabled bool   `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Detail     string `json:"detail"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Agreement string `json:"agreement"`
					Actions   []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						Title   string `json:"title"`
						TitleUk string `json:"title_uk,omitempty"`
						TitleEn string `json:"title_en,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"sber-insurance_success"`
				SignoutPaywall struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"signout_paywall"`
				PlusSubscriptionChangeMonth struct {
					Style    string `json:"style"`
					Messages []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Fail struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"fail"`
							AndroidSubscription string `json:"android_subscription"`
							Name                string `json:"name"`
							Success             struct {
								ID   string `json:"id"`
								Name string `json:"name"`
							} `json:"success"`
							Subscription string `json:"subscription"`
						} `json:"action"`
						Style string `json:"style"`
						Title string `json:"title"`
					} `json:"actions"`
					Source           string      `json:"source"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"plus_subscription_change_month"`
				SberInsurance struct {
					Comment              string `json:"comment"`
					CloseGestureDisabled bool   `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Detail     string `json:"detail"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Agreement string `json:"agreement"`
					Actions   []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						Title   string `json:"title"`
						TitleUk string `json:"title_uk,omitempty"`
						TitleEn string `json:"title_en,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"sber-insurance"`
				AppleGenericUpsellCardPay0 struct {
					Style    string `json:"style"`
					Messages []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL     string `json:"url"`
							Success struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"success"`
							Name string `json:"name"`
							Auth bool   `json:"auth"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_upsell_card_pay0"`
				SberprimeSub struct {
					Comment              string `json:"comment"`
					Style                string `json:"style"`
					CloseGestureDisabled bool   `json:"close_gesture_disabled"`
					Messages             []struct {
						Image   string `json:"image"`
						Bullets []struct {
							Text  string `json:"text"`
							Image string `json:"image"`
						} `json:"bullets"`
						Detail     string `json:"detail"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Type       string `json:"type,omitempty"`
						Checkboxes []struct {
							Text     string `json:"text"`
							Required bool   `json:"required"`
							Checked  bool   `json:"checked"`
							ID       string `json:"id"`
							Error    string `json:"error"`
						} `json:"checkboxes,omitempty"`
						RequiresAccept bool   `json:"requires_accept,omitempty"`
						Style          string `json:"style,omitempty"`
						Subtitle       string `json:"subtitle,omitempty"`
						Title          string `json:"title,omitempty"`
						TextColor      string `json:"text_color,omitempty"`
						BgColor        string `json:"bg_color,omitempty"`
						Action         struct {
							Fail struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"fail"`
							AndroidSubscription string `json:"android_subscription"`
							Name                string `json:"name"`
							Success             struct {
								ID   string `json:"id"`
								Name string `json:"name"`
							} `json:"success"`
							Subscription string `json:"subscription"`
						} `json:"action,omitempty"`
					} `json:"actions"`
					Source           string      `json:"source"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"Sberprime_Sub"`
				AppleGenericAnnualTryInWeb struct {
					Style    string `json:"style"`
					Messages []struct {
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_annual_try_in_web"`
				WavesEmptyRelease struct {
					Style    string `json:"style"`
					Messages []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"waves_empty_release"`
				CollectionPodcastEmptyStateDownloaded struct {
					Style    string `json:"style"`
					Messages []struct {
						Text   string `json:"text"`
						Image  string `json:"image"`
						TextEn string `json:"text_en"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL  string `json:"url"`
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"collection_podcast_empty_state_downloaded"`
				AppleGenericUpsellDismiss149 struct {
					Style    string `json:"style"`
					Messages []struct {
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
							ID   string `json:"id"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_upsell_dismiss149"`
				AppleGenericUpsellWhotAmiPaying0 struct {
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Fail struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"fail"`
							AndroidSubscription string `json:"android_subscription"`
							Name                string `json:"name"`
							Success             struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"success"`
							Subscription string `json:"subscription"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_upsell_whot_ami_paying0"`
				AppleGenericUpsellErrorConfuse struct {
					Style    string `json:"style"`
					Messages []struct {
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
							ID   string `json:"id"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_upsell_error_confuse"`
				DallTracksEmptyStateGridheader struct {
					Comment  string `json:"comment"`
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
							Image string `json:"image"`
						} `json:"background"`
						BrandedBackground struct {
							Gradient      bool   `json:"gradient"`
							Width         int    `json:"width"`
							Height        int    `json:"height"`
							PaletteBottom string `json:"palette_bottom"`
						} `json:"branded_background"`
					} `json:"messages"`
					Actions []struct {
						Style     string `json:"style"`
						Title     string `json:"title"`
						TitleEn   string `json:"title_en"`
						TextColor string `json:"text_color"`
						BgColor   string `json:"bg_color"`
						Action    struct {
							URL  string `json:"url"`
							Fail struct {
								Name string `json:"name"`
							} `json:"fail"`
							Name    string `json:"name"`
							Success struct {
								Name string `json:"name"`
							} `json:"success"`
						} `json:"action"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"dall_tracks_empty_state_gridheader"`
				Videopodcast8102021 struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL             string `json:"url"`
							Title           string `json:"title"`
							ShowsMiniPlayer bool   `json:"shows_mini_player"`
							Name            string `json:"name"`
							InWebkit        bool   `json:"in_webkit"`
						} `json:"action"`
						Style   string `json:"style"`
						Title   string `json:"title"`
						TitleUk string `json:"title_uk,omitempty"`
						TitleEn string `json:"title_en,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"videopodcast_8.10.2021"`
				PrivatePlaylistBlock struct {
					Style    string `json:"style"`
					Messages []struct {
						Text   string `json:"text"`
						Image  string `json:"image"`
						TextEn string `json:"text_en"`
					} `json:"messages"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"private_playlist_block"`
				RestoreSubscriptionFail struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk string `json:"title_uk"`
						TextUk  string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Fail struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"fail"`
							AndroidSubscription string `json:"android_subscription"`
							Name                string `json:"name"`
							Success             struct {
								ID   string `json:"id"`
								Name string `json:"name"`
							} `json:"success"`
							Subscription string `json:"subscription"`
						} `json:"action"`
						TitleUk string `json:"title_uk"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"restore_subscription_fail"`
				ShowAllMusicOffer struct {
					Style    string `json:"style"`
					Messages []struct {
						TextAz     string `json:"text_az"`
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						TitleAz    string `json:"title_az"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Style     string `json:"style"`
						Title     string `json:"title"`
						TitleEn   string `json:"title_en"`
						TitleAz   string `json:"title_az"`
						TextColor string `json:"text_color"`
						BgColor   string `json:"bg_color"`
						Action    struct {
							Name string `json:"name"`
						} `json:"action"`
					} `json:"actions"`
					Source           string      `json:"source"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"show_all_music_offer"`
				TestStory struct {
					CloseGestureDisabled bool `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name         string `json:"name"`
							StoryBlockID string `json:"story_block_id"`
							ID           string `json:"id"`
						} `json:"action"`
						Style   string `json:"style"`
						Title   string `json:"title"`
						TitleUk string `json:"title_uk,omitempty"`
						TitleEn string `json:"title_en,omitempty"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"test_story"`
				AllTracksEmptyStateGridheader struct {
					Comment  string `json:"comment"`
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
							Image string `json:"image"`
						} `json:"background"`
						BrandedBackground struct {
							Gradient      bool   `json:"gradient"`
							Width         int    `json:"width"`
							Height        int    `json:"height"`
							PaletteBottom string `json:"palette_bottom"`
						} `json:"branded_background"`
					} `json:"messages"`
					Actions []struct {
						Style     string `json:"style"`
						Title     string `json:"title"`
						TitleEn   string `json:"title_en"`
						TextColor string `json:"text_color"`
						BgColor   string `json:"bg_color"`
						Action    struct {
							URL  string `json:"url"`
							Fail struct {
								Name string `json:"name"`
							} `json:"fail"`
							Name    string `json:"name"`
							Success struct {
								Name string `json:"name"`
							} `json:"success"`
						} `json:"action"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"all_tracks_empty_state_gridheader"`
				AppleGenericUpsellFree struct {
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Fail struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"fail"`
							AndroidSubscription string `json:"android_subscription"`
							Name                string `json:"name"`
							Success             struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"success"`
							Subscription string `json:"subscription"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_upsell_free"`
				ProfileNoPlaylistsEmptyState struct {
					Style    string `json:"style"`
					Messages []struct {
						Text   string `json:"text"`
						Image  string `json:"image"`
						TextEn string `json:"text_en"`
					} `json:"messages"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"profile_no_playlists_empty_state"`
				AppleGenericUpsellPremiumExplanation struct {
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Fail struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"fail"`
							AndroidSubscription string `json:"android_subscription"`
							Name                string `json:"name"`
							Success             struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"success"`
							Subscription string `json:"subscription"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_upsell_premium_explanation"`
				AppleGenericUpsellAnnualErrorConfuse struct {
					Comment  string `json:"comment"`
					Style    string `json:"style"`
					Messages []struct {
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
							ID   string `json:"id"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_upsell_annual_error_confuse"`
				AppleGenericWhypay struct {
					Style    string `json:"style"`
					Messages []struct {
						WebContentURL string `json:"web_content_url"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style string `json:"style"`
						Title string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_whypay"`
				AppleGenericTrackUnavailable struct {
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk string `json:"title_uk"`
						TextUk  string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							URL  string `json:"url"`
							Fail struct {
								Name string `json:"name"`
							} `json:"fail"`
							Name    string `json:"name"`
							Success struct {
								Name string `json:"name"`
							} `json:"success"`
						} `json:"action"`
						TitleUk string `json:"title_uk"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_track_unavailable"`
				WavesEmptyStateGridheader struct {
					Comment  string `json:"comment"`
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
							Image string `json:"image"`
						} `json:"background"`
						BrandedBackground struct {
							Gradient      bool   `json:"gradient"`
							Width         int    `json:"width"`
							Height        int    `json:"height"`
							PaletteBottom string `json:"palette_bottom"`
						} `json:"branded_background"`
					} `json:"messages"`
					Actions []struct {
						Style     string `json:"style"`
						Title     string `json:"title"`
						TitleEn   string `json:"title_en"`
						TextColor string `json:"text_color"`
						BgColor   string `json:"bg_color"`
						Action    struct {
							URL  string `json:"url"`
							Fail struct {
								Name string `json:"name"`
							} `json:"fail"`
							Name    string `json:"name"`
							Success struct {
								Name string `json:"name"`
							} `json:"success"`
						} `json:"action"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"waves_empty_state_gridheader"`
				TestYotaKit struct {
					Messages []struct {
						Image   string `json:"image"`
						Bullets []struct {
							Color string `json:"color"`
							Text  string `json:"text"`
							Image string `json:"image"`
						} `json:"bullets"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Fail struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"fail"`
							AndroidSubscription string `json:"android_subscription"`
							Name                string `json:"name"`
							Success             struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"success"`
							Subscription string `json:"subscription"`
						} `json:"action"`
						Style string `json:"style"`
						Title string `json:"title"`
					} `json:"actions"`
					PinActions bool   `json:"pin_actions"`
					Type       string `json:"type"`
				} `json:"test_yota_kit"`
				PlusSecret struct {
					Style                string `json:"style"`
					CloseGestureDisabled bool   `json:"close_gesture_disabled"`
					Messages             []struct {
						Text       string `json:"text"`
						Image      string `json:"image"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						Title string `json:"title"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
							ID   string `json:"id"`
						} `json:"action"`
						Style string `json:"style"`
						Title string `json:"title"`
					} `json:"actions"`
					Source           string      `json:"source"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"plus_secret"`
				DplaylistsEmptyStateGridheader struct {
					Comment  string `json:"comment"`
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
							Image string `json:"image"`
						} `json:"background"`
						BrandedBackground struct {
							Gradient      bool   `json:"gradient"`
							Width         int    `json:"width"`
							Height        int    `json:"height"`
							PaletteBottom string `json:"palette_bottom"`
						} `json:"branded_background"`
					} `json:"messages"`
					Actions []struct {
						Style     string `json:"style"`
						Title     string `json:"title"`
						TitleEn   string `json:"title_en"`
						TextColor string `json:"text_color"`
						BgColor   string `json:"bg_color"`
						Action    struct {
							URL  string `json:"url"`
							Fail struct {
								Name string `json:"name"`
							} `json:"fail"`
							Name    string `json:"name"`
							Success struct {
								Name string `json:"name"`
							} `json:"success"`
						} `json:"action"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"dplaylists_empty_state_gridheader"`
				PlaylistsEmptyStateGridheader struct {
					Comment  string `json:"comment"`
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
							Image string `json:"image"`
						} `json:"background"`
						BrandedBackground struct {
							Gradient      bool   `json:"gradient"`
							Width         int    `json:"width"`
							Height        int    `json:"height"`
							PaletteBottom string `json:"palette_bottom"`
						} `json:"branded_background"`
					} `json:"messages"`
					Actions []struct {
						Style     string `json:"style"`
						Title     string `json:"title"`
						TitleEn   string `json:"title_en"`
						TextColor string `json:"text_color"`
						BgColor   string `json:"bg_color"`
						Action    struct {
							URL  string `json:"url"`
							Fail struct {
								Name string `json:"name"`
							} `json:"fail"`
							Name    string `json:"name"`
							Success struct {
								Name string `json:"name"`
							} `json:"success"`
						} `json:"action"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"playlists_empty_state_gridheader"`
				AppleGenericUpsellPremiumExplanation0 struct {
					Style    string `json:"style"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name    string `json:"name"`
							Success struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"success"`
							SubscriptionType    string `json:"subscription_type"`
							AndroidSubscription string `json:"android_subscription"`
							Fail                struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"fail"`
							Subscription string `json:"subscription"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
				} `json:"apple_generic_upsell_premium_explanation0"`
				SubscriptionSuccess struct {
					Comment  string `json:"comment"`
					Messages []struct {
						Title      string `json:"title"`
						Text       string `json:"text"`
						Image      string `json:"image"`
						TitleEn    string `json:"title_en"`
						TextEn     string `json:"text_en"`
						Background struct {
							Color string `json:"color"`
						} `json:"background"`
						TitleUk string `json:"title_uk"`
						TextUk  string `json:"text_uk"`
					} `json:"messages"`
					Actions []struct {
						Action struct {
							Name string `json:"name"`
						} `json:"action"`
						Style   string `json:"style"`
						TitleEn string `json:"title_en"`
						Title   string `json:"title"`
					} `json:"actions"`
					PinActions       bool        `json:"pin_actions"`
					AllowCloseButton interface{} `json:"allow_close_button"`
					Type             string      `json:"type"`
				} `json:"subscription_success"`
			} `json:"action_kit_pages"`
			PlayAdsInPodcasts bool `json:"play_ads_in_podcasts"`
			AbExperiment      []struct {
				Group string `json:"group"`
				Name  string `json:"name"`
			} `json:"ab_experiment"`
			DaysWithoutAds int  `json:"days_without_ads"`
			LyricsEnabled  bool `json:"lyrics_enabled"`
			Events         struct {
				WebviewHowUnsubscribe struct {
					URL   string `json:"url"`
					Name  string `json:"name"`
					URLEn string `json:"url_en"`
				} `json:"webview-how-unsubscribe"`
				OfflineRadio struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"offline-radio"`
				WebviewGiftCode struct {
					URL     string `json:"url"`
					Success struct {
						Name string `json:"name"`
						ID   string `json:"id"`
					} `json:"success"`
					Name     string `json:"name"`
					Auth     bool   `json:"auth"`
					InWebkit bool   `json:"in_webkit"`
				} `json:"webview-gift-code"`
				WebviewSupport struct {
					Name string `json:"name"`
				} `json:"webview-support"`
				SubscriptionExpired struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"subscription_expired"`
				GridZvukplus struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"grid-zvukplus"`
				MultitaskingFinished struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"multitasking-finished"`
				UnregLoginPhone struct {
					Name    string `json:"name"`
					Success struct {
						ID   string `json:"id"`
						Name string `json:"name"`
					} `json:"success"`
					LoginTrigger string `json:"login_trigger"`
				} `json:"unreg-login-phone"`
				ZeroPenIviGrid struct {
					URL  string `json:"url"`
					Name string `json:"name"`
				} `json:"0pen-ivi-grid"`
				EmptyStateSubscribe struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"empty-state-subscribe"`
				OfflineDownload struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"offline-download"`
				LotsOfSubscriptionsAction struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"lots_of_subscriptions-action"`
				Privacy struct {
					URL   string `json:"url"`
					URLUk string `json:"url_uk"`
					Name  string `json:"name"`
				} `json:"privacy"`
				OfflineStreaming struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"offline-streaming"`
				HuaweiHqPromo struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"huawei_hq_promo"`
				NotEnoughStorage struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"not-enough-storage"`
				Advert struct {
					URL  string `json:"url"`
					Name string `json:"name"`
				} `json:"advert"`
				PaywallZvukplus struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"paywall-zvukplus"`
				PlaylistSber1 struct {
					FreebanFeatured bool   `json:"freeban_featured"`
					TrackNumber     int    `json:"track_number"`
					ID              int    `json:"id"`
					Name            string `json:"name"`
					Autoplay        bool   `json:"autoplay"`
				} `json:"playlist-sber-1"`
				PaywallNotEnoughStorage struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"paywall-not-enough-storage"`
				TriggerMicrophone struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"trigger-microphone"`
				GridZvukplusLight struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"grid-zvukplus-light"`
				OfflineCollection struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"offline-collection"`
				Terms struct {
					URL   string `json:"url"`
					Name  string `json:"name"`
					URLEn string `json:"url_en"`
				} `json:"terms"`
				ExplicitBlock struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"explicit-block"`
				PersonalPolicy struct {
					URL  string `json:"url"`
					Name string `json:"name"`
				} `json:"personal-policy"`
				AttachAccountSber struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"attach-account-sber"`
				TrackUnavailable struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"track-unavailable"`
				Logout struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"logout"`
				LoginSuccess struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"login-success"`
				OfflineLyrics struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"offline-lyrics"`
				LoginSuccessPlus struct {
					URL  string `json:"url"`
					Name string `json:"name"`
				} `json:"login-success-plus"`
				UnregLoginOther struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"unreg-login-other"`
				HowUnsubscribe struct {
					URL   string `json:"url"`
					Name  string `json:"name"`
					URLEn string `json:"url_en"`
				} `json:"how_unsubscribe"`
				GiftCode struct {
					URL     string `json:"url"`
					Success struct {
						Name string `json:"name"`
						ID   string `json:"id"`
					} `json:"success"`
					Name     string `json:"name"`
					Auth     bool   `json:"auth"`
					InWebkit bool   `json:"in_webkit"`
				} `json:"gift-code"`
				StoriesError struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"stories-error"`
				PaywallZvukplusLight struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"paywall-zvukplus-light"`
				FirstStartOnUpdate struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"first-start-on-update"`
			} `json:"events"`
			AvailableSubscriptions      string   `json:"available-subscriptions"`
			ClickstreamBgUpdateInterval int      `json:"clickstream_bg_update_interval"`
			ShowHeader                  string   `json:"show_header"`
			ZvukPlusSubscriptions       []string `json:"zvuk_plus_subscriptions"`
			BrandedPlaylists            struct {
				Num123456789 struct {
					Comment              string        `json:"comment"`
					Style                string        `json:"style"`
					Target               []interface{} `json:"target"`
					CloseGestureDisabled bool          `json:"close_gesture_disabled"`
					Messages             []struct {
						Action struct {
							URL  string `json:"url"`
							Name string `json:"name"`
						} `json:"action"`
						BrandedBackground struct {
							Gradient      bool   `json:"gradient"`
							Width         int    `json:"width"`
							Src           string `json:"src"`
							Height        int    `json:"height"`
							PaletteBottom string `json:"palette_bottom"`
						} `json:"branded_background"`
					} `json:"messages"`
					Actions []struct {
						Style     string `json:"style"`
						Rounded   bool   `json:"rounded"`
						Title     string `json:"title"`
						TextColor string `json:"text_color"`
						BgColor   string `json:"bg_color"`
						Action    struct {
							URL  string `json:"url"`
							Name string `json:"name"`
						} `json:"action"`
					} `json:"actions"`
					Source           string `json:"source"`
					PinActions       bool   `json:"pin_actions"`
					AllowCloseButton bool   `json:"allow_close_button"`
					Type             string `json:"type"`
					ID               string `json:"id"`
				} `json:"123456789"`
			} `json:"branded_playlists"`
			SearchView               string `json:"search_view"`
			ShowRestoreButton        bool   `json:"show_restore_button"`
			ClikstreamUpdateDuration int    `json:"clikstream_update_duration"`
			ContentAvailable         string `json:"content_available"`
			ActionAliases            struct {
				AzercellSubscriptionReloadSettings struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"azercell-subscription-reload-settings"`
				Tele2UnsubscribeReloadSettings struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"tele2-unsubscribe-reload-settings"`
				Tele2RuSubscriptionFailureAlias struct {
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"tele2-ru-subscription-failure-alias"`
				MegafonSubscriptionReloadSettings struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"megafon-subscription-reload-settings"`
				UcellUnsubscribeReloadSettings struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"ucell-unsubscribe-reload-settings"`
				BeelineRuUnsubscribeReloadSettings struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"beeline-ru-unsubscribe-reload-settings"`
				AppleGenericSubscriptionReloadSettingsSberprimeSub struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"apple-generic-subscription-reload-settings_sberprime_sub"`
				BeelineKzSubscriptionReloadSettings struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"beeline-kz-subscription-reload-settings"`
				Tele2SubscriptionReloadSettingsSberInsurance struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"tele2-subscription-reload-settings_sber-insurance"`
				YotaSubscriptionReloadSettings struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"yota-subscription-reload-settings"`
				MegafonUnsubscribeReloadSettings struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"megafon-unsubscribe-reload-settings"`
				AppleGenericSubscriptionReloadSettingsSberInsurance struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"apple-generic-subscription-reload-settings_sber-insurance"`
				Tele2ReloadSettings struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"tele2-reload-settings"`
				MegafonSubscriptionReloadSettingsSberInsurance struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"megafon-subscription-reload-settings_sber-insurance"`
				YotaUnsubscribeReloadSettings struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"yota-unsubscribe-reload-settings"`
				Tele2SubscriptionReloadSettings struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"tele2-subscription-reload-settings"`
				AppleGenericSubscriptionReloadSettings struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"apple-generic-subscription-reload-settings"`
				BeelineRuSubscriptionReloadSettings struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"beeline-ru-subscription-reload-settings"`
				BeelineRuSubscriptionReloadSettingsSberInsurance struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"beeline-ru-subscription-reload-settings_sber-insurance"`
				VelcomSubscriptionReloadSettings struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"velcom-subscription-reload-settings"`
				PlusSubscriptionReloadSettings struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"plus-subscription-reload-settings"`
				CodeSubscriptionReloadSettings struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"code-subscription-reload-settings"`
				RestoreSubscriptionReloadSettings struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"restore-subscription-reload-settings"`
				ShowBanner struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"show-banner"`
				Tele2ChangeSubscriptionReloadSettings struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"tele2-change-subscription-reload-settings"`
				UcellSubscriptionReloadSettings struct {
					IDReloadSuccess string `json:"id_reload_success"`
					IDReloadFail    string `json:"id_reload_fail"`
					Name            string `json:"name"`
				} `json:"ucell-subscription-reload-settings"`
			} `json:"action_aliases"`
			XynRules []struct {
				Count          int    `json:"count"`
				RuleIdentifier string `json:"ruleIdentifier"`
				ShouldRepeat   bool   `json:"shouldRepeat"`
				EventType      int    `json:"eventType"`
				ValidPeriod    int    `json:"validPeriod"`
				Rule           struct {
					Action int    `json:"action"`
					Name   string `json:"name"`
				} `json:"rule"`
				Action struct {
					Name string `json:"name"`
					ID   string `json:"id"`
				} `json:"action"`
			} `json:"xyn_rules"`
			ClickstreamUpdateInterval int `json:"clickstream_update_interval"`
		} `json:"settings"`
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
		Tracks    map[string]Track `json:"tracks"`
		Playlists struct {
		} `json:"playlists"`
		RadioWaves struct {
		} `json:"radio_waves"`
		Releases map[string]Release `json:"releases"`
		Artists  struct {
		} `json:"artists"`
		Labels struct {
		} `json:"labels"`
		Users struct {
		} `json:"users"`
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
