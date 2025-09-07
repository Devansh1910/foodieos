package getOutletFood

// Request payload
type RequestPayload struct {
    Platform     string  `json:"platform"`
    Country      string  `json:"country"`
    City         string  `json:"city"`
    State        string  `json:"state"`
    Lat          float64 `json:"lat"`
    Lon          float64 `json:"lon"`
    OutletID     int     `json:"outletid"`
    FoodCategory string  `json:"foodCategory"`
    Date         string  `json:"date"`
}

// Top-level response (matches PDF-ish)
type Response struct {
    Status int    `json:"status"`
    Code   int    `json:"code"`
    Result string `json:"result"`
    Msg    string `json:"msg"`
    Output Output `json:"output"`
}

type Output struct {
    DefaultDisplatCat          string      `json:"defaultDisplatCat"`
    DefaultCatImgUrl           string      `json:"defaultCatImgUrl"`
    PrepareDefaultTimeinMin    int         `json:"prepareDefaultTimeinMin"`
    PrepareCapTimeinMin        int         `json:"prepareCapTimeinMin"`
    HideShowBeforeEndTimeinMin int         `json:"hideShowBeforeEndTimeinMin"`
    StatusUpdateTimeinSec      int         `json:"statusUpdateTimeinSec"`
    OutletName                 string      `json:"outletName"`
    City                       CityInfo    `json:"city"`
    R                          []Item      `json:"r"`
    Cat                        []string    `json:"cat"`
    Repeat                     interface{} `json:"repeat"`
    Cats                       []Cat       `json:"cats"`
    Aqt                        int         `json:"aqt"`
    Nams                       string      `json:"nams"`
    EmptyCart                  string      `json:"emptyCart"`
    ItemQuantityLimit          string      `json:"itemQuantityLimit"`
    DeliveryLocation           string      `json:"deliveryLocation"`
    Pu                         interface{} `json:"pu"`
    Ph                         interface{} `json:"ph"`
    Offers                     []Offer     `json:"offers"`
}

type CityInfo struct {
    ID          int           `json:"id"`
    Name        string        `json:"name"`
    Region      string        `json:"region"`
    HasSubCities bool         `json:"hasSubCities"`
    State       string        `json:"state"`
    Subcities   []interface{} `json:"subcities"`
    Lat         string        `json:"lat"`
    Lng         string        `json:"lng"`
    Image       *string       `json:"image"`
    ImageR      *string       `json:"imageR"`
}

type Item struct {
    Category        string        `json:"category"`
    ItemPackage     interface{}   `json:"itemPackage"`
    FoodType        int           `json:"foodType"`
    ImgData         interface{}   `json:"imgData"`
    BestSeller      bool          `json:"bestSeller"`
    Combo           bool          `json:"combo"`
    ComboItems      []interface{} `json:"comboItems"`
    AddOn           bool          `json:"addOn"`
    AddOnItems      []interface{} `json:"addOnItems"`
    PreparationTime int           `json:"preparationTime"`
    Upgradable      bool          `json:"upgradable"`
    Upsellable      bool          `json:"upsellable"`
    UpgradeItems    []interface{} `json:"upgradeItems"`
    UpsellItems     []interface{} `json:"upsellItems"`
    PreparationType string        `json:"preparationType"`
    ID              string        `json:"id"`
    H               string        `json:"h"`
    Op              int           `json:"op"`
    Dp              int           `json:"dp"`
    Ho              string        `json:"ho"`
    Ct              string        `json:"ct"`
    Veg             bool          `json:"veg"`
    Dis             interface{}   `json:"dis"`
    I               *string       `json:"i"`
    Sf              interface{}   `json:"sf"`
    Wt              string        `json:"wt"`
    En              string        `json:"en"`
    Fa              string        `json:"fa"`
}

type Cat struct {
    Name     string `json:"name"`
    ImageUrl string `json:"imageUrl"`
}

type Offer struct {
    ID                string   `json:"id"`
    ChainKey          string   `json:"chainKey"`
    VouId             int      `json:"vouId"`
    VouDesc           string   `json:"vouDesc"`
    Category          string   `json:"category"`
    Type              string   `json:"type"`
    Bank              string   `json:"bank"`
    BankId            int      `json:"bankId"`
    BankVouType       string   `json:"bankVouType"`
    BankVouTypeSub    string   `json:"bankVouTypeSub"`
    DiscAppPayType    string   `json:"discAppPayType"`
    AppFor            string   `json:"appFor"`
    ForBooklet        string   `json:"forBooklet"`
    Loyalty           bool     `json:"loyalty"`
    Tnc               []string `json:"tnc"`
    ValidFrom         string   `json:"validFrom"`
    ValidTo           string   `json:"validTo"`
    ImageVertical     string   `json:"imageVertical"`
    ImageHorizontal   string   `json:"imageHorizontal"`
    BankImage         string   `json:"bankImage"`
    Cities            string   `json:"cities"`
    Subscription      bool     `json:"subscription"`
    RedemptionOutlet  string   `json:"redemptionOutlet"`
    Theater           string   `json:"theater"`
    OnTheFly          bool     `json:"onTheFly"`
    BestOffer         bool     `json:"bestOffer"`
    BestOfferUrl      string   `json:"bestOfferUrl"`
}
