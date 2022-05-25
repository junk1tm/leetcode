import os
import re
from dataclasses import dataclass


@dataclass
class Problem:
    link: str
    number: int
    title: str
    difficulty: str
    file: str
    test: str = ""


problems: list[Problem] = []
total: dict[str:int] = {"Easy": 0, "Medium": 0, "Hard": 0}

for filename in sorted(os.listdir()):
    if not re.fullmatch(r"^[0-9]{4}[a-z_]+.go$", filename):
        continue

    with open(filename, "r") as f:
        COMMENT_PREFIX: str = "// "
        # the first line is always a LeetCode link.
        link: str = f.readline().removeprefix(COMMENT_PREFIX).removesuffix("\n")
        for line in f:
            # the next comment after the link is a header...
            if line.startswith(COMMENT_PREFIX):
                header: str = line.removeprefix(COMMENT_PREFIX).removesuffix("\n")
                break

    # ...that consists of the task's number, name and difficulty.
    match = re.fullmatch(r"^([0-9]{4})\. ([a-zA-Z ()]+) \[(Easy|Medium|Hard)]$", header)

    number: int = int(match.group(1))
    title: str = match.group(2)
    difficulty: str = match.group(3)
    file: str = f"[👀]({filename})"

    p = Problem(link, number, title, difficulty, file)
    problems.append(p)
    total[p.difficulty] += 1

with open("solutions_test.go", "r") as f:
    i: int = 0
    number: int = 0
    for line in f:
        number += 1
        if re.match(r"^// [0-9]{4}\.", line):
            problems[i].test = f"[✅](solutions_test.go#L{number})"
            i += 1

with open("README.md", "w") as f:
    template = f"""\
# LeetCode

[![ci-img]][ci]
[![license-img]][license]

[ci]: https://github.com/junk1tm/leetcode/actions/workflows/go.yml
[ci-img]: https://github.com/junk1tm/leetcode/actions/workflows/go.yml/badge.svg
[license]: https://github.com/junk1tm/leetcode/blob/main/LICENSE
[license-img]: https://img.shields.io/github/license/junk1tm/leetcode

> A collection of my [LeetCode](https://leetcode.com) solutions written in Go

## Progress

### Total

| Easy | Medium | Hard |
|:----:|:------:|:----:|
|{total["Easy"]}|{total["Medium"]}|{total["Hard"]}|

### Solutions

|  #  | Title | Difficulty | Solution |
|:---:|:------|:----------:|:--------:|
"""

    f.write(template)
    for p in problems:
        row = f"|{p.number}|[{p.title}]({p.link})|{p.difficulty}|{p.file} {p.test}|\n"
        f.write(row)
