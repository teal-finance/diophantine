# Diophantine

[![Go Reference](https://pkg.go.dev/badge/github.com/teal-finance/diophantine.svg "Go documentation for Diophantine")](https://pkg.go.dev/github.com/teal-finance/diophantine)

Diophantine is a pure Go lib and tool to approximate a real number
by a fraction of two rational numbers, also known as rational approximation,
fraction approximation or ratio approximation.

The name is from the [Diophantine approximation](https://wikiless.org/wiki/Diophantine_approximation).

Diophantine use the functions of two other projects:

- [github.com/bodokaiser/approx](https://github.com/bodokaiser/approx)
- [github.com/pnelson/fraction](https://github.com/pnelson/fraction)

## CLI

Example:

    $ go run github.com/teal-finance/diophantine/cmd@latest -list 0.770507 -farey -phil

    List approximated fractions using the Farey series by Bodo
    #0	    770507 / 1000000     0
    #1	    267641 / 347357      3e-12
    #2	    235225 / 305286      7e-12
    #3	     32416 / 42071       7e-11
    #4	      8313 / 10789       2e-09
    #5	      7477 / 9704        7e-09
    #6	       836 / 1085        9e-08
    #7	       789 / 1024        8e-07	Denominator is power of two
    #8	        47 / 61          2e-05
    #9	        37 / 48          0.0003
    #10	        10 / 13          0.001
    #11	         7 / 9           0.007
    
    List approximated fractions using the Phil's implementation
    #0	    770507 / 1000000     0
    #1	    267641 / 347357      3e-12
    #2	    235225 / 305286      7e-12
    #3	     32416 / 42071       7e-11
    #4	      8313 / 10789       2e-09
    #5	      7477 / 9704        7e-09
    #6	       836 / 1085        9e-08
    #7	       789 / 1024        8e-07	Denominator is power of two
    #8	        47 / 61          2e-05
    #9	        37 / 48          0.0003
    #10	        10 / 13          0.001
    #11	         7 / 9           0.007

## Library

Diophantine can also be used as a library.

## Feedback

If you have some suggestions, or need a new feature,
please contact us, using the [issue tracking](https://github.com/teal-finance/diophantine/issues),
or at Teal.Finance[at]pm.me or at [@TealFinance](https://twitter.com/TealFinance).
Feel free to propose a [Pull Request](https://github.com/teal-finance/diophantine/pulls).

## See also

Some online tools to approximate fractions:

- <https://oss.sheetjs.com/frac/>
- <https://www.johndcook.com/rational_approximation.html>
