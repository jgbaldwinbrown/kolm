package kolm

import (
	"strings"
	"fmt"
	"os/exec"
	"iter"
	"regexp"
	"github.com/jgbaldwinbrown/csvh"
	"os"
)

// KstestResult(statistic=0.9021515612265636, pvalue=1.4749825326634794e-14, statistic_location=7.0, statistic_sign=1)

var ksRe = regexp.MustCompile(`KstestResult\(statistic=([^,]*), pvalue=([^,]*), statistic_location=([^,]*), statistic_sign=([^)]*)\)`)

type KolmogorovSmirnovResult struct {
	Statistic float64
	PValue float64
	StatisticLocation float64
	StatisticSign int
}

func ParseKolmogorovSmirnovResult(s string) (KolmogorovSmirnovResult, error) {
	var k KolmogorovSmirnovResult
	fields := ksRe.FindStringSubmatch(s)
	if fields == nil {
		return k, fmt.Errorf("ParseKS: parsing error; %v", s)
	}
	_, e := csvh.Scan(fields, &k.Statistic, &k.PValue, &k.StatisticLocation, &k.StatisticSign)
	return k, e
}

func KolmogorovSmirnovChi2(data iter.Seq[float64]) (KolmogorovSmirnovResult, error) {
	cmd := exec.Command("kstest.py")
	var b strings.Builder
	cmd.Stdout = &b
	cmd.Stderr = os.Stderr
	inp, e := cmd.StdinPipe()
	if e != nil {
		return KolmogorovSmirnovResult{}, e
	}

	errc := make(chan error, 1)
	go func() {
		for val := range data {
			_, e := fmt.Fprintln(inp, val)
			if e != nil {
				errc <- e
				inp.Close()
				return
			}
		}
		inp.Close()
		errc <- nil
	}()

	e = cmd.Start()
	if e != nil {
		return KolmogorovSmirnovResult{}, e
	}

	res, e := ParseKolmogorovSmirnovResult(b.String())
	if e != nil {
		return KolmogorovSmirnovResult{}, e
	}
	return res, cmd.Wait()
}
