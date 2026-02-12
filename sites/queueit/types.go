package queueit

import "time"

type ChallengeResponse struct {
	Function   string `json:"function"`
	SessionID  string `json:"sessionId"`
	Meta       string `json:"meta"`
	Parameters struct {
		Type       string `json:"type"`
		Input      string `json:"input"`
		Runs       int    `json:"runs"`
		Complexity int    `json:"complexity"`
		ZeroCount  int    `json:"zeroCount"`
	} `json:"parameters"`
	ChallengeDetails string `json:"challengeDetails"`
}

type QueueItVerify struct {
	ChallengeType    string `json:"challengeType"`
	SessionID        string `json:"sessionId"`
	ChallengeDetails string `json:"challengeDetails"`
	Solution         string `json:"solution"`
	Stats            Stats  `json:"stats"`
	CustomerID       string `json:"customerId"`
	EventID          string `json:"eventId"`
	Version          int    `json:"version"`
}

type Stats struct {
	UserAgent      string `json:"userAgent"`
	Screen         string `json:"screen"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browserVersion"`
	IsMobile       bool   `json:"isMobile"`
	Os             string `json:"os"`
	OsVersion      string `json:"osVersion"`
	CookiesEnabled bool   `json:"cookiesEnabled"`
	Tries          int    `json:"tries"`
	Duration       int    `json:"duration"`
}

type VerifyResponse struct {
	IsVerified  bool         `json:"isVerified"`
	Timestamp   time.Time    `json:"timestamp"`
	SessionInfo QueueSession `json:"sessionInfo"`
}

type QueueSession struct {
	SessionID     string    `json:"sessionId"`
	Timestamp     time.Time `json:"timestamp"`
	Checksum      string    `json:"checksum"`
	SourceIP      string    `json:"sourceIp"`
	ChallengeType string    `json:"challengeType"`
	Version       int       `json:"version"`
	CustomerID    string    `json:"customerId"`
	WaitingRoomID string    `json:"waitingRoomId"`
}

type Enqueue struct {
	ChallengeSessions []QueueSession `json:"challengeSessions"`
	LayoutName        string         `json:"layoutName"`
	CustomURLParams   string         `json:"customUrlParams"`
	TargetURL         string         `json:"targetUrl"`
	Referrer          string         `json:"Referrer"`
}

type EnqueueResp struct {
	ChallengeFailed              bool   `json:"challengeFailed"`
	CustomDataUniqueKeyViolation bool   `json:"customDataUniqueKeyViolation"`
	InvalidQueueitEnqueueToken   bool   `json:"invalidQueueitEnqueueToken"`
	MissingCustomDataKey         bool   `json:"missingCustomDataKey"`
	QueueID                      string `json:"queueId"`
	RedirectURL                  string `json:"redirectUrl"`
	ServerIsBusy                 bool   `json:"serverIsBusy"`
}

type InQueueResp struct {
	IsBeforeOrIdle bool   `json:"isBeforeOrIdle"`
	PageID         string `json:"pageId"`
	PageClass      string `json:"pageClass"`
	QueueState     int    `json:"QueueState"`
	Ticket         struct {
		QueueNumber             any       `json:"queueNumber"`
		UsersInLineAheadOfYou   string    `json:"usersInLineAheadOfYou"`
		ExpectedServiceTime     string    `json:"expectedServiceTime"`
		QueuePaused             bool      `json:"queuePaused"`
		LastUpdatedUTC          time.Time `json:"lastUpdatedUTC"`
		WhichIsIn               string    `json:"whichIsIn"`
		LastUpdated             string    `json:"lastUpdated"`
		Progress                int       `json:"progress"`
		TimeZonePostfix         string    `json:"timeZonePostfix"`
		ExpectedServiceTimeUTC  time.Time `json:"expectedServiceTimeUTC"`
		CustomURLParams         string    `json:"customUrlParams"`
		SdkVersion              any       `json:"sdkVersion"`
		WindowStartTimeUTC      any       `json:"windowStartTimeUTC"`
		WindowStartTime         string    `json:"windowStartTime"`
		SecondsToStart          int       `json:"secondsToStart"`
		UsersInQueue            int       `json:"usersInQueue"`
		EventStartTimeFormatted string    `json:"eventStartTimeFormatted"`
		EventStartTimeUTC       time.Time `json:"eventStartTimeUTC"`
	} `json:"ticket"`
	Message any `json:"message"`
	Texts   struct {
		CountdownFinishedText        string   `json:"countdownFinishedText"`
		QueueBody                    any      `json:"queueBody"`
		QueueHeader                  any      `json:"queueHeader"`
		Header                       string   `json:"header"`
		Body                         string   `json:"body"`
		DisclaimerText               any      `json:"disclaimerText"`
		StyleSheets                  string   `json:"styleSheets"`
		Javascripts                  []any    `json:"javascripts"`
		LogoSrc                      string   `json:"logoSrc"`
		ToppanelIFrameSrc            any      `json:"toppanelIFrameSrc"`
		SidepanelIFrameSrc           any      `json:"sidepanelIFrameSrc"`
		LeftpanelIFrameSrc           any      `json:"leftpanelIFrameSrc"`
		RightpanelIFrameSrc          any      `json:"rightpanelIFrameSrc"`
		MiddlepanelIFrameSrc         any      `json:"middlepanelIFrameSrc"`
		BottompanelIFrameSrc         any      `json:"bottompanelIFrameSrc"`
		FaviconURL                   string   `json:"faviconUrl"`
		Languages                    []any    `json:"languages"`
		WhatIsThisURL                string   `json:"whatIsThisUrl"`
		QueueItLogoPointsToURL       string   `json:"queueItLogoPointsToUrl"`
		WelcomeSoundUrls             []string `json:"welcomeSoundUrls"`
		CookiesAllowedInfoText       string   `json:"cookiesAllowedInfoText"`
		CookiesNotAllowedInfoText    string   `json:"cookiesNotAllowedInfoText"`
		CookiesAllowedInfoTooltip    string   `json:"cookiesAllowedInfoTooltip"`
		CookiesNotAllowedInfoTooltip string   `json:"cookiesNotAllowedInfoTooltip"`
		AppleTouchIconURL            string   `json:"AppleTouchIconUrl"`
		DocumentTitle                string   `json:"DocumentTitle"`
		Tags                         []any    `json:"tags"`
		MessageUpdatedTimeAgo        string   `json:"messageUpdatedTimeAgo"`
	} `json:"texts"`
	Layout struct {
		LanguageSelectorVisible       bool `json:"languageSelectorVisible"`
		LogoVisible                   bool `json:"logoVisible"`
		DynamicMessageVisible         bool `json:"dynamicMessageVisible"`
		ReminderEmailVisible          bool `json:"reminderEmailVisible"`
		ExpectedServiceTimeVisible    bool `json:"expectedServiceTimeVisible"`
		QueueNumberVisible            bool `json:"queueNumberVisible"`
		UsersInLineAheadOfYouVisible  bool `json:"usersInLineAheadOfYouVisible"`
		WhichIsInVisible              bool `json:"whichIsInVisible"`
		SidePanelVisible              bool `json:"sidePanelVisible"`
		TopPanelVisible               bool `json:"topPanelVisible"`
		LeftPanelVisible              bool `json:"leftPanelVisible"`
		RightPanelVisible             bool `json:"rightPanelVisible"`
		MiddlePanelVisible            bool `json:"middlePanelVisible"`
		BottomPanelVisible            bool `json:"bottomPanelVisible"`
		UsersInQueueVisible           bool `json:"usersInQueueVisible"`
		QueueIsPausedVisible          bool `json:"queueIsPausedVisible"`
		ReminderVisible               bool `json:"reminderVisible"`
		ServicedSoonVisible           bool `json:"servicedSoonVisible"`
		FirstInLineVisible            bool `json:"firstInLineVisible"`
		QueueNumberLoadingVisible     bool `json:"queueNumberLoadingVisible"`
		ProgressVisible               bool `json:"progressVisible"`
		IsRedirectPromptDialogEnabled bool `json:"isRedirectPromptDialogEnabled"`
		IsQueueitFooterHidden         bool `json:"isQueueitFooterHidden"`
	} `json:"layout"`
	ForecastStatus string `json:"forecastStatus"`
	LayoutName     string `json:"layoutName"`
	LayoutVersion  int64  `json:"layoutVersion"`
	UpdateInterval int    `json:"updateInterval"`
}

type LinkFound struct {
	RedirectURL        string `json:"redirectUrl"`
	IsRedirectToTarget bool   `json:"isRedirectToTarget"`
}
