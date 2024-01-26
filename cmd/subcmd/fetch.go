package subcmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"text/tabwriter"
	"time"

	"github.com/silaselisha/coins/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type marketData struct {
	Price            float64     `json:"price"`
	Volume24h        float64     `json:"volume_24h"`
	VolumeChange24h  float64     `json:"volume_change_24h"`
	PercentChange1h  float64     `json:"percent_change_1h"`
	PercentChange24h float64     `json:"percent_change_24h"`
	PercentChange7d  float64     `json:"percent_change_7d"`
	PercentChange30d float64     `json:"percent_change_30d"`
	PercentChange60d float64     `json:"percent_change_60d"`
	PercentChange90d float64     `json:"percent_change_90d"`
	MarketCap        float64     `json:"market_cap"`
	MarketCapDomin   float64     `json:"market_cap_dominance"`
	FullyDilutedCap  float64     `json:"fully_diluted_market_cap"`
	LastUpdated      time.Time   `json:"last_updated"`
}

type cryptoCurrency struct {
	ID                  int64       `json:"id"`
	Name                string      `json:"name"`
	Symbol              string      `json:"symbol"`
	Slug                string      `json:"slug"`
	NumMarketPairs      int64       `json:"num_market_pairs"`
	DateAdded           time.Time   `json:"date_added"`
	Tags                []string    `json:"tags"`
	MaxSupply           float64     `json:"max_supply"`
	CirculatingSupply   float64     `json:"circulating_supply"`
	TotalSupply         float64     `json:"total_supply"`
	InfiniteSupply      bool        `json:"infinite_supply"`
	CMCRank             float64     `json:"cmc_rank"`
	LastUpdated         time.Time   `json:"last_updated"`
	Quote               struct {
		Currency marketData `json:"USD"`
	} `json:"quote"`
}

type Response struct {
	Data []*cryptoCurrency `json:"data"`
}

var (
  Listing string
	FetchCmd = &cobra.Command{
		Use:   "fetch",
		Short: "Fetch crypto coins info",
		Long:  "Fetch coins based on the crypto conin symbol provided",
		Run: func(cmd *cobra.Command, args []string) {
      value := viper.GetString("listing")
      go asciiLoader(100 * time.Millisecond)
			request := serverRequestHandler(value)
      data := clientRequestHandler(request)
			printCryptoData(data)
		},
	}
)

func asciiLoader(delay time.Duration) {
  fmt.Print(`    
 ___  ___  _   _  ___  _____  __     ___  __   _  _   _  ___  
(  _)(   )( \_/ )(   )(_   _)(  )   (  _)(  ) ( )( \ ( )/  _) 
| |  | O  |\   / | O  | | |  /  \   | |  /  \ | || \\| |\_"-. 
( )_ ( _ (  ( )  ( __/  ( ) ( O  )  ( )_( O  )( )( )\\ ) __) )
/___\/_\\_| |_|  /_\    /_\  \__/   /___\\__/ /_\/_\ \_\/___/ 
                                                                
`)
  time.Sleep(delay)
}

func printCryptoData(data []*cryptoCurrency) {
	const format = "%v\t%v\t%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 3, ' ', 0)
	fmt.Fprintf(tw, format, "Name", "Symbol", "Price", "Max Supply", "Market Cap", "Volume in 24h", "Percent Change 24h")
	fmt.Fprintf(tw, format, "-------", "-------", "-------", "-------", "-------", "-------", "-------")
	for _, info := range data {
		fmt.Fprintf(tw, format, info.Name, info.Symbol, info.Quote.Currency.Price, info.MaxSupply, info.Quote.Currency.MarketCap, info.Quote.Currency.Volume24h, info.Quote.Currency.PercentChange24h)
	}
	tw.Flush()
}

func clientRequestHandler(req *http.Request) ([]*cryptoCurrency) {
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    log.Panic(err)
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    log.Panic(err)
  }

  var result Response
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		fmt.Println("error unmarshalling...")
		log.Print(err)
	}
  return result.Data
}


func serverRequestHandler(listing string) *http.Request {
  envs, err := util.LoadEnvs(".")
  if err != nil {
    log.Print(err)
  }
  
  endpoint := fmt.Sprintf("https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/%v", listing)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Print(err)
	}
	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "10")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Set("X-CMC_PRO_API_KEY", envs.ApiKey)
	req.URL.RawQuery = q.Encode()
	return req
}
