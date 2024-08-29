#!/usr/bin/env python3

from scipy import stats
from typing import List
import sys

# def kstest_chi2(data: List[float]) -> stats.KstestResult:

def kstest_chi2(data: List[float]):
	return stats.kstest(data, stats.chi2(len(data) - 1).cdf)

def main():
	data = [float(line) for line in sys.stdin]
	res = kstest_chi2(data)
	print(res)

if __name__ == "__main__":
	main()

# x = stats.norm.rvs(size=100, loc=0.5, random_state=rng)
# 
# stats.kstest(x, stats.norm.cdf, alternative='less')
# KstestResult(statistic=0.17482387821055168,
#              pvalue=0.001913921057766743,
#              statistic_location=0.3713830565352756,
#              statistic_sign=-1)
