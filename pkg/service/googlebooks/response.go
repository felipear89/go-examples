package googlebooks

type ErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Errors  []struct {
			Message string `json:"message"`
			Domain  string `json:"domain"`
			Reason  string `json:"reason"`
		} `json:"errors"`
	} `json:"error"`
}

type Response struct {
	Kind       string `json:"kind"`
	TotalItems int    `json:"totalItems"`
	Items      []struct {
		Kind       string `json:"kind"`
		Id         string `json:"id"`
		Etag       string `json:"etag"`
		SelfLink   string `json:"selfLink"`
		VolumeInfo struct {
			Title               string   `json:"title"`
			Authors             []string `json:"authors"`
			Publisher           string   `json:"publisher,omitempty"`
			PublishedDate       string   `json:"publishedDate"`
			Description         string   `json:"description,omitempty"`
			IndustryIdentifiers []struct {
				Type       string `json:"type"`
				Identifier string `json:"identifier"`
			} `json:"industryIdentifiers"`
			ReadingModes struct {
				Text  bool `json:"text"`
				Image bool `json:"image"`
			} `json:"readingModes"`
			PageCount           int      `json:"pageCount"`
			PrintType           string   `json:"printType"`
			Categories          []string `json:"categories,omitempty"`
			MaturityRating      string   `json:"maturityRating"`
			AllowAnonLogging    bool     `json:"allowAnonLogging"`
			ContentVersion      string   `json:"contentVersion"`
			PanelizationSummary struct {
				ContainsEpubBubbles  bool `json:"containsEpubBubbles"`
				ContainsImageBubbles bool `json:"containsImageBubbles"`
			} `json:"panelizationSummary"`
			ImageLinks struct {
				SmallThumbnail string `json:"smallThumbnail"`
				Thumbnail      string `json:"thumbnail"`
			} `json:"imageLinks,omitempty"`
			Language            string `json:"language"`
			PreviewLink         string `json:"previewLink"`
			InfoLink            string `json:"infoLink"`
			CanonicalVolumeLink string `json:"canonicalVolumeLink"`
			Subtitle            string `json:"subtitle,omitempty"`
		} `json:"volumeInfo"`
		SaleInfo struct {
			Country     string `json:"country"`
			Saleability string `json:"saleability"`
			IsEbook     bool   `json:"isEbook"`
			ListPrice   struct {
				Amount       float64 `json:"amount"`
				CurrencyCode string  `json:"currencyCode"`
			} `json:"listPrice,omitempty"`
			RetailPrice struct {
				Amount       float64 `json:"amount"`
				CurrencyCode string  `json:"currencyCode"`
			} `json:"retailPrice,omitempty"`
			BuyLink string `json:"buyLink,omitempty"`
			Offers  []struct {
				FinskyOfferType int `json:"finskyOfferType"`
				ListPrice       struct {
					AmountInMicros int    `json:"amountInMicros"`
					CurrencyCode   string `json:"currencyCode"`
				} `json:"listPrice"`
				RetailPrice struct {
					AmountInMicros int    `json:"amountInMicros"`
					CurrencyCode   string `json:"currencyCode"`
				} `json:"retailPrice"`
				Giftable bool `json:"giftable"`
			} `json:"offers,omitempty"`
		} `json:"saleInfo"`
		AccessInfo struct {
			Country                string `json:"country"`
			Viewability            string `json:"viewability"`
			Embeddable             bool   `json:"embeddable"`
			PublicDomain           bool   `json:"publicDomain"`
			TextToSpeechPermission string `json:"textToSpeechPermission"`
			Epub                   struct {
				IsAvailable  bool   `json:"isAvailable"`
				AcsTokenLink string `json:"acsTokenLink,omitempty"`
			} `json:"epub"`
			Pdf struct {
				IsAvailable  bool   `json:"isAvailable"`
				AcsTokenLink string `json:"acsTokenLink,omitempty"`
			} `json:"pdf"`
			WebReaderLink       string `json:"webReaderLink"`
			AccessViewStatus    string `json:"accessViewStatus"`
			QuoteSharingAllowed bool   `json:"quoteSharingAllowed"`
		} `json:"accessInfo"`
		SearchInfo struct {
			TextSnippet string `json:"textSnippet"`
		} `json:"searchInfo,omitempty"`
	} `json:"items"`
}
