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
    #0     770507 / 1000000     0
    #1     267641 / 347357      3e-12
    #2     235225 / 305286      7e-12
    #3      32416 / 42071       7e-11
    #4       8313 / 10789       2e-09
    #5       7477 / 9704        7e-09
    #6        836 / 1085        9e-08
    #7        789 / 1024        8e-07 Denominator is power of two
    #8         47 / 61          2e-05
    #9         37 / 48          0.0003
    #10         10 / 13          0.001
    #11          7 / 9           0.007
    
    List approximated fractions using the Phil's implementation
    #0     770507 / 1000000     0
    #1     267641 / 347357      3e-12
    #2     235225 / 305286      7e-12
    #3      32416 / 42071       7e-11
    #4       8313 / 10789       2e-09
    #5       7477 / 9704        7e-09
    #6        836 / 1085        9e-08
    #7        789 / 1024        8e-07 Denominator is power of two
    #8         47 / 61          2e-05
    #9         37 / 48          0.0003
    #10         10 / 13          0.001
    #11          7 / 9           0.007

## Library

Diophantine can also be used as a library.

## Alternative tools

Some online tools to approximate fractions:

- <https://oss.sheetjs.com/frac/>
- <https://www.johndcook.com/rational_approximation.html>

## ✨ Contributions Welcome

This new project needs your help to become better.
Please propose your enhancements,
or even a further refactoring.

We welcome contributions in many forms,
and there's always plenty to do!

## 🗣️ Feedback

If you have some suggestions, or need a new feature, please contact us:

- [issues tracking](https://github.com/teal-finance/diophantine/issues)
- e-mail at Teal.Finance(à)pm.me
- Twitter at [@TealFinance](https://twitter.com/TealFinance)

Feel free to propose a
[Pull Request](https://github.com/teal-finance/diophantine/pulls),
your contributions are welcome. :wink:

## 🗽 Copyright and license

Copyright (c) 2022 Teal.Finance contributors

Diophantine is free software, and can be redistributed
and/or modified under the terms of the MIT License.
SPDX-License-Identifier: MIT

Diophantine is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty
of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

See the [LICENSE](LICENSE) file (alongside the source files)
or <https://opensource.org/licenses/MIT>.
