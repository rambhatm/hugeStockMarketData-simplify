package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var stockList = []string{"a", "aal", "aap", "aapl", "abbv", "abc", "abmd", "abt", "acn", "adbe", "adi", "adm", "adp", "ads", "adsk", "aee", "aep", "aes", "afl", "agn", "aig", "aiv", "aiz", "ajg", "akam", "alb", "algn", "alk", "all", "alle", "alxn", "amat", "amcr", "amd", "ame", "amg", "amgn", "amp", "amt", "amzn", "anet", "anss", "antm", "aon", "aos", "apa", "apd", "aph", "aptv", "are", "arnc", "ato", "atvi", "avb", "avgo", "avy", "awk", "axp", "azo", "ba", "bac", "bax", "bbt", "bby", "bdx", "ben", "bf-b", "bhge", "biib", "bk", "bkng", "blk", "bll", "bmy", "br", "brk-b", "bsx", "bwa", "bxp", "c", "cag", "cah", "cat", "cb", "cboe", "cbre", "cbs", "cci", "ccl", "cdns", "cdw", "ce", "celg", "cern", "cf", "cfg", "chd", "chrw", "chtr", "ci", "cinf", "cl", "clx", "cma", "cmcsa", "cme", "cmg", "cmi", "cms", "cnc", "cnp", "cof", "cog", "coo", "cop", "cost", "coty", "cpb", "cpri", "cprt", "crm", "csco", "csx", "ctas", "ctl", "ctsh", "ctva", "ctxs", "cvs", "cvx", "cxo", "d", "dal", "dd", "de", "dfs", "dg", "dgx", "dhi", "dhr", "dis", "disca", "disck", "dish", "dlr", "dltr", "dov", "dow", "dre", "dri", "dte", "duk", "dva", "dvn", "dxc", "ea", "ebay", "ecl", "ed", "efx", "eix", "el", "emn", "emr", "eog", "eqix", "eqr", "es", "ess", "etfc", "etn", "etr", "evrg", "ew", "exc", "expd", "expe", "exr", "f", "fang", "fast", "fb", "fbhs", "fcx", "fdx", "fe", "ffiv", "fis", "fisv", "fitb", "flir", "fls", "flt", "fmc", "fox", "foxa", "frc", "frt", "fti", "ftnt", "ftv", "gd", "ge", "gild", "gis", "gl", "glw", "gm", "goog", "googl", "gpc", "gpn", "gps", "grmn", "gs", "gww", "hal", "has", "hban", "hbi", "hca", "hcp", "hd", "hes", "hfc", "hig", "hii", "hlt", "hog", "holx", "hon", "hp", "hpe", "hpq", "hrb", "hrl", "hsic", "hst", "hsy", "hum", "ibm", "ice", "idxx", "iex", "iff", "ilmn", "incy", "info", "intc", "intu", "ip", "ipg", "ipgp", "iqv", "ir", "irm", "isrg", "it", "itw", "ivz", "jbht", "jci", "jec", "jkhy", "jnj", "jnpr", "jpm", "jwn", "k", "key", "keys", "khc", "kim", "klac", "kmb", "kmi", "kmx", "ko", "kr", "kss", "ksu", "l", "lb", "ldos", "leg", "len", "lh", "lhx", "lin", "lkq", "lly", "lmt", "lnc", "lnt", "low", "lrcx", "luv", "lvs", "lw", "lyb", "m", "ma", "maa", "mac", "mar", "mas", "mcd", "mchp", "mck", "mco", "mdlz", "mdt", "met", "mgm", "mhk", "mkc", "mktx", "mlm", "mmc", "mmm", "mnst", "mo", "mos", "mpc", "mrk", "mro", "ms", "msci", "msft", "msi", "mtb", "mtd", "mu", "mxim", "myl", "nbl", "nclh", "ndaq", "nee", "nem", "nflx", "ni", "nke", "nlsn", "noc", "nov", "nrg", "nsc", "ntap", "ntrs", "nue", "nvda", "nvr", "nwl", "nws", "nwsa", "o", "oke", "omc", "orcl", "orly", "oxy", "payx", "pbct", "pcar", "peg", "pep", "pfe", "pfg", "pg", "pgr", "ph", "phm", "pkg", "pki", "pld", "pm", "pnc", "pnr", "pnw", "ppg", "ppl", "prgo", "pru", "psa", "psx", "pvh", "pwr", "pxd", "pypl", "qcom", "qrvo", "rcl", "re", "reg", "regn", "rf", "rhi", "rjf", "rl", "rmd", "rok", "rol", "rop", "rost", "rsg", "rtn", "sbac", "sbux", "schw", "see", "shw", "sivb", "sjm", "slb", "slg", "sna", "snps", "so", "spg", "spgi", "sre", "sti", "stt", "stx", "stz", "swk", "swks", "syf", "syk", "symc", "syy", "t", "tap", "tdg", "tel", "tfx", "tgt", "tif", "tjx", "tmo", "tmus", "tpr", "trip", "trow", "trv", "tsco", "tsn", "ttwo", "twtr", "txn", "txt", "ua", "uaa", "ual", "udr", "uhs", "ulta", "unh", "unm", "unp", "ups", "uri", "usb", "utx", "v", "var", "vfc", "viab", "vlo", "vmc", "vno", "vrsk", "vrsn", "vrtx", "vtr", "vz", "wab", "wat", "wba", "wcg", "wdc", "wec", "well", "wfc", "whr", "wltw", "wm", "wmb", "wmt", "wrk", "wu", "wy", "wynn", "xec", "xel", "xlnx", "xom", "xray", "xrx", "xyl", "yum", "zbh", "zion", "zts"}

func sp500only() {
	for _, stock := range stockList {
		oldpath := "./Stocks/" + stock + ".us.txt"
		newpath := "./sp500/" + stock + ".csv"
		err := os.Rename(oldpath, newpath)
		if err != nil {
			log.Printf("%s", err)

		}
	}
}

func processOneFile(name string) {
	start := time.Now()
	readFile, err := os.Open("./sp500/" + name + ".csv") // For read access.
	if err != nil {
		log.Println(err)
		return
	}
	writeFile, err := os.Create("./scratch/" + name + ".csv")
	if err != nil {
		log.Println(err)
	}
	defer readFile.Close()
	defer writeFile.Close()

	r := csv.NewReader(readFile)

	w := csv.NewWriter(writeFile)

	var header = true

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if header {
			header = false
			continue
		}

		date := record[0]
		a, _ := strconv.ParseFloat(record[1], 64)
		b, _ := strconv.ParseFloat(record[2], 64)
		c, _ := strconv.ParseFloat(record[3], 64)
		d, _ := strconv.ParseFloat(record[4], 64)
		avgPrice := (a + b + c + d) / 4.00
		s := fmt.Sprintf("%f", avgPrice)
		var newEntry = []string{date, s}
		if err := w.Write(newEntry); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}

	}
	w.Flush()
	elapsed := time.Since(start)
	log.Printf("[%s]--> %s", name, elapsed)
}

func main() {
	for _, stock := range stockList {
		processOneFile(stock)
	}

}
